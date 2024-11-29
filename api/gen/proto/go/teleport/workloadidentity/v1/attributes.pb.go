// Copyright 2024 Gravitational, Inc
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
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: teleport/workloadidentity/v1/attributes.proto

package workloadidentityv1pb

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

// The collection of attributes that result from the join process.
type JoinAttrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection of attributes that result from the join process but are not
	// specific to any particular join method.
	Meta *JoinAttrsMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Attributes that are specific to the GitLab (`gitlab`) join method.
	Gitlab *JoinAttrsGitLab `protobuf:"bytes,2,opt,name=gitlab,proto3" json:"gitlab,omitempty"`
	// Attributes that are specific to the GitHub (`github`) join method.
	Github *JoinAttrsGitHub `protobuf:"bytes,3,opt,name=github,proto3" json:"github,omitempty"`
	// Attributes that are specific to the AWS IAM (`iam`) join method.
	Iam *JoinAttrsAWSIAM `protobuf:"bytes,4,opt,name=iam,proto3" json:"iam,omitempty"`
	// Attributes that are specific to the TPM (`tpm`) join method.
	Tpm *JoinAttrsTPM `protobuf:"bytes,5,opt,name=tpm,proto3" json:"tpm,omitempty"`
}

func (x *JoinAttrs) Reset() {
	*x = JoinAttrs{}
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinAttrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinAttrs) ProtoMessage() {}

func (x *JoinAttrs) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinAttrs.ProtoReflect.Descriptor instead.
func (*JoinAttrs) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_attributes_proto_rawDescGZIP(), []int{0}
}

func (x *JoinAttrs) GetMeta() *JoinAttrsMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *JoinAttrs) GetGitlab() *JoinAttrsGitLab {
	if x != nil {
		return x.Gitlab
	}
	return nil
}

func (x *JoinAttrs) GetGithub() *JoinAttrsGitHub {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *JoinAttrs) GetIam() *JoinAttrsAWSIAM {
	if x != nil {
		return x.Iam
	}
	return nil
}

func (x *JoinAttrs) GetTpm() *JoinAttrsTPM {
	if x != nil {
		return x.Tpm
	}
	return nil
}

// The collection of attributes that result from the join process but are not
// specific to any particular join method.
type JoinAttrsMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the join token that was used to join.
	//
	// This field is omitted if the join token that was used to join was of the
	// `token` method as in this case, the name of the join token is sensitive.
	//
	// Example: `my-gitlab-join-token`
	JoinTokenName string `protobuf:"bytes,1,opt,name=join_token_name,json=joinTokenName,proto3" json:"join_token_name,omitempty"`
	// The name of the join method that was used to join.
	//
	// Example: `gitlab`
	JoinMethod string `protobuf:"bytes,2,opt,name=join_method,json=joinMethod,proto3" json:"join_method,omitempty"`
}

func (x *JoinAttrsMeta) Reset() {
	*x = JoinAttrsMeta{}
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinAttrsMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinAttrsMeta) ProtoMessage() {}

func (x *JoinAttrsMeta) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinAttrsMeta.ProtoReflect.Descriptor instead.
func (*JoinAttrsMeta) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_attributes_proto_rawDescGZIP(), []int{1}
}

func (x *JoinAttrsMeta) GetJoinTokenName() string {
	if x != nil {
		return x.JoinTokenName
	}
	return ""
}

func (x *JoinAttrsMeta) GetJoinMethod() string {
	if x != nil {
		return x.JoinMethod
	}
	return ""
}

