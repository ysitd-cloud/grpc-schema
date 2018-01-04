// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deployer/models/types.proto

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	deployer/models/types.proto

It has these top-level messages:
	Application
	Deployment
	EnvironmentPair
	Network
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

type Application struct {
	Id          string             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Owner       string             `protobuf:"bytes,2,opt,name=owner" json:"owner,omitempty"`
	Name        string             `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Deployment  *Deployment        `protobuf:"bytes,4,opt,name=deployment" json:"deployment,omitempty"`
	Network     *Network           `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
	Environment []*EnvironmentPair `protobuf:"bytes,6,rep,name=environment" json:"environment,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (m *Application) String() string            { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Application) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Application) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *Application) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Application) GetEnvironment() []*EnvironmentPair {
	if m != nil {
		return m.Environment
	}
	return nil
}

type Deployment struct {
	Image string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Tag   string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
}

func (m *Deployment) Reset()                    { *m = Deployment{} }
func (m *Deployment) String() string            { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()               {}
func (*Deployment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Deployment) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Deployment) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type EnvironmentPair struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *EnvironmentPair) Reset()                    { *m = EnvironmentPair{} }
func (m *EnvironmentPair) String() string            { return proto.CompactTextString(m) }
func (*EnvironmentPair) ProtoMessage()               {}
func (*EnvironmentPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EnvironmentPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EnvironmentPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Network struct {
	Domain string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Network) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func init() {
	proto.RegisterType((*Application)(nil), "deployer.Application")
	proto.RegisterType((*Deployment)(nil), "deployer.Deployment")
	proto.RegisterType((*EnvironmentPair)(nil), "deployer.EnvironmentPair")
	proto.RegisterType((*Network)(nil), "deployer.Network")
}

func init() { proto.RegisterFile("deployer/models/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0x4d, 0x6b, 0xeb, 0x30,
	0x10, 0xc4, 0xf9, 0x7c, 0x6f, 0x0d, 0xef, 0x43, 0x84, 0xe2, 0xd2, 0x4b, 0xea, 0x53, 0xa0, 0x44,
	0x86, 0x34, 0x97, 0xd2, 0x53, 0x4b, 0x7b, 0x2d, 0xc1, 0xc7, 0xde, 0x14, 0x69, 0x71, 0x44, 0x2c,
	0xc9, 0xc8, 0x72, 0x82, 0x7f, 0x71, 0xff, 0x46, 0xb1, 0x6c, 0xd7, 0x21, 0xb7, 0x9d, 0x9d, 0x99,
	0x5d, 0x8d, 0x16, 0xee, 0x04, 0x16, 0xb9, 0xa9, 0xd1, 0x26, 0xca, 0x08, 0xcc, 0xcb, 0xc4, 0xd5,
	0x05, 0x96, 0xb4, 0xb0, 0xc6, 0x19, 0xf2, 0xab, 0x27, 0xe3, 0xaf, 0x00, 0xc2, 0x97, 0xa2, 0xc8,
	0x25, 0x67, 0x4e, 0x1a, 0x4d, 0xfe, 0xc0, 0x48, 0x8a, 0x28, 0x58, 0x06, 0xab, 0xdf, 0xe9, 0x48,
	0x0a, 0xb2, 0x80, 0xa9, 0x39, 0x6b, 0xb4, 0xd1, 0xc8, 0xb7, 0x5a, 0x40, 0x08, 0x4c, 0x34, 0x53,
	0x18, 0x8d, 0x7d, 0xd3, 0xd7, 0x64, 0x0b, 0xd0, 0x4e, 0x55, 0xa8, 0x5d, 0x34, 0x59, 0x06, 0xab,
	0x70, 0xb3, 0xa0, 0xfd, 0x22, 0xfa, 0xf6, 0xc3, 0xa5, 0x17, 0x3a, 0xf2, 0x00, 0x73, 0x8d, 0xee,
	0x6c, 0xec, 0x31, 0x9a, 0x7a, 0xcb, 0xff, 0xc1, 0xf2, 0xd1, 0x12, 0x69, 0xaf, 0x20, 0xcf, 0x10,
	0xa2, 0x3e, 0x49, 0x6b, 0xb4, 0xdf, 0x31, 0x5b, 0x8e, 0x57, 0xe1, 0xe6, 0x76, 0x30, 0xbc, 0x0f,
	0xe4, 0x8e, 0x49, 0x9b, 0x5e, 0xaa, 0xe3, 0x2d, 0xc0, 0xf0, 0x86, 0x26, 0x97, 0x54, 0x2c, 0xc3,
	0x2e, 0x6a, 0x0b, 0xc8, 0x3f, 0x18, 0x3b, 0x96, 0x75, 0x59, 0x9b, 0x32, 0x7e, 0x82, 0xbf, 0x57,
	0x53, 0x1b, 0xd1, 0x11, 0xeb, 0xce, 0xd8, 0x94, 0xcd, 0xb0, 0x13, 0xcb, 0x2b, 0xec, 0x3f, 0xc9,
	0x83, 0xf8, 0x1e, 0xe6, 0x5d, 0x02, 0x72, 0x03, 0x33, 0x61, 0x14, 0x93, 0xba, 0x73, 0x75, 0xe8,
	0x75, 0x07, 0x11, 0xcf, 0x4d, 0x25, 0x68, 0x5d, 0x4a, 0x27, 0x68, 0x66, 0x0b, 0x4e, 0x19, 0xe7,
	0xa6, 0xd2, 0x6e, 0x17, 0x7c, 0x6e, 0x32, 0xe9, 0x0e, 0xd5, 0x9e, 0x72, 0xa3, 0x12, 0x2f, 0x58,
	0x7b, 0x71, 0xd2, 0xc8, 0xd6, 0x25, 0x3f, 0xa0, 0x62, 0xc9, 0xd5, 0x95, 0xf7, 0x33, 0x7f, 0xe0,
	0xc7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xd2, 0x15, 0x5c, 0xff, 0x01, 0x00, 0x00,
}