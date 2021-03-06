// Code generated by protoc-gen-go. DO NOT EDIT.
// source: carvedService.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	carvedService.proto

It has these top-level messages:
	Times
	Sum
*/
package rpc

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

type Times struct {
	Times int32 `protobuf:"varint,1,opt,name=times" json:"times,omitempty"`
}

func (m *Times) Reset()                    { *m = Times{} }
func (m *Times) String() string            { return proto.CompactTextString(m) }
func (*Times) ProtoMessage()               {}
func (*Times) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Times) GetTimes() int32 {
	if m != nil {
		return m.Times
	}
	return 0
}

type Sum struct {
	Sum int32 `protobuf:"varint,1,opt,name=Sum" json:"Sum,omitempty"`
}

func (m *Sum) Reset()                    { *m = Sum{} }
func (m *Sum) String() string            { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()               {}
func (*Sum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Sum) GetSum() int32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func init() {
	proto.RegisterType((*Times)(nil), "twirp.carved.Times")
	proto.RegisterType((*Sum)(nil), "twirp.carved.Sum")
}

func init() { proto.RegisterFile("carvedService.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4e, 0x2c, 0x2a,
	0x4b, 0x4d, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x29, 0x29, 0xcf, 0x2c, 0x2a, 0xd0, 0x83, 0x48, 0x29, 0xc9, 0x72, 0xb1, 0x86, 0x64, 0xe6,
	0xa6, 0x16, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x80, 0x18, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41,
	0x10, 0x8e, 0x92, 0x38, 0x17, 0x73, 0x70, 0x69, 0xae, 0x90, 0x00, 0x98, 0x82, 0x4a, 0x81, 0x98,
	0x46, 0x96, 0x5c, 0x6c, 0xce, 0x60, 0x13, 0x84, 0xf4, 0xb9, 0x58, 0x9d, 0xf3, 0x4b, 0xf3, 0x4a,
	0x84, 0x84, 0xf5, 0x90, 0x4d, 0xd6, 0x03, 0x1b, 0x2b, 0x25, 0x88, 0x2a, 0x18, 0x5c, 0x9a, 0xeb,
	0xc4, 0x1a, 0xc5, 0x5c, 0x54, 0x90, 0x9c, 0xc4, 0x06, 0x76, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x84, 0xca, 0xbe, 0xa5, 0x00, 0x00, 0x00,
}
