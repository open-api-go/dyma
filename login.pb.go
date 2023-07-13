// 抖音小程序开放接口-登录 douyin-miniapp

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.0
// source: douyin-miniapp/login.proto

package douyin_miniapp

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Code2SessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid         string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`                                      // 小程序 ID
	Secret        string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`                                    // 小程序的 APP Secret，可以在开发者后台获取
	Code          string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`                                        // login 接口返回的登录凭证
	AnonymousCode string `protobuf:"bytes,4,opt,name=anonymous_code,json=anonymousCode,proto3" json:"anonymous_code,omitempty"` // login 接口返回的匿名登录凭证 code 和 anonymous_code 至少要有一个
}

func (x *Code2SessionReq) Reset() {
	*x = Code2SessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_miniapp_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code2SessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code2SessionReq) ProtoMessage() {}

func (x *Code2SessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_miniapp_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code2SessionReq.ProtoReflect.Descriptor instead.
func (*Code2SessionReq) Descriptor() ([]byte, []int) {
	return file_douyin_miniapp_login_proto_rawDescGZIP(), []int{0}
}

func (x *Code2SessionReq) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *Code2SessionReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Code2SessionReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Code2SessionReq) GetAnonymousCode() string {
	if x != nil {
		return x.AnonymousCode
	}
	return ""
}

type Code2SessionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo           int64  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips         string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
	SessionKey      string `protobuf:"bytes,10,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`                // 会话密钥，如果请求时有 code 参数才会返回
	Openid          string `protobuf:"bytes,11,opt,name=openid,proto3" json:"openid,omitempty"`                                          // 用户在当前小程序的 ID，如果请求时有 code 参数才会返回
	AnonymousOpenid string `protobuf:"bytes,12,opt,name=anonymous_openid,json=anonymousOpenid,proto3" json:"anonymous_openid,omitempty"` // 匿名用户在当前小程序的 ID，如果请求时有 anonymous_code 参数才会返回
	Unionid         string `protobuf:"bytes,13,opt,name=unionid,proto3" json:"unionid,omitempty"`                                        // 用户在小程序平台的唯一标识符，请求时有 code 参数才会返回。如果开发者拥有多个小程序，可通过 unionid 来区分用户的唯一性
}

func (x *Code2SessionRes) Reset() {
	*x = Code2SessionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_miniapp_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code2SessionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code2SessionRes) ProtoMessage() {}

func (x *Code2SessionRes) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_miniapp_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code2SessionRes.ProtoReflect.Descriptor instead.
func (*Code2SessionRes) Descriptor() ([]byte, []int) {
	return file_douyin_miniapp_login_proto_rawDescGZIP(), []int{1}
}

func (x *Code2SessionRes) GetErrNo() int64 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *Code2SessionRes) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

func (x *Code2SessionRes) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *Code2SessionRes) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *Code2SessionRes) GetAnonymousOpenid() string {
	if x != nil {
		return x.AnonymousOpenid
	}
	return ""
}

func (x *Code2SessionRes) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

var File_douyin_miniapp_login_proto protoreflect.FileDescriptor

var file_douyin_miniapp_login_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x61, 0x70, 0x70,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x61, 0x6f, 0x2e,
	0x63, 0x6f, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7a, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xc1, 0x01,
	0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f,
	0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x54,
	0x69, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69,
	0x64, 0x32, 0x97, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2e,
	0x74, 0x6f, 0x75, 0x74, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x61, 0x6f, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2d, 0x6d, 0x69, 0x6e,
	0x69, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_douyin_miniapp_login_proto_rawDescOnce sync.Once
	file_douyin_miniapp_login_proto_rawDescData = file_douyin_miniapp_login_proto_rawDesc
)

func file_douyin_miniapp_login_proto_rawDescGZIP() []byte {
	file_douyin_miniapp_login_proto_rawDescOnce.Do(func() {
		file_douyin_miniapp_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_douyin_miniapp_login_proto_rawDescData)
	})
	return file_douyin_miniapp_login_proto_rawDescData
}

var file_douyin_miniapp_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_douyin_miniapp_login_proto_goTypes = []interface{}{
	(*Code2SessionReq)(nil), // 0: developer.toutiao.com.Code2SessionReq
	(*Code2SessionRes)(nil), // 1: developer.toutiao.com.Code2SessionRes
}
var file_douyin_miniapp_login_proto_depIdxs = []int32{
	0, // 0: developer.toutiao.com.LoginService.Code2Session:input_type -> developer.toutiao.com.Code2SessionReq
	1, // 1: developer.toutiao.com.LoginService.Code2Session:output_type -> developer.toutiao.com.Code2SessionRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_douyin_miniapp_login_proto_init() }
func file_douyin_miniapp_login_proto_init() {
	if File_douyin_miniapp_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_douyin_miniapp_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code2SessionReq); i {
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
		file_douyin_miniapp_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code2SessionRes); i {
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
			RawDescriptor: file_douyin_miniapp_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_douyin_miniapp_login_proto_goTypes,
		DependencyIndexes: file_douyin_miniapp_login_proto_depIdxs,
		MessageInfos:      file_douyin_miniapp_login_proto_msgTypes,
	}.Build()
	File_douyin_miniapp_login_proto = out.File
	file_douyin_miniapp_login_proto_rawDesc = nil
	file_douyin_miniapp_login_proto_goTypes = nil
	file_douyin_miniapp_login_proto_depIdxs = nil
}
