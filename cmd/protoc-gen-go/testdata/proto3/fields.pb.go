// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3/fields.proto

package proto3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldTestMessage_Enum int32

const (
	FieldTestMessage_ZERO FieldTestMessage_Enum = 0
)

var FieldTestMessage_Enum_name = map[int32]string{
	0: "ZERO",
}

var FieldTestMessage_Enum_value = map[string]int32{
	"ZERO": 0,
}

func (x FieldTestMessage_Enum) String() string {
	return proto.EnumName(FieldTestMessage_Enum_name, int32(x))
}

func (FieldTestMessage_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1e3ea068187307c, []int{0, 0}
}

type FieldTestMessage struct {
	OptionalBool         string                               `protobuf:"bytes,1,opt,name=optional_bool,json=optionalBool,proto3" json:"optional_bool,omitempty"`
	OptionalEnum         FieldTestMessage_Enum                `protobuf:"varint,2,opt,name=optional_enum,json=optionalEnum,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum" json:"optional_enum,omitempty"`
	OptionalInt32        int32                                `protobuf:"varint,3,opt,name=optional_int32,json=optionalInt32,proto3" json:"optional_int32,omitempty"`
	OptionalSint32       int32                                `protobuf:"zigzag32,4,opt,name=optional_sint32,json=optionalSint32,proto3" json:"optional_sint32,omitempty"`
	OptionalUint32       uint32                               `protobuf:"varint,5,opt,name=optional_uint32,json=optionalUint32,proto3" json:"optional_uint32,omitempty"`
	OptionalInt64        int64                                `protobuf:"varint,6,opt,name=optional_int64,json=optionalInt64,proto3" json:"optional_int64,omitempty"`
	OptionalSint64       int64                                `protobuf:"zigzag64,7,opt,name=optional_sint64,json=optionalSint64,proto3" json:"optional_sint64,omitempty"`
	OptionalUint64       uint64                               `protobuf:"varint,8,opt,name=optional_uint64,json=optionalUint64,proto3" json:"optional_uint64,omitempty"`
	OptionalSfixed32     int32                                `protobuf:"fixed32,9,opt,name=optional_sfixed32,json=optionalSfixed32,proto3" json:"optional_sfixed32,omitempty"`
	OptionalFixed32      uint32                               `protobuf:"fixed32,10,opt,name=optional_fixed32,json=optionalFixed32,proto3" json:"optional_fixed32,omitempty"`
	OptionalFloat        float32                              `protobuf:"fixed32,11,opt,name=optional_float,json=optionalFloat,proto3" json:"optional_float,omitempty"`
	OptionalSfixed64     int64                                `protobuf:"fixed64,12,opt,name=optional_sfixed64,json=optionalSfixed64,proto3" json:"optional_sfixed64,omitempty"`
	OptionalFixed64      uint64                               `protobuf:"fixed64,13,opt,name=optional_fixed64,json=optionalFixed64,proto3" json:"optional_fixed64,omitempty"`
	OptionalDouble       float64                              `protobuf:"fixed64,14,opt,name=optional_double,json=optionalDouble,proto3" json:"optional_double,omitempty"`
	OptionalString       string                               `protobuf:"bytes,15,opt,name=optional_string,json=optionalString,proto3" json:"optional_string,omitempty"`
	OptionalBytes        []byte                               `protobuf:"bytes,16,opt,name=optional_bytes,json=optionalBytes,proto3" json:"optional_bytes,omitempty"`
	Optional_Message     *FieldTestMessage_Message            `protobuf:"bytes,17,opt,name=optional_Message,json=optionalMessage,proto3" json:"optional_Message,omitempty"`
	RepeatedBool         []bool                               `protobuf:"varint,201,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedEnum         []FieldTestMessage_Enum              `protobuf:"varint,202,rep,packed,name=repeated_enum,json=repeatedEnum,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum" json:"repeated_enum,omitempty"`
	RepeatedInt32        []int32                              `protobuf:"varint,203,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedSint32       []int32                              `protobuf:"zigzag32,204,rep,packed,name=repeated_sint32,json=repeatedSint32,proto3" json:"repeated_sint32,omitempty"`
	RepeatedUint32       []uint32                             `protobuf:"varint,205,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedInt64        []int64                              `protobuf:"varint,206,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedSint64       []int64                              `protobuf:"zigzag64,207,rep,packed,name=repeated_sint64,json=repeatedSint64,proto3" json:"repeated_sint64,omitempty"`
	RepeatedUint64       []uint64                             `protobuf:"varint,208,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedSfixed32     []int32                              `protobuf:"fixed32,209,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedFixed32      []uint32                             `protobuf:"fixed32,210,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFloat        []float32                            `protobuf:"fixed32,211,rep,packed,name=repeated_float,json=repeatedFloat,proto3" json:"repeated_float,omitempty"`
	RepeatedSfixed64     []int64                              `protobuf:"fixed64,212,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedFixed64      []uint64                             `protobuf:"fixed64,213,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedDouble       []float64                            `protobuf:"fixed64,214,rep,packed,name=repeated_double,json=repeatedDouble,proto3" json:"repeated_double,omitempty"`
	RepeatedString       []string                             `protobuf:"bytes,215,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes        [][]byte                             `protobuf:"bytes,216,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	Repeated_Message     []*FieldTestMessage_Message          `protobuf:"bytes,217,rep,name=repeated_Message,json=repeatedMessage,proto3" json:"repeated_Message,omitempty"`
	MapInt32Int64        map[int32]int64                      `protobuf:"bytes,500,rep,name=map_int32_int64,json=mapInt32Int64,proto3" json:"map_int32_int64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStringMessage     map[string]*FieldTestMessage_Message `protobuf:"bytes,501,rep,name=map_string_message,json=mapStringMessage,proto3" json:"map_string_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapFixed64Enum       map[uint64]FieldTestMessage_Enum     `protobuf:"bytes,502,rep,name=map_fixed64_enum,json=mapFixed64Enum,proto3" json:"map_fixed64_enum,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *FieldTestMessage) Reset()         { *m = FieldTestMessage{} }
