// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared/internal.proto

package roleypoly

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	discord "github.com/roleypoly/rpc/discord"
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

type RoleypolySession_SessionSource int32

const (
	RoleypolySession_OAUTH RoleypolySession_SessionSource = 0
	RoleypolySession_DM    RoleypolySession_SessionSource = 1
)

var RoleypolySession_SessionSource_name = map[int32]string{
	0: "OAUTH",
	1: "DM",
}

var RoleypolySession_SessionSource_value = map[string]int32{
	"OAUTH": 0,
	"DM":    1,
}

func (x RoleypolySession_SessionSource) String() string {
	return proto.EnumName(RoleypolySession_SessionSource_name, int32(x))
}

func (RoleypolySession_SessionSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c67730e27cff8995, []int{1, 0}
}

type RoleypolyUser struct {
	DiscordUser          *discord.User `protobuf:"bytes,1,opt,name=discordUser,proto3" json:"discordUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RoleypolyUser) Reset()         { *m = RoleypolyUser{} }
func (m *RoleypolyUser) String() string { return proto.CompactTextString(m) }
func (*RoleypolyUser) ProtoMessage()    {}
func (*RoleypolyUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67730e27cff8995, []int{0}
}

func (m *RoleypolyUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleypolyUser.Unmarshal(m, b)
}
func (m *RoleypolyUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleypolyUser.Marshal(b, m, deterministic)
}
func (m *RoleypolyUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleypolyUser.Merge(m, src)
}
func (m *RoleypolyUser) XXX_Size() int {
	return xxx_messageInfo_RoleypolyUser.Size(m)
}
func (m *RoleypolyUser) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleypolyUser.DiscardUnknown(m)
}

var xxx_messageInfo_RoleypolyUser proto.InternalMessageInfo

func (m *RoleypolyUser) GetDiscordUser() *discord.User {
	if m != nil {
		return m.DiscordUser
	}
	return nil
}

type RoleypolySession struct {
	Id                   string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *RoleypolyUser                 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Source               RoleypolySession_SessionSource `protobuf:"varint,3,opt,name=source,proto3,enum=roleypoly.RoleypolySession_SessionSource" json:"source,omitempty"`
	CreatedAt            int64                          `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresIn            int64                          `protobuf:"varint,5,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	Extra                map[string]string              `protobuf:"bytes,6,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RoleypolySession) Reset()         { *m = RoleypolySession{} }
func (m *RoleypolySession) String() string { return proto.CompactTextString(m) }
func (*RoleypolySession) ProtoMessage()    {}
func (*RoleypolySession) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67730e27cff8995, []int{1}
}

func (m *RoleypolySession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleypolySession.Unmarshal(m, b)
}
func (m *RoleypolySession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleypolySession.Marshal(b, m, deterministic)
}
func (m *RoleypolySession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleypolySession.Merge(m, src)
}
func (m *RoleypolySession) XXX_Size() int {
	return xxx_messageInfo_RoleypolySession.Size(m)
}
func (m *RoleypolySession) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleypolySession.DiscardUnknown(m)
}

var xxx_messageInfo_RoleypolySession proto.InternalMessageInfo

func (m *RoleypolySession) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RoleypolySession) GetUser() *RoleypolyUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RoleypolySession) GetSource() RoleypolySession_SessionSource {
	if m != nil {
		return m.Source
	}
	return RoleypolySession_OAUTH
}

func (m *RoleypolySession) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *RoleypolySession) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *RoleypolySession) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterEnum("roleypoly.RoleypolySession_SessionSource", RoleypolySession_SessionSource_name, RoleypolySession_SessionSource_value)
	proto.RegisterType((*RoleypolyUser)(nil), "roleypoly.RoleypolyUser")
	proto.RegisterType((*RoleypolySession)(nil), "roleypoly.RoleypolySession")
	proto.RegisterMapType((map[string]string)(nil), "roleypoly.RoleypolySession.ExtraEntry")
}

func init() { proto.RegisterFile("shared/internal.proto", fileDescriptor_c67730e27cff8995) }

