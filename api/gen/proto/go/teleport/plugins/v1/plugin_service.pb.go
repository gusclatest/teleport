// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: teleport/plugins/v1/plugin_service.proto

package pluginsv1

import (
	types "github.com/gravitational/teleport/api/types"
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

// PluginType represents a single type of hosted plugin
// that can be onboarded.
type PluginType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is a string corresponding to api.PluginTypeXXX constants
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// OAuthClientID contains the client ID of the OAuth application
	// that is used with this plugin's API provider.
	// For plugins that are not authenticated via OAuth,
	// this will be empty.
	OauthClientId string `protobuf:"bytes,2,opt,name=oauth_client_id,json=oauthClientId,proto3" json:"oauth_client_id,omitempty"`
}

func (x *PluginType) Reset() {
	*x = PluginType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginType) ProtoMessage() {}

func (x *PluginType) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginType.ProtoReflect.Descriptor instead.
func (*PluginType) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{0}
}

func (x *PluginType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PluginType) GetOauthClientId() string {
	if x != nil {
		return x.OauthClientId
	}
	return ""
}

// CreatePluginRequest creates a new plugin from the given spec and initial credentials.
type CreatePluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugin is the plugin object without live credentials.
	Plugin *types.PluginV1 `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	// BootstrapCredentials are the initial credentials
	// issued by e.g. OAuth2 authorization code flow.
	// In the scope of processing this request, these are exchanged for
	// short-lived renewable credentials, which are stored in the Plugin.
	BootstrapCredentials *types.PluginBootstrapCredentialsV1 `protobuf:"bytes,2,opt,name=bootstrap_credentials,json=bootstrapCredentials,proto3" json:"bootstrap_credentials,omitempty"`
}

func (x *CreatePluginRequest) Reset() {
	*x = CreatePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePluginRequest) ProtoMessage() {}

func (x *CreatePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePluginRequest.ProtoReflect.Descriptor instead.
func (*CreatePluginRequest) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePluginRequest) GetPlugin() *types.PluginV1 {
	if x != nil {
		return x.Plugin
	}
	return nil
}

func (x *CreatePluginRequest) GetBootstrapCredentials() *types.PluginBootstrapCredentialsV1 {
	if x != nil {
		return x.BootstrapCredentials
	}
	return nil
}

// GetPluginRequest is a request to return a plugin instance by name.
type GetPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the plugin instance.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPluginRequest) Reset() {
	*x = GetPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginRequest) ProtoMessage() {}

func (x *GetPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginRequest.ProtoReflect.Descriptor instead.
func (*GetPluginRequest) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPluginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListPluginsRequest is a paginated request to list all plugin instances.
type ListPluginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageSize is the maximum number of plugins to return in a single response.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// StartKey is the value of NextKey received in the last ListPluginsResponse.
	// When making the initial request, this should be left empty.
	StartKey string `protobuf:"bytes,2,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
}

func (x *ListPluginsRequest) Reset() {
	*x = ListPluginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginsRequest) ProtoMessage() {}

func (x *ListPluginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginsRequest.ProtoReflect.Descriptor instead.
func (*ListPluginsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListPluginsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPluginsRequest) GetStartKey() string {
	if x != nil {
		return x.StartKey
	}
	return ""
}

// ListPluginsResponse is a paginated response to a ListPluginsRequest.
type ListPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugins is the list of plugins.
	Plugins []*types.PluginV1 `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
	// NextKey is a token to retrieve the next page of results, or empty
	// if there are no more results.
	NextKey string `protobuf:"bytes,2,opt,name=next_key,json=nextKey,proto3" json:"next_key,omitempty"`
}

func (x *ListPluginsResponse) Reset() {
	*x = ListPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginsResponse) ProtoMessage() {}

func (x *ListPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginsResponse.ProtoReflect.Descriptor instead.
func (*ListPluginsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListPluginsResponse) GetPlugins() []*types.PluginV1 {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *ListPluginsResponse) GetNextKey() string {
	if x != nil {
		return x.NextKey
	}
	return ""
}

// DeletePluginRequest is a request to delete a plugin instance by name.
type DeletePluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the plugin instance.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePluginRequest) Reset() {
	*x = DeletePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePluginRequest) ProtoMessage() {}

func (x *DeletePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePluginRequest.ProtoReflect.Descriptor instead.
func (*DeletePluginRequest) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePluginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SetPluginCredentialsRequest is a request to set credentials for an existing plugin
type SetPluginCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the plugin instance.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Credentials are the credentials obtained after exchanging the initial credentials,
	// and after successive credential renewals.
	Credentials *types.PluginCredentialsV1 `protobuf:"bytes,2,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (x *SetPluginCredentialsRequest) Reset() {
	*x = SetPluginCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPluginCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPluginCredentialsRequest) ProtoMessage() {}

