// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zetasql/public/function.proto

package zetasql

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SignatureArgumentKind int32

const (
	SignatureArgumentKind_ARG_TYPE_FIXED                                        SignatureArgumentKind = 0
	SignatureArgumentKind_ARG_TYPE_ANY_1                                        SignatureArgumentKind = 1
	SignatureArgumentKind_ARG_TYPE_ANY_2                                        SignatureArgumentKind = 2
	SignatureArgumentKind_ARG_ARRAY_TYPE_ANY_1                                  SignatureArgumentKind = 3
	SignatureArgumentKind_ARG_ARRAY_TYPE_ANY_2                                  SignatureArgumentKind = 4
	SignatureArgumentKind_ARG_PROTO_MAP_ANY                                     SignatureArgumentKind = 14
	SignatureArgumentKind_ARG_PROTO_MAP_KEY_ANY                                 SignatureArgumentKind = 15
	SignatureArgumentKind_ARG_PROTO_MAP_VALUE_ANY                               SignatureArgumentKind = 16
	SignatureArgumentKind_ARG_PROTO_ANY                                         SignatureArgumentKind = 5
	SignatureArgumentKind_ARG_STRUCT_ANY                                        SignatureArgumentKind = 6
	SignatureArgumentKind_ARG_ENUM_ANY                                          SignatureArgumentKind = 7
	SignatureArgumentKind_ARG_TYPE_ARBITRARY                                    SignatureArgumentKind = 8
	SignatureArgumentKind_ARG_TYPE_RELATION                                     SignatureArgumentKind = 9
	SignatureArgumentKind_ARG_TYPE_VOID                                         SignatureArgumentKind = 10
	SignatureArgumentKind_ARG_TYPE_MODEL                                        SignatureArgumentKind = 11
	SignatureArgumentKind_ARG_TYPE_CONNECTION                                   SignatureArgumentKind = 12
	SignatureArgumentKind_ARG_TYPE_DESCRIPTOR                                   SignatureArgumentKind = 13
	SignatureArgumentKind_ARG_TYPE_LAMBDA                                       SignatureArgumentKind = 17
	SignatureArgumentKind___SignatureArgumentKind__switch_must_have_a_default__ SignatureArgumentKind = -1
)

var SignatureArgumentKind_name = map[int32]string{
	0:  "ARG_TYPE_FIXED",
	1:  "ARG_TYPE_ANY_1",
	2:  "ARG_TYPE_ANY_2",
	3:  "ARG_ARRAY_TYPE_ANY_1",
	4:  "ARG_ARRAY_TYPE_ANY_2",
	14: "ARG_PROTO_MAP_ANY",
	15: "ARG_PROTO_MAP_KEY_ANY",
	16: "ARG_PROTO_MAP_VALUE_ANY",
	5:  "ARG_PROTO_ANY",
	6:  "ARG_STRUCT_ANY",
	7:  "ARG_ENUM_ANY",
	8:  "ARG_TYPE_ARBITRARY",
	9:  "ARG_TYPE_RELATION",
	10: "ARG_TYPE_VOID",
	11: "ARG_TYPE_MODEL",
	12: "ARG_TYPE_CONNECTION",
	13: "ARG_TYPE_DESCRIPTOR",
	17: "ARG_TYPE_LAMBDA",
	-1: "__SignatureArgumentKind__switch_must_have_a_default__",
}

var SignatureArgumentKind_value = map[string]int32{
	"ARG_TYPE_FIXED":          0,
	"ARG_TYPE_ANY_1":          1,
	"ARG_TYPE_ANY_2":          2,
	"ARG_ARRAY_TYPE_ANY_1":    3,
	"ARG_ARRAY_TYPE_ANY_2":    4,
	"ARG_PROTO_MAP_ANY":       14,
	"ARG_PROTO_MAP_KEY_ANY":   15,
	"ARG_PROTO_MAP_VALUE_ANY": 16,
	"ARG_PROTO_ANY":           5,
	"ARG_STRUCT_ANY":          6,
	"ARG_ENUM_ANY":            7,
	"ARG_TYPE_ARBITRARY":      8,
	"ARG_TYPE_RELATION":       9,
	"ARG_TYPE_VOID":           10,
	"ARG_TYPE_MODEL":          11,
	"ARG_TYPE_CONNECTION":     12,
	"ARG_TYPE_DESCRIPTOR":     13,
	"ARG_TYPE_LAMBDA":         17,
	"__SignatureArgumentKind__switch_must_have_a_default__": -1,
}

