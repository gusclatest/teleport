//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: teleport/lib/teleterm/v1/access_request.proto

package teletermv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the request state of Access Request. option of PENDING, APPROVED, DENIED, PROMOTED, NONE
	State         string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	ResolveReason string `protobuf:"bytes,3,opt,name=resolve_reason,json=resolveReason,proto3" json:"resolve_reason,omitempty"`
	RequestReason string `protobuf:"bytes,4,opt,name=request_reason,json=requestReason,proto3" json:"request_reason,omitempty"`
	// user is the user who submitted the Access Request
	User string `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	// a list of roles requested
	Roles              []string               `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
	Created            *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created,proto3" json:"created,omitempty"`
	Expires            *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expires,proto3" json:"expires,omitempty"`
	Reviews            []*AccessRequestReview `protobuf:"bytes,9,rep,name=reviews,proto3" json:"reviews,omitempty"`
	SuggestedReviewers []string               `protobuf:"bytes,10,rep,name=suggested_reviewers,json=suggestedReviewers,proto3" json:"suggested_reviewers,omitempty"`
	// thresholds specifies minimum amount of approvers or deniers. Defaults to 'default'
	ThresholdNames []string `protobuf:"bytes,11,rep,name=threshold_names,json=thresholdNames,proto3" json:"threshold_names,omitempty"`
	// TODO(avatus) remove the resource_ids field once the changes to rely on resources instead is merged
	// a list of resourceIDs requested in the AccessRequest
	ResourceIds []*ResourceID `protobuf:"bytes,12,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
	Resources   []*Resource   `protobuf:"bytes,13,rep,name=resources,proto3" json:"resources,omitempty"`
	// promoted_access_list_title is the title of the access
	// list that this access request was promoted to.
	PromotedAccessListTitle string `protobuf:"bytes,14,opt,name=promoted_access_list_title,json=promotedAccessListTitle,proto3" json:"promoted_access_list_title,omitempty"`
	// assume_start_time is the time after which the requested access can be assumed.
	AssumeStartTime *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=assume_start_time,json=assumeStartTime,proto3" json:"assume_start_time,omitempty"`
	// max_duration is the maximum duration for which the request is valid.
	MaxDuration *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=max_duration,json=maxDuration,proto3" json:"max_duration,omitempty"`
	// request_ttl is the expiration time of the request (how long it will await
	// approval).
	RequestTtl *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=request_ttl,json=requestTtl,proto3" json:"request_ttl,omitempty"`
	// session_ttl indicates how long a certificate for a session should be valid for.
	SessionTtl *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=session_ttl,json=sessionTtl,proto3" json:"session_ttl,omitempty"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_access_request_proto_rawDescGZIP(), []int{0}
}

func (x *AccessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccessRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AccessRequest) GetResolveReason() string {
	if x != nil {
		return x.ResolveReason
	}
	return ""
}

func (x *AccessRequest) GetRequestReason() string {
	if x != nil {
		return x.RequestReason
	}
	return ""
}

func (x *AccessRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AccessRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AccessRequest) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *AccessRequest) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

func (x *AccessRequest) GetReviews() []*AccessRequestReview {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *AccessRequest) GetSuggestedReviewers() []string {
	if x != nil {
		return x.SuggestedReviewers
	}
	return nil
}

func (x *AccessRequest) GetThresholdNames() []string {
	if x != nil {
		return x.ThresholdNames
	}
	return nil
}

func (x *AccessRequest) GetResourceIds() []*ResourceID {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *AccessRequest) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *AccessRequest) GetPromotedAccessListTitle() string {
	if x != nil {
		return x.PromotedAccessListTitle
	}
	return ""
}

func (x *AccessRequest) GetAssumeStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AssumeStartTime
	}
	return nil
}

func (x *AccessRequest) GetMaxDuration() *timestamppb.Timestamp {
	if x != nil {
		return x.MaxDuration
	}
	return nil
}

func (x *AccessRequest) GetRequestTtl() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestTtl
	}
	return nil
}

func (x *AccessRequest) GetSessionTtl() *timestamppb.Timestamp {
	if x != nil {
		return x.SessionTtl
	}
	return nil
}

type AccessRequestReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// author is the creator of the AccessRequestReview.
	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	// list of roles approved
	Roles []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	// the state of the review, either APPROVED or DENIED
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// reason is why the request was approved or denied
	Reason  string                 `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Created *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	// promoted_access_list_title is the title of the access
	// list that the access request was promoted to.
	PromotedAccessListTitle string `protobuf:"bytes,6,opt,name=promoted_access_list_title,json=promotedAccessListTitle,proto3" json:"promoted_access_list_title,omitempty"`
	// if not a nil value, this reviewer overwrote
	// the requested start time.
	AssumeStartTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=assume_start_time,json=assumeStartTime,proto3" json:"assume_start_time,omitempty"`
}

func (x *AccessRequestReview) Reset() {
	*x = AccessRequestReview{}
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessRequestReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequestReview) ProtoMessage() {}

func (x *AccessRequestReview) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequestReview.ProtoReflect.Descriptor instead.
func (*AccessRequestReview) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_access_request_proto_rawDescGZIP(), []int{1}
}

func (x *AccessRequestReview) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *AccessRequestReview) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AccessRequestReview) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AccessRequestReview) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *AccessRequestReview) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *AccessRequestReview) GetPromotedAccessListTitle() string {
	if x != nil {
		return x.PromotedAccessListTitle
	}
	return ""
}

func (x *AccessRequestReview) GetAssumeStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AssumeStartTime
	}
	return nil
}

type ResourceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind            string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ClusterName     string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	SubResourceName string `protobuf:"bytes,4,opt,name=sub_resource_name,json=subResourceName,proto3" json:"sub_resource_name,omitempty"`
}

