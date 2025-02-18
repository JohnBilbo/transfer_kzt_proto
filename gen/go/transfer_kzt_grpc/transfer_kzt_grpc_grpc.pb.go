// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: transfer_kzt_grpc/transfer_kzt_grpc.proto

package transfer_kzt_grpc_v1

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
	TransferKZTService_Transfer_FullMethodName       = "/transfer_kzt_grpc.TransferKZTService/Transfer"
	TransferKZTService_CreateTransfer_FullMethodName = "/transfer_kzt_grpc.TransferKZTService/CreateTransfer"
)

// TransferKZTServiceClient is the client API for TransferKZTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferKZTServiceClient interface {
	Transfer(ctx context.Context, in *RequestTransfer, opts ...grpc.CallOption) (*ResponseTransfer, error)
	CreateTransfer(ctx context.Context, in *RequestCreateTransfer, opts ...grpc.CallOption) (*ResponseCreateTransfer, error)
}

type transferKZTServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferKZTServiceClient(cc grpc.ClientConnInterface) TransferKZTServiceClient {
	return &transferKZTServiceClient{cc}
}

func (c *transferKZTServiceClient) Transfer(ctx context.Context, in *RequestTransfer, opts ...grpc.CallOption) (*ResponseTransfer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseTransfer)
	err := c.cc.Invoke(ctx, TransferKZTService_Transfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferKZTServiceClient) CreateTransfer(ctx context.Context, in *RequestCreateTransfer, opts ...grpc.CallOption) (*ResponseCreateTransfer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseCreateTransfer)
	err := c.cc.Invoke(ctx, TransferKZTService_CreateTransfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferKZTServiceServer is the server API for TransferKZTService service.
// All implementations must embed UnimplementedTransferKZTServiceServer
// for forward compatibility.
type TransferKZTServiceServer interface {
	Transfer(context.Context, *RequestTransfer) (*ResponseTransfer, error)
	CreateTransfer(context.Context, *RequestCreateTransfer) (*ResponseCreateTransfer, error)
	mustEmbedUnimplementedTransferKZTServiceServer()
}

// UnimplementedTransferKZTServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransferKZTServiceServer struct{}

func (UnimplementedTransferKZTServiceServer) Transfer(context.Context, *RequestTransfer) (*ResponseTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedTransferKZTServiceServer) CreateTransfer(context.Context, *RequestCreateTransfer) (*ResponseCreateTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedTransferKZTServiceServer) mustEmbedUnimplementedTransferKZTServiceServer() {}
func (UnimplementedTransferKZTServiceServer) testEmbeddedByValue()                            {}

// UnsafeTransferKZTServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferKZTServiceServer will
// result in compilation errors.
type UnsafeTransferKZTServiceServer interface {
	mustEmbedUnimplementedTransferKZTServiceServer()
}

func RegisterTransferKZTServiceServer(s grpc.ServiceRegistrar, srv TransferKZTServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransferKZTServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransferKZTService_ServiceDesc, srv)
}

func _TransferKZTService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferKZTServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferKZTService_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferKZTServiceServer).Transfer(ctx, req.(*RequestTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferKZTService_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferKZTServiceServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferKZTService_CreateTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferKZTServiceServer).CreateTransfer(ctx, req.(*RequestCreateTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

// TransferKZTService_ServiceDesc is the grpc.ServiceDesc for TransferKZTService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferKZTService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transfer_kzt_grpc.TransferKZTService",
	HandlerType: (*TransferKZTServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _TransferKZTService_Transfer_Handler,
		},
		{
			MethodName: "CreateTransfer",
			Handler:    _TransferKZTService_CreateTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer_kzt_grpc/transfer_kzt_grpc.proto",
}
