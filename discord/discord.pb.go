// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discord.proto

package discord

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37f1fb6677774d8, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

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
	return fileDescriptor_d37f1fb6677774d8, []int{1}
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
	return fileDescriptor_d37f1fb6677774d8, []int{2}
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
	return fileDescriptor_d37f1fb6677774d8, []int{3}
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
	return fileDescriptor_d37f1fb6677774d8, []int{4}
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

type GuildMembers struct {
	ID                   string    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Members              []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GuildMembers) Reset()         { *m = GuildMembers{} }
func (m *GuildMembers) String() string { return proto.CompactTextString(m) }
func (*GuildMembers) ProtoMessage()    {}
func (*GuildMembers) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37f1fb6677774d8, []int{5}
}

func (m *GuildMembers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuildMembers.Unmarshal(m, b)
}
func (m *GuildMembers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuildMembers.Marshal(b, m, deterministic)
}
func (m *GuildMembers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuildMembers.Merge(m, src)
}
func (m *GuildMembers) XXX_Size() int {
	return xxx_messageInfo_GuildMembers.Size(m)
}
func (m *GuildMembers) XXX_DiscardUnknown() {
	xxx_messageInfo_GuildMembers.DiscardUnknown(m)
}

var xxx_messageInfo_GuildMembers proto.InternalMessageInfo

