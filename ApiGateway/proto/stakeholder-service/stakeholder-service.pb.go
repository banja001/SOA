// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: stakeholder-service/stakeholder-service.proto

package stakeholder_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholder_service_stakeholder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholder_service_stakeholder_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_stakeholder_service_stakeholder_service_proto_rawDescGZIP(), []int{0}
}

func (x *Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthenticationTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *AuthenticationTokensResponse) Reset() {
	*x = AuthenticationTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholder_service_stakeholder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationTokensResponse) ProtoMessage() {}

func (x *AuthenticationTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholder_service_stakeholder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationTokensResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationTokensResponse) Descriptor() ([]byte, []int) {
	return file_stakeholder_service_stakeholder_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationTokensResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthenticationTokensResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_stakeholder_service_stakeholder_service_proto protoreflect.FileDescriptor

var file_stakeholder_service_stakeholder_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x50, 0x0a, 0x1c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x63, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x1a, 0x1d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x1b, 0x5a, 0x19,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_stakeholder_service_stakeholder_service_proto_rawDescOnce sync.Once
	file_stakeholder_service_stakeholder_service_proto_rawDescData = file_stakeholder_service_stakeholder_service_proto_rawDesc
)

func file_stakeholder_service_stakeholder_service_proto_rawDescGZIP() []byte {
	file_stakeholder_service_stakeholder_service_proto_rawDescOnce.Do(func() {
		file_stakeholder_service_stakeholder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stakeholder_service_stakeholder_service_proto_rawDescData)
	})
	return file_stakeholder_service_stakeholder_service_proto_rawDescData
}

var file_stakeholder_service_stakeholder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stakeholder_service_stakeholder_service_proto_goTypes = []interface{}{
	(*Credentials)(nil),                  // 0: Credentials
	(*AuthenticationTokensResponse)(nil), // 1: AuthenticationTokensResponse
}
var file_stakeholder_service_stakeholder_service_proto_depIdxs = []int32{
	0, // 0: StakeholderService.Login:input_type -> Credentials
	1, // 1: StakeholderService.Login:output_type -> AuthenticationTokensResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stakeholder_service_stakeholder_service_proto_init() }
func file_stakeholder_service_stakeholder_service_proto_init() {
	if File_stakeholder_service_stakeholder_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stakeholder_service_stakeholder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_stakeholder_service_stakeholder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationTokensResponse); i {
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
			RawDescriptor: file_stakeholder_service_stakeholder_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stakeholder_service_stakeholder_service_proto_goTypes,
		DependencyIndexes: file_stakeholder_service_stakeholder_service_proto_depIdxs,
		MessageInfos:      file_stakeholder_service_stakeholder_service_proto_msgTypes,
	}.Build()
	File_stakeholder_service_stakeholder_service_proto = out.File
	file_stakeholder_service_stakeholder_service_proto_rawDesc = nil
	file_stakeholder_service_stakeholder_service_proto_goTypes = nil
	file_stakeholder_service_stakeholder_service_proto_depIdxs = nil
}