func (x *SetPluginCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPluginCredentialsRequest.ProtoReflect.Descriptor instead.
func (*SetPluginCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{6}
}

func (x *SetPluginCredentialsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetPluginCredentialsRequest) GetCredentials() *types.PluginCredentialsV1 {
	if x != nil {
		return x.Credentials
	}
	return nil
}

// SetPluginStatusRequest is a request to set the status for an existing plugin
type SetPluginStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the plugin instance.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Status is the plugin status.
	Status *types.PluginStatusV1 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetPluginStatusRequest) Reset() {
	*x = SetPluginStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPluginStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPluginStatusRequest) ProtoMessage() {}

func (x *SetPluginStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPluginStatusRequest.ProtoReflect.Descriptor instead.
func (*SetPluginStatusRequest) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{7}
}

func (x *SetPluginStatusRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetPluginStatusRequest) GetStatus() *types.PluginStatusV1 {
	if x != nil {
		return x.Status
	}
	return nil
}

// GetAvailablePluginTypesRequest is the request type for GetAvailablePluginTypes
type GetAvailablePluginTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAvailablePluginTypesRequest) Reset() {
	*x = GetAvailablePluginTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailablePluginTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailablePluginTypesRequest) ProtoMessage() {}

func (x *GetAvailablePluginTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailablePluginTypesRequest.ProtoReflect.Descriptor instead.
func (*GetAvailablePluginTypesRequest) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{8}
}

// GetAvailablePluginTypesResponse is a response to for GetAvailablePluginTypes
type GetAvailablePluginTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PluginTypes is a list of hosted plugins
	// that the auth service supports.
	PluginTypes []*PluginType `protobuf:"bytes,1,rep,name=plugin_types,json=pluginTypes,proto3" json:"plugin_types,omitempty"`
}

func (x *GetAvailablePluginTypesResponse) Reset() {
	*x = GetAvailablePluginTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailablePluginTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailablePluginTypesResponse) ProtoMessage() {}

func (x *GetAvailablePluginTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_plugins_v1_plugin_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailablePluginTypesResponse.ProtoReflect.Descriptor instead.
func (*GetAvailablePluginTypesResponse) Descriptor() ([]byte, []int) {
	return file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetAvailablePluginTypesResponse) GetPluginTypes() []*PluginType {
	if x != nil {
		return x.PluginTypes
	}
	return nil
}

var File_teleport_plugins_v1_plugin_service_proto protoreflect.FileDescriptor

var file_teleport_plugins_v1_plugin_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x48, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x56, 0x31, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x58, 0x0a, 0x15, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x56, 0x31, 0x52, 0x14,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x5b, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x56, 0x31, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x56, 0x31, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x5b, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x31, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x32, 0x9b, 0x05, 0x0a, 0x0d,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x28, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x43, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x56, 0x31, 0x12, 0x50, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x30, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x56, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x84, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x33,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_plugins_v1_plugin_service_proto_rawDescOnce sync.Once
	file_teleport_plugins_v1_plugin_service_proto_rawDescData = file_teleport_plugins_v1_plugin_service_proto_rawDesc
)

func file_teleport_plugins_v1_plugin_service_proto_rawDescGZIP() []byte {
	file_teleport_plugins_v1_plugin_service_proto_rawDescOnce.Do(func() {
		file_teleport_plugins_v1_plugin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_plugins_v1_plugin_service_proto_rawDescData)
	})
	return file_teleport_plugins_v1_plugin_service_proto_rawDescData
}

