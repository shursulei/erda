// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: projectpipeline.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/projectpipeline/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ProjectPipelineService projectpipeline.proto
	ProjectPipelineService() pb.ProjectPipelineServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		projectPipelineService: pb.NewProjectPipelineServiceClient(cc),
	}
}

type serviceClients struct {
	projectPipelineService pb.ProjectPipelineServiceClient
}

func (c *serviceClients) ProjectPipelineService() pb.ProjectPipelineServiceClient {
	return c.projectPipelineService
}

type projectPipelineServiceWrapper struct {
	client pb.ProjectPipelineServiceClient
	opts   []grpc1.CallOption
}

func (s *projectPipelineServiceWrapper) Create(ctx context.Context, req *pb.CreateProjectPipelineRequest) (*pb.CreateProjectPipelineResponse, error) {
	return s.client.Create(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *projectPipelineServiceWrapper) ListApp(ctx context.Context, req *pb.ListAppRequest) (*pb.ListAppResponse, error) {
	return s.client.ListApp(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *projectPipelineServiceWrapper) ListPipelineYml(ctx context.Context, req *pb.ListAppPipelineYmlRequest) (*pb.ListAppPipelineYmlResponse, error) {
	return s.client.ListPipelineYml(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *projectPipelineServiceWrapper) CreateNamePreCheck(ctx context.Context, req *pb.CreateProjectPipelineNamePreCheckRequest) (*pb.CreateProjectPipelineNamePreCheckResponse, error) {
	return s.client.CreateNamePreCheck(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *projectPipelineServiceWrapper) CreateSourcePreCheck(ctx context.Context, req *pb.CreateProjectPipelineSourcePreCheckRequest) (*pb.CreateProjectPipelineSourcePreCheckResponse, error) {
	return s.client.CreateSourcePreCheck(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *projectPipelineServiceWrapper) ListPipelineCategory(ctx context.Context, req *pb.ListPipelineCategoryRequest) (*pb.ListPipelineCategoryResponse, error) {
	return s.client.ListPipelineCategory(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *projectPipelineServiceWrapper) Update(ctx context.Context, req *pb.UpdateProjectPipelineRequest) (*pb.UpdateProjectPipelineResponse, error) {
	return s.client.Update(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
