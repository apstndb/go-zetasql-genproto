// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zetasql/public/type.proto

package public

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

type TypeKind int32

const (
	TypeKind___TypeKind__switch_must_have_a_default__ TypeKind = -1
	TypeKind_TYPE_UNKNOWN                             TypeKind = 0
	TypeKind_TYPE_INT32                               TypeKind = 1
	TypeKind_TYPE_INT64                               TypeKind = 2
	TypeKind_TYPE_UINT32                              TypeKind = 3
	TypeKind_TYPE_UINT64                              TypeKind = 4
	TypeKind_TYPE_BOOL                                TypeKind = 5
	TypeKind_TYPE_FLOAT                               TypeKind = 6
	TypeKind_TYPE_DOUBLE                              TypeKind = 7
	TypeKind_TYPE_STRING                              TypeKind = 8
	TypeKind_TYPE_BYTES                               TypeKind = 9
	TypeKind_TYPE_DATE                                TypeKind = 10
	TypeKind_TYPE_TIMESTAMP                           TypeKind = 19
	TypeKind_TYPE_ENUM                                TypeKind = 15
	TypeKind_TYPE_ARRAY                               TypeKind = 16
	TypeKind_TYPE_STRUCT                              TypeKind = 17
	TypeKind_TYPE_PROTO                               TypeKind = 18
	TypeKind_TYPE_TIME                                TypeKind = 20
	TypeKind_TYPE_DATETIME                            TypeKind = 21
	TypeKind_TYPE_GEOGRAPHY                           TypeKind = 22
	TypeKind_TYPE_NUMERIC                             TypeKind = 23
	TypeKind_TYPE_BIGNUMERIC                          TypeKind = 24
	TypeKind_TYPE_EXTENDED                            TypeKind = 25
	TypeKind_TYPE_JSON                                TypeKind = 26
	TypeKind_TYPE_INTERVAL                            TypeKind = 27
)

var TypeKind_name = map[int32]string{
	-1: "__TypeKind__switch_must_have_a_default__",
	0:  "TYPE_UNKNOWN",
	1:  "TYPE_INT32",
	2:  "TYPE_INT64",
	3:  "TYPE_UINT32",
	4:  "TYPE_UINT64",
	5:  "TYPE_BOOL",
	6:  "TYPE_FLOAT",
	7:  "TYPE_DOUBLE",
	8:  "TYPE_STRING",
	9:  "TYPE_BYTES",
	10: "TYPE_DATE",
	19: "TYPE_TIMESTAMP",
	15: "TYPE_ENUM",
	16: "TYPE_ARRAY",
	17: "TYPE_STRUCT",
	18: "TYPE_PROTO",
	20: "TYPE_TIME",
	21: "TYPE_DATETIME",
	22: "TYPE_GEOGRAPHY",
	23: "TYPE_NUMERIC",
	24: "TYPE_BIGNUMERIC",
	25: "TYPE_EXTENDED",
	26: "TYPE_JSON",
	27: "TYPE_INTERVAL",
}

var TypeKind_value = map[string]int32{
	"__TypeKind__switch_must_have_a_default__": -1,
	"TYPE_UNKNOWN":    0,
	"TYPE_INT32":      1,
	"TYPE_INT64":      2,
	"TYPE_UINT32":     3,
	"TYPE_UINT64":     4,
	"TYPE_BOOL":       5,
	"TYPE_FLOAT":      6,
	"TYPE_DOUBLE":     7,
	"TYPE_STRING":     8,
	"TYPE_BYTES":      9,
	"TYPE_DATE":       10,
	"TYPE_TIMESTAMP":  19,
	"TYPE_ENUM":       15,
	"TYPE_ARRAY":      16,
	"TYPE_STRUCT":     17,
	"TYPE_PROTO":      18,
	"TYPE_TIME":       20,
	"TYPE_DATETIME":   21,
	"TYPE_GEOGRAPHY":  22,
	"TYPE_NUMERIC":    23,
	"TYPE_BIGNUMERIC": 24,
	"TYPE_EXTENDED":   25,
	"TYPE_JSON":       26,
	"TYPE_INTERVAL":   27,
}

func (x TypeKind) Enum() *TypeKind {
	p := new(TypeKind)
	*p = x
	return p
}