var file_teleport_plugins_v1_plugin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_teleport_plugins_v1_plugin_service_proto_goTypes = []interface{}{
	(*PluginType)(nil),                         // 0: teleport.plugins.v1.PluginType
	(*CreatePluginRequest)(nil),                // 1: teleport.plugins.v1.CreatePluginRequest
	(*GetPluginRequest)(nil),                   // 2: teleport.plugins.v1.GetPluginRequest
	(*ListPluginsRequest)(nil),                 // 3: teleport.plugins.v1.ListPluginsRequest
	(*ListPluginsResponse)(nil),                // 4: teleport.plugins.v1.ListPluginsResponse
	(*DeletePluginRequest)(nil),                // 5: teleport.plugins.v1.DeletePluginRequest
	(*SetPluginCredentialsRequest)(nil),        // 6: teleport.plugins.v1.SetPluginCredentialsRequest
	(*SetPluginStatusRequest)(nil),             // 7: teleport.plugins.v1.SetPluginStatusRequest
	(*GetAvailablePluginTypesRequest)(nil),     // 8: teleport.plugins.v1.GetAvailablePluginTypesRequest
	(*GetAvailablePluginTypesResponse)(nil),    // 9: teleport.plugins.v1.GetAvailablePluginTypesResponse
	(*types.PluginV1)(nil),                     // 10: types.PluginV1
	(*types.PluginBootstrapCredentialsV1)(nil), // 11: types.PluginBootstrapCredentialsV1
	(*types.PluginCredentialsV1)(nil),          // 12: types.PluginCredentialsV1
	(*types.PluginStatusV1)(nil),               // 13: types.PluginStatusV1
	(*emptypb.Empty)(nil),                      // 14: google.protobuf.Empty
}
var file_teleport_plugins_v1_plugin_service_proto_depIdxs = []int32{
	10, // 0: teleport.plugins.v1.CreatePluginRequest.plugin:type_name -> types.PluginV1
	11, // 1: teleport.plugins.v1.CreatePluginRequest.bootstrap_credentials:type_name -> types.PluginBootstrapCredentialsV1
	10, // 2: teleport.plugins.v1.ListPluginsResponse.plugins:type_name -> types.PluginV1
	12, // 3: teleport.plugins.v1.SetPluginCredentialsRequest.credentials:type_name -> types.PluginCredentialsV1
	13, // 4: teleport.plugins.v1.SetPluginStatusRequest.status:type_name -> types.PluginStatusV1
	0,  // 5: teleport.plugins.v1.GetAvailablePluginTypesResponse.plugin_types:type_name -> teleport.plugins.v1.PluginType
	1,  // 6: teleport.plugins.v1.PluginService.CreatePlugin:input_type -> teleport.plugins.v1.CreatePluginRequest
	2,  // 7: teleport.plugins.v1.PluginService.GetPlugin:input_type -> teleport.plugins.v1.GetPluginRequest
	5,  // 8: teleport.plugins.v1.PluginService.DeletePlugin:input_type -> teleport.plugins.v1.DeletePluginRequest
	3,  // 9: teleport.plugins.v1.PluginService.ListPlugins:input_type -> teleport.plugins.v1.ListPluginsRequest
	6,  // 10: teleport.plugins.v1.PluginService.SetPluginCredentials:input_type -> teleport.plugins.v1.SetPluginCredentialsRequest
	7,  // 11: teleport.plugins.v1.PluginService.SetPluginStatus:input_type -> teleport.plugins.v1.SetPluginStatusRequest
	8,  // 12: teleport.plugins.v1.PluginService.GetAvailablePluginTypes:input_type -> teleport.plugins.v1.GetAvailablePluginTypesRequest
	14, // 13: teleport.plugins.v1.PluginService.CreatePlugin:output_type -> google.protobuf.Empty
	10, // 14: teleport.plugins.v1.PluginService.GetPlugin:output_type -> types.PluginV1
	14, // 15: teleport.plugins.v1.PluginService.DeletePlugin:output_type -> google.protobuf.Empty
	4,  // 16: teleport.plugins.v1.PluginService.ListPlugins:output_type -> teleport.plugins.v1.ListPluginsResponse
	14, // 17: teleport.plugins.v1.PluginService.SetPluginCredentials:output_type -> google.protobuf.Empty
	14, // 18: teleport.plugins.v1.PluginService.SetPluginStatus:output_type -> google.protobuf.Empty
	9,  // 19: teleport.plugins.v1.PluginService.GetAvailablePluginTypes:output_type -> teleport.plugins.v1.GetAvailablePluginTypesResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_teleport_plugins_v1_plugin_service_proto_init() }
func file_teleport_plugins_v1_plugin_service_proto_init() {
	if File_teleport_plugins_v1_plugin_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginType); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePluginRequest); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginRequest); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginsRequest); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginsResponse); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePluginRequest); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPluginCredentialsRequest); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPluginStatusRequest); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailablePluginTypesRequest); i {
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
		file_teleport_plugins_v1_plugin_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailablePluginTypesResponse); i {
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
			RawDescriptor: file_teleport_plugins_v1_plugin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_plugins_v1_plugin_service_proto_goTypes,
		DependencyIndexes: file_teleport_plugins_v1_plugin_service_proto_depIdxs,
		MessageInfos:      file_teleport_plugins_v1_plugin_service_proto_msgTypes,
	}.Build()
	File_teleport_plugins_v1_plugin_service_proto = out.File
	file_teleport_plugins_v1_plugin_service_proto_rawDesc = nil
	file_teleport_plugins_v1_plugin_service_proto_goTypes = nil
	file_teleport_plugins_v1_plugin_service_proto_depIdxs = nil
}
