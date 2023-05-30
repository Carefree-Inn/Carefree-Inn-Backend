// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: proto/post.proto

package post

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{1}
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content    string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	SearchType string `protobuf:"bytes,2,opt,name=SearchType,proto3" json:"SearchType,omitempty"`
	Account    string `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{2}
}

func (x *SearchRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SearchRequest) GetSearchType() string {
	if x != nil {
		return x.SearchType
	}
	return ""
}

func (x *SearchRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string        `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Title    string        `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content  string        `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	Category *CategoryInfo `protobuf:"bytes,4,opt,name=Category,proto3" json:"Category,omitempty"`
	Tag      []*TagInfo    `protobuf:"bytes,5,rep,name=Tag,proto3" json:"Tag,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePostRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreatePostRequest) GetCategory() *CategoryInfo {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *CreatePostRequest) GetTag() []*TagInfo {
	if x != nil {
		return x.Tag
	}
	return nil
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId  uint32 `protobuf:"varint,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePostRequest) GetPostId() uint32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *DeletePostRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type PostOfCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *CategoryInfo `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Page     uint32        `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit    uint32        `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Account  string        `protobuf:"bytes,4,opt,name=Account,proto3" json:"Account,omitempty"`
}

func (x *PostOfCategoryRequest) Reset() {
	*x = PostOfCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostOfCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOfCategoryRequest) ProtoMessage() {}

