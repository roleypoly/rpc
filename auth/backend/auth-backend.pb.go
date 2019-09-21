// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth-backend.proto

package com_roleypoly_auth_backend

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserSlug struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSlug) Reset()         { *m = UserSlug{} }
func (m *UserSlug) String() string { return proto.CompactTextString(m) }
func (*UserSlug) ProtoMessage()    {}
func (*UserSlug) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7c305325d5462b7, []int{0}
}

func (m *UserSlug) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSlug.Unmarshal(m, b)
}
func (m *UserSlug) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSlug.Marshal(b, m, deterministic)
}
func (m *UserSlug) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSlug.Merge(m, src)
}
func (m *UserSlug) XXX_Size() int {
	return xxx_messageInfo_UserSlug.Size(m)
}
func (m *UserSlug) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSlug.DiscardUnknown(m)
}

var xxx_messageInfo_UserSlug proto.InternalMessageInfo

func (m *UserSlug) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type AuthChallenge struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	MagicUrl             string   `protobuf:"bytes,2,opt,name=magicUrl,proto3" json:"magicUrl,omitempty"`
	MagicWords           string   `protobuf:"bytes,3,opt,name=magicWords,proto3" json:"magicWords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthChallenge) Reset()         { *m = AuthChallenge{} }
func (m *AuthChallenge) String() string { return proto.CompactTextString(m) }
func (*AuthChallenge) ProtoMessage()    {}
func (*AuthChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7c305325d5462b7, []int{1}
}

func (m *AuthChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthChallenge.Unmarshal(m, b)
}
func (m *AuthChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthChallenge.Marshal(b, m, deterministic)
}
func (m *AuthChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthChallenge.Merge(m, src)
}
func (m *AuthChallenge) XXX_Size() int {
	return xxx_messageInfo_AuthChallenge.Size(m)
}
func (m *AuthChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_AuthChallenge proto.InternalMessageInfo

func (m *AuthChallenge) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AuthChallenge) GetMagicUrl() string {
	if m != nil {
		return m.MagicUrl
	}
	return ""
}

func (m *AuthChallenge) GetMagicWords() string {
	if m != nil {
		return m.MagicWords
	}
	return ""
}

func init() {
	proto.RegisterType((*UserSlug)(nil), "com.roleypoly.auth.backend.UserSlug")
	proto.RegisterType((*AuthChallenge)(nil), "com.roleypoly.auth.backend.AuthChallenge")
}

func init() { proto.RegisterFile("auth-backend.proto", fileDescriptor_e7c305325d5462b7) }

var fileDescriptor_e7c305325d5462b7 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x4d, 0x4a, 0x4c, 0xce, 0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x4a, 0xce, 0xcf, 0xd5, 0x2b, 0xca, 0xcf, 0x49, 0xad, 0x2c, 0xc8, 0xcf, 0xa9, 0xd4, 0x03, 0xa9,
	0xd0, 0x83, 0xaa, 0x50, 0x52, 0xe2, 0xe2, 0x08, 0x2d, 0x4e, 0x2d, 0x0a, 0xce, 0x29, 0x4d, 0x17,
	0x12, 0xe3, 0x62, 0x2b, 0x2d, 0x4e, 0x2d, 0xf2, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0xf2, 0x94, 0x92, 0xb9, 0x78, 0x1d, 0x4b, 0x4b, 0x32, 0x9c, 0x33, 0x12, 0x73, 0x72, 0x52,
	0xf3, 0xd2, 0x53, 0x71, 0x29, 0x14, 0x92, 0xe2, 0xe2, 0xc8, 0x4d, 0x4c, 0xcf, 0x4c, 0x0e, 0x2d,
	0xca, 0x91, 0x60, 0x02, 0xcb, 0xc0, 0xf9, 0x42, 0x72, 0x5c, 0x5c, 0x60, 0x76, 0x78, 0x7e, 0x51,
	0x4a, 0xb1, 0x04, 0x33, 0x58, 0x16, 0x49, 0xc4, 0xa8, 0x9c, 0x8b, 0x1b, 0x64, 0x89, 0x13, 0xc4,
	0x5d, 0x42, 0x19, 0x5c, 0xc2, 0xee, 0xa9, 0x25, 0xc1, 0xa9, 0xc5, 0xc5, 0x99, 0xf9, 0x79, 0x08,
	0x9b, 0x55, 0xf4, 0x70, 0xfb, 0x45, 0x0f, 0xe6, 0x11, 0x29, 0x4d, 0x7c, 0xaa, 0x50, 0xbc, 0xa2,
	0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x24, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xc9, 0x7e,
	0x2e, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthBackendClient is the client API for AuthBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthBackendClient interface {
	GetSessionChallenge(ctx context.Context, in *UserSlug, opts ...grpc.CallOption) (*AuthChallenge, error)
}

type authBackendClient struct {
	cc *grpc.ClientConn
}

func NewAuthBackendClient(cc *grpc.ClientConn) AuthBackendClient {
	return &authBackendClient{cc}
}

func (c *authBackendClient) GetSessionChallenge(ctx context.Context, in *UserSlug, opts ...grpc.CallOption) (*AuthChallenge, error) {
	out := new(AuthChallenge)
	err := c.cc.Invoke(ctx, "/com.roleypoly.auth.backend.AuthBackend/GetSessionChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthBackendServer is the server API for AuthBackend service.
type AuthBackendServer interface {
	GetSessionChallenge(context.Context, *UserSlug) (*AuthChallenge, error)
}

// UnimplementedAuthBackendServer can be embedded to have forward compatible implementations.
type UnimplementedAuthBackendServer struct {
}

func (*UnimplementedAuthBackendServer) GetSessionChallenge(ctx context.Context, req *UserSlug) (*AuthChallenge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionChallenge not implemented")
}

func RegisterAuthBackendServer(s *grpc.Server, srv AuthBackendServer) {
	s.RegisterService(&_AuthBackend_serviceDesc, srv)
}

func _AuthBackend_GetSessionChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSlug)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthBackendServer).GetSessionChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.roleypoly.auth.backend.AuthBackend/GetSessionChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthBackendServer).GetSessionChallenge(ctx, req.(*UserSlug))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.roleypoly.auth.backend.AuthBackend",
	HandlerType: (*AuthBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSessionChallenge",
			Handler:    _AuthBackend_GetSessionChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-backend.proto",
}