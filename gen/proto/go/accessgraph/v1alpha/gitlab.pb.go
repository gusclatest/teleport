//
// Teleport
// Copyright (C) 2024  Gravitational, Inc.
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
// source: accessgraph/v1alpha/gitlab.proto

package accessgraphv1alpha

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

// AccessLevelType defines the access level a user has
// to a project or Gitlab group.
type AccessLevelType int32

const (
	// ACCESS_LEVEL_TYPE_UNSPECIFIED is an unspecified permissions.
	AccessLevelType_ACCESS_LEVEL_TYPE_UNSPECIFIED AccessLevelType = 0
	// ACCESS_LEVEL_TYPE_NO_PERMISSIONS represents no permissions.
	AccessLevelType_ACCESS_LEVEL_TYPE_NO_PERMISSIONS AccessLevelType = 1
	// ACCESS_LEVEL_TYPE_MINIMAL represents "minimal" permissions to a project/group.
	AccessLevelType_ACCESS_LEVEL_TYPE_MINIMAL AccessLevelType = 2
	// ACCESS_LEVEL_TYPE_GUEST represents "guest" permissions to a project/group.
	AccessLevelType_ACCESS_LEVEL_TYPE_GUEST AccessLevelType = 3
	// ACCESS_LEVEL_TYPE_REPORTER represents "reporter" permissions to a project/group.
	AccessLevelType_ACCESS_LEVEL_TYPE_REPORTER AccessLevelType = 4
	// ACCESS_LEVEL_TYPE_DEVELOPER represents "developer" permissions to a project/group.
	AccessLevelType_ACCESS_LEVEL_TYPE_DEVELOPER AccessLevelType = 5
	// ACCESS_LEVEL_TYPE_MAINTAINER represents "master" permissions to a project/group.
	AccessLevelType_ACCESS_LEVEL_TYPE_MAINTAINER AccessLevelType = 6
	// ACCESS_LEVEL_TYPE_OWNER represents "owner" permissions to a project/group.
	AccessLevelType_ACCESS_LEVEL_TYPE_OWNER AccessLevelType = 7
)

// Enum value maps for AccessLevelType.
var (
	AccessLevelType_name = map[int32]string{
		0: "ACCESS_LEVEL_TYPE_UNSPECIFIED",
		1: "ACCESS_LEVEL_TYPE_NO_PERMISSIONS",
		2: "ACCESS_LEVEL_TYPE_MINIMAL",
		3: "ACCESS_LEVEL_TYPE_GUEST",
		4: "ACCESS_LEVEL_TYPE_REPORTER",
		5: "ACCESS_LEVEL_TYPE_DEVELOPER",
		6: "ACCESS_LEVEL_TYPE_MAINTAINER",
		7: "ACCESS_LEVEL_TYPE_OWNER",
	}
	AccessLevelType_value = map[string]int32{
		"ACCESS_LEVEL_TYPE_UNSPECIFIED":    0,
		"ACCESS_LEVEL_TYPE_NO_PERMISSIONS": 1,
		"ACCESS_LEVEL_TYPE_MINIMAL":        2,
		"ACCESS_LEVEL_TYPE_GUEST":          3,
		"ACCESS_LEVEL_TYPE_REPORTER":       4,
		"ACCESS_LEVEL_TYPE_DEVELOPER":      5,
		"ACCESS_LEVEL_TYPE_MAINTAINER":     6,
		"ACCESS_LEVEL_TYPE_OWNER":          7,
	}
)

func (x AccessLevelType) Enum() *AccessLevelType {
	p := new(AccessLevelType)
	*p = x
	return p
}

func (x AccessLevelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessLevelType) Descriptor() protoreflect.EnumDescriptor {
	return file_accessgraph_v1alpha_gitlab_proto_enumTypes[0].Descriptor()
}

func (AccessLevelType) Type() protoreflect.EnumType {
	return &file_accessgraph_v1alpha_gitlab_proto_enumTypes[0]
}

