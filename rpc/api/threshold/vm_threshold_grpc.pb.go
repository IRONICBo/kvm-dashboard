// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: vm_threshold.proto

package threshold

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

// RpcDomainThvClient is the client API for RpcDomainThv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcDomainThvClient interface {
	GetDomainThv(ctx context.Context, in *GetDomainThvRequest, opts ...grpc.CallOption) (*GetDomainThvResponse, error)
}

type rpcDomainThvClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcDomainThvClient(cc grpc.ClientConnInterface) RpcDomainThvClient {
	return &rpcDomainThvClient{cc}
}

func (c *rpcDomainThvClient) GetDomainThv(ctx context.Context, in *GetDomainThvRequest, opts ...grpc.CallOption) (*GetDomainThvResponse, error) {
	out := new(GetDomainThvResponse)
	err := c.cc.Invoke(ctx, "/RpcDomainThv/GetDomainThv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcDomainThvServer is the server API for RpcDomainThv service.
// All implementations must embed UnimplementedRpcDomainThvServer
// for forward compatibility
type RpcDomainThvServer interface {
	GetDomainThv(context.Context, *GetDomainThvRequest) (*GetDomainThvResponse, error)
	mustEmbedUnimplementedRpcDomainThvServer()
}

// UnimplementedRpcDomainThvServer must be embedded to have forward compatible implementations.
type UnimplementedRpcDomainThvServer struct {
}

func (UnimplementedRpcDomainThvServer) GetDomainThv(context.Context, *GetDomainThvRequest) (*GetDomainThvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainThv not implemented")
}
func (UnimplementedRpcDomainThvServer) mustEmbedUnimplementedRpcDomainThvServer() {}

// UnsafeRpcDomainThvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcDomainThvServer will
// result in compilation errors.
type UnsafeRpcDomainThvServer interface {
	mustEmbedUnimplementedRpcDomainThvServer()
}

func RegisterRpcDomainThvServer(s grpc.ServiceRegistrar, srv RpcDomainThvServer) {
	s.RegisterService(&RpcDomainThv_ServiceDesc, srv)
}

func _RpcDomainThv_GetDomainThv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainThvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDomainThvServer).GetDomainThv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcDomainThv/GetDomainThv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDomainThvServer).GetDomainThv(ctx, req.(*GetDomainThvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcDomainThv_ServiceDesc is the grpc.ServiceDesc for RpcDomainThv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcDomainThv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RpcDomainThv",
	HandlerType: (*RpcDomainThvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDomainThv",
			Handler:    _RpcDomainThv_GetDomainThv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vm_threshold.proto",
}

// RpcSetDomainThvClient is the client API for RpcSetDomainThv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcSetDomainThvClient interface {
	SetDomainThv(ctx context.Context, in *SetDomainRequest, opts ...grpc.CallOption) (*SetDomainResponse, error)
}

type rpcSetDomainThvClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcSetDomainThvClient(cc grpc.ClientConnInterface) RpcSetDomainThvClient {
	return &rpcSetDomainThvClient{cc}
}

func (c *rpcSetDomainThvClient) SetDomainThv(ctx context.Context, in *SetDomainRequest, opts ...grpc.CallOption) (*SetDomainResponse, error) {
	out := new(SetDomainResponse)
	err := c.cc.Invoke(ctx, "/RpcSetDomainThv/SetDomainThv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcSetDomainThvServer is the server API for RpcSetDomainThv service.
// All implementations must embed UnimplementedRpcSetDomainThvServer
// for forward compatibility
type RpcSetDomainThvServer interface {
	SetDomainThv(context.Context, *SetDomainRequest) (*SetDomainResponse, error)
	mustEmbedUnimplementedRpcSetDomainThvServer()
}

// UnimplementedRpcSetDomainThvServer must be embedded to have forward compatible implementations.
type UnimplementedRpcSetDomainThvServer struct {
}

func (UnimplementedRpcSetDomainThvServer) SetDomainThv(context.Context, *SetDomainRequest) (*SetDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDomainThv not implemented")
}
func (UnimplementedRpcSetDomainThvServer) mustEmbedUnimplementedRpcSetDomainThvServer() {}

// UnsafeRpcSetDomainThvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcSetDomainThvServer will
// result in compilation errors.
type UnsafeRpcSetDomainThvServer interface {
	mustEmbedUnimplementedRpcSetDomainThvServer()
}

func RegisterRpcSetDomainThvServer(s grpc.ServiceRegistrar, srv RpcSetDomainThvServer) {
	s.RegisterService(&RpcSetDomainThv_ServiceDesc, srv)
}

func _RpcSetDomainThv_SetDomainThv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcSetDomainThvServer).SetDomainThv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcSetDomainThv/SetDomainThv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcSetDomainThvServer).SetDomainThv(ctx, req.(*SetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcSetDomainThv_ServiceDesc is the grpc.ServiceDesc for RpcSetDomainThv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcSetDomainThv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RpcSetDomainThv",
	HandlerType: (*RpcSetDomainThvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDomainThv",
			Handler:    _RpcSetDomainThv_SetDomainThv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vm_threshold.proto",
}