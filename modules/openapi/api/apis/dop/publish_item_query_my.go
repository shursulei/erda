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

package dop

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var PUBLISH_ITEM_QUERY_MY = apis.ApiSpec{
	Path:          "/api/my-publish-items",
	BackendPath:   "/api/my-publish-items",
	Host:          "dop.marathon.l4lb.thisdcos.directory:9527",
	Scheme:        "http",
	Method:        "GET",
	TryCheckLogin: true,
	RequestType:   apistructs.QueryPublishItemRequest{},
	ResponseType:  apistructs.QueryPublishItemResponse{},
	Doc:           "summary: 查询我的发布",
}