func (x AccessLevelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessLevelType.Descriptor instead.
func (AccessLevelType) EnumDescriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{0}
}

// GitlabSyncOperation is a request to sync Gitlab resources
type GitlabSyncOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GitlabSyncOperation) Reset() {
	*x = GitlabSyncOperation{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabSyncOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabSyncOperation) ProtoMessage() {}

func (x *GitlabSyncOperation) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabSyncOperation.ProtoReflect.Descriptor instead.
func (*GitlabSyncOperation) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{0}
}

// GitlabResourceList is a request that contains resources to be sync.
type GitlabResourceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resources is a list of gitlab resources to sync.
	Resources []*GitlabResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *GitlabResourceList) Reset() {
	*x = GitlabResourceList{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabResourceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabResourceList) ProtoMessage() {}

func (x *GitlabResourceList) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabResourceList.ProtoReflect.Descriptor instead.
func (*GitlabResourceList) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{1}
}

func (x *GitlabResourceList) GetResources() []*GitlabResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

// GitlabResource represents a Gitlab resource
type GitlabResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resource:
	//
	//	*GitlabResource_Group
	//	*GitlabResource_Project
	//	*GitlabResource_ProjectMember
	//	*GitlabResource_GroupMember
	//	*GitlabResource_User
	Resource isGitlabResource_Resource `protobuf_oneof:"resource"`
}

func (x *GitlabResource) Reset() {
	*x = GitlabResource{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabResource) ProtoMessage() {}

func (x *GitlabResource) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabResource.ProtoReflect.Descriptor instead.
func (*GitlabResource) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{2}
}

func (m *GitlabResource) GetResource() isGitlabResource_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (x *GitlabResource) GetGroup() *GitlabGroup {
	if x, ok := x.GetResource().(*GitlabResource_Group); ok {
		return x.Group
	}
	return nil
}

func (x *GitlabResource) GetProject() *GitlabProject {
	if x, ok := x.GetResource().(*GitlabResource_Project); ok {
		return x.Project
	}
	return nil
}

func (x *GitlabResource) GetProjectMember() *GitlabProjectMember {
	if x, ok := x.GetResource().(*GitlabResource_ProjectMember); ok {
		return x.ProjectMember
	}
	return nil
}

func (x *GitlabResource) GetGroupMember() *GitlabGroupMember {
	if x, ok := x.GetResource().(*GitlabResource_GroupMember); ok {
		return x.GroupMember
	}
	return nil
}

func (x *GitlabResource) GetUser() *GitlabUser {
	if x, ok := x.GetResource().(*GitlabResource_User); ok {
		return x.User
	}
	return nil
}

type isGitlabResource_Resource interface {
	isGitlabResource_Resource()
}

type GitlabResource_Group struct {
	// group represents a gitlab group or subgroup in an organization.
	Group *GitlabGroup `protobuf:"bytes,1,opt,name=group,proto3,oneof"`
}

type GitlabResource_Project struct {
	// project represents a gitlab repository.
	Project *GitlabProject `protobuf:"bytes,2,opt,name=project,proto3,oneof"`
}

type GitlabResource_ProjectMember struct {
	// project_member represents a user with certain access levels to a project.
	ProjectMember *GitlabProjectMember `protobuf:"bytes,3,opt,name=project_member,json=projectMember,proto3,oneof"`
}

type GitlabResource_GroupMember struct {
	// group_member represents a user with certain access levels to a group and all subgroups/projects within.
	GroupMember *GitlabGroupMember `protobuf:"bytes,4,opt,name=group_member,json=groupMember,proto3,oneof"`
}

type GitlabResource_User struct {
	// user represents a gitlab user.
	User *GitlabUser `protobuf:"bytes,5,opt,name=user,proto3,oneof"`
}

func (*GitlabResource_Group) isGitlabResource_Resource() {}

func (*GitlabResource_Project) isGitlabResource_Resource() {}

func (*GitlabResource_ProjectMember) isGitlabResource_Resource() {}

func (*GitlabResource_GroupMember) isGitlabResource_Resource() {}