func (x SignatureArgumentKind) Enum() *SignatureArgumentKind {
	p := new(SignatureArgumentKind)
	*p = x
	return p
}

func (x SignatureArgumentKind) String() string {
	return proto.EnumName(SignatureArgumentKind_name, int32(x))
}

func (x *SignatureArgumentKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SignatureArgumentKind_value, data, "SignatureArgumentKind")
	if err != nil {
		return err
	}
	*x = SignatureArgumentKind(value)
	return nil
}

func (SignatureArgumentKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0}
}

type FunctionEnums_ArgumentCardinality int32

const (
	FunctionEnums_REQUIRED FunctionEnums_ArgumentCardinality = 0
	FunctionEnums_REPEATED FunctionEnums_ArgumentCardinality = 1
	FunctionEnums_OPTIONAL FunctionEnums_ArgumentCardinality = 2
)

var FunctionEnums_ArgumentCardinality_name = map[int32]string{
	0: "REQUIRED",
	1: "REPEATED",
	2: "OPTIONAL",
}

var FunctionEnums_ArgumentCardinality_value = map[string]int32{
	"REQUIRED": 0,
	"REPEATED": 1,
	"OPTIONAL": 2,
}

func (x FunctionEnums_ArgumentCardinality) Enum() *FunctionEnums_ArgumentCardinality {
	p := new(FunctionEnums_ArgumentCardinality)
	*p = x
	return p
}

func (x FunctionEnums_ArgumentCardinality) String() string {
	return proto.EnumName(FunctionEnums_ArgumentCardinality_name, int32(x))
}

func (x *FunctionEnums_ArgumentCardinality) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FunctionEnums_ArgumentCardinality_value, data, "FunctionEnums_ArgumentCardinality")
	if err != nil {
		return err
	}
	*x = FunctionEnums_ArgumentCardinality(value)
	return nil
}

func (FunctionEnums_ArgumentCardinality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0, 0}
}

type FunctionEnums_ProcedureArgumentMode int32

const (
	FunctionEnums_NOT_SET FunctionEnums_ProcedureArgumentMode = 0
	FunctionEnums_IN      FunctionEnums_ProcedureArgumentMode = 1
	FunctionEnums_OUT     FunctionEnums_ProcedureArgumentMode = 2
	FunctionEnums_INOUT   FunctionEnums_ProcedureArgumentMode = 3
)

var FunctionEnums_ProcedureArgumentMode_name = map[int32]string{
	0: "NOT_SET",
	1: "IN",
	2: "OUT",
	3: "INOUT",
}

var FunctionEnums_ProcedureArgumentMode_value = map[string]int32{
	"NOT_SET": 0,
	"IN":      1,
	"OUT":     2,
	"INOUT":   3,
}

func (x FunctionEnums_ProcedureArgumentMode) Enum() *FunctionEnums_ProcedureArgumentMode {
	p := new(FunctionEnums_ProcedureArgumentMode)
	*p = x
	return p
}

func (x FunctionEnums_ProcedureArgumentMode) String() string {
	return proto.EnumName(FunctionEnums_ProcedureArgumentMode_name, int32(x))
}

func (x *FunctionEnums_ProcedureArgumentMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FunctionEnums_ProcedureArgumentMode_value, data, "FunctionEnums_ProcedureArgumentMode")
	if err != nil {
		return err
	}
	*x = FunctionEnums_ProcedureArgumentMode(value)
	return nil
}

func (FunctionEnums_ProcedureArgumentMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0, 1}
}

type FunctionEnums_WindowOrderSupport int32

const (
	FunctionEnums_ORDER_UNSUPPORTED FunctionEnums_WindowOrderSupport = 0
	FunctionEnums_ORDER_OPTIONAL    FunctionEnums_WindowOrderSupport = 1
	FunctionEnums_ORDER_REQUIRED    FunctionEnums_WindowOrderSupport = 2
)

var FunctionEnums_WindowOrderSupport_name = map[int32]string{
	0: "ORDER_UNSUPPORTED",
	1: "ORDER_OPTIONAL",
	2: "ORDER_REQUIRED",
}

