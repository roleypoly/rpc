// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared/shared.proto

package roleypoly

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type IDQuery struct {
	MemberID             string   `protobuf:"bytes,1,opt,name=MemberID,proto3" json:"MemberID,omitempty"`
	GuildID              string   `protobuf:"bytes,2,opt,name=GuildID,proto3" json:"GuildID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDQuery) Reset()         { *m = IDQuery{} }
func (m *IDQuery) String() string { return proto.CompactTextString(m) }
func (*IDQuery) ProtoMessage()    {}
func (*IDQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9301c8e954a55bc6, []int{0}
}

func (m *IDQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDQuery.Unmarshal(m, b)
}
func (m *IDQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDQuery.Marshal(b, m, deterministic)
}
func (m *IDQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDQuery.Merge(m, src)
}
func (m *IDQuery) XXX_Size() int {
	return xxx_messageInfo_IDQuery.Size(m)
}
func (m *IDQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_IDQuery.DiscardUnknown(m)
}

var xxx_messageInfo_IDQuery proto.InternalMessageInfo

func (m *IDQuery) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

func (m *IDQuery) GetGuildID() string {
	if m != nil {
		return m.GuildID
	}
	return ""
}

type GuildList struct {
	Guilds               []*Guild `protobuf:"bytes,1,rep,name=guilds,proto3" json:"guilds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuildList) Reset()         { *m = GuildList{} }
func (m *GuildList) String() string { return proto.CompactTextString(m) }
func (*GuildList) ProtoMessage()    {}
func (*GuildList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9301c8e954a55bc6, []int{1}
}

func (m *GuildList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuildList.Unmarshal(m, b)
}
func (m *GuildList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuildList.Marshal(b, m, deterministic)
}
func (m *GuildList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuildList.Merge(m, src)
}
func (m *GuildList) XXX_Size() int {
	return xxx_messageInfo_GuildList.Size(m)
}
func (m *GuildList) XXX_DiscardUnknown() {
	xxx_messageInfo_GuildList.DiscardUnknown(m)
}

var xxx_messageInfo_GuildList proto.InternalMessageInfo

func (m *GuildList) GetGuilds() []*Guild {
	if m != nil {
		return m.Guilds
	}
	return nil
}

type Guild struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	OwnerID              string   `protobuf:"bytes,4,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
	MemberCount          int32    `protobuf:"varint,5,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	Splash               string   `protobuf:"bytes,6,opt,name=splash,proto3" json:"splash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Guild) Reset()         { *m = Guild{} }
func (m *Guild) String() string { return proto.CompactTextString(m) }
func (*Guild) ProtoMessage()    {}
func (*Guild) Descriptor() ([]byte, []int) {
	return fileDescriptor_9301c8e954a55bc6, []int{2}
}

func (m *Guild) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Guild.Unmarshal(m, b)
}
func (m *Guild) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Guild.Marshal(b, m, deterministic)
}
func (m *Guild) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Guild.Merge(m, src)
}
func (m *Guild) XXX_Size() int {
	return xxx_messageInfo_Guild.Size(m)
}
func (m *Guild) XXX_DiscardUnknown() {
	xxx_messageInfo_Guild.DiscardUnknown(m)
}

var xxx_messageInfo_Guild proto.InternalMessageInfo

func (m *Guild) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Guild) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Guild) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Guild) GetOwnerID() string {
	if m != nil {
		return m.OwnerID
	}
	return ""
}

func (m *Guild) GetMemberCount() int32 {
	if m != nil {
		return m.MemberCount
	}
	return 0
}

func (m *Guild) GetSplash() string {
	if m != nil {
		return m.Splash
	}
	return ""
}

type GuildRoles struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Roles                []*Role  `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuildRoles) Reset()         { *m = GuildRoles{} }
func (m *GuildRoles) String() string { return proto.CompactTextString(m) }
func (*GuildRoles) ProtoMessage()    {}
func (*GuildRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_9301c8e954a55bc6, []int{3}
}

func (m *GuildRoles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuildRoles.Unmarshal(m, b)
}
func (m *GuildRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuildRoles.Marshal(b, m, deterministic)
}
func (m *GuildRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuildRoles.Merge(m, src)
}
func (m *GuildRoles) XXX_Size() int {
	return xxx_messageInfo_GuildRoles.Size(m)
}
func (m *GuildRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_GuildRoles.DiscardUnknown(m)
}

var xxx_messageInfo_GuildRoles proto.InternalMessageInfo

func (m *GuildRoles) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GuildRoles) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Role struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions          int64    `protobuf:"varint,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Color                int32    `protobuf:"varint,4,opt,name=color,proto3" json:"color,omitempty"`
	Position             int32    `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
	Managed              bool     `protobuf:"varint,6,opt,name=managed,proto3" json:"managed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_9301c8e954a55bc6, []int{4}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetPermissions() int64 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

