// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform.proto

package com_roleypoly_platform

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

func init() { proto.RegisterFile("platform.proto", fileDescriptor_918f3d50bfb447e4) }

var fileDescriptor_918f3d50bfb447e4 = []byte{
	// 70 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc8, 0x49, 0x2c,
	0x49, 0xcb, 0x2f, 0xca, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4b, 0xce, 0xcf, 0xd5,
	0x2b, 0xca, 0xcf, 0x49, 0xad, 0x2c, 0xc8, 0xcf, 0xa9, 0xd4, 0x83, 0xc9, 0x1a, 0x71, 0x71, 0x71,
	0x04, 0x40, 0xd9, 0x49, 0x6c, 0x60, 0xa5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0xca,
	0x63, 0xb1, 0x3c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformClient interface {
}

type platformClient struct {
	cc *grpc.ClientConn
}

func NewPlatformClient(cc *grpc.ClientConn) PlatformClient {
	return &platformClient{cc}
}

// PlatformServer is the server API for Platform service.
type PlatformServer interface {
}

// UnimplementedPlatformServer can be embedded to have forward compatible implementations.
type UnimplementedPlatformServer struct {
}

func RegisterPlatformServer(s *grpc.Server, srv PlatformServer) {
	s.RegisterService(&_Platform_serviceDesc, srv)
}

var _Platform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.roleypoly.platform.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "platform.proto",
}
