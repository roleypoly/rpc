// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: platform/platform.proto

package platform

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	discord "github.com/roleypoly/rpc/discord"
	shared "github.com/roleypoly/rpc/shared"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type Category_CategoryType int32

const (
	Category_multi  Category_CategoryType = 0
	Category_single Category_CategoryType = 1
)

// Enum value maps for Category_CategoryType.
var (
	Category_CategoryType_name = map[int32]string{
		0: "multi",
		1: "single",
	}
	Category_CategoryType_value = map[string]int32{
		"multi":  0,
		"single": 1,
	}
)

func (x Category_CategoryType) Enum() *Category_CategoryType {
	p := new(Category_CategoryType)
	*p = x
	return p
}

func (x Category_CategoryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category_CategoryType) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_platform_proto_enumTypes[0].Descriptor()
}

func (Category_CategoryType) Type() protoreflect.EnumType {
	return &file_platform_platform_proto_enumTypes[0]
}

func (x Category_CategoryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category_CategoryType.Descriptor instead.
func (Category_CategoryType) EnumDescriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{5, 0}
}

type GuildEnumeration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guilds []*PresentableGuild `protobuf:"bytes,1,rep,name=guilds,proto3" json:"guilds,omitempty"`
}

func (x *GuildEnumeration) Reset() {
	*x = GuildEnumeration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuildEnumeration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuildEnumeration) ProtoMessage() {}

func (x *GuildEnumeration) ProtoReflect() protoreflect.Message {
	mi := &file_platform_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuildEnumeration.ProtoReflect.Descriptor instead.
func (*GuildEnumeration) Descriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{0}
}

func (x *GuildEnumeration) GetGuilds() []*PresentableGuild {
	if x != nil {
		return x.Guilds
	}
	return nil
}

type PresentableGuild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string             `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Guild  *shared.Guild      `protobuf:"bytes,2,opt,name=guild,proto3" json:"guild,omitempty"`
	Data   *GuildData         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Member *discord.Member    `protobuf:"bytes,4,opt,name=member,proto3" json:"member,omitempty"`
	Roles  *shared.GuildRoles `protobuf:"bytes,5,opt,name=roles,proto3" json:"roles,omitempty"`
}

func (x *PresentableGuild) Reset() {
	*x = PresentableGuild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresentableGuild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresentableGuild) ProtoMessage() {}

func (x *PresentableGuild) ProtoReflect() protoreflect.Message {
	mi := &file_platform_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresentableGuild.ProtoReflect.Descriptor instead.
func (*PresentableGuild) Descriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{1}
}

func (x *PresentableGuild) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PresentableGuild) GetGuild() *shared.Guild {
	if x != nil {
		return x.Guild
	}
	return nil
}

func (x *PresentableGuild) GetData() *GuildData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PresentableGuild) GetMember() *discord.Member {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *PresentableGuild) GetRoles() *shared.GuildRoles {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GuildData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Message      string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Categories   []*Category `protobuf:"bytes,3,rep,name=categories,proto3" json:"categories,omitempty"`
	Entitlements []string    `protobuf:"bytes,4,rep,name=entitlements,proto3" json:"entitlements,omitempty"`
}

func (x *GuildData) Reset() {
	*x = GuildData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuildData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuildData) ProtoMessage() {}

func (x *GuildData) ProtoReflect() protoreflect.Message {
	mi := &file_platform_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuildData.ProtoReflect.Descriptor instead.
func (*GuildData) Descriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{2}
}

func (x *GuildData) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *GuildData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GuildData) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *GuildData) GetEntitlements() []string {
	if x != nil {
		return x.Entitlements
	}
	return nil
}

type UpdateRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildID string `protobuf:"bytes,1,opt,name=guildID,proto3" json:"guildID,omitempty"`
	Roles   *Roles `protobuf:"bytes,2,opt,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UpdateRoles) Reset() {
	*x = UpdateRoles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoles) ProtoMessage() {}

