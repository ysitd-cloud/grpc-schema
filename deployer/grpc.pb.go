// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deployer/grpc.proto

/*
Package deployer is a generated protocol buffer package.

It is generated from these files:
	deployer/grpc.proto

It has these top-level messages:
*/
package deployer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import deployer2 "code.ysitd.cloud/grpc/schema/deployer/actions"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Deployer service

type DeployerClient interface {
	ListApplicationsByUsername(ctx context.Context, in *deployer2.ListApplicationsByUsernameRequest, opts ...grpc.CallOption) (*deployer2.ListApplicationsByUsernameReply, error)
	CreateApplication(ctx context.Context, in *deployer2.CreateApplicationRequest, opts ...grpc.CallOption) (*deployer2.CreateApplicationReply, error)
	UpdateDeploymentImage(ctx context.Context, in *deployer2.UpdateDeploymentImageRequest, opts ...grpc.CallOption) (*deployer2.UpdateDeploymentImageReply, error)
	GetApplicationById(ctx context.Context, in *deployer2.GetApplicationByIdRequest, opts ...grpc.CallOption) (*deployer2.GetApplicationByIdResponse, error)
}

type deployerClient struct {
	cc *grpc.ClientConn
}

func NewDeployerClient(cc *grpc.ClientConn) DeployerClient {
	return &deployerClient{cc}
}

