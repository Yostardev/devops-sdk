// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: artifact/v1/products.proto

package artifact

import (
	general "github.com/Yostardev/devops-sdk/general"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	GppInfo *general.GppInfo `protobuf:"bytes,1,opt,name=GppInfo,proto3" json:"GppInfo,omitempty" validate:"required"`
	// @gotags: validate:"required"
	PaginateQuery *general.PaginateQuery `protobuf:"bytes,2,opt,name=PaginateQuery,proto3" json:"PaginateQuery,omitempty" validate:"required"`
	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,3,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
}

func (x *ProductListRequest) Reset() {
	*x = ProductListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListRequest) ProtoMessage() {}

func (x *ProductListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListRequest.ProtoReflect.Descriptor instead.
func (*ProductListRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{0}
}

func (x *ProductListRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *ProductListRequest) GetPaginateQuery() *general.PaginateQuery {
	if x != nil {
		return x.PaginateQuery
	}
	return nil
}

func (x *ProductListRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"created_at"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=CreatedAt,proto3" json:"created_at"`
	// @gotags: json:"updated_at"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=UpdatedAt,proto3" json:"updated_at"`
	// @gotags: json:"id"
	ID int64 `protobuf:"varint,3,opt,name=ID,proto3" json:"id"`
	// @gotags: json:"name"
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"name"`
	// @gotags: json:"repositories_id"
	RepositoriesID int64 `protobuf:"varint,5,opt,name=RepositoriesID,proto3" json:"repositories_id"`
	// @gotags: json:"branch"
	Branch string `protobuf:"bytes,6,opt,name=Branch,proto3" json:"branch"`
	// @gotags: json:"url"
	Url string `protobuf:"bytes,7,opt,name=Url,proto3" json:"url"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProductResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProductResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetRepositoriesID() int64 {
	if x != nil {
		return x.RepositoriesID
	}
	return 0
}

func (x *ProductResponse) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *ProductResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ProductListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"paginate"
	Paginate *general.PaginateInfo `protobuf:"bytes,1,opt,name=Paginate,proto3" json:"paginate"`
	// @gotags: json:"items"
	Items []*ProductResponse `protobuf:"bytes,2,rep,name=Items,proto3" json:"items"`
}

func (x *ProductListResponse) Reset() {
	*x = ProductListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListResponse) ProtoMessage() {}

func (x *ProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListResponse.ProtoReflect.Descriptor instead.
func (*ProductListResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{2}
}

func (x *ProductListResponse) GetPaginate() *general.PaginateInfo {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *ProductListResponse) GetItems() []*ProductResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" validate:"required"`
	// @gotags: json:"branch" validate:"required"
	Branch string `protobuf:"bytes,2,opt,name=Branch,proto3" json:"branch" validate:"required"`
	// @gotags: json:"type" validate:"required,oneof=image binary"
	Type string `protobuf:"bytes,3,opt,name=Type,proto3" json:"type" validate:"required,oneof=image binary"`
	// @gotags: json:"repositories_name" validate:"required"
	RepositoriesName string `protobuf:"bytes,4,opt,name=RepositoriesName,proto3" json:"repositories_name" validate:"required"`
	// @gotags: json:"url" validate:"required"
	Url string `protobuf:"bytes,5,opt,name=Url,proto3" json:"url" validate:"required"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{3}
}

func (x *ProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *ProductRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ProductRequest) GetRepositoriesName() string {
	if x != nil {
		return x.RepositoriesName
	}
	return ""
}

func (x *ProductRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ProductCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	GppInfo *general.GppInfo `protobuf:"bytes,1,opt,name=GppInfo,proto3" json:"GppInfo,omitempty" validate:"required"`
	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,2,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Request *ProductRequest `protobuf:"bytes,3,opt,name=Request,proto3" json:"Request,omitempty" validate:"required"`
}

func (x *ProductCreateRequest) Reset() {
	*x = ProductCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCreateRequest) ProtoMessage() {}

func (x *ProductCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCreateRequest.ProtoReflect.Descriptor instead.
func (*ProductCreateRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{4}
}

func (x *ProductCreateRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *ProductCreateRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *ProductCreateRequest) GetRequest() *ProductRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type ProductDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	GppInfo *general.GppInfo `protobuf:"bytes,1,opt,name=GppInfo,proto3" json:"GppInfo,omitempty" validate:"required"`
	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,2,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
	// @gotags: validate:"required"
	ID int64 `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty" validate:"required"`
}

func (x *ProductDownloadRequest) Reset() {
	*x = ProductDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDownloadRequest) ProtoMessage() {}

func (x *ProductDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDownloadRequest.ProtoReflect.Descriptor instead.
func (*ProductDownloadRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{5}
}

func (x *ProductDownloadRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *ProductDownloadRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *ProductDownloadRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ProductDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"data"
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"data"`
}

func (x *ProductDownloadResponse) Reset() {
	*x = ProductDownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDownloadResponse) ProtoMessage() {}

func (x *ProductDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDownloadResponse.ProtoReflect.Descriptor instead.
func (*ProductDownloadResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{6}
}

func (x *ProductDownloadResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProductObjectMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty" validate:"required"`
	// @gotags: validate:"required"
	GppInfo *general.GppInfo `protobuf:"bytes,2,opt,name=GppInfo,proto3" json:"GppInfo,omitempty" validate:"required"`
	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,3,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
}

func (x *ProductObjectMetaRequest) Reset() {
	*x = ProductObjectMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductObjectMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductObjectMetaRequest) ProtoMessage() {}

func (x *ProductObjectMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductObjectMetaRequest.ProtoReflect.Descriptor instead.
func (*ProductObjectMetaRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{7}
}

func (x *ProductObjectMetaRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductObjectMetaRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *ProductObjectMetaRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

type ProductObjectMetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"HTTPHeaderContentLength"
	HTTPHeaderContentLength string `protobuf:"bytes,1,opt,name=HTTPHeaderContentLength,proto3" json:"HTTPHeaderContentLength"`
	// @gotags: json:"file_name"
	FileName string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"file_name"`
}

func (x *ProductObjectMetaResponse) Reset() {
	*x = ProductObjectMetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_products_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductObjectMetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductObjectMetaResponse) ProtoMessage() {}

func (x *ProductObjectMetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_products_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductObjectMetaResponse.ProtoReflect.Descriptor instead.
func (*ProductObjectMetaResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_products_proto_rawDescGZIP(), []int{8}
}

func (x *ProductObjectMetaResponse) GetHTTPHeaderContentLength() string {
	if x != nil {
		return x.HTTPHeaderContentLength
	}
	return ""
}

func (x *ProductObjectMetaResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_artifact_v1_products_proto protoreflect.FileDescriptor

var file_artifact_v1_products_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x1a, 0x21, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x64, 0x65, 0x76, 0x6f, 0x70,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x70,
	0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x64, 0x65,
	0x76, 0x6f, 0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x47, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x44, 0x22, 0xfb, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x55, 0x72, 0x6c, 0x22, 0x79, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x8e,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x22,
	0xa8, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x47, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x2e, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x22, 0x71, 0x0a,
	0x19, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x48, 0x54,
	0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x48, 0x54, 0x54,
	0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x32, 0xd3, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x22, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x65, 0x76, 0x2f, 0x64,
	0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_v1_products_proto_rawDescOnce sync.Once
	file_artifact_v1_products_proto_rawDescData = file_artifact_v1_products_proto_rawDesc
)

func file_artifact_v1_products_proto_rawDescGZIP() []byte {
	file_artifact_v1_products_proto_rawDescOnce.Do(func() {
		file_artifact_v1_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_v1_products_proto_rawDescData)
	})
	return file_artifact_v1_products_proto_rawDescData
}

var file_artifact_v1_products_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_artifact_v1_products_proto_goTypes = []interface{}{
	(*ProductListRequest)(nil),        // 0: artifact.ProductListRequest
	(*ProductResponse)(nil),           // 1: artifact.ProductResponse
	(*ProductListResponse)(nil),       // 2: artifact.ProductListResponse
	(*ProductRequest)(nil),            // 3: artifact.ProductRequest
	(*ProductCreateRequest)(nil),      // 4: artifact.ProductCreateRequest
	(*ProductDownloadRequest)(nil),    // 5: artifact.ProductDownloadRequest
	(*ProductDownloadResponse)(nil),   // 6: artifact.ProductDownloadResponse
	(*ProductObjectMetaRequest)(nil),  // 7: artifact.ProductObjectMetaRequest
	(*ProductObjectMetaResponse)(nil), // 8: artifact.ProductObjectMetaResponse
	(*general.GppInfo)(nil),           // 9: general.GppInfo
	(*general.PaginateQuery)(nil),     // 10: general.PaginateQuery
	(*general.RequestID)(nil),         // 11: general.RequestID
	(*timestamppb.Timestamp)(nil),     // 12: google.protobuf.Timestamp
	(*general.PaginateInfo)(nil),      // 13: general.PaginateInfo
}
var file_artifact_v1_products_proto_depIdxs = []int32{
	9,  // 0: artifact.ProductListRequest.GppInfo:type_name -> general.GppInfo
	10, // 1: artifact.ProductListRequest.PaginateQuery:type_name -> general.PaginateQuery
	11, // 2: artifact.ProductListRequest.RequestID:type_name -> general.RequestID
	12, // 3: artifact.ProductResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	12, // 4: artifact.ProductResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	13, // 5: artifact.ProductListResponse.Paginate:type_name -> general.PaginateInfo
	1,  // 6: artifact.ProductListResponse.Items:type_name -> artifact.ProductResponse
	9,  // 7: artifact.ProductCreateRequest.GppInfo:type_name -> general.GppInfo
	11, // 8: artifact.ProductCreateRequest.RequestID:type_name -> general.RequestID
	3,  // 9: artifact.ProductCreateRequest.Request:type_name -> artifact.ProductRequest
	9,  // 10: artifact.ProductDownloadRequest.GppInfo:type_name -> general.GppInfo
	11, // 11: artifact.ProductDownloadRequest.RequestID:type_name -> general.RequestID
	9,  // 12: artifact.ProductObjectMetaRequest.GppInfo:type_name -> general.GppInfo
	11, // 13: artifact.ProductObjectMetaRequest.RequestID:type_name -> general.RequestID
	0,  // 14: artifact.ProductService.List:input_type -> artifact.ProductListRequest
	4,  // 15: artifact.ProductService.Create:input_type -> artifact.ProductCreateRequest
	7,  // 16: artifact.ProductService.GetObjectMeta:input_type -> artifact.ProductObjectMetaRequest
	5,  // 17: artifact.ProductService.DownloadFile:input_type -> artifact.ProductDownloadRequest
	2,  // 18: artifact.ProductService.List:output_type -> artifact.ProductListResponse
	1,  // 19: artifact.ProductService.Create:output_type -> artifact.ProductResponse
	8,  // 20: artifact.ProductService.GetObjectMeta:output_type -> artifact.ProductObjectMetaResponse
	6,  // 21: artifact.ProductService.DownloadFile:output_type -> artifact.ProductDownloadResponse
	18, // [18:22] is the sub-list for method output_type
	14, // [14:18] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_artifact_v1_products_proto_init() }
func file_artifact_v1_products_proto_init() {
	if File_artifact_v1_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_v1_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListRequest); i {
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
		file_artifact_v1_products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
		file_artifact_v1_products_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductListResponse); i {
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
		file_artifact_v1_products_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
		file_artifact_v1_products_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCreateRequest); i {
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
		file_artifact_v1_products_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDownloadRequest); i {
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
		file_artifact_v1_products_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDownloadResponse); i {
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
		file_artifact_v1_products_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductObjectMetaRequest); i {
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
		file_artifact_v1_products_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductObjectMetaResponse); i {
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
			RawDescriptor: file_artifact_v1_products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_artifact_v1_products_proto_goTypes,
		DependencyIndexes: file_artifact_v1_products_proto_depIdxs,
		MessageInfos:      file_artifact_v1_products_proto_msgTypes,
	}.Build()
	File_artifact_v1_products_proto = out.File
	file_artifact_v1_products_proto_rawDesc = nil
	file_artifact_v1_products_proto_goTypes = nil
	file_artifact_v1_products_proto_depIdxs = nil
}
