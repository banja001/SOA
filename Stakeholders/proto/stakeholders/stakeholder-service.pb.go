// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: stakeholders/stakeholder-service.proto

package stakeholders

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
		mi := &file_stakeholders_stakeholder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_stakeholder_service_proto_msgTypes[0]
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
	return file_stakeholders_stakeholder_service_proto_rawDescGZIP(), []int{0}
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

type AuthenticationTokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *AuthenticationTokens) Reset() {
	*x = AuthenticationTokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_stakeholder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationTokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationTokens) ProtoMessage() {}

func (x *AuthenticationTokens) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_stakeholder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationTokens.ProtoReflect.Descriptor instead.
func (*AuthenticationTokens) Descriptor() ([]byte, []int) {
	return file_stakeholders_stakeholder_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationTokens) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthenticationTokens) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_stakeholder_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_stakeholder_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_stakeholders_stakeholder_service_proto_rawDescGZIP(), []int{2}
}

func (x *AccessToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TokenClaims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jti      string `protobuf:"bytes,1,opt,name=jti,proto3" json:"jti,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	PersonId string `protobuf:"bytes,4,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	Role     string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Exp      int64  `protobuf:"varint,6,opt,name=exp,proto3" json:"exp,omitempty"`
	Iss      string `protobuf:"bytes,7,opt,name=iss,proto3" json:"iss,omitempty"`
	Aud      string `protobuf:"bytes,8,opt,name=aud,proto3" json:"aud,omitempty"`
}

func (x *TokenClaims) Reset() {
	*x = TokenClaims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stakeholders_stakeholder_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenClaims) ProtoMessage() {}

func (x *TokenClaims) ProtoReflect() protoreflect.Message {
	mi := &file_stakeholders_stakeholder_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenClaims.ProtoReflect.Descriptor instead.
func (*TokenClaims) Descriptor() ([]byte, []int) {
	return file_stakeholders_stakeholder_service_proto_rawDescGZIP(), []int{3}
}

func (x *TokenClaims) GetJti() string {
	if x != nil {
		return x.Jti
	}
	return ""
}

func (x *TokenClaims) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TokenClaims) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TokenClaims) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

func (x *TokenClaims) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *TokenClaims) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *TokenClaims) GetIss() string {
	if x != nil {
		return x.Iss
	}
	return ""
}

func (x *TokenClaims) GetAud() string {
	if x != nil {
		return x.Aud
	}
	return ""
}

var File_stakeholders_stakeholder_service_proto protoreflect.FileDescriptor

var file_stakeholders_stakeholder_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x48, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x23, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb2,
	0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6a, 0x74, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x74, 0x69,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x78, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x75, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x75, 0x64, 0x32, 0x75, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x1a, 0x15, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0c, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stakeholders_stakeholder_service_proto_rawDescOnce sync.Once
	file_stakeholders_stakeholder_service_proto_rawDescData = file_stakeholders_stakeholder_service_proto_rawDesc
)

func file_stakeholders_stakeholder_service_proto_rawDescGZIP() []byte {
	file_stakeholders_stakeholder_service_proto_rawDescOnce.Do(func() {
		file_stakeholders_stakeholder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stakeholders_stakeholder_service_proto_rawDescData)
	})
	return file_stakeholders_stakeholder_service_proto_rawDescData
}

var file_stakeholders_stakeholder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stakeholders_stakeholder_service_proto_goTypes = []interface{}{
	(*Credentials)(nil),          // 0: Credentials
	(*AuthenticationTokens)(nil), // 1: AuthenticationTokens
	(*AccessToken)(nil),          // 2: AccessToken
	(*TokenClaims)(nil),          // 3: TokenClaims
}
var file_stakeholders_stakeholder_service_proto_depIdxs = []int32{
	0, // 0: StakeholderService.Login:input_type -> Credentials
	2, // 1: StakeholderService.ValidateAccessToken:input_type -> AccessToken
	1, // 2: StakeholderService.Login:output_type -> AuthenticationTokens
	3, // 3: StakeholderService.ValidateAccessToken:output_type -> TokenClaims
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stakeholders_stakeholder_service_proto_init() }
func file_stakeholders_stakeholder_service_proto_init() {
	if File_stakeholders_stakeholder_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stakeholders_stakeholder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_stakeholders_stakeholder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationTokens); i {
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
		file_stakeholders_stakeholder_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessToken); i {
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
		file_stakeholders_stakeholder_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenClaims); i {
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
			RawDescriptor: file_stakeholders_stakeholder_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stakeholders_stakeholder_service_proto_goTypes,
		DependencyIndexes: file_stakeholders_stakeholder_service_proto_depIdxs,
		MessageInfos:      file_stakeholders_stakeholder_service_proto_msgTypes,
	}.Build()
	File_stakeholders_stakeholder_service_proto = out.File
	file_stakeholders_stakeholder_service_proto_rawDesc = nil
	file_stakeholders_stakeholder_service_proto_goTypes = nil
	file_stakeholders_stakeholder_service_proto_depIdxs = nil
}
