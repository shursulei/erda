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

package odata

import (
	"fmt"
)

// bytes representation of ObservableData for performance
type Raws []*Raw

type Raw struct {
	Data []byte    `json:"data"`
	Meta *Metadata `json:"meta"`
}

func NewRaw(item []byte) *Raw {
	return &Raw{Data: item, Meta: NewMetadata()}
}

func (r *Raw) HandleKeyValuePair(_ func(pairs map[string]interface{}) map[string]interface{}) {
}

func (r *Raw) Pairs() map[string]interface{} { return nil }

func (r *Raw) Name() string {
	return ""
}

func (r *Raw) Metadata() *Metadata {
	return r.Meta
}

func (r *Raw) Clone() ObservableData {
	item := make([]byte, len(r.Data))
	copy(item, r.Data)
	return &Raw{
		Data: item,
		Meta: r.Meta.Clone(),
	}
}

func (r *Raw) Source() interface{} {
	return r.Data
}

func (r *Raw) SourceCompatibility() interface{} {
	return r.Data
}

func (r *Raw) SourceType() SourceType {
	return RawType
}

func (r *Raw) String() string {
	return fmt.Sprintf("raw(%d)", len(r.Data))
}
