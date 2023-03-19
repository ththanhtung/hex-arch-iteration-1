// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: arithmetic_srv.proto

package pb

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

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	GetAdd(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error)
	GetSub(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error)
	GetMul(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error)
	GetDiv(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error)
}

type arithmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticServiceClient(cc grpc.ClientConnInterface) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) GetAdd(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetSub(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetSub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetMul(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetMul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) GetDiv(ctx context.Context, in *OperationParams, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArithmeticService/GetDiv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
// All implementations should embed UnimplementedArithmeticServiceServer
// for forward compatibility
type ArithmeticServiceServer interface {
	GetAdd(context.Context, *OperationParams) (*Answer, error)
	GetSub(context.Context, *OperationParams) (*Answer, error)
	GetMul(context.Context, *OperationParams) (*Answer, error)
	GetDiv(context.Context, *OperationParams) (*Answer, error)
}

// UnimplementedArithmeticServiceServer should be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (UnimplementedArithmeticServiceServer) GetAdd(context.Context, *OperationParams) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdd not implemented")
}
func (UnimplementedArithmeticServiceServer) GetSub(context.Context, *OperationParams) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSub not implemented")
}
func (UnimplementedArithmeticServiceServer) GetMul(context.Context, *OperationParams) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMul not implemented")
}
func (UnimplementedArithmeticServiceServer) GetDiv(context.Context, *OperationParams) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiv not implemented")
}

// UnsafeArithmeticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticServiceServer will
// result in compilation errors.
type UnsafeArithmeticServiceServer interface {
	mustEmbedUnimplementedArithmeticServiceServer()
}

func RegisterArithmeticServiceServer(s grpc.ServiceRegistrar, srv ArithmeticServiceServer) {
	s.RegisterService(&ArithmeticService_ServiceDesc, srv)
}

func _ArithmeticService_GetAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetAdd(ctx, req.(*OperationParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetSub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetSub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetSub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetSub(ctx, req.(*OperationParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetMul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetMul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetMul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetMul(ctx, req.(*OperationParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_GetDiv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).GetDiv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArithmeticService/GetDiv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).GetDiv(ctx, req.(*OperationParams))
	}
	return interceptor(ctx, in, info, handler)
}

// ArithmeticService_ServiceDesc is the grpc.ServiceDesc for ArithmeticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithmeticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdd",
			Handler:    _ArithmeticService_GetAdd_Handler,
		},
		{
			MethodName: "GetSub",
			Handler:    _ArithmeticService_GetSub_Handler,
		},
		{
			MethodName: "GetMul",
			Handler:    _ArithmeticService_GetMul_Handler,
		},
		{
			MethodName: "GetDiv",
			Handler:    _ArithmeticService_GetDiv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic_srv.proto",
}
