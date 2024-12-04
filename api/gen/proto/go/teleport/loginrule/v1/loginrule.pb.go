// Copyright 2022 Gravitational, Inc
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
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: teleport/loginrule/v1/loginrule.proto

package loginrulev1

import (
	types "github.com/gravitational/teleport/api/types"
	wrappers "github.com/gravitational/teleport/api/types/wrappers"
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

// LoginRule is a resource to configure rules and logic which should run during
// Teleport user login.
type LoginRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata is resource metadata.
	Metadata *types.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Version is the resource version.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Priority is the priority of the login rule relative to other login rules
	// in the same cluster. Login rules with a lower numbered priority will be
	// evaluated first.
	Priority int32 `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	// TraitsMap is a map of trait keys to lists of predicate expressions which
	// should evaluate to the desired values for that trait.
	TraitsMap map[string]*wrappers.StringValues `protobuf:"bytes,4,rep,name=traits_map,json=traitsMap,proto3" json:"traits_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// TraitsExpression is a predicate expression which should return the
	// desired traits for the user upon login.
	TraitsExpression string `protobuf:"bytes,5,opt,name=traits_expression,json=traitsExpression,proto3" json:"traits_expression,omitempty"`
}

func (x *LoginRule) Reset() {
	*x = LoginRule{}
	mi := &file_teleport_loginrule_v1_loginrule_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRule) ProtoMessage() {}

func (x *LoginRule) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_loginrule_v1_loginrule_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRule.ProtoReflect.Descriptor instead.
func (*LoginRule) Descriptor() ([]byte, []int) {
	return file_teleport_loginrule_v1_loginrule_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRule) GetMetadata() *types.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *LoginRule) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *LoginRule) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *LoginRule) GetTraitsMap() map[string]*wrappers.StringValues {
	if x != nil {
		return x.TraitsMap
	}
	return nil
}

func (x *LoginRule) GetTraitsExpression() string {
	if x != nil {
		return x.TraitsExpression
	}
	return ""
}

var File_teleport_loginrule_v1_loginrule_proto protoreflect.FileDescriptor

var file_teleport_loginrule_v1_loginrule_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x21,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc1, 0x02, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2b,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x4e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x4d, 0x61,
	0x70, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x54,
	0x0a, 0x0e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x75, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_loginrule_v1_loginrule_proto_rawDescOnce sync.Once
	file_teleport_loginrule_v1_loginrule_proto_rawDescData = file_teleport_loginrule_v1_loginrule_proto_rawDesc
)

func file_teleport_loginrule_v1_loginrule_proto_rawDescGZIP() []byte {
	file_teleport_loginrule_v1_loginrule_proto_rawDescOnce.Do(func() {
		file_teleport_loginrule_v1_loginrule_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_loginrule_v1_loginrule_proto_rawDescData)
	})
	return file_teleport_loginrule_v1_loginrule_proto_rawDescData
}

var file_teleport_loginrule_v1_loginrule_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_loginrule_v1_loginrule_proto_goTypes = []any{
	(*LoginRule)(nil),             // 0: teleport.loginrule.v1.LoginRule
	nil,                           // 1: teleport.loginrule.v1.LoginRule.TraitsMapEntry
	(*types.Metadata)(nil),        // 2: types.Metadata
	(*wrappers.StringValues)(nil), // 3: wrappers.StringValues
}
var file_teleport_loginrule_v1_loginrule_proto_depIdxs = []int32{
	2, // 0: teleport.loginrule.v1.LoginRule.metadata:type_name -> types.Metadata
	1, // 1: teleport.loginrule.v1.LoginRule.traits_map:type_name -> teleport.loginrule.v1.LoginRule.TraitsMapEntry
	3, // 2: teleport.loginrule.v1.LoginRule.TraitsMapEntry.value:type_name -> wrappers.StringValues
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_teleport_loginrule_v1_loginrule_proto_init() }
func file_teleport_loginrule_v1_loginrule_proto_init() {
	if File_teleport_loginrule_v1_loginrule_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_loginrule_v1_loginrule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_loginrule_v1_loginrule_proto_goTypes,
		DependencyIndexes: file_teleport_loginrule_v1_loginrule_proto_depIdxs,
		MessageInfos:      file_teleport_loginrule_v1_loginrule_proto_msgTypes,
	}.Build()
	File_teleport_loginrule_v1_loginrule_proto = out.File
	file_teleport_loginrule_v1_loginrule_proto_rawDesc = nil
	file_teleport_loginrule_v1_loginrule_proto_goTypes = nil
	file_teleport_loginrule_v1_loginrule_proto_depIdxs = nil
}
