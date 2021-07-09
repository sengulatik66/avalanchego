// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gsubnetlookupproto

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

// SubnetLookupClient is the client API for SubnetLookup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubnetLookupClient interface {
	SubnetID(ctx context.Context, in *SubnetIDRequest, opts ...grpc.CallOption) (*SubnetIDResponse, error)
}

type subnetLookupClient struct {
	cc grpc.ClientConnInterface
}

func NewSubnetLookupClient(cc grpc.ClientConnInterface) SubnetLookupClient {
	return &subnetLookupClient{cc}
}

func (c *subnetLookupClient) SubnetID(ctx context.Context, in *SubnetIDRequest, opts ...grpc.CallOption) (*SubnetIDResponse, error) {
	out := new(SubnetIDResponse)
	err := c.cc.Invoke(ctx, "/gsubnetlookupproto.SubnetLookup/SubnetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubnetLookupServer is the server API for SubnetLookup service.
// All implementations must embed UnimplementedSubnetLookupServer
// for forward compatibility
type SubnetLookupServer interface {
	SubnetID(context.Context, *SubnetIDRequest) (*SubnetIDResponse, error)
	mustEmbedUnimplementedSubnetLookupServer()
}

// UnimplementedSubnetLookupServer must be embedded to have forward compatible implementations.
type UnimplementedSubnetLookupServer struct {
}

func (UnimplementedSubnetLookupServer) SubnetID(context.Context, *SubnetIDRequest) (*SubnetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubnetID not implemented")
}
func (UnimplementedSubnetLookupServer) mustEmbedUnimplementedSubnetLookupServer() {}

// UnsafeSubnetLookupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubnetLookupServer will
// result in compilation errors.
type UnsafeSubnetLookupServer interface {
	mustEmbedUnimplementedSubnetLookupServer()
}

func RegisterSubnetLookupServer(s grpc.ServiceRegistrar, srv SubnetLookupServer) {
	s.RegisterService(&SubnetLookup_ServiceDesc, srv)
}

func _SubnetLookup_SubnetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubnetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetLookupServer).SubnetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsubnetlookupproto.SubnetLookup/SubnetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetLookupServer).SubnetID(ctx, req.(*SubnetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubnetLookup_ServiceDesc is the grpc.ServiceDesc for SubnetLookup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubnetLookup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsubnetlookupproto.SubnetLookup",
	HandlerType: (*SubnetLookupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubnetID",
			Handler:    _SubnetLookup_SubnetID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gsubnetlookup.proto",
}
