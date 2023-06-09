// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: general/paginate.proto

package general

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

type PaginateQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64             `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`                                                                                            // 页码
	PageSize int64             `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`                                                                                    // 每页大小
	Order    string            `protobuf:"bytes,3,opt,name=Order,proto3" json:"Order,omitempty"`                                                                                           // 排序
	OrderBy  string            `protobuf:"bytes,4,opt,name=OrderBy,proto3" json:"OrderBy,omitempty"`                                                                                       // 排序对象
	Search   string            `protobuf:"bytes,5,opt,name=Search,proto3" json:"Search,omitempty"`                                                                                         // 模糊搜索
	Params   map[string]*Param `protobuf:"bytes,6,rep,name=Params,proto3" json:"Params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // query对象字典
	AllData  bool              `protobuf:"varint,7,opt,name=AllData,proto3" json:"AllData,omitempty"`                                                                                      // 是否查询全量数据
}

func (x *PaginateQuery) Reset() {
	*x = PaginateQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_paginate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginateQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginateQuery) ProtoMessage() {}

func (x *PaginateQuery) ProtoReflect() protoreflect.Message {
	mi := &file_general_paginate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginateQuery.ProtoReflect.Descriptor instead.
func (*PaginateQuery) Descriptor() ([]byte, []int) {
	return file_general_paginate_proto_rawDescGZIP(), []int{0}
}

func (x *PaginateQuery) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginateQuery) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginateQuery) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *PaginateQuery) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *PaginateQuery) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *PaginateQuery) GetParams() map[string]*Param {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *PaginateQuery) GetAllData() bool {
	if x != nil {
		return x.AllData
	}
	return false
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=Values,proto3" json:"Values,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_paginate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_general_paginate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_general_paginate_proto_rawDescGZIP(), []int{1}
}

func (x *Param) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type PaginateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page int64 `protobuf:"varint,1,opt,name=Page,proto3" json:"page"` // 页码
	// @gotags: json:"page_size"
	PageSize int64 `protobuf:"varint,2,opt,name=PageSize,proto3" json:"page_size"` // 每页大小
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,3,opt,name=Total,proto3" json:"total"` // 数据总数
}

func (x *PaginateInfo) Reset() {
	*x = PaginateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_paginate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginateInfo) ProtoMessage() {}

func (x *PaginateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_general_paginate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginateInfo.ProtoReflect.Descriptor instead.
func (*PaginateInfo) Descriptor() ([]byte, []int) {
	return file_general_paginate_proto_rawDescGZIP(), []int{2}
}

func (x *PaginateInfo) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginateInfo) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginateInfo) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_general_paginate_proto protoreflect.FileDescriptor

var file_general_paginate_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x22, 0xa8, 0x02, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x3a, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x49, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x05,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x54, 0x0a,
	0x0c, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x59, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_general_paginate_proto_rawDescOnce sync.Once
	file_general_paginate_proto_rawDescData = file_general_paginate_proto_rawDesc
)

func file_general_paginate_proto_rawDescGZIP() []byte {
	file_general_paginate_proto_rawDescOnce.Do(func() {
		file_general_paginate_proto_rawDescData = protoimpl.X.CompressGZIP(file_general_paginate_proto_rawDescData)
	})
	return file_general_paginate_proto_rawDescData
}

var file_general_paginate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_general_paginate_proto_goTypes = []interface{}{
	(*PaginateQuery)(nil), // 0: general.PaginateQuery
	(*Param)(nil),         // 1: general.Param
	(*PaginateInfo)(nil),  // 2: general.PaginateInfo
	nil,                   // 3: general.PaginateQuery.ParamsEntry
}
var file_general_paginate_proto_depIdxs = []int32{
	3, // 0: general.PaginateQuery.Params:type_name -> general.PaginateQuery.ParamsEntry
	1, // 1: general.PaginateQuery.ParamsEntry.value:type_name -> general.Param
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_general_paginate_proto_init() }
func file_general_paginate_proto_init() {
	if File_general_paginate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_general_paginate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginateQuery); i {
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
		file_general_paginate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_general_paginate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginateInfo); i {
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
			RawDescriptor: file_general_paginate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_general_paginate_proto_goTypes,
		DependencyIndexes: file_general_paginate_proto_depIdxs,
		MessageInfos:      file_general_paginate_proto_msgTypes,
	}.Build()
	File_general_paginate_proto = out.File
	file_general_paginate_proto_rawDesc = nil
	file_general_paginate_proto_goTypes = nil
	file_general_paginate_proto_depIdxs = nil
}
