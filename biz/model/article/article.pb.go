// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: article.proto

package article

import (
	_ "github.com/cold-runner/skylark/biz/model/api"
	response "github.com/cold-runner/skylark/biz/model/response"
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

type GetArticleFeedRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Articles []*ArticlesInfo  `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty" form:"articles" query:"articles"`
}

func (x *GetArticleFeedRes) Reset() {
	*x = GetArticleFeedRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleFeedRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleFeedRes) ProtoMessage() {}

func (x *GetArticleFeedRes) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetArticleFeedRes.ProtoReflect.Descriptor instead.
func (*GetArticleFeedRes) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{1}
}

func (x *GetArticleFeedRes) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetArticleFeedRes) GetArticles() []*ArticlesInfo {
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

type GetArticleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Article *ArticlesInfo    `protobuf:"bytes,2,opt,name=article,proto3" json:"article,omitempty" form:"article" query:"article"`
}

func (x *GetArticleRes) Reset() {
	*x = GetArticleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRes) ProtoMessage() {}

func (x *GetArticleRes) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetArticleRes.ProtoReflect.Descriptor instead.
func (*GetArticleRes) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleRes) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetArticleRes) GetArticle() *ArticlesInfo {
	if x != nil {
		return x.Article
	}
	return nil
}

type SearchArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword    string `protobuf:"bytes,1,opt,name=keyword,proto3" form:"keyword" json:"keyword,omitempty"`
	From       int32  `protobuf:"varint,2,opt,name=from,proto3" form:"from" json:"from,omitempty"`
	MaxResults int32  `protobuf:"varint,3,opt,name=max_results,json=maxResults,proto3" form:"max_results" json:"max_results,omitempty"`
	StartTime  string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" form:"start_time" json:"start_time,omitempty"`
	EndTime    string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" form:"end_time" json:"end_time,omitempty"`
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

type SearchArticleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Articles []*ArticlesInfo  `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty" form:"articles" query:"articles"`
}

func (x *SearchArticleRes) Reset() {
	*x = SearchArticleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchArticleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchArticleRes) ProtoMessage() {}

func (x *SearchArticleRes) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SearchArticleRes.ProtoReflect.Descriptor instead.
func (*SearchArticleRes) Descriptor() ([]byte, []int) {
	return file_article_proto_rawDescGZIP(), []int{5}
}

func (x *SearchArticleRes) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SearchArticleRes) GetArticles() []*ArticlesInfo {
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
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06,
	0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x70, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0a, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0xd2, 0xbb, 0x18, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x52,
	0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb,
	0x18, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x08, 0xca, 0xbb, 0x18, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x30, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xca, 0xbb, 0x18, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x2d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65,
	0x71, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49,
	0x64, 0x22, 0x6f, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x31, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x22, 0x26, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x0e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0xd6, 0x02, 0x0a, 0x07, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x22,
	0x09, 0xca, 0xc1, 0x18, 0x05, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x12, 0x56, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x18, 0xca, 0xc1, 0x18, 0x14, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x3a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x12, 0x52, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x0b, 0xca, 0xc1, 0x18, 0x07, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x4a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x16, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x0d, 0xca, 0xc1, 0x18, 0x09, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6c, 0x64, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x6b, 0x79,
	0x6c, 0x61, 0x72, 0x6b, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*GetArticleFeedReq)(nil), // 0: article.GetArticleFeedReq
	(*GetArticleFeedRes)(nil), // 1: article.GetArticleFeedRes
	(*GetArticleReq)(nil),     // 2: article.GetArticleReq
	(*GetArticleRes)(nil),     // 3: article.GetArticleRes
	(*SearchArticleReq)(nil),  // 4: article.SearchArticleReq
	(*SearchArticleRes)(nil),  // 5: article.SearchArticleRes
	(*CategoriesReq)(nil),     // 6: article.CategoriesReq
	(*CategoriesResp)(nil),    // 7: article.CategoriesResp
	(*response.Status)(nil),   // 8: response.Status
	(*ArticlesInfo)(nil),      // 9: article.ArticlesInfo
	(*Category)(nil),          // 10: article.Category
}
var file_article_proto_depIdxs = []int32{
	8,  // 0: article.GetArticleFeedRes.status:type_name -> response.Status
	9,  // 1: article.GetArticleFeedRes.articles:type_name -> article.ArticlesInfo
	8,  // 2: article.GetArticleRes.status:type_name -> response.Status
	9,  // 3: article.GetArticleRes.article:type_name -> article.ArticlesInfo
	8,  // 4: article.SearchArticleRes.status:type_name -> response.Status
	9,  // 5: article.SearchArticleRes.articles:type_name -> article.ArticlesInfo
	8,  // 6: article.CategoriesResp.status:type_name -> response.Status
	10, // 7: article.CategoriesResp.categories:type_name -> article.Category
	0,  // 8: article.Article.GetArticleFeed:input_type -> article.GetArticleFeedReq
	2,  // 9: article.Article.GetArticle:input_type -> article.GetArticleReq
	4,  // 10: article.Article.SearchArticle:input_type -> article.SearchArticleReq
	6,  // 11: article.Article.Category:input_type -> article.CategoriesReq
	1,  // 12: article.Article.GetArticleFeed:output_type -> article.GetArticleFeedRes
	3,  // 13: article.Article.GetArticle:output_type -> article.GetArticleRes
	5,  // 14: article.Article.SearchArticle:output_type -> article.SearchArticleRes
	7,  // 15: article.Article.Category:output_type -> article.CategoriesResp
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
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
			switch v := v.(*GetArticleFeedRes); i {
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
			switch v := v.(*GetArticleRes); i {
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
			switch v := v.(*SearchArticleRes); i {
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
