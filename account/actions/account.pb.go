// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/actions/account.proto

/*
Package actions is a generated protocol buffer package.

It is generated from these files:
	account/actions/account.proto

It has these top-level messages:
	ValidateUserRequest
	ValidateUserReply
	GetUserInfoRequest
	GetUserInfoReply
	GetTokenInfoRequest
	GetTokenInfoReply
*/
package actions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import account "code.ysitd.cloud/grpc/schema/account/models"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ValidateUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *ValidateUserRequest) Reset()                    { *m = ValidateUserRequest{} }
func (m *ValidateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateUserRequest) ProtoMessage()               {}
func (*ValidateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ValidateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ValidateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ValidateUserReply struct {
	Valid bool          `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	User  *account.User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *ValidateUserReply) Reset()                    { *m = ValidateUserReply{} }
func (m *ValidateUserReply) String() string            { return proto.CompactTextString(m) }
func (*ValidateUserReply) ProtoMessage()               {}
func (*ValidateUserReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidateUserReply) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *ValidateUserReply) GetUser() *account.User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserInfoRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *GetUserInfoRequest) Reset()                    { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()               {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserInfoRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserInfoReply struct {
	Exists bool          `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
	User   *account.User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *GetUserInfoReply) Reset()                    { *m = GetUserInfoReply{} }
func (m *GetUserInfoReply) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoReply) ProtoMessage()               {}
func (*GetUserInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetUserInfoReply) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *GetUserInfoReply) GetUser() *account.User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetTokenInfoRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetTokenInfoRequest) Reset()                    { *m = GetTokenInfoRequest{} }
func (m *GetTokenInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTokenInfoRequest) ProtoMessage()               {}
func (*GetTokenInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetTokenInfoRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetTokenInfoReply struct {
	Exists bool           `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
	Token  *account.Token `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *GetTokenInfoReply) Reset()                    { *m = GetTokenInfoReply{} }
func (m *GetTokenInfoReply) String() string            { return proto.CompactTextString(m) }
func (*GetTokenInfoReply) ProtoMessage()               {}
func (*GetTokenInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetTokenInfoReply) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *GetTokenInfoReply) GetToken() *account.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidateUserRequest)(nil), "account.ValidateUserRequest")
	proto.RegisterType((*ValidateUserReply)(nil), "account.ValidateUserReply")
	proto.RegisterType((*GetUserInfoRequest)(nil), "account.GetUserInfoRequest")
	proto.RegisterType((*GetUserInfoReply)(nil), "account.GetUserInfoReply")
	proto.RegisterType((*GetTokenInfoRequest)(nil), "account.GetTokenInfoRequest")
	proto.RegisterType((*GetTokenInfoReply)(nil), "account.GetTokenInfoReply")
}

func init() { proto.RegisterFile("account/actions/account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xa9, 0xd8, 0x5a, 0x47, 0x14, 0x9b, 0x16, 0x29, 0x05, 0x41, 0x83, 0x07, 0x41, 0xc9,
	0x8a, 0xfe, 0x03, 0x2f, 0x45, 0xb1, 0xa0, 0x41, 0x3d, 0x78, 0x5b, 0x77, 0x47, 0x0d, 0x26, 0xd9,
	0x98, 0x99, 0xa8, 0xf9, 0xf7, 0xb2, 0x9b, 0x4d, 0x31, 0x1e, 0xa4, 0xc7, 0x6f, 0xe6, 0xcd, 0x7b,
	0x0f, 0x76, 0x61, 0x5f, 0x2a, 0x65, 0xaa, 0x9c, 0x85, 0x54, 0x9c, 0x98, 0x9c, 0x84, 0xe7, 0xa8,
	0x28, 0x0d, 0x9b, 0x60, 0xc3, 0xe3, 0x6c, 0xd6, 0xea, 0x32, 0xa3, 0x31, 0x25, 0xc1, 0x75, 0x81,
	0xd4, 0x88, 0xc2, 0x05, 0x8c, 0x1f, 0x65, 0x9a, 0x68, 0xc9, 0xf8, 0x40, 0x58, 0xc6, 0xf8, 0x51,
	0x21, 0x71, 0x30, 0x83, 0x61, 0x45, 0x58, 0xe6, 0x32, 0xc3, 0x69, 0xef, 0xa0, 0x77, 0xbc, 0x19,
	0x2f, 0xd9, 0xee, 0x0a, 0x49, 0xf4, 0x65, 0x4a, 0x3d, 0x5d, 0x6b, 0x76, 0x2d, 0x87, 0x37, 0x30,
	0xea, 0xda, 0x15, 0x69, 0x1d, 0x4c, 0xa0, 0xff, 0x69, 0x87, 0xce, 0x69, 0x18, 0x37, 0x10, 0x1c,
	0xc2, 0xba, 0xb5, 0x74, 0x16, 0x5b, 0xe7, 0xdb, 0x51, 0x5b, 0xde, 0xdd, 0xb9, 0x55, 0x78, 0x06,
	0xc1, 0x1c, 0xd9, 0x0e, 0xae, 0xf2, 0x17, 0xb3, 0x42, 0xb7, 0x70, 0x01, 0xbb, 0x9d, 0x0b, 0x1b,
	0xbf, 0x07, 0x03, 0xfc, 0x4e, 0x88, 0xc9, 0xe7, 0x7b, 0x5a, 0xa5, 0xc0, 0x09, 0x8c, 0xe7, 0xc8,
	0xf7, 0xe6, 0x1d, 0xf3, 0xdf, 0x0d, 0x26, 0xd0, 0x67, 0x3b, 0xf3, 0xf1, 0x0d, 0x84, 0x77, 0x30,
	0xea, 0x8a, 0xff, 0x0b, 0x3f, 0x6a, 0x2d, 0x9a, 0xf4, 0x9d, 0x65, 0xba, 0xbb, 0xf7, 0x96, 0x97,
	0xd7, 0x30, 0x55, 0xa9, 0xa9, 0x74, 0x54, 0x53, 0xc2, 0x3a, 0x7a, 0x2d, 0x0b, 0xd5, 0x0a, 0x6f,
	0x7b, 0x4f, 0xa7, 0xca, 0x68, 0xf4, 0x2b, 0x27, 0x13, 0x56, 0x20, 0x48, 0xbd, 0x61, 0x26, 0xc5,
	0x9f, 0xcf, 0xf1, 0x3c, 0x70, 0x0f, 0x7e, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xb7, 0x5d,
	0x3d, 0x36, 0x02, 0x00, 0x00,
}
