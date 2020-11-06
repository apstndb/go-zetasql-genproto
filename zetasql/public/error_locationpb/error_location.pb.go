// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zetasql/public/error_location.proto

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

type ErrorLocation struct {
	Line                 *int32         `protobuf:"varint,1,opt,name=line,def=1" json:"line,omitempty"`
	Column               *int32         `protobuf:"varint,2,opt,name=column,def=1" json:"column,omitempty"`
	Filename             *string        `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	ErrorSource          []*ErrorSource `protobuf:"bytes,4,rep,name=error_source,json=errorSource" json:"error_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ErrorLocation) Reset()         { *m = ErrorLocation{} }
func (m *ErrorLocation) String() string { return proto.CompactTextString(m) }
func (*ErrorLocation) ProtoMessage()    {}
func (*ErrorLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2518f6889f67388c, []int{0}
}

func (m *ErrorLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorLocation.Unmarshal(m, b)
}
func (m *ErrorLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorLocation.Marshal(b, m, deterministic)
}
func (m *ErrorLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorLocation.Merge(m, src)
}
func (m *ErrorLocation) XXX_Size() int {
	return xxx_messageInfo_ErrorLocation.Size(m)
}
func (m *ErrorLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorLocation.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorLocation proto.InternalMessageInfo

const Default_ErrorLocation_Line int32 = 1
const Default_ErrorLocation_Column int32 = 1

func (m *ErrorLocation) GetLine() int32 {
	if m != nil && m.Line != nil {
		return *m.Line
	}
	return Default_ErrorLocation_Line
}

func (m *ErrorLocation) GetColumn() int32 {
	if m != nil && m.Column != nil {
		return *m.Column
	}
	return Default_ErrorLocation_Column
}

func (m *ErrorLocation) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *ErrorLocation) GetErrorSource() []*ErrorSource {
	if m != nil {
		return m.ErrorSource
	}
	return nil
}

type ErrorSource struct {
	ErrorMessage            *string        `protobuf:"bytes,1,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	ErrorMessageCaretString *string        `protobuf:"bytes,2,opt,name=error_message_caret_string,json=errorMessageCaretString" json:"error_message_caret_string,omitempty"`
	ErrorLocation           *ErrorLocation `protobuf:"bytes,3,opt,name=error_location,json=errorLocation" json:"error_location,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}       `json:"-"`
	XXX_unrecognized        []byte         `json:"-"`
	XXX_sizecache           int32          `json:"-"`
}

func (m *ErrorSource) Reset()         { *m = ErrorSource{} }
func (m *ErrorSource) String() string { return proto.CompactTextString(m) }
func (*ErrorSource) ProtoMessage()    {}
func (*ErrorSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_2518f6889f67388c, []int{1}
}

func (m *ErrorSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorSource.Unmarshal(m, b)
}
func (m *ErrorSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorSource.Marshal(b, m, deterministic)
}
func (m *ErrorSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorSource.Merge(m, src)
}
func (m *ErrorSource) XXX_Size() int {
	return xxx_messageInfo_ErrorSource.Size(m)
}
func (m *ErrorSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorSource.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorSource proto.InternalMessageInfo

func (m *ErrorSource) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func (m *ErrorSource) GetErrorMessageCaretString() string {
	if m != nil && m.ErrorMessageCaretString != nil {
		return *m.ErrorMessageCaretString
	}
	return ""
}

func (m *ErrorSource) GetErrorLocation() *ErrorLocation {
	if m != nil {
		return m.ErrorLocation
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorLocation)(nil), "zetasql.ErrorLocation")
	proto.RegisterType((*ErrorSource)(nil), "zetasql.ErrorSource")
}

func init() {
	proto.RegisterFile("zetasql/public/error_location.proto", fileDescriptor_2518f6889f67388c)
}

var fileDescriptor_2518f6889f67388c = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x65, 0x5a, 0xfe, 0xe4, 0x42, 0x19, 0xac, 0x02, 0xa1, 0x53, 0xd4, 0x2e, 0x99, 0x52,
	0xd1, 0x05, 0x09, 0xc4, 0x02, 0x62, 0x83, 0xc5, 0x7d, 0x80, 0x28, 0x58, 0x47, 0x64, 0xc9, 0xf1,
	0x15, 0xdb, 0x59, 0x78, 0x10, 0x1e, 0x83, 0x67, 0x44, 0xb5, 0x13, 0x68, 0xc6, 0xef, 0xbe, 0xef,
	0xee, 0x7e, 0x77, 0xb0, 0xfa, 0x42, 0x5f, 0xbb, 0x4f, 0xbd, 0xde, 0x75, 0xef, 0x5a, 0xc9, 0x35,
	0x5a, 0x4b, 0xb6, 0xd2, 0x24, 0x6b, 0xaf, 0xc8, 0x94, 0x3b, 0x4b, 0x9e, 0xf8, 0x69, 0x1f, 0x5a,
	0x7e, 0x33, 0x98, 0xbd, 0xec, 0x13, 0xaf, 0x7d, 0x80, 0x5f, 0xc2, 0x54, 0x2b, 0x83, 0x19, 0xcb,
	0x59, 0x71, 0x7c, 0xcf, 0x6e, 0x45, 0x90, 0xfc, 0x06, 0x4e, 0x24, 0xe9, 0xae, 0x35, 0xd9, 0xd1,
	0x60, 0xf4, 0x05, 0xbe, 0x80, 0xb3, 0x0f, 0xa5, 0xd1, 0xd4, 0x2d, 0x66, 0x93, 0x9c, 0x15, 0x89,
	0xf8, 0xd3, 0xfc, 0x0e, 0xce, 0x23, 0x80, 0xa3, 0xce, 0x4a, 0xcc, 0xa6, 0xf9, 0xa4, 0x48, 0x37,
	0xf3, 0xb2, 0xdf, 0x5f, 0x86, 0xdd, 0xdb, 0xe0, 0x89, 0x14, 0xff, 0xc5, 0xf2, 0x87, 0x41, 0x7a,
	0x60, 0xf2, 0x15, 0xcc, 0xe2, 0xa0, 0x16, 0x9d, 0xab, 0x9b, 0xc8, 0x97, 0x88, 0x38, 0xfd, 0x2d,
	0xd6, 0xf8, 0x03, 0x2c, 0x46, 0xa1, 0x4a, 0xd6, 0x16, 0x7d, 0xe5, 0xbc, 0x55, 0xa6, 0x09, 0xe0,
	0x89, 0xb8, 0x3e, 0xec, 0x78, 0xde, 0xfb, 0xdb, 0x60, 0xf3, 0x47, 0xb8, 0x18, 0xff, 0x2a, 0x1c,
	0x93, 0x6e, 0xae, 0xc6, 0xb0, 0xc3, 0xa3, 0x44, 0xe4, 0x19, 0xe4, 0xd3, 0x1c, 0xb8, 0xa4, 0xb6,
	0x6c, 0x88, 0x1a, 0x8d, 0x43, 0xcb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xeb, 0x31, 0x32,
	0x8e, 0x01, 0x00, 0x00,
}