var fileDescriptor_c67730e27cff8995 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4b, 0xfb, 0x30,
	0x10, 0xc7, 0x7f, 0x6d, 0xd7, 0x42, 0x6f, 0x6c, 0x94, 0xf0, 0x1b, 0x96, 0x81, 0x50, 0xfa, 0x20,
	0x15, 0xa4, 0x83, 0xf9, 0x32, 0xc5, 0x97, 0x81, 0x03, 0x15, 0x44, 0xc8, 0xdc, 0xf3, 0xa8, 0xeb,
	0x81, 0xc1, 0xd2, 0x8c, 0x24, 0x95, 0xf5, 0x8f, 0xf6, 0x7f, 0x90, 0xa4, 0xd9, 0xdc, 0x40, 0x7c,
	0xba, 0xdc, 0xdd, 0xe7, 0xfb, 0xbd, 0xe4, 0x02, 0x23, 0xf9, 0x5e, 0x08, 0x2c, 0x27, 0xac, 0x56,
	0x28, 0xea, 0xa2, 0xca, 0xb7, 0x82, 0x2b, 0x4e, 0x42, 0xc1, 0x2b, 0x6c, 0xb7, 0xbc, 0x6a, 0xc7,
	0xa3, 0x92, 0xc9, 0x0d, 0x17, 0xe5, 0xc4, 0xc6, 0x8e, 0x48, 0x9f, 0x60, 0x40, 0xf7, 0xcc, 0x4a,
	0xa2, 0x20, 0x37, 0xd0, 0xb7, 0x84, 0x4e, 0x63, 0x27, 0x71, 0xb2, 0xfe, 0xf4, 0x2c, 0x3f, 0x18,
	0xe5, 0x7b, 0xbd, 0x6e, 0xd3, 0x63, 0x36, 0xfd, 0x72, 0x21, 0x3a, 0x98, 0x2d, 0x51, 0x4a, 0xc6,
	0x6b, 0x32, 0x04, 0x97, 0x95, 0xc6, 0x26, 0xa4, 0x2e, 0x2b, 0xc9, 0x15, 0xf4, 0x1a, 0x6d, 0xec,
	0x1a, 0xe3, 0xf8, 0xc8, 0xf8, 0xe4, 0x1e, 0xd4, 0x50, 0x64, 0x0e, 0x81, 0xe4, 0x8d, 0xd8, 0x60,
	0xec, 0x25, 0x4e, 0x36, 0x9c, 0x5e, 0xfe, 0xc6, 0xdb, 0x51, 0xb9, 0x8d, 0x4b, 0x23, 0xa0, 0x56,
	0x48, 0xce, 0x01, 0x36, 0x02, 0x0b, 0x85, 0xe5, 0xba, 0x50, 0x71, 0x2f, 0x71, 0x32, 0x8f, 0x86,
	0xb6, 0x32, 0x57, 0xba, 0x8d, 0xbb, 0x2d, 0x13, 0x28, 0xd7, 0xac, 0x8e, 0xfd, 0xae, 0x6d, 0x2b,
	0x8f, 0x35, 0xb9, 0x03, 0x1f, 0x77, 0x4a, 0x14, 0x71, 0x90, 0x78, 0x59, 0x7f, 0x7a, 0xf1, 0xd7,
	0xfc, 0x85, 0x06, 0x17, 0xb5, 0x12, 0x2d, 0xed, 0x44, 0xe3, 0x19, 0xc0, 0x4f, 0x91, 0x44, 0xe0,
	0x7d, 0x60, 0x6b, 0x77, 0xa1, 0x8f, 0xe4, 0x3f, 0xf8, 0x9f, 0x45, 0xd5, 0xa0, 0xd9, 0x46, 0x48,
	0xbb, 0xe4, 0xd6, 0x9d, 0x39, 0x69, 0x0a, 0x83, 0x93, 0xe7, 0x90, 0x10, 0xfc, 0x97, 0xf9, 0xea,
	0xf5, 0x21, 0xfa, 0x47, 0x02, 0x70, 0xef, 0x9f, 0x23, 0xe7, 0x2d, 0x30, 0x5f, 0x78, 0xfd, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x65, 0xbb, 0x84, 0x59, 0xfd, 0x01, 0x00, 0x00,
}
