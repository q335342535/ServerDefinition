// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: source/user_server/protobuf.proto

package user_server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	common "target/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtocolID int32

const (
	ProtocolID_UNDEFINED                        ProtocolID = 0
	ProtocolID_REQUEST_USER_DATA                ProtocolID = 1 // 请求用户数据
	ProtocolID_REPLY_USER_DATA                  ProtocolID = 2 // 回复用户数据
	ProtocolID_REQUEST_REAL_NAME_AUTHENTICATION ProtocolID = 3 // 请求实名认证
	ProtocolID_REPLY_REAL_NAME_AUTHENTICATION   ProtocolID = 4 // 回复实名认证
	ProtocolID_REQUEST_MOBILE_BINDING           ProtocolID = 5 // 请求手机号绑定
	ProtocolID_REPLY_MOBILE_BINDING             ProtocolID = 6 // 回复手机号绑定
	ProtocolID_REQUEST_USER_GAME_POSITION       ProtocolID = 7 // 请求用户游戏位置
	ProtocolID_REPLY_USER_GAME_POSITION         ProtocolID = 8 // 回复用户游戏位置
)

// Enum value maps for ProtocolID.
var (
	ProtocolID_name = map[int32]string{
		0: "UNDEFINED",
		1: "REQUEST_USER_DATA",
		2: "REPLY_USER_DATA",
		3: "REQUEST_REAL_NAME_AUTHENTICATION",
		4: "REPLY_REAL_NAME_AUTHENTICATION",
		5: "REQUEST_MOBILE_BINDING",
		6: "REPLY_MOBILE_BINDING",
		7: "REQUEST_USER_GAME_POSITION",
		8: "REPLY_USER_GAME_POSITION",
	}
	ProtocolID_value = map[string]int32{
		"UNDEFINED":                        0,
		"REQUEST_USER_DATA":                1,
		"REPLY_USER_DATA":                  2,
		"REQUEST_REAL_NAME_AUTHENTICATION": 3,
		"REPLY_REAL_NAME_AUTHENTICATION":   4,
		"REQUEST_MOBILE_BINDING":           5,
		"REPLY_MOBILE_BINDING":             6,
		"REQUEST_USER_GAME_POSITION":       7,
		"REPLY_USER_GAME_POSITION":         8,
	}
)

func (x ProtocolID) Enum() *ProtocolID {
	p := new(ProtocolID)
	*p = x
	return p
}

func (x ProtocolID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtocolID) Descriptor() protoreflect.EnumDescriptor {
	return file_source_user_server_protobuf_proto_enumTypes[0].Descriptor()
}

func (ProtocolID) Type() protoreflect.EnumType {
	return &file_source_user_server_protobuf_proto_enumTypes[0]
}

func (x ProtocolID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtocolID.Descriptor instead.
func (ProtocolID) EnumDescriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{0}
}

type RequestUserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestUserData) Reset() {
	*x = RequestUserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUserData) ProtoMessage() {}

func (x *RequestUserData) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUserData.ProtoReflect.Descriptor instead.
func (*RequestUserData) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{0}
}

type ReplyUserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName                 string `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`                                                      //昵称
	HeadUrl                  string `protobuf:"bytes,2,opt,name=head_url,json=headUrl,proto3" json:"head_url,omitempty"`                                                         //头像
	Sex                      int32  `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`                                                                               //性别
	IsRealNameAuthentication bool   `protobuf:"varint,4,opt,name=is_real_name_authentication,json=isRealNameAuthentication,proto3" json:"is_real_name_authentication,omitempty"` //是否实名
	Mobile                   string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`                                                                          //手机号
}

func (x *ReplyUserData) Reset() {
	*x = ReplyUserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyUserData) ProtoMessage() {}

func (x *ReplyUserData) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyUserData.ProtoReflect.Descriptor instead.
func (*ReplyUserData) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyUserData) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *ReplyUserData) GetHeadUrl() string {
	if x != nil {
		return x.HeadUrl
	}
	return ""
}

func (x *ReplyUserData) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *ReplyUserData) GetIsRealNameAuthentication() bool {
	if x != nil {
		return x.IsRealNameAuthentication
	}
	return false
}

func (x *ReplyUserData) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type RequestRealNameAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RealName string `protobuf:"bytes,1,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"` //真实姓名
	IdCard   string `protobuf:"bytes,2,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`       //身份证号
}

