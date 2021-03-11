// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/carnot/planner/compilerpb/compiler_status.proto

package compilerpb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	planpb "pixielabs.ai/pixielabs/src/carnot/planpb"
	proto1 "pixielabs.ai/pixielabs/src/common/base/proto"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LineColError struct {
	Line    uint64 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Column  uint64 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *LineColError) Reset()      { *m = LineColError{} }
func (*LineColError) ProtoMessage() {}
func (*LineColError) Descriptor() ([]byte, []int) {
	return fileDescriptor_7758d5b44fb349f6, []int{0}
}
func (m *LineColError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LineColError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LineColError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LineColError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineColError.Merge(m, src)
}
func (m *LineColError) XXX_Size() int {
	return m.Size()
}
func (m *LineColError) XXX_DiscardUnknown() {
	xxx_messageInfo_LineColError.DiscardUnknown(m)
}

var xxx_messageInfo_LineColError proto.InternalMessageInfo

func (m *LineColError) GetLine() uint64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *LineColError) GetColumn() uint64 {
	if m != nil {
		return m.Column
	}
	return 0
}

func (m *LineColError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CompilerError struct {
	// Types that are valid to be assigned to Error:
	//	*CompilerError_LineColError
	Error isCompilerError_Error `protobuf_oneof:"error"`
}

func (m *CompilerError) Reset()      { *m = CompilerError{} }
func (*CompilerError) ProtoMessage() {}
func (*CompilerError) Descriptor() ([]byte, []int) {
	return fileDescriptor_7758d5b44fb349f6, []int{1}
}
func (m *CompilerError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompilerError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompilerError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompilerError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompilerError.Merge(m, src)
}
func (m *CompilerError) XXX_Size() int {
	return m.Size()
}
func (m *CompilerError) XXX_DiscardUnknown() {
	xxx_messageInfo_CompilerError.DiscardUnknown(m)
}

var xxx_messageInfo_CompilerError proto.InternalMessageInfo

type isCompilerError_Error interface {
	isCompilerError_Error()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type CompilerError_LineColError struct {
	LineColError *LineColError `protobuf:"bytes,1,opt,name=line_col_error,json=lineColError,proto3,oneof" json:"line_col_error,omitempty"`
}

func (*CompilerError_LineColError) isCompilerError_Error() {}

func (m *CompilerError) GetError() isCompilerError_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CompilerError) GetLineColError() *LineColError {
	if x, ok := m.GetError().(*CompilerError_LineColError); ok {
		return x.LineColError
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CompilerError) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CompilerError_LineColError)(nil),
	}
}

type CompilerResult struct {
	Status      *proto1.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	LogicalPlan *planpb.Plan   `protobuf:"bytes,2,opt,name=logical_plan,json=logicalPlan,proto3" json:"logical_plan,omitempty"`
}

func (m *CompilerResult) Reset()      { *m = CompilerResult{} }
func (*CompilerResult) ProtoMessage() {}
func (*CompilerResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7758d5b44fb349f6, []int{2}
}
func (m *CompilerResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompilerResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompilerResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompilerResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompilerResult.Merge(m, src)
}
func (m *CompilerResult) XXX_Size() int {
	return m.Size()
}
func (m *CompilerResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CompilerResult.DiscardUnknown(m)
}

var xxx_messageInfo_CompilerResult proto.InternalMessageInfo

func (m *CompilerResult) GetStatus() *proto1.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CompilerResult) GetLogicalPlan() *planpb.Plan {
	if m != nil {
		return m.LogicalPlan
	}
	return nil
}

