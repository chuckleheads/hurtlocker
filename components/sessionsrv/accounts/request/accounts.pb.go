// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sessionsrv/accounts/request/accounts.proto

/*
Package request is a generated protocol buffer package.

It is generated from these files:
	components/sessionsrv/accounts/request/accounts.proto

It has these top-level messages:
	Account
	GetAccountReq
	CreateAccountReq
	UpdateAccountReq
*/
package request

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

type Account struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type GetAccountReq struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *GetAccountReq) Reset()                    { *m = GetAccountReq{} }
func (m *GetAccountReq) String() string            { return proto.CompactTextString(m) }
func (*GetAccountReq) ProtoMessage()               {}
func (*GetAccountReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetAccountReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CreateAccountReq struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *CreateAccountReq) Reset()                    { *m = CreateAccountReq{} }
func (m *CreateAccountReq) String() string            { return proto.CompactTextString(m) }
func (*CreateAccountReq) ProtoMessage()               {}
func (*CreateAccountReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateAccountReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateAccountReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UpdateAccountReq struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *UpdateAccountReq) Reset()                    { *m = UpdateAccountReq{} }
func (m *UpdateAccountReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccountReq) ProtoMessage()               {}
func (*UpdateAccountReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateAccountReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateAccountReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "hurtlocker.sessionsrv.accounts.request.Account")
	proto.RegisterType((*GetAccountReq)(nil), "hurtlocker.sessionsrv.accounts.request.GetAccountReq")
	proto.RegisterType((*CreateAccountReq)(nil), "hurtlocker.sessionsrv.accounts.request.CreateAccountReq")
	proto.RegisterType((*UpdateAccountReq)(nil), "hurtlocker.sessionsrv.accounts.request.UpdateAccountReq")
}

func init() {
	proto.RegisterFile("components/sessionsrv/accounts/request/accounts.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x2d, 0xfe, 0x1b, 0x50, 0xca, 0xe2, 0x61, 0xf1, 0x54, 0xf6, 0x20, 0x05, 0x21,
	0x39, 0x88, 0x1f, 0x40, 0x2d, 0x88, 0xf4, 0xb6, 0xe0, 0xc5, 0x5b, 0x9a, 0x0c, 0x26, 0x74, 0x93,
	0xd9, 0x66, 0x12, 0x3f, 0xbf, 0xd8, 0x96, 0x5d, 0x7b, 0x5b, 0xf0, 0x38, 0x2f, 0xef, 0xfd, 0x1e,
	0xe4, 0xc1, 0x93, 0x26, 0xdf, 0x53, 0xc0, 0x90, 0x58, 0x32, 0x32, 0x3b, 0x0a, 0x1c, 0xbf, 0xa5,
	0xd2, 0x9a, 0xf2, 0xaf, 0x16, 0x71, 0x97, 0x91, 0xd3, 0x20, 0x88, 0x3e, 0x52, 0xa2, 0xea, 0xde,
	0xe6, 0x98, 0x3a, 0xd2, 0x5b, 0x8c, 0x62, 0x8c, 0x89, 0xc1, 0x75, 0x8c, 0x35, 0x6b, 0xb8, 0x78,
	0x3e, 0x68, 0xd5, 0x0d, 0x94, 0xce, 0xd4, 0xc5, 0xa2, 0x58, 0xce, 0xda, 0xd2, 0x99, 0xea, 0x0e,
	0x2e, 0x33, 0x63, 0x0c, 0xca, 0x63, 0x5d, 0x2e, 0x8a, 0xe5, 0x55, 0x3b, 0xdc, 0xd5, 0x2d, 0x9c,
	0xa1, 0x57, 0xae, 0xab, 0x67, 0xfb, 0x87, 0xc3, 0xd1, 0x3c, 0xc0, 0xf5, 0x1b, 0xa6, 0x23, 0xaf,
	0xc5, 0xdd, 0x09, 0xa2, 0x38, 0x45, 0x34, 0x2b, 0x98, 0xbf, 0x46, 0x54, 0x09, 0xa7, 0xf9, 0xc7,
	0xca, 0xf2, 0x6f, 0xe5, 0x0a, 0xe6, 0x1f, 0xbd, 0xf9, 0x27, 0xe5, 0x65, 0xfd, 0xf9, 0xfe, 0xe5,
	0x92, 0xcd, 0x1b, 0xa1, 0xc9, 0x4b, 0x6d, 0xb3, 0xde, 0x76, 0x68, 0x51, 0x19, 0x96, 0xe3, 0x3f,
	0xca, 0x69, 0x4b, 0x6c, 0xce, 0xf7, 0x0b, 0x3c, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x27,
	0x78, 0x7b, 0xba, 0x01, 0x00, 0x00,
}
