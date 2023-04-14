// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/page/page.proto

package page

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

type NewListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListKey string `protobuf:"bytes,1,opt,name=listKey,proto3" json:"listKey,omitempty"`
	UserID  int64  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *NewListRequest) Reset() {
	*x = NewListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewListRequest) ProtoMessage() {}

func (x *NewListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewListRequest.ProtoReflect.Descriptor instead.
func (*NewListRequest) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{0}
}

func (x *NewListRequest) GetListKey() string {
	if x != nil {
		return x.ListKey
	}
	return ""
}

func (x *NewListRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type NewListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *NewListResponse) Reset() {
	*x = NewListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewListResponse) ProtoMessage() {}

func (x *NewListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewListResponse.ProtoReflect.Descriptor instead.
func (*NewListResponse) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{1}
}

func (x *NewListResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetHeadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListKey string `protobuf:"bytes,1,opt,name=listKey,proto3" json:"listKey,omitempty"`
	UserID  int64  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetHeadRequest) Reset() {
	*x = GetHeadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeadRequest) ProtoMessage() {}

func (x *GetHeadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeadRequest.ProtoReflect.Descriptor instead.
func (*GetHeadRequest) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{2}
}

func (x *GetHeadRequest) GetListKey() string {
	if x != nil {
		return x.ListKey
	}
	return ""
}

func (x *GetHeadRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetHeadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageKey string `protobuf:"bytes,1,opt,name=pageKey,proto3" json:"pageKey,omitempty"`
}

func (x *GetHeadResponse) Reset() {
	*x = GetHeadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeadResponse) ProtoMessage() {}

func (x *GetHeadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeadResponse.ProtoReflect.Descriptor instead.
func (*GetHeadResponse) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{3}
}

func (x *GetHeadResponse) GetPageKey() string {
	if x != nil {
		return x.PageKey
	}
	return ""
}

type GetPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListKey string `protobuf:"bytes,1,opt,name=listKey,proto3" json:"listKey,omitempty"`
	UserID  int64  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	PageKey string `protobuf:"bytes,3,opt,name=pageKey,proto3" json:"pageKey,omitempty"`
}

func (x *GetPageRequest) Reset() {
	*x = GetPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageRequest) ProtoMessage() {}

func (x *GetPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageRequest.ProtoReflect.Descriptor instead.
func (*GetPageRequest) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{4}
}

func (x *GetPageRequest) GetListKey() string {
	if x != nil {
		return x.ListKey
	}
	return ""
}

func (x *GetPageRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetPageRequest) GetPageKey() string {
	if x != nil {
		return x.PageKey
	}
	return ""
}

type GetPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	PageContent string `protobuf:"bytes,2,opt,name=pageContent,proto3" json:"pageContent,omitempty"`
	Next        string `protobuf:"bytes,3,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *GetPageResponse) Reset() {
	*x = GetPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageResponse) ProtoMessage() {}

func (x *GetPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageResponse.ProtoReflect.Descriptor instead.
func (*GetPageResponse) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{5}
}

func (x *GetPageResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetPageResponse) GetPageContent() string {
	if x != nil {
		return x.PageContent
	}
	return ""
}

func (x *GetPageResponse) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

type SetPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ListKey     string `protobuf:"bytes,2,opt,name=listKey,proto3" json:"listKey,omitempty"`
	PageContent string `protobuf:"bytes,3,opt,name=pageContent,proto3" json:"pageContent,omitempty"`
}

func (x *SetPageRequest) Reset() {
	*x = SetPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPageRequest) ProtoMessage() {}

func (x *SetPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPageRequest.ProtoReflect.Descriptor instead.
func (*SetPageRequest) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{6}
}

func (x *SetPageRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SetPageRequest) GetListKey() string {
	if x != nil {
		return x.ListKey
	}
	return ""
}

func (x *SetPageRequest) GetPageContent() string {
	if x != nil {
		return x.PageContent
	}
	return ""
}

type SetPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageKey string `protobuf:"bytes,1,opt,name=pageKey,proto3" json:"pageKey,omitempty"`
}

func (x *SetPageResponse) Reset() {
	*x = SetPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_page_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPageResponse) ProtoMessage() {}

func (x *SetPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_page_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPageResponse.ProtoReflect.Descriptor instead.
func (*SetPageResponse) Descriptor() ([]byte, []int) {
	return file_proto_page_page_proto_rawDescGZIP(), []int{7}
}

func (x *SetPageResponse) GetPageKey() string {
	if x != nil {
		return x.PageKey
	}
	return ""
}

var File_proto_page_page_proto protoreflect.FileDescriptor

var file_proto_page_page_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x42, 0x0a,
	0x0e, 0x4e, 0x65, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x29, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x42, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x5c, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x59, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x22, 0x64, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x0f,
	0x53, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x32, 0xfd, 0x01, 0x0a, 0x0b, 0x50, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x4e, 0x65, 0x77,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x4e, 0x65, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x12, 0x14,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x07, 0x53,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x74,
	0x70, 0x6f, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_page_page_proto_rawDescOnce sync.Once
	file_proto_page_page_proto_rawDescData = file_proto_page_page_proto_rawDesc
)

func file_proto_page_page_proto_rawDescGZIP() []byte {
	file_proto_page_page_proto_rawDescOnce.Do(func() {
		file_proto_page_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_page_page_proto_rawDescData)
	})
	return file_proto_page_page_proto_rawDescData
}

var file_proto_page_page_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_page_page_proto_goTypes = []interface{}{
	(*NewListRequest)(nil),  // 0: page.NewListRequest
	(*NewListResponse)(nil), // 1: page.NewListResponse
	(*GetHeadRequest)(nil),  // 2: page.GetHeadRequest
	(*GetHeadResponse)(nil), // 3: page.GetHeadResponse
	(*GetPageRequest)(nil),  // 4: page.GetPageRequest
	(*GetPageResponse)(nil), // 5: page.GetPageResponse
	(*SetPageRequest)(nil),  // 6: page.SetPageRequest
	(*SetPageResponse)(nil), // 7: page.SetPageResponse
}
var file_proto_page_page_proto_depIdxs = []int32{
	0, // 0: page.PageService.NewList:input_type -> page.NewListRequest
	2, // 1: page.PageService.GetHead:input_type -> page.GetHeadRequest
	4, // 2: page.PageService.GetPage:input_type -> page.GetPageRequest
	6, // 3: page.PageService.SetPage:input_type -> page.SetPageRequest
	1, // 4: page.PageService.NewList:output_type -> page.NewListResponse
	3, // 5: page.PageService.GetHead:output_type -> page.GetHeadResponse
	5, // 6: page.PageService.GetPage:output_type -> page.GetPageResponse
	7, // 7: page.PageService.SetPage:output_type -> page.SetPageResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_page_page_proto_init() }
func file_proto_page_page_proto_init() {
	if File_proto_page_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_page_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewListRequest); i {
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
		file_proto_page_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewListResponse); i {
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
		file_proto_page_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeadRequest); i {
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
		file_proto_page_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeadResponse); i {
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
		file_proto_page_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageRequest); i {
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
		file_proto_page_page_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageResponse); i {
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
		file_proto_page_page_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPageRequest); i {
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
		file_proto_page_page_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPageResponse); i {
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
			RawDescriptor: file_proto_page_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_page_page_proto_goTypes,
		DependencyIndexes: file_proto_page_page_proto_depIdxs,
		MessageInfos:      file_proto_page_page_proto_msgTypes,
	}.Build()
	File_proto_page_page_proto = out.File
	file_proto_page_page_proto_rawDesc = nil
	file_proto_page_page_proto_goTypes = nil
	file_proto_page_page_proto_depIdxs = nil
}