type CompilerErrorGroup struct {
	Errors []*CompilerError `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (m *CompilerErrorGroup) Reset()      { *m = CompilerErrorGroup{} }
func (*CompilerErrorGroup) ProtoMessage() {}
func (*CompilerErrorGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_7758d5b44fb349f6, []int{3}
}
func (m *CompilerErrorGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompilerErrorGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompilerErrorGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompilerErrorGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompilerErrorGroup.Merge(m, src)
}
func (m *CompilerErrorGroup) XXX_Size() int {
	return m.Size()
}
func (m *CompilerErrorGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_CompilerErrorGroup.DiscardUnknown(m)
}

var xxx_messageInfo_CompilerErrorGroup proto.InternalMessageInfo

func (m *CompilerErrorGroup) GetErrors() []*CompilerError {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*LineColError)(nil), "pl.carnot.planner.compilerpb.LineColError")
	proto.RegisterType((*CompilerError)(nil), "pl.carnot.planner.compilerpb.CompilerError")
	proto.RegisterType((*CompilerResult)(nil), "pl.carnot.planner.compilerpb.CompilerResult")
	proto.RegisterType((*CompilerErrorGroup)(nil), "pl.carnot.planner.compilerpb.CompilerErrorGroup")
}

func init() {
	proto.RegisterFile("src/carnot/planner/compilerpb/compiler_status.proto", fileDescriptor_7758d5b44fb349f6)
}

var fileDescriptor_7758d5b44fb349f6 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbf, 0x6a, 0xe3, 0x30,
	0x18, 0xb7, 0x2e, 0x39, 0x87, 0x93, 0x73, 0x19, 0x74, 0x10, 0x4c, 0x08, 0x22, 0x78, 0x0a, 0x17,
	0x90, 0x21, 0x99, 0x6e, 0x3b, 0x12, 0x4a, 0x3b, 0x74, 0x28, 0x6a, 0x97, 0x76, 0x31, 0xb6, 0x11,
	0xc1, 0x20, 0x4b, 0x42, 0xb6, 0xa1, 0x63, 0x1f, 0xa1, 0x8f, 0xd1, 0x47, 0xe9, 0x98, 0x31, 0x63,
	0xe3, 0x2c, 0x1d, 0xf3, 0x08, 0xc5, 0xb2, 0xd3, 0xfc, 0x19, 0x32, 0xf9, 0xfb, 0x7d, 0xfe, 0x7e,
	0xff, 0x6c, 0x38, 0xcb, 0x74, 0xec, 0xc7, 0xa1, 0x16, 0x32, 0xf7, 0x15, 0x0f, 0x85, 0x60, 0xda,
	0x8f, 0x65, 0xaa, 0x12, 0xce, 0xb4, 0x8a, 0xbe, 0xc7, 0x20, 0xcb, 0xc3, 0xbc, 0xc8, 0x88, 0xd2,
	0x32, 0x97, 0x68, 0xa8, 0x38, 0xa9, 0x39, 0xa4, 0xe1, 0x90, 0x03, 0x67, 0xe0, 0x19, 0x49, 0x99,
	0xa6, 0x52, 0xf8, 0x51, 0x98, 0x31, 0xdf, 0x90, 0xfc, 0x63, 0x85, 0xc1, 0xf0, 0xcc, 0x56, 0x45,
	0xe6, 0x51, 0xbf, 0xf5, 0x1e, 0x60, 0xf7, 0x36, 0x11, 0x6c, 0x21, 0xf9, 0x95, 0xd6, 0x52, 0x23,
	0x04, 0xdb, 0x3c, 0x11, 0xcc, 0x05, 0x23, 0x30, 0x6e, 0x53, 0x33, 0xa3, 0x3e, 0xb4, 0x63, 0xc9,
	0x8b, 0x54, 0xb8, 0x3f, 0xcc, 0xb6, 0x41, 0xc8, 0x85, 0x9d, 0x94, 0x65, 0x59, 0xb8, 0x64, 0x6e,
	0x6b, 0x04, 0xc6, 0xbf, 0xe8, 0x1e, 0x7a, 0x1c, 0xfe, 0x5e, 0x34, 0x29, 0x6b, 0x59, 0x0a, 0x7b,
	0x95, 0x54, 0x10, 0x4b, 0x1e, 0xb0, 0x6a, 0x63, 0x0c, 0x9c, 0xe9, 0x5f, 0x72, 0xa9, 0x1f, 0x39,
	0x8e, 0x76, 0x63, 0xd1, 0x2e, 0x3f, 0xc2, 0xf3, 0x0e, 0xfc, 0x69, 0xa4, 0xbc, 0x67, 0xd8, 0xdb,
	0xbb, 0x51, 0x96, 0x15, 0x3c, 0x47, 0x13, 0x68, 0xd7, 0xdf, 0xa0, 0xb1, 0xf9, 0x53, 0xd9, 0xd4,
	0x1b, 0x15, 0x91, 0x7b, 0x33, 0xd0, 0xe6, 0x04, 0xfd, 0x83, 0x5d, 0x2e, 0x97, 0x49, 0x1c, 0xf2,
	0xa0, 0x8a, 0x60, 0x4a, 0x3a, 0xd3, 0xfe, 0x59, 0x32, 0x15, 0x91, 0x3b, 0x1e, 0x0a, 0xea, 0x34,
	0xb7, 0x15, 0xf0, 0x1e, 0x21, 0x3a, 0xe9, 0x79, 0xad, 0x65, 0xa1, 0xd0, 0x02, 0xda, 0x26, 0x58,
	0xe5, 0xde, 0x1a, 0x3b, 0xd3, 0xc9, 0xe5, 0x92, 0x27, 0x0a, 0xb4, 0xa1, 0xce, 0xff, 0xaf, 0x36,
	0xd8, 0x5a, 0x6f, 0xb0, 0xb5, 0xdb, 0x60, 0xf0, 0x52, 0x62, 0xf0, 0x56, 0x62, 0xf0, 0x5e, 0x62,
	0xb0, 0x2a, 0x31, 0xf8, 0x28, 0x31, 0xf8, 0x2c, 0xb1, 0xb5, 0x2b, 0x31, 0x78, 0xdd, 0x62, 0x6b,
	0xb5, 0xc5, 0xd6, 0x7a, 0x8b, 0xad, 0x27, 0x78, 0xd0, 0x8d, 0x6c, 0xf3, 0x87, 0x67, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf7, 0x3c, 0x1d, 0xb3, 0x78, 0x02, 0x00, 0x00,
}

func (this *LineColError) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LineColError)
	if !ok {
		that2, ok := that.(LineColError)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Line != that1.Line {
		return false
	}
	if this.Column != that1.Column {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *CompilerError) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CompilerError)
	if !ok {
		that2, ok := that.(CompilerError)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Error == nil {
		if this.Error != nil {
			return false
		}
	} else if this.Error == nil {
		return false
	} else if !this.Error.Equal(that1.Error) {
		return false
	}
	return true
}
func (this *CompilerError_LineColError) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CompilerError_LineColError)
	if !ok {
		that2, ok := that.(CompilerError_LineColError)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.LineColError.Equal(that1.LineColError) {
		return false
	}
	return true
}
func (this *CompilerResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CompilerResult)
	if !ok {
		that2, ok := that.(CompilerResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !this.LogicalPlan.Equal(that1.LogicalPlan) {
		return false
	}
	return true
}
func (this *CompilerErrorGroup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CompilerErrorGroup)
	if !ok {
		that2, ok := that.(CompilerErrorGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Errors) != len(that1.Errors) {
		return false
	}
	for i := range this.Errors {
		if !this.Errors[i].Equal(that1.Errors[i]) {
			return false
		}
	}
	return true
}
func (this *LineColError) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&compilerpb.LineColError{")
	s = append(s, "Line: "+fmt.Sprintf("%#v", this.Line)+",\n")
	s = append(s, "Column: "+fmt.Sprintf("%#v", this.Column)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CompilerError) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&compilerpb.CompilerError{")
	if this.Error != nil {
		s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CompilerError_LineColError) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&compilerpb.CompilerError_LineColError{` +
		`LineColError:` + fmt.Sprintf("%#v", this.LineColError) + `}`}, ", ")
	return s
}
func (this *CompilerResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&compilerpb.CompilerResult{")
	if this.Status != nil {
		s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	}
	if this.LogicalPlan != nil {
		s = append(s, "LogicalPlan: "+fmt.Sprintf("%#v", this.LogicalPlan)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CompilerErrorGroup) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&compilerpb.CompilerErrorGroup{")
	if this.Errors != nil {
		s = append(s, "Errors: "+fmt.Sprintf("%#v", this.Errors)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCompilerStatus(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *LineColError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LineColError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LineColError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintCompilerStatus(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Column != 0 {
		i = encodeVarintCompilerStatus(dAtA, i, uint64(m.Column))
		i--
		dAtA[i] = 0x10
	}
	if m.Line != 0 {
		i = encodeVarintCompilerStatus(dAtA, i, uint64(m.Line))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CompilerError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompilerError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompilerError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		{
			size := m.Error.Size()
			i -= size
			if _, err := m.Error.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *CompilerError_LineColError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompilerError_LineColError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LineColError != nil {
		{
			size, err := m.LineColError.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCompilerStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *CompilerResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompilerResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompilerResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LogicalPlan != nil {
		{
			size, err := m.LogicalPlan.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCompilerStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCompilerStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CompilerErrorGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompilerErrorGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompilerErrorGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for iNdEx := len(m.Errors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Errors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCompilerStatus(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCompilerStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovCompilerStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LineColError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Line != 0 {
		n += 1 + sovCompilerStatus(uint64(m.Line))
	}
	if m.Column != 0 {
		n += 1 + sovCompilerStatus(uint64(m.Column))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCompilerStatus(uint64(l))
	}
	return n
}

func (m *CompilerError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != nil {
		n += m.Error.Size()
	}
	return n
}

func (m *CompilerError_LineColError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LineColError != nil {
		l = m.LineColError.Size()
		n += 1 + l + sovCompilerStatus(uint64(l))
	}
	return n
}
func (m *CompilerResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovCompilerStatus(uint64(l))
	}
	if m.LogicalPlan != nil {
		l = m.LogicalPlan.Size()
		n += 1 + l + sovCompilerStatus(uint64(l))
	}
	return n
}

func (m *CompilerErrorGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for _, e := range m.Errors {
			l = e.Size()
			n += 1 + l + sovCompilerStatus(uint64(l))
		}
	}
	return n
}

func sovCompilerStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCompilerStatus(x uint64) (n int) {
	return sovCompilerStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LineColError) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LineColError{`,
		`Line:` + fmt.Sprintf("%v", this.Line) + `,`,
		`Column:` + fmt.Sprintf("%v", this.Column) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CompilerError) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CompilerError{`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CompilerError_LineColError) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CompilerError_LineColError{`,
		`LineColError:` + strings.Replace(fmt.Sprintf("%v", this.LineColError), "LineColError", "LineColError", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CompilerResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CompilerResult{`,
		`Status:` + strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "proto1.Status", 1) + `,`,
		`LogicalPlan:` + strings.Replace(fmt.Sprintf("%v", this.LogicalPlan), "Plan", "planpb.Plan", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CompilerErrorGroup) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForErrors := "[]*CompilerError{"
	for _, f := range this.Errors {
		repeatedStringForErrors += strings.Replace(f.String(), "CompilerError", "CompilerError", 1) + ","
	}
	repeatedStringForErrors += "}"
	s := strings.Join([]string{`&CompilerErrorGroup{`,
		`Errors:` + repeatedStringForErrors + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCompilerStatus(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LineColError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompilerStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LineColError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LineColError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			m.Line = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Line |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Column", wireType)
			}
			m.Column = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Column |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCompilerStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompilerError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompilerStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CompilerError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompilerError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LineColError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &LineColError{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Error = &CompilerError_LineColError{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCompilerStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompilerResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompilerStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CompilerResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompilerResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &proto1.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogicalPlan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LogicalPlan == nil {
				m.LogicalPlan = &planpb.Plan{}
			}
			if err := m.LogicalPlan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCompilerStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompilerErrorGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompilerStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CompilerErrorGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompilerErrorGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, &CompilerError{})
			if err := m.Errors[len(m.Errors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCompilerStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompilerStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCompilerStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCompilerStatus
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCompilerStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCompilerStatus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCompilerStatus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCompilerStatus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCompilerStatus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCompilerStatus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCompilerStatus = fmt.Errorf("proto: unexpected end of group")
)
