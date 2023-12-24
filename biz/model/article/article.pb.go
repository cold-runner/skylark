// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: article.proto

package article

import (
	_ "github.com/cold-runner/skylark/biz/model/api"
	response "github.com/cold-runner/skylark/biz/model/response"
	user "github.com/cold-runner/skylark/biz/model/user"
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

type GetArticleFeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId  string `protobuf:"bytes,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty" form:"req_id" query:"req_id"`
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty" form:"cursor" query:"cursor"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty" form:"limit" query:"limit"`
}

func (x *GetArticleFeedReq) Reset() {
	*x = GetArticleFeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleFeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleFeedReq) ProtoMessage() {}

func (x *GetArticleFeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleFeedReq.ProtoReflect.Descriptor instead.
func (*GetArticleFeedReq) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{0}
}

func (x *GetArticleFeedReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *GetArticleFeedReq) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *GetArticleFeedReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetArticleFeedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Articles []*ArticlesInfo  `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty" form:"articles" query:"articles"`
}

func (x *GetArticleFeedResp) Reset() {
	*x = GetArticleFeedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleFeedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleFeedResp) ProtoMessage() {}

func (x *GetArticleFeedResp) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleFeedResp.ProtoReflect.Descriptor instead.
func (*GetArticleFeedResp) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{1}
}

func (x *GetArticleFeedResp) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetArticleFeedResp) GetArticles() []*ArticlesInfo {
	if x != nil {
		return x.Articles
	}
	return nil
}

type GetArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId     string `protobuf:"bytes,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty" form:"req_id" query:"req_id"`
	ArticleId string `protobuf:"bytes,2,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty" path:"article_id"`
}

func (x *GetArticleReq) Reset() {
	*x = GetArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleReq) ProtoMessage() {}

func (x *GetArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleReq.ProtoReflect.Descriptor instead.
func (*GetArticleReq) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticleReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *GetArticleReq) GetArticleId() string {
	if x != nil {
		return x.ArticleId
	}
	return ""
}

type GetArticleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      *response.Status     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Content     string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty" form:"content" query:"content"`
	AuthorInfo  *user.Basic          `protobuf:"bytes,3,opt,name=author_info,json=authorInfo,proto3" json:"author_info,omitempty" form:"author_info" query:"author_info"`
	Interaction *UserInteractionInfo `protobuf:"bytes,4,opt,name=interaction,proto3" json:"interaction,omitempty" form:"interaction" query:"interaction"`
}

func (x *GetArticleResp) Reset() {
	*x = GetArticleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleResp) ProtoMessage() {}

func (x *GetArticleResp) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleResp.ProtoReflect.Descriptor instead.
func (*GetArticleResp) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleResp) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetArticleResp) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetArticleResp) GetAuthorInfo() *user.Basic {
	if x != nil {
		return x.AuthorInfo
	}
	return nil
}

func (x *GetArticleResp) GetInteraction() *UserInteractionInfo {
	if x != nil {
		return x.Interaction
	}
	return nil
}

type SearchArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword    string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty" query:"keyword"`
	From       int32  `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty" query:"from"`
	MaxResults int32  `protobuf:"varint,3,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty" query:"max_results"`
	StartTime  string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" query:"start_time"`
	EndTime    string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" query:"end_time"`
	ReqId      string `protobuf:"bytes,6,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty" form:"req_id" query:"req_id"`
}

func (x *SearchArticleReq) Reset() {
	*x = SearchArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchArticleReq) ProtoMessage() {}

func (x *SearchArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchArticleReq.ProtoReflect.Descriptor instead.
func (*SearchArticleReq) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{4}
}

func (x *SearchArticleReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchArticleReq) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *SearchArticleReq) GetMaxResults() int32 {
	if x != nil {
		return x.MaxResults
	}
	return 0
}

func (x *SearchArticleReq) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *SearchArticleReq) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *SearchArticleReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

type SearchArticleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Articles []*ArticlesInfo  `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty" form:"articles" query:"articles"`
}

func (x *SearchArticleResp) Reset() {
	*x = SearchArticleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchArticleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchArticleResp) ProtoMessage() {}

func (x *SearchArticleResp) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchArticleResp.ProtoReflect.Descriptor instead.
func (*SearchArticleResp) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{5}
}

func (x *SearchArticleResp) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SearchArticleResp) GetArticles() []*ArticlesInfo {
	if x != nil {
		return x.Articles
	}
	return nil
}

type CategoriesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId string `protobuf:"bytes,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty" form:"req_id" query:"req_id"`
}

func (x *CategoriesReq) Reset() {
	*x = CategoriesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesReq) ProtoMessage() {}

func (x *CategoriesReq) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesReq.ProtoReflect.Descriptor instead.
func (*CategoriesReq) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{6}
}

func (x *CategoriesReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

type CategoriesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Categories []*Category      `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty" form:"categories" query:"categories"`
}

