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
	Environment     []string `protobuf:"bytes,1,rep,name=environment" json:"environment,omitempty"`
	PlanPath        string   `protobuf:"bytes,2,opt,name=plan_path,json=planPath" json:"plan_path,omitempty"`
	RepoUrl         string   `protobuf:"bytes,3,opt,name=repo_url,json=repoUrl" json:"repo_url,omitempty"`
	Channel         string   `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
	EnableLogStream bool     `protobuf:"varint,5,opt,name=enable_log_stream,json=enableLogStream" json:"enable_log_stream,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Build) GetEnvironment() []string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *Build) GetPlanPath() string {
	if m != nil {
		return m.PlanPath
	}
	return ""
}

func (m *Build) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *Build) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Build) GetEnableLogStream() bool {
	if m != nil {
		return m.EnableLogStream
	}
	return false
}

func init() {
	proto.RegisterType((*Build)(nil), "hurtlocker.agent.build.Build")
}

func init() { proto.RegisterFile("components/agent/proto/build/build.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x15, 0x4a, 0x69, 0x63, 0x06, 0x84, 0x07, 0x64, 0xc4, 0x12, 0x31, 0x45, 0x0c, 0xf1,
	0xc0, 0x07, 0x20, 0x3a, 0x33, 0xa0, 0x20, 0x16, 0x96, 0xc8, 0x71, 0x4f, 0x71, 0x54, 0xe7, 0xce,
	0x72, 0xce, 0xfc, 0x12, 0xbf, 0x89, 0xe2, 0x20, 0xc1, 0xc4, 0x62, 0xf9, 0xbd, 0x77, 0xcb, 0x9d,
	0xa8, 0x2d, 0x4d, 0x81, 0x10, 0x90, 0x67, 0x6d, 0x06, 0x40, 0xd6, 0x21, 0x12, 0x93, 0xee, 0xd3,
	0xe8, 0x8f, 0xeb, 0xdb, 0x64, 0x23, 0x6f, 0x5c, 0x8a, 0xec, 0xc9, 0x9e, 0x20, 0x36, 0x79, 0xb2,
	0xc9, 0xf5, 0xfe, 0xab, 0x10, 0xdb, 0xc3, 0xf2, 0x93, 0x95, 0xb8, 0x04, 0xfc, 0x1c, 0x23, 0xe1,
	0x04, 0xc8, 0xaa, 0xa8, 0x36, 0x75, 0xd9, 0xfe, 0x55, 0xf2, 0x4e, 0x94, 0xc1, 0x1b, 0xec, 0x82,
	0x61, 0xa7, 0xce, 0xaa, 0xa2, 0x2e, 0xdb, 0xfd, 0x22, 0x5e, 0x0d, 0x3b, 0x79, 0x2b, 0xf6, 0x11,
	0x02, 0x75, 0x29, 0x7a, 0xb5, 0xc9, 0x6d, 0xb7, 0xf0, 0x7b, 0xf4, 0x52, 0x89, 0x9d, 0x75, 0x06,
	0x11, 0xbc, 0x3a, 0x5f, 0xcb, 0x0f, 0xca, 0x07, 0x71, 0x0d, 0x68, 0x7a, 0x0f, 0x9d, 0xa7, 0xa1,
	0x9b, 0x39, 0x82, 0x99, 0xd4, 0xb6, 0x2a, 0xea, 0x7d, 0x7b, 0xb5, 0x86, 0x17, 0x1a, 0xde, 0xb2,
	0x3e, 0x3c, 0x7f, 0x3c, 0x0d, 0x23, 0xbb, 0xd4, 0x37, 0x96, 0x26, 0x6d, 0x5d, 0xb2, 0x27, 0x0f,
	0x0e, 0xcc, 0x71, 0xd6, 0xbf, 0xbb, 0xe9, 0xff, 0x0e, 0xd2, 0x5f, 0x64, 0x78, 0xfc, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x79, 0x04, 0x4d, 0x37, 0x01, 0x00, 0x00,
}
