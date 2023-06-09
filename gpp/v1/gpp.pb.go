// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: gpp/v1/gpp.proto

package gpp

import (
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

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	GameId string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty" validate:"required"`
	// @gotags: validate:"required"
	GameName string `protobuf:"bytes,2,opt,name=game_name,json=gameName,proto3" json:"game_name,omitempty" validate:"required"`
	// @gotags: validate:"required"
	GameShortName string `protobuf:"bytes,3,opt,name=game_short_name,json=gameShortName,proto3" json:"game_short_name,omitempty" validate:"required"`
	TeamIndex     int64  `protobuf:"varint,4,opt,name=team_index,json=teamIndex,proto3" json:"team_index,omitempty"`
	// @gotags: validate:"required"
	TeamId   string `protobuf:"bytes,5,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty" validate:"required"`
	TeamName string `protobuf:"bytes,6,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	// @gotags: validate:"required"
	GameRegion int64 `protobuf:"varint,7,opt,name=game_region,json=gameRegion,proto3" json:"game_region,omitempty" validate:"required"`
	// @gotags: validate:"required"
	GameRegionDesc string `protobuf:"bytes,8,opt,name=game_region_desc,json=gameRegionDesc,proto3" json:"game_region_desc,omitempty" validate:"required"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpp_v1_gpp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_gpp_v1_gpp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_gpp_v1_gpp_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *Game) GetGameName() string {
	if x != nil {
		return x.GameName
	}
	return ""
}

func (x *Game) GetGameShortName() string {
	if x != nil {
		return x.GameShortName
	}
	return ""
}

func (x *Game) GetTeamIndex() int64 {
	if x != nil {
		return x.TeamIndex
	}
	return 0
}

func (x *Game) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Game) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *Game) GetGameRegion() int64 {
	if x != nil {
		return x.GameRegion
	}
	return 0
}

func (x *Game) GetGameRegionDesc() string {
	if x != nil {
		return x.GameRegionDesc
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// @gotags: validate:"required"
	Username        string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" validate:"required"`
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	NameFirstLetter string `protobuf:"bytes,4,opt,name=name_first_letter,json=nameFirstLetter,proto3" json:"name_first_letter,omitempty"`
	Branch          string `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpp_v1_gpp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_gpp_v1_gpp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_gpp_v1_gpp_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetNameFirstLetter() string {
	if x != nil {
		return x.NameFirstLetter
	}
	return ""
}

func (x *User) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	Game *Game `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty" validate:"required"`
	// @gotags: validate:"required"
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Cookie string `protobuf:"bytes,3,opt,name=cookie,proto3" json:"cookie,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty" validate:"required"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpp_v1_gpp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_gpp_v1_gpp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_gpp_v1_gpp_proto_rawDescGZIP(), []int{2}
}

func (x *Info) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

func (x *Info) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Info) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

func (x *Info) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_gpp_v1_gpp_proto protoreflect.FileDescriptor

var file_gpp_v1_gpp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x67, 0x70, 0x70, 0x22, 0x84, 0x02, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61,
	0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x67, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x22, 0x8c,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x74, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x70, 0x70, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x59, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x70,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gpp_v1_gpp_proto_rawDescOnce sync.Once
	file_gpp_v1_gpp_proto_rawDescData = file_gpp_v1_gpp_proto_rawDesc
)

func file_gpp_v1_gpp_proto_rawDescGZIP() []byte {
	file_gpp_v1_gpp_proto_rawDescOnce.Do(func() {
		file_gpp_v1_gpp_proto_rawDescData = protoimpl.X.CompressGZIP(file_gpp_v1_gpp_proto_rawDescData)
	})
	return file_gpp_v1_gpp_proto_rawDescData
}

var file_gpp_v1_gpp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gpp_v1_gpp_proto_goTypes = []interface{}{
	(*Game)(nil), // 0: gpp.Game
	(*User)(nil), // 1: gpp.User
	(*Info)(nil), // 2: gpp.Info
}
var file_gpp_v1_gpp_proto_depIdxs = []int32{
	0, // 0: gpp.Info.game:type_name -> gpp.Game
	1, // 1: gpp.Info.user:type_name -> gpp.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gpp_v1_gpp_proto_init() }
func file_gpp_v1_gpp_proto_init() {
	if File_gpp_v1_gpp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gpp_v1_gpp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_gpp_v1_gpp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_gpp_v1_gpp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
			RawDescriptor: file_gpp_v1_gpp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gpp_v1_gpp_proto_goTypes,
		DependencyIndexes: file_gpp_v1_gpp_proto_depIdxs,
		MessageInfos:      file_gpp_v1_gpp_proto_msgTypes,
	}.Build()
	File_gpp_v1_gpp_proto = out.File
	file_gpp_v1_gpp_proto_rawDesc = nil
	file_gpp_v1_gpp_proto_goTypes = nil
	file_gpp_v1_gpp_proto_depIdxs = nil
}
