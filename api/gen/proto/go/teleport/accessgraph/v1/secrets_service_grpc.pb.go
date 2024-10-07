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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/access_graph/v1/secrets_service.proto

package accessgraphv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SecretsScannerService_ReportAuthorizedKeys_FullMethodName = "/teleport.access_graph.v1.SecretsScannerService/ReportAuthorizedKeys"
	SecretsScannerService_ReportSecrets_FullMethodName        = "/teleport.access_graph.v1.SecretsScannerService/ReportSecrets"
)

// SecretsScannerServiceClient is the client API for SecretsScannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretsScannerServiceClient interface {
	// ReportAuthorizedKeys is used by Teleport SSH nodes to report authorized keys
	// that could be used to bypass Teleport.
	// The client (Teleport SSH Node) should authenticate using the certificate-key pair signed by Teleport HostCA.
	ReportAuthorizedKeys(ctx context.Context, opts ...grpc.CallOption) (SecretsScannerService_ReportAuthorizedKeysClient, error)
	// ReportSecrets is used by trusted devices to report secrets found on the host that could be used to bypass Teleport.
	// The client (device) should first authenticate using the [ReportSecretsRequest.device_assertion] flow. Please refer to
	// the [teleport.devicetrust.v1.AssertDeviceRequest] and [teleport.devicetrust.v1.AssertDeviceResponse] messages for more details.
	//
	// Once the device is asserted, the client can send the secrets using the [ReportSecretsRequest.private_keys] field
	// and then close the client side of the stream.
	//
	// -> ReportSecrets (client) [1 or more]
	// -> CloseStream (client)
	// <- TerminateStream (server)
	//
	// Any failure in the assertion ceremony will result in the stream being terminated by the server. All secrets
	// reported by the client before the assertion terminates will be ignored and result in the stream being terminated.
	ReportSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretsScannerService_ReportSecretsClient, error)
}

type secretsScannerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretsScannerServiceClient(cc grpc.ClientConnInterface) SecretsScannerServiceClient {
	return &secretsScannerServiceClient{cc}
}

func (c *secretsScannerServiceClient) ReportAuthorizedKeys(ctx context.Context, opts ...grpc.CallOption) (SecretsScannerService_ReportAuthorizedKeysClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecretsScannerService_ServiceDesc.Streams[0], SecretsScannerService_ReportAuthorizedKeys_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &secretsScannerServiceReportAuthorizedKeysClient{stream}
	return x, nil
}

type SecretsScannerService_ReportAuthorizedKeysClient interface {
	Send(*ReportAuthorizedKeysRequest) error
	Recv() (*ReportAuthorizedKeysResponse, error)
	grpc.ClientStream
}

type secretsScannerServiceReportAuthorizedKeysClient struct {
	grpc.ClientStream
}

func (x *secretsScannerServiceReportAuthorizedKeysClient) Send(m *ReportAuthorizedKeysRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretsScannerServiceReportAuthorizedKeysClient) Recv() (*ReportAuthorizedKeysResponse, error) {
	m := new(ReportAuthorizedKeysResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretsScannerServiceClient) ReportSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretsScannerService_ReportSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecretsScannerService_ServiceDesc.Streams[1], SecretsScannerService_ReportSecrets_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &secretsScannerServiceReportSecretsClient{stream}
	return x, nil
}

type SecretsScannerService_ReportSecretsClient interface {
	Send(*ReportSecretsRequest) error
	Recv() (*ReportSecretsResponse, error)
	grpc.ClientStream
}

type secretsScannerServiceReportSecretsClient struct {
	grpc.ClientStream
}

