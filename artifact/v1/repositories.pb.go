// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: artifact/v1/repositories.proto

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

type RepositoriesListRequest struct {
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

func (x *RepositoriesListRequest) Reset() {
	*x = RepositoriesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_repositories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoriesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoriesListRequest) ProtoMessage() {}

func (x *RepositoriesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_repositories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoriesListRequest.ProtoReflect.Descriptor instead.
func (*RepositoriesListRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_repositories_proto_rawDescGZIP(), []int{0}
}

func (x *RepositoriesListRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *RepositoriesListRequest) GetPaginateQuery() *general.PaginateQuery {
	if x != nil {
		return x.PaginateQuery
	}
	return nil
}

func (x *RepositoriesListRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

type RepositoriesCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	GppInfo *general.GppInfo `protobuf:"bytes,1,opt,name=GppInfo,proto3" json:"GppInfo,omitempty" validate:"required"`
	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,2,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
	// @gotags: validate:"required"
	Request *RepositoriesRequest `protobuf:"bytes,3,opt,name=Request,proto3" json:"Request,omitempty" validate:"required"`
}

func (x *RepositoriesCreateRequest) Reset() {
	*x = RepositoriesCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_repositories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoriesCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoriesCreateRequest) ProtoMessage() {}

func (x *RepositoriesCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_repositories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoriesCreateRequest.ProtoReflect.Descriptor instead.
func (*RepositoriesCreateRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_repositories_proto_rawDescGZIP(), []int{1}
}

func (x *RepositoriesCreateRequest) GetGppInfo() *general.GppInfo {
	if x != nil {
		return x.GppInfo
	}
	return nil
}

func (x *RepositoriesCreateRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *RepositoriesCreateRequest) GetRequest() *RepositoriesRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type RepositoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty" validate:"required"`
	// @gotags: validate:"required,oneof=image binary"
	RepositoriesType string `protobuf:"bytes,2,opt,name=RepositoriesType,proto3" json:"RepositoriesType,omitempty" validate:"required,oneof=image binary"`
	// @gotags: validate:"required,oneof=dev test staging"
	Branch string `protobuf:"bytes,3,opt,name=Branch,proto3" json:"Branch,omitempty" validate:"required,oneof=dev test staging"`
}

func (x *RepositoriesRequest) Reset() {
	*x = RepositoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_repositories_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoriesRequest) ProtoMessage() {}

func (x *RepositoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_repositories_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoriesRequest.ProtoReflect.Descriptor instead.
func (*RepositoriesRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_repositories_proto_rawDescGZIP(), []int{2}
}

func (x *RepositoriesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoriesRequest) GetRepositoriesType() string {
	if x != nil {
		return x.RepositoriesType
	}
	return ""
}

func (x *RepositoriesRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type RepositoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @gotags: json:"created_at"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"created_at"`
	// @gotags: json:"updated_at"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=UpdatedAt,proto3" json:"updated_at"`
	// @gotags: json:"name"
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"name"`
	// @gotags: json:"repositories_address"
	RepositoriesAddress string `protobuf:"bytes,5,opt,name=RepositoriesAddress,proto3" json:"repositories_address"`
}

func (x *RepositoriesResponse) Reset() {
	*x = RepositoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_repositories_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoriesResponse) ProtoMessage() {}

func (x *RepositoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_repositories_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoriesResponse.ProtoReflect.Descriptor instead.
func (*RepositoriesResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_repositories_proto_rawDescGZIP(), []int{3}
}

func (x *RepositoriesResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RepositoriesResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RepositoriesResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *RepositoriesResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoriesResponse) GetRepositoriesAddress() string {
	if x != nil {
		return x.RepositoriesAddress
	}
	return ""
}

type RepositoriesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"paginate"
	Paginate *general.PaginateInfo `protobuf:"bytes,1,opt,name=Paginate,proto3" json:"paginate"`
	// @gotags: json:"items"
	Items []*RepositoriesResponse `protobuf:"bytes,2,rep,name=Items,proto3" json:"items"`
}

func (x *RepositoriesListResponse) Reset() {
	*x = RepositoriesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_repositories_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoriesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoriesListResponse) ProtoMessage() {}

func (x *RepositoriesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_repositories_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoriesListResponse.ProtoReflect.Descriptor instead.
func (*RepositoriesListResponse) Descriptor() ([]byte, []int) {
	return file_artifact_v1_repositories_proto_rawDescGZIP(), []int{4}
}

func (x *RepositoriesListResponse) GetPaginate() *general.PaginateInfo {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *RepositoriesListResponse) GetItems() []*RepositoriesResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type RepositoriesGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	RequestID *general.RequestID `protobuf:"bytes,1,opt,name=RequestID,proto3" json:"RequestID,omitempty" validate:"required"`
	// @gotags: validate:"required"
	RepositoriesName string `protobuf:"bytes,2,opt,name=RepositoriesName,proto3" json:"RepositoriesName,omitempty" validate:"required"`
}

func (x *RepositoriesGetRequest) Reset() {
	*x = RepositoriesGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_v1_repositories_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoriesGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoriesGetRequest) ProtoMessage() {}

func (x *RepositoriesGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_v1_repositories_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoriesGetRequest.ProtoReflect.Descriptor instead.
func (*RepositoriesGetRequest) Descriptor() ([]byte, []int) {
	return file_artifact_v1_repositories_proto_rawDescGZIP(), []int{5}
}

func (x *RepositoriesGetRequest) GetRequestID() *general.RequestID {
	if x != nil {
		return x.RequestID
	}
	return nil
}

func (x *RepositoriesGetRequest) GetRepositoriesName() string {
	if x != nil {
		return x.RepositoriesName
	}
	return ""
}

var File_artifact_v1_repositories_proto protoreflect.FileDescriptor

var file_artifact_v1_repositories_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x1a, 0x21, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x64,
	0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2f, 0x67, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c,
	0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x0d, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x09,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x22, 0xb2,
	0x01, 0x0a, 0x19, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07,
	0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x47, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x22, 0xe0, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x76, 0x0a, 0x16, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x52, 0x09, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x32, 0x82, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x64, 0x65, 0x76,
	0x2f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_v1_repositories_proto_rawDescOnce sync.Once
	file_artifact_v1_repositories_proto_rawDescData = file_artifact_v1_repositories_proto_rawDesc
)

func file_artifact_v1_repositories_proto_rawDescGZIP() []byte {
	file_artifact_v1_repositories_proto_rawDescOnce.Do(func() {
		file_artifact_v1_repositories_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_v1_repositories_proto_rawDescData)
	})
	return file_artifact_v1_repositories_proto_rawDescData
}

var file_artifact_v1_repositories_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_artifact_v1_repositories_proto_goTypes = []interface{}{
	(*RepositoriesListRequest)(nil),   // 0: artifact.RepositoriesListRequest
	(*RepositoriesCreateRequest)(nil), // 1: artifact.RepositoriesCreateRequest
	(*RepositoriesRequest)(nil),       // 2: artifact.RepositoriesRequest
	(*RepositoriesResponse)(nil),      // 3: artifact.RepositoriesResponse
	(*RepositoriesListResponse)(nil),  // 4: artifact.RepositoriesListResponse
	(*RepositoriesGetRequest)(nil),    // 5: artifact.RepositoriesGetRequest
	(*general.GppInfo)(nil),           // 6: general.GppInfo
	(*general.PaginateQuery)(nil),     // 7: general.PaginateQuery
	(*general.RequestID)(nil),         // 8: general.RequestID
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
	(*general.PaginateInfo)(nil),      // 10: general.PaginateInfo
}
var file_artifact_v1_repositories_proto_depIdxs = []int32{
	6,  // 0: artifact.RepositoriesListRequest.GppInfo:type_name -> general.GppInfo
	7,  // 1: artifact.RepositoriesListRequest.PaginateQuery:type_name -> general.PaginateQuery
	8,  // 2: artifact.RepositoriesListRequest.RequestID:type_name -> general.RequestID
	6,  // 3: artifact.RepositoriesCreateRequest.GppInfo:type_name -> general.GppInfo
	8,  // 4: artifact.RepositoriesCreateRequest.RequestID:type_name -> general.RequestID
	2,  // 5: artifact.RepositoriesCreateRequest.Request:type_name -> artifact.RepositoriesRequest
	9,  // 6: artifact.RepositoriesResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	9,  // 7: artifact.RepositoriesResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	10, // 8: artifact.RepositoriesListResponse.Paginate:type_name -> general.PaginateInfo
	3,  // 9: artifact.RepositoriesListResponse.Items:type_name -> artifact.RepositoriesResponse
	8,  // 10: artifact.RepositoriesGetRequest.RequestID:type_name -> general.RequestID
	0,  // 11: artifact.RepositoriesService.List:input_type -> artifact.RepositoriesListRequest
	1,  // 12: artifact.RepositoriesService.Create:input_type -> artifact.RepositoriesCreateRequest
	5,  // 13: artifact.RepositoriesService.Get:input_type -> artifact.RepositoriesGetRequest
	4,  // 14: artifact.RepositoriesService.List:output_type -> artifact.RepositoriesListResponse
	3,  // 15: artifact.RepositoriesService.Create:output_type -> artifact.RepositoriesResponse
	3,  // 16: artifact.RepositoriesService.Get:output_type -> artifact.RepositoriesResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_artifact_v1_repositories_proto_init() }
func file_artifact_v1_repositories_proto_init() {
	if File_artifact_v1_repositories_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_v1_repositories_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoriesListRequest); i {
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
		file_artifact_v1_repositories_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoriesCreateRequest); i {
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
		file_artifact_v1_repositories_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoriesRequest); i {
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
		file_artifact_v1_repositories_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoriesResponse); i {
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
		file_artifact_v1_repositories_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoriesListResponse); i {
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
		file_artifact_v1_repositories_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoriesGetRequest); i {
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
			RawDescriptor: file_artifact_v1_repositories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_artifact_v1_repositories_proto_goTypes,
		DependencyIndexes: file_artifact_v1_repositories_proto_depIdxs,
		MessageInfos:      file_artifact_v1_repositories_proto_msgTypes,
	}.Build()
	File_artifact_v1_repositories_proto = out.File
	file_artifact_v1_repositories_proto_rawDesc = nil
	file_artifact_v1_repositories_proto_goTypes = nil
	file_artifact_v1_repositories_proto_depIdxs = nil
}