var FunctionEnums_WindowOrderSupport_value = map[string]int32{
	"ORDER_UNSUPPORTED": 0,
	"ORDER_OPTIONAL":    1,
	"ORDER_REQUIRED":    2,
}

func (x FunctionEnums_WindowOrderSupport) Enum() *FunctionEnums_WindowOrderSupport {
	p := new(FunctionEnums_WindowOrderSupport)
	*p = x
	return p
}

func (x FunctionEnums_WindowOrderSupport) String() string {
	return proto.EnumName(FunctionEnums_WindowOrderSupport_name, int32(x))
}

func (x *FunctionEnums_WindowOrderSupport) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FunctionEnums_WindowOrderSupport_value, data, "FunctionEnums_WindowOrderSupport")
	if err != nil {
		return err
	}
	*x = FunctionEnums_WindowOrderSupport(value)
	return nil
}

func (FunctionEnums_WindowOrderSupport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0, 2}
}

type FunctionEnums_Mode int32

const (
	FunctionEnums_SCALAR    FunctionEnums_Mode = 1
	FunctionEnums_AGGREGATE FunctionEnums_Mode = 2
	FunctionEnums_ANALYTIC  FunctionEnums_Mode = 3
)

var FunctionEnums_Mode_name = map[int32]string{
	1: "SCALAR",
	2: "AGGREGATE",
	3: "ANALYTIC",
}

var FunctionEnums_Mode_value = map[string]int32{
	"SCALAR":    1,
	"AGGREGATE": 2,
	"ANALYTIC":  3,
}

func (x FunctionEnums_Mode) Enum() *FunctionEnums_Mode {
	p := new(FunctionEnums_Mode)
	*p = x
	return p
}

func (x FunctionEnums_Mode) String() string {
	return proto.EnumName(FunctionEnums_Mode_name, int32(x))
}

func (x *FunctionEnums_Mode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FunctionEnums_Mode_value, data, "FunctionEnums_Mode")
	if err != nil {
		return err
	}
	*x = FunctionEnums_Mode(value)
	return nil
}

func (FunctionEnums_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0, 3}
}

type FunctionEnums_Volatility int32

const (
	FunctionEnums_IMMUTABLE FunctionEnums_Volatility = 0
	FunctionEnums_STABLE    FunctionEnums_Volatility = 1
	FunctionEnums_VOLATILE  FunctionEnums_Volatility = 2
)

var FunctionEnums_Volatility_name = map[int32]string{
	0: "IMMUTABLE",
	1: "STABLE",
	2: "VOLATILE",
}

var FunctionEnums_Volatility_value = map[string]int32{
	"IMMUTABLE": 0,
	"STABLE":    1,
	"VOLATILE":  2,
}

func (x FunctionEnums_Volatility) Enum() *FunctionEnums_Volatility {
	p := new(FunctionEnums_Volatility)
	*p = x
	return p
}

func (x FunctionEnums_Volatility) String() string {
	return proto.EnumName(FunctionEnums_Volatility_name, int32(x))
}

func (x *FunctionEnums_Volatility) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FunctionEnums_Volatility_value, data, "FunctionEnums_Volatility")
	if err != nil {
		return err
	}
	*x = FunctionEnums_Volatility(value)
	return nil
}

func (FunctionEnums_Volatility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0, 4}
}

type FunctionEnums_TableValuedFunctionType int32

const (
	FunctionEnums_INVALID                                                     FunctionEnums_TableValuedFunctionType = 0
	FunctionEnums_FIXED_OUTPUT_SCHEMA_TVF                                     FunctionEnums_TableValuedFunctionType = 1
	FunctionEnums_FORWARD_INPUT_SCHEMA_TO_OUTPUT_SCHEMA_TVF                   FunctionEnums_TableValuedFunctionType = 2
	FunctionEnums_TEMPLATED_SQL_TVF                                           FunctionEnums_TableValuedFunctionType = 3
	FunctionEnums_FORWARD_INPUT_SCHEMA_TO_OUTPUT_SCHEMA_WITH_APPENDED_COLUMNS FunctionEnums_TableValuedFunctionType = 7
)

