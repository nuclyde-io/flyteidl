// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/training_job.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type InputMode_Value int32

const (
	InputMode_FILE InputMode_Value = 0
	InputMode_PIPE InputMode_Value = 1
)

var InputMode_Value_name = map[int32]string{
	0: "FILE",
	1: "PIPE",
}

var InputMode_Value_value = map[string]int32{
	"FILE": 0,
	"PIPE": 1,
}

func (x InputMode_Value) String() string {
	return proto.EnumName(InputMode_Value_name, int32(x))
}

func (InputMode_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0, 0}
}

type AlgorithmName_Value int32

const (
	AlgorithmName_CUSTOM  AlgorithmName_Value = 0
	AlgorithmName_XGBOOST AlgorithmName_Value = 1
)

var AlgorithmName_Value_name = map[int32]string{
	0: "CUSTOM",
	1: "XGBOOST",
}

var AlgorithmName_Value_value = map[string]int32{
	"CUSTOM":  0,
	"XGBOOST": 1,
}

func (x AlgorithmName_Value) String() string {
	return proto.EnumName(AlgorithmName_Value_name, int32(x))
}

func (AlgorithmName_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1, 0}
}

type InputContentType_Value int32

const (
	InputContentType_TEXT_CSV InputContentType_Value = 0
)

var InputContentType_Value_name = map[int32]string{
	0: "TEXT_CSV",
}

var InputContentType_Value_value = map[string]int32{
	"TEXT_CSV": 0,
}

func (x InputContentType_Value) String() string {
	return proto.EnumName(InputContentType_Value_name, int32(x))
}

func (InputContentType_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{2, 0}
}

type DistributedProtocol_Value int32

const (
	// Use this value if the user wishes to use framework-native distributed training interfaces.
	// If this value is used, Flyte won't configure SageMaker to initialize unnecessary components such as
	// OpenMPI or Parameter Server.
	DistributedProtocol_UNSPECIFIED DistributedProtocol_Value = 0
	// Use this value if the user wishes to use MPI as the underlying protocol for her distributed training job
	// MPI is a framework-agnostic distributed protocol. It has multiple implementations. Currently, we have only
	// tested the OpenMPI implementation, which is the recommended implementation for Horovod.
	DistributedProtocol_MPI DistributedProtocol_Value = 1
)

var DistributedProtocol_Value_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "MPI",
}

var DistributedProtocol_Value_value = map[string]int32{
	"UNSPECIFIED": 0,
	"MPI":         1,
}

func (x DistributedProtocol_Value) String() string {
	return proto.EnumName(DistributedProtocol_Value_name, int32(x))
}

func (DistributedProtocol_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{5, 0}
}

// The input mode that the algorithm supports. When using the File input mode, SageMaker downloads
// the training data from S3 to the provisioned ML storage Volume, and mounts the directory to docker
// volume for training container. When using Pipe input mode, Amazon SageMaker streams data directly
// from S3 to the container.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
// For the input modes that different SageMaker algorithms support, see:
// https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html
type InputMode struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputMode) Reset()         { *m = InputMode{} }
func (m *InputMode) String() string { return proto.CompactTextString(m) }
func (*InputMode) ProtoMessage()    {}
func (*InputMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0}
}

func (m *InputMode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputMode.Unmarshal(m, b)
}
func (m *InputMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputMode.Marshal(b, m, deterministic)
}
func (m *InputMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputMode.Merge(m, src)
}
func (m *InputMode) XXX_Size() int {
	return xxx_messageInfo_InputMode.Size(m)
}
func (m *InputMode) XXX_DiscardUnknown() {
	xxx_messageInfo_InputMode.DiscardUnknown(m)
}

var xxx_messageInfo_InputMode proto.InternalMessageInfo

