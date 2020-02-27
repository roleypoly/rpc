// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/backend/auth-backend.proto

package backend

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	shared "github.com/roleypoly/rpc/shared"
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

type SessionQuery struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionQuery) Reset()         { *m = SessionQuery{} }
func (m *SessionQuery) String() string { return proto.CompactTextString(m) }
func (*SessionQuery) ProtoMessage()    {}
func (*SessionQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_5436ad1a948af874, []int{0}
}

func (m *SessionQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionQuery.Unmarshal(m, b)
}
func (m *SessionQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionQuery.Marshal(b, m, deterministic)
}
func (m *SessionQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionQuery.Merge(m, src)
}
func (m *SessionQuery) XXX_Size() int {
	return xxx_messageInfo_SessionQuery.Size(m)
}
func (m *SessionQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SessionQuery proto.InternalMessageInfo

func (m *SessionQuery) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

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
	return fileDescriptor_5436ad1a948af874, []int{1}
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
	return fileDescriptor_5436ad1a948af874, []int{2}
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
	proto.RegisterType((*SessionQuery)(nil), "roleypoly.auth.backend.SessionQuery")
	proto.RegisterType((*UserSlug)(nil), "roleypoly.auth.backend.UserSlug")
	proto.RegisterType((*AuthChallenge)(nil), "roleypoly.auth.backend.AuthChallenge")
}

func init() {
	proto.RegisterFile("auth/backend/auth-backend.proto", fileDescriptor_5436ad1a948af874)
}

var fileDescriptor_5436ad1a948af874 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x37, 0xc5, 0x31, 0xaf, 0xfa, 0x12, 0x75, 0x8c, 0x0a, 0x3a, 0xc2, 0x86, 0xbe, 0x98,
	0x80, 0x7e, 0x02, 0xa7, 0x20, 0x3e, 0x6e, 0x63, 0x08, 0x3e, 0x99, 0x75, 0x97, 0xb6, 0x2c, 0x4b,
	0x4a, 0xfe, 0x3c, 0xf4, 0xeb, 0xf9, 0xc9, 0xa4, 0x6d, 0xb6, 0x55, 0xb0, 0x6f, 0xf7, 0xe4, 0xfc,
	0xc2, 0xe1, 0xdc, 0x0b, 0x77, 0xc2, 0xbb, 0x94, 0xaf, 0x44, 0xbc, 0x41, 0xb5, 0xe6, 0xa5, 0x78,
	0x0c, 0x82, 0xe5, 0x46, 0x3b, 0x4d, 0x06, 0x46, 0x4b, 0x2c, 0x72, 0x2d, 0x0b, 0x56, 0xba, 0x2c,
	0xb8, 0xd1, 0xb5, 0x4d, 0x85, 0xc1, 0x35, 0xcf, 0x94, 0x43, 0xa3, 0x84, 0xac, 0x71, 0x3a, 0x86,
	0xf3, 0x05, 0x5a, 0x9b, 0x69, 0x35, 0xf3, 0x68, 0x0a, 0x72, 0x05, 0x27, 0x4e, 0x6f, 0x50, 0x0d,
	0xbb, 0xa3, 0xee, 0xc3, 0xe9, 0xbc, 0x16, 0x94, 0x42, 0x7f, 0x69, 0xd1, 0x2c, 0xa4, 0x4f, 0xc8,
	0x00, 0x7a, 0xde, 0xa2, 0xf9, 0x78, 0x0b, 0x48, 0x50, 0x34, 0x86, 0x8b, 0x17, 0xef, 0xd2, 0xd7,
	0x54, 0x48, 0x89, 0x2a, 0xc1, 0x36, 0x90, 0x44, 0xd0, 0xdf, 0x8a, 0x24, 0x8b, 0x97, 0x46, 0x0e,
	0x8f, 0x2a, 0x67, 0xaf, 0xc9, 0x2d, 0x40, 0x35, 0x7f, 0x6a, 0xb3, 0xb6, 0xc3, 0xe3, 0xca, 0x6d,
	0xbc, 0x3c, 0xfd, 0x74, 0xe1, 0xac, 0x4c, 0x99, 0xd6, 0xad, 0xc8, 0x37, 0x5c, 0xbe, 0xa3, 0x0b,
	0x0d, 0x0e, 0xd1, 0x23, 0xf6, 0xff, 0x16, 0xd8, 0xae, 0x45, 0x34, 0x69, 0x23, 0xfe, 0x74, 0xa0,
	0x1d, 0x32, 0x03, 0x38, 0x24, 0x90, 0x71, 0xdb, 0xb7, 0xe6, 0x12, 0xa3, 0x9b, 0x06, 0x35, 0xdf,
	0x4d, 0x81, 0xa0, 0x9d, 0xe9, 0xfd, 0xd7, 0x24, 0xc9, 0x5c, 0xea, 0x57, 0x2c, 0xd6, 0x5b, 0xbe,
	0x47, 0xb9, 0xc9, 0x63, 0xde, 0x3c, 0xef, 0xaa, 0x57, 0xdd, 0xe8, 0xf9, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xa0, 0x7e, 0x4c, 0x0e, 0xf5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthBackendClient is the client API for AuthBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthBackendClient interface {
	GetSessionChallenge(ctx context.Context, in *UserSlug, opts ...grpc.CallOption) (*AuthChallenge, error)
	GetSession(ctx context.Context, in *SessionQuery, opts ...grpc.CallOption) (*shared.RoleypolySession, error)
}

type authBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthBackendClient(cc grpc.ClientConnInterface) AuthBackendClient {
	return &authBackendClient{cc}
}

func (c *authBackendClient) GetSessionChallenge(ctx context.Context, in *UserSlug, opts ...grpc.CallOption) (*AuthChallenge, error) {
	out := new(AuthChallenge)
	err := c.cc.Invoke(ctx, "/roleypoly.auth.backend.AuthBackend/GetSessionChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authBackendClient) GetSession(ctx context.Context, in *SessionQuery, opts ...grpc.CallOption) (*shared.RoleypolySession, error) {
	out := new(shared.RoleypolySession)
	err := c.cc.Invoke(ctx, "/roleypoly.auth.backend.AuthBackend/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthBackendServer is the server API for AuthBackend service.
type AuthBackendServer interface {
	GetSessionChallenge(context.Context, *UserSlug) (*AuthChallenge, error)
	GetSession(context.Context, *SessionQuery) (*shared.RoleypolySession, error)
}

// UnimplementedAuthBackendServer can be embedded to have forward compatible implementations.
type UnimplementedAuthBackendServer struct {
}

func (*UnimplementedAuthBackendServer) GetSessionChallenge(ctx context.Context, req *UserSlug) (*AuthChallenge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionChallenge not implemented")
}
func (*UnimplementedAuthBackendServer) GetSession(ctx context.Context, req *SessionQuery) (*shared.RoleypolySession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
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
		FullMethod: "/roleypoly.auth.backend.AuthBackend/GetSessionChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthBackendServer).GetSessionChallenge(ctx, req.(*UserSlug))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthBackend_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthBackendServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roleypoly.auth.backend.AuthBackend/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthBackendServer).GetSession(ctx, req.(*SessionQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roleypoly.auth.backend.AuthBackend",
	HandlerType: (*AuthBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSessionChallenge",
			Handler:    _AuthBackend_GetSessionChallenge_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _AuthBackend_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/backend/auth-backend.proto",
}
