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
// source: teleport/devicetrust/v1/usage.proto

package devicetrustv1

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

// AccountUsageType is the type of the underlying account, either limitless or
// limit-based.
type AccountUsageType int32

const (
	AccountUsageType_ACCOUNT_USAGE_TYPE_UNSPECIFIED AccountUsageType = 0
	AccountUsageType_ACCOUNT_USAGE_TYPE_UNLIMITED   AccountUsageType = 1
	AccountUsageType_ACCOUNT_USAGE_TYPE_USAGE_BASED AccountUsageType = 2
)

// Enum value maps for AccountUsageType.
var (
	AccountUsageType_name = map[int32]string{
		0: "ACCOUNT_USAGE_TYPE_UNSPECIFIED",
		1: "ACCOUNT_USAGE_TYPE_UNLIMITED",
		2: "ACCOUNT_USAGE_TYPE_USAGE_BASED",
	}
	AccountUsageType_value = map[string]int32{
		"ACCOUNT_USAGE_TYPE_UNSPECIFIED": 0,
		"ACCOUNT_USAGE_TYPE_UNLIMITED":   1,
		"ACCOUNT_USAGE_TYPE_USAGE_BASED": 2,
	}
)

func (x AccountUsageType) Enum() *AccountUsageType {
	p := new(AccountUsageType)
	*p = x
	return p
}

func (x AccountUsageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountUsageType) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_devicetrust_v1_usage_proto_enumTypes[0].Descriptor()
}

func (AccountUsageType) Type() protoreflect.EnumType {
	return &file_teleport_devicetrust_v1_usage_proto_enumTypes[0]
}

func (x AccountUsageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountUsageType.Descriptor instead.
func (AccountUsageType) EnumDescriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_usage_proto_rawDescGZIP(), []int{0}
}

// DevicesUsage holds aggregated information about trusted device usage.
type DevicesUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Usage type of the underlying account.
	// UNLIMITED accounts have no limits on the number of trusted devices,
	// therefore all usage-based limits are data are zeroed when read.
	AccountUsageType AccountUsageType `protobuf:"varint,1,opt,name=account_usage_type,json=accountUsageType,proto3,enum=teleport.devicetrust.v1.AccountUsageType" json:"account_usage_type,omitempty"`
	// Devices usage limit.
	// Always zero if the usage type is UNLIMITED.
	DevicesUsageLimit int32 `protobuf:"varint,2,opt,name=devices_usage_limit,json=devicesUsageLimit,proto3" json:"devices_usage_limit,omitempty"`
	// Devices in use.
	// May be greater than [devices_usage_limit] in some cases.
	// Always zero if the usage type is UNLIMITED.
	DevicesInUse int32 `protobuf:"varint,3,opt,name=devices_in_use,json=devicesInUse,proto3" json:"devices_in_use,omitempty"`
}

func (x *DevicesUsage) Reset() {
	*x = DevicesUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_devicetrust_v1_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicesUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicesUsage) ProtoMessage() {}

func (x *DevicesUsage) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicesUsage.ProtoReflect.Descriptor instead.
func (*DevicesUsage) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_usage_proto_rawDescGZIP(), []int{0}
}

func (x *DevicesUsage) GetAccountUsageType() AccountUsageType {
	if x != nil {
		return x.AccountUsageType
	}
	return AccountUsageType_ACCOUNT_USAGE_TYPE_UNSPECIFIED
}

func (x *DevicesUsage) GetDevicesUsageLimit() int32 {
	if x != nil {
		return x.DevicesUsageLimit
	}
	return 0
}

func (x *DevicesUsage) GetDevicesInUse() int32 {
	if x != nil {
		return x.DevicesInUse
	}
	return 0
}

var File_teleport_devicetrust_v1_usage_proto protoreflect.FileDescriptor

var file_teleport_devicetrust_v1_usage_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xbd,
	0x01, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x57, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x2a, 0x7c,
	0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x55, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4c,
	0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x44, 0x10, 0x02, 0x42, 0x5a, 0x5a, 0x58,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_devicetrust_v1_usage_proto_rawDescOnce sync.Once
	file_teleport_devicetrust_v1_usage_proto_rawDescData = file_teleport_devicetrust_v1_usage_proto_rawDesc
)

func file_teleport_devicetrust_v1_usage_proto_rawDescGZIP() []byte {
	file_teleport_devicetrust_v1_usage_proto_rawDescOnce.Do(func() {
		file_teleport_devicetrust_v1_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_devicetrust_v1_usage_proto_rawDescData)
	})
	return file_teleport_devicetrust_v1_usage_proto_rawDescData
}

var file_teleport_devicetrust_v1_usage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teleport_devicetrust_v1_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teleport_devicetrust_v1_usage_proto_goTypes = []interface{}{
	(AccountUsageType)(0), // 0: teleport.devicetrust.v1.AccountUsageType
	(*DevicesUsage)(nil),  // 1: teleport.devicetrust.v1.DevicesUsage
}
var file_teleport_devicetrust_v1_usage_proto_depIdxs = []int32{
	0, // 0: teleport.devicetrust.v1.DevicesUsage.account_usage_type:type_name -> teleport.devicetrust.v1.AccountUsageType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_devicetrust_v1_usage_proto_init() }
func file_teleport_devicetrust_v1_usage_proto_init() {
	if File_teleport_devicetrust_v1_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_devicetrust_v1_usage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicesUsage); i {
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
			RawDescriptor: file_teleport_devicetrust_v1_usage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_devicetrust_v1_usage_proto_goTypes,
		DependencyIndexes: file_teleport_devicetrust_v1_usage_proto_depIdxs,
		EnumInfos:         file_teleport_devicetrust_v1_usage_proto_enumTypes,
		MessageInfos:      file_teleport_devicetrust_v1_usage_proto_msgTypes,
	}.Build()
	File_teleport_devicetrust_v1_usage_proto = out.File
	file_teleport_devicetrust_v1_usage_proto_rawDesc = nil
	file_teleport_devicetrust_v1_usage_proto_goTypes = nil
	file_teleport_devicetrust_v1_usage_proto_depIdxs = nil
}