// The algorithm name is used for deciding which pre-built image to point to.
// This is only required for use cases where SageMaker's built-in algorithm mode is used.
// While we currently only support a subset of the algorithms, more will be added to the list.
// See: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
type AlgorithmName struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlgorithmName) Reset()         { *m = AlgorithmName{} }
func (m *AlgorithmName) String() string { return proto.CompactTextString(m) }
func (*AlgorithmName) ProtoMessage()    {}
func (*AlgorithmName) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1}
}

func (m *AlgorithmName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmName.Unmarshal(m, b)
}
func (m *AlgorithmName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmName.Marshal(b, m, deterministic)
}
func (m *AlgorithmName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmName.Merge(m, src)
}
func (m *AlgorithmName) XXX_Size() int {
	return xxx_messageInfo_AlgorithmName.Size(m)
}
func (m *AlgorithmName) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmName.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmName proto.InternalMessageInfo

// Specifies the type of file for input data. Different SageMaker built-in algorithms require different file types of input data
// See https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-training.html
// https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html
type InputContentType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputContentType) Reset()         { *m = InputContentType{} }
func (m *InputContentType) String() string { return proto.CompactTextString(m) }
func (*InputContentType) ProtoMessage()    {}
func (*InputContentType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{2}
}

func (m *InputContentType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputContentType.Unmarshal(m, b)
}
func (m *InputContentType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputContentType.Marshal(b, m, deterministic)
}
func (m *InputContentType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputContentType.Merge(m, src)
}
func (m *InputContentType) XXX_Size() int {
	return xxx_messageInfo_InputContentType.Size(m)
}
func (m *InputContentType) XXX_DiscardUnknown() {
	xxx_messageInfo_InputContentType.DiscardUnknown(m)
}

var xxx_messageInfo_InputContentType proto.InternalMessageInfo

// Specifies a metric that the training algorithm writes to stderr or stdout.
// This object is a pass-through.
// See this for details: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_MetricDefinition.html
type MetricDefinition struct {
	// User-defined name of the metric
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SageMaker hyperparameter tuning parses your algorithm’s stdout and stderr streams to find algorithm metrics
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricDefinition) Reset()         { *m = MetricDefinition{} }
func (m *MetricDefinition) String() string { return proto.CompactTextString(m) }
func (*MetricDefinition) ProtoMessage()    {}
func (*MetricDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{3}
}

