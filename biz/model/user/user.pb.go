// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: user.proto

package user

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

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StuNum          string `protobuf:"bytes,1,opt,name=stu_num,json=stuNum,proto3" form:"stu_num" json:"stu_num,omitempty" vd:"stuNum"`
	StuName         string `protobuf:"bytes,2,opt,name=stu_name,json=stuName,proto3" form:"stu_name" json:"stu_name,omitempty" vd:"stuName"`
	Gender          string `protobuf:"bytes,3,opt,name=gender,proto3" form:"gender" json:"gender,omitempty" vd:"gender"`
	College         string `protobuf:"bytes,4,opt,name=college,proto3" form:"college" json:"college,omitempty" vd:"college"`
	Major           string `protobuf:"bytes,5,opt,name=major,proto3" form:"major" json:"major,omitempty" vd:"major"`
	Grade           string `protobuf:"bytes,6,opt,name=grade,proto3" form:"grade" json:"grade,omitempty" vd:"grade"`
	Password        string `protobuf:"bytes,7,opt,name=password,proto3" form:"password" json:"password,omitempty" vd:"password"`
	StuCardPhotoUrl string `protobuf:"bytes,8,opt,name=stu_card_photo_url,json=stuCardPhotoUrl,proto3" form:"stu_card_photo_url" json:"stu_card_photo_url,omitempty" vd:"http_url"`
	Phone           string `protobuf:"bytes,9,opt,name=phone,proto3" form:"phone" json:"phone,omitempty" vd:"phone"`
	SmsCode         string `protobuf:"bytes,10,opt,name=sms_code,json=smsCode,proto3" form:"sms_code" json:"sms_code,omitempty" vd:"smsCode"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetStuNum() string {
	if x != nil {
		return x.StuNum
	}
	return ""
}

func (x *RegisterReq) GetStuName() string {
	if x != nil {
		return x.StuName
	}
	return ""
}

func (x *RegisterReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *RegisterReq) GetCollege() string {
	if x != nil {
		return x.College
	}
	return ""
}

func (x *RegisterReq) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *RegisterReq) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetStuCardPhotoUrl() string {
	if x != nil {
		return x.StuCardPhotoUrl
	}
	return ""
}

func (x *RegisterReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterReq) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResp) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type PasswordLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StuNum   string `protobuf:"bytes,1,opt,name=stu_num,json=stuNum,proto3" form:"stu_num" json:"stu_num,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" form:"password" json:"password,omitempty"`
}

func (x *PasswordLoginReq) Reset() {
	*x = PasswordLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordLoginReq) ProtoMessage() {}

func (x *PasswordLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordLoginReq.ProtoReflect.Descriptor instead.
func (*PasswordLoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *PasswordLoginReq) GetStuNum() string {
	if x != nil {
		return x.StuNum
	}
	return ""
}

func (x *PasswordLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	Token  string           `protobuf:"bytes,2,opt,name=token,proto3" form:"token" json:"token,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResp) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PhoneLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" form:"phone" json:"phone,omitempty"`
	SmsCode string `protobuf:"bytes,2,opt,name=sms_code,json=smsCode,proto3" form:"sms_code" json:"sms_code,omitempty"`
}

func (x *PhoneLoginReq) Reset() {
	*x = PhoneLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneLoginReq) ProtoMessage() {}

func (x *PhoneLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneLoginReq.ProtoReflect.Descriptor instead.
func (*PhoneLoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *PhoneLoginReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PhoneLoginReq) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

type SendSmsCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty" form:"phone" query:"phone"`
	ReqId string `protobuf:"bytes,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty" form:"req_id" query:"req_id"`
}

func (x *SendSmsCodeReq) Reset() {
	*x = SendSmsCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsCodeReq) ProtoMessage() {}