var FunctionEnums_TableValuedFunctionType_name = map[int32]string{
	0: "INVALID",
	1: "FIXED_OUTPUT_SCHEMA_TVF",
	2: "FORWARD_INPUT_SCHEMA_TO_OUTPUT_SCHEMA_TVF",
	3: "TEMPLATED_SQL_TVF",
	7: "FORWARD_INPUT_SCHEMA_TO_OUTPUT_SCHEMA_WITH_APPENDED_COLUMNS",
}

var FunctionEnums_TableValuedFunctionType_value = map[string]int32{
	"INVALID":                 0,
	"FIXED_OUTPUT_SCHEMA_TVF": 1,
	"FORWARD_INPUT_SCHEMA_TO_OUTPUT_SCHEMA_TVF":                   2,
	"TEMPLATED_SQL_TVF":                                           3,
	"FORWARD_INPUT_SCHEMA_TO_OUTPUT_SCHEMA_WITH_APPENDED_COLUMNS": 7,
}

func (x FunctionEnums_TableValuedFunctionType) Enum() *FunctionEnums_TableValuedFunctionType {
	p := new(FunctionEnums_TableValuedFunctionType)
	*p = x
	return p
}

func (x FunctionEnums_TableValuedFunctionType) String() string {
	return proto.EnumName(FunctionEnums_TableValuedFunctionType_name, int32(x))
}

func (x *FunctionEnums_TableValuedFunctionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FunctionEnums_TableValuedFunctionType_value, data, "FunctionEnums_TableValuedFunctionType")
	if err != nil {
		return err
	}
	*x = FunctionEnums_TableValuedFunctionType(value)
	return nil
}

func (FunctionEnums_TableValuedFunctionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0, 5}
}

type FunctionEnums struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionEnums) Reset()         { *m = FunctionEnums{} }
func (m *FunctionEnums) String() string { return proto.CompactTextString(m) }
func (*FunctionEnums) ProtoMessage()    {}
func (*FunctionEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_b53802616b9fed8e, []int{0}
}

func (m *FunctionEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionEnums.Unmarshal(m, b)
}
func (m *FunctionEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionEnums.Marshal(b, m, deterministic)
}
func (m *FunctionEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionEnums.Merge(m, src)
}
func (m *FunctionEnums) XXX_Size() int {
	return xxx_messageInfo_FunctionEnums.Size(m)
}
func (m *FunctionEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionEnums.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionEnums proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("zetasql.SignatureArgumentKind", SignatureArgumentKind_name, SignatureArgumentKind_value)
	proto.RegisterEnum("zetasql.FunctionEnums_ArgumentCardinality", FunctionEnums_ArgumentCardinality_name, FunctionEnums_ArgumentCardinality_value)
	proto.RegisterEnum("zetasql.FunctionEnums_ProcedureArgumentMode", FunctionEnums_ProcedureArgumentMode_name, FunctionEnums_ProcedureArgumentMode_value)
	proto.RegisterEnum("zetasql.FunctionEnums_WindowOrderSupport", FunctionEnums_WindowOrderSupport_name, FunctionEnums_WindowOrderSupport_value)
	proto.RegisterEnum("zetasql.FunctionEnums_Mode", FunctionEnums_Mode_name, FunctionEnums_Mode_value)
	proto.RegisterEnum("zetasql.FunctionEnums_Volatility", FunctionEnums_Volatility_name, FunctionEnums_Volatility_value)
	proto.RegisterEnum("zetasql.FunctionEnums_TableValuedFunctionType", FunctionEnums_TableValuedFunctionType_name, FunctionEnums_TableValuedFunctionType_value)
	proto.RegisterType((*FunctionEnums)(nil), "zetasql.FunctionEnums")
}

func init() { proto.RegisterFile("zetasql/public/function.proto", fileDescriptor_b53802616b9fed8e) }