func (x *CategoriesResp) Reset() {
	*x = CategoriesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesResp) ProtoMessage() {}

func (x *CategoriesResp) ProtoReflect() protoreflect.Message {
	mi := &file_article_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesResp.ProtoReflect.Descriptor instead.
func (*CategoriesResp) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{7}
}

func (x *CategoriesResp) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CategoriesResp) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_article_proto protoreflect.FileDescriptor

var file_article_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a,
	0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x65, 0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x71, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xd2, 0xbb, 0x18, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xc2, 0x01, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3e, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xf8, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xb2, 0xbb, 0x18, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xb2, 0xbb, 0x18,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x30, 0x0a, 0x0b, 0x6d,
	0x61, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0f, 0xb2, 0xbb, 0x18, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xb2, 0xbb, 0x18, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xb2, 0xbb, 0x18, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x11,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x26,
	0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0xd9, 0x02, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x54, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x1b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x09, 0xca, 0xc1,
	0x18, 0x05, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x12, 0x57, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0xca, 0xc1, 0x18, 0x14, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2f, 0x3a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x12, 0x53, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x19, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0b, 0xca, 0xc1, 0x18, 0x07, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x4a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x0d, 0xca, 0xc1, 0x18, 0x09, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6c, 0x64, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x6b, 0x79, 0x6c,
	0x61, 0x72, 0x6b, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_article_proto_rawDescOnce sync.Once
	file_article_proto_rawDescData = file_article_proto_rawDesc
)

func file_article_proto_rawDescGZIP() []byte {
	file_article_proto_rawDescOnce.Do(func() {
		file_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_proto_rawDescData)
	})
	return file_article_proto_rawDescData
}

var file_article_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_article_proto_goTypes = []interface{}{
	(*GetArticleFeedReq)(nil),   // 0: article.GetArticleFeedReq
	(*GetArticleFeedResp)(nil),  // 1: article.GetArticleFeedResp
	(*GetArticleReq)(nil),       // 2: article.GetArticleReq
	(*GetArticleResp)(nil),      // 3: article.GetArticleResp
	(*SearchArticleReq)(nil),    // 4: article.SearchArticleReq
	(*SearchArticleResp)(nil),   // 5: article.SearchArticleResp
	(*CategoriesReq)(nil),       // 6: article.CategoriesReq
	(*CategoriesResp)(nil),      // 7: article.CategoriesResp
	(*response.Status)(nil),     // 8: response.Status
	(*ArticlesInfo)(nil),        // 9: article.ArticlesInfo
	(*user.Basic)(nil),          // 10: user.Basic
	(*UserInteractionInfo)(nil), // 11: article.UserInteractionInfo
	(*Category)(nil),            // 12: article.Category
}
var file_article_proto_depIdxs = []int32{
	8,  // 0: article.GetArticleFeedResp.status:type_name -> response.Status
	9,  // 1: article.GetArticleFeedResp.articles:type_name -> article.ArticlesInfo
	8,  // 2: article.GetArticleResp.status:type_name -> response.Status
	10, // 3: article.GetArticleResp.author_info:type_name -> user.Basic
	11, // 4: article.GetArticleResp.interaction:type_name -> article.UserInteractionInfo
	8,  // 5: article.SearchArticleResp.status:type_name -> response.Status
	9,  // 6: article.SearchArticleResp.articles:type_name -> article.ArticlesInfo
	8,  // 7: article.CategoriesResp.status:type_name -> response.Status
	12, // 8: article.CategoriesResp.categories:type_name -> article.Category
	0,  // 9: article.Article.GetArticleFeed:input_type -> article.GetArticleFeedReq
	2,  // 10: article.Article.GetArticle:input_type -> article.GetArticleReq
	4,  // 11: article.Article.SearchArticle:input_type -> article.SearchArticleReq
	6,  // 12: article.Article.Category:input_type -> article.CategoriesReq
	1,  // 13: article.Article.GetArticleFeed:output_type -> article.GetArticleFeedResp
	3,  // 14: article.Article.GetArticle:output_type -> article.GetArticleResp
	5,  // 15: article.Article.SearchArticle:output_type -> article.SearchArticleResp
	7,  // 16: article.Article.Category:output_type -> article.CategoriesResp
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_article_proto_init() }
func file_article_proto_init() {
	if File_article_proto != nil {
		return
	}
	file_articleBasic_proto_init()
	file_category_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleFeedReq); i {
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
		file_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleFeedResp); i {
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
		file_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleReq); i {
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
		file_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleResp); i {
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
		file_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchArticleReq); i {
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
		file_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchArticleResp); i {
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
		file_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesReq); i {
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
		file_article_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesResp); i {
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
			RawDescriptor: file_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_proto_goTypes,
		DependencyIndexes: file_article_proto_depIdxs,
		MessageInfos:      file_article_proto_msgTypes,
	}.Build()
	File_article_proto = out.File
	file_article_proto_rawDesc = nil
	file_article_proto_goTypes = nil
	file_article_proto_depIdxs = nil
}