func (*GitlabResource_User) isGitlabResource_Resource() {}

// GitlabGroup represents a Gitlab group
type GitlabGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the group name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// path is the universal identifier for the group location.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// full_name is the group full name.
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// description is the group description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GitlabGroup) Reset() {
	*x = GitlabGroup{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabGroup) ProtoMessage() {}

func (x *GitlabGroup) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabGroup.ProtoReflect.Descriptor instead.
func (*GitlabGroup) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{3}
}

func (x *GitlabGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GitlabGroup) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GitlabGroup) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *GitlabGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// GitlabProject represents a Gitlab project
type GitlabProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the repository name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// path is the universal identifier for the project location.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// description is the project description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GitlabProject) Reset() {
	*x = GitlabProject{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabProject) ProtoMessage() {}

func (x *GitlabProject) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabProject.ProtoReflect.Descriptor instead.
func (*GitlabProject) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{4}
}

func (x *GitlabProject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GitlabProject) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GitlabProject) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// GitlabProjectMember represents a Gitlab project member
type GitlabProjectMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// username is the username of the user.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// access_level defines the permissions the user has to the project.
	AccessLevel AccessLevelType `protobuf:"varint,2,opt,name=access_level,json=accessLevel,proto3,enum=accessgraph.v1alpha.AccessLevelType" json:"access_level,omitempty"`
	// project identifies the project that the user is member of.
	Project *GitlabProject `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GitlabProjectMember) Reset() {
	*x = GitlabProjectMember{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabProjectMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabProjectMember) ProtoMessage() {}

func (x *GitlabProjectMember) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabProjectMember.ProtoReflect.Descriptor instead.
func (*GitlabProjectMember) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{5}
}

func (x *GitlabProjectMember) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GitlabProjectMember) GetAccessLevel() AccessLevelType {
	if x != nil {
		return x.AccessLevel
	}
	return AccessLevelType_ACCESS_LEVEL_TYPE_UNSPECIFIED
}

func (x *GitlabProjectMember) GetProject() *GitlabProject {
	if x != nil {
		return x.Project
	}
	return nil
}

// GitlabGroupMember represents a Gitlab group member
type GitlabGroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// username is the username of the user.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// access_level defines the permissions the user has to the group and all projects within.
	AccessLevel AccessLevelType `protobuf:"varint,2,opt,name=access_level,json=accessLevel,proto3,enum=accessgraph.v1alpha.AccessLevelType" json:"access_level,omitempty"`
	// project identifies the project that the user is member of.
	Group *GitlabGroup `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GitlabGroupMember) Reset() {
	*x = GitlabGroupMember{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabGroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabGroupMember) ProtoMessage() {}

func (x *GitlabGroupMember) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabGroupMember.ProtoReflect.Descriptor instead.
func (*GitlabGroupMember) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{6}
}

func (x *GitlabGroupMember) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GitlabGroupMember) GetAccessLevel() AccessLevelType {
	if x != nil {
		return x.AccessLevel
	}
	return AccessLevelType_ACCESS_LEVEL_TYPE_UNSPECIFIED
}

func (x *GitlabGroupMember) GetGroup() *GitlabGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

// GitlabGroupMember represents a Gitlab user.
type GitlabUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// username is the username of the user.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// email is the user's email.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// name is the user's name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// is_admin indicates if a user is admin.
	IsAdmin bool `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// organization is the user's organization.
	Organization string `protobuf:"bytes,5,opt,name=organization,proto3" json:"organization,omitempty"`
	// last_sign_in_at identifies the last sign in date.
	LastSignInAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_sign_in_at,json=lastSignInAt,proto3" json:"last_sign_in_at,omitempty"`
	// can_create_group identifies if the user can create groups.
	CanCreateGroup bool `protobuf:"varint,7,opt,name=can_create_group,json=canCreateGroup,proto3" json:"can_create_group,omitempty"`
	// can_create_project identifies if the user can create projects.
	CanCreateProject bool `protobuf:"varint,8,opt,name=can_create_project,json=canCreateProject,proto3" json:"can_create_project,omitempty"`
	// two_factor_enabled identifies if the user has two factor authentication enabled.
	TwoFactorEnabled bool `protobuf:"varint,9,opt,name=two_factor_enabled,json=twoFactorEnabled,proto3" json:"two_factor_enabled,omitempty"`
	// identities represents the identity source for the user.
	Identities []*GitlabUserIdentity `protobuf:"bytes,10,rep,name=identities,proto3" json:"identities,omitempty"`
}

