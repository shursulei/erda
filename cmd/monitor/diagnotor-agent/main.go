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

package main

import (
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda/conf"
	"github.com/erda-project/erda/pkg/common"

	// modules
	_ "github.com/erda-project/erda/modules/core/monitor/diagnotor/agent"
	"github.com/erda-project/erda/modules/core/monitor/diagnotor/agent/bootstrap"

	// providers
	_ "github.com/erda-project/erda-infra/providers"
)

func main() {
	bootstrap.Run(func() {
		common.Run(&servicehub.RunOptions{
			ConfigFile: conf.DiagnotorAgentConfigFilePath,
			Content:    conf.DiagnotorAgentDefaultConfig,
		})
	})
}
