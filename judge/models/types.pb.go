// Code generated by protoc-gen-go. DO NOT EDIT.
// source: judge/models/types.proto

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	judge/models/types.proto

It has these top-level messages:
	Subject
	SubjectSelector
	UpdateSubjectRequest
	SubjectMutationReply
	Resource
	ResourceSelector
	ResourceMutationReply
	UpdateResourceRequest
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

type SubjectType int32

const (
	SubjectType_APP  SubjectType = 0
	SubjectType_USER SubjectType = 1
)

var SubjectType_name = map[int32]string{
	0: "APP",
	1: "USER",
}
var SubjectType_value = map[string]int32{
	"APP":  0,
	"USER": 1,
}

func (x SubjectType) String() string {
	return proto.EnumName(SubjectType_name, int32(x))
}
func (SubjectType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Subject struct {
	Id   string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type SubjectType `protobuf:"varint,2,opt,name=type,enum=judge.SubjectType" json:"type,omitempty"`
}

func (m *Subject) Reset()                    { *m = Subject{} }
func (m *Subject) String() string            { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()               {}
func (*Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Subject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subject) GetType() SubjectType {
	if m != nil {
		return m.Type
	}
	return SubjectType_APP
}

type SubjectSelector struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SubjectSelector) Reset()                    { *m = SubjectSelector{} }
func (m *SubjectSelector) String() string            { return proto.CompactTextString(m) }
func (*SubjectSelector) ProtoMessage()               {}
func (*SubjectSelector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubjectSelector) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateSubjectRequest struct {
	Selector *SubjectSelector `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	Content  *Subject         `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *UpdateSubjectRequest) Reset()                    { *m = UpdateSubjectRequest{} }
func (m *UpdateSubjectRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSubjectRequest) ProtoMessage()               {}
func (*UpdateSubjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateSubjectRequest) GetSelector() *SubjectSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpdateSubjectRequest) GetContent() *Subject {
	if m != nil {
		return m.Content
	}
	return nil
}

type SubjectMutationReply struct {
	Error   bool `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	Message bool `protobuf:"varint,2,opt,name=message" json:"message,omitempty"`
}

func (m *SubjectMutationReply) Reset()                    { *m = SubjectMutationReply{} }
func (m *SubjectMutationReply) String() string            { return proto.CompactTextString(m) }
func (*SubjectMutationReply) ProtoMessage()               {}
func (*SubjectMutationReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubjectMutationReply) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *SubjectMutationReply) GetMessage() bool {
	if m != nil {
		return m.Message
	}
	return false
}

type Resource struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Owner       string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Resource) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ResourceSelector struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ResourceSelector) Reset()                    { *m = ResourceSelector{} }
func (m *ResourceSelector) String() string            { return proto.CompactTextString(m) }
func (*ResourceSelector) ProtoMessage()               {}
func (*ResourceSelector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ResourceSelector) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ResourceMutationReply struct {
	Error   bool `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	Message bool `protobuf:"varint,2,opt,name=message" json:"message,omitempty"`
}

func (m *ResourceMutationReply) Reset()                    { *m = ResourceMutationReply{} }
func (m *ResourceMutationReply) String() string            { return proto.CompactTextString(m) }
func (*ResourceMutationReply) ProtoMessage()               {}
func (*ResourceMutationReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResourceMutationReply) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *ResourceMutationReply) GetMessage() bool {
	if m != nil {
		return m.Message
	}
	return false
}

type UpdateResourceRequest struct {
	Selector *ResourceSelector `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	Content  *Resource         `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *UpdateResourceRequest) Reset()                    { *m = UpdateResourceRequest{} }
func (m *UpdateResourceRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateResourceRequest) ProtoMessage()               {}
func (*UpdateResourceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateResourceRequest) GetSelector() *ResourceSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpdateResourceRequest) GetContent() *Resource {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Subject)(nil), "judge.Subject")
	proto.RegisterType((*SubjectSelector)(nil), "judge.SubjectSelector")
	proto.RegisterType((*UpdateSubjectRequest)(nil), "judge.UpdateSubjectRequest")
	proto.RegisterType((*SubjectMutationReply)(nil), "judge.SubjectMutationReply")
	proto.RegisterType((*Resource)(nil), "judge.Resource")
	proto.RegisterType((*ResourceSelector)(nil), "judge.ResourceSelector")
	proto.RegisterType((*ResourceMutationReply)(nil), "judge.ResourceMutationReply")
	proto.RegisterType((*UpdateResourceRequest)(nil), "judge.UpdateResourceRequest")
	proto.RegisterEnum("judge.SubjectType", SubjectType_name, SubjectType_value)
}

