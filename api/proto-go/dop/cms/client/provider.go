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

// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: cicdcms.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/cms/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.dop.cms",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType              = reflect.TypeOf((*Client)(nil)).Elem()
	cicdcmsServiceClientType = reflect.TypeOf((*pb.CICDCmsServiceClient)(nil)).Elem()
	cicdcmsServiceServerType = reflect.TypeOf((*pb.CICDCmsServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.dop.cms-client":
		return p.client
	case "erda.dop.cms.CICDCmsService":
		return &cicdcmsServiceWrapper{client: p.client.CICDCmsService(), opts: opts}
	case "erda.dop.cms.CICDCmsService.client":
		return p.client.CICDCmsService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case cicdcmsServiceClientType:
		return p.client.CICDCmsService()
	case cicdcmsServiceServerType:
		return &cicdcmsServiceWrapper{client: p.client.CICDCmsService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.cms-client", &servicehub.Spec{
		Services: []string{
			"erda.dop.cms.CICDCmsService",
			"erda.dop.cms.CICDCmsService.client",
			"erda.dop.cms-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			cicdcmsServiceClientType,
			// server types
			cicdcmsServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
