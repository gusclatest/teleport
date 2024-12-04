// Copyright 2024 Gravitational, Inc.
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
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: teleport/autoupdate/v1/autoupdate.proto

package autoupdate

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
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

// AutoUpdateConfig is a config singleton used to configure cluster
// autoupdate settings.
type AutoUpdateConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     string                `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	SubKind  string                `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	Version  string                `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Metadata *v1.Metadata          `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *AutoUpdateConfigSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *AutoUpdateConfig) Reset() {
	*x = AutoUpdateConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateConfig) ProtoMessage() {}

func (x *AutoUpdateConfig) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateConfig.ProtoReflect.Descriptor instead.
func (*AutoUpdateConfig) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{0}
}

func (x *AutoUpdateConfig) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AutoUpdateConfig) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *AutoUpdateConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AutoUpdateConfig) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AutoUpdateConfig) GetSpec() *AutoUpdateConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// AutoUpdateConfigSpec encodes the parameters of the autoupdate config object.
type AutoUpdateConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tools *AutoUpdateConfigSpecTools `protobuf:"bytes,2,opt,name=tools,proto3" json:"tools,omitempty"`
}

func (x *AutoUpdateConfigSpec) Reset() {
	*x = AutoUpdateConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateConfigSpec) ProtoMessage() {}

func (x *AutoUpdateConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateConfigSpec.ProtoReflect.Descriptor instead.
func (*AutoUpdateConfigSpec) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{1}
}

func (x *AutoUpdateConfigSpec) GetTools() *AutoUpdateConfigSpecTools {
	if x != nil {
		return x.Tools
	}
	return nil
}