func (m *Role) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *Role) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Role) GetManaged() bool {
	if m != nil {
		return m.Managed
	}
	return false
}

func init() {
	proto.RegisterType((*IDQuery)(nil), "roleypoly.IDQuery")
	proto.RegisterType((*GuildList)(nil), "roleypoly.GuildList")
	proto.RegisterType((*Guild)(nil), "roleypoly.Guild")
	proto.RegisterType((*GuildRoles)(nil), "roleypoly.GuildRoles")
	proto.RegisterType((*Role)(nil), "roleypoly.Role")
}

func init() { proto.RegisterFile("shared/shared.proto", fileDescriptor_9301c8e954a55bc6) }

var fileDescriptor_9301c8e954a55bc6 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0xb1, 0x13, 0x39, 0xc9, 0x04, 0xfe, 0xbf, 0xa8, 0xa5, 0x88, 0xae, 0x8c, 0xa1, 0xe0,
	0x55, 0x0a, 0x2d, 0x5d, 0x77, 0x11, 0x43, 0x31, 0xb4, 0x8b, 0xea, 0x06, 0x4e, 0x22, 0x12, 0x81,
	0xac, 0x31, 0x92, 0x43, 0xc9, 0x39, 0xba, 0xe9, 0x71, 0x8b, 0x46, 0x8e, 0x09, 0x74, 0xd3, 0x95,
	0xe7, 0x7b, 0xe3, 0x61, 0xde, 0xd3, 0xc0, 0xb5, 0x3f, 0x34, 0x4e, 0xed, 0x1e, 0xe2, 0x67, 0xd5,
	0x39, 0xec, 0x91, 0x2f, 0x1c, 0x1a, 0x75, 0xea, 0xd0, 0x9c, 0x8a, 0x17, 0x98, 0xd5, 0xd5, 0xc7,
	0x51, 0xb9, 0x13, 0xbf, 0x83, 0xf9, 0xbb, 0x6a, 0x37, 0xca, 0xd5, 0x95, 0x48, 0xf2, 0xa4, 0x5c,
	0xc8, 0x91, 0xb9, 0x80, 0xd9, 0xeb, 0x51, 0x9b, 0x5d, 0x5d, 0x89, 0x94, 0x5a, 0x67, 0x2c, 0x9e,
	0x61, 0x41, 0xe5, 0x9b, 0xf6, 0x3d, 0x2f, 0x21, 0xdb, 0x07, 0xf0, 0x22, 0xc9, 0x27, 0xe5, 0xf2,
	0xf1, 0x6a, 0x35, 0x6e, 0x5a, 0xd1, 0x5f, 0x72, 0xe8, 0x17, 0x5f, 0x09, 0x30, 0x52, 0xf8, 0x3f,
	0x48, 0xc7, 0x85, 0x69, 0x5d, 0x71, 0x0e, 0x53, 0xdb, 0xb4, 0x6a, 0xd8, 0x43, 0x75, 0xd0, 0xf4,
	0x16, 0xad, 0x98, 0x44, 0x2d, 0xd4, 0xc1, 0x12, 0x7e, 0x5a, 0x72, 0x3b, 0x8d, 0x96, 0x06, 0xe4,
	0x39, 0x2c, 0x5b, 0x32, 0xbe, 0xc6, 0xa3, 0xed, 0x05, 0xcb, 0x93, 0x92, 0xc9, 0x4b, 0x89, 0xdf,
	0x42, 0xe6, 0x3b, 0xd3, 0xf8, 0x83, 0xc8, 0x68, 0x74, 0xa0, 0x62, 0x0d, 0x10, 0x6d, 0xa2, 0x51,
	0xfe, 0x97, 0xb3, 0x7b, 0x60, 0x21, 0x8e, 0x17, 0x29, 0x85, 0xfb, 0x7f, 0x11, 0x2e, 0x0c, 0xc8,
	0xd8, 0x2d, 0xbe, 0x13, 0x98, 0x06, 0xfe, 0x53, 0xb2, 0x1c, 0x96, 0x9d, 0x72, 0xad, 0xf6, 0x5e,
	0xa3, 0xf5, 0x14, 0x70, 0x22, 0x2f, 0x25, 0x7e, 0x03, 0x6c, 0x8b, 0x06, 0x1d, 0xa5, 0x64, 0x32,
	0x42, 0x38, 0x56, 0x87, 0x5e, 0xf7, 0x1a, 0xed, 0x10, 0x70, 0xe4, 0xf0, 0x32, 0x6d, 0x63, 0x9b,
	0xbd, 0xda, 0x51, 0xbc, 0xb9, 0x3c, 0xe3, 0x26, 0xa3, 0xfb, 0x3f, 0xfd, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x5d, 0xf0, 0x97, 0x45, 0x16, 0x02, 0x00, 0x00,
}
