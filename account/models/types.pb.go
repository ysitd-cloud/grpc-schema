// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/models/types.proto

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	account/models/types.proto

It has these top-level messages:
	User
*/
package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Username    string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	AvatarUrl   string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "account.User")
}

func init() { proto.RegisterFile("account/models/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xbd, 0x4a, 0xc6, 0x30,
	0x18, 0x85, 0xa9, 0xd6, 0x9f, 0xbe, 0x3a, 0x05, 0x87, 0x50, 0x10, 0xd4, 0xc9, 0xa5, 0x09, 0xe2,
	0x1d, 0x78, 0x01, 0x2a, 0x42, 0x17, 0x97, 0xf2, 0x36, 0x09, 0x6d, 0x20, 0x69, 0x42, 0x7e, 0x84,
	0x7a, 0xf5, 0xd2, 0xb4, 0x08, 0xdf, 0x78, 0xce, 0xf3, 0x0c, 0xe7, 0x40, 0x8b, 0x42, 0xb8, 0xbc,
	0x24, 0x6e, 0x9d, 0x54, 0x26, 0xf2, 0xb4, 0x7a, 0x15, 0x99, 0x0f, 0x2e, 0x39, 0x72, 0x75, 0xb0,
	0xa7, 0x5f, 0xa8, 0xfb, 0xa8, 0x02, 0x69, 0xe1, 0x3a, 0x47, 0x15, 0x16, 0xb4, 0x8a, 0x56, 0x0f,
	0xd5, 0x73, 0xf3, 0xf5, 0x9f, 0xc9, 0x23, 0xdc, 0x4a, 0x1d, 0xbd, 0xc1, 0x75, 0x28, 0xfc, 0xac,
	0xf0, 0x9b, 0xa3, 0x7b, 0xdf, 0x94, 0x3b, 0xb8, 0x50, 0x16, 0xb5, 0xa1, 0xe7, 0x85, 0xed, 0x81,
	0xdc, 0x03, 0xe0, 0x0f, 0x26, 0x0c, 0x43, 0x0e, 0x86, 0xd6, 0x05, 0x35, 0x7b, 0xd3, 0x07, 0xf3,
	0xf6, 0x01, 0x54, 0x18, 0x97, 0x25, 0x5b, 0xa3, 0x4e, 0x92, 0x4d, 0xc1, 0x0b, 0x76, 0xec, 0xfa,
	0xac, 0xbe, 0x5f, 0x26, 0x9d, 0xe6, 0x3c, 0x32, 0xe1, 0x2c, 0x2f, 0x42, 0x57, 0x64, 0xbe, 0x69,
	0x5d, 0x14, 0xb3, 0xb2, 0xc8, 0x4f, 0x1f, 0x8e, 0x97, 0xe5, 0xdc, 0xeb, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x27, 0x26, 0x6a, 0x10, 0xfa, 0x00, 0x00, 0x00,
}