// AutoUpdateConfigSpecTools encodes the parameters for client tools auto updates.
type AutoUpdateConfigSpecTools struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mode defines state of the client tools auto update.
	Mode string `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *AutoUpdateConfigSpecTools) Reset() {
	*x = AutoUpdateConfigSpecTools{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateConfigSpecTools) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateConfigSpecTools) ProtoMessage() {}

func (x *AutoUpdateConfigSpecTools) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateConfigSpecTools.ProtoReflect.Descriptor instead.
func (*AutoUpdateConfigSpecTools) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{2}
}

func (x *AutoUpdateConfigSpecTools) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

// AutoUpdateVersion is a resource singleton with version required for
// tools autoupdate.
type AutoUpdateVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	SubKind  string                 `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	Version  string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Metadata *v1.Metadata           `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *AutoUpdateVersionSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *AutoUpdateVersion) Reset() {
	*x = AutoUpdateVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateVersion) ProtoMessage() {}

func (x *AutoUpdateVersion) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateVersion.ProtoReflect.Descriptor instead.
func (*AutoUpdateVersion) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{3}
}

func (x *AutoUpdateVersion) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AutoUpdateVersion) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *AutoUpdateVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AutoUpdateVersion) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AutoUpdateVersion) GetSpec() *AutoUpdateVersionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// AutoUpdateVersionSpec encodes the parameters of the autoupdate versions.
type AutoUpdateVersionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tools *AutoUpdateVersionSpecTools `protobuf:"bytes,2,opt,name=tools,proto3" json:"tools,omitempty"`
}

func (x *AutoUpdateVersionSpec) Reset() {
	*x = AutoUpdateVersionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateVersionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateVersionSpec) ProtoMessage() {}

func (x *AutoUpdateVersionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateVersionSpec.ProtoReflect.Descriptor instead.
func (*AutoUpdateVersionSpec) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{4}
}

func (x *AutoUpdateVersionSpec) GetTools() *AutoUpdateVersionSpecTools {
	if x != nil {
		return x.Tools
	}
	return nil
}

// AutoUpdateVersionSpecTools encodes the parameters for client tools auto updates.
type AutoUpdateVersionSpecTools struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TargetVersion specifies the semantic version required for tools to establish a connection with the cluster.
	// Client tools after connection to the cluster going to be updated to this version automatically.
	TargetVersion string `protobuf:"bytes,1,opt,name=target_version,json=targetVersion,proto3" json:"target_version,omitempty"`
}

func (x *AutoUpdateVersionSpecTools) Reset() {
	*x = AutoUpdateVersionSpecTools{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateVersionSpecTools) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateVersionSpecTools) ProtoMessage() {}

func (x *AutoUpdateVersionSpecTools) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateVersionSpecTools.ProtoReflect.Descriptor instead.
func (*AutoUpdateVersionSpecTools) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP(), []int{5}
}

func (x *AutoUpdateVersionSpecTools) GetTargetVersion() string {
	if x != nil {
		return x.TargetVersion
	}
	return ""
}

var File_teleport_autoupdate_v1_autoupdate_proto protoreflect.FileDescriptor

var file_teleport_autoupdate_v1_autoupdate_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x77,
	0x0a, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x47, 0x0a, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x70, 0x65, 0x63, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x10, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x61, 0x75, 0x74,
	0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2f, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x54,
	0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74,
	0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x41, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x76, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x48, 0x0a,
	0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x54, 0x6f, 0x6f, 0x6c, 0x73,
	0x52, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x0d, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x1a,
	0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_teleport_autoupdate_v1_autoupdate_proto_rawDescOnce sync.Once
	file_teleport_autoupdate_v1_autoupdate_proto_rawDescData = file_teleport_autoupdate_v1_autoupdate_proto_rawDesc
)

func file_teleport_autoupdate_v1_autoupdate_proto_rawDescGZIP() []byte {
	file_teleport_autoupdate_v1_autoupdate_proto_rawDescOnce.Do(func() {
		file_teleport_autoupdate_v1_autoupdate_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_autoupdate_v1_autoupdate_proto_rawDescData)
	})
	return file_teleport_autoupdate_v1_autoupdate_proto_rawDescData
}

var file_teleport_autoupdate_v1_autoupdate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_teleport_autoupdate_v1_autoupdate_proto_goTypes = []any{
	(*AutoUpdateConfig)(nil),           // 0: teleport.autoupdate.v1.AutoUpdateConfig
	(*AutoUpdateConfigSpec)(nil),       // 1: teleport.autoupdate.v1.AutoUpdateConfigSpec
	(*AutoUpdateConfigSpecTools)(nil),  // 2: teleport.autoupdate.v1.AutoUpdateConfigSpecTools
	(*AutoUpdateVersion)(nil),          // 3: teleport.autoupdate.v1.AutoUpdateVersion
	(*AutoUpdateVersionSpec)(nil),      // 4: teleport.autoupdate.v1.AutoUpdateVersionSpec
	(*AutoUpdateVersionSpecTools)(nil), // 5: teleport.autoupdate.v1.AutoUpdateVersionSpecTools
	(*v1.Metadata)(nil),                // 6: teleport.header.v1.Metadata
}
var file_teleport_autoupdate_v1_autoupdate_proto_depIdxs = []int32{
	6, // 0: teleport.autoupdate.v1.AutoUpdateConfig.metadata:type_name -> teleport.header.v1.Metadata
	1, // 1: teleport.autoupdate.v1.AutoUpdateConfig.spec:type_name -> teleport.autoupdate.v1.AutoUpdateConfigSpec
	2, // 2: teleport.autoupdate.v1.AutoUpdateConfigSpec.tools:type_name -> teleport.autoupdate.v1.AutoUpdateConfigSpecTools
	6, // 3: teleport.autoupdate.v1.AutoUpdateVersion.metadata:type_name -> teleport.header.v1.Metadata
	4, // 4: teleport.autoupdate.v1.AutoUpdateVersion.spec:type_name -> teleport.autoupdate.v1.AutoUpdateVersionSpec
	5, // 5: teleport.autoupdate.v1.AutoUpdateVersionSpec.tools:type_name -> teleport.autoupdate.v1.AutoUpdateVersionSpecTools
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_teleport_autoupdate_v1_autoupdate_proto_init() }
func file_teleport_autoupdate_v1_autoupdate_proto_init() {
	if File_teleport_autoupdate_v1_autoupdate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateConfig); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateConfigSpec); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateConfigSpecTools); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateVersion); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateVersionSpec); i {
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
		file_teleport_autoupdate_v1_autoupdate_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateVersionSpecTools); i {
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
			RawDescriptor: file_teleport_autoupdate_v1_autoupdate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_autoupdate_v1_autoupdate_proto_goTypes,
		DependencyIndexes: file_teleport_autoupdate_v1_autoupdate_proto_depIdxs,
		MessageInfos:      file_teleport_autoupdate_v1_autoupdate_proto_msgTypes,
	}.Build()
	File_teleport_autoupdate_v1_autoupdate_proto = out.File
	file_teleport_autoupdate_v1_autoupdate_proto_rawDesc = nil
	file_teleport_autoupdate_v1_autoupdate_proto_goTypes = nil
	file_teleport_autoupdate_v1_autoupdate_proto_depIdxs = nil
}