func (x *PostOfCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOfCategoryRequest.ProtoReflect.Descriptor instead.
func (*PostOfCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{5}
}

func (x *PostOfCategoryRequest) GetCategory() *CategoryInfo {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *PostOfCategoryRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PostOfCategoryRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PostOfCategoryRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type CategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*CategoryInfo `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{6}
}

func (x *CategoryResponse) GetCategories() []*CategoryInfo {
	if x != nil {
		return x.Categories
	}
	return nil
}

type PostOfTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Page    uint32 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit   uint32 `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Account string `protobuf:"bytes,4,opt,name=Account,proto3" json:"Account,omitempty"`
}

func (x *PostOfTagRequest) Reset() {
	*x = PostOfTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostOfTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOfTagRequest) ProtoMessage() {}

func (x *PostOfTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOfTagRequest.ProtoReflect.Descriptor instead.
func (*PostOfTagRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{7}
}

func (x *PostOfTagRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostOfTagRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PostOfTagRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PostOfTagRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type CategoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	CategoryId uint32 `protobuf:"varint,2,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
}

func (x *CategoryInfo) Reset() {
	*x = CategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryInfo) ProtoMessage() {}

func (x *CategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryInfo.ProtoReflect.Descriptor instead.
func (*CategoryInfo) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{8}
}

func (x *CategoryInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CategoryInfo) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

type TagInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	TagId uint32 `protobuf:"varint,2,opt,name=TagId,proto3" json:"TagId,omitempty"`
}

func (x *TagInfo) Reset() {
	*x = TagInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagInfo) ProtoMessage() {}

func (x *TagInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagInfo.ProtoReflect.Descriptor instead.
func (*TagInfo) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{9}
}

func (x *TagInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TagInfo) GetTagId() uint32 {
	if x != nil {
		return x.TagId
	}
	return 0
}

type PostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account    string        `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	PostId     uint32        `protobuf:"varint,2,opt,name=PostId,proto3" json:"PostId,omitempty"`
	Title      string        `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Content    string        `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Category   *CategoryInfo `protobuf:"bytes,5,opt,name=Category,proto3" json:"Category,omitempty"`
	Likes      uint32        `protobuf:"varint,6,opt,name=Likes,proto3" json:"Likes,omitempty"`
	Comments   uint32        `protobuf:"varint,7,opt,name=Comments,proto3" json:"Comments,omitempty"`
	CreateTime string        `protobuf:"bytes,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	Tag        []*TagInfo    `protobuf:"bytes,9,rep,name=tag,proto3" json:"tag,omitempty"`
	Liked      bool          `protobuf:"varint,10,opt,name=Liked,proto3" json:"Liked,omitempty"`
}

func (x *PostInfo) Reset() {
	*x = PostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostInfo) ProtoMessage() {}

func (x *PostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostInfo.ProtoReflect.Descriptor instead.
func (*PostInfo) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{10}
}

func (x *PostInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PostInfo) GetPostId() uint32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PostInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostInfo) GetCategory() *CategoryInfo {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *PostInfo) GetLikes() uint32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *PostInfo) GetComments() uint32 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *PostInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *PostInfo) GetTag() []*TagInfo {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *PostInfo) GetLiked() bool {
	if x != nil {
		return x.Liked
	}
	return false
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*PostInfo `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{11}
}

func (x *PostResponse) GetPosts() []*PostInfo {
	if x != nil {
		return x.Posts
	}
	return nil
}

type PostOfUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Page    int32  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit   int32  `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *PostOfUserRequest) Reset() {
	*x = PostOfUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostOfUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostOfUserRequest) ProtoMessage() {}

func (x *PostOfUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostOfUserRequest.ProtoReflect.Descriptor instead.
func (*PostOfUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_proto_rawDescGZIP(), []int{12}
}

func (x *PostOfUserRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PostOfUserRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PostOfUserRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_proto_post_proto protoreflect.FileDescriptor

var file_proto_post_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x63, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x03, 0x54, 0x61, 0x67, 0x22, 0x45, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8b, 0x01, 0x0a,
	0x15, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x6c, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x44, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x61, 0x67, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x61, 0x67, 0x49, 0x64, 0x22, 0xa5, 0x02,
	0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x08,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x57, 0x0a, 0x11, 0x50,
	0x6f, 0x73, 0x74, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x32, 0xf4, 0x03, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0d,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x4f, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x54, 0x61, 0x67, 0x12,
	0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x6b, 0x65, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4f,
	0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_post_proto_rawDescOnce sync.Once
	file_proto_post_proto_rawDescData = file_proto_post_proto_rawDesc
)

func file_proto_post_proto_rawDescGZIP() []byte {
	file_proto_post_proto_rawDescOnce.Do(func() {
		file_proto_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_post_proto_rawDescData)
	})
	return file_proto_post_proto_rawDescData
}

var file_proto_post_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_post_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: post.Request
	(*Response)(nil),              // 1: post.Response
	(*SearchRequest)(nil),         // 2: post.SearchRequest
	(*CreatePostRequest)(nil),     // 3: post.CreatePostRequest
	(*DeletePostRequest)(nil),     // 4: post.DeletePostRequest
	(*PostOfCategoryRequest)(nil), // 5: post.PostOfCategoryRequest
	(*CategoryResponse)(nil),      // 6: post.CategoryResponse
	(*PostOfTagRequest)(nil),      // 7: post.PostOfTagRequest
	(*CategoryInfo)(nil),          // 8: post.CategoryInfo
	(*TagInfo)(nil),               // 9: post.TagInfo
	(*PostInfo)(nil),              // 10: post.PostInfo
	(*PostResponse)(nil),          // 11: post.PostResponse
	(*PostOfUserRequest)(nil),     // 12: post.PostOfUserRequest
}
var file_proto_post_proto_depIdxs = []int32{
	8,  // 0: post.CreatePostRequest.Category:type_name -> post.CategoryInfo
	9,  // 1: post.CreatePostRequest.Tag:type_name -> post.TagInfo
	8,  // 2: post.PostOfCategoryRequest.category:type_name -> post.CategoryInfo
	8,  // 3: post.CategoryResponse.categories:type_name -> post.CategoryInfo
	8,  // 4: post.PostInfo.Category:type_name -> post.CategoryInfo
	9,  // 5: post.PostInfo.tag:type_name -> post.TagInfo
	10, // 6: post.PostResponse.posts:type_name -> post.PostInfo
	3,  // 7: post.Post.CreatePost:input_type -> post.CreatePostRequest
	4,  // 8: post.Post.DeletePost:input_type -> post.DeletePostRequest
	0,  // 9: post.Post.GetCategory:input_type -> post.Request
	5,  // 10: post.Post.GetPostOfCategory:input_type -> post.PostOfCategoryRequest
	7,  // 11: post.Post.GetPostOfTag:input_type -> post.PostOfTagRequest
	12, // 12: post.Post.GetPostOfUser:input_type -> post.PostOfUserRequest
	12, // 13: post.Post.GetPostOfUserLiked:input_type -> post.PostOfUserRequest
	2,  // 14: post.Post.SearchPost:input_type -> post.SearchRequest
	1,  // 15: post.Post.CreatePost:output_type -> post.Response
	1,  // 16: post.Post.DeletePost:output_type -> post.Response
	6,  // 17: post.Post.GetCategory:output_type -> post.CategoryResponse
	11, // 18: post.Post.GetPostOfCategory:output_type -> post.PostResponse
	11, // 19: post.Post.GetPostOfTag:output_type -> post.PostResponse
	11, // 20: post.Post.GetPostOfUser:output_type -> post.PostResponse
	11, // 21: post.Post.GetPostOfUserLiked:output_type -> post.PostResponse
	11, // 22: post.Post.SearchPost:output_type -> post.PostResponse
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_post_proto_init() }
func file_proto_post_proto_init() {
	if File_proto_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_proto_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_proto_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_proto_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostOfCategoryRequest); i {
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
		file_proto_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryResponse); i {
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
		file_proto_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostOfTagRequest); i {
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
		file_proto_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryInfo); i {
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
		file_proto_post_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagInfo); i {
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
		file_proto_post_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostInfo); i {
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
		file_proto_post_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
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
		file_proto_post_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostOfUserRequest); i {
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
			RawDescriptor: file_proto_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_post_proto_goTypes,
		DependencyIndexes: file_proto_post_proto_depIdxs,
		MessageInfos:      file_proto_post_proto_msgTypes,
	}.Build()
	File_proto_post_proto = out.File
	file_proto_post_proto_rawDesc = nil
	file_proto_post_proto_goTypes = nil
	file_proto_post_proto_depIdxs = nil
}
