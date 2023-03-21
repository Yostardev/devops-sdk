// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: artifact/v1/binary.proto

package artifact

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BinaryServiceClient is the client API for BinaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BinaryServiceClient interface {
	List(ctx context.Context, in *BinaryListRequest, opts ...grpc.CallOption) (*BinaryListResponse, error)
	Create(ctx context.Context, in *BinaryCreateRequest, opts ...grpc.CallOption) (*BinaryResponse, error)
	GetObjectMeta(ctx context.Context, in *BinaryObjectMetaRequest, opts ...grpc.CallOption) (*BinaryObjectMetaResponse, error)
	DownloadFile(ctx context.Context, in *BinaryDownloadRequest, opts ...grpc.CallOption) (BinaryService_DownloadFileClient, error)
	DeleteCollection(ctx context.Context, in *BinaryObjectDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type binaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBinaryServiceClient(cc grpc.ClientConnInterface) BinaryServiceClient {
	return &binaryServiceClient{cc}
}

func (c *binaryServiceClient) List(ctx context.Context, in *BinaryListRequest, opts ...grpc.CallOption) (*BinaryListResponse, error) {
	out := new(BinaryListResponse)
	err := c.cc.Invoke(ctx, "/artifact.BinaryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryServiceClient) Create(ctx context.Context, in *BinaryCreateRequest, opts ...grpc.CallOption) (*BinaryResponse, error) {
	out := new(BinaryResponse)
	err := c.cc.Invoke(ctx, "/artifact.BinaryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryServiceClient) GetObjectMeta(ctx context.Context, in *BinaryObjectMetaRequest, opts ...grpc.CallOption) (*BinaryObjectMetaResponse, error) {
	out := new(BinaryObjectMetaResponse)
	err := c.cc.Invoke(ctx, "/artifact.BinaryService/GetObjectMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryServiceClient) DownloadFile(ctx context.Context, in *BinaryDownloadRequest, opts ...grpc.CallOption) (BinaryService_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &BinaryService_ServiceDesc.Streams[0], "/artifact.BinaryService/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &binaryServiceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BinaryService_DownloadFileClient interface {
	Recv() (*BinaryDownloadResponse, error)
	grpc.ClientStream
}

type binaryServiceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *binaryServiceDownloadFileClient) Recv() (*BinaryDownloadResponse, error) {
	m := new(BinaryDownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *binaryServiceClient) DeleteCollection(ctx context.Context, in *BinaryObjectDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/artifact.BinaryService/DeleteCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinaryServiceServer is the server API for BinaryService service.
// All implementations must embed UnimplementedBinaryServiceServer
// for forward compatibility
type BinaryServiceServer interface {
	List(context.Context, *BinaryListRequest) (*BinaryListResponse, error)
	Create(context.Context, *BinaryCreateRequest) (*BinaryResponse, error)
	GetObjectMeta(context.Context, *BinaryObjectMetaRequest) (*BinaryObjectMetaResponse, error)
	DownloadFile(*BinaryDownloadRequest, BinaryService_DownloadFileServer) error
	DeleteCollection(context.Context, *BinaryObjectDeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBinaryServiceServer()
}

// UnimplementedBinaryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBinaryServiceServer struct {
}

func (UnimplementedBinaryServiceServer) List(context.Context, *BinaryListRequest) (*BinaryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBinaryServiceServer) Create(context.Context, *BinaryCreateRequest) (*BinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBinaryServiceServer) GetObjectMeta(context.Context, *BinaryObjectMetaRequest) (*BinaryObjectMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectMeta not implemented")
}
func (UnimplementedBinaryServiceServer) DownloadFile(*BinaryDownloadRequest, BinaryService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedBinaryServiceServer) DeleteCollection(context.Context, *BinaryObjectDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollection not implemented")
}
func (UnimplementedBinaryServiceServer) mustEmbedUnimplementedBinaryServiceServer() {}

// UnsafeBinaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BinaryServiceServer will
// result in compilation errors.
type UnsafeBinaryServiceServer interface {
	mustEmbedUnimplementedBinaryServiceServer()
}

func RegisterBinaryServiceServer(s grpc.ServiceRegistrar, srv BinaryServiceServer) {
	s.RegisterService(&BinaryService_ServiceDesc, srv)
}

func _BinaryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifact.BinaryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryServiceServer).List(ctx, req.(*BinaryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifact.BinaryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryServiceServer).Create(ctx, req.(*BinaryCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryService_GetObjectMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryObjectMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryServiceServer).GetObjectMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifact.BinaryService/GetObjectMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryServiceServer).GetObjectMeta(ctx, req.(*BinaryObjectMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BinaryDownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BinaryServiceServer).DownloadFile(m, &binaryServiceDownloadFileServer{stream})
}

type BinaryService_DownloadFileServer interface {
	Send(*BinaryDownloadResponse) error
	grpc.ServerStream
}

type binaryServiceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *binaryServiceDownloadFileServer) Send(m *BinaryDownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BinaryService_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryObjectDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryServiceServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifact.BinaryService/DeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryServiceServer).DeleteCollection(ctx, req.(*BinaryObjectDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BinaryService_ServiceDesc is the grpc.ServiceDesc for BinaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BinaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "artifact.BinaryService",
	HandlerType: (*BinaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BinaryService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BinaryService_Create_Handler,
		},
		{
			MethodName: "GetObjectMeta",
			Handler:    _BinaryService_GetObjectMeta_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _BinaryService_DeleteCollection_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadFile",
			Handler:       _BinaryService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "artifact/v1/binary.proto",
}