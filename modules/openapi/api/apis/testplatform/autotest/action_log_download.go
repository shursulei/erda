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

package autotest

import "github.com/erda-project/erda/modules/openapi/api/apis"

var ACTION_LOG_DOWNLOAD = apis.ApiSpec{
	Path:        "/api/autotests/filetree/<inode>/actions/download-action-log",
	BackendPath: "/api/logs/actions/download",
	Host:        "monitor.marathon.l4lb.thisdcos.directory:7096",
	Scheme:      "http",
	Method:      "GET",
	CheckLogin:  true,
	CheckToken:  true,
	ChunkAPI:    true,
	Doc:         "下载自动化测试节点 action 日志",
}