func (x TypeKind) String() string {
	return proto.EnumName(TypeKind_name, int32(x))
}

func (x *TypeKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TypeKind_value, data, "TypeKind")
	if err != nil {
		return err
	}
	*x = TypeKind(value)
	return nil
}

func (TypeKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3cd30f7a11376b6e, []int{0}
}

type TypeProto struct {
	TypeKind             *TypeKind                       `protobuf:"varint,1,opt,name=type_kind,json=typeKind,enum=zetasql.TypeKind" json:"type_kind,omitempty"`
	ArrayType            *ArrayTypeProto                 `protobuf:"bytes,2,opt,name=array_type,json=arrayType" json:"array_type,omitempty"`
	StructType           *StructTypeProto                `protobuf:"bytes,3,opt,name=struct_type,json=structType" json:"struct_type,omitempty"`
	ProtoType            *ProtoTypeProto                 `protobuf:"bytes,4,opt,name=proto_type,json=protoType" json:"proto_type,omitempty"`
	EnumType             *EnumTypeProto                  `protobuf:"bytes,5,opt,name=enum_type,json=enumType" json:"enum_type,omitempty"`
	FileDescriptorSet    []*descriptor.FileDescriptorSet `protobuf:"bytes,6,rep,name=file_descriptor_set,json=fileDescriptorSet" json:"file_descriptor_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TypeProto) Reset()         { *m = TypeProto{} }
func (m *TypeProto) String() string { return proto.CompactTextString(m) }
func (*TypeProto) ProtoMessage()    {}
func (*TypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd30f7a11376b6e, []int{0}
}

func (m *TypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeProto.Unmarshal(m, b)
}
func (m *TypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeProto.Marshal(b, m, deterministic)
}
func (m *TypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeProto.Merge(m, src)
}
func (m *TypeProto) XXX_Size() int {
	return xxx_messageInfo_TypeProto.Size(m)
}
func (m *TypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_TypeProto proto.InternalMessageInfo

func (m *TypeProto) GetTypeKind() TypeKind {
	if m != nil && m.TypeKind != nil {
		return *m.TypeKind
	}
	return TypeKind___TypeKind__switch_must_have_a_default__
}

func (m *TypeProto) GetArrayType() *ArrayTypeProto {
	if m != nil {
		return m.ArrayType
	}
	return nil
}

func (m *TypeProto) GetStructType() *StructTypeProto {
	if m != nil {
		return m.StructType
	}
	return nil
}

func (m *TypeProto) GetProtoType() *ProtoTypeProto {
	if m != nil {
		return m.ProtoType
	}
	return nil
}

func (m *TypeProto) GetEnumType() *EnumTypeProto {
	if m != nil {
		return m.EnumType
	}
	return nil
}

func (m *TypeProto) GetFileDescriptorSet() []*descriptor.FileDescriptorSet {
	if m != nil {
		return m.FileDescriptorSet
	}
	return nil
}

type ArrayTypeProto struct {
	ElementType          *TypeProto `protobuf:"bytes,1,opt,name=element_type,json=elementType" json:"element_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ArrayTypeProto) Reset()         { *m = ArrayTypeProto{} }
func (m *ArrayTypeProto) String() string { return proto.CompactTextString(m) }
func (*ArrayTypeProto) ProtoMessage()    {}
func (*ArrayTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd30f7a11376b6e, []int{1}
}

func (m *ArrayTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayTypeProto.Unmarshal(m, b)
}
func (m *ArrayTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayTypeProto.Marshal(b, m, deterministic)
}
func (m *ArrayTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayTypeProto.Merge(m, src)
}
func (m *ArrayTypeProto) XXX_Size() int {
	return xxx_messageInfo_ArrayTypeProto.Size(m)
}
func (m *ArrayTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayTypeProto proto.InternalMessageInfo

func (m *ArrayTypeProto) GetElementType() *TypeProto {
	if m != nil {
		return m.ElementType
	}
	return nil
}

type StructFieldProto struct {
	FieldName            *string    `protobuf:"bytes,1,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	FieldType            *TypeProto `protobuf:"bytes,2,opt,name=field_type,json=fieldType" json:"field_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StructFieldProto) Reset()         { *m = StructFieldProto{} }
func (m *StructFieldProto) String() string { return proto.CompactTextString(m) }
func (*StructFieldProto) ProtoMessage()    {}
func (*StructFieldProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd30f7a11376b6e, []int{2}
}

func (m *StructFieldProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructFieldProto.Unmarshal(m, b)
}
func (m *StructFieldProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructFieldProto.Marshal(b, m, deterministic)
}
func (m *StructFieldProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructFieldProto.Merge(m, src)
}
func (m *StructFieldProto) XXX_Size() int {
	return xxx_messageInfo_StructFieldProto.Size(m)
}
func (m *StructFieldProto) XXX_DiscardUnknown() {
	xxx_messageInfo_StructFieldProto.DiscardUnknown(m)
}

var xxx_messageInfo_StructFieldProto proto.InternalMessageInfo

func (m *StructFieldProto) GetFieldName() string {
	if m != nil && m.FieldName != nil {
		return *m.FieldName
	}
	return ""
}

func (m *StructFieldProto) GetFieldType() *TypeProto {
	if m != nil {
		return m.FieldType
	}
	return nil
}

type StructTypeProto struct {
	Field                []*StructFieldProto `protobuf:"bytes,1,rep,name=field" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StructTypeProto) Reset()         { *m = StructTypeProto{} }
func (m *StructTypeProto) String() string { return proto.CompactTextString(m) }
func (*StructTypeProto) ProtoMessage()    {}
func (*StructTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd30f7a11376b6e, []int{3}
}

func (m *StructTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructTypeProto.Unmarshal(m, b)
}
func (m *StructTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructTypeProto.Marshal(b, m, deterministic)
}
func (m *StructTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructTypeProto.Merge(m, src)
}
func (m *StructTypeProto) XXX_Size() int {
	return xxx_messageInfo_StructTypeProto.Size(m)
}
func (m *StructTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_StructTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_StructTypeProto proto.InternalMessageInfo

func (m *StructTypeProto) GetField() []*StructFieldProto {
	if m != nil {
		return m.Field
	}
	return nil
}

type ProtoTypeProto struct {
	ProtoName              *string  `protobuf:"bytes,1,opt,name=proto_name,json=protoName" json:"proto_name,omitempty"`
	ProtoFileName          *string  `protobuf:"bytes,2,opt,name=proto_file_name,json=protoFileName" json:"proto_file_name,omitempty"`
	FileDescriptorSetIndex *int32   `protobuf:"varint,3,opt,name=file_descriptor_set_index,json=fileDescriptorSetIndex,def=0" json:"file_descriptor_set_index,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ProtoTypeProto) Reset()         { *m = ProtoTypeProto{} }
func (m *ProtoTypeProto) String() string { return proto.CompactTextString(m) }
func (*ProtoTypeProto) ProtoMessage()    {}
func (*ProtoTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd30f7a11376b6e, []int{4}
}

func (m *ProtoTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoTypeProto.Unmarshal(m, b)
}
func (m *ProtoTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoTypeProto.Marshal(b, m, deterministic)
}
func (m *ProtoTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoTypeProto.Merge(m, src)
}
func (m *ProtoTypeProto) XXX_Size() int {
	return xxx_messageInfo_ProtoTypeProto.Size(m)
}
func (m *ProtoTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoTypeProto proto.InternalMessageInfo

const Default_ProtoTypeProto_FileDescriptorSetIndex int32 = 0

func (m *ProtoTypeProto) GetProtoName() string {
	if m != nil && m.ProtoName != nil {
		return *m.ProtoName
	}
	return ""
}

func (m *ProtoTypeProto) GetProtoFileName() string {
	if m != nil && m.ProtoFileName != nil {
		return *m.ProtoFileName
	}
	return ""
}

func (m *ProtoTypeProto) GetFileDescriptorSetIndex() int32 {
	if m != nil && m.FileDescriptorSetIndex != nil {
		return *m.FileDescriptorSetIndex
	}
	return Default_ProtoTypeProto_FileDescriptorSetIndex
}

type EnumTypeProto struct {
	EnumName               *string  `protobuf:"bytes,1,opt,name=enum_name,json=enumName" json:"enum_name,omitempty"`
	EnumFileName           *string  `protobuf:"bytes,2,opt,name=enum_file_name,json=enumFileName" json:"enum_file_name,omitempty"`
	FileDescriptorSetIndex *int32   `protobuf:"varint,3,opt,name=file_descriptor_set_index,json=fileDescriptorSetIndex,def=0" json:"file_descriptor_set_index,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *EnumTypeProto) Reset()         { *m = EnumTypeProto{} }
func (m *EnumTypeProto) String() string { return proto.CompactTextString(m) }
func (*EnumTypeProto) ProtoMessage()    {}
func (*EnumTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd30f7a11376b6e, []int{5}
}

func (m *EnumTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumTypeProto.Unmarshal(m, b)
}
func (m *EnumTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumTypeProto.Marshal(b, m, deterministic)
}
func (m *EnumTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumTypeProto.Merge(m, src)
}
func (m *EnumTypeProto) XXX_Size() int {
	return xxx_messageInfo_EnumTypeProto.Size(m)
}
func (m *EnumTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_EnumTypeProto proto.InternalMessageInfo

const Default_EnumTypeProto_FileDescriptorSetIndex int32 = 0

func (m *EnumTypeProto) GetEnumName() string {
	if m != nil && m.EnumName != nil {
		return *m.EnumName
	}
	return ""
}

func (m *EnumTypeProto) GetEnumFileName() string {
	if m != nil && m.EnumFileName != nil {
		return *m.EnumFileName
	}
	return ""
}

func (m *EnumTypeProto) GetFileDescriptorSetIndex() int32 {
	if m != nil && m.FileDescriptorSetIndex != nil {
		return *m.FileDescriptorSetIndex
	}
	return Default_EnumTypeProto_FileDescriptorSetIndex
}

func init() {
	proto.RegisterEnum("zetasql.TypeKind", TypeKind_name, TypeKind_value)
	proto.RegisterType((*TypeProto)(nil), "zetasql.TypeProto")
	proto.RegisterType((*ArrayTypeProto)(nil), "zetasql.ArrayTypeProto")
	proto.RegisterType((*StructFieldProto)(nil), "zetasql.StructFieldProto")
	proto.RegisterType((*StructTypeProto)(nil), "zetasql.StructTypeProto")
	proto.RegisterType((*ProtoTypeProto)(nil), "zetasql.ProtoTypeProto")
	proto.RegisterType((*EnumTypeProto)(nil), "zetasql.EnumTypeProto")
}

func init() { proto.RegisterFile("zetasql/public/type.proto", fileDescriptor_3cd30f7a11376b6e) }

var fileDescriptor_3cd30f7a11376b6e = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x73, 0xfa, 0x54,
	0x14, 0x35, 0xf0, 0xe3, 0xf7, 0x23, 0x97, 0x02, 0x8f, 0x57, 0x6d, 0xa1, 0x1d, 0x67, 0x90, 0x71,
	0x1c, 0xc6, 0x45, 0x50, 0xfa, 0x67, 0x46, 0xc7, 0x4d, 0x28, 0x29, 0x62, 0x21, 0xc1, 0x10, 0x54,
	0xba, 0x79, 0x13, 0xc8, 0xa3, 0xcd, 0x98, 0x04, 0x24, 0x89, 0x5a, 0x3f, 0x85, 0x1b, 0x3f, 0x9d,
	0x1f, 0xc2, 0xad, 0x3b, 0x9d, 0xbc, 0x97, 0x3c, 0xfe, 0xd8, 0xa5, 0xac, 0xb8, 0xe7, 0x9e, 0x73,
	0xef, 0x99, 0x73, 0x93, 0x40, 0xe3, 0x37, 0x1a, 0xd9, 0xe1, 0x4f, 0x5e, 0x67, 0x13, 0x2f, 0x3c,
	0x77, 0xd9, 0x89, 0x5e, 0x36, 0x54, 0xd9, 0x6c, 0xd7, 0xd1, 0x1a, 0xbf, 0x4b, 0x5b, 0x17, 0xcd,
	0xa7, 0xf5, 0xfa, 0xc9, 0xa3, 0x1d, 0x06, 0x2f, 0xe2, 0x55, 0xc7, 0xa1, 0xe1, 0x72, 0xeb, 0x6e,
	0xa2, 0xf5, 0x96, 0x53, 0x5b, 0x7f, 0xe5, 0x40, 0xb6, 0x5e, 0x36, 0x74, 0xc2, 0x84, 0x0a, 0xc8,
	0xc9, 0x18, 0xf2, 0xa3, 0x1b, 0x38, 0x75, 0xa9, 0x29, 0xb5, 0x2b, 0xdd, 0x9a, 0x92, 0x0e, 0x53,
	0x12, 0xda, 0x83, 0x1b, 0x38, 0x66, 0x31, 0x4a, 0xff, 0xe1, 0x5b, 0x00, 0x7b, 0xbb, 0xb5, 0x5f,
	0x48, 0x82, 0xd4, 0x73, 0x4d, 0xa9, 0x5d, 0xea, 0x9e, 0x0b, 0x81, 0x9a, 0xb4, 0xc4, 0x70, 0x53,
	0xb6, 0xb3, 0x1a, 0x7f, 0x01, 0xa5, 0x30, 0xda, 0xc6, 0xcb, 0x88, 0x0b, 0xf3, 0x4c, 0x58, 0x17,
	0xc2, 0x29, 0xeb, 0xed, 0x94, 0x10, 0x0a, 0x20, 0x59, 0xc9, 0x9c, 0x73, 0xe5, 0x9b, 0xa3, 0x95,
	0x8c, 0xbf, 0xb7, 0x72, 0x93, 0xd5, 0xf8, 0x0a, 0x64, 0x1a, 0xc4, 0x3e, 0x97, 0x15, 0x98, 0xec,
	0x4c, 0xc8, 0xb4, 0x20, 0xf6, 0x77, 0xaa, 0x22, 0x4d, 0x4b, 0x6c, 0xc2, 0xe9, 0xca, 0xf5, 0x28,
	0xd9, 0xc5, 0x46, 0x42, 0x1a, 0xd5, 0xdf, 0x36, 0xf3, 0xed, 0x52, 0xb7, 0xa5, 0xf0, 0x74, 0x95,
	0x2c, 0x5d, 0xe5, 0xde, 0xf5, 0x68, 0x5f, 0x50, 0xa7, 0x34, 0x32, 0x6b, 0xab, 0x63, 0xa8, 0x35,
	0x80, 0xca, 0x61, 0x30, 0xf8, 0x06, 0x4e, 0xa8, 0x47, 0x7d, 0x1a, 0xa4, 0x71, 0x48, 0xcc, 0x1d,
	0x3e, 0x08, 0x9e, 0x3b, 0x2b, 0xa5, 0xbc, 0x04, 0x69, 0x39, 0x80, 0x78, 0x50, 0xf7, 0x2e, 0xf5,
	0x1c, 0x3e, 0xea, 0x43, 0x80, 0x55, 0x52, 0x91, 0xc0, 0xf6, 0xf9, 0x20, 0xd9, 0x94, 0x19, 0xa2,
	0xdb, 0x3e, 0xc5, 0x9f, 0x67, 0xed, 0xbd, 0x7b, 0xbd, 0xb6, 0x87, 0x4b, 0xd8, 0x96, 0x1e, 0x54,
	0x8f, 0xce, 0x81, 0x3b, 0x50, 0x60, 0xfd, 0xba, 0xc4, 0x72, 0x68, 0x1c, 0xdd, 0x6d, 0x67, 0xc7,
	0xe4, 0xbc, 0xd6, 0x1f, 0x12, 0x54, 0x0e, 0x2f, 0x93, 0x18, 0xe5, 0x67, 0xdc, 0x37, 0xca, 0x10,
	0x66, 0xf4, 0x13, 0xa8, 0xf2, 0x36, 0x8b, 0x9f, 0x71, 0x72, 0x8c, 0x53, 0x66, 0x70, 0x12, 0x34,
	0xe3, 0x7d, 0x05, 0x8d, 0x57, 0x0e, 0x44, 0xdc, 0xc0, 0xa1, 0xbf, 0xb2, 0xc7, 0xaa, 0xf0, 0xa5,
	0xf4, 0x99, 0x79, 0xf6, 0x9f, 0x2b, 0x0c, 0x13, 0x42, 0xeb, 0x77, 0x09, 0xca, 0x07, 0xa7, 0xc7,
	0x97, 0xe9, 0x53, 0xb2, 0xe7, 0x8a, 0x3d, 0x0d, 0x6c, 0xd9, 0xc7, 0x50, 0x61, 0xcd, 0x63, 0x4f,
	0x27, 0x09, 0xfa, 0xff, 0x58, 0xfa, 0xf4, 0xcf, 0x3c, 0x14, 0xb3, 0x17, 0x0d, 0xdf, 0x40, 0x9b,
	0x90, 0xac, 0x22, 0x24, 0xfc, 0xc5, 0x8d, 0x96, 0xcf, 0xc4, 0x8f, 0xc3, 0x88, 0x3c, 0xdb, 0x3f,
	0x53, 0x62, 0x13, 0x87, 0xae, 0xec, 0xd8, 0x8b, 0x08, 0x41, 0xff, 0x64, 0x3f, 0x09, 0x23, 0x38,
	0xb1, 0xe6, 0x13, 0x8d, 0xcc, 0xf4, 0x07, 0xdd, 0xf8, 0x5e, 0x47, 0xef, 0xe1, 0x0a, 0x00, 0x43,
	0x86, 0xba, 0x75, 0xd5, 0x45, 0xd2, 0x7e, 0x7d, 0x7b, 0x8d, 0x72, 0xb8, 0x0a, 0x25, 0xae, 0xe0,
	0x84, 0xfc, 0x01, 0x70, 0x7b, 0x8d, 0xde, 0xe0, 0x32, 0xc8, 0x0c, 0xe8, 0x19, 0xc6, 0x08, 0x15,
	0xc4, 0x80, 0xfb, 0x91, 0xa1, 0x5a, 0xe8, 0xad, 0xe0, 0xf7, 0x8d, 0x59, 0x6f, 0xa4, 0xa1, 0x77,
	0x02, 0x98, 0x5a, 0xe6, 0x50, 0x1f, 0xa0, 0xa2, 0x50, 0xf4, 0xe6, 0x96, 0x36, 0x45, 0xb2, 0x18,
	0xd8, 0x57, 0x2d, 0x0d, 0x01, 0xc6, 0x50, 0x61, 0xa5, 0x35, 0x1c, 0x6b, 0x53, 0x4b, 0x1d, 0x4f,
	0xd0, 0xa9, 0xa0, 0x68, 0xfa, 0x6c, 0x8c, 0xaa, 0x62, 0x82, 0x6a, 0x9a, 0xea, 0x1c, 0xa1, 0xfd,
	0x15, 0xb3, 0x3b, 0x0b, 0xd5, 0x04, 0x61, 0x62, 0x1a, 0x96, 0x81, 0xb0, 0xd0, 0x27, 0x33, 0xd1,
	0xfb, 0xb8, 0x06, 0x65, 0xb1, 0x91, 0x41, 0x1f, 0x88, 0xad, 0x03, 0xcd, 0x18, 0x98, 0xea, 0xe4,
	0xeb, 0x39, 0x3a, 0x13, 0xe9, 0xe9, 0xb3, 0xb1, 0x66, 0x0e, 0xef, 0xd0, 0x39, 0x3e, 0x85, 0x2a,
	0xb7, 0x3e, 0x1c, 0x64, 0x60, 0x5d, 0x4c, 0xd3, 0x7e, 0xb0, 0x34, 0xbd, 0xaf, 0xf5, 0x51, 0x43,
	0xec, 0xfb, 0x66, 0x6a, 0xe8, 0xe8, 0x42, 0x30, 0x86, 0xba, 0xa5, 0x99, 0xdf, 0xa9, 0x23, 0x74,
	0xd9, 0x1b, 0x00, 0x5e, 0xae, 0xfd, 0xec, 0xbb, 0x91, 0xbe, 0x36, 0xbd, 0xd2, 0x23, 0x8d, 0xec,
	0xe9, 0xb7, 0xa3, 0xe4, 0xd2, 0x8f, 0x1f, 0x3d, 0xb9, 0xd1, 0x73, 0xbc, 0x50, 0x96, 0x6b, 0xbf,
	0x63, 0x6f, 0xc2, 0x28, 0x70, 0x16, 0x9d, 0xc3, 0x2f, 0xfd, 0xdf, 0x92, 0xf4, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7f, 0xdc, 0xc6, 0xf2, 0xfd, 0x05, 0x00, 0x00,
}