// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zetasql/public/proto/wire_format_annotation.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type TableType int32

const (
	TableType_DEFAULT_TABLE_TYPE TableType = 0
	TableType_SQL_TABLE          TableType = 1
	TableType_VALUE_TABLE        TableType = 2
)

var TableType_name = map[int32]string{
	0: "DEFAULT_TABLE_TYPE",
	1: "SQL_TABLE",
	2: "VALUE_TABLE",
}

var TableType_value = map[string]int32{
	"DEFAULT_TABLE_TYPE": 0,
	"SQL_TABLE":          1,
	"VALUE_TABLE":        2,
}

func (x TableType) Enum() *TableType {
	p := new(TableType)
	*p = x
	return p
}

func (x TableType) String() string {
	return proto.EnumName(TableType_name, int32(x))
}

func (x *TableType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TableType_value, data, "TableType")
	if err != nil {
		return err
	}
	*x = TableType(value)
	return nil
}

func (TableType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a62354060f04f71f, []int{0}
}

type WireFormatAnnotationEmptyMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WireFormatAnnotationEmptyMessage) Reset()         { *m = WireFormatAnnotationEmptyMessage{} }
func (m *WireFormatAnnotationEmptyMessage) String() string { return proto.CompactTextString(m) }
func (*WireFormatAnnotationEmptyMessage) ProtoMessage()    {}
func (*WireFormatAnnotationEmptyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a62354060f04f71f, []int{0}
}

func (m *WireFormatAnnotationEmptyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WireFormatAnnotationEmptyMessage.Unmarshal(m, b)
}
func (m *WireFormatAnnotationEmptyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WireFormatAnnotationEmptyMessage.Marshal(b, m, deterministic)
}
func (m *WireFormatAnnotationEmptyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WireFormatAnnotationEmptyMessage.Merge(m, src)
}
func (m *WireFormatAnnotationEmptyMessage) XXX_Size() int {
	return xxx_messageInfo_WireFormatAnnotationEmptyMessage.Size(m)
}
func (m *WireFormatAnnotationEmptyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WireFormatAnnotationEmptyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WireFormatAnnotationEmptyMessage proto.InternalMessageInfo

var E_StructFieldName = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         69661525,
	Name:          "zetasql.struct_field_name",
	Tag:           "bytes,69661525,opt,name=struct_field_name",
	Filename:      "zetasql/public/proto/wire_format_annotation.proto",
}

var E_IsRawProto = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         163760138,
	Name:          "zetasql.is_raw_proto",
	Tag:           "varint,163760138,opt,name=is_raw_proto",
	Filename:      "zetasql/public/proto/wire_format_annotation.proto",
}

var E_IsWrapper = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         49730018,
	Name:          "zetasql.is_wrapper",
	Tag:           "varint,49730018,opt,name=is_wrapper",
	Filename:      "zetasql/public/proto/wire_format_annotation.proto",
}

var E_IsStruct = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         49727810,
	Name:          "zetasql.is_struct",
	Tag:           "varint,49727810,opt,name=is_struct",
	Filename:      "zetasql/public/proto/wire_format_annotation.proto",
}

var E_TableType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*TableType)(nil),
	Field:         86830980,
	Name:          "zetasql.table_type",
	Tag:           "varint,86830980,opt,name=table_type,enum=zetasql.TableType",
	Filename:      "zetasql/public/proto/wire_format_annotation.proto",
}

func init() {
	proto.RegisterEnum("zetasql.TableType", TableType_name, TableType_value)
	proto.RegisterType((*WireFormatAnnotationEmptyMessage)(nil), "zetasql.WireFormatAnnotationEmptyMessage")
	proto.RegisterExtension(E_StructFieldName)
	proto.RegisterExtension(E_IsRawProto)
	proto.RegisterExtension(E_IsWrapper)
	proto.RegisterExtension(E_IsStruct)
	proto.RegisterExtension(E_TableType)
}

func init() {
	proto.RegisterFile("zetasql/public/proto/wire_format_annotation.proto", fileDescriptor_a62354060f04f71f)
}

