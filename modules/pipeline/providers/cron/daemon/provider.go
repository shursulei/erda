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

package daemon

import (
	"context"
	"reflect"
	"strconv"
	"sync"

	"github.com/coreos/etcd/clientv3"

	logs "github.com/erda-project/erda-infra/base/logs"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/etcd"
	"github.com/erda-project/erda-infra/providers/mysqlxorm"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/pipeline/providers/cron/db"
	"github.com/erda-project/erda/modules/pipeline/providers/leaderworker"
	"github.com/erda-project/erda/pkg/cron"
)

type config struct {
}

// +provider
type provider struct {
	Cfg *config
	Log logs.Logger

	ETCD         etcd.Interface // autowired
	EtcdClient   *clientv3.Client
	MySQL        mysqlxorm.Interface    `autowired:"mysql-xorm"`
	LeaderWorker leaderworker.Interface `autowired:"leader-worker"`

	createPipelineFunc CreatePipelineFunc
	bdl                *bundle.Bundle
	dbClient           *db.Client
	crond              *cron.Cron
	mu                 *sync.Mutex
}

func (p *provider) WithPipelineFunc(createPipelineFunc CreatePipelineFunc) {
	p.createPipelineFunc = createPipelineFunc
}

func (p *provider) ReloadCrond(ctx context.Context) ([]string, error) {
	return p.reloadCrond(ctx)
}

func (p *provider) CrondSnapshot() []string {
	return p.crondSnapshot()
}

func (p *provider) AddIntoPipelineCrond(cronID uint64) error {
	if cronID <= 0 {
		return nil
	}
	_, err := p.EtcdClient.Put(context.Background(), etcdCronPrefixAddKey+strconv.FormatUint(cronID, 10), "")
	return err
}

func (p *provider) DeletePipelineCrond(cronID uint64) error {
	if cronID <= 0 {
		return nil
	}
	_, err := p.EtcdClient.Put(context.Background(), etcdCronPrefixDeleteKey+strconv.FormatUint(cronID, 10), "")
	return err
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.bdl = bundle.New(bundle.WithCMP())
	p.dbClient = &db.Client{Interface: p.MySQL}
	p.crond = cron.New()
	p.mu = &sync.Mutex{}
	return nil
}

func (p *provider) Run(ctx context.Context) error {
	p.LeaderWorker.OnLeader(func(ctx context.Context) {
		p.DoCrondAbout(ctx)
	})
	return nil
}

func init() {
	servicehub.Register("cron-daemon", &servicehub.Spec{
		Services:   []string{"cron-daemon"},
		Types:      []reflect.Type{reflect.TypeOf((*Interface)(nil)).Elem()},
		ConfigFunc: func() interface{} { return &config{} },
		Creator:    func() servicehub.Provider { return &provider{} },
	})
}