func (c *deployerClient) ListApplicationsByUsername(ctx context.Context, in *deployer2.ListApplicationsByUsernameRequest, opts ...grpc.CallOption) (*deployer2.ListApplicationsByUsernameReply, error) {
	out := new(deployer2.ListApplicationsByUsernameReply)
	err := grpc.Invoke(ctx, "/deployer.Deployer/listApplicationsByUsername", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployerClient) CreateApplication(ctx context.Context, in *deployer2.CreateApplicationRequest, opts ...grpc.CallOption) (*deployer2.CreateApplicationReply, error) {
	out := new(deployer2.CreateApplicationReply)
	err := grpc.Invoke(ctx, "/deployer.Deployer/createApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployerClient) UpdateDeploymentImage(ctx context.Context, in *deployer2.UpdateDeploymentImageRequest, opts ...grpc.CallOption) (*deployer2.UpdateDeploymentImageReply, error) {
	out := new(deployer2.UpdateDeploymentImageReply)
	err := grpc.Invoke(ctx, "/deployer.Deployer/updateDeploymentImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployerClient) GetApplicationById(ctx context.Context, in *deployer2.GetApplicationByIdRequest, opts ...grpc.CallOption) (*deployer2.GetApplicationByIdResponse, error) {
	out := new(deployer2.GetApplicationByIdResponse)
	err := grpc.Invoke(ctx, "/deployer.Deployer/getApplicationById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Deployer service

type DeployerServer interface {
	ListApplicationsByUsername(context.Context, *deployer2.ListApplicationsByUsernameRequest) (*deployer2.ListApplicationsByUsernameReply, error)
	CreateApplication(context.Context, *deployer2.CreateApplicationRequest) (*deployer2.CreateApplicationReply, error)
	UpdateDeploymentImage(context.Context, *deployer2.UpdateDeploymentImageRequest) (*deployer2.UpdateDeploymentImageReply, error)
	GetApplicationById(context.Context, *deployer2.GetApplicationByIdRequest) (*deployer2.GetApplicationByIdResponse, error)
}

func RegisterDeployerServer(s *grpc.Server, srv DeployerServer) {
	s.RegisterService(&_Deployer_serviceDesc, srv)
}

func _Deployer_ListApplicationsByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(deployer2.ListApplicationsByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerServer).ListApplicationsByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployer.Deployer/ListApplicationsByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerServer).ListApplicationsByUsername(ctx, req.(*deployer2.ListApplicationsByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployer_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(deployer2.CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployer.Deployer/CreateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerServer).CreateApplication(ctx, req.(*deployer2.CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployer_UpdateDeploymentImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(deployer2.UpdateDeploymentImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerServer).UpdateDeploymentImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployer.Deployer/UpdateDeploymentImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerServer).UpdateDeploymentImage(ctx, req.(*deployer2.UpdateDeploymentImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployer_GetApplicationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(deployer2.GetApplicationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerServer).GetApplicationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployer.Deployer/GetApplicationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerServer).GetApplicationById(ctx, req.(*deployer2.GetApplicationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deployer.Deployer",
	HandlerType: (*DeployerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "listApplicationsByUsername",
			Handler:    _Deployer_ListApplicationsByUsername_Handler,
		},
		{
			MethodName: "createApplication",
			Handler:    _Deployer_CreateApplication_Handler,
		},
		{
			MethodName: "updateDeploymentImage",
			Handler:    _Deployer_UpdateDeploymentImage_Handler,
		},
		{
			MethodName: "getApplicationById",
			Handler:    _Deployer_GetApplicationById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deployer/grpc.proto",
}

func init() { proto.RegisterFile("deployer/grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x57, 0x04, 0x59, 0x02, 0x1e, 0x8c, 0x08, 0x52, 0x10, 0xa4, 0xfe, 0x01, 0x11, 0x5a,
	0xd0, 0x4f, 0x60, 0x15, 0x64, 0xc1, 0xc3, 0xb2, 0xb0, 0x17, 0x3d, 0x8d, 0x93, 0xa1, 0x16, 0xda,
	0x26, 0x66, 0x52, 0x21, 0x9f, 0xc9, 0x2f, 0x29, 0xa6, 0x1b, 0xbb, 0x68, 0xd5, 0xbd, 0xe6, 0xfd,
	0xde, 0x7b, 0x84, 0x37, 0x62, 0x5f, 0x91, 0xa9, 0xb5, 0x27, 0x9b, 0x97, 0xd6, 0x60, 0x66, 0xac,
	0x76, 0x5a, 0x4e, 0xe3, 0x63, 0x72, 0xf4, 0x25, 0x03, 0xba, 0x4a, 0xb7, 0x9c, 0x03, 0x22, 0x31,
	0xf7, 0xe0, 0xd5, 0xfb, 0xb6, 0x98, 0xde, 0xad, 0x08, 0xf9, 0x26, 0x92, 0xba, 0x62, 0x77, 0x63,
	0x4c, 0x5d, 0x21, 0x04, 0xba, 0xf0, 0x4b, 0x26, 0xdb, 0x42, 0x43, 0xf2, 0x32, 0x8b, 0x51, 0xd9,
	0xc3, 0xaf, 0xd4, 0x82, 0x5e, 0x3b, 0x62, 0x97, 0x5c, 0x6c, 0x06, 0x9b, 0xda, 0xa7, 0x13, 0xf9,
	0x24, 0xf6, 0xd0, 0x12, 0x38, 0x5a, 0xc3, 0x64, 0x3a, 0x24, 0xdc, 0x7e, 0x17, 0x63, 0xcb, 0xf1,
	0x9f, 0x4c, 0x1f, 0x5e, 0x8a, 0x83, 0xce, 0x28, 0x70, 0xd4, 0x7f, 0xb3, 0xa1, 0xd6, 0xcd, 0x1a,
	0x28, 0x49, 0x9e, 0x0f, 0xe6, 0xe5, 0x18, 0x10, 0x4b, 0x4e, 0xff, 0xe5, 0xfa, 0x22, 0x10, 0xb2,
	0xa4, 0xf5, 0x9f, 0x16, 0x7e, 0xa6, 0xe4, 0xc9, 0xe0, 0xbe, 0xff, 0xa1, 0x8e, 0x54, 0x8c, 0x41,
	0x6c, 0x74, 0xcb, 0x94, 0x4e, 0x8a, 0x85, 0x38, 0xc4, 0x5a, 0x77, 0x2a, 0xf3, 0x5c, 0x39, 0x95,
	0x85, 0xc1, 0x01, 0x51, 0x77, 0xad, 0x2b, 0x76, 0xe3, 0x8c, 0xf3, 0xcf, 0x61, 0xe7, 0x5b, 0x8f,
	0x67, 0xa8, 0x15, 0xad, 0xc8, 0xe0, 0x0a, 0x07, 0x92, 0x33, 0xbe, 0x50, 0x03, 0x79, 0x2c, 0x7c,
	0xde, 0x09, 0x87, 0x70, 0xfd, 0x11, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xce, 0x9a, 0xe5, 0x48, 0x02,
	0x00, 0x00,
}
