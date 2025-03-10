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

package monitor

import "github.com/erda-project/erda/modules/openapi/api/apis"

var SPOT_JVM_PROFILER_THREAD_SNAPSHOT_STATE = apis.ApiSpec{
	Path:        "/api/monitor/jvmprofiler/<addonID>/profiling/<profilingID>/thread-snapshots/<snapshotsID>/state",
	BackendPath: "/api/jvmprofiler/<addonID>/profiling/<profilingID>/thread_snapshots/<snapshotsID>/state",
	Host:        "monitor.marathon.l4lb.thisdcos.directory:7096",
	Scheme:      "http",
	Method:      "GET",
	CheckLogin:  true,
	CheckToken:  true,
	Doc:         "summary: 获取线程dump状态",
}