func (x *UpdateRoles) ProtoReflect() protoreflect.Message {
	mi := &file_platform_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoles.ProtoReflect.Descriptor instead.
func (*UpdateRoles) Descriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRoles) GetGuildID() string {
	if x != nil {
		return x.GuildID
	}
	return ""
}

func (x *UpdateRoles) GetRoles() *Roles {
	if x != nil {
		return x.Roles
	}
	return nil
}

type Roles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []string `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Roles) Reset() {
	*x = Roles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_platform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roles) ProtoMessage() {}

func (x *Roles) ProtoReflect() protoreflect.Message {
	mi := &file_platform_platform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roles.ProtoReflect.Descriptor instead.
func (*Roles) Descriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{4}
}

func (x *Roles) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string                `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Roles    []string              `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Hidden   bool                  `protobuf:"varint,4,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Type     Category_CategoryType `protobuf:"varint,5,opt,name=type,proto3,enum=roleypoly.platform.Category_CategoryType" json:"type,omitempty"`
	Position int32                 `protobuf:"varint,6,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_platform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_platform_platform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{5}
}

func (x *Category) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Category) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *Category) GetType() Category_CategoryType {
	if x != nil {
		return x.Type
	}
	return Category_multi
}

func (x *Category) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type UpdateEntitlement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *shared.IDQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Name  string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State bool            `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *UpdateEntitlement) Reset() {
	*x = UpdateEntitlement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_platform_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntitlement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntitlement) ProtoMessage() {}

func (x *UpdateEntitlement) ProtoReflect() protoreflect.Message {
	mi := &file_platform_platform_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntitlement.ProtoReflect.Descriptor instead.
func (*UpdateEntitlement) Descriptor() ([]byte, []int) {
	return file_platform_platform_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEntitlement) GetQuery() *shared.IDQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *UpdateEntitlement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEntitlement) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

var File_platform_platform_proto protoreflect.FileDescriptor

var file_platform_platform_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x6f, 0x6c, 0x65, 0x79,
	0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x1a, 0x13, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x10, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x45,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x06, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x05,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79,
	0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x09, 0x47, 0x75, 0x69,
	0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x58, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x1d, 0x0a, 0x05,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x08,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79,
	0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x10, 0x01, 0x22, 0x67, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x49, 0x44, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xf7, 0x02, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x53, 0x0a, 0x11, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d,
	0x79, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x24, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70,
	0x6f, 0x6c, 0x79, 0x2e, 0x49, 0x44, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x10, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x12, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x49, 0x44, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x24, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x47, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79,
	0x70, 0x6f, 0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x75, 0x69,
	0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x79, 0x70, 0x6f,
	0x6c, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x47, 0x75, 0x69, 0x6c,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x79, 0x70, 0x6f, 0x6c, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_platform_proto_rawDescOnce sync.Once
	file_platform_platform_proto_rawDescData = file_platform_platform_proto_rawDesc
)

func file_platform_platform_proto_rawDescGZIP() []byte {
	file_platform_platform_proto_rawDescOnce.Do(func() {
		file_platform_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_platform_proto_rawDescData)
	})
	return file_platform_platform_proto_rawDescData
}