func (x *RequestRealNameAuthentication) Reset() {
	*x = RequestRealNameAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRealNameAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRealNameAuthentication) ProtoMessage() {}

func (x *RequestRealNameAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRealNameAuthentication.ProtoReflect.Descriptor instead.
func (*RequestRealNameAuthentication) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{2}
}

func (x *RequestRealNameAuthentication) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *RequestRealNameAuthentication) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

type ReplyRealNameAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ErrorCode `protobuf:"varint,100,opt,name=error_code,json=errorCode,proto3,enum=protobuf.ErrorCode" json:"error_code,omitempty"`
}

func (x *ReplyRealNameAuthentication) Reset() {
	*x = ReplyRealNameAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyRealNameAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyRealNameAuthentication) ProtoMessage() {}

func (x *ReplyRealNameAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyRealNameAuthentication.ProtoReflect.Descriptor instead.
func (*ReplyRealNameAuthentication) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyRealNameAuthentication) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_SUCCESS
}

type RequestMobileBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile     string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	MobileCode string `protobuf:"bytes,2,opt,name=mobile_code,json=mobileCode,proto3" json:"mobile_code,omitempty"`
}

func (x *RequestMobileBinding) Reset() {
	*x = RequestMobileBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMobileBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMobileBinding) ProtoMessage() {}

func (x *RequestMobileBinding) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMobileBinding.ProtoReflect.Descriptor instead.
func (*RequestMobileBinding) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{4}
}

func (x *RequestMobileBinding) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RequestMobileBinding) GetMobileCode() string {
	if x != nil {
		return x.MobileCode
	}
	return ""
}

type ReplyMobileBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ErrorCode `protobuf:"varint,100,opt,name=error_code,json=errorCode,proto3,enum=protobuf.ErrorCode" json:"error_code,omitempty"`
}

func (x *ReplyMobileBinding) Reset() {
	*x = ReplyMobileBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMobileBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMobileBinding) ProtoMessage() {}

func (x *ReplyMobileBinding) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMobileBinding.ProtoReflect.Descriptor instead.
func (*ReplyMobileBinding) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{5}
}

func (x *ReplyMobileBinding) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_SUCCESS
}

type RequestUserGamePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestUserGamePosition) Reset() {
	*x = RequestUserGamePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUserGamePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUserGamePosition) ProtoMessage() {}

func (x *RequestUserGamePosition) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUserGamePosition.ProtoReflect.Descriptor instead.
func (*RequestUserGamePosition) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{6}
}

type ReplyUserGamePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GamePosition *common.GamePosition `protobuf:"bytes,1,opt,name=game_position,json=gamePosition,proto3" json:"game_position,omitempty"`
	ErrorCode    ErrorCode            `protobuf:"varint,100,opt,name=error_code,json=errorCode,proto3,enum=protobuf.ErrorCode" json:"error_code,omitempty"`
}

func (x *ReplyUserGamePosition) Reset() {
	*x = ReplyUserGamePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_user_server_protobuf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyUserGamePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyUserGamePosition) ProtoMessage() {}

func (x *ReplyUserGamePosition) ProtoReflect() protoreflect.Message {
	mi := &file_source_user_server_protobuf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyUserGamePosition.ProtoReflect.Descriptor instead.
func (*ReplyUserGamePosition) Descriptor() ([]byte, []int) {
	return file_source_user_server_protobuf_proto_rawDescGZIP(), []int{7}
}

func (x *ReplyUserGamePosition) GetGamePosition() *common.GamePosition {
	if x != nil {
		return x.GamePosition
	}
	return nil
}

func (x *ReplyUserGamePosition) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_SUCCESS
}

var File_source_user_server_protobuf_proto protoreflect.FileDescriptor