func (m *FieldTestMessage) String() string { return proto.CompactTextString(m) }
func (*FieldTestMessage) ProtoMessage()    {}
func (*FieldTestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1e3ea068187307c, []int{0}
}

func (m *FieldTestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldTestMessage.Unmarshal(m, b)
}
func (m *FieldTestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldTestMessage.Marshal(b, m, deterministic)
}
func (m *FieldTestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldTestMessage.Merge(m, src)
}
func (m *FieldTestMessage) XXX_Size() int {
	return xxx_messageInfo_FieldTestMessage.Size(m)
}
func (m *FieldTestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldTestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FieldTestMessage proto.InternalMessageInfo

func (m *FieldTestMessage) GetOptionalBool() string {
	if m != nil {
		return m.OptionalBool
	}
	return ""
}

func (m *FieldTestMessage) GetOptionalEnum() FieldTestMessage_Enum {
	if m != nil {
		return m.OptionalEnum
	}
	return FieldTestMessage_ZERO
}

func (m *FieldTestMessage) GetOptionalInt32() int32 {
	if m != nil {
		return m.OptionalInt32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSint32() int32 {
	if m != nil {
		return m.OptionalSint32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalUint32() uint32 {
	if m != nil {
		return m.OptionalUint32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalInt64() int64 {
	if m != nil {
		return m.OptionalInt64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSint64() int64 {
	if m != nil {
		return m.OptionalSint64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalUint64() uint64 {
	if m != nil {
		return m.OptionalUint64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSfixed32() int32 {
	if m != nil {
		return m.OptionalSfixed32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalFixed32() uint32 {
	if m != nil {
		return m.OptionalFixed32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalFloat() float32 {
	if m != nil {
		return m.OptionalFloat
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSfixed64() int64 {
	if m != nil {
		return m.OptionalSfixed64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalFixed64() uint64 {
	if m != nil {
		return m.OptionalFixed64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalDouble() float64 {
	if m != nil {
		return m.OptionalDouble
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalString() string {
	if m != nil {
		return m.OptionalString
	}
	return ""
}

func (m *FieldTestMessage) GetOptionalBytes() []byte {
	if m != nil {
		return m.OptionalBytes
	}
	return nil
}

func (m *FieldTestMessage) GetOptional_Message() *FieldTestMessage_Message {
	if m != nil {
		return m.Optional_Message
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedBool() []bool {
	if m != nil {
		return m.RepeatedBool
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedEnum() []FieldTestMessage_Enum {
	if m != nil {
		return m.RepeatedEnum
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSint32() []int32 {
	if m != nil {
		return m.RepeatedSint32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedUint32() []uint32 {
	if m != nil {
		return m.RepeatedUint32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedInt64() []int64 {
	if m != nil {
		return m.RepeatedInt64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSint64() []int64 {
	if m != nil {
		return m.RepeatedSint64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedUint64() []uint64 {
	if m != nil {
		return m.RepeatedUint64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSfixed32() []int32 {
	if m != nil {
		return m.RepeatedSfixed32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedFixed32() []uint32 {
	if m != nil {
		return m.RepeatedFixed32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedFloat() []float32 {
	if m != nil {
		return m.RepeatedFloat
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSfixed64() []int64 {
	if m != nil {
		return m.RepeatedSfixed64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedFixed64() []uint64 {
	if m != nil {
		return m.RepeatedFixed64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedDouble() []float64 {
	if m != nil {
		return m.RepeatedDouble
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedString() []string {
	if m != nil {
		return m.RepeatedString
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedBytes() [][]byte {
	if m != nil {
		return m.RepeatedBytes
	}
	return nil
}

func (m *FieldTestMessage) GetRepeated_Message() []*FieldTestMessage_Message {
	if m != nil {
		return m.Repeated_Message
	}
	return nil
}

func (m *FieldTestMessage) GetMapInt32Int64() map[int32]int64 {
	if m != nil {
		return m.MapInt32Int64
	}
	return nil
}

func (m *FieldTestMessage) GetMapStringMessage() map[string]*FieldTestMessage_Message {
	if m != nil {
		return m.MapStringMessage
	}
	return nil
}

func (m *FieldTestMessage) GetMapFixed64Enum() map[uint64]FieldTestMessage_Enum {
	if m != nil {
		return m.MapFixed64Enum
	}
	return nil
}

type FieldTestMessage_Message struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldTestMessage_Message) Reset()         { *m = FieldTestMessage_Message{} }
func (m *FieldTestMessage_Message) String() string { return proto.CompactTextString(m) }
func (*FieldTestMessage_Message) ProtoMessage()    {}
func (*FieldTestMessage_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1e3ea068187307c, []int{0, 3}
}

func (m *FieldTestMessage_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldTestMessage_Message.Unmarshal(m, b)
}
func (m *FieldTestMessage_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldTestMessage_Message.Marshal(b, m, deterministic)
}
func (m *FieldTestMessage_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldTestMessage_Message.Merge(m, src)
}
func (m *FieldTestMessage_Message) XXX_Size() int {
	return xxx_messageInfo_FieldTestMessage_Message.Size(m)
}
func (m *FieldTestMessage_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldTestMessage_Message.DiscardUnknown(m)
}

var xxx_messageInfo_FieldTestMessage_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FieldTestMessage)(nil), "goproto.protoc.proto3.FieldTestMessage")
	proto.RegisterMapType((map[uint64]FieldTestMessage_Enum)(nil), "goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry")
	proto.RegisterMapType((map[int32]int64)(nil), "goproto.protoc.proto3.FieldTestMessage.MapInt32Int64Entry")
	proto.RegisterMapType((map[string]*FieldTestMessage_Message)(nil), "goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry")
	proto.RegisterType((*FieldTestMessage_Message)(nil), "goproto.protoc.proto3.FieldTestMessage.Message")
	proto.RegisterEnum("goproto.protoc.proto3.FieldTestMessage_Enum", FieldTestMessage_Enum_name, FieldTestMessage_Enum_value)
}

func init() { proto.RegisterFile("proto3/fields.proto", fileDescriptor_f1e3ea068187307c) }

var fileDescriptor_f1e3ea068187307c = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xdb, 0x6f, 0xda, 0x48,
	0x14, 0xc6, 0x77, 0xb8, 0x86, 0x09, 0x17, 0x33, 0xd9, 0x48, 0xa3, 0x3c, 0x8d, 0xb2, 0xb7, 0xd9,
	0xdd, 0x04, 0x24, 0xb0, 0xac, 0xdd, 0x44, 0x2b, 0xad, 0xd0, 0x12, 0x69, 0x1f, 0xaa, 0xaa, 0x6e,
	0xfb, 0x92, 0x3e, 0x44, 0x06, 0x06, 0x0b, 0xd5, 0xf6, 0x20, 0x30, 0x55, 0xf3, 0x27, 0xf6, 0x9e,
	0xde, 0xdb, 0xff, 0xa3, 0xed, 0x73, 0x35, 0x9e, 0xb1, 0x3d, 0x36, 0x3c, 0x84, 0x3c, 0xd9, 0x3e,
	0x7c, 0x7c, 0xbf, 0x39, 0xc7, 0xe7, 0x03, 0xb8, 0x37, 0x5f, 0xf0, 0x90, 0xf7, 0xbb, 0xd3, 0x19,
	0xf3, 0x26, 0xcb, 0x4e, 0xf4, 0x84, 0xf6, 0x5d, 0x1e, 0xdd, 0xc8, 0xc7, 0xb1, 0xbc, 0xf4, 0x0f,
	0xaf, 0xda, 0xd0, 0x38, 0x13, 0xba, 0x7b, 0x6c, 0x19, 0xde, 0x62, 0xcb, 0xa5, 0xe3, 0x32, 0xf4,
	0x13, 0x6c, 0xf0, 0x79, 0x38, 0xe3, 0x81, 0xe3, 0x5d, 0x8c, 0x38, 0xf7, 0x30, 0x20, 0x80, 0xd6,
	0xec, 0x7a, 0x5c, 0x1c, 0x70, 0xee, 0xa1, 0x3b, 0x9a, 0x88, 0x05, 0x2b, 0x1f, 0x17, 0x08, 0xa0,
	0xcd, 0xde, 0x51, 0x67, 0x23, 0xa8, 0x93, 0x87, 0x74, 0x86, 0xc1, 0xca, 0x4f, 0x2d, 0xc5, 0x13,
	0xfa, 0x05, 0x36, 0x13, 0xcb, 0x59, 0x10, 0xf6, 0x7b, 0xb8, 0x48, 0x00, 0x2d, 0xdb, 0x09, 0xe8,
	0x7f, 0x51, 0x44, 0xbf, 0xc1, 0x56, 0x22, 0x5b, 0x4a, 0x5d, 0x89, 0x00, 0xda, 0xb6, 0x93, 0x6f,
	0xdf, 0x9d, 0xad, 0x09, 0x57, 0x52, 0x58, 0x26, 0x80, 0x36, 0x52, 0xe1, 0x7d, 0x29, 0xcc, 0x81,
	0x2d, 0x13, 0x57, 0x08, 0xa0, 0xc5, 0x0c, 0xd8, 0x32, 0xd7, 0xc0, 0x96, 0x89, 0xab, 0x04, 0x50,
	0x94, 0x05, 0xe7, 0x84, 0x2b, 0x29, 0xdc, 0x21, 0x80, 0x96, 0xb2, 0x60, 0xcb, 0x44, 0x7f, 0xc2,
	0x76, 0xea, 0x38, 0x9d, 0x3d, 0x66, 0x93, 0x7e, 0x0f, 0xd7, 0x08, 0xa0, 0x2d, 0xdb, 0x48, 0x3c,
	0x55, 0x1d, 0xfd, 0x0e, 0x93, 0xda, 0x45, 0xac, 0x85, 0x04, 0xd0, 0xaa, 0x9d, 0xd0, 0xce, 0x94,
	0x54, 0x6f, 0x68, 0xea, 0x71, 0x27, 0xc4, 0xbb, 0x04, 0xd0, 0x42, 0xda, 0xd0, 0x99, 0x28, 0x6e,
	0xc0, 0x5b, 0x26, 0xae, 0x13, 0x40, 0x8d, 0x3c, 0xde, 0x32, 0xd7, 0xf1, 0x96, 0x89, 0x1b, 0x04,
	0xd0, 0x4a, 0x0e, 0x9f, 0xeb, 0x7f, 0xc2, 0x57, 0x23, 0x8f, 0xe1, 0x26, 0x01, 0x14, 0xa4, 0xfd,
	0xff, 0x17, 0x55, 0xb3, 0x13, 0x0d, 0x17, 0xb3, 0xc0, 0xc5, 0xad, 0x68, 0xd7, 0xd2, 0x89, 0x46,
	0xd5, 0x4c, 0x43, 0xa3, 0xcb, 0x90, 0x2d, 0xb1, 0x41, 0x00, 0xad, 0xa7, 0x0d, 0x0d, 0x44, 0x11,
	0x9d, 0x6b, 0x67, 0x54, 0x8b, 0x86, 0xdb, 0x04, 0xd0, 0xdd, 0x5e, 0xf7, 0xba, 0x7b, 0xa9, 0xae,
	0x69, 0x53, 0x71, 0x2a, 0x7e, 0x86, 0x8d, 0x05, 0x9b, 0x33, 0x27, 0x64, 0x13, 0x99, 0x8a, 0x27,
	0x80, 0x14, 0xe9, 0x8e, 0x5d, 0x8f, 0xab, 0x51, 0x2c, 0x6c, 0x4d, 0x15, 0xc5, 0xe2, 0xa9, 0x50,
	0x6d, 0x9d, 0x8b, 0xd8, 0x23, 0xca, 0xc5, 0xaf, 0xb0, 0x99, 0x78, 0xca, 0x35, 0x7e, 0x26, 0x4c,
	0xcb, 0x76, 0x82, 0x92, 0xc1, 0xa0, 0xb0, 0x95, 0xe8, 0x54, 0x30, 0x9e, 0x0b, 0x61, 0xdb, 0x4e,
	0xbe, 0xaf, 0x92, 0xa1, 0x2b, 0x55, 0x32, 0x5e, 0x08, 0x65, 0x23, 0x55, 0xaa, 0x68, 0xe4, 0xd8,
	0x96, 0x89, 0x5f, 0x0a, 0x61, 0x31, 0xc3, 0xb6, 0xcc, 0x35, 0xb6, 0x65, 0xe2, 0x57, 0x42, 0x88,
	0xb2, 0xec, 0x9c, 0x52, 0x85, 0xe3, 0x4a, 0x28, 0x4b, 0x59, 0xb6, 0x65, 0xa2, 0x23, 0xd8, 0x4e,
	0x3d, 0xe3, 0x8d, 0x7f, 0x2d, 0xb4, 0x2d, 0xdb, 0x48, 0x5c, 0xe3, 0x78, 0xfc, 0x01, 0x93, 0x5a,
	0x12, 0x8f, 0x37, 0x42, 0x5c, 0xb5, 0x13, 0x60, 0x9c, 0x0f, 0xbd, 0x2b, 0x99, 0x8f, 0xb7, 0x42,
	0x59, 0x48, 0xbb, 0x92, 0x01, 0x59, 0x3f, 0x81, 0x65, 0xe2, 0x77, 0x42, 0x6a, 0xe4, 0x4f, 0x60,
	0x99, 0xeb, 0x27, 0xb0, 0x4c, 0xfc, 0x5e, 0x88, 0x2b, 0xb9, 0x13, 0xe4, 0xa6, 0xa0, 0x22, 0xf2,
	0x41, 0x48, 0x41, 0x3a, 0x05, 0x95, 0x91, 0xcc, 0x64, 0x65, 0x46, 0x3e, 0x0a, 0x65, 0x4d, 0x9b,
	0xac, 0x0c, 0x89, 0xde, 0x95, 0x0c, 0xc9, 0x27, 0x21, 0xac, 0xa7, 0x5d, 0xc9, 0x94, 0x3c, 0xd0,
	0xce, 0x19, 0xa7, 0xe4, 0xb3, 0x50, 0xde, 0x24, 0x26, 0xb1, 0x53, 0x1c, 0x93, 0x31, 0x6c, 0xf9,
	0xce, 0x5c, 0xee, 0xa9, 0xda, 0x98, 0x2f, 0xc5, 0xc8, 0xfb, 0xe4, 0xda, 0xde, 0xce, 0x3c, 0x5a,
	0xe8, 0x68, 0xb3, 0x86, 0x41, 0xb8, 0xb8, 0xb4, 0x1b, 0xbe, 0x5e, 0x43, 0x1e, 0x44, 0x02, 0x22,
	0xc7, 0x71, 0xe1, 0xab, 0x1e, 0xbe, 0x4a, 0xce, 0x3f, 0x5b, 0x70, 0xe4, 0xe4, 0x54, 0x41, 0xa2,
	0x0c, 0x3f, 0x57, 0x46, 0x53, 0x28, 0x6a, 0xf1, 0x2b, 0x95, 0xb1, 0xfe, 0x26, 0x59, 0xa7, 0x5b,
	0xb0, 0xd4, 0xab, 0x17, 0x91, 0x96, 0xa4, 0xa6, 0x9f, 0x29, 0x1e, 0xfc, 0x0b, 0xd1, 0x7a, 0xeb,
	0xc8, 0x80, 0xc5, 0x87, 0xec, 0x32, 0xfa, 0x0f, 0x2e, 0xdb, 0xe2, 0x16, 0xfd, 0x08, 0xcb, 0x8f,
	0x1c, 0x6f, 0xc5, 0xa2, 0xbf, 0xdc, 0xa2, 0x2d, 0x1f, 0x4e, 0x0a, 0x7f, 0x81, 0x83, 0x10, 0xee,
	0x6f, 0x6c, 0x4a, 0x37, 0xa9, 0x49, 0x93, 0xa1, 0x6e, 0x72, 0x83, 0x17, 0xaf, 0x51, 0x39, 0xdc,
	0xdb, 0xd0, 0x9e, 0xce, 0xac, 0x48, 0xe6, 0x40, 0x67, 0x6e, 0xfb, 0x9b, 0xa8, 0x01, 0x6b, 0xb0,
	0xaa, 0x3e, 0x3a, 0x34, 0x60, 0x29, 0xfa, 0x8d, 0xdc, 0x81, 0xa5, 0xf3, 0xa1, 0x7d, 0xdb, 0xf8,
	0x61, 0x70, 0x7a, 0xfe, 0xb7, 0xcb, 0xb9, 0xeb, 0xb1, 0x8e, 0xcb, 0x3d, 0x27, 0x70, 0x3b, 0x7c,
	0xe1, 0x76, 0x23, 0xe7, 0xee, 0xd8, 0x9f, 0xc8, 0xbb, 0xf1, 0xb1, 0xcb, 0x82, 0x63, 0x97, 0x77,
	0x43, 0xb6, 0x0c, 0x27, 0x4e, 0xe8, 0xc8, 0x72, 0x7f, 0x54, 0x91, 0xd7, 0xef, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0x70, 0x3f, 0x3d, 0x44, 0x09, 0x00, 0x00,
}
