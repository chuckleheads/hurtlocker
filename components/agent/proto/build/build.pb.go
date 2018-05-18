// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/agent/proto/build/build.proto

/*
Package build is a generated protocol buffer package.

It is generated from these files:
	components/agent/proto/build/build.proto

It has these top-level messages:
	Build
*/
package build

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

type Build struct {
	PackageName string `protobuf:"bytes,1,opt,name=packageName" json:"packageName,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Build) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func init() {
	proto.RegisterType((*Build)(nil), "hurtlocker.agent.build.Build")
}

func init() { proto.RegisterFile("components/agent/proto/build/build.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0xd7, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0x81, 0x90, 0x7a, 0x60, 0x11, 0x21, 0xb1, 0x8c,
	0xd2, 0xa2, 0x92, 0x9c, 0xfc, 0xe4, 0xec, 0xd4, 0x22, 0x3d, 0xb0, 0x4a, 0x3d, 0xb0, 0xac, 0x92,
	0x26, 0x17, 0xab, 0x13, 0x88, 0x21, 0xa4, 0xc0, 0xc5, 0x5d, 0x90, 0x98, 0x9c, 0x9d, 0x98, 0x9e,
	0xea, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2c, 0xe4, 0xe4, 0x18,
	0x65, 0x9f, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x51, 0x9a,
	0x9c, 0x9d, 0x93, 0x9a, 0x91, 0x9a, 0x98, 0x52, 0xac, 0x8f, 0x30, 0x5c, 0x1f, 0x9f, 0x8b, 0x92,
	0xd8, 0xc0, 0x1c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xd0, 0x43, 0x50, 0xb8, 0x00,
	0x00, 0x00,
}
