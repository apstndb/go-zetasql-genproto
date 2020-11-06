// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zetasql/public/simple_table.proto

package public

import (
	fmt "fmt"
	typepb "github.com/apstndb/go-zetasql-genproto/zetasql/public/typepb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SimpleTableProto struct {
	Name                      *string              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SerializationId           *int64               `protobuf:"varint,2,opt,name=serialization_id,json=serializationId" json:"serialization_id,omitempty"`
	IsValueTable              *bool                `protobuf:"varint,3,opt,name=is_value_table,json=isValueTable" json:"is_value_table,omitempty"`
	Column                    []*SimpleColumnProto `protobuf:"bytes,4,rep,name=column" json:"column,omitempty"`
	PrimaryKeyColumnIndex     []int32              `protobuf:"varint,9,rep,name=primary_key_column_index,json=primaryKeyColumnIndex" json:"primary_key_column_index,omitempty"`
	NameInCatalog             *string              `protobuf:"bytes,5,opt,name=name_in_catalog,json=nameInCatalog" json:"name_in_catalog,omitempty"`
	AllowAnonymousColumnName  *bool                `protobuf:"varint,6,opt,name=allow_anonymous_column_name,json=allowAnonymousColumnName" json:"allow_anonymous_column_name,omitempty"`
	AllowDuplicateColumnNames *bool                `protobuf:"varint,7,opt,name=allow_duplicate_column_names,json=allowDuplicateColumnNames" json:"allow_duplicate_column_names,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}             `json:"-"`
	XXX_unrecognized          []byte               `json:"-"`
	XXX_sizecache             int32                `json:"-"`
}

func (m *SimpleTableProto) Reset()         { *m = SimpleTableProto{} }
func (m *SimpleTableProto) String() string { return proto.CompactTextString(m) }
func (*SimpleTableProto) ProtoMessage()    {}
func (*SimpleTableProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c39c4d202c369d41, []int{0}
}

func (m *SimpleTableProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleTableProto.Unmarshal(m, b)
}
func (m *SimpleTableProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleTableProto.Marshal(b, m, deterministic)
}
func (m *SimpleTableProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleTableProto.Merge(m, src)
}
func (m *SimpleTableProto) XXX_Size() int {
	return xxx_messageInfo_SimpleTableProto.Size(m)
}
func (m *SimpleTableProto) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleTableProto.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleTableProto proto.InternalMessageInfo

func (m *SimpleTableProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SimpleTableProto) GetSerializationId() int64 {
	if m != nil && m.SerializationId != nil {
		return *m.SerializationId
	}
	return 0
}

func (m *SimpleTableProto) GetIsValueTable() bool {
	if m != nil && m.IsValueTable != nil {
		return *m.IsValueTable
	}
	return false
}

func (m *SimpleTableProto) GetColumn() []*SimpleColumnProto {
	if m != nil {
		return m.Column
	}
	return nil
}

func (m *SimpleTableProto) GetPrimaryKeyColumnIndex() []int32 {
	if m != nil {
		return m.PrimaryKeyColumnIndex
	}
	return nil
}

func (m *SimpleTableProto) GetNameInCatalog() string {
	if m != nil && m.NameInCatalog != nil {
		return *m.NameInCatalog
	}
	return ""
}

func (m *SimpleTableProto) GetAllowAnonymousColumnName() bool {
	if m != nil && m.AllowAnonymousColumnName != nil {
		return *m.AllowAnonymousColumnName
	}
	return false
}

func (m *SimpleTableProto) GetAllowDuplicateColumnNames() bool {
	if m != nil && m.AllowDuplicateColumnNames != nil {
		return *m.AllowDuplicateColumnNames
	}
	return false
}

type SimpleColumnProto struct {
	Name                 *string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type                 *typepb.TypeProto `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	IsPseudoColumn       *bool             `protobuf:"varint,3,opt,name=is_pseudo_column,json=isPseudoColumn" json:"is_pseudo_column,omitempty"`
	IsWritableColumn     *bool             `protobuf:"varint,4,opt,name=is_writable_column,json=isWritableColumn,def=1" json:"is_writable_column,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SimpleColumnProto) Reset()         { *m = SimpleColumnProto{} }
func (m *SimpleColumnProto) String() string { return proto.CompactTextString(m) }
func (*SimpleColumnProto) ProtoMessage()    {}
func (*SimpleColumnProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c39c4d202c369d41, []int{1}
}

func (m *SimpleColumnProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleColumnProto.Unmarshal(m, b)
}
func (m *SimpleColumnProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleColumnProto.Marshal(b, m, deterministic)
}
func (m *SimpleColumnProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleColumnProto.Merge(m, src)
}
func (m *SimpleColumnProto) XXX_Size() int {
	return xxx_messageInfo_SimpleColumnProto.Size(m)
}
func (m *SimpleColumnProto) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleColumnProto.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleColumnProto proto.InternalMessageInfo

const Default_SimpleColumnProto_IsWritableColumn bool = true

func (m *SimpleColumnProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SimpleColumnProto) GetType() *typepb.TypeProto {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SimpleColumnProto) GetIsPseudoColumn() bool {
	if m != nil && m.IsPseudoColumn != nil {
		return *m.IsPseudoColumn
	}
	return false
}

func (m *SimpleColumnProto) GetIsWritableColumn() bool {
	if m != nil && m.IsWritableColumn != nil {
		return *m.IsWritableColumn
	}
	return Default_SimpleColumnProto_IsWritableColumn
}

func init() {
	proto.RegisterType((*SimpleTableProto)(nil), "zetasql.SimpleTableProto")
	proto.RegisterType((*SimpleColumnProto)(nil), "zetasql.SimpleColumnProto")
}

func init() { proto.RegisterFile("zetasql/public/simple_table.proto", fileDescriptor_c39c4d202c369d41) }

var fileDescriptor_c39c4d202c369d41 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xdd, 0x6a, 0x14, 0x31,
	0x14, 0x66, 0x9c, 0xed, 0x5f, 0xaa, 0xed, 0x36, 0x20, 0xa4, 0xd5, 0x8b, 0x69, 0x91, 0x32, 0xde,
	0xcc, 0xc2, 0xde, 0x08, 0x82, 0x88, 0xad, 0x37, 0xab, 0x20, 0x65, 0x2c, 0x0a, 0xde, 0x84, 0xec,
	0x4c, 0x58, 0x0f, 0x66, 0x92, 0x38, 0x49, 0xac, 0xd3, 0x67, 0xf0, 0x55, 0x7c, 0x47, 0x99, 0x93,
	0x6c, 0xb5, 0xb5, 0x77, 0xc3, 0xf7, 0x33, 0xe7, 0x3b, 0xe7, 0x0b, 0x39, 0xbe, 0x96, 0x5e, 0xb8,
	0xef, 0x6a, 0x66, 0xc3, 0x52, 0x41, 0x33, 0x73, 0xd0, 0x59, 0x25, 0xb9, 0x17, 0x4b, 0x25, 0x2b,
	0xdb, 0x1b, 0x6f, 0xe8, 0x56, 0x92, 0x1c, 0x1d, 0xde, 0xd1, 0xfa, 0xc1, 0x26, 0xcd, 0xc9, 0xaf,
	0x9c, 0x4c, 0x3f, 0xa2, 0xf5, 0x72, 0x74, 0x5e, 0xa0, 0x91, 0x92, 0x89, 0x16, 0x9d, 0x64, 0x59,
	0x91, 0x95, 0x3b, 0x35, 0x7e, 0xd3, 0xe7, 0x64, 0xea, 0x64, 0x0f, 0x42, 0xc1, 0xb5, 0xf0, 0x60,
	0x34, 0x87, 0x96, 0x3d, 0x28, 0xb2, 0x32, 0xaf, 0xf7, 0x6f, 0xe1, 0x8b, 0x96, 0x3e, 0x23, 0x7b,
	0xe0, 0xf8, 0x0f, 0xa1, 0x42, 0xca, 0xc3, 0xf2, 0x22, 0x2b, 0xb7, 0xeb, 0x87, 0xe0, 0x3e, 0x8d,
	0x20, 0x4e, 0xa2, 0x73, 0xb2, 0xd9, 0x18, 0x15, 0x3a, 0xcd, 0x26, 0x45, 0x5e, 0xee, 0xce, 0x8f,
	0xaa, 0x94, 0xb2, 0x8a, 0x79, 0xce, 0x91, 0xc4, 0x40, 0x75, 0x52, 0xd2, 0x17, 0x84, 0xd9, 0x1e,
	0x3a, 0xd1, 0x0f, 0xfc, 0x9b, 0x1c, 0x78, 0x44, 0x39, 0xe8, 0x56, 0xfe, 0x64, 0x3b, 0x45, 0x5e,
	0x6e, 0xd4, 0x8f, 0x13, 0xff, 0x5e, 0x0e, 0xf1, 0x07, 0x8b, 0x91, 0xa4, 0xa7, 0x64, 0x7f, 0xdc,
	0x82, 0x83, 0xe6, 0x8d, 0xf0, 0x42, 0x99, 0x15, 0xdb, 0xc0, 0xe5, 0x1e, 0x8d, 0xf0, 0x42, 0x9f,
	0x47, 0x90, 0xbe, 0x22, 0x4f, 0x84, 0x52, 0xe6, 0x8a, 0x0b, 0x6d, 0xf4, 0xd0, 0x99, 0xe0, 0xd6,
	0x43, 0xf0, 0x20, 0x9b, 0xb8, 0x07, 0x43, 0xc9, 0x9b, 0xb5, 0x22, 0xce, 0xf9, 0x30, 0x1e, 0xe9,
	0x35, 0x79, 0x1a, 0xed, 0x6d, 0xb0, 0x0a, 0x1a, 0xe1, 0xe5, 0xbf, 0x76, 0xc7, 0xb6, 0xd0, 0x7f,
	0x88, 0x9a, 0xb7, 0x6b, 0xc9, 0x5f, 0xbf, 0x3b, 0xf9, 0x9d, 0x91, 0x83, 0xff, 0xd6, 0xbf, 0xb7,
	0x8f, 0x53, 0x32, 0x19, 0x6b, 0xc4, 0x0e, 0x76, 0xe7, 0xf4, 0xe6, 0x78, 0x97, 0x83, 0x8d, 0x2d,
	0xd6, 0xc8, 0xd3, 0x92, 0x4c, 0xc1, 0x71, 0xeb, 0x64, 0x68, 0x4d, 0x0a, 0x93, 0xea, 0xd8, 0x03,
	0x77, 0x81, 0x70, 0x1c, 0x45, 0xe7, 0x84, 0x82, 0xe3, 0x57, 0x3d, 0x60, 0x67, 0xfc, 0xa6, 0x9c,
	0xac, 0xdc, 0x7e, 0x39, 0xf1, 0x7d, 0x90, 0xf5, 0x14, 0xdc, 0xe7, 0x44, 0x47, 0xcf, 0xd9, 0x3b,
	0x42, 0x1b, 0xd3, 0x55, 0x2b, 0x63, 0x56, 0x4a, 0xae, 0x33, 0x9c, 0x1d, 0xdc, 0x7d, 0x51, 0xee,
	0xcb, 0xf1, 0x0a, 0xfc, 0xd7, 0xb0, 0xac, 0x1a, 0xd3, 0xcd, 0x84, 0x75, 0x5e, 0xb7, 0xcb, 0xd9,
	0xed, 0x57, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x0d, 0x80, 0xdd, 0xd2, 0x02, 0x00, 0x00,
}
