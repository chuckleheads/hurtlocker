// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/originsrv/origins/origins.proto

/*
Package origins is a generated protocol buffer package.

It is generated from these files:
	components/originsrv/origins/origins.proto

It has these top-level messages:
*/
package origins

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hurtlocker_originsrv_origins_request "github.com/chuckleheads/hurtlocker/components/originsrv/origins/request"
import hurtlocker_originsrv_origins_response "github.com/chuckleheads/hurtlocker/components/originsrv/origins/response"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Origins service

type OriginsClient interface {
	GetOrigin(ctx context.Context, in *hurtlocker_originsrv_origins_request.GetOriginReq, opts ...grpc.CallOption) (*hurtlocker_originsrv_origins_response.OriginResp, error)
	CreateOrigin(ctx context.Context, in *hurtlocker_originsrv_origins_request.CreateOriginReq, opts ...grpc.CallOption) (*hurtlocker_originsrv_origins_response.OriginResp, error)
	UpdateOrigin(ctx context.Context, in *hurtlocker_originsrv_origins_request.UpdateOriginReq, opts ...grpc.CallOption) (*hurtlocker_originsrv_origins_response.OriginResp, error)
}

type originsClient struct {
	cc *grpc.ClientConn
}

func NewOriginsClient(cc *grpc.ClientConn) OriginsClient {
	return &originsClient{cc}
}

func (c *originsClient) GetOrigin(ctx context.Context, in *hurtlocker_originsrv_origins_request.GetOriginReq, opts ...grpc.CallOption) (*hurtlocker_originsrv_origins_response.OriginResp, error) {
	out := new(hurtlocker_originsrv_origins_response.OriginResp)
	err := grpc.Invoke(ctx, "/hurtlocker.originsrv.origins.Origins/GetOrigin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originsClient) CreateOrigin(ctx context.Context, in *hurtlocker_originsrv_origins_request.CreateOriginReq, opts ...grpc.CallOption) (*hurtlocker_originsrv_origins_response.OriginResp, error) {
	out := new(hurtlocker_originsrv_origins_response.OriginResp)
	err := grpc.Invoke(ctx, "/hurtlocker.originsrv.origins.Origins/CreateOrigin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originsClient) UpdateOrigin(ctx context.Context, in *hurtlocker_originsrv_origins_request.UpdateOriginReq, opts ...grpc.CallOption) (*hurtlocker_originsrv_origins_response.OriginResp, error) {
	out := new(hurtlocker_originsrv_origins_response.OriginResp)
	err := grpc.Invoke(ctx, "/hurtlocker.originsrv.origins.Origins/UpdateOrigin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Origins service

type OriginsServer interface {
	GetOrigin(context.Context, *hurtlocker_originsrv_origins_request.GetOriginReq) (*hurtlocker_originsrv_origins_response.OriginResp, error)
	CreateOrigin(context.Context, *hurtlocker_originsrv_origins_request.CreateOriginReq) (*hurtlocker_originsrv_origins_response.OriginResp, error)
	UpdateOrigin(context.Context, *hurtlocker_originsrv_origins_request.UpdateOriginReq) (*hurtlocker_originsrv_origins_response.OriginResp, error)
}

func RegisterOriginsServer(s *grpc.Server, srv OriginsServer) {
	s.RegisterService(&_Origins_serviceDesc, srv)
}

func _Origins_GetOrigin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hurtlocker_originsrv_origins_request.GetOriginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginsServer).GetOrigin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hurtlocker.originsrv.origins.Origins/GetOrigin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginsServer).GetOrigin(ctx, req.(*hurtlocker_originsrv_origins_request.GetOriginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Origins_CreateOrigin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hurtlocker_originsrv_origins_request.CreateOriginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginsServer).CreateOrigin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hurtlocker.originsrv.origins.Origins/CreateOrigin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginsServer).CreateOrigin(ctx, req.(*hurtlocker_originsrv_origins_request.CreateOriginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Origins_UpdateOrigin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hurtlocker_originsrv_origins_request.UpdateOriginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginsServer).UpdateOrigin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hurtlocker.originsrv.origins.Origins/UpdateOrigin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginsServer).UpdateOrigin(ctx, req.(*hurtlocker_originsrv_origins_request.UpdateOriginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Origins_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hurtlocker.originsrv.origins.Origins",
	HandlerType: (*OriginsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrigin",
			Handler:    _Origins_GetOrigin_Handler,
		},
		{
			MethodName: "CreateOrigin",
			Handler:    _Origins_CreateOrigin_Handler,
		},
		{
			MethodName: "UpdateOrigin",
			Handler:    _Origins_UpdateOrigin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/originsrv/origins/origins.proto",
}

func init() { proto.RegisterFile("components/originsrv/origins/origins.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0xcf, 0x2f, 0xca, 0x4c, 0xcf, 0xcc, 0x2b, 0x2e, 0x2a,
	0x83, 0xb1, 0x60, 0xb4, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x4c, 0x46, 0x69, 0x51, 0x49,
	0x4e, 0x7e, 0x72, 0x76, 0x6a, 0x91, 0x1e, 0x5c, 0x2d, 0x8c, 0x25, 0x65, 0x84, 0xd7, 0xa4, 0xa2,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x54, 0x13, 0xa5, 0x8c, 0x09, 0xe8, 0x29, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0x45, 0xd5, 0x64, 0xf4, 0x91, 0x89, 0x8b, 0xdd, 0x1f, 0x22, 0x22, 0x54, 0xc4, 0xc5,
	0xe9, 0x9e, 0x5a, 0x02, 0xe1, 0x09, 0x19, 0xe9, 0xe1, 0x73, 0xa0, 0x1e, 0xd4, 0x09, 0x7a, 0x70,
	0x0d, 0x41, 0xa9, 0x85, 0x52, 0x86, 0x84, 0xf4, 0x40, 0x9c, 0xa0, 0x07, 0xd3, 0x51, 0x5c, 0x20,
	0x54, 0xc1, 0xc5, 0xe3, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x0a, 0xb5, 0xd6, 0x94, 0x38, 0x6b, 0x91,
	0xf5, 0x90, 0x6f, 0x73, 0x68, 0x41, 0x0a, 0xc9, 0x36, 0x23, 0xeb, 0x21, 0xcf, 0x66, 0x27, 0xc7,
	0x28, 0xfb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe4, 0x8c, 0xd2,
	0xe4, 0xec, 0x9c, 0xd4, 0x8c, 0xd4, 0xc4, 0x94, 0x62, 0x7d, 0x84, 0x59, 0xfa, 0xf8, 0x62, 0x33,
	0x89, 0x0d, 0x1c, 0x7b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x23, 0x84, 0xd2, 0x72,
	0x02, 0x00, 0x00,
}
