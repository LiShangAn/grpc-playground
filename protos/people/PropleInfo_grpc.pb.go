// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package people

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

// PeopleInfoServiceClient is the client API for PeopleInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeopleInfoServiceClient interface {
	PeopleInfo(ctx context.Context, opts ...grpc.CallOption) (PeopleInfoService_PeopleInfoClient, error)
}

type peopleInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeopleInfoServiceClient(cc grpc.ClientConnInterface) PeopleInfoServiceClient {
	return &peopleInfoServiceClient{cc}
}

func (c *peopleInfoServiceClient) PeopleInfo(ctx context.Context, opts ...grpc.CallOption) (PeopleInfoService_PeopleInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &PeopleInfoService_ServiceDesc.Streams[0], "/people.PeopleInfoService/PeopleInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &peopleInfoServicePeopleInfoClient{stream}
	return x, nil
}

type PeopleInfoService_PeopleInfoClient interface {
	Send(*PeopleInfoRequest) error
	CloseAndRecv() (*PeopleInfoResponse, error)
	grpc.ClientStream
}

type peopleInfoServicePeopleInfoClient struct {
	grpc.ClientStream
}

func (x *peopleInfoServicePeopleInfoClient) Send(m *PeopleInfoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peopleInfoServicePeopleInfoClient) CloseAndRecv() (*PeopleInfoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PeopleInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PeopleInfoServiceServer is the server API for PeopleInfoService service.
// All implementations must embed UnimplementedPeopleInfoServiceServer
// for forward compatibility
type PeopleInfoServiceServer interface {
	PeopleInfo(PeopleInfoService_PeopleInfoServer) error
	mustEmbedUnimplementedPeopleInfoServiceServer()
}

// UnimplementedPeopleInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPeopleInfoServiceServer struct {
}

func (UnimplementedPeopleInfoServiceServer) PeopleInfo(PeopleInfoService_PeopleInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method PeopleInfo not implemented")
}
func (UnimplementedPeopleInfoServiceServer) mustEmbedUnimplementedPeopleInfoServiceServer() {}

// UnsafePeopleInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeopleInfoServiceServer will
// result in compilation errors.
type UnsafePeopleInfoServiceServer interface {
	mustEmbedUnimplementedPeopleInfoServiceServer()
}

func RegisterPeopleInfoServiceServer(s grpc.ServiceRegistrar, srv PeopleInfoServiceServer) {
	s.RegisterService(&PeopleInfoService_ServiceDesc, srv)
}

func _PeopleInfoService_PeopleInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeopleInfoServiceServer).PeopleInfo(&peopleInfoServicePeopleInfoServer{stream})
}

type PeopleInfoService_PeopleInfoServer interface {
	SendAndClose(*PeopleInfoResponse) error
	Recv() (*PeopleInfoRequest, error)
	grpc.ServerStream
}

type peopleInfoServicePeopleInfoServer struct {
	grpc.ServerStream
}

func (x *peopleInfoServicePeopleInfoServer) SendAndClose(m *PeopleInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peopleInfoServicePeopleInfoServer) Recv() (*PeopleInfoRequest, error) {
	m := new(PeopleInfoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PeopleInfoService_ServiceDesc is the grpc.ServiceDesc for PeopleInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeopleInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "people.PeopleInfoService",
	HandlerType: (*PeopleInfoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PeopleInfo",
			Handler:       _PeopleInfoService_PeopleInfo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/PropleInfo.proto",
}