func (m *MetricDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricDefinition.Unmarshal(m, b)
}
func (m *MetricDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricDefinition.Marshal(b, m, deterministic)
}
func (m *MetricDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricDefinition.Merge(m, src)
}
func (m *MetricDefinition) XXX_Size() int {
	return xxx_messageInfo_MetricDefinition.Size(m)
}
func (m *MetricDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_MetricDefinition proto.InternalMessageInfo

func (m *MetricDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricDefinition) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// Specifies the training algorithm to be used in the training job
// This object is mostly a pass-through, with a couple of exceptions include: (1) in Flyte, users don't need to specify
// TrainingImage; either use the built-in algorithm mode by using Flytekit's Simple Training Job and specifying an algorithm
// name and an algorithm version or (2) when users want to supply custom algorithms they should set algorithm_name field to
// CUSTOM. In this case, the value of the algorithm_version field has no effect
// For pass-through use cases: refer to this AWS official document for more details
// https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
type AlgorithmSpecification struct {
	// The input mode can be either PIPE or FILE
	InputMode InputMode_Value `protobuf:"varint,1,opt,name=input_mode,json=inputMode,proto3,enum=flyteidl.plugins.sagemaker.InputMode_Value" json:"input_mode,omitempty"`
	// The algorithm name is used for deciding which pre-built image to point to
	AlgorithmName AlgorithmName_Value `protobuf:"varint,2,opt,name=algorithm_name,json=algorithmName,proto3,enum=flyteidl.plugins.sagemaker.AlgorithmName_Value" json:"algorithm_name,omitempty"`
	// The algorithm version field is used for deciding which pre-built image to point to
	// This is only needed for use cases where SageMaker's built-in algorithm mode is chosen
	AlgorithmVersion string `protobuf:"bytes,3,opt,name=algorithm_version,json=algorithmVersion,proto3" json:"algorithm_version,omitempty"`
	// A list of metric definitions for SageMaker to evaluate/track on the progress of the training job
	// See this: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
	// and this: https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics.html
	MetricDefinitions []*MetricDefinition `protobuf:"bytes,4,rep,name=metric_definitions,json=metricDefinitions,proto3" json:"metric_definitions,omitempty"`
	// The content type of the input
	// See https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-training.html
	// https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html
	InputContentType     InputContentType_Value `protobuf:"varint,5,opt,name=input_content_type,json=inputContentType,proto3,enum=flyteidl.plugins.sagemaker.InputContentType_Value" json:"input_content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AlgorithmSpecification) Reset()         { *m = AlgorithmSpecification{} }
func (m *AlgorithmSpecification) String() string { return proto.CompactTextString(m) }
func (*AlgorithmSpecification) ProtoMessage()    {}
func (*AlgorithmSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{4}
}

func (m *AlgorithmSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmSpecification.Unmarshal(m, b)
}
func (m *AlgorithmSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmSpecification.Marshal(b, m, deterministic)
}
func (m *AlgorithmSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmSpecification.Merge(m, src)
}
func (m *AlgorithmSpecification) XXX_Size() int {
	return xxx_messageInfo_AlgorithmSpecification.Size(m)
}
func (m *AlgorithmSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmSpecification proto.InternalMessageInfo

func (m *AlgorithmSpecification) GetInputMode() InputMode_Value {
	if m != nil {
		return m.InputMode
	}
	return InputMode_FILE
}

func (m *AlgorithmSpecification) GetAlgorithmName() AlgorithmName_Value {
	if m != nil {
		return m.AlgorithmName
	}
	return AlgorithmName_CUSTOM
}

func (m *AlgorithmSpecification) GetAlgorithmVersion() string {
	if m != nil {
		return m.AlgorithmVersion
	}
	return ""
}

func (m *AlgorithmSpecification) GetMetricDefinitions() []*MetricDefinition {
	if m != nil {
		return m.MetricDefinitions
	}
	return nil
}

func (m *AlgorithmSpecification) GetInputContentType() InputContentType_Value {
	if m != nil {
		return m.InputContentType
	}
	return InputContentType_TEXT_CSV
}

// When enabling distributed training on a training job, the user should use this message to tell Flyte and SageMaker
// what kind of distributed protocol he/she wants to use to distribute the work.
type DistributedProtocol struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributedProtocol) Reset()         { *m = DistributedProtocol{} }
func (m *DistributedProtocol) String() string { return proto.CompactTextString(m) }
func (*DistributedProtocol) ProtoMessage()    {}
func (*DistributedProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{5}
}

func (m *DistributedProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedProtocol.Unmarshal(m, b)
}
func (m *DistributedProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedProtocol.Marshal(b, m, deterministic)
}
func (m *DistributedProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedProtocol.Merge(m, src)
}
func (m *DistributedProtocol) XXX_Size() int {
	return xxx_messageInfo_DistributedProtocol.Size(m)
}
func (m *DistributedProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedProtocol proto.InternalMessageInfo

// TrainingJobResourceConfig is a pass-through, specifying the instance type to use for the training job, the
// number of instances to launch, and the size of the ML storage volume the user wants to provision
// Refer to SageMaker official doc for more details: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
type TrainingJobResourceConfig struct {
	// The number of ML compute instances to use. For distributed training, provide a value greater than 1.
	InstanceCount int64 `protobuf:"varint,1,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
	// The ML compute instance type
	InstanceType string `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	// The size of the ML storage volume that you want to provision.
	VolumeSizeInGb int64 `protobuf:"varint,3,opt,name=volume_size_in_gb,json=volumeSizeInGb,proto3" json:"volume_size_in_gb,omitempty"`
	// When users specify an instance_count > 1, Flyte will try to configure SageMaker to enable distributed training.
	// If the users wish to use framework-agnostic distributed protocol such as MPI or Parameter Server, this
	// field should be set to the corresponding enum value
	DistributedProtocol  DistributedProtocol_Value `protobuf:"varint,4,opt,name=distributed_protocol,json=distributedProtocol,proto3,enum=flyteidl.plugins.sagemaker.DistributedProtocol_Value" json:"distributed_protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TrainingJobResourceConfig) Reset()         { *m = TrainingJobResourceConfig{} }
func (m *TrainingJobResourceConfig) String() string { return proto.CompactTextString(m) }
func (*TrainingJobResourceConfig) ProtoMessage()    {}
func (*TrainingJobResourceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{6}
}

func (m *TrainingJobResourceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJobResourceConfig.Unmarshal(m, b)
}
func (m *TrainingJobResourceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJobResourceConfig.Marshal(b, m, deterministic)
}
func (m *TrainingJobResourceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJobResourceConfig.Merge(m, src)
}
func (m *TrainingJobResourceConfig) XXX_Size() int {
	return xxx_messageInfo_TrainingJobResourceConfig.Size(m)
}
func (m *TrainingJobResourceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJobResourceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJobResourceConfig proto.InternalMessageInfo

func (m *TrainingJobResourceConfig) GetInstanceCount() int64 {
	if m != nil {
		return m.InstanceCount
	}
	return 0
}

func (m *TrainingJobResourceConfig) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *TrainingJobResourceConfig) GetVolumeSizeInGb() int64 {
	if m != nil {
		return m.VolumeSizeInGb
	}
	return 0
}

func (m *TrainingJobResourceConfig) GetDistributedProtocol() DistributedProtocol_Value {
	if m != nil {
		return m.DistributedProtocol
	}
	return DistributedProtocol_UNSPECIFIED
}

type CheckpointConfig struct {
	LocalPath            string   `protobuf:"bytes,1,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	S3Uri                string   `protobuf:"bytes,2,opt,name=s3_uri,json=s3Uri,proto3" json:"s3_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckpointConfig) Reset()         { *m = CheckpointConfig{} }
func (m *CheckpointConfig) String() string { return proto.CompactTextString(m) }
func (*CheckpointConfig) ProtoMessage()    {}
func (*CheckpointConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{7}
}

func (m *CheckpointConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckpointConfig.Unmarshal(m, b)
}
func (m *CheckpointConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckpointConfig.Marshal(b, m, deterministic)
}
func (m *CheckpointConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckpointConfig.Merge(m, src)
}
func (m *CheckpointConfig) XXX_Size() int {
	return xxx_messageInfo_CheckpointConfig.Size(m)
}
func (m *CheckpointConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckpointConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CheckpointConfig proto.InternalMessageInfo

func (m *CheckpointConfig) GetLocalPath() string {
	if m != nil {
		return m.LocalPath
	}
	return ""
}

func (m *CheckpointConfig) GetS3Uri() string {
	if m != nil {
		return m.S3Uri
	}
	return ""
}

// The spec of a training job. This is mostly a pass-through object
// https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
type TrainingJob struct {
	AlgorithmSpecification    *AlgorithmSpecification    `protobuf:"bytes,1,opt,name=algorithm_specification,json=algorithmSpecification,proto3" json:"algorithm_specification,omitempty"`
	TrainingJobResourceConfig *TrainingJobResourceConfig `protobuf:"bytes,2,opt,name=training_job_resource_config,json=trainingJobResourceConfig,proto3" json:"training_job_resource_config,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                   `json:"-"`
	XXX_unrecognized          []byte                     `json:"-"`
	XXX_sizecache             int32                      `json:"-"`
}

func (m *TrainingJob) Reset()         { *m = TrainingJob{} }
func (m *TrainingJob) String() string { return proto.CompactTextString(m) }
func (*TrainingJob) ProtoMessage()    {}
func (*TrainingJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{8}
}

func (m *TrainingJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJob.Unmarshal(m, b)
}
func (m *TrainingJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJob.Marshal(b, m, deterministic)
}
func (m *TrainingJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJob.Merge(m, src)
}
func (m *TrainingJob) XXX_Size() int {
	return xxx_messageInfo_TrainingJob.Size(m)
}
func (m *TrainingJob) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJob.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJob proto.InternalMessageInfo

func (m *TrainingJob) GetAlgorithmSpecification() *AlgorithmSpecification {
	if m != nil {
		return m.AlgorithmSpecification
	}
	return nil
}

func (m *TrainingJob) GetTrainingJobResourceConfig() *TrainingJobResourceConfig {
	if m != nil {
		return m.TrainingJobResourceConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.sagemaker.InputMode_Value", InputMode_Value_name, InputMode_Value_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.AlgorithmName_Value", AlgorithmName_Value_name, AlgorithmName_Value_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.InputContentType_Value", InputContentType_Value_name, InputContentType_Value_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.DistributedProtocol_Value", DistributedProtocol_Value_name, DistributedProtocol_Value_value)
	proto.RegisterType((*InputMode)(nil), "flyteidl.plugins.sagemaker.InputMode")
	proto.RegisterType((*AlgorithmName)(nil), "flyteidl.plugins.sagemaker.AlgorithmName")
	proto.RegisterType((*InputContentType)(nil), "flyteidl.plugins.sagemaker.InputContentType")
	proto.RegisterType((*MetricDefinition)(nil), "flyteidl.plugins.sagemaker.MetricDefinition")
	proto.RegisterType((*AlgorithmSpecification)(nil), "flyteidl.plugins.sagemaker.AlgorithmSpecification")
	proto.RegisterType((*DistributedProtocol)(nil), "flyteidl.plugins.sagemaker.DistributedProtocol")
	proto.RegisterType((*TrainingJobResourceConfig)(nil), "flyteidl.plugins.sagemaker.TrainingJobResourceConfig")
	proto.RegisterType((*CheckpointConfig)(nil), "flyteidl.plugins.sagemaker.CheckpointConfig")
	proto.RegisterType((*TrainingJob)(nil), "flyteidl.plugins.sagemaker.TrainingJob")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/sagemaker/training_job.proto", fileDescriptor_6a68f64d8fd9fe30)
}

var fileDescriptor_6a68f64d8fd9fe30 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcb, 0x4e, 0xdb, 0x4c,
	0x14, 0x4e, 0x48, 0xb8, 0xe4, 0x04, 0xf2, 0x9b, 0xe1, 0xf2, 0x07, 0x7a, 0x11, 0x75, 0x55, 0x09,
	0x44, 0x89, 0x55, 0x22, 0xa4, 0x2e, 0xba, 0x29, 0x21, 0xd0, 0xa0, 0x06, 0x22, 0x27, 0x44, 0xa8,
	0x5d, 0xb8, 0x63, 0x7b, 0xe2, 0x4c, 0xb1, 0x67, 0x2c, 0x7b, 0x8c, 0x1a, 0x9e, 0xa8, 0xcf, 0xd7,
	0x75, 0x17, 0x55, 0xc6, 0x17, 0xd2, 0x28, 0xa4, 0x3b, 0xe7, 0x9b, 0x33, 0x27, 0xdf, 0xe5, 0x9c,
	0x81, 0xa3, 0x81, 0x3b, 0x12, 0x84, 0xda, 0xae, 0xe6, 0xbb, 0x91, 0x43, 0x59, 0xa8, 0x85, 0xd8,
	0x21, 0x1e, 0xbe, 0x23, 0x81, 0x26, 0x02, 0x4c, 0x19, 0x65, 0x8e, 0xf1, 0x9d, 0x9b, 0x35, 0x3f,
	0xe0, 0x82, 0xa3, 0xdd, 0xb4, 0xbc, 0x96, 0x94, 0xd7, 0xb2, 0xf2, 0xdd, 0x97, 0x0e, 0xe7, 0x8e,
	0x4b, 0x34, 0x59, 0x69, 0x46, 0x03, 0xcd, 0x8e, 0x02, 0x2c, 0x28, 0x67, 0xf1, 0x5d, 0x75, 0x1f,
	0x4a, 0x2d, 0xe6, 0x47, 0xa2, 0xcd, 0x6d, 0xa2, 0x3e, 0x83, 0xc5, 0x3e, 0x76, 0x23, 0x82, 0x56,
	0xa0, 0x78, 0xde, 0xfa, 0xdc, 0x54, 0x72, 0xe3, 0xaf, 0x4e, 0xab, 0xd3, 0x54, 0xf2, 0xea, 0x3b,
	0x58, 0xfb, 0xe8, 0x3a, 0x3c, 0xa0, 0x62, 0xe8, 0x5d, 0x61, 0x8f, 0xa8, 0x7b, 0x69, 0x35, 0xc0,
	0x52, 0xe3, 0xa6, 0xdb, 0xbb, 0x6e, 0x2b, 0x39, 0x54, 0x86, 0xe5, 0xdb, 0x8b, 0xd3, 0xeb, 0xeb,
	0x6e, 0x4f, 0xc9, 0xab, 0x07, 0xa0, 0xc8, 0xe6, 0x0d, 0xce, 0x04, 0x61, 0xa2, 0x37, 0xf2, 0x89,
	0xba, 0x95, 0xde, 0x5a, 0x85, 0x95, 0x5e, 0xf3, 0xb6, 0x67, 0x34, 0xba, 0x7d, 0x25, 0xa7, 0x7e,
	0x00, 0xa5, 0x4d, 0x44, 0x40, 0xad, 0x33, 0x32, 0xa0, 0x8c, 0x8e, 0x19, 0x22, 0x04, 0x45, 0x86,
	0x3d, 0x52, 0xcd, 0xef, 0xe5, 0xf7, 0x4b, 0xba, 0xfc, 0x46, 0x9b, 0xb0, 0x18, 0x10, 0x87, 0xfc,
	0xa8, 0x2e, 0x48, 0x30, 0xfe, 0xa1, 0xfe, 0x2c, 0xc0, 0x76, 0x46, 0xae, 0xeb, 0x13, 0x8b, 0x0e,
	0xa8, 0x25, 0x65, 0xa2, 0x4b, 0x00, 0x3a, 0xe6, 0x60, 0x78, 0xdc, 0x8e, 0x5b, 0x55, 0x8e, 0x0f,
	0x6b, 0x4f, 0x3b, 0x56, 0xcb, 0xec, 0xa8, 0x49, 0x9e, 0x7a, 0x89, 0xa6, 0x00, 0xea, 0x43, 0x05,
	0xa7, 0xff, 0x62, 0x48, 0x6a, 0x0b, 0xb2, 0x9f, 0x36, 0xaf, 0xdf, 0x5f, 0xa6, 0x25, 0x3d, 0xd7,
	0xf0, 0x24, 0x88, 0x0e, 0x61, 0xfd, 0xb1, 0xef, 0x3d, 0x09, 0x42, 0xca, 0x59, 0xb5, 0x20, 0x05,
	0x2a, 0xd9, 0x41, 0x3f, 0xc6, 0xd1, 0x57, 0x40, 0x9e, 0x74, 0xca, 0xb0, 0x33, 0xab, 0xc2, 0x6a,
	0x71, 0xaf, 0xb0, 0x5f, 0x3e, 0x7e, 0x3b, 0x8f, 0xc8, 0xb4, 0xbf, 0xfa, 0xba, 0x37, 0x85, 0x84,
	0xe8, 0x1b, 0xa0, 0xd8, 0x2d, 0x2b, 0x8e, 0xcc, 0x10, 0x23, 0x9f, 0x54, 0x17, 0xa5, 0xca, 0xe3,
	0x7f, 0xba, 0x36, 0x91, 0x73, 0x22, 0x54, 0xa1, 0xd3, 0xf9, 0xbf, 0x87, 0x8d, 0x33, 0x1a, 0x8a,
	0x80, 0x9a, 0x91, 0x20, 0x76, 0x67, 0x3c, 0x84, 0x16, 0x77, 0xd5, 0x57, 0xe9, 0x58, 0xfc, 0x07,
	0xe5, 0x9b, 0xab, 0x6e, 0xa7, 0xd9, 0x68, 0x9d, 0xb7, 0x9a, 0x67, 0x4a, 0x0e, 0x2d, 0x43, 0xa1,
	0xdd, 0x69, 0x29, 0x79, 0xf5, 0x77, 0x1e, 0x76, 0x7a, 0xc9, 0xf4, 0x5f, 0x72, 0x53, 0x27, 0x21,
	0x8f, 0x02, 0x8b, 0x34, 0x38, 0x1b, 0x50, 0x07, 0xbd, 0x81, 0x0a, 0x65, 0xa1, 0xc0, 0xcc, 0x22,
	0x86, 0xc5, 0x23, 0x26, 0x64, 0xd6, 0x05, 0x7d, 0x2d, 0x45, 0x1b, 0x63, 0x10, 0xbd, 0x86, 0x0c,
	0x88, 0xb5, 0xc5, 0x73, 0xb4, 0x9a, 0x82, 0x63, 0x8e, 0xe8, 0x00, 0xd6, 0xef, 0xb9, 0x1b, 0x79,
	0xc4, 0x08, 0xe9, 0x03, 0x31, 0x28, 0x33, 0x1c, 0x53, 0xe6, 0x51, 0xd0, 0x2b, 0xf1, 0x41, 0x97,
	0x3e, 0x90, 0x16, 0xbb, 0x30, 0xd1, 0x10, 0x36, 0xed, 0x47, 0x39, 0x86, 0x9f, 0xe8, 0xa9, 0x16,
	0xa5, 0x65, 0x27, 0xf3, 0x2c, 0x9b, 0x61, 0x43, 0xe2, 0xda, 0x86, 0x3d, 0xc3, 0xa1, 0x4f, 0xa0,
	0x34, 0x86, 0xc4, 0xba, 0xf3, 0x39, 0x65, 0x22, 0x11, 0xfd, 0x02, 0xc0, 0xe5, 0x16, 0x76, 0x0d,
	0x1f, 0x8b, 0x61, 0xb2, 0x27, 0x25, 0x89, 0x74, 0xb0, 0x18, 0xa2, 0x2d, 0x58, 0x0a, 0xeb, 0x46,
	0x14, 0xd0, 0x74, 0x5b, 0xc2, 0xfa, 0x4d, 0x40, 0xd5, 0x5f, 0x79, 0x28, 0x4f, 0x18, 0x89, 0xee,
	0xe0, 0xff, 0xc7, 0xf1, 0x0b, 0x27, 0xb7, 0x47, 0xb6, 0x2c, 0xcf, 0x4f, 0x7e, 0xf6, 0xde, 0xe9,
	0xdb, 0x78, 0xf6, 0x3e, 0xde, 0xc3, 0xf3, 0xc9, 0x27, 0xcc, 0x08, 0x92, 0x18, 0xc7, 0x13, 0x37,
	0xa0, 0x8e, 0x64, 0x5a, 0x9e, 0x6f, 0xdc, 0x93, 0x43, 0xa0, 0xef, 0x88, 0xa7, 0x8e, 0x4e, 0x4f,
	0xbe, 0xd4, 0x1d, 0x2a, 0x86, 0x91, 0x59, 0xb3, 0xb8, 0xa7, 0xb9, 0xa3, 0x81, 0xd0, 0xb2, 0x57,
	0xd6, 0x21, 0x4c, 0xf3, 0xcd, 0x23, 0x87, 0x6b, 0xd3, 0x0f, 0xaf, 0xb9, 0x24, 0x13, 0xad, 0xff,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xa3, 0x20, 0xe7, 0x93, 0x05, 0x00, 0x00,
}
