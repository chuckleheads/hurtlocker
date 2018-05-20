// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sessionsrv/accounts/response/accounts.proto

/*
Package response is a generated protocol buffer package.

It is generated from these files:
	components/sessionsrv/accounts/response/accounts.proto

It has these top-level messages:
	Account
	Accounts
	AccountResp
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

// TODO - this shouldn't be duplicated with the request
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

type Accounts struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Accounts) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type AccountResp struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *AccountResp) Reset()                    { *m = AccountResp{} }
func (m *AccountResp) String() string            { return proto.CompactTextString(m) }
func (*AccountResp) ProtoMessage()               {}
func (*AccountResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AccountResp) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "hurtlocker.sessionsrv.accounts.response.Account")
	proto.RegisterType((*Accounts)(nil), "hurtlocker.sessionsrv.accounts.response.Accounts")
	proto.RegisterType((*AccountResp)(nil), "hurtlocker.sessionsrv.accounts.response.AccountResp")
}

func init() {
	proto.RegisterFile("components/sessionsrv/accounts/response/accounts.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x44, 0xd0, 0x70, 0x95, 0x18, 0x2c, 0x86, 0x88, 0x29, 0xca, 0x42, 0x26, 0x1b,
	0x81, 0xc4, 0x0e, 0x63, 0xe9, 0x94, 0x09, 0xd8, 0x52, 0xe7, 0x44, 0xac, 0x26, 0xbe, 0xc8, 0x17,
	0xf3, 0xfb, 0x11, 0x69, 0xec, 0xae, 0x15, 0xe3, 0x7d, 0xe7, 0xf7, 0xf9, 0xc9, 0x86, 0x17, 0x4d,
	0xe3, 0x44, 0x16, 0xed, 0xcc, 0x8a, 0x91, 0xd9, 0x90, 0x65, 0xf7, 0xa3, 0x5a, 0xad, 0xc9, 0xff,
	0x31, 0x87, 0x3c, 0x91, 0x65, 0x8c, 0x44, 0x4e, 0x8e, 0x66, 0x12, 0x0f, 0xbd, 0x77, 0xf3, 0x40,
	0xfa, 0x88, 0x4e, 0x9e, 0x73, 0x32, 0x9e, 0x0a, 0xb9, 0xea, 0x1d, 0x36, 0xaf, 0x27, 0x28, 0x6e,
	0x21, 0x35, 0x5d, 0x91, 0x94, 0x49, 0x9d, 0x35, 0xa9, 0xe9, 0xc4, 0x3d, 0xe4, 0x9e, 0xd1, 0xd9,
	0x76, 0xc4, 0x22, 0x2d, 0x93, 0xfa, 0xa6, 0x89, 0xb3, 0xb8, 0x83, 0x2b, 0x1c, 0x5b, 0x33, 0x14,
	0xd9, 0xb2, 0x38, 0x0d, 0xd5, 0x07, 0xe4, 0xab, 0x8c, 0xc5, 0x1e, 0xf2, 0x70, 0x5b, 0x91, 0x94,
	0x59, 0xbd, 0x7d, 0x7a, 0x94, 0x17, 0x96, 0x92, 0xab, 0xa4, 0x89, 0x86, 0xea, 0x13, 0xb6, 0x01,
	0x22, 0x4f, 0x62, 0x07, 0x9b, 0x75, 0xb5, 0xf4, 0xfd, 0x8f, 0x3b, 0x08, 0xde, 0xf6, 0x5f, 0xbb,
	0x6f, 0x33, 0xf7, 0xfe, 0x20, 0x35, 0x8d, 0x4a, 0xf7, 0x5e, 0x1f, 0x07, 0xec, 0xb1, 0xed, 0x58,
	0x9d, 0x9d, 0xea, 0xc2, 0x7f, 0x38, 0x5c, 0x2f, 0xef, 0xff, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x61, 0xb1, 0x5a, 0x91, 0xb9, 0x01, 0x00, 0x00,
}