func (x *GitlabUser) Reset() {
	*x = GitlabUser{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabUser) ProtoMessage() {}

func (x *GitlabUser) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabUser.ProtoReflect.Descriptor instead.
func (*GitlabUser) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{7}
}

func (x *GitlabUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GitlabUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GitlabUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GitlabUser) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *GitlabUser) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *GitlabUser) GetLastSignInAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSignInAt
	}
	return nil
}

func (x *GitlabUser) GetCanCreateGroup() bool {
	if x != nil {
		return x.CanCreateGroup
	}
	return false
}

func (x *GitlabUser) GetCanCreateProject() bool {
	if x != nil {
		return x.CanCreateProject
	}
	return false
}

func (x *GitlabUser) GetTwoFactorEnabled() bool {
	if x != nil {
		return x.TwoFactorEnabled
	}
	return false
}

func (x *GitlabUser) GetIdentities() []*GitlabUserIdentity {
	if x != nil {
		return x.Identities
	}
	return nil
}

// GitlabUserIdentity identifies the external identity of the user.
type GitlabUserIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// provider identifies the identity provider.
	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	// extern_uid identifies the external uid of the identity.
	ExternUid string `protobuf:"bytes,2,opt,name=extern_uid,json=externUid,proto3" json:"extern_uid,omitempty"`
}

func (x *GitlabUserIdentity) Reset() {
	*x = GitlabUserIdentity{}
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitlabUserIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitlabUserIdentity) ProtoMessage() {}

func (x *GitlabUserIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_gitlab_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitlabUserIdentity.ProtoReflect.Descriptor instead.
func (*GitlabUserIdentity) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP(), []int{8}
}

func (x *GitlabUserIdentity) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *GitlabUserIdentity) GetExternUid() string {
	if x != nil {
		return x.ExternUid
	}
	return ""
}

var File_accessgraph_v1alpha_gitlab_proto protoreflect.FileDescriptor

var file_accessgraph_v1alpha_gitlab_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x57, 0x0a, 0x12, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xed, 0x02, 0x0a, 0x0e, 0x47, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3e, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x51, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x74, 0x0a, 0x0b, 0x47, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x59,
	0x0a, 0x0d, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x47, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x36, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0xa3, 0x03, 0x0a, 0x0a, 0x47, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x41, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x6e, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x74, 0x77, 0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x4f, 0x0a,
	0x12, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x55, 0x69, 0x64, 0x2a, 0x96,
	0x02, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56,
	0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41,
	0x49, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x07, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x67, 0x72, 0x61, 0x70, 0x68, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accessgraph_v1alpha_gitlab_proto_rawDescOnce sync.Once
	file_accessgraph_v1alpha_gitlab_proto_rawDescData = file_accessgraph_v1alpha_gitlab_proto_rawDesc
)

func file_accessgraph_v1alpha_gitlab_proto_rawDescGZIP() []byte {
	file_accessgraph_v1alpha_gitlab_proto_rawDescOnce.Do(func() {
		file_accessgraph_v1alpha_gitlab_proto_rawDescData = protoimpl.X.CompressGZIP(file_accessgraph_v1alpha_gitlab_proto_rawDescData)
	})
	return file_accessgraph_v1alpha_gitlab_proto_rawDescData
}

