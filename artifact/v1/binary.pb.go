// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: artifact/v1/binary.proto

package artifact

import (
	general "github.com/Yostardev/devops-sdk/general"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type BinaryListRequest struct {
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

func (x *BinaryListRequest) Reset() {
	*x = BinaryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryListRequest) ProtoMessage() {}

func (x *BinaryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryListRequest.ProtoReflect.Descriptor instead.
func (*BinaryListRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{0}
}

func (x *BinaryListRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *BinaryListRequest) GetPaginateQuery() *general.PaginateQuery {
	if x != nil {
		return x.PaginateQuery
	}
	return nil
}

func (x *BinaryListRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

type BinaryResponse struct {
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
	// @gotags: json:"relative_path"
	RelativePath string `protobuf:"bytes,7,opt,name=RelativePath,proto3" json:"relative_path"`
}

func (x *BinaryResponse) Reset() {
	*x = BinaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryResponse) ProtoMessage() {}

func (x *BinaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryResponse.ProtoReflect.Descriptor instead.
func (*BinaryResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{1}
}

func (x *BinaryResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BinaryResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *BinaryResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BinaryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BinaryResponse) GetRepositoriesID() int64 {
	if x != nil {
		return x.RepositoriesID
	}
	return 0
}

func (x *BinaryResponse) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *BinaryResponse) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

type BinaryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"paginate"
	Paginate *general.PaginateInfo `protobuf:"bytes,1,opt,name=Paginate,proto3" json:"paginate"`
	// @gotags: json:"items"
	Items []*BinaryResponse `protobuf:"bytes,2,rep,name=Items,proto3" json:"items"`
}

func (x *BinaryListResponse) Reset() {
	*x = BinaryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryListResponse) ProtoMessage() {}

func (x *BinaryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryListResponse.ProtoReflect.Descriptor instead.
func (*BinaryListResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{2}
}

func (x *BinaryListResponse) GetPaginate() *general.PaginateInfo {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *BinaryListResponse) GetItems() []*BinaryResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type BinaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name" validate:"required"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" validate:"required"`
	// @gotags: json:"branch" validate:"required"
	Branch string `protobuf:"bytes,2,opt,name=Branch,proto3" json:"branch" validate:"required"`
	// @gotags: json:"repositories_name" validate:"required"
	RepositoriesName string `protobuf:"bytes,4,opt,name=RepositoriesName,proto3" json:"repositories_name" validate:"required"`
}

func (x *BinaryRequest) Reset() {
	*x = BinaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryRequest) ProtoMessage() {}

func (x *BinaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryRequest.ProtoReflect.Descriptor instead.
func (*BinaryRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{3}
}

func (x *BinaryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BinaryRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *BinaryRequest) GetRepositoriesName() string {
	if x != nil {
		return x.RepositoriesName
	}
	return ""
}

type BinaryCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	GppInfo *general.GppInfo `protobuf:"bytes,1,opt,name=GppInfo,proto3" json:"GppInfo,omitempty" validate:"required"`
	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,2,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Request *BinaryRequest `protobuf:"bytes,3,opt,name=Request,proto3" json:"Request,omitempty" validate:"required"`
}

func (x *BinaryCreateRequest) Reset() {
	*x = BinaryCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryCreateRequest) ProtoMessage() {}

func (x *BinaryCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryCreateRequest.ProtoReflect.Descriptor instead.
func (*BinaryCreateRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{4}
}

func (x *BinaryCreateRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *BinaryCreateRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *BinaryCreateRequest) GetRequest() *BinaryRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type BinaryDownloadRequest struct {
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

func (x *BinaryDownloadRequest) Reset() {
	*x = BinaryDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryDownloadRequest) ProtoMessage() {}

func (x *BinaryDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryDownloadRequest.ProtoReflect.Descriptor instead.
func (*BinaryDownloadRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{5}
}

func (x *BinaryDownloadRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *BinaryDownloadRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *BinaryDownloadRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type BinaryDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"data"
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"data"`
}

func (x *BinaryDownloadResponse) Reset() {
	*x = BinaryDownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryDownloadResponse) ProtoMessage() {}

func (x *BinaryDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryDownloadResponse.ProtoReflect.Descriptor instead.
func (*BinaryDownloadResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{6}
}

func (x *BinaryDownloadResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BinaryObjectMetaRequest struct {
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

func (x *BinaryObjectMetaRequest) Reset() {
	*x = BinaryObjectMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryObjectMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryObjectMetaRequest) ProtoMessage() {}

func (x *BinaryObjectMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryObjectMetaRequest.ProtoReflect.Descriptor instead.
func (*BinaryObjectMetaRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{7}
}

func (x *BinaryObjectMetaRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BinaryObjectMetaRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *BinaryObjectMetaRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

type BinaryObjectMetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"HTTPHeaderContentLength"
	HTTPHeaderContentLength string `protobuf:"bytes,1,opt,name=HTTPHeaderContentLength,proto3" json:"HTTPHeaderContentLength"`
	// @gotags: json:"file_name"
	FileName string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"file_name"`
}

func (x *BinaryObjectMetaResponse) Reset() {
	*x = BinaryObjectMetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryObjectMetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryObjectMetaResponse) ProtoMessage() {}

func (x *BinaryObjectMetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryObjectMetaResponse.ProtoReflect.Descriptor instead.
func (*BinaryObjectMetaResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{8}
}

func (x *BinaryObjectMetaResponse) GetHTTPHeaderContentLength() string {
	if x != nil {
		return x.HTTPHeaderContentLength
	}
	return ""
}

func (x *BinaryObjectMetaResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type BinaryObjectDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,1,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
	// @gotags: json:"branch" validate:"required,oneof=dev test"
	Branch string `protobuf:"bytes,3,opt,name=Branch,proto3" json:"branch" validate:"required,oneof=dev test"`
}

func (x *BinaryObjectDeleteRequest) Reset() {
	*x = BinaryObjectDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_binary_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryObjectDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryObjectDeleteRequest) ProtoMessage() {}

func (x *BinaryObjectDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_binary_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryObjectDeleteRequest.ProtoReflect.Descriptor instead.
func (*BinaryObjectDeleteRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_binary_proto_rawDescGZIP(), []int{9}
}

func (x *BinaryObjectDeleteRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *BinaryObjectDeleteRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

var File_artifact_v1_binary_proto protoreflect.FileDescriptor

var file_artifact_v1_binary_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x1a, 0x21, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x67, 0x70, 0x70, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01,
	0x0a, 0x11, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47,
	0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x3c, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0d,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x22,
	0x8c, 0x02, 0x0a, 0x0e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x77,
	0x0a, 0x12, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x67, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xa6, 0x01, 0x0a, 0x13, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x47, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47,
	0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x2c, 0x0a, 0x16, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x87, 0x01, 0x0a, 0x17, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x07, 0x47,
	0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x22, 0x70, 0x0a, 0x18, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x19, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x32, 0x9d, 0x03, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x21, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x51, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x59, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x65, 0x76, 0x6f, 0x70,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_v1_binary_proto_rawDescOnce sync.Once
	file_artifact_v1_binary_proto_rawDescData = file_artifact_v1_binary_proto_rawDesc
)

func file_artifact_v1_binary_proto_rawDescGZIP() []byte {
	file_artifact_v1_binary_proto_rawDescOnce.Do(func() {
		file_artifact_v1_binary_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_v1_binary_proto_rawDescData)
	})
	return file_artifact_v1_binary_proto_rawDescData
}

var file_artifact_v1_binary_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_artifact_v1_binary_proto_goTypes = []interface{}{
	(*BinaryListRequest)(nil),         // 0: artifact.BinaryListRequest
	(*BinaryResponse)(nil),            // 1: artifact.BinaryResponse
	(*BinaryListResponse)(nil),        // 2: artifact.BinaryListResponse
	(*BinaryRequest)(nil),             // 3: artifact.BinaryRequest
	(*BinaryCreateRequest)(nil),       // 4: artifact.BinaryCreateRequest
	(*BinaryDownloadRequest)(nil),     // 5: artifact.BinaryDownloadRequest
	(*BinaryDownloadResponse)(nil),    // 6: artifact.BinaryDownloadResponse
	(*BinaryObjectMetaRequest)(nil),   // 7: artifact.BinaryObjectMetaRequest
	(*BinaryObjectMetaResponse)(nil),  // 8: artifact.BinaryObjectMetaResponse
	(*BinaryObjectDeleteRequest)(nil), // 9: artifact.BinaryObjectDeleteRequest
	(*general.GppInfo)(nil),           // 10: general.GppInfo
	(*general.PaginateQuery)(nil),     // 11: general.PaginateQuery
	(*general.RequestID)(nil),         // 12: general.RequestID
	(*timestamppb.Timestamp)(nil),     // 13: google.protobuf.Timestamp
	(*general.PaginateInfo)(nil),      // 14: general.PaginateInfo
	(*emptypb.Empty)(nil),             // 15: google.protobuf.Empty
}
var file_artifact_v1_binary_proto_depIdxs = []int32{
	10, // 0: artifact.BinaryListRequest.GppInfo:type_name -> general.GppInfo
	11, // 1: artifact.BinaryListRequest.PaginateQuery:type_name -> general.PaginateQuery
	12, // 2: artifact.BinaryListRequest.RequestID:type_name -> general.RequestID
	13, // 3: artifact.BinaryResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	13, // 4: artifact.BinaryResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	14, // 5: artifact.BinaryListResponse.Paginate:type_name -> general.PaginateInfo
	1,  // 6: artifact.BinaryListResponse.Items:type_name -> artifact.BinaryResponse
	10, // 7: artifact.BinaryCreateRequest.GppInfo:type_name -> general.GppInfo
	12, // 8: artifact.BinaryCreateRequest.RequestID:type_name -> general.RequestID
	3,  // 9: artifact.BinaryCreateRequest.Request:type_name -> artifact.BinaryRequest
	10, // 10: artifact.BinaryDownloadRequest.GppInfo:type_name -> general.GppInfo
	12, // 11: artifact.BinaryDownloadRequest.RequestID:type_name -> general.RequestID
	10, // 12: artifact.BinaryObjectMetaRequest.GppInfo:type_name -> general.GppInfo
	12, // 13: artifact.BinaryObjectMetaRequest.RequestID:type_name -> general.RequestID
	12, // 14: artifact.BinaryObjectDeleteRequest.RequestID:type_name -> general.RequestID
	0,  // 15: artifact.BinaryService.List:input_type -> artifact.BinaryListRequest
	4,  // 16: artifact.BinaryService.Create:input_type -> artifact.BinaryCreateRequest
	7,  // 17: artifact.BinaryService.GetObjectMeta:input_type -> artifact.BinaryObjectMetaRequest
	5,  // 18: artifact.BinaryService.DownloadFile:input_type -> artifact.BinaryDownloadRequest
	9,  // 19: artifact.BinaryService.DeleteCollection:input_type -> artifact.BinaryObjectDeleteRequest
	2,  // 20: artifact.BinaryService.List:output_type -> artifact.BinaryListResponse
	1,  // 21: artifact.BinaryService.Create:output_type -> artifact.BinaryResponse
	8,  // 22: artifact.BinaryService.GetObjectMeta:output_type -> artifact.BinaryObjectMetaResponse
	6,  // 23: artifact.BinaryService.DownloadFile:output_type -> artifact.BinaryDownloadResponse
	15, // 24: artifact.BinaryService.DeleteCollection:output_type -> google.protobuf.Empty
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_artifact_v1_binary_proto_init() }
func file_artifact_v1_binary_proto_init() {
	if File_artifact_v1_binary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_v1_binary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryListRequest); i {
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
		file_artifact_v1_binary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryResponse); i {
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
		file_artifact_v1_binary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryListResponse); i {
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
		file_artifact_v1_binary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryRequest); i {
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
		file_artifact_v1_binary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryCreateRequest); i {
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
		file_artifact_v1_binary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryDownloadRequest); i {
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
		file_artifact_v1_binary_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryDownloadResponse); i {
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
		file_artifact_v1_binary_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryObjectMetaRequest); i {
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
		file_artifact_v1_binary_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryObjectMetaResponse); i {
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
		file_artifact_v1_binary_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryObjectDeleteRequest); i {
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
			RawDescriptor: file_artifact_v1_binary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_artifact_v1_binary_proto_goTypes,
		DependencyIndexes: file_artifact_v1_binary_proto_depIdxs,
		MessageInfos:      file_artifact_v1_binary_proto_msgTypes,
	}.Build()
	File_artifact_v1_binary_proto = out.File
	file_artifact_v1_binary_proto_rawDesc = nil
	file_artifact_v1_binary_proto_goTypes = nil
	file_artifact_v1_binary_proto_depIdxs = nil
}