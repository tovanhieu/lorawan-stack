// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: lorawan-stack/api/devicerepository.proto

package ttnpb

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
	DeviceRepository_ListBrands_FullMethodName         = "/ttn.lorawan.v3.DeviceRepository/ListBrands"
	DeviceRepository_GetBrand_FullMethodName           = "/ttn.lorawan.v3.DeviceRepository/GetBrand"
	DeviceRepository_ListModels_FullMethodName         = "/ttn.lorawan.v3.DeviceRepository/ListModels"
	DeviceRepository_GetModel_FullMethodName           = "/ttn.lorawan.v3.DeviceRepository/GetModel"
	DeviceRepository_GetTemplate_FullMethodName        = "/ttn.lorawan.v3.DeviceRepository/GetTemplate"
	DeviceRepository_GetUplinkDecoder_FullMethodName   = "/ttn.lorawan.v3.DeviceRepository/GetUplinkDecoder"
	DeviceRepository_GetDownlinkDecoder_FullMethodName = "/ttn.lorawan.v3.DeviceRepository/GetDownlinkDecoder"
	DeviceRepository_GetDownlinkEncoder_FullMethodName = "/ttn.lorawan.v3.DeviceRepository/GetDownlinkEncoder"
)

// DeviceRepositoryClient is the client API for DeviceRepository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceRepositoryClient interface {
	ListBrands(ctx context.Context, in *ListEndDeviceBrandsRequest, opts ...grpc.CallOption) (*ListEndDeviceBrandsResponse, error)
	GetBrand(ctx context.Context, in *GetEndDeviceBrandRequest, opts ...grpc.CallOption) (*EndDeviceBrand, error)
	ListModels(ctx context.Context, in *ListEndDeviceModelsRequest, opts ...grpc.CallOption) (*ListEndDeviceModelsResponse, error)
	GetModel(ctx context.Context, in *GetEndDeviceModelRequest, opts ...grpc.CallOption) (*EndDeviceModel, error)
	GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*EndDeviceTemplate, error)
	GetUplinkDecoder(ctx context.Context, in *GetPayloadFormatterRequest, opts ...grpc.CallOption) (*MessagePayloadDecoder, error)
	GetDownlinkDecoder(ctx context.Context, in *GetPayloadFormatterRequest, opts ...grpc.CallOption) (*MessagePayloadDecoder, error)
	GetDownlinkEncoder(ctx context.Context, in *GetPayloadFormatterRequest, opts ...grpc.CallOption) (*MessagePayloadEncoder, error)
}

type deviceRepositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceRepositoryClient(cc grpc.ClientConnInterface) DeviceRepositoryClient {
	return &deviceRepositoryClient{cc}
}

