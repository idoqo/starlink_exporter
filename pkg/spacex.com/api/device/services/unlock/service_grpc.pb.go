// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: spacex/api/device/services/unlock/service.proto

package unlock

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

// UnlockServiceClient is the client API for UnlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnlockServiceClient interface {
	StartUnlock(ctx context.Context, in *StartUnlockRequest, opts ...grpc.CallOption) (*StartUnlockResponse, error)
	FinishUnlock(ctx context.Context, in *FinishUnlockRequest, opts ...grpc.CallOption) (*FinishUnlockResponse, error)
}

type unlockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnlockServiceClient(cc grpc.ClientConnInterface) UnlockServiceClient {
	return &unlockServiceClient{cc}
}

func (c *unlockServiceClient) StartUnlock(ctx context.Context, in *StartUnlockRequest, opts ...grpc.CallOption) (*StartUnlockResponse, error) {
	out := new(StartUnlockResponse)
	err := c.cc.Invoke(ctx, "/SpaceX.API.Device.Services.Unlock.UnlockService/StartUnlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unlockServiceClient) FinishUnlock(ctx context.Context, in *FinishUnlockRequest, opts ...grpc.CallOption) (*FinishUnlockResponse, error) {
	out := new(FinishUnlockResponse)
	err := c.cc.Invoke(ctx, "/SpaceX.API.Device.Services.Unlock.UnlockService/FinishUnlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnlockServiceServer is the server API for UnlockService service.
// All implementations must embed UnimplementedUnlockServiceServer
// for forward compatibility
type UnlockServiceServer interface {
	StartUnlock(context.Context, *StartUnlockRequest) (*StartUnlockResponse, error)
	FinishUnlock(context.Context, *FinishUnlockRequest) (*FinishUnlockResponse, error)
	mustEmbedUnimplementedUnlockServiceServer()
}

// UnimplementedUnlockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUnlockServiceServer struct {
}

func (UnimplementedUnlockServiceServer) StartUnlock(context.Context, *StartUnlockRequest) (*StartUnlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartUnlock not implemented")
}
func (UnimplementedUnlockServiceServer) FinishUnlock(context.Context, *FinishUnlockRequest) (*FinishUnlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishUnlock not implemented")
}
func (UnimplementedUnlockServiceServer) mustEmbedUnimplementedUnlockServiceServer() {}

// UnsafeUnlockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnlockServiceServer will
// result in compilation errors.
type UnsafeUnlockServiceServer interface {
	mustEmbedUnimplementedUnlockServiceServer()
}

func RegisterUnlockServiceServer(s grpc.ServiceRegistrar, srv UnlockServiceServer) {
	s.RegisterService(&UnlockService_ServiceDesc, srv)
}

func _UnlockService_StartUnlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartUnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnlockServiceServer).StartUnlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SpaceX.API.Device.Services.Unlock.UnlockService/StartUnlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnlockServiceServer).StartUnlock(ctx, req.(*StartUnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnlockService_FinishUnlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishUnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnlockServiceServer).FinishUnlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SpaceX.API.Device.Services.Unlock.UnlockService/FinishUnlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnlockServiceServer).FinishUnlock(ctx, req.(*FinishUnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UnlockService_ServiceDesc is the grpc.ServiceDesc for UnlockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnlockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SpaceX.API.Device.Services.Unlock.UnlockService",
	HandlerType: (*UnlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartUnlock",
			Handler:    _UnlockService_StartUnlock_Handler,
		},
		{
			MethodName: "FinishUnlock",
			Handler:    _UnlockService_FinishUnlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacex/api/device/services/unlock/service.proto",
}