func init() { proto.RegisterFile("judge/models/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcd, 0x8e, 0x9b, 0x30,
	0x10, 0x2e, 0xd9, 0xdd, 0xc2, 0x0e, 0x52, 0x36, 0xb2, 0xb2, 0x5b, 0x8e, 0x94, 0x43, 0x95, 0x56,
	0x5a, 0x90, 0xc8, 0x13, 0xa4, 0x52, 0xdb, 0x43, 0x55, 0x29, 0x72, 0x9a, 0x4b, 0x6f, 0xc4, 0x1e,
	0x11, 0x22, 0xc0, 0xd4, 0x36, 0x8a, 0x78, 0xfb, 0x0a, 0x03, 0x11, 0x21, 0xea, 0x69, 0x8f, 0x33,
	0xdf, 0xcf, 0xd8, 0xdf, 0x0c, 0x78, 0xa7, 0x9a, 0xa7, 0x18, 0x15, 0x82, 0x63, 0xae, 0x22, 0xdd,
	0x54, 0xa8, 0xc2, 0x4a, 0x0a, 0x2d, 0xc8, 0x83, 0x41, 0x82, 0x0d, 0xd8, 0xbb, 0xfa, 0x70, 0x42,
	0xa6, 0xc9, 0x1c, 0x66, 0x19, 0xf7, 0x2c, 0xdf, 0x5a, 0x3d, 0xd2, 0x59, 0xc6, 0xc9, 0x27, 0xb8,
	0x6f, 0x05, 0xde, 0xcc, 0xb7, 0x56, 0xf3, 0x98, 0x84, 0x46, 0x10, 0xf6, 0xec, 0xdf, 0x4d, 0x85,
	0xd4, 0xe0, 0xc1, 0x47, 0x78, 0xea, 0x9b, 0x3b, 0xcc, 0x91, 0x69, 0x21, 0xa7, 0x56, 0x81, 0x86,
	0xe5, 0xbe, 0xe2, 0x89, 0xc6, 0x9e, 0x48, 0xf1, 0x6f, 0x8d, 0x4a, 0x93, 0x18, 0x1c, 0xd5, 0x6b,
	0x0c, 0xdb, 0x8d, 0x5f, 0xae, 0xc7, 0x0c, 0x8e, 0xf4, 0xc2, 0x23, 0x2b, 0xb0, 0x99, 0x28, 0x35,
	0x96, 0xda, 0xbc, 0xcc, 0x8d, 0xe7, 0xd7, 0x12, 0x3a, 0xc0, 0xc1, 0x77, 0x58, 0xf6, 0xbd, 0x5f,
	0xb5, 0x4e, 0x74, 0x26, 0x4a, 0x8a, 0x55, 0xde, 0x90, 0x25, 0x3c, 0xa0, 0x94, 0xfd, 0x48, 0x87,
	0x76, 0x05, 0xf1, 0xc0, 0x2e, 0x50, 0xa9, 0x24, 0xed, 0x7e, 0xec, 0xd0, 0xa1, 0x0c, 0x28, 0x38,
	0x14, 0x95, 0xa8, 0x25, 0xc3, 0x9b, 0x90, 0x7c, 0x70, 0x39, 0x2a, 0x26, 0xb3, 0xaa, 0xf5, 0x37,
	0xca, 0x47, 0x3a, 0x6e, 0xb5, 0xd3, 0xc4, 0xb9, 0x44, 0xe9, 0xdd, 0x19, 0xac, 0x2b, 0x82, 0x00,
	0x16, 0x83, 0xe7, 0x7f, 0x53, 0xfb, 0x01, 0xcf, 0x03, 0xe7, 0x6d, 0x1f, 0x38, 0xc3, 0x73, 0x17,
	0xff, 0x60, 0x37, 0xe4, 0xbf, 0xbe, 0xc9, 0xff, 0x43, 0x1f, 0xe6, 0xf4, 0x71, 0xa3, 0x05, 0x7c,
	0x9e, 0x2e, 0xe0, 0x69, 0xa2, 0xb9, 0x6c, 0xe0, 0x8b, 0x0f, 0xee, 0xe8, 0x5e, 0x88, 0x0d, 0x77,
	0x9b, 0xed, 0x76, 0xf1, 0x8e, 0x38, 0x70, 0xbf, 0xdf, 0x7d, 0xa3, 0x0b, 0xeb, 0xeb, 0x4f, 0x78,
	0x61, 0xb9, 0xa8, 0x79, 0xd8, 0xa8, 0x4c, 0xf3, 0x30, 0x95, 0x15, 0xeb, 0xdc, 0xb6, 0xd6, 0x9f,
	0x28, 0xcd, 0xf4, 0xb1, 0x3e, 0x84, 0x4c, 0x14, 0x91, 0x81, 0x5f, 0x0d, 0x35, 0x6a, 0x49, 0xaf,
	0x8a, 0x1d, 0xb1, 0x48, 0xa2, 0xf1, 0x7d, 0x1f, 0xde, 0x9b, 0xd3, 0x5e, 0xff, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0xc4, 0xc7, 0xbf, 0xf6, 0x02, 0x00, 0x00,
}