func (x *ResourceID) Reset() {
	*x = ResourceID{}
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceID) ProtoMessage() {}

func (x *ResourceID) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceID.ProtoReflect.Descriptor instead.
func (*ResourceID) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_access_request_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceID) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ResourceID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceID) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ResourceID) GetSubResourceName() string {
	if x != nil {
		return x.SubResourceName
	}
	return ""
}

type ResourceDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname     string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	FriendlyName string `protobuf:"bytes,2,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
}

func (x *ResourceDetails) Reset() {
	*x = ResourceDetails{}
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDetails) ProtoMessage() {}

func (x *ResourceDetails) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDetails.ProtoReflect.Descriptor instead.
func (*ResourceDetails) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_access_request_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceDetails) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ResourceDetails) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *ResourceID      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Details *ResourceDetails `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_access_request_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_access_request_proto_rawDescGZIP(), []int{4}
}

func (x *Resource) GetId() *ResourceID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Resource) GetDetails() *ResourceDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_teleport_lib_teleterm_v1_access_request_proto protoreflect.FileDescriptor

var file_teleport_lib_teleterm_v1_access_request_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x07, 0x0a, 0x0d, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x34, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12,
	0x2f, 0x0a, 0x13, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x65, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x46, 0x0a, 0x11, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6d, 0x61, 0x78,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6d, 0x61, 0x78,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x74, 0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x74, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x74, 0x6c, 0x22, 0xac, 0x02, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x1a,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x61, 0x73, 0x73,
	0x75, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0f, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x83, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x75, 0x62, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c,
	0x69, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_teleport_lib_teleterm_v1_access_request_proto_rawDescOnce sync.Once
	file_teleport_lib_teleterm_v1_access_request_proto_rawDescData = file_teleport_lib_teleterm_v1_access_request_proto_rawDesc
)

func file_teleport_lib_teleterm_v1_access_request_proto_rawDescGZIP() []byte {
	file_teleport_lib_teleterm_v1_access_request_proto_rawDescOnce.Do(func() {
		file_teleport_lib_teleterm_v1_access_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_lib_teleterm_v1_access_request_proto_rawDescData)
	})
	return file_teleport_lib_teleterm_v1_access_request_proto_rawDescData
}

var file_teleport_lib_teleterm_v1_access_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teleport_lib_teleterm_v1_access_request_proto_goTypes = []any{
	(*AccessRequest)(nil),         // 0: teleport.lib.teleterm.v1.AccessRequest
	(*AccessRequestReview)(nil),   // 1: teleport.lib.teleterm.v1.AccessRequestReview
	(*ResourceID)(nil),            // 2: teleport.lib.teleterm.v1.ResourceID
	(*ResourceDetails)(nil),       // 3: teleport.lib.teleterm.v1.ResourceDetails
	(*Resource)(nil),              // 4: teleport.lib.teleterm.v1.Resource
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_teleport_lib_teleterm_v1_access_request_proto_depIdxs = []int32{
	5,  // 0: teleport.lib.teleterm.v1.AccessRequest.created:type_name -> google.protobuf.Timestamp
	5,  // 1: teleport.lib.teleterm.v1.AccessRequest.expires:type_name -> google.protobuf.Timestamp
	1,  // 2: teleport.lib.teleterm.v1.AccessRequest.reviews:type_name -> teleport.lib.teleterm.v1.AccessRequestReview
	2,  // 3: teleport.lib.teleterm.v1.AccessRequest.resource_ids:type_name -> teleport.lib.teleterm.v1.ResourceID
	4,  // 4: teleport.lib.teleterm.v1.AccessRequest.resources:type_name -> teleport.lib.teleterm.v1.Resource
	5,  // 5: teleport.lib.teleterm.v1.AccessRequest.assume_start_time:type_name -> google.protobuf.Timestamp
	5,  // 6: teleport.lib.teleterm.v1.AccessRequest.max_duration:type_name -> google.protobuf.Timestamp
	5,  // 7: teleport.lib.teleterm.v1.AccessRequest.request_ttl:type_name -> google.protobuf.Timestamp
	5,  // 8: teleport.lib.teleterm.v1.AccessRequest.session_ttl:type_name -> google.protobuf.Timestamp
	5,  // 9: teleport.lib.teleterm.v1.AccessRequestReview.created:type_name -> google.protobuf.Timestamp
	5,  // 10: teleport.lib.teleterm.v1.AccessRequestReview.assume_start_time:type_name -> google.protobuf.Timestamp
	2,  // 11: teleport.lib.teleterm.v1.Resource.id:type_name -> teleport.lib.teleterm.v1.ResourceID
	3,  // 12: teleport.lib.teleterm.v1.Resource.details:type_name -> teleport.lib.teleterm.v1.ResourceDetails
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_teleport_lib_teleterm_v1_access_request_proto_init() }
func file_teleport_lib_teleterm_v1_access_request_proto_init() {
	if File_teleport_lib_teleterm_v1_access_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_lib_teleterm_v1_access_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_lib_teleterm_v1_access_request_proto_goTypes,
		DependencyIndexes: file_teleport_lib_teleterm_v1_access_request_proto_depIdxs,
		MessageInfos:      file_teleport_lib_teleterm_v1_access_request_proto_msgTypes,
	}.Build()
	File_teleport_lib_teleterm_v1_access_request_proto = out.File
	file_teleport_lib_teleterm_v1_access_request_proto_rawDesc = nil
	file_teleport_lib_teleterm_v1_access_request_proto_goTypes = nil
	file_teleport_lib_teleterm_v1_access_request_proto_depIdxs = nil
}