var fileDescriptor_a62354060f04f71f = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbd, 0x8e, 0xd3, 0x40,
	0x14, 0x85, 0x31, 0x0d, 0xf1, 0xf0, 0x93, 0x30, 0x05, 0x7f, 0x12, 0xc2, 0xa4, 0x21, 0x50, 0xd8,
	0x82, 0xd2, 0x0d, 0xb2, 0xc1, 0xa9, 0x4c, 0x08, 0x8e, 0x43, 0x80, 0x66, 0x34, 0x76, 0x26, 0x66,
	0x24, 0xdb, 0x33, 0xcc, 0x8c, 0x65, 0x85, 0x9a, 0x0a, 0x1a, 0x44, 0x4b, 0xb9, 0x8f, 0xb0, 0x6f,
	0xb0, 0xef, 0xb0, 0xd5, 0x16, 0x2b, 0x6d, 0xb1, 0xcf, 0xb1, 0xb2, 0xc7, 0x4e, 0xb3, 0x91, 0xd2,
	0x8d, 0xee, 0xbd, 0xe7, 0xd3, 0xb9, 0xf7, 0x0c, 0x78, 0xfd, 0x93, 0x28, 0x2c, 0x7f, 0xe4, 0x0e,
	0xaf, 0x92, 0x9c, 0xa6, 0x0e, 0x17, 0x4c, 0x31, 0xa7, 0xa6, 0x82, 0xa0, 0x0d, 0x13, 0x05, 0x56,
	0x08, 0x97, 0x25, 0x53, 0x58, 0x51, 0x56, 0xda, 0x6d, 0x13, 0xde, 0xea, 0x24, 0x4f, 0xac, 0x8c,
	0xb1, 0x2c, 0x27, 0x5a, 0x93, 0x54, 0x1b, 0x67, 0x4d, 0x64, 0x2a, 0x28, 0x57, 0x4c, 0xe8, 0xd1,
	0xf1, 0x18, 0x58, 0x2b, 0x2a, 0xc8, 0xb4, 0x25, 0x79, 0x3b, 0x50, 0x50, 0x70, 0xb5, 0xfd, 0x40,
	0xa4, 0xc4, 0x19, 0x79, 0xf5, 0x0e, 0x98, 0x31, 0x4e, 0x72, 0x12, 0x6f, 0x39, 0x81, 0x0f, 0x00,
	0x7c, 0x1f, 0x4c, 0xbd, 0x65, 0x18, 0xa3, 0xd8, 0xf3, 0xc3, 0x00, 0xc5, 0x5f, 0xe7, 0xc1, 0xe8,
	0x06, 0xbc, 0x0b, 0xcc, 0xc5, 0xa7, 0x50, 0xd7, 0x46, 0x06, 0x1c, 0x82, 0xdb, 0x9f, 0xbd, 0x70,
	0x19, 0x74, 0x85, 0x9b, 0x6e, 0x08, 0xee, 0x4b, 0x25, 0xaa, 0x54, 0xa1, 0x0d, 0x25, 0xf9, 0x1a,
	0x95, 0xb8, 0x20, 0xf0, 0xa9, 0xad, 0x0d, 0xda, 0xbd, 0x41, 0x7b, 0xda, 0x34, 0x3f, 0xf2, 0xc6,
	0x83, 0x7c, 0x74, 0x7a, 0xf9, 0xff, 0xb9, 0x65, 0x4c, 0xcc, 0x68, 0xa8, 0xa5, 0x6d, 0x73, 0x86,
	0x0b, 0xe2, 0xfa, 0xe0, 0x0e, 0x95, 0x48, 0xe0, 0x1a, 0xe9, 0x8d, 0x0f, 0x80, 0x7e, 0xff, 0xfd,
	0x33, 0xb3, 0x8c, 0xc9, 0x20, 0x02, 0x54, 0x46, 0xb8, 0x9e, 0x37, 0x43, 0xae, 0x07, 0x00, 0x95,
	0xa8, 0x16, 0x98, 0x73, 0x22, 0xe0, 0xb3, 0x6b, 0x84, 0x6e, 0xfd, 0x9e, 0x71, 0x71, 0x74, 0xf6,
	0xb0, 0x65, 0x98, 0x54, 0xae, 0xb4, 0xc8, 0x7d, 0x0b, 0x4c, 0x2a, 0x91, 0x36, 0x77, 0x98, 0x70,
	0xf2, 0xaf, 0x23, 0x0c, 0xa8, 0x5c, 0xb4, 0x1a, 0x77, 0x09, 0x80, 0x6a, 0x4e, 0x8b, 0x54, 0x73,
	0xdb, 0x83, 0x84, 0x5f, 0xe7, 0xc7, 0x2f, 0x2d, 0x63, 0x72, 0xef, 0x0d, 0xb4, 0xbb, 0x88, 0xed,
	0x5d, 0x32, 0x91, 0xa9, 0xfa, 0xa7, 0xff, 0x05, 0xc0, 0x94, 0x15, 0x3d, 0xad, 0x1b, 0xf5, 0x1f,
	0xef, 0x4b, 0xba, 0xbd, 0xc5, 0xb7, 0x17, 0x19, 0x55, 0xdf, 0xab, 0xc4, 0x4e, 0x59, 0xe1, 0x60,
	0x2e, 0x55, 0xb9, 0x4e, 0x9c, 0x7d, 0xff, 0xee, 0x2a, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x70, 0xbb,
	0xa0, 0x8e, 0x02, 0x00, 0x00,
}
