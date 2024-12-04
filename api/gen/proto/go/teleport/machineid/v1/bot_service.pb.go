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
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: teleport/machineid/v1/bot_service.proto

package machineidv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request for CreateBot.
type CreateBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bot to create.
	Bot *Bot `protobuf:"bytes,1,opt,name=bot,proto3" json:"bot,omitempty"`
}

func (x *CreateBotRequest) Reset() {
	*x = CreateBotRequest{}
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBotRequest) ProtoMessage() {}

func (x *CreateBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBotRequest.ProtoReflect.Descriptor instead.
func (*CreateBotRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBotRequest) GetBot() *Bot {
	if x != nil {
		return x.Bot
	}
	return nil
}

// The request for GetBot.
type GetBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the bot to fetch.
	BotName string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
}

func (x *GetBotRequest) Reset() {
	*x = GetBotRequest{}
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBotRequest) ProtoMessage() {}

func (x *GetBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBotRequest.ProtoReflect.Descriptor instead.
func (*GetBotRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetBotRequest) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

// The request for ListBots.
type ListBotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBotsRequest) Reset() {
	*x = ListBotsRequest{}
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBotsRequest) ProtoMessage() {}

func (x *ListBotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBotsRequest.ProtoReflect.Descriptor instead.
func (*ListBotsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBotsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBotsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response for ListBots.
type ListBotsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The page of Bots that matched the request.
	Bots []*Bot `protobuf:"bytes,1,rep,name=bots,proto3" json:"bots,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBotsResponse) Reset() {
	*x = ListBotsResponse{}
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBotsResponse) ProtoMessage() {}

func (x *ListBotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBotsResponse.ProtoReflect.Descriptor instead.
func (*ListBotsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListBotsResponse) GetBots() []*Bot {
	if x != nil {
		return x.Bots
	}
	return nil
}

func (x *ListBotsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request for UpdateBot.
type UpdateBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The values to apply based on the update mask. The name must be specified.
	Bot *Bot `protobuf:"bytes,1,opt,name=bot,proto3" json:"bot,omitempty"`
	// The update mask applied to a Bot.
	// Fields are masked according to their proto name.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateBotRequest) Reset() {
	*x = UpdateBotRequest{}
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBotRequest) ProtoMessage() {}

func (x *UpdateBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBotRequest.ProtoReflect.Descriptor instead.
func (*UpdateBotRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBotRequest) GetBot() *Bot {
	if x != nil {
		return x.Bot
	}
	return nil
}

func (x *UpdateBotRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// The request for UpsertBot.
type UpsertBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bot to create or replace.
	Bot *Bot `protobuf:"bytes,1,opt,name=bot,proto3" json:"bot,omitempty"`
}

func (x *UpsertBotRequest) Reset() {
	*x = UpsertBotRequest{}
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertBotRequest) ProtoMessage() {}

func (x *UpsertBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertBotRequest.ProtoReflect.Descriptor instead.
func (*UpsertBotRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertBotRequest) GetBot() *Bot {
	if x != nil {
		return x.Bot
	}
	return nil
}

// The request for DeleteBot.
type DeleteBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the bot to delete.
	BotName string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
}

func (x *DeleteBotRequest) Reset() {
	*x = DeleteBotRequest{}
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBotRequest) ProtoMessage() {}

func (x *DeleteBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_machineid_v1_bot_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBotRequest.ProtoReflect.Descriptor instead.
func (*DeleteBotRequest) Descriptor() ([]byte, []int) {
	return file_teleport_machineid_v1_bot_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBotRequest) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

var File_teleport_machineid_v1_bot_service_proto protoreflect.FileDescriptor

var file_teleport_machineid_v1_bot_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x62, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x52, 0x03, 0x62,
	0x6f, 0x74, 0x22, 0x2a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6a, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x52, 0x04, 0x62, 0x6f, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x03, 0x62, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x52, 0x03, 0x62, 0x6f, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x40, 0x0a, 0x10, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x03,
	0x62, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x74, 0x52, 0x03, 0x62, 0x6f, 0x74, 0x22, 0x2d, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xf9, 0x03, 0x0a, 0x0a, 0x42, 0x6f,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x74, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6f, 0x74, 0x12, 0x5b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x74, 0x73,
	0x12, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x12, 0x27,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x74,
	0x12, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x42,
	0x6f, 0x74, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x74, 0x12, 0x4c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x74, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_machineid_v1_bot_service_proto_rawDescOnce sync.Once
	file_teleport_machineid_v1_bot_service_proto_rawDescData = file_teleport_machineid_v1_bot_service_proto_rawDesc
)

func file_teleport_machineid_v1_bot_service_proto_rawDescGZIP() []byte {
	file_teleport_machineid_v1_bot_service_proto_rawDescOnce.Do(func() {
		file_teleport_machineid_v1_bot_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_machineid_v1_bot_service_proto_rawDescData)
	})
	return file_teleport_machineid_v1_bot_service_proto_rawDescData
}

var file_teleport_machineid_v1_bot_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_teleport_machineid_v1_bot_service_proto_goTypes = []any{
	(*CreateBotRequest)(nil),      // 0: teleport.machineid.v1.CreateBotRequest
	(*GetBotRequest)(nil),         // 1: teleport.machineid.v1.GetBotRequest
	(*ListBotsRequest)(nil),       // 2: teleport.machineid.v1.ListBotsRequest
	(*ListBotsResponse)(nil),      // 3: teleport.machineid.v1.ListBotsResponse
	(*UpdateBotRequest)(nil),      // 4: teleport.machineid.v1.UpdateBotRequest
	(*UpsertBotRequest)(nil),      // 5: teleport.machineid.v1.UpsertBotRequest
	(*DeleteBotRequest)(nil),      // 6: teleport.machineid.v1.DeleteBotRequest
	(*Bot)(nil),                   // 7: teleport.machineid.v1.Bot
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_teleport_machineid_v1_bot_service_proto_depIdxs = []int32{
	7,  // 0: teleport.machineid.v1.CreateBotRequest.bot:type_name -> teleport.machineid.v1.Bot
	7,  // 1: teleport.machineid.v1.ListBotsResponse.bots:type_name -> teleport.machineid.v1.Bot
	7,  // 2: teleport.machineid.v1.UpdateBotRequest.bot:type_name -> teleport.machineid.v1.Bot
	8,  // 3: teleport.machineid.v1.UpdateBotRequest.update_mask:type_name -> google.protobuf.FieldMask
	7,  // 4: teleport.machineid.v1.UpsertBotRequest.bot:type_name -> teleport.machineid.v1.Bot
	1,  // 5: teleport.machineid.v1.BotService.GetBot:input_type -> teleport.machineid.v1.GetBotRequest
	2,  // 6: teleport.machineid.v1.BotService.ListBots:input_type -> teleport.machineid.v1.ListBotsRequest
	0,  // 7: teleport.machineid.v1.BotService.CreateBot:input_type -> teleport.machineid.v1.CreateBotRequest
	4,  // 8: teleport.machineid.v1.BotService.UpdateBot:input_type -> teleport.machineid.v1.UpdateBotRequest
	5,  // 9: teleport.machineid.v1.BotService.UpsertBot:input_type -> teleport.machineid.v1.UpsertBotRequest
	6,  // 10: teleport.machineid.v1.BotService.DeleteBot:input_type -> teleport.machineid.v1.DeleteBotRequest
	7,  // 11: teleport.machineid.v1.BotService.GetBot:output_type -> teleport.machineid.v1.Bot
	3,  // 12: teleport.machineid.v1.BotService.ListBots:output_type -> teleport.machineid.v1.ListBotsResponse
	7,  // 13: teleport.machineid.v1.BotService.CreateBot:output_type -> teleport.machineid.v1.Bot
	7,  // 14: teleport.machineid.v1.BotService.UpdateBot:output_type -> teleport.machineid.v1.Bot
	7,  // 15: teleport.machineid.v1.BotService.UpsertBot:output_type -> teleport.machineid.v1.Bot
	9,  // 16: teleport.machineid.v1.BotService.DeleteBot:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_teleport_machineid_v1_bot_service_proto_init() }
func file_teleport_machineid_v1_bot_service_proto_init() {
	if File_teleport_machineid_v1_bot_service_proto != nil {
		return
	}
	file_teleport_machineid_v1_bot_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_machineid_v1_bot_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_machineid_v1_bot_service_proto_goTypes,
		DependencyIndexes: file_teleport_machineid_v1_bot_service_proto_depIdxs,
		MessageInfos:      file_teleport_machineid_v1_bot_service_proto_msgTypes,
	}.Build()
	File_teleport_machineid_v1_bot_service_proto = out.File
	file_teleport_machineid_v1_bot_service_proto_rawDesc = nil
	file_teleport_machineid_v1_bot_service_proto_goTypes = nil
	file_teleport_machineid_v1_bot_service_proto_depIdxs = nil
}
