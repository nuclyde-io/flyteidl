// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/schedule.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

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

type FixedRateUnit int32

const (
	FixedRateUnit_MINUTE FixedRateUnit = 0
	FixedRateUnit_HOUR   FixedRateUnit = 1
	FixedRateUnit_DAY    FixedRateUnit = 2
)

var FixedRateUnit_name = map[int32]string{
	0: "MINUTE",
	1: "HOUR",
	2: "DAY",
}
var FixedRateUnit_value = map[string]int32{
	"MINUTE": 0,
	"HOUR":   1,
	"DAY":    2,
}

func (x FixedRateUnit) String() string {
	return proto.EnumName(FixedRateUnit_name, int32(x))
}
func (FixedRateUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schedule_7021bfb4f8197528, []int{0}
}

type FixedRate struct {
	Value                uint32        `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 FixedRateUnit `protobuf:"varint,2,opt,name=unit,proto3,enum=flyteidl.admin.FixedRateUnit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FixedRate) Reset()         { *m = FixedRate{} }
func (m *FixedRate) String() string { return proto.CompactTextString(m) }
func (*FixedRate) ProtoMessage()    {}
func (*FixedRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_7021bfb4f8197528, []int{0}
}
func (m *FixedRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedRate.Unmarshal(m, b)
}
func (m *FixedRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedRate.Marshal(b, m, deterministic)
}
func (dst *FixedRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedRate.Merge(dst, src)
}
func (m *FixedRate) XXX_Size() int {
	return xxx_messageInfo_FixedRate.Size(m)
}
func (m *FixedRate) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedRate.DiscardUnknown(m)
}

var xxx_messageInfo_FixedRate proto.InternalMessageInfo

func (m *FixedRate) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *FixedRate) GetUnit() FixedRateUnit {
	if m != nil {
		return m.Unit
	}
	return FixedRateUnit_MINUTE
}

