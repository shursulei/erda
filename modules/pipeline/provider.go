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

package pipeline

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/base/servicehub"
	_ "github.com/erda-project/erda-infra/providers/etcd"
	"github.com/erda-project/erda-infra/providers/httpserver"
	"github.com/erda-project/erda-proto-go/core/pipeline/cms/pb"
	cronpb "github.com/erda-project/erda-proto-go/core/pipeline/cron/pb"
	_ "github.com/erda-project/erda/modules/pipeline/aop/plugins"
	"github.com/erda-project/erda/modules/pipeline/providers/clusterinfo"
	"github.com/erda-project/erda/modules/pipeline/providers/cron/compensator"
	"github.com/erda-project/erda/modules/pipeline/providers/cron/daemon"
	"github.com/erda-project/erda/modules/pipeline/providers/dbgc"
	_ "github.com/erda-project/erda/modules/pipeline/providers/dispatcher"
	"github.com/erda-project/erda/modules/pipeline/providers/engine"
	"github.com/erda-project/erda/modules/pipeline/providers/leaderworker"
	"github.com/erda-project/erda/modules/pipeline/providers/queuemanager"
	"github.com/erda-project/erda/modules/pipeline/providers/reconciler"
	"github.com/erda-project/erda/modules/pipeline/providers/resourcegc"
	"github.com/erda-project/erda/providers/metrics/report"
)

type provider struct {
	CmsService     pb.CmsServiceServer      `autowired:"erda.core.pipeline.cms.CmsService"`
	MetricReport   report.MetricReport      `autowired:"metric-report-client" optional:"true"`
	Router         httpserver.Router        `autowired:"http-router"`
	CronService    cronpb.CronServiceServer `autowired:"erda.core.pipeline.cron.CronService" required:"true"`
	CronDaemon     daemon.Interface
	CronCompensate compensator.Interface

	Engine       engine.Interface
	QueueManager queuemanager.Interface
	Reconciler   reconciler.Interface
	LeaderWorker leaderworker.Interface
	ClusterInfo  clusterinfo.Interface
	DBGC         dbgc.Interface
	ResourceGC   resourcegc.Interface
}

func (p *provider) Run(ctx context.Context) error {
	logrus.Infof("[alert] starting pipeline instance")
	var err error

	select {
	case <-ctx.Done():
	}
	return err
}

func init() {
	servicehub.Register("pipeline", &servicehub.Spec{
		Services:     []string{"pipeline"},
		Dependencies: []string{"etcd"},
		Creator:      func() servicehub.Provider { return &provider{} },
	})
}