func (x *secretsScannerServiceReportSecretsClient) Send(m *ReportSecretsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretsScannerServiceReportSecretsClient) Recv() (*ReportSecretsResponse, error) {
	m := new(ReportSecretsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SecretsScannerServiceServer is the server API for SecretsScannerService service.
// All implementations must embed UnimplementedSecretsScannerServiceServer
// for forward compatibility
type SecretsScannerServiceServer interface {
	// ReportAuthorizedKeys is used by Teleport SSH nodes to report authorized keys
	// that could be used to bypass Teleport.
	// The client (Teleport SSH Node) should authenticate using the certificate-key pair signed by Teleport HostCA.
	ReportAuthorizedKeys(SecretsScannerService_ReportAuthorizedKeysServer) error
	// ReportSecrets is used by trusted devices to report secrets found on the host that could be used to bypass Teleport.
	// The client (device) should first authenticate using the [ReportSecretsRequest.device_assertion] flow. Please refer to
	// the [teleport.devicetrust.v1.AssertDeviceRequest] and [teleport.devicetrust.v1.AssertDeviceResponse] messages for more details.
	//
	// Once the device is asserted, the client can send the secrets using the [ReportSecretsRequest.private_keys] field
	// and then close the client side of the stream.
	//
	// -> ReportSecrets (client) [1 or more]
	// -> CloseStream (client)
	// <- TerminateStream (server)
	//
	// Any failure in the assertion ceremony will result in the stream being terminated by the server. All secrets
	// reported by the client before the assertion terminates will be ignored and result in the stream being terminated.
	ReportSecrets(SecretsScannerService_ReportSecretsServer) error
	mustEmbedUnimplementedSecretsScannerServiceServer()
}

// UnimplementedSecretsScannerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretsScannerServiceServer struct {
}

func (UnimplementedSecretsScannerServiceServer) ReportAuthorizedKeys(SecretsScannerService_ReportAuthorizedKeysServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportAuthorizedKeys not implemented")
}
func (UnimplementedSecretsScannerServiceServer) ReportSecrets(SecretsScannerService_ReportSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportSecrets not implemented")
}
func (UnimplementedSecretsScannerServiceServer) mustEmbedUnimplementedSecretsScannerServiceServer() {}

// UnsafeSecretsScannerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretsScannerServiceServer will
// result in compilation errors.
type UnsafeSecretsScannerServiceServer interface {
	mustEmbedUnimplementedSecretsScannerServiceServer()
}

func RegisterSecretsScannerServiceServer(s grpc.ServiceRegistrar, srv SecretsScannerServiceServer) {
	s.RegisterService(&SecretsScannerService_ServiceDesc, srv)
}

func _SecretsScannerService_ReportAuthorizedKeys_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretsScannerServiceServer).ReportAuthorizedKeys(&secretsScannerServiceReportAuthorizedKeysServer{stream})
}

type SecretsScannerService_ReportAuthorizedKeysServer interface {
	Send(*ReportAuthorizedKeysResponse) error
	Recv() (*ReportAuthorizedKeysRequest, error)
	grpc.ServerStream
}

type secretsScannerServiceReportAuthorizedKeysServer struct {
	grpc.ServerStream
}

func (x *secretsScannerServiceReportAuthorizedKeysServer) Send(m *ReportAuthorizedKeysResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretsScannerServiceReportAuthorizedKeysServer) Recv() (*ReportAuthorizedKeysRequest, error) {
	m := new(ReportAuthorizedKeysRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretsScannerService_ReportSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretsScannerServiceServer).ReportSecrets(&secretsScannerServiceReportSecretsServer{stream})
}

type SecretsScannerService_ReportSecretsServer interface {
	Send(*ReportSecretsResponse) error
	Recv() (*ReportSecretsRequest, error)
	grpc.ServerStream
}

type secretsScannerServiceReportSecretsServer struct {
	grpc.ServerStream
}

func (x *secretsScannerServiceReportSecretsServer) Send(m *ReportSecretsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretsScannerServiceReportSecretsServer) Recv() (*ReportSecretsRequest, error) {
	m := new(ReportSecretsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SecretsScannerService_ServiceDesc is the grpc.ServiceDesc for SecretsScannerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretsScannerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.access_graph.v1.SecretsScannerService",
	HandlerType: (*SecretsScannerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReportAuthorizedKeys",
			Handler:       _SecretsScannerService_ReportAuthorizedKeys_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReportSecrets",
			Handler:       _SecretsScannerService_ReportSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "teleport/access_graph/v1/secrets_service.proto",
}
