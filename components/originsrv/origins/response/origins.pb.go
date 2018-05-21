// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/originsrv/origins/response/origins.proto

/*
Package response is a generated protocol buffer package.

It is generated from these files:
	components/originsrv/origins/response/origins.proto

It has these top-level messages:
	Origin
	Origins
	OriginResp
	OriginsResp
*/
package response

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

type Origin struct {
	Id                       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	DefaultPackageVisibility string `protobuf:"bytes,3,opt,name=default_package_visibility,json=defaultPackageVisibility" json:"default_package_visibility,omitempty"`
}

func (m *Origin) Reset()                    { *m = Origin{} }
func (m *Origin) String() string            { return proto.CompactTextString(m) }
func (*Origin) ProtoMessage()               {}
func (*Origin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Origin) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Origin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Origin) GetDefaultPackageVisibility() string {
	if m != nil {
		return m.DefaultPackageVisibility
	}
	return ""
}

type Origins struct {
	Origins []*Origin `protobuf:"bytes,1,rep,name=origins" json:"origins,omitempty"`
}

func (m *Origins) Reset()                    { *m = Origins{} }
func (m *Origins) String() string            { return proto.CompactTextString(m) }
func (*Origins) ProtoMessage()               {}
func (*Origins) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Origins) GetOrigins() []*Origin {
	if m != nil {
		return m.Origins
	}
	return nil
}

type OriginResp struct {
	Origin *Origin `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
}

func (m *OriginResp) Reset()                    { *m = OriginResp{} }
func (m *OriginResp) String() string            { return proto.CompactTextString(m) }
func (*OriginResp) ProtoMessage()               {}
func (*OriginResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OriginResp) GetOrigin() *Origin {
	if m != nil {
		return m.Origin
	}
	return nil
}

type OriginsResp struct {
	Origins *Origins `protobuf:"bytes,1,opt,name=origins" json:"origins,omitempty"`
}

func (m *OriginsResp) Reset()                    { *m = OriginsResp{} }
func (m *OriginsResp) String() string            { return proto.CompactTextString(m) }
func (*OriginsResp) ProtoMessage()               {}
func (*OriginsResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OriginsResp) GetOrigins() *Origins {
	if m != nil {
		return m.Origins
	}
	return nil
}

func init() {
	proto.RegisterType((*Origin)(nil), "hurtlocker.originsrv.origins.response.Origin")
	proto.RegisterType((*Origins)(nil), "hurtlocker.originsrv.origins.response.Origins")
	proto.RegisterType((*OriginResp)(nil), "hurtlocker.originsrv.origins.response.OriginResp")
	proto.RegisterType((*OriginsResp)(nil), "hurtlocker.originsrv.origins.response.OriginsResp")
}

func init() {
	proto.RegisterFile("components/originsrv/origins/response/origins.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0x69, 0xfb, 0xa5, 0xcb, 0x77, 0x16, 0x3c, 0xe4, 0x54, 0x3c, 0x95, 0x82, 0xd0, 0x8b,
	0x29, 0xb8, 0x57, 0x4f, 0x82, 0xb8, 0x78, 0x51, 0x2a, 0x28, 0x78, 0x59, 0xda, 0x74, 0x6c, 0x63,
	0x7f, 0x24, 0x24, 0xe9, 0x82, 0xff, 0xbd, 0x98, 0x26, 0xd5, 0xe3, 0xee, 0x6d, 0x78, 0xf3, 0xde,
	0x27, 0x2f, 0x0c, 0xec, 0x98, 0x18, 0xa5, 0x98, 0x70, 0x32, 0xba, 0x10, 0x8a, 0xb7, 0x7c, 0xd2,
	0xea, 0xe8, 0xa7, 0x42, 0xa1, 0x96, 0x62, 0xd2, 0xe8, 0x05, 0x2a, 0x95, 0x30, 0x82, 0x5c, 0x75,
	0xb3, 0x32, 0x83, 0x60, 0x3d, 0x2a, 0xba, 0x86, 0xfc, 0x44, 0x7d, 0x28, 0xfb, 0x84, 0xf8, 0xc9,
	0x6a, 0xe4, 0x02, 0x42, 0xde, 0x24, 0x41, 0x1a, 0xe4, 0x51, 0x19, 0xf2, 0x86, 0x10, 0xf8, 0x37,
	0x55, 0x23, 0x26, 0x61, 0x1a, 0xe4, 0xff, 0x4b, 0x3b, 0x93, 0x5b, 0xb8, 0x6c, 0xf0, 0xa3, 0x9a,
	0x07, 0x73, 0x90, 0x15, 0xeb, 0xab, 0x16, 0x0f, 0x47, 0xae, 0x79, 0xcd, 0x07, 0x6e, 0xbe, 0x92,
	0xc8, 0x3a, 0x13, 0xe7, 0x78, 0x5e, 0x0c, 0xaf, 0xeb, 0x3e, 0x2b, 0x61, 0xb3, 0xbc, 0xa5, 0xc9,
	0x03, 0x6c, 0x5c, 0x95, 0x24, 0x48, 0xa3, 0x7c, 0x7b, 0x73, 0x4d, 0x4f, 0xea, 0x4b, 0x17, 0x40,
	0xe9, 0xd3, 0xd9, 0x0b, 0x80, 0x93, 0x50, 0x4b, 0x72, 0x0f, 0xf1, 0xb2, 0xb0, 0xff, 0x38, 0x9b,
	0xea, 0xc2, 0xd9, 0x1b, 0x6c, 0x5d, 0x51, 0x4b, 0xdd, 0xff, 0x2d, 0xfb, 0x83, 0xa5, 0x67, 0x61,
	0xf5, 0xda, 0xf6, 0xee, 0xf1, 0x7d, 0xdf, 0x72, 0xd3, 0xcd, 0x35, 0x65, 0x62, 0x2c, 0x58, 0x37,
	0xb3, 0x7e, 0xc0, 0x0e, 0xab, 0x46, 0x17, 0xbf, 0xc4, 0xe2, 0xa4, 0x73, 0xd7, 0xb1, 0xbd, 0xf3,
	0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x11, 0xe9, 0xd5, 0xd8, 0x1e, 0x02, 0x00, 0x00,
}