func (x *SendSmsCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsCodeReq.ProtoReflect.Descriptor instead.
func (*SendSmsCodeReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *SendSmsCodeReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SendSmsCodeReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

type SendSmsCodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
}

func (x *SendSmsCodeRes) Reset() {
	*x = SendSmsCodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsCodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsCodeRes) ProtoMessage() {}

func (x *SendSmsCodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsCodeRes.ProtoReflect.Descriptor instead.
func (*SendSmsCodeRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *SendSmsCodeRes) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetUserInfoByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	ReqId string `protobuf:"bytes,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty" form:"req_id" query:"req_id"`
}

func (x *GetUserInfoByIdReq) Reset() {
	*x = GetUserInfoByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoByIdReq) ProtoMessage() {}

func (x *GetUserInfoByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoByIdReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoByIdReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserInfoByIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUserInfoByIdReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

type GetUserInfoByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    *response.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
	BasicInfo *Basic           `protobuf:"bytes,2,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty" form:"basic_info" query:"basic_info"`
}

func (x *GetUserInfoByIdRes) Reset() {
	*x = GetUserInfoByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoByIdRes) ProtoMessage() {}

func (x *GetUserInfoByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoByIdRes.ProtoReflect.Descriptor instead.
func (*GetUserInfoByIdRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserInfoByIdRes) GetStatus() *response.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetUserInfoByIdRes) GetBasicInfo() *Basic {
	if x != nil {
		return x.BasicInfo
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x04, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x2e,
	0x0a, 0x07, 0x73, 0x74, 0x75, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x15, 0xca, 0xbb, 0x18, 0x07, 0x73, 0x74, 0x75, 0x5f, 0x6e, 0x75, 0x6d, 0xda, 0xbb, 0x18, 0x06,
	0x73, 0x74, 0x75, 0x4e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x75, 0x4e, 0x75, 0x6d, 0x12, 0x32,
	0x0a, 0x08, 0x73, 0x74, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x17, 0xca, 0xbb, 0x18, 0x08, 0x73, 0x74, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xda, 0xbb,
	0x18, 0x07, 0x73, 0x74, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x07, 0x73, 0x74, 0x75, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x14, 0xca, 0xbb, 0x18, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0xda, 0xbb,
	0x18, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0xca, 0xbb, 0x18, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0xda, 0xbb,
	0x18, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x67, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x12, 0xca, 0xbb, 0x18, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0xda, 0xbb, 0x18, 0x05,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xca, 0xbb, 0x18,
	0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0xda, 0xbb, 0x18, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xca, 0xbb, 0x18, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0xda, 0xbb, 0x18, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x4f, 0x0a, 0x12,
	0x73, 0x74, 0x75, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xca, 0xbb, 0x18, 0x12, 0x73, 0x74,
	0x75, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0xda, 0xbb, 0x18, 0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x52, 0x0f, 0x73, 0x74,
	0x75, 0x43, 0x61, 0x72, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xca, 0xbb,
	0x18, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0xda, 0xbb, 0x18, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x6d, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xca, 0xbb, 0x18, 0x08, 0x73,
	0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0xda, 0xbb, 0x18, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x62, 0x0a, 0x10, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x07, 0x73, 0x74, 0x75,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07,
	0x73, 0x74, 0x75, 0x5f, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x75, 0x4e, 0x75, 0x6d, 0x12,
	0x28, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x56, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xca, 0xbb, 0x18, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x59, 0x0a, 0x0d, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xca, 0xbb, 0x18, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x73, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xbb, 0x18, 0x08, 0x73, 0x6d, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3d, 0x0a, 0x0e,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x65, 0x71, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x09, 0x62, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0x84, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0xd2, 0xc1,
	0x18, 0x09, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0d, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0xca, 0xc1, 0x18, 0x0f, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x44, 0x0a, 0x0a, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10,
	0xca, 0xc1, 0x18, 0x0c, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x4b, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22, 0x10, 0xca, 0xc1, 0x18,
	0x0c, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x58, 0x0a,
	0x0b, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x22, 0x15, 0xca, 0xc1, 0x18, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x64, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x72, 0x6b, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_proto_goTypes = []interface{}{
	(*RegisterReq)(nil),        // 0: user.RegisterReq
	(*RegisterResp)(nil),       // 1: user.RegisterResp
	(*PasswordLoginReq)(nil),   // 2: user.PasswordLoginReq
	(*LoginResp)(nil),          // 3: user.LoginResp
	(*PhoneLoginReq)(nil),      // 4: user.PhoneLoginReq
	(*SendSmsCodeReq)(nil),     // 5: user.SendSmsCodeReq
	(*SendSmsCodeRes)(nil),     // 6: user.SendSmsCodeRes
	(*GetUserInfoByIdReq)(nil), // 7: user.GetUserInfoByIdReq
	(*GetUserInfoByIdRes)(nil), // 8: user.GetUserInfoByIdRes
	(*response.Status)(nil),    // 9: response.Status
	(*Basic)(nil),              // 10: user.Basic
}
var file_user_proto_depIdxs = []int32{
	9,  // 0: user.RegisterResp.status:type_name -> response.Status
	9,  // 1: user.LoginResp.status:type_name -> response.Status
	9,  // 2: user.SendSmsCodeRes.status:type_name -> response.Status
	9,  // 3: user.GetUserInfoByIdRes.status:type_name -> response.Status
	10, // 4: user.GetUserInfoByIdRes.basic_info:type_name -> user.Basic
	0,  // 5: user.User.register:input_type -> user.RegisterReq
	2,  // 6: user.User.passwordLogin:input_type -> user.PasswordLoginReq
	4,  // 7: user.User.phoneLogin:input_type -> user.PhoneLoginReq
	5,  // 8: user.User.SendSmsCode:input_type -> user.SendSmsCodeReq
	7,  // 9: user.User.getUserInfo:input_type -> user.GetUserInfoByIdReq
	1,  // 10: user.User.register:output_type -> user.RegisterResp
	3,  // 11: user.User.passwordLogin:output_type -> user.LoginResp
	3,  // 12: user.User.phoneLogin:output_type -> user.LoginResp
	6,  // 13: user.User.SendSmsCode:output_type -> user.SendSmsCodeRes
	8,  // 14: user.User.getUserInfo:output_type -> user.GetUserInfoByIdRes
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_userBasic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordLoginReq); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneLoginReq); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsCodeReq); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsCodeRes); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoByIdReq); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoByIdRes); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}