var file_source_user_server_protobuf_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0xb0, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65,
	0x78, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x73, 0x52, 0x65, 0x61, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x55, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x22,
	0x51, 0x0a, 0x1b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x4f, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x19, 0x0a,
	0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x2a, 0x85, 0x02, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x49, 0x44, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x50, 0x4c,
	0x59, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x24, 0x0a,
	0x20, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x41,
	0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x4d, 0x4f, 0x42,
	0x49, 0x4c, 0x45, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x1e, 0x0a,
	0x1a, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x41,
	0x4d, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x1c, 0x0a,
	0x18, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x42, 0x14, 0x5a, 0x12, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_user_server_protobuf_proto_rawDescOnce sync.Once
	file_source_user_server_protobuf_proto_rawDescData = file_source_user_server_protobuf_proto_rawDesc
)

func file_source_user_server_protobuf_proto_rawDescGZIP() []byte {
	file_source_user_server_protobuf_proto_rawDescOnce.Do(func() {
		file_source_user_server_protobuf_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_user_server_protobuf_proto_rawDescData)
	})
	return file_source_user_server_protobuf_proto_rawDescData
}

var file_source_user_server_protobuf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_source_user_server_protobuf_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_source_user_server_protobuf_proto_goTypes = []interface{}{
	(ProtocolID)(0),                       // 0: protobuf.protocolID
	(*RequestUserData)(nil),               // 1: protobuf.RequestUserData
	(*ReplyUserData)(nil),                 // 2: protobuf.ReplyUserData
	(*RequestRealNameAuthentication)(nil), // 3: protobuf.RequestRealNameAuthentication
	(*ReplyRealNameAuthentication)(nil),   // 4: protobuf.ReplyRealNameAuthentication
	(*RequestMobileBinding)(nil),          // 5: protobuf.RequestMobileBinding
	(*ReplyMobileBinding)(nil),            // 6: protobuf.ReplyMobileBinding
	(*RequestUserGamePosition)(nil),       // 7: protobuf.RequestUserGamePosition
	(*ReplyUserGamePosition)(nil),         // 8: protobuf.ReplyUserGamePosition
	(ErrorCode)(0),                        // 9: protobuf.ErrorCode
	(*common.GamePosition)(nil),           // 10: protobuf.GamePosition
}
var file_source_user_server_protobuf_proto_depIdxs = []int32{
	9,  // 0: protobuf.ReplyRealNameAuthentication.error_code:type_name -> protobuf.ErrorCode
	9,  // 1: protobuf.ReplyMobileBinding.error_code:type_name -> protobuf.ErrorCode
	10, // 2: protobuf.ReplyUserGamePosition.game_position:type_name -> protobuf.GamePosition
	9,  // 3: protobuf.ReplyUserGamePosition.error_code:type_name -> protobuf.ErrorCode
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_source_user_server_protobuf_proto_init() }
func file_source_user_server_protobuf_proto_init() {
	if File_source_user_server_protobuf_proto != nil {
		return
	}
	file_source_user_server_errors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_source_user_server_protobuf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUserData); i {
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
		file_source_user_server_protobuf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyUserData); i {
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
		file_source_user_server_protobuf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRealNameAuthentication); i {
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
		file_source_user_server_protobuf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyRealNameAuthentication); i {
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
		file_source_user_server_protobuf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMobileBinding); i {
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
		file_source_user_server_protobuf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMobileBinding); i {
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
		file_source_user_server_protobuf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUserGamePosition); i {
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
		file_source_user_server_protobuf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyUserGamePosition); i {
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
			RawDescriptor: file_source_user_server_protobuf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_source_user_server_protobuf_proto_goTypes,
		DependencyIndexes: file_source_user_server_protobuf_proto_depIdxs,
		EnumInfos:         file_source_user_server_protobuf_proto_enumTypes,
		MessageInfos:      file_source_user_server_protobuf_proto_msgTypes,
	}.Build()
	File_source_user_server_protobuf_proto = out.File
	file_source_user_server_protobuf_proto_rawDesc = nil
	file_source_user_server_protobuf_proto_goTypes = nil
	file_source_user_server_protobuf_proto_depIdxs = nil
}
