// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: buildsafe/v1/search.proto

package bsfv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	// ListPackages returns a list of packages available. Mostly meant for browsing if a package is available or not
	ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error)
	// FetchPackage returns a package based on the name
	FetchPackages(ctx context.Context, in *FetchPackagesRequest, opts ...grpc.CallOption) (*FetchPackagesResponse, error)
	// FetchPackageVersion returns a package based on the name and version
	FetchPackageVersion(ctx context.Context, in *FetchPackageVersionRequest, opts ...grpc.CallOption) (*FetchPackageVersionResponse, error)
	// FetchVulnerabilities returns a list of vulnerabilities for a package based on the name and version
	FetchVulnerabilities(ctx context.Context, in *FetchVulnerabilitiesRequest, opts ...grpc.CallOption) (*FetchVulnerabilitiesResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error) {
	out := new(ListPackagesResponse)
	err := c.cc.Invoke(ctx, "/buildsafe.v1.SearchService/ListPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) FetchPackages(ctx context.Context, in *FetchPackagesRequest, opts ...grpc.CallOption) (*FetchPackagesResponse, error) {
	out := new(FetchPackagesResponse)
	err := c.cc.Invoke(ctx, "/buildsafe.v1.SearchService/FetchPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) FetchPackageVersion(ctx context.Context, in *FetchPackageVersionRequest, opts ...grpc.CallOption) (*FetchPackageVersionResponse, error) {
	out := new(FetchPackageVersionResponse)
	err := c.cc.Invoke(ctx, "/buildsafe.v1.SearchService/FetchPackageVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) FetchVulnerabilities(ctx context.Context, in *FetchVulnerabilitiesRequest, opts ...grpc.CallOption) (*FetchVulnerabilitiesResponse, error) {
	out := new(FetchVulnerabilitiesResponse)
	err := c.cc.Invoke(ctx, "/buildsafe.v1.SearchService/FetchVulnerabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	// ListPackages returns a list of packages available. Mostly meant for browsing if a package is available or not
	ListPackages(context.Context, *ListPackagesRequest) (*ListPackagesResponse, error)
	// FetchPackage returns a package based on the name
	FetchPackages(context.Context, *FetchPackagesRequest) (*FetchPackagesResponse, error)
	// FetchPackageVersion returns a package based on the name and version
	FetchPackageVersion(context.Context, *FetchPackageVersionRequest) (*FetchPackageVersionResponse, error)
	// FetchVulnerabilities returns a list of vulnerabilities for a package based on the name and version
	FetchVulnerabilities(context.Context, *FetchVulnerabilitiesRequest) (*FetchVulnerabilitiesResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) ListPackages(context.Context, *ListPackagesRequest) (*ListPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPackages not implemented")
}
func (UnimplementedSearchServiceServer) FetchPackages(context.Context, *FetchPackagesRequest) (*FetchPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPackages not implemented")
}
func (UnimplementedSearchServiceServer) FetchPackageVersion(context.Context, *FetchPackageVersionRequest) (*FetchPackageVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPackageVersion not implemented")
}
func (UnimplementedSearchServiceServer) FetchVulnerabilities(context.Context, *FetchVulnerabilitiesRequest) (*FetchVulnerabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchVulnerabilities not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_ListPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ListPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildsafe.v1.SearchService/ListPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ListPackages(ctx, req.(*ListPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_FetchPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).FetchPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildsafe.v1.SearchService/FetchPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).FetchPackages(ctx, req.(*FetchPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_FetchPackageVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchPackageVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).FetchPackageVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildsafe.v1.SearchService/FetchPackageVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).FetchPackageVersion(ctx, req.(*FetchPackageVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_FetchVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchVulnerabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).FetchVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildsafe.v1.SearchService/FetchVulnerabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).FetchVulnerabilities(ctx, req.(*FetchVulnerabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildsafe.v1.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPackages",
			Handler:    _SearchService_ListPackages_Handler,
		},
		{
			MethodName: "FetchPackages",
			Handler:    _SearchService_FetchPackages_Handler,
		},
		{
			MethodName: "FetchPackageVersion",
			Handler:    _SearchService_FetchPackageVersion_Handler,
		},
		{
			MethodName: "FetchVulnerabilities",
			Handler:    _SearchService_FetchVulnerabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buildsafe/v1/search.proto",
}