// Attributes that are specific to the GitLab join method.
//
// Typically, these are mapped directly from the claims of the GitLab JWT that
// was used to join. You can view the documentation for those claims at:
// https://docs.gitlab.com/ee/ci/secrets/id_token_authentication.html#token-payload
type JoinAttrsGitLab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The `sub` claim of the GitLab JWT that was used to join.
	// For example: `project_path:mygroup/my-project:ref_type:branch:ref:main`
	Sub string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	// The ref that the pipeline is running against.
	// For example: `main`
	Ref string `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	// The type of ref that the pipeline is running against.
	// This is typically `branch` or `tag`.
	RefType string `protobuf:"bytes,3,opt,name=ref_type,json=refType,proto3" json:"ref_type,omitempty"`
	// Whether or not the ref that the pipeline is running against is protected.
	RefProtected bool `protobuf:"varint,4,opt,name=ref_protected,json=refProtected,proto3" json:"ref_protected,omitempty"`
	// The path of the namespace of the project that the pipeline is running within.
	// For example: `mygroup`
	NamespacePath string `protobuf:"bytes,5,opt,name=namespace_path,json=namespacePath,proto3" json:"namespace_path,omitempty"`
	// The full qualified path of the project that the pipeline is running within.
	// This includes the namespace path.
	// For example: `mygroup/my-project`
	ProjectPath string `protobuf:"bytes,6,opt,name=project_path,json=projectPath,proto3" json:"project_path,omitempty"`
	// The name of the user that triggered the pipeline run.
	UserLogin string `protobuf:"bytes,7,opt,name=user_login,json=userLogin,proto3" json:"user_login,omitempty"`
	// The email of the user that triggered the pipeline run.
	UserEmail string `protobuf:"bytes,8,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// The ID of the pipeline.
	PipelineId string `protobuf:"bytes,9,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	// The source of the pipeline.
	// For example: `push` or `web`
	PipelineSource string `protobuf:"bytes,10,opt,name=pipeline_source,json=pipelineSource,proto3" json:"pipeline_source,omitempty"`
	// The environment the pipeline is running against, if any.
	Environment string `protobuf:"bytes,11,opt,name=environment,proto3" json:"environment,omitempty"`
	// Whether or not the pipeline is running against a protected environment.
	// If there is no configured environment, this field is false.
	EnvironmentProtected bool `protobuf:"varint,12,opt,name=environment_protected,json=environmentProtected,proto3" json:"environment_protected,omitempty"`
	// The ID of the runner that this pipeline is running on.
	RunnerId int64 `protobuf:"varint,13,opt,name=runner_id,json=runnerId,proto3" json:"runner_id,omitempty"`
	// The type of runner that is processing the pipeline.
	// Either `gitlab-hosted` or `self-hosted`.
	RunnerEnvironment string `protobuf:"bytes,14,opt,name=runner_environment,json=runnerEnvironment,proto3" json:"runner_environment,omitempty"`
	// The SHA of the commit that triggered the pipeline run.
	Sha string `protobuf:"bytes,15,opt,name=sha,proto3" json:"sha,omitempty"`
	// The ref URI of the CI config configuring the pipeline.
	CiConfigRefUri string `protobuf:"bytes,16,opt,name=ci_config_ref_uri,json=ciConfigRefUri,proto3" json:"ci_config_ref_uri,omitempty"`
	// The Git SHA of the CI config ref configuring the pipeline.
	CiConfigSha string `protobuf:"bytes,17,opt,name=ci_config_sha,json=ciConfigSha,proto3" json:"ci_config_sha,omitempty"`
}

func (x *JoinAttrsGitLab) Reset() {
	*x = JoinAttrsGitLab{}
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinAttrsGitLab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinAttrsGitLab) ProtoMessage() {}

func (x *JoinAttrsGitLab) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinAttrsGitLab.ProtoReflect.Descriptor instead.
func (*JoinAttrsGitLab) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_attributes_proto_rawDescGZIP(), []int{2}
}

func (x *JoinAttrsGitLab) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *JoinAttrsGitLab) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *JoinAttrsGitLab) GetRefType() string {
	if x != nil {
		return x.RefType
	}
	return ""
}

func (x *JoinAttrsGitLab) GetRefProtected() bool {
	if x != nil {
		return x.RefProtected
	}
	return false
}

func (x *JoinAttrsGitLab) GetNamespacePath() string {
	if x != nil {
		return x.NamespacePath
	}
	return ""
}

func (x *JoinAttrsGitLab) GetProjectPath() string {
	if x != nil {
		return x.ProjectPath
	}
	return ""
}

func (x *JoinAttrsGitLab) GetUserLogin() string {
	if x != nil {
		return x.UserLogin
	}
	return ""
}

func (x *JoinAttrsGitLab) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *JoinAttrsGitLab) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *JoinAttrsGitLab) GetPipelineSource() string {
	if x != nil {
		return x.PipelineSource
	}
	return ""
}

func (x *JoinAttrsGitLab) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *JoinAttrsGitLab) GetEnvironmentProtected() bool {
	if x != nil {
		return x.EnvironmentProtected
	}
	return false
}

func (x *JoinAttrsGitLab) GetRunnerId() int64 {
	if x != nil {
		return x.RunnerId
	}
	return 0
}

func (x *JoinAttrsGitLab) GetRunnerEnvironment() string {
	if x != nil {
		return x.RunnerEnvironment
	}
	return ""
}

func (x *JoinAttrsGitLab) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *JoinAttrsGitLab) GetCiConfigRefUri() string {
	if x != nil {
		return x.CiConfigRefUri
	}
	return ""
}

func (x *JoinAttrsGitLab) GetCiConfigSha() string {
	if x != nil {
		return x.CiConfigSha
	}
	return ""
}

// Attributes that are specific to the GitHub (`github`) join method.
//
// Typically, these are mapped directly from the claims of the GitHub JWT that
// was used to join. You can view the documentation for those claims at:
// https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect#understanding-the-oidc-token
type JoinAttrsGitHub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The `sub` claim of the GitHub JWT that was used to join.
	Sub string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	// The username of the actor that initiated the workflow run.
	Actor string `protobuf:"bytes,2,opt,name=actor,proto3" json:"actor,omitempty"`
	// The name of the environment that the workflow is running against, if any.
	Environment string `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	// The ref that the workflow is running against.
	Ref string `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	// The type of ref that the workflow is running against.
	// For example, `branch`.
	RefType string `protobuf:"bytes,5,opt,name=ref_type,json=refType,proto3" json:"ref_type,omitempty"`
	// The name of the repository that the workflow is running within.
	Repository string `protobuf:"bytes,6,opt,name=repository,proto3" json:"repository,omitempty"`
	// The name of the owner of the repository that the workflow is running within.
	RepositoryOwner string `protobuf:"bytes,7,opt,name=repository_owner,json=repositoryOwner,proto3" json:"repository_owner,omitempty"`
	// The name of the workflow that is running.
	Workflow string `protobuf:"bytes,8,opt,name=workflow,proto3" json:"workflow,omitempty"`
	// The name of the event that triggered the workflow run.
	EventName string `protobuf:"bytes,9,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	// The SHA of the commit that triggered the workflow run.
	Sha string `protobuf:"bytes,10,opt,name=sha,proto3" json:"sha,omitempty"`
	// The ID of this GitHub actions workflow run.
	RunId string `protobuf:"bytes,11,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *JoinAttrsGitHub) Reset() {
	*x = JoinAttrsGitHub{}
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinAttrsGitHub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinAttrsGitHub) ProtoMessage() {}

func (x *JoinAttrsGitHub) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinAttrsGitHub.ProtoReflect.Descriptor instead.
func (*JoinAttrsGitHub) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_attributes_proto_rawDescGZIP(), []int{3}
}

func (x *JoinAttrsGitHub) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *JoinAttrsGitHub) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

func (x *JoinAttrsGitHub) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *JoinAttrsGitHub) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *JoinAttrsGitHub) GetRefType() string {
	if x != nil {
		return x.RefType
	}
	return ""
}

func (x *JoinAttrsGitHub) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *JoinAttrsGitHub) GetRepositoryOwner() string {
	if x != nil {
		return x.RepositoryOwner
	}
	return ""
}

func (x *JoinAttrsGitHub) GetWorkflow() string {
	if x != nil {
		return x.Workflow
	}
	return ""
}

func (x *JoinAttrsGitHub) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *JoinAttrsGitHub) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *JoinAttrsGitHub) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

// Attributes that are specific to the AWS IAM (`iam`) join method.
//
// Typically, these are mapped directly from the results of the
// STS GetCallerIdentity call that is made as part of the join process.
type JoinAttrsAWSIAM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the account that the joining entity is a part of.
	// For example: `123456789012`
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The AWS ARN of the joining entity.
	// For example: `arn:aws:sts::123456789012:assumed-role/my-role-name/my-role-session-name`
	Arn string `protobuf:"bytes,2,opt,name=arn,proto3" json:"arn,omitempty"`
}

func (x *JoinAttrsAWSIAM) Reset() {
	*x = JoinAttrsAWSIAM{}
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinAttrsAWSIAM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinAttrsAWSIAM) ProtoMessage() {}

func (x *JoinAttrsAWSIAM) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinAttrsAWSIAM.ProtoReflect.Descriptor instead.
func (*JoinAttrsAWSIAM) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_attributes_proto_rawDescGZIP(), []int{4}
}

func (x *JoinAttrsAWSIAM) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *JoinAttrsAWSIAM) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

// Attributes that are specific to the TPM (`tpm`) join method.
type JoinAttrsTPM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The SHA256 hash of the PKIX formatted EK public key, encoded in hex.
	// This effectively identifies a specific TPM.
	EkPubHash string `protobuf:"bytes,1,opt,name=ek_pub_hash,json=ekPubHash,proto3" json:"ek_pub_hash,omitempty"`
	// The serial number of the EK certificate, if present.
	EkCertSerial string `protobuf:"bytes,2,opt,name=ek_cert_serial,json=ekCertSerial,proto3" json:"ek_cert_serial,omitempty"`
	// Whether or not the EK certificate was verified against a certificate
	// authority.
	EkCertVerified bool `protobuf:"varint,3,opt,name=ek_cert_verified,json=ekCertVerified,proto3" json:"ek_cert_verified,omitempty"`
}

func (x *JoinAttrsTPM) Reset() {
	*x = JoinAttrsTPM{}
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinAttrsTPM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinAttrsTPM) ProtoMessage() {}

func (x *JoinAttrsTPM) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_attributes_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinAttrsTPM.ProtoReflect.Descriptor instead.
func (*JoinAttrsTPM) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_attributes_proto_rawDescGZIP(), []int{5}
}

func (x *JoinAttrsTPM) GetEkPubHash() string {
	if x != nil {
		return x.EkPubHash
	}
	return ""
}

func (x *JoinAttrsTPM) GetEkCertSerial() string {
	if x != nil {
		return x.EkCertSerial
	}
	return ""
}

func (x *JoinAttrsTPM) GetEkCertVerified() bool {
	if x != nil {
		return x.EkCertVerified
	}
	return false
}

var File_teleport_workloadidentity_v1_attributes_proto protoreflect.FileDescriptor

var file_teleport_workloadidentity_v1_attributes_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0xd9, 0x02,
	0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74,
	0x72, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x06,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x41, 0x74, 0x74, 0x72, 0x73, 0x47, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x52, 0x06, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x12, 0x45, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x47, 0x69, 0x74, 0x48,
	0x75, 0x62, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x3f, 0x0a, 0x03, 0x69, 0x61,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73,
	0x41, 0x57, 0x53, 0x49, 0x41, 0x4d, 0x52, 0x03, 0x69, 0x61, 0x6d, 0x12, 0x3c, 0x0a, 0x03, 0x74,
	0x70, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72,
	0x73, 0x54, 0x50, 0x4d, 0x52, 0x03, 0x74, 0x70, 0x6d, 0x22, 0x58, 0x0a, 0x0d, 0x4a, 0x6f, 0x69,
	0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x6a, 0x6f,
	0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x22, 0xcb, 0x04, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72,
	0x73, 0x47, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x33, 0x0a, 0x15, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x68, 0x61, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x72, 0x65, 0x66, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x66, 0x55, 0x72, 0x69, 0x12, 0x22, 0x0a,
	0x0d, 0x63, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x68,
	0x61, 0x22, 0xb7, 0x02, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x47,
	0x69, 0x74, 0x48, 0x75, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65,
	0x66, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x68, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x0f, 0x4a,
	0x6f, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x41, 0x57, 0x53, 0x49, 0x41, 0x4d, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x22, 0x7e, 0x0a, 0x0c, 0x4a, 0x6f,
	0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x54, 0x50, 0x4d, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6b,
	0x5f, 0x70, 0x75, 0x62, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x6b, 0x50, 0x75, 0x62, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6b,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x6b, 0x43, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x12, 0x28, 0x0a, 0x10, 0x65, 0x6b, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6b, 0x43, 0x65,
	0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x66, 0x5a, 0x64, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x76, 0x31,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_workloadidentity_v1_attributes_proto_rawDescOnce sync.Once
	file_teleport_workloadidentity_v1_attributes_proto_rawDescData = file_teleport_workloadidentity_v1_attributes_proto_rawDesc
)

func file_teleport_workloadidentity_v1_attributes_proto_rawDescGZIP() []byte {
	file_teleport_workloadidentity_v1_attributes_proto_rawDescOnce.Do(func() {
		file_teleport_workloadidentity_v1_attributes_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_workloadidentity_v1_attributes_proto_rawDescData)
	})
	return file_teleport_workloadidentity_v1_attributes_proto_rawDescData
}

var file_teleport_workloadidentity_v1_attributes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_teleport_workloadidentity_v1_attributes_proto_goTypes = []any{
	(*JoinAttrs)(nil),       // 0: teleport.workloadidentity.v1.JoinAttrs
	(*JoinAttrsMeta)(nil),   // 1: teleport.workloadidentity.v1.JoinAttrsMeta
	(*JoinAttrsGitLab)(nil), // 2: teleport.workloadidentity.v1.JoinAttrsGitLab
	(*JoinAttrsGitHub)(nil), // 3: teleport.workloadidentity.v1.JoinAttrsGitHub
	(*JoinAttrsAWSIAM)(nil), // 4: teleport.workloadidentity.v1.JoinAttrsAWSIAM
	(*JoinAttrsTPM)(nil),    // 5: teleport.workloadidentity.v1.JoinAttrsTPM
}
var file_teleport_workloadidentity_v1_attributes_proto_depIdxs = []int32{
	1, // 0: teleport.workloadidentity.v1.JoinAttrs.meta:type_name -> teleport.workloadidentity.v1.JoinAttrsMeta
	2, // 1: teleport.workloadidentity.v1.JoinAttrs.gitlab:type_name -> teleport.workloadidentity.v1.JoinAttrsGitLab
	3, // 2: teleport.workloadidentity.v1.JoinAttrs.github:type_name -> teleport.workloadidentity.v1.JoinAttrsGitHub
	4, // 3: teleport.workloadidentity.v1.JoinAttrs.iam:type_name -> teleport.workloadidentity.v1.JoinAttrsAWSIAM
	5, // 4: teleport.workloadidentity.v1.JoinAttrs.tpm:type_name -> teleport.workloadidentity.v1.JoinAttrsTPM
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_teleport_workloadidentity_v1_attributes_proto_init() }
func file_teleport_workloadidentity_v1_attributes_proto_init() {
	if File_teleport_workloadidentity_v1_attributes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_workloadidentity_v1_attributes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_workloadidentity_v1_attributes_proto_goTypes,
		DependencyIndexes: file_teleport_workloadidentity_v1_attributes_proto_depIdxs,
		MessageInfos:      file_teleport_workloadidentity_v1_attributes_proto_msgTypes,
	}.Build()
	File_teleport_workloadidentity_v1_attributes_proto = out.File
	file_teleport_workloadidentity_v1_attributes_proto_rawDesc = nil
	file_teleport_workloadidentity_v1_attributes_proto_goTypes = nil
	file_teleport_workloadidentity_v1_attributes_proto_depIdxs = nil
}
