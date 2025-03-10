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

package stackChartFilter

import (
	"context"
	"encoding/json"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/filter"
)

type Filter struct {
	filter.CommonFilter

	State State `json:"state,omitempty"`
}

type State struct {
	Conditions []filter.PropCondition `json:"conditions,omitempty"`
	Values     Values                 `json:"values,omitempty"`
}

type Values struct {
	Type string `json:"type"`
}

const OperationKeyFilter filter.OperationKey = "filter"
const OperationOwnerSelectMe filter.OperationKey = "ownerSelectMe"

func (f *Filter) SetToProtocolComponent(c *cptype.Component) error {
	b, err := json.Marshal(f)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &c)
}

func (f *Filter) InitFromProtocol(ctx context.Context, c *cptype.Component) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, f)
}
