// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leaderworker

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"

	"github.com/erda-project/erda/modules/pipeline/providers/leaderworker/worker"
)

func (p *provider) RegisterCandidateWorker(ctx context.Context, w worker.Worker) error {
	p.Log.Infof("begin register candidate worker, workerID: %s", w.GetID())

	// check leader can be worker
	if p.Election.IsLeader() && !p.Cfg.Leader.IsWorker {
		p.Log.Warnf("leader cannot be worker, skip register candidate worker, workerID: %s", w.GetID())
		return nil
	}

	// check worker fields
	if err := p.checkWorkerFields(w); err != nil {
		p.Log.Errorf("failed to check worker fields, workerID: %s, err: %v", w.GetID(), err)
		return err
	}

	// register worker
	if err := p.registerWorker(ctx, w, worker.Candidate); err != nil {
		return err
	}

	p.lock.Lock()
	wctx, wcancel := context.WithCancel(ctx)
	p.forWorkerUse.myWorkers[w.GetID()] = workerWithCancel{Worker: w, Ctx: wctx, CancelFunc: wcancel}
	p.lock.Unlock()

	// promote to official
	go func() {
		p.promoteCandidateWorker(wctx, w)
		// begin listen after promoted
		go p.workerListenIncomingLogicTask(wctx, w)
		// handle worker delete
		go p.workerListenOfficialWorkerSelfDelete(wctx, w)
	}()

	// heartbeat report
	go p.workerContinueReportHeartbeat(wctx, w)

	return nil
}

func (p *provider) WorkerHookOnWorkerDelete(h WorkerDeleteHandler) {
	p.mustNotStarted()
	p.lock.Lock()
	defer p.lock.Unlock()
	p.forWorkerUse.handlersOnWorkerDelete = append(p.forWorkerUse.handlersOnWorkerDelete, h)
}

func (p *provider) promoteCandidateWorker(ctx context.Context, w worker.Worker) {
	ticker := time.NewTicker(p.Cfg.Worker.Candidate.ThresholdToBecomeOfficial)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.SetType(worker.Official)
			if err := p.registerWorker(ctx, w, w.GetType()); err != nil {
				p.Log.Errorf("failed to promote worker to official(auto retry), workerID: %s, err: %v", w.GetID(), err)
				continue
			}
			p.Log.Infof("promote worker to official, workerID: %s", w.GetID())
			return
		}
	}
}

func (p *provider) registerWorker(ctx context.Context, w worker.Worker, typ worker.Type) error {
	workerBytes, err := w.MarshalJSON()
	if err != nil {
		return err
	}

	// report heartbeat before add
	if err := p.workerOnceReportHeartbeat(ctx, w); err != nil {
		return err
	}

	var ops []clientv3.Op
	switch typ {
	case worker.Candidate:
		ops = append(ops,
			clientv3.OpPut(p.makeEtcdWorkerKey(w.GetID(), worker.Candidate), string(workerBytes)),
		)
	case worker.Official:
		ops = append(ops,
			clientv3.OpDelete(p.makeEtcdWorkerKey(w.GetID(), worker.Candidate)),
			clientv3.OpPut(p.makeEtcdWorkerKey(w.GetID(), worker.Official), string(workerBytes)),
		)
	}

	_, err = p.EtcdClient.Txn(ctx).Then(ops...).Commit()
	if err != nil {
		p.Log.Errorf("failed to notify worker add, workerID: %s, err: %v", w.GetID(), err)
		return err
	}

	return nil
}

func (p *provider) workerListenOfficialWorkerSelfDelete(ctx context.Context, w worker.Worker) {
	p.forWorkerUse.handlersOnWorkerDelete = append(p.forWorkerUse.handlersOnWorkerDelete, p.workerIntervalCleanupOnDelete)
	key := p.makeEtcdWorkerKey(w.GetID(), worker.Official)
	p.ListenPrefix(ctx, key, nil, func(ctx context.Context, event *clientv3.Event) {
		if string(event.Kv.Key) != key {
			return
		}
		ww, ok := p.forWorkerUse.myWorkers[w.GetID()]
		if ok && ww.CancelFunc != nil {
			ww.CancelFunc()
		}
		for _, h := range p.forWorkerUse.handlersOnWorkerDelete {
			h := h
			go h(ctx, Event{Type: event.Type, WorkerID: w.GetID()})
		}
	})
}

func (p *provider) workerIntervalCleanupOnDelete(ctx context.Context, ev Event) {
	// delete heartbeat key
	go func() {
		for {
			_, err := p.EtcdClient.Delete(ctx, p.makeEtcdWorkerHeartbeatKey(ev.WorkerID))
			if err == nil {
				return
			}
			p.Log.Errorf("failed to do worker interval cleanup on delete(auto retry), step: delete heartbeat key, workerID: %s, err: %v", ev.WorkerID, err)
			time.Sleep(p.Cfg.Worker.RetryInterval)
		}
	}()
	// delete dispatch key
	go func() {
		for {
			_, err := p.EtcdClient.Delete(ctx, p.makeEtcdWorkerLogicTaskListenPrefix(ev.WorkerID), clientv3.WithPrefix())
			if err == nil {
				return
			}
			p.Log.Errorf("failed to do worker interval cleanup on delete(auto retry), step: delete dispatch key, workerID: %s, err: %v", ev.WorkerID, err)
			time.Sleep(p.Cfg.Worker.RetryInterval)
		}
	}()
}

func (p *provider) workerListenIncomingLogicTask(ctx context.Context, w worker.Worker) {
	prefix := p.makeEtcdWorkerLogicTaskListenPrefix(w.GetID())
	p.Log.Infof("worker begin listen incoming logic task, workerID: %s", w.GetID())
	defer p.Log.Infof("worker stop listen incoming logic task, workerID: %s", w.GetID())

	p.ListenPrefix(ctx, prefix, func(ctx context.Context, event *clientv3.Event) {
		go func() {
			// key added, do logic
			key := string(event.Kv.Key)
			taskLogicID := p.getWorkerTaskLogicIDFromIncomingKey(w.GetID(), key)
			taskData := event.Kv.Value
			p.Log.Infof("logic task received and begin handle it, workerID: %s, logicTaskID: %s", w.GetID(), taskLogicID)
			taskDoneCh := make(chan struct{})
			go func() {
				w.Handle(ctx, worker.NewLogicTask(taskLogicID, taskData))
				taskDoneCh <- struct{}{}
			}()
			select {
			case <-ctx.Done():
				p.Log.Warnf("task canceled, workerID: %s, logicTaskID: %s", w.GetID(), taskLogicID)
			case <-taskDoneCh:
				p.Log.Infof("task done, workerID: %s, logicTaskID: %s", w.GetID(), taskLogicID)
				// delete task key means task done
				for {
					_, err := p.EtcdClient.Delete(context.Background(), key)
					if err == nil {
						break
					}
					p.Log.Warnf("failed to delete incoming logic task key after done(auto retry), key: %s, logicTaskID: %s, err: %v", key, taskLogicID, err)
					time.Sleep(p.Cfg.Worker.Task.RetryDeleteTaskInterval)
				}
			}
		}()
	},
		nil,
	)
}
