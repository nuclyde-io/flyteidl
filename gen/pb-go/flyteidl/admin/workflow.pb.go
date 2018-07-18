// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/workflow.proto

package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowCreateRequest struct {
	Id                   *Identifier   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version              string        `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Spec                 *WorkflowSpec `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WorkflowCreateRequest) Reset()         { *m = WorkflowCreateRequest{} }
func (m *WorkflowCreateRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowCreateRequest) ProtoMessage()    {}
func (*WorkflowCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_365b0efa67a90c64, []int{0}
}
func (m *WorkflowCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowCreateRequest.Unmarshal(m, b)
}
func (m *WorkflowCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowCreateRequest.Marshal(b, m, deterministic)
}
func (dst *WorkflowCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowCreateRequest.Merge(dst, src)
}
func (m *WorkflowCreateRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowCreateRequest.Size(m)
}
func (m *WorkflowCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowCreateRequest proto.InternalMessageInfo

func (m *WorkflowCreateRequest) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *WorkflowCreateRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *WorkflowCreateRequest) GetSpec() *WorkflowSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type WorkflowCreateResponse struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowCreateResponse) Reset()         { *m = WorkflowCreateResponse{} }
func (m *WorkflowCreateResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowCreateResponse) ProtoMessage()    {}
func (*WorkflowCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_365b0efa67a90c64, []int{1}
}
func (m *WorkflowCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowCreateResponse.Unmarshal(m, b)
}
func (m *WorkflowCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowCreateResponse.Marshal(b, m, deterministic)
}
func (dst *WorkflowCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowCreateResponse.Merge(dst, src)
}
func (m *WorkflowCreateResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowCreateResponse.Size(m)
}
func (m *WorkflowCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowCreateResponse proto.InternalMessageInfo

func (m *WorkflowCreateResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

type Workflow struct {
	Id                   *Identifier   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version              string        `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Urn                  string        `protobuf:"bytes,3,opt,name=urn" json:"urn,omitempty"`
	Spec                 *WorkflowSpec `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Workflow) Reset()         { *m = Workflow{} }
func (m *Workflow) String() string { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()    {}
func (*Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_365b0efa67a90c64, []int{2}
}
func (m *Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workflow.Unmarshal(m, b)
}
func (m *Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workflow.Marshal(b, m, deterministic)
}
func (dst *Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workflow.Merge(dst, src)
}
func (m *Workflow) XXX_Size() int {
	return xxx_messageInfo_Workflow.Size(m)
}
func (m *Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Workflow proto.InternalMessageInfo

func (m *Workflow) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Workflow) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Workflow) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *Workflow) GetSpec() *WorkflowSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type WorkflowList struct {
	Workflows            []*Workflow `protobuf:"bytes,1,rep,name=workflows" json:"workflows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WorkflowList) Reset()         { *m = WorkflowList{} }
func (m *WorkflowList) String() string { return proto.CompactTextString(m) }
func (*WorkflowList) ProtoMessage()    {}
func (*WorkflowList) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_365b0efa67a90c64, []int{3}
}
func (m *WorkflowList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowList.Unmarshal(m, b)
}
func (m *WorkflowList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowList.Marshal(b, m, deterministic)
}
func (dst *WorkflowList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowList.Merge(dst, src)
}
func (m *WorkflowList) XXX_Size() int {
	return xxx_messageInfo_WorkflowList.Size(m)
}
func (m *WorkflowList) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowList.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowList proto.InternalMessageInfo

func (m *WorkflowList) GetWorkflows() []*Workflow {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type WorkflowSpec struct {
	WorkflowTemplate     *core.WorkflowTemplate `protobuf:"bytes,1,opt,name=workflow_template,json=workflowTemplate" json:"workflow_template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WorkflowSpec) Reset()         { *m = WorkflowSpec{} }
func (m *WorkflowSpec) String() string { return proto.CompactTextString(m) }
func (*WorkflowSpec) ProtoMessage()    {}
func (*WorkflowSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_365b0efa67a90c64, []int{4}
}
func (m *WorkflowSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowSpec.Unmarshal(m, b)
}
func (m *WorkflowSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowSpec.Marshal(b, m, deterministic)
}
func (dst *WorkflowSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowSpec.Merge(dst, src)
}
func (m *WorkflowSpec) XXX_Size() int {
	return xxx_messageInfo_WorkflowSpec.Size(m)
}
func (m *WorkflowSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowSpec.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowSpec proto.InternalMessageInfo

func (m *WorkflowSpec) GetWorkflowTemplate() *core.WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkflowCreateRequest)(nil), "admin.WorkflowCreateRequest")
	proto.RegisterType((*WorkflowCreateResponse)(nil), "admin.WorkflowCreateResponse")
	proto.RegisterType((*Workflow)(nil), "admin.Workflow")
	proto.RegisterType((*WorkflowList)(nil), "admin.WorkflowList")
	proto.RegisterType((*WorkflowSpec)(nil), "admin.WorkflowSpec")
}

