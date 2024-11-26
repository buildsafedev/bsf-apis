// (-- api-linter: core::0191::java-package=disabled
// (-- api-linter: core::0191::java-outer-classname=disabled
// (-- api-linter: core::0191::java-multiple-files=disabled
//     aip.dev/not-precedent: We're doing This as this is a Go server --)
// (-- api-linter: core::0191::proto-package=disabled

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: buildsafe/v1/auth.proto

package bsfv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Provider int32

const (
	Provider_PROVIDER_UNSPECIFIED Provider = 0
	Provider_PROVIDER_GITHUB      Provider = 1
)

// Enum value maps for Provider.
var (
	Provider_name = map[int32]string{
		0: "PROVIDER_UNSPECIFIED",
		1: "PROVIDER_GITHUB",
	}
	Provider_value = map[string]int32{
		"PROVIDER_UNSPECIFIED": 0,
		"PROVIDER_GITHUB":      1,
	}
)

func (x Provider) Enum() *Provider {
	p := new(Provider)
	*p = x
	return p
}

func (x Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_buildsafe_v1_auth_proto_enumTypes[0].Descriptor()
}

func (Provider) Type() protoreflect.EnumType {
	return &file_buildsafe_v1_auth_proto_enumTypes[0]
}

func (x Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Provider.Descriptor instead.
func (Provider) EnumDescriptor() ([]byte, []int) {
	return file_buildsafe_v1_auth_proto_rawDescGZIP(), []int{0}
}

// OAuthAuthenticateRequest is the request message for OAuthAuthenticate.
type OAuthAuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider Provider `protobuf:"varint,1,opt,name=provider,proto3,enum=buildsafe.v1.Provider" json:"provider,omitempty"` // e.g., "google", "github"
}

func (x *OAuthAuthenticateRequest) Reset() {
	*x = OAuthAuthenticateRequest{}
	mi := &file_buildsafe_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuthAuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthAuthenticateRequest) ProtoMessage() {}

func (x *OAuthAuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthAuthenticateRequest.ProtoReflect.Descriptor instead.
func (*OAuthAuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *OAuthAuthenticateRequest) GetProvider() Provider {
	if x != nil {
		return x.Provider
	}
	return Provider_PROVIDER_UNSPECIFIED
}

// Response message for OAuth-based authentication
type OAuthAuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // The authentication token
}

func (x *OAuthAuthenticateResponse) Reset() {
	*x = OAuthAuthenticateResponse{}
	mi := &file_buildsafe_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuthAuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthAuthenticateResponse) ProtoMessage() {}

func (x *OAuthAuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthAuthenticateResponse.ProtoReflect.Descriptor instead.
func (*OAuthAuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthAuthenticateResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IdentityRequest) Reset() {
	*x = IdentityRequest{}
	mi := &file_buildsafe_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityRequest) ProtoMessage() {}

func (x *IdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityRequest.ProtoReflect.Descriptor instead.
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_auth_proto_rawDescGZIP(), []int{2}
}

type IdentityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *IdentityResponse) Reset() {
	*x = IdentityResponse{}
	mi := &file_buildsafe_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityResponse) ProtoMessage() {}

func (x *IdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityResponse.ProtoReflect.Descriptor instead.
func (*IdentityResponse) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *IdentityResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *IdentityResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IdentityResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_buildsafe_v1_auth_proto protoreflect.FileDescriptor

var file_buildsafe_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x18, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x19, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x11, 0x0a,
	0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x55, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x39, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42,
	0x10, 0x01, 0x32, 0xda, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x65, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x04, 0x4a, 0x57, 0x4b,
	0x53, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x2d,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x6a, 0x77, 0x6b, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x64, 0x65, 0x76, 0x2f, 0x62, 0x73, 0x66, 0x2d, 0x61,
	0x70, 0x69, 0x73, 0x3b, 0x62, 0x73, 0x66, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_buildsafe_v1_auth_proto_rawDescOnce sync.Once
	file_buildsafe_v1_auth_proto_rawDescData = file_buildsafe_v1_auth_proto_rawDesc
)

func file_buildsafe_v1_auth_proto_rawDescGZIP() []byte {
	file_buildsafe_v1_auth_proto_rawDescOnce.Do(func() {
		file_buildsafe_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_buildsafe_v1_auth_proto_rawDescData)
	})
	return file_buildsafe_v1_auth_proto_rawDescData
}

var file_buildsafe_v1_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buildsafe_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buildsafe_v1_auth_proto_goTypes = []any{
	(Provider)(0),                     // 0: buildsafe.v1.Provider
	(*OAuthAuthenticateRequest)(nil),  // 1: buildsafe.v1.OAuthAuthenticateRequest
	(*OAuthAuthenticateResponse)(nil), // 2: buildsafe.v1.OAuthAuthenticateResponse
	(*IdentityRequest)(nil),           // 3: buildsafe.v1.IdentityRequest
	(*IdentityResponse)(nil),          // 4: buildsafe.v1.IdentityResponse
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
	(*httpbody.HttpBody)(nil),         // 6: google.api.HttpBody
}
var file_buildsafe_v1_auth_proto_depIdxs = []int32{
	0, // 0: buildsafe.v1.OAuthAuthenticateRequest.provider:type_name -> buildsafe.v1.Provider
	1, // 1: buildsafe.v1.AuthService.OAuthAuthenticate:input_type -> buildsafe.v1.OAuthAuthenticateRequest
	3, // 2: buildsafe.v1.AuthService.Identity:input_type -> buildsafe.v1.IdentityRequest
	5, // 3: buildsafe.v1.AuthService.JWKS:input_type -> google.protobuf.Empty
	2, // 4: buildsafe.v1.AuthService.OAuthAuthenticate:output_type -> buildsafe.v1.OAuthAuthenticateResponse
	4, // 5: buildsafe.v1.AuthService.Identity:output_type -> buildsafe.v1.IdentityResponse
	6, // 6: buildsafe.v1.AuthService.JWKS:output_type -> google.api.HttpBody
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buildsafe_v1_auth_proto_init() }
func file_buildsafe_v1_auth_proto_init() {
	if File_buildsafe_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buildsafe_v1_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buildsafe_v1_auth_proto_goTypes,
		DependencyIndexes: file_buildsafe_v1_auth_proto_depIdxs,
		EnumInfos:         file_buildsafe_v1_auth_proto_enumTypes,
		MessageInfos:      file_buildsafe_v1_auth_proto_msgTypes,
	}.Build()
	File_buildsafe_v1_auth_proto = out.File
	file_buildsafe_v1_auth_proto_rawDesc = nil
	file_buildsafe_v1_auth_proto_goTypes = nil
	file_buildsafe_v1_auth_proto_depIdxs = nil
}
