// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: auth/backend/auth-backend.proto

package backend

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	auth "github.com/roleypoly/rpc/auth"
	shared "github.com/roleypoly/rpc/shared"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_auth_backend_auth_backend_proto protoreflect.FileDescriptor

var file_auth_backend_auth_backend_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x1a, 0x15, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9d, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x4a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12,
	0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x49, 0x44, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1b, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_backend_auth_backend_proto_goTypes = []interface{}{
	(*shared.IDQuery)(nil),          // 0: roleypoly.IDQuery
	(*auth.Token)(nil),              // 1: roleypoly.auth.Token
	(*auth.AuthChallenge)(nil),      // 2: roleypoly.auth.AuthChallenge
	(*shared.RoleypolySession)(nil), // 3: roleypoly.RoleypolySession
}
var file_auth_backend_auth_backend_proto_depIdxs = []int32{
	0, // 0: roleypoly.auth.backend.AuthBackend.GetNewAuthChallenge:input_type -> roleypoly.IDQuery
	1, // 1: roleypoly.auth.backend.AuthBackend.GetSession:input_type -> roleypoly.auth.Token
	2, // 2: roleypoly.auth.backend.AuthBackend.GetNewAuthChallenge:output_type -> roleypoly.auth.AuthChallenge
	3, // 3: roleypoly.auth.backend.AuthBackend.GetSession:output_type -> roleypoly.RoleypolySession
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_backend_auth_backend_proto_init() }
func file_auth_backend_auth_backend_proto_init() {
	if File_auth_backend_auth_backend_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_backend_auth_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_backend_auth_backend_proto_goTypes,
		DependencyIndexes: file_auth_backend_auth_backend_proto_depIdxs,
	}.Build()
	File_auth_backend_auth_backend_proto = out.File
	file_auth_backend_auth_backend_proto_rawDesc = nil
	file_auth_backend_auth_backend_proto_goTypes = nil
	file_auth_backend_auth_backend_proto_depIdxs = nil
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
	GetNewAuthChallenge(ctx context.Context, in *shared.IDQuery, opts ...grpc.CallOption) (*auth.AuthChallenge, error)
	GetSession(ctx context.Context, in *auth.Token, opts ...grpc.CallOption) (*shared.RoleypolySession, error)
}

type authBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthBackendClient(cc grpc.ClientConnInterface) AuthBackendClient {
	return &authBackendClient{cc}
}

func (c *authBackendClient) GetNewAuthChallenge(ctx context.Context, in *shared.IDQuery, opts ...grpc.CallOption) (*auth.AuthChallenge, error) {
	out := new(auth.AuthChallenge)
	err := c.cc.Invoke(ctx, "/roleypoly.auth.backend.AuthBackend/GetNewAuthChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authBackendClient) GetSession(ctx context.Context, in *auth.Token, opts ...grpc.CallOption) (*shared.RoleypolySession, error) {
	out := new(shared.RoleypolySession)
	err := c.cc.Invoke(ctx, "/roleypoly.auth.backend.AuthBackend/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthBackendServer is the server API for AuthBackend service.
type AuthBackendServer interface {
	GetNewAuthChallenge(context.Context, *shared.IDQuery) (*auth.AuthChallenge, error)
	GetSession(context.Context, *auth.Token) (*shared.RoleypolySession, error)
}

// UnimplementedAuthBackendServer can be embedded to have forward compatible implementations.
type UnimplementedAuthBackendServer struct {
}

func (*UnimplementedAuthBackendServer) GetNewAuthChallenge(context.Context, *shared.IDQuery) (*auth.AuthChallenge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewAuthChallenge not implemented")
}
func (*UnimplementedAuthBackendServer) GetSession(context.Context, *auth.Token) (*shared.RoleypolySession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}

func RegisterAuthBackendServer(s *grpc.Server, srv AuthBackendServer) {
	s.RegisterService(&_AuthBackend_serviceDesc, srv)
}

func _AuthBackend_GetNewAuthChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.IDQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthBackendServer).GetNewAuthChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roleypoly.auth.backend.AuthBackend/GetNewAuthChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthBackendServer).GetNewAuthChallenge(ctx, req.(*shared.IDQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthBackend_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.Token)
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
		return srv.(AuthBackendServer).GetSession(ctx, req.(*auth.Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roleypoly.auth.backend.AuthBackend",
	HandlerType: (*AuthBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewAuthChallenge",
			Handler:    _AuthBackend_GetNewAuthChallenge_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _AuthBackend_GetSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/backend/auth-backend.proto",
}
