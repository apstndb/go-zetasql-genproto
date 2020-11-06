// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zetasql/public/function.proto

package public

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
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x4e, 0xe3, 0x46,
	0x14, 0x86, 0x89, 0x03, 0x84, 0x1c, 0x08, 0x0c, 0x43, 0x53, 0x5a, 0x55, 0xbd, 0x68, 0xee, 0x5a,
	0xa9, 0x89, 0x8a, 0xc4, 0x55, 0x2f, 0xe8, 0xc4, 0x9e, 0x84, 0x11, 0xb6, 0xc7, 0x8c, 0xc7, 0xa1,
	0xe1, 0x66, 0xe4, 0xc4, 0x26, 0x58, 0x72, 0xec, 0xd4, 0xb1, 0x8b, 0xd8, 0x67, 0xd9, 0x27, 0xda,
	0xa7, 0xd9, 0x37, 0xd8, 0x95, 0x1d, 0x12, 0x88, 0xc4, 0x4a, 0x9b, 0xab, 0x99, 0xef, 0x9f, 0x1c,
	0xff, 0xe7, 0x3f, 0xa3, 0x81, 0x5f, 0x3f, 0x84, 0xb9, 0xbf, 0xfc, 0x2f, 0xee, 0x2d, 0x8a, 0x49,
	0x1c, 0x4d, 0x7b, 0x0f, 0x45, 0x32, 0xcd, 0xa3, 0x34, 0xe9, 0x2e, 0xb2, 0x34, 0x4f, 0x71, 0xe3,
	0x45, 0xee, 0x7c, 0xdc, 0x85, 0xd6, 0xe0, 0x45, 0xa3, 0x49, 0x31, 0x5f, 0x76, 0xae, 0xe0, 0x8c,
	0x64, 0xb3, 0x62, 0x1e, 0x26, 0xb9, 0xee, 0x67, 0x41, 0x94, 0xf8, 0x71, 0x94, 0x3f, 0xe3, 0x23,
	0x38, 0x10, 0xf4, 0xd6, 0x63, 0x82, 0x1a, 0x68, 0x67, 0xb5, 0x73, 0x28, 0x91, 0xd4, 0x40, 0xb5,
	0x72, 0xc7, 0x1d, 0xc9, 0xb8, 0x4d, 0x4c, 0xa4, 0x75, 0xfe, 0x81, 0xb6, 0x93, 0xa5, 0xd3, 0x30,
	0x28, 0xb2, 0x70, 0x5d, 0xc9, 0x4a, 0x83, 0x10, 0x1f, 0x42, 0xc3, 0xe6, 0x52, 0xb9, 0x54, 0xa2,
	0x1d, 0xbc, 0x0f, 0x1a, 0xb3, 0x51, 0x0d, 0x37, 0xa0, 0xce, 0x3d, 0x89, 0x34, 0xdc, 0x84, 0x3d,
	0x66, 0x97, 0xcb, 0x7a, 0xc7, 0x05, 0x7c, 0x17, 0x25, 0x41, 0xfa, 0xc4, 0xb3, 0x20, 0xcc, 0xdc,
	0x62, 0xb1, 0x48, 0xb3, 0x1c, 0xb7, 0xe1, 0x94, 0x0b, 0x83, 0x0a, 0xe5, 0xd9, 0xae, 0xe7, 0x38,
	0x5c, 0xc8, 0xca, 0x0a, 0x86, 0xe3, 0x15, 0xde, 0x58, 0xa8, 0xbd, 0xb2, 0x8d, 0x65, 0xad, 0xd3,
	0x83, 0xdd, 0xca, 0x05, 0xc0, 0xbe, 0xab, 0x13, 0x93, 0x08, 0x54, 0xc3, 0x2d, 0x68, 0x92, 0xe1,
	0x50, 0xd0, 0x21, 0x91, 0x14, 0x69, 0x65, 0x1f, 0xc4, 0x26, 0xe6, 0x58, 0x32, 0x1d, 0xd5, 0x3b,
	0x97, 0x00, 0xa3, 0x34, 0xf6, 0xf3, 0xa8, 0xea, 0xbf, 0x05, 0x4d, 0x66, 0x59, 0x9e, 0x24, 0x7d,
	0x93, 0xa2, 0x9d, 0xaa, 0xca, 0x6a, 0x5d, 0xb5, 0x3f, 0xe2, 0x26, 0x91, 0xcc, 0xa4, 0x48, 0xeb,
	0x7c, 0xaa, 0xc1, 0xb9, 0xf4, 0x27, 0x71, 0x38, 0xf2, 0xe3, 0x22, 0x0c, 0xd6, 0xe1, 0xca, 0xe7,
	0x45, 0x95, 0x00, 0xb3, 0x47, 0xc4, 0x64, 0xa5, 0xf1, 0x5f, 0xe0, 0x7c, 0xc0, 0xfe, 0xa5, 0x86,
	0xe2, 0x9e, 0x74, 0x3c, 0xa9, 0x5c, 0xfd, 0x9a, 0x5a, 0x44, 0xc9, 0xd1, 0x00, 0xd5, 0xf0, 0x9f,
	0xf0, 0xfb, 0x80, 0x8b, 0x3b, 0x22, 0x0c, 0xc5, 0xec, 0xb7, 0x2a, 0x7f, 0xe7, 0xb8, 0x56, 0x66,
	0x23, 0xa9, 0xe5, 0x98, 0xe5, 0x40, 0x94, 0x7b, 0x6b, 0x56, 0xb8, 0x8e, 0xaf, 0xe0, 0xef, 0xef,
	0xab, 0x72, 0xc7, 0xe4, 0xb5, 0x22, 0x8e, 0x43, 0x6d, 0x83, 0x1a, 0x4a, 0xe7, 0xa6, 0x67, 0xd9,
	0x2e, 0x6a, 0xfc, 0xf1, 0xb9, 0x0e, 0x6d, 0x37, 0x9a, 0x25, 0x7e, 0xfe, 0x66, 0x98, 0x37, 0x51,
	0x12, 0x94, 0x11, 0x13, 0x31, 0x54, 0x72, 0xec, 0x50, 0x55, 0xb5, 0xb1, 0x1a, 0xc5, 0x86, 0x11,
	0x7b, 0xac, 0xfe, 0x5a, 0x8d, 0x62, 0x8b, 0x5d, 0x20, 0x0d, 0xff, 0x04, 0x3f, 0x94, 0x8c, 0x08,
	0x41, 0xc6, 0x6f, 0x4f, 0xd7, 0xbf, 0xa1, 0x5c, 0xa0, 0xdd, 0xb2, 0xc3, 0x52, 0x71, 0x04, 0x97,
	0x5c, 0x59, 0xc4, 0x29, 0x05, 0x74, 0x8c, 0x7f, 0x86, 0xf6, 0x36, 0xbe, 0xa1, 0xe3, 0x4a, 0x3a,
	0x29, 0xf3, 0xdd, 0x96, 0x46, 0xc4, 0xf4, 0xaa, 0x82, 0x08, 0xe1, 0x53, 0x68, 0xbd, 0x8a, 0x25,
	0xda, 0x5b, 0x3b, 0x75, 0xa5, 0xf0, 0x74, 0x59, 0xb1, 0x7d, 0x8c, 0xe0, 0xa8, 0x64, 0xd4, 0xf6,
	0xac, 0x8a, 0x34, 0xf0, 0x8f, 0x80, 0x5f, 0xfb, 0x11, 0x7d, 0x26, 0x05, 0x11, 0x63, 0x74, 0xb0,
	0xf6, 0x57, 0x71, 0x41, 0xcb, 0xdb, 0xc0, 0x6d, 0xd4, 0x5c, 0x7f, 0xa7, 0xc2, 0x23, 0xce, 0x0c,
	0x04, 0x5b, 0x89, 0x58, 0xdc, 0xa0, 0x26, 0x3a, 0xc4, 0xe7, 0x70, 0xb6, 0x61, 0x3a, 0xb7, 0x6d,
	0xaa, 0x57, 0xff, 0x3f, 0xda, 0x12, 0x0c, 0xea, 0xea, 0x82, 0x39, 0x92, 0x0b, 0xd4, 0xc2, 0x67,
	0x70, 0xb2, 0x11, 0x4c, 0x62, 0xf5, 0x0d, 0x82, 0x4e, 0x71, 0x1f, 0x2e, 0x95, 0x7a, 0x77, 0x5e,
	0x4a, 0x2d, 0x9f, 0xa2, 0x7c, 0xfa, 0xa8, 0xe6, 0xc5, 0x32, 0x57, 0x8f, 0xfe, 0xff, 0xa1, 0xf2,
	0x55, 0x10, 0x3e, 0xf8, 0x45, 0x9c, 0x2b, 0x85, 0xbe, 0xac, 0x7f, 0xb5, 0x3e, 0x03, 0x3c, 0x4d,
	0xe7, 0xdd, 0x59, 0x9a, 0xce, 0xe2, 0xb0, 0xfb, 0xf2, 0x4e, 0xf4, 0xd1, 0x7d, 0x98, 0xfb, 0xee,
	0xad, 0xb9, 0xbe, 0xce, 0xcb, 0xfb, 0xdf, 0x66, 0x51, 0xfe, 0x58, 0x4c, 0xba, 0xd3, 0x74, 0xde,
	0xf3, 0x17, 0xcb, 0x3c, 0x09, 0x26, 0xbd, 0xed, 0x47, 0xe7, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x3e, 0xc0, 0x00, 0x85, 0x04, 0x00, 0x00,
}
