// (-- api-linter: core::0191::java-package=disabled
// (-- api-linter: core::0191::java-outer-classname=disabled
// (-- api-linter: core::0191::java-multiple-files=disabled
//     aip.dev/not-precedent: We're doing This as this is a Go server --)
// (-- api-linter: core::0191::proto-package=disabled

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: buildsafe/v1/app.proto

package bsfv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type RegisterAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	RepoId    uint64 `protobuf:"varint,2,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
}

func (x *RegisterAppRequest) Reset() {
	*x = RegisterAppRequest{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAppRequest) ProtoMessage() {}

func (x *RegisterAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAppRequest.ProtoReflect.Descriptor instead.
func (*RegisterAppRequest) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterAppRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *RegisterAppRequest) GetRepoId() uint64 {
	if x != nil {
		return x.RepoId
	}
	return 0
}

type RegisterAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId uint64 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *RegisterAppResponse) Reset() {
	*x = RegisterAppResponse{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAppResponse) ProtoMessage() {}

func (x *RegisterAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAppResponse.ProtoReflect.Descriptor instead.
func (*RegisterAppResponse) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterAppResponse) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type GetAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	AppId     uint64 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetAppRequest) Reset() {
	*x = GetAppRequest{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRequest) ProtoMessage() {}

func (x *GetAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRequest.ProtoReflect.Descriptor instead.
func (*GetAppRequest) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{2}
}

func (x *GetAppRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *GetAppRequest) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     uint64  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ProjectId uint64  `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	RepoId    uint64  `protobuf:"varint,3,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	Branch    *string `protobuf:"bytes,4,opt,name=branch,proto3,oneof" json:"branch,omitempty"`
	Folder    *string `protobuf:"bytes,5,opt,name=folder,proto3,oneof" json:"folder,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{3}
}

func (x *App) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *App) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *App) GetRepoId() uint64 {
	if x != nil {
		return x.RepoId
	}
	return 0
}

func (x *App) GetBranch() string {
	if x != nil && x.Branch != nil {
		return *x.Branch
	}
	return ""
}

func (x *App) GetFolder() string {
	if x != nil && x.Folder != nil {
		return *x.Folder
	}
	return ""
}

type GetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *GetAppResponse) Reset() {
	*x = GetAppResponse{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppResponse) ProtoMessage() {}

func (x *GetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppResponse.ProtoReflect.Descriptor instead.
func (*GetAppResponse) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppResponse) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type ListAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ListAppsRequest) Reset() {
	*x = ListAppsRequest{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppsRequest) ProtoMessage() {}

func (x *ListAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppsRequest.ProtoReflect.Descriptor instead.
func (*ListAppsRequest) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{5}
}

func (x *ListAppsRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type ListAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*App `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
}

func (x *ListAppsResponse) Reset() {
	*x = ListAppsResponse{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppsResponse) ProtoMessage() {}

func (x *ListAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppsResponse.ProtoReflect.Descriptor instead.
func (*ListAppsResponse) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{6}
}

func (x *ListAppsResponse) GetApps() []*App {
	if x != nil {
		return x.Apps
	}
	return nil
}

type UnregisterAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	AppId     uint64 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *UnregisterAppRequest) Reset() {
	*x = UnregisterAppRequest{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnregisterAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterAppRequest) ProtoMessage() {}

func (x *UnregisterAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterAppRequest.ProtoReflect.Descriptor instead.
func (*UnregisterAppRequest) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{7}
}

func (x *UnregisterAppRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *UnregisterAppRequest) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type UnregisterAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnregisterAppResponse) Reset() {
	*x = UnregisterAppResponse{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnregisterAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterAppResponse) ProtoMessage() {}

func (x *UnregisterAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterAppResponse.ProtoReflect.Descriptor instead.
func (*UnregisterAppResponse) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{8}
}

type AddAppDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	AppId     uint64 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Branch    string `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	Folder    string `protobuf:"bytes,4,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *AddAppDetailsRequest) Reset() {
	*x = AddAppDetailsRequest{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAppDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppDetailsRequest) ProtoMessage() {}

func (x *AddAppDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppDetailsRequest.ProtoReflect.Descriptor instead.
func (*AddAppDetailsRequest) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{9}
}

func (x *AddAppDetailsRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *AddAppDetailsRequest) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AddAppDetailsRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *AddAppDetailsRequest) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

type AddAppDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddAppDetailsResponse) Reset() {
	*x = AddAppDetailsResponse{}
	mi := &file_buildsafe_v1_app_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAppDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAppDetailsResponse) ProtoMessage() {}

func (x *AddAppDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildsafe_v1_app_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAppDetailsResponse.ProtoReflect.Descriptor instead.
func (*AddAppDetailsResponse) Descriptor() ([]byte, []int) {
	return file_buildsafe_v1_app_proto_rawDescGZIP(), []int{10}
}

var File_buildsafe_v1_app_proto protoreflect.FileDescriptor

var file_buildsafe_v1_app_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73,
	0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6c, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b,
	0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32,
	0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x22,
	0x3c, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06,
	0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x65, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84, 0x3d,
	0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba,
	0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x22, 0xd4, 0x01, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x25, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48,
	0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32,
	0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84,
	0x3d, 0x20, 0x00, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61,
	0x70, 0x70, 0x22, 0x40, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01,
	0x01, 0x32, 0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61,
	0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22,
	0x6c, 0x0a, 0x14, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b,
	0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06,
	0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x41, 0x70,
	0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84,
	0x3d, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0e,
	0xba, 0x48, 0x0b, 0xc8, 0x01, 0x01, 0x32, 0x06, 0x10, 0xc0, 0x84, 0x3d, 0x20, 0x00, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x72, 0x05, 0x10,
	0x01, 0x18, 0x80, 0x01, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x25, 0x0a, 0x06,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48,
	0x0a, 0xc8, 0x01, 0x01, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x06, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfa, 0x04, 0x0a,
	0x0a, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0b, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x12, 0x20, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x3d, 0x2a, 0x7d, 0x12, 0x6f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x12, 0x1b, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24,
	0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x6a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73,
	0x12, 0x1d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d,
	0x12, 0x84, 0x01, 0x0a, 0x0d, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x12, 0x22, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61,
	0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x24, 0x2a, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x7b, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x8f, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x41,
	0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x61, 0x66,
	0x65, 0x64, 0x65, 0x76, 0x2f, 0x62, 0x73, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x3b, 0x62, 0x73,
	0x66, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buildsafe_v1_app_proto_rawDescOnce sync.Once
	file_buildsafe_v1_app_proto_rawDescData = file_buildsafe_v1_app_proto_rawDesc
)

func file_buildsafe_v1_app_proto_rawDescGZIP() []byte {
	file_buildsafe_v1_app_proto_rawDescOnce.Do(func() {
		file_buildsafe_v1_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_buildsafe_v1_app_proto_rawDescData)
	})
	return file_buildsafe_v1_app_proto_rawDescData
}

var file_buildsafe_v1_app_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_buildsafe_v1_app_proto_goTypes = []any{
	(*RegisterAppRequest)(nil),    // 0: buildsafe.v1.RegisterAppRequest
	(*RegisterAppResponse)(nil),   // 1: buildsafe.v1.RegisterAppResponse
	(*GetAppRequest)(nil),         // 2: buildsafe.v1.GetAppRequest
	(*App)(nil),                   // 3: buildsafe.v1.App
	(*GetAppResponse)(nil),        // 4: buildsafe.v1.GetAppResponse
	(*ListAppsRequest)(nil),       // 5: buildsafe.v1.ListAppsRequest
	(*ListAppsResponse)(nil),      // 6: buildsafe.v1.ListAppsResponse
	(*UnregisterAppRequest)(nil),  // 7: buildsafe.v1.UnregisterAppRequest
	(*UnregisterAppResponse)(nil), // 8: buildsafe.v1.UnregisterAppResponse
	(*AddAppDetailsRequest)(nil),  // 9: buildsafe.v1.AddAppDetailsRequest
	(*AddAppDetailsResponse)(nil), // 10: buildsafe.v1.AddAppDetailsResponse
}
var file_buildsafe_v1_app_proto_depIdxs = []int32{
	3,  // 0: buildsafe.v1.GetAppResponse.app:type_name -> buildsafe.v1.App
	3,  // 1: buildsafe.v1.ListAppsResponse.apps:type_name -> buildsafe.v1.App
	0,  // 2: buildsafe.v1.AppService.RegisterApp:input_type -> buildsafe.v1.RegisterAppRequest
	2,  // 3: buildsafe.v1.AppService.GetApp:input_type -> buildsafe.v1.GetAppRequest
	5,  // 4: buildsafe.v1.AppService.ListApps:input_type -> buildsafe.v1.ListAppsRequest
	7,  // 5: buildsafe.v1.AppService.UnregisterApp:input_type -> buildsafe.v1.UnregisterAppRequest
	9,  // 6: buildsafe.v1.AppService.AddAppDetails:input_type -> buildsafe.v1.AddAppDetailsRequest
	1,  // 7: buildsafe.v1.AppService.RegisterApp:output_type -> buildsafe.v1.RegisterAppResponse
	4,  // 8: buildsafe.v1.AppService.GetApp:output_type -> buildsafe.v1.GetAppResponse
	6,  // 9: buildsafe.v1.AppService.ListApps:output_type -> buildsafe.v1.ListAppsResponse
	8,  // 10: buildsafe.v1.AppService.UnregisterApp:output_type -> buildsafe.v1.UnregisterAppResponse
	10, // 11: buildsafe.v1.AppService.AddAppDetails:output_type -> buildsafe.v1.AddAppDetailsResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_buildsafe_v1_app_proto_init() }
func file_buildsafe_v1_app_proto_init() {
	if File_buildsafe_v1_app_proto != nil {
		return
	}
	file_buildsafe_v1_app_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buildsafe_v1_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buildsafe_v1_app_proto_goTypes,
		DependencyIndexes: file_buildsafe_v1_app_proto_depIdxs,
		MessageInfos:      file_buildsafe_v1_app_proto_msgTypes,
	}.Build()
	File_buildsafe_v1_app_proto = out.File
	file_buildsafe_v1_app_proto_rawDesc = nil
	file_buildsafe_v1_app_proto_goTypes = nil
	file_buildsafe_v1_app_proto_depIdxs = nil
}
