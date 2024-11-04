// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: internal/api/operate/app.proto

package __

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

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 內容ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 內容標題
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 影片URL
	VideoUrl string `protobuf:"bytes,3,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	// 作者
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	// 內容描述
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// 封面URL
	Thumbnail string `protobuf:"bytes,6,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// 內容分類
	Category string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	// 內容時長
	Duration string `protobuf:"bytes,8,opt,name=duration,proto3" json:"duration,omitempty"`
	// 分辨率
	Resolution string `protobuf:"bytes,9,opt,name=resolution,proto3" json:"resolution,omitempty"`
	// 文件大小
	FileSize int64 `protobuf:"varint,10,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// 文件格式，如:MP4、AVI
	Format string `protobuf:"bytes,11,opt,name=format,proto3" json:"format,omitempty"`
	// 影片質量，1-1080P 2-640P
	Quality int32 `protobuf:"varint,12,opt,name=quality,proto3" json:"quality,omitempty"`
	// 審核狀態，1-審核中 2-審核通過 3-審核失敗
	ApprovalStatus int32 `protobuf:"varint,13,opt,name=approval_status,json=approvalStatus,proto3" json:"approval_status,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	mi := &file_internal_api_operate_app_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *Content) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Content) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Content) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Content) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Content) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *Content) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Content) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Content) GetQuality() int32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

func (x *Content) GetApprovalStatus() int32 {
	if x != nil {
		return x.ApprovalStatus
	}
	return 0
}

type CreateContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 內容
	Content *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateContentReq) Reset() {
	*x = CreateContentReq{}
	mi := &file_internal_api_operate_app_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentReq) ProtoMessage() {}

func (x *CreateContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentReq.ProtoReflect.Descriptor instead.
func (*CreateContentReq) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type UpdateContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 內容
	Content *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateContentReq) Reset() {
	*x = UpdateContentReq{}
	mi := &file_internal_api_operate_app_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentReq) ProtoMessage() {}

func (x *UpdateContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentReq.ProtoReflect.Descriptor instead.
func (*UpdateContentReq) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type DeleteContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteContentReq) Reset() {
	*x = DeleteContentReq{}
	mi := &file_internal_api_operate_app_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContentReq) ProtoMessage() {}

func (x *DeleteContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContentReq.ProtoReflect.Descriptor instead.
func (*DeleteContentReq) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteContentReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 內容ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 作者
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// 標題
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 頁
	Page int32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	// 每頁數量
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *FindContentReq) Reset() {
	*x = FindContentReq{}
	mi := &file_internal_api_operate_app_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContentReq) ProtoMessage() {}

func (x *FindContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContentReq.ProtoReflect.Descriptor instead.
func (*FindContentReq) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{4}
}

func (x *FindContentReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindContentReq) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *FindContentReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FindContentReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindContentReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type CreateContentRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateContentRep) Reset() {
	*x = CreateContentRep{}
	mi := &file_internal_api_operate_app_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateContentRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentRep) ProtoMessage() {}

func (x *CreateContentRep) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentRep.ProtoReflect.Descriptor instead.
func (*CreateContentRep) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{5}
}

type UpdateContentRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateContentRep) Reset() {
	*x = UpdateContentRep{}
	mi := &file_internal_api_operate_app_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateContentRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentRep) ProtoMessage() {}

func (x *UpdateContentRep) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentRep.ProtoReflect.Descriptor instead.
func (*UpdateContentRep) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{6}
}

type DeleteContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteContentRsp) Reset() {
	*x = DeleteContentRsp{}
	mi := &file_internal_api_operate_app_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContentRsp) ProtoMessage() {}

func (x *DeleteContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContentRsp.ProtoReflect.Descriptor instead.
func (*DeleteContentRsp) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{7}
}

type FindContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 內容總數
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 內容
	Contents []*Content `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *FindContentRsp) Reset() {
	*x = FindContentRsp{}
	mi := &file_internal_api_operate_app_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContentRsp) ProtoMessage() {}

func (x *FindContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_operate_app_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContentRsp.ProtoReflect.Descriptor instead.
func (*FindContentRsp) Descriptor() ([]byte, []int) {
	return file_internal_api_operate_app_proto_rawDescGZIP(), []int{8}
}

func (x *FindContentRsp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FindContentRsp) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_internal_api_operate_app_proto protoreflect.FileDescriptor

var file_internal_api_operate_app_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x22, 0xf4, 0x02,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x7f, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x58, 0x0a,
	0x0e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xbb, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12,
	0x4d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x12, 0x4d,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x12, 0x4d, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x47, 0x0a, 0x0b,
	0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x73, 0x70, 0x42, 0x13, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x50, 0x01, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_api_operate_app_proto_rawDescOnce sync.Once
	file_internal_api_operate_app_proto_rawDescData = file_internal_api_operate_app_proto_rawDesc
)

func file_internal_api_operate_app_proto_rawDescGZIP() []byte {
	file_internal_api_operate_app_proto_rawDescOnce.Do(func() {
		file_internal_api_operate_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_operate_app_proto_rawDescData)
	})
	return file_internal_api_operate_app_proto_rawDescData
}

var file_internal_api_operate_app_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_api_operate_app_proto_goTypes = []any{
	(*Content)(nil),          // 0: api.operate.Content
	(*CreateContentReq)(nil), // 1: api.operate.CreateContentReq
	(*UpdateContentReq)(nil), // 2: api.operate.UpdateContentReq
	(*DeleteContentReq)(nil), // 3: api.operate.DeleteContentReq
	(*FindContentReq)(nil),   // 4: api.operate.FindContentReq
	(*CreateContentRep)(nil), // 5: api.operate.CreateContentRep
	(*UpdateContentRep)(nil), // 6: api.operate.UpdateContentRep
	(*DeleteContentRsp)(nil), // 7: api.operate.DeleteContentRsp
	(*FindContentRsp)(nil),   // 8: api.operate.FindContentRsp
}
var file_internal_api_operate_app_proto_depIdxs = []int32{
	0, // 0: api.operate.CreateContentReq.content:type_name -> api.operate.Content
	0, // 1: api.operate.UpdateContentReq.content:type_name -> api.operate.Content
	0, // 2: api.operate.FindContentRsp.contents:type_name -> api.operate.Content
	1, // 3: api.operate.App.CreateContent:input_type -> api.operate.CreateContentReq
	2, // 4: api.operate.App.UpdateContent:input_type -> api.operate.UpdateContentReq
	3, // 5: api.operate.App.DeleteContent:input_type -> api.operate.DeleteContentReq
	4, // 6: api.operate.App.FindContent:input_type -> api.operate.FindContentReq
	5, // 7: api.operate.App.CreateContent:output_type -> api.operate.CreateContentRep
	6, // 8: api.operate.App.UpdateContent:output_type -> api.operate.UpdateContentRep
	7, // 9: api.operate.App.DeleteContent:output_type -> api.operate.DeleteContentRsp
	8, // 10: api.operate.App.FindContent:output_type -> api.operate.FindContentRsp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_api_operate_app_proto_init() }
func file_internal_api_operate_app_proto_init() {
	if File_internal_api_operate_app_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_api_operate_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_operate_app_proto_goTypes,
		DependencyIndexes: file_internal_api_operate_app_proto_depIdxs,
		MessageInfos:      file_internal_api_operate_app_proto_msgTypes,
	}.Build()
	File_internal_api_operate_app_proto = out.File
	file_internal_api_operate_app_proto_rawDesc = nil
	file_internal_api_operate_app_proto_goTypes = nil
	file_internal_api_operate_app_proto_depIdxs = nil
}