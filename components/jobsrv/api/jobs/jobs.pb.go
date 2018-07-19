// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/jobsrv/api/jobs/jobs.proto

/*
Package jobs is a generated protocol buffer package.

It is generated from these files:
	components/jobsrv/api/jobs/jobs.proto

It has these top-level messages:
*/
package jobs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import hurtlocker_jobsrv_jobs_request "github.com/chuckleheads/hurtlocker/components/jobsrv/jobs/request"
import hurtlocker_jobsrv_jobs_response "github.com/chuckleheads/hurtlocker/components/jobsrv/jobs/response"

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

// Client API for Jobs service

type JobsClient interface {
	GetJob(ctx context.Context, in *hurtlocker_jobsrv_jobs_request.GetJobReq, opts ...grpc.CallOption) (*hurtlocker_jobsrv_jobs_response.JobResp, error)
	CreateJob(ctx context.Context, in *hurtlocker_jobsrv_jobs_request.CreateJobReq, opts ...grpc.CallOption) (*hurtlocker_jobsrv_jobs_response.JobResp, error)
}

type jobsClient struct {
	cc *grpc.ClientConn
}

func NewJobsClient(cc *grpc.ClientConn) JobsClient {
	return &jobsClient{cc}
}

func (c *jobsClient) GetJob(ctx context.Context, in *hurtlocker_jobsrv_jobs_request.GetJobReq, opts ...grpc.CallOption) (*hurtlocker_jobsrv_jobs_response.JobResp, error) {
	out := new(hurtlocker_jobsrv_jobs_response.JobResp)
	err := grpc.Invoke(ctx, "/hurtlocker.jobsrv.jobs.Jobs/GetJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) CreateJob(ctx context.Context, in *hurtlocker_jobsrv_jobs_request.CreateJobReq, opts ...grpc.CallOption) (*hurtlocker_jobsrv_jobs_response.JobResp, error) {
	out := new(hurtlocker_jobsrv_jobs_response.JobResp)
	err := grpc.Invoke(ctx, "/hurtlocker.jobsrv.jobs.Jobs/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jobs service

type JobsServer interface {
	GetJob(context.Context, *hurtlocker_jobsrv_jobs_request.GetJobReq) (*hurtlocker_jobsrv_jobs_response.JobResp, error)
	CreateJob(context.Context, *hurtlocker_jobsrv_jobs_request.CreateJobReq) (*hurtlocker_jobsrv_jobs_response.JobResp, error)
}

func RegisterJobsServer(s *grpc.Server, srv JobsServer) {
	s.RegisterService(&_Jobs_serviceDesc, srv)
}

func _Jobs_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hurtlocker_jobsrv_jobs_request.GetJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hurtlocker.jobsrv.jobs.Jobs/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).GetJob(ctx, req.(*hurtlocker_jobsrv_jobs_request.GetJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hurtlocker_jobsrv_jobs_request.CreateJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hurtlocker.jobsrv.jobs.Jobs/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).CreateJob(ctx, req.(*hurtlocker_jobsrv_jobs_request.CreateJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Jobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hurtlocker.jobsrv.jobs.Jobs",
	HandlerType: (*JobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _Jobs_GetJob_Handler,
		},
		{
			MethodName: "CreateJob",
			Handler:    _Jobs_CreateJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/jobsrv/api/jobs/jobs.proto",
}

func init() { proto.RegisterFile("components/jobsrv/api/jobs/jobs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0xcf, 0xca, 0x4f, 0x2a, 0x2e, 0x2a, 0xd3, 0x4f, 0x2c,
	0xc8, 0x04, 0x33, 0xc1, 0x84, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x58, 0x46, 0x69, 0x51,
	0x49, 0x4e, 0x7e, 0x72, 0x76, 0x6a, 0x91, 0x1e, 0x44, 0x19, 0x98, 0x92, 0x92, 0x49, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0x05, 0xeb, 0x49, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x83,
	0xea, 0x92, 0xd2, 0xc4, 0x34, 0x1c, 0x6c, 0x70, 0x51, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0x92,
	0x05, 0x52, 0x5a, 0x38, 0x95, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x22, 0xa9, 0x35, 0xfa, 0xc0,
	0xc8, 0xc5, 0xe2, 0x95, 0x9f, 0x54, 0x2c, 0x54, 0xc8, 0xc5, 0xe6, 0x9e, 0x5a, 0xe2, 0x95, 0x9f,
	0x24, 0xa4, 0xa9, 0x87, 0xdd, 0x81, 0x7a, 0x50, 0xab, 0xf4, 0x20, 0xea, 0x82, 0x52, 0x0b, 0xa5,
	0x34, 0x70, 0x2b, 0x85, 0x58, 0xa5, 0x07, 0x56, 0x58, 0x5c, 0xa0, 0x24, 0xd4, 0x74, 0xf9, 0xc9,
	0x64, 0x26, 0x1e, 0x21, 0x2e, 0x88, 0x53, 0xaa, 0x33, 0x53, 0x6a, 0x85, 0x4a, 0xb9, 0x38, 0x9d,
	0x8b, 0x52, 0x13, 0x4b, 0x52, 0x41, 0xb6, 0xea, 0x10, 0xb2, 0x15, 0xae, 0x94, 0x34, 0x8b, 0x05,
	0xc0, 0x16, 0x73, 0x29, 0xb1, 0x82, 0x2d, 0xb6, 0x62, 0xd4, 0x72, 0xb2, 0x8f, 0xb2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x28, 0x4d, 0xce, 0xce, 0x49,
	0xcd, 0x48, 0x4d, 0x4c, 0x29, 0xd6, 0x47, 0x18, 0xaa, 0x8f, 0x3b, 0x2e, 0x93, 0xd8, 0xc0, 0x41,
	0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x51, 0x8e, 0x0d, 0x25, 0xf0, 0x01, 0x00, 0x00,
}