func (m *GuildMembers) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GuildMembers) GetMembers() []*Member {
	if m != nil {
		return m.Members
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
	return fileDescriptor_d37f1fb6677774d8, []int{6}
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

type Member struct {
	GuildID              string   `protobuf:"bytes,1,opt,name=guildID,proto3" json:"guildID,omitempty"`
	Roles                []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Nick                 string   `protobuf:"bytes,4,opt,name=nick,proto3" json:"nick,omitempty"`
	User                 *User    `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37f1fb6677774d8, []int{7}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetGuildID() string {
	if m != nil {
		return m.GuildID
	}
	return ""
}

func (m *Member) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Member) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *Member) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Discriminator        string   `protobuf:"bytes,3,opt,name=discriminator,proto3" json:"discriminator,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Bot                  bool     `protobuf:"varint,5,opt,name=bot,proto3" json:"bot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37f1fb6677774d8, []int{8}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetDiscriminator() string {
	if m != nil {
		return m.Discriminator
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetBot() bool {
	if m != nil {
		return m.Bot
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*IDQuery)(nil), "IDQuery")
	proto.RegisterType((*GuildList)(nil), "GuildList")
	proto.RegisterType((*Guild)(nil), "Guild")
	proto.RegisterType((*GuildRoles)(nil), "GuildRoles")
	proto.RegisterType((*GuildMembers)(nil), "GuildMembers")
	proto.RegisterType((*Role)(nil), "Role")
	proto.RegisterType((*Member)(nil), "Member")
	proto.RegisterType((*User)(nil), "User")
}

func init() { proto.RegisterFile("discord.proto", fileDescriptor_d37f1fb6677774d8) }

var fileDescriptor_d37f1fb6677774d8 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x6d, 0x9a, 0xcf, 0x4e, 0x17, 0x84, 0x2c, 0x84, 0x42, 0x91, 0x56, 0xc1, 0x70, 0xa8, 0x84,
	0x94, 0xc3, 0x72, 0xe2, 0x84, 0x10, 0x41, 0x55, 0x25, 0x38, 0x60, 0xb4, 0x3f, 0xc0, 0x6d, 0xac,
	0x62, 0xd1, 0xc4, 0x91, 0xed, 0x2e, 0xda, 0x2b, 0x7f, 0x01, 0x0e, 0xfc, 0x5c, 0xe4, 0xb1, 0x13,
	0x8a, 0x7a, 0xe1, 0x36, 0xef, 0x79, 0x1c, 0xbf, 0xbc, 0x79, 0x03, 0x0f, 0x5a, 0x69, 0xf6, 0x4a,
	0xb7, 0xf5, 0xa0, 0x95, 0x55, 0x34, 0x87, 0xf4, 0x43, 0x37, 0xd8, 0x7b, 0xfa, 0x16, 0xf2, 0x6d,
	0xf3, 0xf9, 0x24, 0xf4, 0x3d, 0x59, 0x41, 0xf1, 0x49, 0x74, 0x3b, 0xa1, 0xb7, 0x4d, 0x19, 0x55,
	0xd1, 0x7a, 0xc1, 0x26, 0x4c, 0x4a, 0xc8, 0x37, 0x27, 0x79, 0x6c, 0xb7, 0x4d, 0x39, 0xc7, 0xa3,
	0x11, 0xd2, 0x57, 0xb0, 0xc0, 0xf2, 0xa3, 0x34, 0x96, 0x5c, 0x43, 0x76, 0x70, 0xc0, 0x94, 0x51,
	0x15, 0xaf, 0x97, 0x37, 0x59, 0x8d, 0x67, 0x2c, 0xb0, 0xf4, 0x67, 0x04, 0x29, 0x32, 0xe4, 0x21,
	0xcc, 0xa7, 0x67, 0xe6, 0xdb, 0x86, 0x10, 0x48, 0x7a, 0xde, 0x89, 0xf0, 0x75, 0xac, 0x1d, 0x27,
	0xf7, 0xaa, 0x2f, 0x63, 0xcf, 0xb9, 0xda, 0x09, 0x51, 0xdf, 0x7b, 0xd4, 0x98, 0x78, 0x21, 0x01,
	0x92, 0x0a, 0x96, 0x1d, 0xca, 0x7d, 0xaf, 0x4e, 0xbd, 0x2d, 0xd3, 0x2a, 0x5a, 0xa7, 0xec, 0x9c,
	0x22, 0x4f, 0x20, 0x33, 0xc3, 0x91, 0x9b, 0xaf, 0x65, 0x86, 0x57, 0x03, 0xa2, 0x6f, 0x00, 0xbc,
	0x4c, 0x75, 0x14, 0xe6, 0x42, 0xd9, 0x33, 0x48, 0xb5, 0x3b, 0x28, 0xe7, 0xf8, 0x4b, 0x69, 0xed,
	0xda, 0x98, 0xe7, 0xe8, 0x3b, 0xb8, 0xc2, 0xab, 0xde, 0xa8, 0xcb, 0xcb, 0xcf, 0x21, 0xf7, 0x0a,
	0xc6, 0xeb, 0x79, 0xed, 0x5b, 0xd9, 0xc8, 0xd3, 0xdf, 0x11, 0x24, 0xee, 0x93, 0xff, 0x65, 0x49,
	0x05, 0xcb, 0x41, 0xe8, 0x4e, 0x1a, 0x23, 0x55, 0x6f, 0xd0, 0x99, 0x98, 0x9d, 0x53, 0xe4, 0x31,
	0xa4, 0x7b, 0x75, 0x54, 0x1a, 0xed, 0x49, 0x99, 0x07, 0x6e, 0xb6, 0x83, 0x32, 0xd2, 0x4a, 0xd5,
	0x07, 0x67, 0x26, 0xec, 0x2c, 0xed, 0x78, 0xcf, 0x0f, 0xa2, 0x45, 0x5f, 0x0a, 0x36, 0x42, 0x7a,
	0x80, 0xcc, 0xab, 0x75, 0x3d, 0x87, 0x30, 0x7f, 0x2f, 0x70, 0x84, 0xee, 0x3d, 0x6f, 0x4f, 0x5c,
	0xc5, 0xeb, 0x45, 0xf0, 0x05, 0xb5, 0xcb, 0xfd, 0xb7, 0x30, 0x23, 0xac, 0xc9, 0x53, 0x48, 0x4e,
	0x46, 0x68, 0x7c, 0xdf, 0xf9, 0x78, 0x6b, 0x84, 0x66, 0x48, 0xd1, 0x1f, 0x11, 0x24, 0x0e, 0x5e,
	0x78, 0xb0, 0x82, 0xc2, 0x35, 0x9c, 0xf9, 0x30, 0x61, 0xf2, 0xd2, 0x87, 0x5a, 0xcb, 0x4e, 0xf6,
	0xdc, 0x2a, 0x1d, 0x72, 0xf2, 0x2f, 0xe9, 0x86, 0xce, 0xef, 0xb8, 0xe5, 0x3a, 0x68, 0x09, 0x88,
	0x3c, 0x82, 0x78, 0xa7, 0x7c, 0x4c, 0x0a, 0xe6, 0xca, 0x9b, 0x5f, 0x11, 0xe4, 0x8d, 0xdf, 0x12,
	0xf2, 0x02, 0x96, 0x2e, 0xd0, 0x5f, 0x84, 0xbe, 0x73, 0x63, 0xcd, 0x6a, 0xdc, 0x96, 0x15, 0xd4,
	0x53, 0xd6, 0xe9, 0x8c, 0x5c, 0x43, 0xb1, 0x11, 0xd6, 0xe7, 0xb9, 0xa8, 0xc3, 0x1a, 0xad, 0x42,
	0xe6, 0xe9, 0x8c, 0x54, 0xb0, 0xd8, 0x08, 0x1b, 0x1c, 0xfc, 0xdb, 0x30, 0x46, 0x80, 0xce, 0x08,
	0x85, 0xab, 0xdb, 0xa1, 0xe5, 0x56, 0x84, 0xa6, 0xf1, 0xe8, 0xac, 0x67, 0x97, 0xe1, 0xc6, 0xbe,
	0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xb7, 0x12, 0x5d, 0xc2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscordClient is the client API for Discord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscordClient interface {
	ListServers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GuildList, error)
	GetGuild(ctx context.Context, in *IDQuery, opts ...grpc.CallOption) (*Guild, error)
	GetMember(ctx context.Context, in *IDQuery, opts ...grpc.CallOption) (*Member, error)
	UpdateMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error)
}

type discordClient struct {
	cc *grpc.ClientConn
}

func NewDiscordClient(cc *grpc.ClientConn) DiscordClient {
	return &discordClient{cc}
}

func (c *discordClient) ListServers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GuildList, error) {
	out := new(GuildList)
	err := c.cc.Invoke(ctx, "/Discord/ListServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordClient) GetGuild(ctx context.Context, in *IDQuery, opts ...grpc.CallOption) (*Guild, error) {
	out := new(Guild)
	err := c.cc.Invoke(ctx, "/Discord/GetGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordClient) GetMember(ctx context.Context, in *IDQuery, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/Discord/GetMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discordClient) UpdateMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/Discord/UpdateMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscordServer is the server API for Discord service.
type DiscordServer interface {
	ListServers(context.Context, *Empty) (*GuildList, error)
	GetGuild(context.Context, *IDQuery) (*Guild, error)
	GetMember(context.Context, *IDQuery) (*Member, error)
	UpdateMember(context.Context, *Member) (*Member, error)
}

// UnimplementedDiscordServer can be embedded to have forward compatible implementations.
type UnimplementedDiscordServer struct {
}

func (*UnimplementedDiscordServer) ListServers(ctx context.Context, req *Empty) (*GuildList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServers not implemented")
}
func (*UnimplementedDiscordServer) GetGuild(ctx context.Context, req *IDQuery) (*Guild, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuild not implemented")
}
func (*UnimplementedDiscordServer) GetMember(ctx context.Context, req *IDQuery) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (*UnimplementedDiscordServer) UpdateMember(ctx context.Context, req *Member) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMember not implemented")
}

func RegisterDiscordServer(s *grpc.Server, srv DiscordServer) {
	s.RegisterService(&_Discord_serviceDesc, srv)
}

func _Discord_ListServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServer).ListServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Discord/ListServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServer).ListServers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discord_GetGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServer).GetGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Discord/GetGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServer).GetGuild(ctx, req.(*IDQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discord_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Discord/GetMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServer).GetMember(ctx, req.(*IDQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discord_UpdateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordServer).UpdateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Discord/UpdateMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordServer).UpdateMember(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Discord",
	HandlerType: (*DiscordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServers",
			Handler:    _Discord_ListServers_Handler,
		},
		{
			MethodName: "GetGuild",
			Handler:    _Discord_GetGuild_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _Discord_GetMember_Handler,
		},
		{
			MethodName: "UpdateMember",
			Handler:    _Discord_UpdateMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discord.proto",
}
