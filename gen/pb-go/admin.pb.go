// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.proto

package lyft_flyte_flyteadmin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import admin "flyteidl/admin"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error)
	CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error)
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateTask(ctx context.Context, in *admin.TaskCreateRequest, opts ...grpc.CallOption) (*admin.TaskCreateResponse, error) {
	out := new(admin.TaskCreateResponse)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateWorkflow(ctx context.Context, in *admin.WorkflowCreateRequest, opts ...grpc.CallOption) (*admin.WorkflowCreateResponse, error) {
	out := new(admin.WorkflowCreateResponse)
	err := c.cc.Invoke(ctx, "/lyft.flyte.flyteadmin.AdminService/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	CreateTask(context.Context, *admin.TaskCreateRequest) (*admin.TaskCreateResponse, error)
	CreateWorkflow(context.Context, *admin.WorkflowCreateRequest) (*admin.WorkflowCreateResponse, error)
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateTask(ctx, req.(*admin.TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.WorkflowCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lyft.flyte.flyteadmin.AdminService/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateWorkflow(ctx, req.(*admin.WorkflowCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lyft.flyte.flyteadmin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _AdminService_CreateTask_Handler,
		},
		{
			MethodName: "CreateWorkflow",
			Handler:    _AdminService_CreateWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor_admin_5b4bd88e80836183) }

var fileDescriptor_admin_5b4bd88e80836183 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0xc9, 0xcd,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcd, 0xa9, 0x4c, 0x2b, 0xd1, 0x4b, 0xcb,
	0xa9, 0x2c, 0x49, 0x85, 0x90, 0x60, 0x49, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd,
	0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88,
	0x26, 0x29, 0x49, 0xb0, 0xca, 0xcc, 0x94, 0x1c, 0x7d, 0xb0, 0x6a, 0xfd, 0x92, 0xc4, 0xe2, 0x6c,
	0xa8, 0x94, 0x2c, 0x9a, 0x54, 0x79, 0x7e, 0x51, 0x76, 0x5a, 0x4e, 0x7e, 0x39, 0x44, 0xda, 0xe8,
	0x06, 0x23, 0x17, 0x8f, 0x23, 0x48, 0x22, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x9a,
	0x8b, 0xcb, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0x24, 0xb1, 0x38, 0x5b, 0x48, 0x42, 0x0f, 0xe2,
	0x36, 0x10, 0x07, 0x22, 0x1c, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x25, 0x89, 0x45, 0xa6,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x49, 0xa2, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x42, 0x4a, 0xbc,
	0x60, 0xd7, 0x96, 0x19, 0x82, 0x9d, 0x53, 0x6c, 0xc5, 0xa8, 0x25, 0x94, 0xcd, 0xc5, 0x07, 0x51,
	0x1b, 0x0e, 0x75, 0x85, 0x90, 0x0c, 0xd4, 0x18, 0x98, 0x00, 0xaa, 0x25, 0xb2, 0x38, 0x64, 0xa1,
	0x16, 0xc9, 0x80, 0x2d, 0x12, 0x53, 0x12, 0x84, 0x59, 0x04, 0xf3, 0x1c, 0xc8, 0xb2, 0x24, 0x36,
	0xb0, 0x0f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x7e, 0xdb, 0x7c, 0x5f, 0x01, 0x00,
	0x00,
}