func (c *deviceRepositoryClient) ListBrands(ctx context.Context, in *ListEndDeviceBrandsRequest, opts ...grpc.CallOption) (*ListEndDeviceBrandsResponse, error) {
	out := new(ListEndDeviceBrandsResponse)
	err := c.cc.Invoke(ctx, DeviceRepository_ListBrands_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRepositoryClient) GetBrand(ctx context.Context, in *GetEndDeviceBrandRequest, opts ...grpc.CallOption) (*EndDeviceBrand, error) {
	out := new(EndDeviceBrand)
	err := c.cc.Invoke(ctx, DeviceRepository_GetBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRepositoryClient) ListModels(ctx context.Context, in *ListEndDeviceModelsRequest, opts ...grpc.CallOption) (*ListEndDeviceModelsResponse, error) {
	out := new(ListEndDeviceModelsResponse)
	err := c.cc.Invoke(ctx, DeviceRepository_ListModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRepositoryClient) GetModel(ctx context.Context, in *GetEndDeviceModelRequest, opts ...grpc.CallOption) (*EndDeviceModel, error) {
	out := new(EndDeviceModel)
	err := c.cc.Invoke(ctx, DeviceRepository_GetModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRepositoryClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*EndDeviceTemplate, error) {
	out := new(EndDeviceTemplate)
	err := c.cc.Invoke(ctx, DeviceRepository_GetTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRepositoryClient) GetUplinkDecoder(ctx context.Context, in *GetPayloadFormatterRequest, opts ...grpc.CallOption) (*MessagePayloadDecoder, error) {
	out := new(MessagePayloadDecoder)
	err := c.cc.Invoke(ctx, DeviceRepository_GetUplinkDecoder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRepositoryClient) GetDownlinkDecoder(ctx context.Context, in *GetPayloadFormatterRequest, opts ...grpc.CallOption) (*MessagePayloadDecoder, error) {
	out := new(MessagePayloadDecoder)
	err := c.cc.Invoke(ctx, DeviceRepository_GetDownlinkDecoder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRepositoryClient) GetDownlinkEncoder(ctx context.Context, in *GetPayloadFormatterRequest, opts ...grpc.CallOption) (*MessagePayloadEncoder, error) {
	out := new(MessagePayloadEncoder)
	err := c.cc.Invoke(ctx, DeviceRepository_GetDownlinkEncoder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceRepositoryServer is the server API for DeviceRepository service.
// All implementations must embed UnimplementedDeviceRepositoryServer
// for forward compatibility
type DeviceRepositoryServer interface {
	ListBrands(context.Context, *ListEndDeviceBrandsRequest) (*ListEndDeviceBrandsResponse, error)
	GetBrand(context.Context, *GetEndDeviceBrandRequest) (*EndDeviceBrand, error)
	ListModels(context.Context, *ListEndDeviceModelsRequest) (*ListEndDeviceModelsResponse, error)
	GetModel(context.Context, *GetEndDeviceModelRequest) (*EndDeviceModel, error)
	GetTemplate(context.Context, *GetTemplateRequest) (*EndDeviceTemplate, error)
	GetUplinkDecoder(context.Context, *GetPayloadFormatterRequest) (*MessagePayloadDecoder, error)
	GetDownlinkDecoder(context.Context, *GetPayloadFormatterRequest) (*MessagePayloadDecoder, error)
	GetDownlinkEncoder(context.Context, *GetPayloadFormatterRequest) (*MessagePayloadEncoder, error)
	mustEmbedUnimplementedDeviceRepositoryServer()
}

// UnimplementedDeviceRepositoryServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceRepositoryServer struct {
}

func (UnimplementedDeviceRepositoryServer) ListBrands(context.Context, *ListEndDeviceBrandsRequest) (*ListEndDeviceBrandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBrands not implemented")
}
func (UnimplementedDeviceRepositoryServer) GetBrand(context.Context, *GetEndDeviceBrandRequest) (*EndDeviceBrand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}
func (UnimplementedDeviceRepositoryServer) ListModels(context.Context, *ListEndDeviceModelsRequest) (*ListEndDeviceModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModels not implemented")
}
func (UnimplementedDeviceRepositoryServer) GetModel(context.Context, *GetEndDeviceModelRequest) (*EndDeviceModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModel not implemented")
}
func (UnimplementedDeviceRepositoryServer) GetTemplate(context.Context, *GetTemplateRequest) (*EndDeviceTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedDeviceRepositoryServer) GetUplinkDecoder(context.Context, *GetPayloadFormatterRequest) (*MessagePayloadDecoder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUplinkDecoder not implemented")
}
func (UnimplementedDeviceRepositoryServer) GetDownlinkDecoder(context.Context, *GetPayloadFormatterRequest) (*MessagePayloadDecoder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownlinkDecoder not implemented")
}
func (UnimplementedDeviceRepositoryServer) GetDownlinkEncoder(context.Context, *GetPayloadFormatterRequest) (*MessagePayloadEncoder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownlinkEncoder not implemented")
}
func (UnimplementedDeviceRepositoryServer) mustEmbedUnimplementedDeviceRepositoryServer() {}

// UnsafeDeviceRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceRepositoryServer will
// result in compilation errors.
type UnsafeDeviceRepositoryServer interface {
	mustEmbedUnimplementedDeviceRepositoryServer()
}

func RegisterDeviceRepositoryServer(s grpc.ServiceRegistrar, srv DeviceRepositoryServer) {
	s.RegisterService(&DeviceRepository_ServiceDesc, srv)
}

func _DeviceRepository_ListBrands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndDeviceBrandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).ListBrands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_ListBrands_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).ListBrands(ctx, req.(*ListEndDeviceBrandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRepository_GetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).GetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_GetBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).GetBrand(ctx, req.(*GetEndDeviceBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRepository_ListModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndDeviceModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).ListModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_ListModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).ListModels(ctx, req.(*ListEndDeviceModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRepository_GetModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndDeviceModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).GetModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_GetModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).GetModel(ctx, req.(*GetEndDeviceModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRepository_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_GetTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).GetTemplate(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRepository_GetUplinkDecoder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayloadFormatterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).GetUplinkDecoder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_GetUplinkDecoder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).GetUplinkDecoder(ctx, req.(*GetPayloadFormatterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRepository_GetDownlinkDecoder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayloadFormatterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).GetDownlinkDecoder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_GetDownlinkDecoder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).GetDownlinkDecoder(ctx, req.(*GetPayloadFormatterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRepository_GetDownlinkEncoder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayloadFormatterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRepositoryServer).GetDownlinkEncoder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRepository_GetDownlinkEncoder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRepositoryServer).GetDownlinkEncoder(ctx, req.(*GetPayloadFormatterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceRepository_ServiceDesc is the grpc.ServiceDesc for DeviceRepository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceRepository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.DeviceRepository",
	HandlerType: (*DeviceRepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBrands",
			Handler:    _DeviceRepository_ListBrands_Handler,
		},
		{
			MethodName: "GetBrand",
			Handler:    _DeviceRepository_GetBrand_Handler,
		},
		{
			MethodName: "ListModels",
			Handler:    _DeviceRepository_ListModels_Handler,
		},
		{
			MethodName: "GetModel",
			Handler:    _DeviceRepository_GetModel_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _DeviceRepository_GetTemplate_Handler,
		},
		{
			MethodName: "GetUplinkDecoder",
			Handler:    _DeviceRepository_GetUplinkDecoder_Handler,
		},
		{
			MethodName: "GetDownlinkDecoder",
			Handler:    _DeviceRepository_GetDownlinkDecoder_Handler,
		},
		{
			MethodName: "GetDownlinkEncoder",
			Handler:    _DeviceRepository_GetDownlinkEncoder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/devicerepository.proto",
}
