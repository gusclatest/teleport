// Teleport
// Copyright (C) 2024 Gravitational, Inc.
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/lib/teleterm/vnet/v1/vnet_service.proto

package vnetv1

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
	VnetService_Start_FullMethodName                   = "/teleport.lib.teleterm.vnet.v1.VnetService/Start"
	VnetService_Stop_FullMethodName                    = "/teleport.lib.teleterm.vnet.v1.VnetService/Stop"
	VnetService_ListDNSZones_FullMethodName            = "/teleport.lib.teleterm.vnet.v1.VnetService/ListDNSZones"
	VnetService_GetBackgroundItemStatus_FullMethodName = "/teleport.lib.teleterm.vnet.v1.VnetService/GetBackgroundItemStatus"
)

// VnetServiceClient is the client API for VnetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VnetServiceClient interface {
	// Start starts VNet.
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	// Stop stops VNet.
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	// ListDNSZones returns DNS zones of all root and leaf clusters with non-expired user certs. This
	// includes the proxy service hostnames and custom DNS zones configured in vnet_config.
	//
	// This is fetched independently of what the Electron app thinks the current state of the cluster
	// looks like, since the VNet admin process also fetches this data independently of the Electron
	// app.
	//
	// Just like the admin process, it skips root and leaf clusters for which the vnet_config couldn't
	// be fetched (due to e.g., a network error or an expired cert).
	ListDNSZones(ctx context.Context, in *ListDNSZonesRequest, opts ...grpc.CallOption) (*ListDNSZonesResponse, error)
	// GetBackgroundItemStatus returns the status of the background item responsible for launching
	// VNet daemon. macOS only. tsh must be compiled with the vnetdaemon build tag.
	GetBackgroundItemStatus(ctx context.Context, in *GetBackgroundItemStatusRequest, opts ...grpc.CallOption) (*GetBackgroundItemStatusResponse, error)
}

type vnetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVnetServiceClient(cc grpc.ClientConnInterface) VnetServiceClient {
	return &vnetServiceClient{cc}
}

func (c *vnetServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, VnetService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnetServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, VnetService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnetServiceClient) ListDNSZones(ctx context.Context, in *ListDNSZonesRequest, opts ...grpc.CallOption) (*ListDNSZonesResponse, error) {
	out := new(ListDNSZonesResponse)
	err := c.cc.Invoke(ctx, VnetService_ListDNSZones_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnetServiceClient) GetBackgroundItemStatus(ctx context.Context, in *GetBackgroundItemStatusRequest, opts ...grpc.CallOption) (*GetBackgroundItemStatusResponse, error) {
	out := new(GetBackgroundItemStatusResponse)
	err := c.cc.Invoke(ctx, VnetService_GetBackgroundItemStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VnetServiceServer is the server API for VnetService service.
// All implementations must embed UnimplementedVnetServiceServer
// for forward compatibility
type VnetServiceServer interface {
	// Start starts VNet.
	Start(context.Context, *StartRequest) (*StartResponse, error)
	// Stop stops VNet.
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	// ListDNSZones returns DNS zones of all root and leaf clusters with non-expired user certs. This
	// includes the proxy service hostnames and custom DNS zones configured in vnet_config.
	//
	// This is fetched independently of what the Electron app thinks the current state of the cluster
	// looks like, since the VNet admin process also fetches this data independently of the Electron
	// app.
	//
	// Just like the admin process, it skips root and leaf clusters for which the vnet_config couldn't
	// be fetched (due to e.g., a network error or an expired cert).
	ListDNSZones(context.Context, *ListDNSZonesRequest) (*ListDNSZonesResponse, error)
	// GetBackgroundItemStatus returns the status of the background item responsible for launching
	// VNet daemon. macOS only. tsh must be compiled with the vnetdaemon build tag.
	GetBackgroundItemStatus(context.Context, *GetBackgroundItemStatusRequest) (*GetBackgroundItemStatusResponse, error)
	mustEmbedUnimplementedVnetServiceServer()
}

// UnimplementedVnetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVnetServiceServer struct {
}

func (UnimplementedVnetServiceServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedVnetServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedVnetServiceServer) ListDNSZones(context.Context, *ListDNSZonesRequest) (*ListDNSZonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDNSZones not implemented")
}
func (UnimplementedVnetServiceServer) GetBackgroundItemStatus(context.Context, *GetBackgroundItemStatusRequest) (*GetBackgroundItemStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackgroundItemStatus not implemented")
}
func (UnimplementedVnetServiceServer) mustEmbedUnimplementedVnetServiceServer() {}

// UnsafeVnetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VnetServiceServer will
// result in compilation errors.
type UnsafeVnetServiceServer interface {
	mustEmbedUnimplementedVnetServiceServer()
}

func RegisterVnetServiceServer(s grpc.ServiceRegistrar, srv VnetServiceServer) {
	s.RegisterService(&VnetService_ServiceDesc, srv)
}

func _VnetService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnetServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VnetService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnetServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VnetService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnetServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VnetService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnetServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VnetService_ListDNSZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDNSZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnetServiceServer).ListDNSZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VnetService_ListDNSZones_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnetServiceServer).ListDNSZones(ctx, req.(*ListDNSZonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VnetService_GetBackgroundItemStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackgroundItemStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnetServiceServer).GetBackgroundItemStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VnetService_GetBackgroundItemStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnetServiceServer).GetBackgroundItemStatus(ctx, req.(*GetBackgroundItemStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VnetService_ServiceDesc is the grpc.ServiceDesc for VnetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VnetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.lib.teleterm.vnet.v1.VnetService",
	HandlerType: (*VnetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _VnetService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _VnetService_Stop_Handler,
		},
		{
			MethodName: "ListDNSZones",
			Handler:    _VnetService_ListDNSZones_Handler,
		},
		{
			MethodName: "GetBackgroundItemStatus",
			Handler:    _VnetService_GetBackgroundItemStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/lib/teleterm/vnet/v1/vnet_service.proto",
}