var file_platform_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_platform_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_platform_platform_proto_goTypes = []interface{}{
	(Category_CategoryType)(0), // 0: roleypoly.platform.Category.CategoryType
	(*GuildEnumeration)(nil),   // 1: roleypoly.platform.GuildEnumeration
	(*PresentableGuild)(nil),   // 2: roleypoly.platform.PresentableGuild
	(*GuildData)(nil),          // 3: roleypoly.platform.GuildData
	(*UpdateRoles)(nil),        // 4: roleypoly.platform.UpdateRoles
	(*Roles)(nil),              // 5: roleypoly.platform.Roles
	(*Category)(nil),           // 6: roleypoly.platform.Category
	(*UpdateEntitlement)(nil),  // 7: roleypoly.platform.UpdateEntitlement
	(*shared.Guild)(nil),       // 8: roleypoly.Guild
	(*discord.Member)(nil),     // 9: roleypoly.discord.Member
	(*shared.GuildRoles)(nil),  // 10: roleypoly.GuildRoles
	(*shared.IDQuery)(nil),     // 11: roleypoly.IDQuery
	(*empty.Empty)(nil),        // 12: google.protobuf.Empty
}
var file_platform_platform_proto_depIdxs = []int32{
	2,  // 0: roleypoly.platform.GuildEnumeration.guilds:type_name -> roleypoly.platform.PresentableGuild
	8,  // 1: roleypoly.platform.PresentableGuild.guild:type_name -> roleypoly.Guild
	3,  // 2: roleypoly.platform.PresentableGuild.data:type_name -> roleypoly.platform.GuildData
	9,  // 3: roleypoly.platform.PresentableGuild.member:type_name -> roleypoly.discord.Member
	10, // 4: roleypoly.platform.PresentableGuild.roles:type_name -> roleypoly.GuildRoles
	6,  // 5: roleypoly.platform.GuildData.categories:type_name -> roleypoly.platform.Category
	5,  // 6: roleypoly.platform.UpdateRoles.roles:type_name -> roleypoly.platform.Roles
	0,  // 7: roleypoly.platform.Category.type:type_name -> roleypoly.platform.Category.CategoryType
	11, // 8: roleypoly.platform.UpdateEntitlement.query:type_name -> roleypoly.IDQuery
	12, // 9: roleypoly.platform.Platform.EnumerateMyGuilds:input_type -> google.protobuf.Empty
	11, // 10: roleypoly.platform.Platform.GetGuildSlug:input_type -> roleypoly.IDQuery
	11, // 11: roleypoly.platform.Platform.GetGuild:input_type -> roleypoly.IDQuery
	4,  // 12: roleypoly.platform.Platform.UpdateMyRoles:input_type -> roleypoly.platform.UpdateRoles
	3,  // 13: roleypoly.platform.Platform.UpdateGuildData:input_type -> roleypoly.platform.GuildData
	1,  // 14: roleypoly.platform.Platform.EnumerateMyGuilds:output_type -> roleypoly.platform.GuildEnumeration
	8,  // 15: roleypoly.platform.Platform.GetGuildSlug:output_type -> roleypoly.Guild
	2,  // 16: roleypoly.platform.Platform.GetGuild:output_type -> roleypoly.platform.PresentableGuild
	12, // 17: roleypoly.platform.Platform.UpdateMyRoles:output_type -> google.protobuf.Empty
	12, // 18: roleypoly.platform.Platform.UpdateGuildData:output_type -> google.protobuf.Empty
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_platform_platform_proto_init() }
func file_platform_platform_proto_init() {
	if File_platform_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuildEnumeration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_platform_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresentableGuild); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_platform_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuildData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_platform_platform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoles); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_platform_platform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roles); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_platform_platform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_platform_platform_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntitlement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_platform_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_platform_proto_goTypes,
		DependencyIndexes: file_platform_platform_proto_depIdxs,
		EnumInfos:         file_platform_platform_proto_enumTypes,
		MessageInfos:      file_platform_platform_proto_msgTypes,
	}.Build()
	File_platform_platform_proto = out.File
	file_platform_platform_proto_rawDesc = nil
	file_platform_platform_proto_goTypes = nil
	file_platform_platform_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformClient interface {
	EnumerateMyGuilds(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GuildEnumeration, error)
	GetGuildSlug(ctx context.Context, in *shared.IDQuery, opts ...grpc.CallOption) (*shared.Guild, error)
	GetGuild(ctx context.Context, in *shared.IDQuery, opts ...grpc.CallOption) (*PresentableGuild, error)
	UpdateMyRoles(ctx context.Context, in *UpdateRoles, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateGuildData(ctx context.Context, in *GuildData, opts ...grpc.CallOption) (*empty.Empty, error)
}

type platformClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformClient(cc grpc.ClientConnInterface) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) EnumerateMyGuilds(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GuildEnumeration, error) {
	out := new(GuildEnumeration)
	err := c.cc.Invoke(ctx, "/roleypoly.platform.Platform/EnumerateMyGuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) GetGuildSlug(ctx context.Context, in *shared.IDQuery, opts ...grpc.CallOption) (*shared.Guild, error) {
	out := new(shared.Guild)
	err := c.cc.Invoke(ctx, "/roleypoly.platform.Platform/GetGuildSlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) GetGuild(ctx context.Context, in *shared.IDQuery, opts ...grpc.CallOption) (*PresentableGuild, error) {
	out := new(PresentableGuild)
	err := c.cc.Invoke(ctx, "/roleypoly.platform.Platform/GetGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) UpdateMyRoles(ctx context.Context, in *UpdateRoles, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/roleypoly.platform.Platform/UpdateMyRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) UpdateGuildData(ctx context.Context, in *GuildData, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/roleypoly.platform.Platform/UpdateGuildData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
type PlatformServer interface {
	EnumerateMyGuilds(context.Context, *empty.Empty) (*GuildEnumeration, error)
	GetGuildSlug(context.Context, *shared.IDQuery) (*shared.Guild, error)
	GetGuild(context.Context, *shared.IDQuery) (*PresentableGuild, error)
	UpdateMyRoles(context.Context, *UpdateRoles) (*empty.Empty, error)
	UpdateGuildData(context.Context, *GuildData) (*empty.Empty, error)
}

// UnimplementedPlatformServer can be embedded to have forward compatible implementations.
type UnimplementedPlatformServer struct {
}

func (*UnimplementedPlatformServer) EnumerateMyGuilds(context.Context, *empty.Empty) (*GuildEnumeration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnumerateMyGuilds not implemented")
}
func (*UnimplementedPlatformServer) GetGuildSlug(context.Context, *shared.IDQuery) (*shared.Guild, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuildSlug not implemented")
}
func (*UnimplementedPlatformServer) GetGuild(context.Context, *shared.IDQuery) (*PresentableGuild, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuild not implemented")
}
func (*UnimplementedPlatformServer) UpdateMyRoles(context.Context, *UpdateRoles) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMyRoles not implemented")
}
func (*UnimplementedPlatformServer) UpdateGuildData(context.Context, *GuildData) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuildData not implemented")
}

func RegisterPlatformServer(s *grpc.Server, srv PlatformServer) {
	s.RegisterService(&_Platform_serviceDesc, srv)
}

func _Platform_EnumerateMyGuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).EnumerateMyGuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roleypoly.platform.Platform/EnumerateMyGuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).EnumerateMyGuilds(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_GetGuildSlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.IDQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).GetGuildSlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roleypoly.platform.Platform/GetGuildSlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).GetGuildSlug(ctx, req.(*shared.IDQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_GetGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.IDQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).GetGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roleypoly.platform.Platform/GetGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).GetGuild(ctx, req.(*shared.IDQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_UpdateMyRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).UpdateMyRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roleypoly.platform.Platform/UpdateMyRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).UpdateMyRoles(ctx, req.(*UpdateRoles))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_UpdateGuildData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuildData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).UpdateGuildData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roleypoly.platform.Platform/UpdateGuildData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).UpdateGuildData(ctx, req.(*GuildData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Platform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roleypoly.platform.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnumerateMyGuilds",
			Handler:    _Platform_EnumerateMyGuilds_Handler,
		},
		{
			MethodName: "GetGuildSlug",
			Handler:    _Platform_GetGuildSlug_Handler,
		},
		{
			MethodName: "GetGuild",
			Handler:    _Platform_GetGuild_Handler,
		},
		{
			MethodName: "UpdateMyRoles",
			Handler:    _Platform_UpdateMyRoles_Handler,
		},
		{
			MethodName: "UpdateGuildData",
			Handler:    _Platform_UpdateGuildData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform/platform.proto",
}