var file_accessgraph_v1alpha_gitlab_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accessgraph_v1alpha_gitlab_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_accessgraph_v1alpha_gitlab_proto_goTypes = []any{
	(AccessLevelType)(0),          // 0: accessgraph.v1alpha.AccessLevelType
	(*GitlabSyncOperation)(nil),   // 1: accessgraph.v1alpha.GitlabSyncOperation
	(*GitlabResourceList)(nil),    // 2: accessgraph.v1alpha.GitlabResourceList
	(*GitlabResource)(nil),        // 3: accessgraph.v1alpha.GitlabResource
	(*GitlabGroup)(nil),           // 4: accessgraph.v1alpha.GitlabGroup
	(*GitlabProject)(nil),         // 5: accessgraph.v1alpha.GitlabProject
	(*GitlabProjectMember)(nil),   // 6: accessgraph.v1alpha.GitlabProjectMember
	(*GitlabGroupMember)(nil),     // 7: accessgraph.v1alpha.GitlabGroupMember
	(*GitlabUser)(nil),            // 8: accessgraph.v1alpha.GitlabUser
	(*GitlabUserIdentity)(nil),    // 9: accessgraph.v1alpha.GitlabUserIdentity
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_accessgraph_v1alpha_gitlab_proto_depIdxs = []int32{
	3,  // 0: accessgraph.v1alpha.GitlabResourceList.resources:type_name -> accessgraph.v1alpha.GitlabResource
	4,  // 1: accessgraph.v1alpha.GitlabResource.group:type_name -> accessgraph.v1alpha.GitlabGroup
	5,  // 2: accessgraph.v1alpha.GitlabResource.project:type_name -> accessgraph.v1alpha.GitlabProject
	6,  // 3: accessgraph.v1alpha.GitlabResource.project_member:type_name -> accessgraph.v1alpha.GitlabProjectMember
	7,  // 4: accessgraph.v1alpha.GitlabResource.group_member:type_name -> accessgraph.v1alpha.GitlabGroupMember
	8,  // 5: accessgraph.v1alpha.GitlabResource.user:type_name -> accessgraph.v1alpha.GitlabUser
	0,  // 6: accessgraph.v1alpha.GitlabProjectMember.access_level:type_name -> accessgraph.v1alpha.AccessLevelType
	5,  // 7: accessgraph.v1alpha.GitlabProjectMember.project:type_name -> accessgraph.v1alpha.GitlabProject
	0,  // 8: accessgraph.v1alpha.GitlabGroupMember.access_level:type_name -> accessgraph.v1alpha.AccessLevelType
	4,  // 9: accessgraph.v1alpha.GitlabGroupMember.group:type_name -> accessgraph.v1alpha.GitlabGroup
	10, // 10: accessgraph.v1alpha.GitlabUser.last_sign_in_at:type_name -> google.protobuf.Timestamp
	9,  // 11: accessgraph.v1alpha.GitlabUser.identities:type_name -> accessgraph.v1alpha.GitlabUserIdentity
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_accessgraph_v1alpha_gitlab_proto_init() }
func file_accessgraph_v1alpha_gitlab_proto_init() {
	if File_accessgraph_v1alpha_gitlab_proto != nil {
		return
	}
	file_accessgraph_v1alpha_gitlab_proto_msgTypes[2].OneofWrappers = []any{
		(*GitlabResource_Group)(nil),
		(*GitlabResource_Project)(nil),
		(*GitlabResource_ProjectMember)(nil),
		(*GitlabResource_GroupMember)(nil),
		(*GitlabResource_User)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accessgraph_v1alpha_gitlab_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accessgraph_v1alpha_gitlab_proto_goTypes,
		DependencyIndexes: file_accessgraph_v1alpha_gitlab_proto_depIdxs,
		EnumInfos:         file_accessgraph_v1alpha_gitlab_proto_enumTypes,
		MessageInfos:      file_accessgraph_v1alpha_gitlab_proto_msgTypes,
	}.Build()
	File_accessgraph_v1alpha_gitlab_proto = out.File
	file_accessgraph_v1alpha_gitlab_proto_rawDesc = nil
	file_accessgraph_v1alpha_gitlab_proto_goTypes = nil
	file_accessgraph_v1alpha_gitlab_proto_depIdxs = nil
}