type Schedule struct {
	// Types that are valid to be assigned to ScheduleExpression:
	//	*Schedule_CronExpression
	//	*Schedule_Rate
	ScheduleExpression   isSchedule_ScheduleExpression `protobuf_oneof:"ScheduleExpression"`
	KickoffTimeInputArg  string                        `protobuf:"bytes,3,opt,name=kickoff_time_input_arg,json=kickoffTimeInputArg,proto3" json:"kickoff_time_input_arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_7021bfb4f8197528, []int{1}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (dst *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(dst, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

type isSchedule_ScheduleExpression interface {
	isSchedule_ScheduleExpression()
}

type Schedule_CronExpression struct {
	CronExpression string `protobuf:"bytes,1,opt,name=cron_expression,json=cronExpression,proto3,oneof"`
}

type Schedule_Rate struct {
	Rate *FixedRate `protobuf:"bytes,2,opt,name=rate,proto3,oneof"`
}

func (*Schedule_CronExpression) isSchedule_ScheduleExpression() {}

func (*Schedule_Rate) isSchedule_ScheduleExpression() {}

func (m *Schedule) GetScheduleExpression() isSchedule_ScheduleExpression {
	if m != nil {
		return m.ScheduleExpression
	}
	return nil
}

func (m *Schedule) GetCronExpression() string {
	if x, ok := m.GetScheduleExpression().(*Schedule_CronExpression); ok {
		return x.CronExpression
	}
	return ""
}

func (m *Schedule) GetRate() *FixedRate {
	if x, ok := m.GetScheduleExpression().(*Schedule_Rate); ok {
		return x.Rate
	}
	return nil
}

func (m *Schedule) GetKickoffTimeInputArg() string {
	if m != nil {
		return m.KickoffTimeInputArg
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Schedule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Schedule_OneofMarshaler, _Schedule_OneofUnmarshaler, _Schedule_OneofSizer, []interface{}{
		(*Schedule_CronExpression)(nil),
		(*Schedule_Rate)(nil),
	}
}

func _Schedule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Schedule)
	// ScheduleExpression
	switch x := m.ScheduleExpression.(type) {
	case *Schedule_CronExpression:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.CronExpression)
	case *Schedule_Rate:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Schedule.ScheduleExpression has unexpected type %T", x)
	}
	return nil
}

func _Schedule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Schedule)
	switch tag {
	case 1: // ScheduleExpression.cron_expression
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ScheduleExpression = &Schedule_CronExpression{x}
		return true, err
	case 2: // ScheduleExpression.rate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FixedRate)
		err := b.DecodeMessage(msg)
		m.ScheduleExpression = &Schedule_Rate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Schedule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Schedule)
	// ScheduleExpression
	switch x := m.ScheduleExpression.(type) {
	case *Schedule_CronExpression:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CronExpression)))
		n += len(x.CronExpression)
	case *Schedule_Rate:
		s := proto.Size(x.Rate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FixedRate)(nil), "flyteidl.admin.FixedRate")
	proto.RegisterType((*Schedule)(nil), "flyteidl.admin.Schedule")
	proto.RegisterEnum("flyteidl.admin.FixedRateUnit", FixedRateUnit_name, FixedRateUnit_value)
}

func init() {
	proto.RegisterFile("flyteidl/admin/schedule.proto", fileDescriptor_schedule_7021bfb4f8197528)
}

var fileDescriptor_schedule_7021bfb4f8197528 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0x86, 0x19, 0x20, 0xc2, 0x67, 0x40, 0x52, 0x89, 0xc1, 0x03, 0x09, 0xe1, 0x84, 0x26, 0xb6,
	0x01, 0x7e, 0x01, 0x44, 0x0c, 0x1c, 0xd4, 0xa4, 0xc2, 0x41, 0x2f, 0xcb, 0xd8, 0xbe, 0x8d, 0x86,
	0xad, 0x5d, 0x4a, 0x67, 0xe0, 0x67, 0xf9, 0x0f, 0x0d, 0x15, 0x30, 0x3b, 0x78, 0x6c, 0x9f, 0xaf,
	0xef, 0xd3, 0xf6, 0x85, 0x4e, 0x18, 0xef, 0x0d, 0x8a, 0x20, 0x66, 0x5e, 0x90, 0x08, 0xc9, 0xb6,
	0xfe, 0x1a, 0x83, 0x2c, 0x46, 0x9a, 0x6a, 0x65, 0x14, 0x69, 0x9c, 0x30, 0xb5, 0xb8, 0xb7, 0x80,
	0xda, 0xb3, 0xd8, 0x61, 0xc0, 0x3d, 0x83, 0xa4, 0x05, 0x17, 0x5f, 0x5e, 0x9c, 0x61, 0xdb, 0xe9,
	0x3a, 0xfd, 0x3a, 0xff, 0x5d, 0x90, 0x01, 0x94, 0x33, 0x29, 0x4c, 0xbb, 0xd8, 0x75, 0xfa, 0x8d,
	0x61, 0x87, 0xe6, 0x13, 0xe8, 0xf9, 0xf8, 0x52, 0x0a, 0xc3, 0xed, 0x68, 0xef, 0xdb, 0x81, 0xea,
	0xfb, 0x51, 0x4c, 0xee, 0xe1, 0xda, 0xd7, 0x4a, 0xba, 0xb8, 0x4b, 0x35, 0x6e, 0xb7, 0x42, 0x49,
	0x9b, 0x5f, 0x9b, 0x15, 0x78, 0xe3, 0x00, 0xa6, 0xe7, 0x7d, 0xc2, 0xa0, 0xac, 0x3d, 0x83, 0x56,
	0x75, 0x35, 0xbc, 0xfb, 0x57, 0x35, 0x2b, 0x70, 0x3b, 0x48, 0x46, 0x70, 0xbb, 0x11, 0xfe, 0x46,
	0x85, 0xa1, 0x6b, 0x44, 0x82, 0xae, 0x90, 0x69, 0x66, 0x5c, 0x4f, 0x47, 0xed, 0xd2, 0x41, 0xc1,
	0x6f, 0x8e, 0x74, 0x21, 0x12, 0x9c, 0x1f, 0xd8, 0x58, 0x47, 0x93, 0x16, 0x90, 0xd3, 0xe5, 0xfe,
	0xdc, 0x0f, 0x14, 0xea, 0xb9, 0xa7, 0x10, 0x80, 0xca, 0xcb, 0xfc, 0x75, 0xb9, 0x98, 0x36, 0x0b,
	0xa4, 0x0a, 0xe5, 0xd9, 0xdb, 0x92, 0x37, 0x1d, 0x72, 0x09, 0xa5, 0xa7, 0xf1, 0x47, 0xb3, 0x38,
	0x19, 0x7d, 0x0e, 0x22, 0x61, 0xd6, 0xd9, 0x8a, 0xfa, 0x2a, 0x61, 0xf1, 0x3e, 0x34, 0xec, 0xfc,
	0xf5, 0x11, 0x4a, 0x96, 0xae, 0x1e, 0x23, 0xc5, 0xf2, 0x6d, 0xac, 0x2a, 0xb6, 0x85, 0xd1, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x31, 0x4f, 0x38, 0xa6, 0x01, 0x00, 0x00,
}