func init() {
	proto.RegisterFile("flyteidl/admin/workflow.proto", fileDescriptor_workflow_365b0efa67a90c64)
}

var fileDescriptor_workflow_365b0efa67a90c64 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0x49, 0xd2, 0xef, 0xab, 0x99, 0x0a, 0xb6, 0x2b, 0x96, 0x50, 0x15, 0x62, 0x2e, 0x06,
	0xc1, 0x14, 0xea, 0xd9, 0x8b, 0x3d, 0x09, 0x9e, 0xb6, 0x82, 0xe0, 0x45, 0x62, 0x32, 0x81, 0xc5,
	0x64, 0x37, 0xee, 0x6e, 0x0d, 0x1e, 0xbc, 0xf8, 0x97, 0x4b, 0xb6, 0xd9, 0x14, 0x72, 0xf1, 0xe0,
	0x2d, 0x99, 0xf9, 0xbd, 0x37, 0x6f, 0x77, 0x16, 0xce, 0x8b, 0xf2, 0x53, 0x23, 0xcb, 0xcb, 0x65,
	0x9a, 0x57, 0x8c, 0x2f, 0x1b, 0x21, 0xdf, 0x8a, 0x52, 0x34, 0x49, 0x2d, 0x85, 0x16, 0xe4, 0x9f,
	0xa9, 0x2e, 0xce, 0x7a, 0x2a, 0x13, 0x12, 0x07, 0xd0, 0xe2, 0x74, 0xe0, 0x91, 0x89, 0xaa, 0x12,
	0x7c, 0xd7, 0x8c, 0xbe, 0xe0, 0xe4, 0xa9, 0xc3, 0xd7, 0x12, 0x53, 0x8d, 0x14, 0xdf, 0xb7, 0xa8,
	0x34, 0xb9, 0x00, 0x97, 0xe5, 0x81, 0x13, 0x3a, 0xf1, 0x64, 0x35, 0x4b, 0x8c, 0x32, 0xb9, 0xcf,
	0x91, 0x6b, 0x56, 0x30, 0x94, 0xd4, 0x65, 0x39, 0x09, 0x60, 0xfc, 0x81, 0x52, 0x31, 0xc1, 0x03,
	0x37, 0x74, 0x62, 0x9f, 0xda, 0x5f, 0x72, 0x09, 0x23, 0x55, 0x63, 0x16, 0x78, 0x46, 0x7e, 0xdc,
	0xc9, 0xed, 0xa0, 0x4d, 0x8d, 0x19, 0x35, 0x40, 0x74, 0x05, 0xf3, 0xe1, 0x78, 0x55, 0x0b, 0xae,
	0x90, 0x4c, 0xc1, 0xdb, 0x4a, 0x6e, 0x02, 0xf8, 0xb4, 0xfd, 0x8c, 0xbe, 0x1d, 0x38, 0xb0, 0xf0,
	0xdf, 0xe2, 0x75, 0xde, 0x5e, 0xef, 0xdd, 0x07, 0x1e, 0xfd, 0x16, 0xf8, 0x16, 0x0e, 0x6d, 0xf5,
	0x81, 0x29, 0x4d, 0xae, 0xc1, 0xb7, 0xd7, 0xad, 0x02, 0x27, 0xf4, 0xe2, 0xc9, 0xea, 0x68, 0xa0,
	0xa6, 0x7b, 0x22, 0xda, 0xec, 0xe5, 0xad, 0x29, 0x59, 0xc3, 0xcc, 0x36, 0x5f, 0x34, 0x56, 0x75,
	0x99, 0x6a, 0xec, 0x4e, 0x35, 0x4f, 0xda, 0x65, 0xf6, 0x2e, 0x8f, 0x5d, 0x97, 0x4e, 0x9b, 0x41,
	0xe5, 0x6e, 0xfc, 0xbc, 0x7b, 0x07, 0xaf, 0xff, 0xcd, 0x4e, 0x6f, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x58, 0xed, 0x4e, 0xff, 0x36, 0x02, 0x00, 0x00,
}