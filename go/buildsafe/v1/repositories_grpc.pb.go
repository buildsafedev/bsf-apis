// (-- api-linter: core::0191::java-package=disabled
// (-- api-linter: core::0191::java-outer-classname=disabled
// (-- api-linter: core::0191::java-multiple-files=disabled
//     aip.dev/not-precedent: We're doing This as this is a Go server --)
// (-- api-linter: core::0191::proto-package=disabled

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: buildsafe/v1/repositories.proto

package bsfv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RepositoryService_ListRepositories_FullMethodName = "/buildsafe.v1.RepositoryService/ListRepositories"
)

// RepositoryServiceClient is the client API for RepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryServiceClient interface {
	ListRepositories(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error)
}

type repositoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryServiceClient(cc grpc.ClientConnInterface) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) ListRepositories(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRepositoriesResponse)
	err := c.cc.Invoke(ctx, RepositoryService_ListRepositories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServiceServer is the server API for RepositoryService service.
// All implementations must embed UnimplementedRepositoryServiceServer
// for forward compatibility.
type RepositoryServiceServer interface {
	ListRepositories(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error)
	mustEmbedUnimplementedRepositoryServiceServer()
}

// UnimplementedRepositoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRepositoryServiceServer struct{}

func (UnimplementedRepositoryServiceServer) ListRepositories(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositories not implemented")
}
func (UnimplementedRepositoryServiceServer) mustEmbedUnimplementedRepositoryServiceServer() {}
func (UnimplementedRepositoryServiceServer) testEmbeddedByValue()                           {}

// UnsafeRepositoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryServiceServer will
// result in compilation errors.
type UnsafeRepositoryServiceServer interface {
	mustEmbedUnimplementedRepositoryServiceServer()
}

func RegisterRepositoryServiceServer(s grpc.ServiceRegistrar, srv RepositoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedRepositoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RepositoryService_ServiceDesc, srv)
}

func _RepositoryService_ListRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).ListRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RepositoryService_ListRepositories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).ListRepositories(ctx, req.(*ListRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepositoryService_ServiceDesc is the grpc.ServiceDesc for RepositoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepositoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildsafe.v1.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRepositories",
			Handler:    _RepositoryService_ListRepositories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buildsafe/v1/repositories.proto",
}