var fileDescriptor_b53802616b9fed8e = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x4e, 0xdb, 0x3e,
	0x14, 0xc6, 0x69, 0x0a, 0x94, 0x1e, 0x28, 0x18, 0xf3, 0xef, 0x9f, 0x4d, 0xd3, 0x6e, 0x7a, 0x31,
	0x69, 0x93, 0x06, 0x1a, 0x12, 0x57, 0xbb, 0x60, 0x6e, 0xe2, 0x96, 0x88, 0x24, 0x0e, 0x8e, 0x53,
	0xd6, 0xdd, 0x58, 0xa1, 0x09, 0x25, 0x52, 0x9a, 0x74, 0x69, 0x32, 0xc4, 0x9e, 0x65, 0x4f, 0xb4,
	0xa7, 0xd9, 0x1b, 0x6c, 0x8a, 0x4b, 0x0b, 0x95, 0x98, 0xb4, 0x5c, 0xd9, 0xbf, 0xcf, 0x39, 0xf9,
	0xce, 0x77, 0xac, 0xc0, 0xeb, 0xef, 0x51, 0x11, 0xcc, 0xbe, 0x26, 0xc7, 0xd3, 0xf2, 0x3a, 0x89,
	0x47, 0xc7, 0x37, 0x65, 0x3a, 0x2a, 0xe2, 0x2c, 0x3d, 0x9a, 0xe6, 0x59, 0x91, 0xe1, 0xc6, 0x83,
	0xdc, 0xf9, 0xb1, 0x0e, 0xad, 0xde, 0x83, 0x46, 0xd3, 0x72, 0x32, 0xeb, 0x9c, 0xc1, 0x01, 0xc9,
	0xc7, 0xe5, 0x24, 0x4a, 0x0b, 0x3d, 0xc8, 0xc3, 0x38, 0x0d, 0x92, 0xb8, 0xb8, 0xc7, 0x3b, 0xb0,
	0xc5, 0xe9, 0xa5, 0x6f, 0x72, 0x6a, 0xa0, 0xb5, 0xf9, 0xce, 0xa5, 0x44, 0x50, 0x03, 0xd5, 0xaa,
	0x1d, 0x73, 0x85, 0xc9, 0x1c, 0x62, 0x21, 0xad, 0xf3, 0x09, 0xda, 0x6e, 0x9e, 0x8d, 0xa2, 0xb0,
	0xcc, 0xa3, 0x45, 0x25, 0x3b, 0x0b, 0x23, 0xbc, 0x0d, 0x0d, 0x87, 0x09, 0xe9, 0x51, 0x81, 0xd6,
	0xf0, 0x26, 0x68, 0xa6, 0x83, 0x6a, 0xb8, 0x01, 0x75, 0xe6, 0x0b, 0xa4, 0xe1, 0x26, 0x6c, 0x98,
	0x4e, 0xb5, 0xac, 0x77, 0x3c, 0xc0, 0x57, 0x71, 0x1a, 0x66, 0x77, 0x2c, 0x0f, 0xa3, 0xdc, 0x2b,
	0xa7, 0xd3, 0x2c, 0x2f, 0x70, 0x1b, 0xf6, 0x19, 0x37, 0x28, 0x97, 0xbe, 0xe3, 0xf9, 0xae, 0xcb,
	0xb8, 0x50, 0x56, 0x30, 0xec, 0xce, 0xf1, 0xd2, 0x42, 0xed, 0x91, 0x2d, 0x2d, 0x6b, 0x9d, 0x63,
	0x58, 0x57, 0x2e, 0x00, 0x36, 0x3d, 0x9d, 0x58, 0x84, 0xa3, 0x1a, 0x6e, 0x41, 0x93, 0xf4, 0xfb,
	0x9c, 0xf6, 0x89, 0xa0, 0x48, 0xab, 0xfa, 0x20, 0x0e, 0xb1, 0x86, 0xc2, 0xd4, 0x51, 0xbd, 0x73,
	0x0a, 0x30, 0xc8, 0x92, 0xa0, 0x88, 0x55, 0xff, 0x2d, 0x68, 0x9a, 0xb6, 0xed, 0x0b, 0xd2, 0xb5,
	0x28, 0x5a, 0x53, 0x55, 0xe6, 0x6b, 0xd5, 0xfe, 0x80, 0x59, 0x44, 0x98, 0x16, 0x45, 0x5a, 0xe7,
	0x67, 0x0d, 0x0e, 0x45, 0x70, 0x9d, 0x44, 0x83, 0x20, 0x29, 0xa3, 0x70, 0x11, 0xae, 0xb8, 0x9f,
	0xaa, 0x04, 0x4c, 0x67, 0x40, 0x2c, 0xb3, 0x32, 0xfe, 0x0a, 0x0e, 0x7b, 0xe6, 0x67, 0x6a, 0x48,
	0xe6, 0x0b, 0xd7, 0x17, 0xd2, 0xd3, 0xcf, 0xa9, 0x4d, 0xa4, 0x18, 0xf4, 0x50, 0x0d, 0xbf, 0x87,
	0xb7, 0x3d, 0xc6, 0xaf, 0x08, 0x37, 0xa4, 0xe9, 0x3c, 0x55, 0xd9, 0x33, 0xc7, 0xb5, 0x2a, 0x1b,
	0x41, 0x6d, 0xd7, 0xaa, 0x06, 0x22, 0xbd, 0x4b, 0x4b, 0xe1, 0x3a, 0x3e, 0x83, 0x8f, 0xff, 0x56,
	0xe5, 0xca, 0x14, 0xe7, 0x92, 0xb8, 0x2e, 0x75, 0x0c, 0x6a, 0x48, 0x9d, 0x59, 0xbe, 0xed, 0x78,
	0xa8, 0xf1, 0xee, 0x57, 0x1d, 0xda, 0x5e, 0x3c, 0x4e, 0x83, 0xe2, 0xc9, 0x30, 0x2f, 0xe2, 0x34,
	0xac, 0x22, 0x26, 0xbc, 0x2f, 0xc5, 0xd0, 0xa5, 0x52, 0xb5, 0x31, 0x1f, 0xc5, 0x92, 0x11, 0x67,
	0x28, 0x3f, 0xcc, 0x47, 0xb1, 0xc2, 0x4e, 0x90, 0x86, 0x5f, 0xc0, 0x7f, 0x15, 0x23, 0x9c, 0x93,
	0xe1, 0xd3, 0xd3, 0xf5, 0xbf, 0x28, 0x27, 0x68, 0xbd, 0xea, 0xb0, 0x52, 0x5c, 0xce, 0x04, 0x93,
	0x36, 0x71, 0x2b, 0x01, 0xed, 0xe2, 0x97, 0xd0, 0x5e, 0xc5, 0x17, 0x74, 0xa8, 0xa4, 0xbd, 0x2a,
	0xdf, 0x55, 0x69, 0x40, 0x2c, 0x5f, 0x15, 0x44, 0x08, 0xef, 0x43, 0xeb, 0x51, 0xac, 0xd0, 0xc6,
	0xc2, 0xa9, 0x27, 0xb8, 0xaf, 0x0b, 0xc5, 0x36, 0x31, 0x82, 0x9d, 0x8a, 0x51, 0xc7, 0xb7, 0x15,
	0x69, 0xe0, 0xff, 0x01, 0x3f, 0xf6, 0xc3, 0xbb, 0xa6, 0xe0, 0x84, 0x0f, 0xd1, 0xd6, 0xc2, 0x9f,
	0xe2, 0x9c, 0x56, 0xb7, 0x81, 0x39, 0xa8, 0xb9, 0xf8, 0x8e, 0xc2, 0x03, 0x66, 0x1a, 0x08, 0x56,
	0x12, 0xb1, 0x99, 0x41, 0x2d, 0xb4, 0x8d, 0x0f, 0xe1, 0x60, 0xc9, 0x74, 0xe6, 0x38, 0x54, 0x57,
	0xef, 0xef, 0xac, 0x08, 0x06, 0xf5, 0x74, 0x6e, 0xba, 0x82, 0x71, 0xd4, 0xc2, 0x07, 0xb0, 0xb7,
	0x14, 0x2c, 0x62, 0x77, 0x0d, 0x82, 0xf6, 0x71, 0x17, 0x4e, 0xa5, 0x7c, 0x76, 0x5e, 0x52, 0xce,
	0xee, 0xe2, 0x62, 0x74, 0x2b, 0x27, 0xe5, 0xac, 0x90, 0xb7, 0xc1, 0xb7, 0x48, 0x06, 0x32, 0x8c,
	0x6e, 0x82, 0x32, 0x29, 0xa4, 0x44, 0xbf, 0x17, 0x4f, 0xad, 0xfb, 0x06, 0xf0, 0x28, 0x9b, 0x1c,
	0x8d, 0xb3, 0x6c, 0x9c, 0x44, 0x47, 0x0f, 0xff, 0x89, 0x2e, 0xfa, 0x12, 0x15, 0x81, 0x77, 0x69,
	0x2d, 0xae, 0xf3, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0xdc, 0xad, 0x35, 0x62, 0x04,
	0x00, 0x00,
}
