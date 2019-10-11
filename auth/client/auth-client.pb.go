// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/client/auth-client.proto

package client

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("auth/client/auth-client.proto", fileDescriptor_1833648300932cc9) }

var fileDescriptor_1833648300932cc9 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0x2d, 0xc9,
	0xd0, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x07, 0xb1, 0x75, 0x21, 0x6c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0xc9, 0xe4, 0xfc, 0x5c, 0xbd, 0xa2, 0xfc, 0x9c, 0xd4, 0xca, 0x82, 0xfc,
	0x9c, 0x4a, 0x3d, 0x90, 0x02, 0x3d, 0x88, 0x02, 0x23, 0x1e, 0x2e, 0x2e, 0xc7, 0xd2, 0x92, 0x0c,
	0x67, 0x30, 0xcf, 0x49, 0x2d, 0x4a, 0x25, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0x1f, 0xae, 0x43, 0xbf, 0xa8, 0x20, 0x59, 0x1f, 0xc9, 0x8a, 0x24, 0x36, 0xb0, 0xb9, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xd3, 0xb2, 0xed, 0x78, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClientClient is the client API for AuthClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClientClient interface {
}

type authClientClient struct {
	cc *grpc.ClientConn
}

func NewAuthClientClient(cc *grpc.ClientConn) AuthClientClient {
	return &authClientClient{cc}
}

// AuthClientServer is the server API for AuthClient service.
type AuthClientServer interface {
}

// UnimplementedAuthClientServer can be embedded to have forward compatible implementations.
type UnimplementedAuthClientServer struct {
}

func RegisterAuthClientServer(s *grpc.Server, srv AuthClientServer) {
	s.RegisterService(&_AuthClient_serviceDesc, srv)
}

var _AuthClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.roleypoly.auth.client.AuthClient",
	HandlerType: (*AuthClientServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "auth/client/auth-client.proto",
}
