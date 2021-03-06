// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/grpc.proto

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	account/grpc.proto

It has these top-level messages:
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import account2 "code.ysitd.cloud/grpc/schema/account/actions"

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

// Client API for Account service

type AccountClient interface {
	ValidateUserPassword(ctx context.Context, in *account2.ValidateUserRequest, opts ...grpc.CallOption) (*account2.ValidateUserReply, error)
	GetUserInfo(ctx context.Context, in *account2.GetUserInfoRequest, opts ...grpc.CallOption) (*account2.GetUserInfoReply, error)
	GetTokenInfo(ctx context.Context, in *account2.GetTokenInfoRequest, opts ...grpc.CallOption) (*account2.GetTokenInfoReply, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) ValidateUserPassword(ctx context.Context, in *account2.ValidateUserRequest, opts ...grpc.CallOption) (*account2.ValidateUserReply, error) {
	out := new(account2.ValidateUserReply)
	err := grpc.Invoke(ctx, "/account.Account/validateUserPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUserInfo(ctx context.Context, in *account2.GetUserInfoRequest, opts ...grpc.CallOption) (*account2.GetUserInfoReply, error) {
	out := new(account2.GetUserInfoReply)
	err := grpc.Invoke(ctx, "/account.Account/getUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetTokenInfo(ctx context.Context, in *account2.GetTokenInfoRequest, opts ...grpc.CallOption) (*account2.GetTokenInfoReply, error) {
	out := new(account2.GetTokenInfoReply)
	err := grpc.Invoke(ctx, "/account.Account/getTokenInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountServer interface {
	ValidateUserPassword(context.Context, *account2.ValidateUserRequest) (*account2.ValidateUserReply, error)
	GetUserInfo(context.Context, *account2.GetUserInfoRequest) (*account2.GetUserInfoReply, error)
	GetTokenInfo(context.Context, *account2.GetTokenInfoRequest) (*account2.GetTokenInfoReply, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_ValidateUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account2.ValidateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).ValidateUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/ValidateUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).ValidateUserPassword(ctx, req.(*account2.ValidateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account2.GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUserInfo(ctx, req.(*account2.GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account2.GetTokenInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/GetTokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetTokenInfo(ctx, req.(*account2.GetTokenInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "validateUserPassword",
			Handler:    _Account_ValidateUserPassword_Handler,
		},
		{
			MethodName: "getUserInfo",
			Handler:    _Account_GetUserInfo_Handler,
		},
		{
			MethodName: "getTokenInfo",
			Handler:    _Account_GetTokenInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/grpc.proto",
}

func init() { proto.RegisterFile("account/grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0x85, 0x30,
	0x14, 0x86, 0xbd, 0x8b, 0x37, 0xa9, 0x77, 0x6a, 0x1c, 0x14, 0x71, 0x31, 0xce, 0x25, 0xd1, 0x27,
	0x90, 0x85, 0xe8, 0xd4, 0x10, 0x75, 0x70, 0xab, 0xed, 0x11, 0x89, 0xd8, 0x83, 0xed, 0x41, 0xc3,
	0x6b, 0xfb, 0x04, 0x86, 0xd2, 0x22, 0x46, 0x1d, 0xfb, 0xff, 0xdf, 0xf9, 0xd3, 0x7c, 0x8c, 0x2b,
	0xad, 0x71, 0xb0, 0x54, 0x34, 0xae, 0xd7, 0xa2, 0x77, 0x48, 0xc8, 0xb7, 0x31, 0xcb, 0x4e, 0x53,
	0xa9, 0x34, 0xb5, 0x68, 0x7d, 0x11, 0xdf, 0x33, 0x77, 0xf1, 0xb9, 0x61, 0xdb, 0xab, 0x39, 0xe1,
	0x35, 0x3b, 0x7c, 0x57, 0x5d, 0x6b, 0x14, 0xc1, 0x9d, 0x07, 0x27, 0x95, 0xf7, 0x1f, 0xe8, 0x0c,
	0xcf, 0x45, 0xba, 0xb9, 0x5f, 0xd5, 0x35, 0xbc, 0x0d, 0xe0, 0x29, 0xcb, 0xfe, 0x69, 0xfb, 0x6e,
	0x3c, 0xdb, 0xe3, 0x15, 0x3b, 0x68, 0x80, 0xa6, 0xe4, 0xda, 0x3e, 0x21, 0x3f, 0x59, 0xe0, 0xea,
	0x3b, 0x4d, 0x4b, 0xc7, 0x7f, 0x97, 0xf3, 0xd0, 0x0d, 0xdb, 0x35, 0x40, 0xb7, 0xf8, 0x02, 0x36,
	0x2c, 0xe5, 0x6b, 0x78, 0x89, 0x7f, 0x7f, 0xea, 0x67, 0x1b, 0xb6, 0x4a, 0xc9, 0x8e, 0x74, 0x87,
	0x83, 0x11, 0xa3, 0x6f, 0xc9, 0x88, 0xa0, 0x2d, 0xf2, 0xe5, 0x2e, 0xda, 0x90, 0x93, 0x1e, 0xb9,
	0x79, 0x38, 0xd7, 0x68, 0x20, 0x82, 0xe1, 0x28, 0x58, 0x2e, 0xbc, 0x7e, 0x86, 0x57, 0x95, 0x64,
	0x3e, 0xee, 0x07, 0x9b, 0x97, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x91, 0x7a, 0x30, 0xac, 0x8b,
	0x01, 0x00, 0x00,
}
