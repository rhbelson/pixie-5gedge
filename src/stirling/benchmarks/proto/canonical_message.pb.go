// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/stirling/benchmarks/proto/canonical_message.proto

package stirlingpb

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Canonical struct {
	TimeStamp int64   `protobuf:"varint,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Data1     float64 `protobuf:"fixed64,2,opt,name=data1,proto3" json:"data1,omitempty"`
	Data2     int64   `protobuf:"varint,3,opt,name=data2,proto3" json:"data2,omitempty"`
}

func (m *Canonical) Reset()      { *m = Canonical{} }
func (*Canonical) ProtoMessage() {}
func (*Canonical) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28f452d22ee1157, []int{0}
}
func (m *Canonical) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Canonical) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Canonical.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Canonical) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Canonical.Merge(m, src)
}
func (m *Canonical) XXX_Size() int {
	return m.Size()
}
func (m *Canonical) XXX_DiscardUnknown() {
	xxx_messageInfo_Canonical.DiscardUnknown(m)
}

var xxx_messageInfo_Canonical proto.InternalMessageInfo

func (m *Canonical) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *Canonical) GetData1() float64 {
	if m != nil {
		return m.Data1
	}
	return 0
}

func (m *Canonical) GetData2() int64 {
	if m != nil {
		return m.Data2
	}
	return 0
}

type CanonicalRepeatedColumn struct {
	TimeStamp []int64   `protobuf:"varint,1,rep,packed,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Data1     []float64 `protobuf:"fixed64,2,rep,packed,name=data1,proto3" json:"data1,omitempty"`
	Data2     []int64   `protobuf:"varint,3,rep,packed,name=data2,proto3" json:"data2,omitempty"`
}

func (m *CanonicalRepeatedColumn) Reset()      { *m = CanonicalRepeatedColumn{} }
func (*CanonicalRepeatedColumn) ProtoMessage() {}
func (*CanonicalRepeatedColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28f452d22ee1157, []int{1}
}
func (m *CanonicalRepeatedColumn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalRepeatedColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalRepeatedColumn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalRepeatedColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalRepeatedColumn.Merge(m, src)
}
func (m *CanonicalRepeatedColumn) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalRepeatedColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalRepeatedColumn.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalRepeatedColumn proto.InternalMessageInfo

func (m *CanonicalRepeatedColumn) GetTimeStamp() []int64 {
	if m != nil {
		return m.TimeStamp
	}
	return nil
}

func (m *CanonicalRepeatedColumn) GetData1() []float64 {
	if m != nil {
		return m.Data1
	}
	return nil
}

func (m *CanonicalRepeatedColumn) GetData2() []int64 {
	if m != nil {
		return m.Data2
	}
	return nil
}

type CanonicalStream struct {
	DataStream []*Canonical `protobuf:"bytes,1,rep,name=data_stream,json=dataStream,proto3" json:"data_stream,omitempty"`
}

func (m *CanonicalStream) Reset()      { *m = CanonicalStream{} }
func (*CanonicalStream) ProtoMessage() {}
func (*CanonicalStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28f452d22ee1157, []int{2}
}
func (m *CanonicalStream) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CanonicalStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CanonicalStream.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CanonicalStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CanonicalStream.Merge(m, src)
}
func (m *CanonicalStream) XXX_Size() int {
	return m.Size()
}
func (m *CanonicalStream) XXX_DiscardUnknown() {
	xxx_messageInfo_CanonicalStream.DiscardUnknown(m)
}

var xxx_messageInfo_CanonicalStream proto.InternalMessageInfo

func (m *CanonicalStream) GetDataStream() []*Canonical {
	if m != nil {
		return m.DataStream
	}
	return nil
}

func init() {
	proto.RegisterType((*Canonical)(nil), "pl.stirling.stirlingpb.Canonical")
	proto.RegisterType((*CanonicalRepeatedColumn)(nil), "pl.stirling.stirlingpb.CanonicalRepeatedColumn")
	proto.RegisterType((*CanonicalStream)(nil), "pl.stirling.stirlingpb.CanonicalStream")
}

func init() {
	proto.RegisterFile("src/stirling/benchmarks/proto/canonical_message.proto", fileDescriptor_f28f452d22ee1157)
}

var fileDescriptor_f28f452d22ee1157 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xfd, 0x88, 0x40, 0xea, 0xeb, 0x80, 0x14, 0x21, 0xe8, 0xc2, 0x53, 0xe9, 0xd4, 0x29,
	0x11, 0x45, 0xec, 0xa8, 0xbd, 0x41, 0x2a, 0x18, 0x58, 0x22, 0x27, 0xb1, 0x4a, 0x44, 0x9c, 0x44,
	0xb1, 0xd9, 0x39, 0x02, 0xc7, 0xe0, 0x28, 0x8c, 0x19, 0x3b, 0x12, 0x67, 0x61, 0xec, 0x11, 0x50,
	0x12, 0x25, 0x19, 0x28, 0x9b, 0xff, 0xdf, 0xfe, 0xff, 0xcf, 0xef, 0xe1, 0xbd, 0x2a, 0x42, 0x57,
	0xe9, 0xb8, 0x48, 0xe2, 0x74, 0xe7, 0x06, 0x22, 0x0d, 0x5f, 0x24, 0x2f, 0x5e, 0x95, 0x9b, 0x17,
	0x99, 0xce, 0xdc, 0x90, 0xa7, 0x59, 0x1a, 0x87, 0x3c, 0xf1, 0xa5, 0x50, 0x8a, 0xef, 0x84, 0xd3,
	0xfa, 0xf6, 0x65, 0x9e, 0x38, 0x7d, 0x6a, 0x38, 0xe4, 0xc1, 0xe2, 0x09, 0x27, 0x9b, 0x3e, 0x62,
	0x5f, 0x23, 0xea, 0x58, 0x0a, 0x5f, 0x69, 0x2e, 0xf3, 0x19, 0xcc, 0x61, 0x69, 0x79, 0x93, 0xc6,
	0xd9, 0x36, 0x86, 0x7d, 0x81, 0xa7, 0x11, 0xd7, 0xfc, 0x76, 0x76, 0x32, 0x87, 0x25, 0x78, 0x9d,
	0xe8, 0xdd, 0xd5, 0xcc, 0x6a, 0xdf, 0x77, 0x62, 0x11, 0xe1, 0xd5, 0xd0, 0xeb, 0x89, 0x5c, 0x70,
	0x2d, 0xa2, 0x4d, 0x96, 0xbc, 0xc9, 0xf4, 0x0f, 0xc5, 0xfa, 0x97, 0x62, 0x1d, 0xa5, 0x58, 0x23,
	0xe5, 0x11, 0xcf, 0x07, 0xca, 0x56, 0x17, 0x82, 0x4b, 0x7b, 0x8d, 0xd3, 0xe6, 0xce, 0x57, 0xad,
	0x6c, 0xeb, 0xa7, 0xab, 0x1b, 0xe7, 0xf8, 0xf8, 0xce, 0xf8, 0x47, 0x6c, 0x52, 0x5d, 0xc7, 0xfa,
	0xa1, 0xac, 0x88, 0xed, 0x2b, 0x62, 0x87, 0x8a, 0xe0, 0xdd, 0x10, 0x7c, 0x1a, 0x82, 0x2f, 0x43,
	0x50, 0x1a, 0x82, 0x6f, 0x43, 0xf0, 0x63, 0x88, 0x1d, 0x0c, 0xc1, 0x47, 0x4d, 0xac, 0xac, 0x89,
	0xed, 0x6b, 0x62, 0xcf, 0x38, 0xf6, 0x06, 0x67, 0xed, 0xd6, 0xef, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x95, 0x97, 0x7c, 0xad, 0xae, 0x01, 0x00, 0x00,
}

func (this *Canonical) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Canonical)
	if !ok {
		that2, ok := that.(Canonical)
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
	if this.TimeStamp != that1.TimeStamp {
		return false
	}
	if this.Data1 != that1.Data1 {
		return false
	}
	if this.Data2 != that1.Data2 {
		return false
	}
	return true
}
func (this *CanonicalRepeatedColumn) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CanonicalRepeatedColumn)
	if !ok {
		that2, ok := that.(CanonicalRepeatedColumn)
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
	if len(this.TimeStamp) != len(that1.TimeStamp) {
		return false
	}
	for i := range this.TimeStamp {
		if this.TimeStamp[i] != that1.TimeStamp[i] {
			return false
		}
	}
	if len(this.Data1) != len(that1.Data1) {
		return false
	}
	for i := range this.Data1 {
		if this.Data1[i] != that1.Data1[i] {
			return false
		}
	}
	if len(this.Data2) != len(that1.Data2) {
		return false
	}
	for i := range this.Data2 {
		if this.Data2[i] != that1.Data2[i] {
			return false
		}
	}
	return true
}
func (this *CanonicalStream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CanonicalStream)
	if !ok {
		that2, ok := that.(CanonicalStream)
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
	if len(this.DataStream) != len(that1.DataStream) {
		return false
	}
	for i := range this.DataStream {
		if !this.DataStream[i].Equal(that1.DataStream[i]) {
			return false
		}
	}
	return true
}
func (this *Canonical) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&stirlingpb.Canonical{")
	s = append(s, "TimeStamp: "+fmt.Sprintf("%#v", this.TimeStamp)+",\n")
	s = append(s, "Data1: "+fmt.Sprintf("%#v", this.Data1)+",\n")
	s = append(s, "Data2: "+fmt.Sprintf("%#v", this.Data2)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CanonicalRepeatedColumn) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&stirlingpb.CanonicalRepeatedColumn{")
	s = append(s, "TimeStamp: "+fmt.Sprintf("%#v", this.TimeStamp)+",\n")
	s = append(s, "Data1: "+fmt.Sprintf("%#v", this.Data1)+",\n")
	s = append(s, "Data2: "+fmt.Sprintf("%#v", this.Data2)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CanonicalStream) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&stirlingpb.CanonicalStream{")
	if this.DataStream != nil {
		s = append(s, "DataStream: "+fmt.Sprintf("%#v", this.DataStream)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCanonicalMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Canonical) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Canonical) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Canonical) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data2 != 0 {
		i = encodeVarintCanonicalMessage(dAtA, i, uint64(m.Data2))
		i--
		dAtA[i] = 0x18
	}
	if m.Data1 != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Data1))))
		i--
		dAtA[i] = 0x11
	}
	if m.TimeStamp != 0 {
		i = encodeVarintCanonicalMessage(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalRepeatedColumn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalRepeatedColumn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalRepeatedColumn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data2) > 0 {
		dAtA2 := make([]byte, len(m.Data2)*10)
		var j1 int
		for _, num1 := range m.Data2 {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintCanonicalMessage(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data1) > 0 {
		for iNdEx := len(m.Data1) - 1; iNdEx >= 0; iNdEx-- {
			f3 := math.Float64bits(float64(m.Data1[iNdEx]))
			i -= 8
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f3))
		}
		i = encodeVarintCanonicalMessage(dAtA, i, uint64(len(m.Data1)*8))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TimeStamp) > 0 {
		dAtA5 := make([]byte, len(m.TimeStamp)*10)
		var j4 int
		for _, num1 := range m.TimeStamp {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		i -= j4
		copy(dAtA[i:], dAtA5[:j4])
		i = encodeVarintCanonicalMessage(dAtA, i, uint64(j4))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CanonicalStream) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CanonicalStream) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CanonicalStream) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DataStream) > 0 {
		for iNdEx := len(m.DataStream) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataStream[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCanonicalMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCanonicalMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovCanonicalMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Canonical) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeStamp != 0 {
		n += 1 + sovCanonicalMessage(uint64(m.TimeStamp))
	}
	if m.Data1 != 0 {
		n += 9
	}
	if m.Data2 != 0 {
		n += 1 + sovCanonicalMessage(uint64(m.Data2))
	}
	return n
}

func (m *CanonicalRepeatedColumn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TimeStamp) > 0 {
		l = 0
		for _, e := range m.TimeStamp {
			l += sovCanonicalMessage(uint64(e))
		}
		n += 1 + sovCanonicalMessage(uint64(l)) + l
	}
	if len(m.Data1) > 0 {
		n += 1 + sovCanonicalMessage(uint64(len(m.Data1)*8)) + len(m.Data1)*8
	}
	if len(m.Data2) > 0 {
		l = 0
		for _, e := range m.Data2 {
			l += sovCanonicalMessage(uint64(e))
		}
		n += 1 + sovCanonicalMessage(uint64(l)) + l
	}
	return n
}

func (m *CanonicalStream) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DataStream) > 0 {
		for _, e := range m.DataStream {
			l = e.Size()
			n += 1 + l + sovCanonicalMessage(uint64(l))
		}
	}
	return n
}

func sovCanonicalMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCanonicalMessage(x uint64) (n int) {
	return sovCanonicalMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Canonical) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Canonical{`,
		`TimeStamp:` + fmt.Sprintf("%v", this.TimeStamp) + `,`,
		`Data1:` + fmt.Sprintf("%v", this.Data1) + `,`,
		`Data2:` + fmt.Sprintf("%v", this.Data2) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CanonicalRepeatedColumn) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CanonicalRepeatedColumn{`,
		`TimeStamp:` + fmt.Sprintf("%v", this.TimeStamp) + `,`,
		`Data1:` + fmt.Sprintf("%v", this.Data1) + `,`,
		`Data2:` + fmt.Sprintf("%v", this.Data2) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CanonicalStream) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForDataStream := "[]*Canonical{"
	for _, f := range this.DataStream {
		repeatedStringForDataStream += strings.Replace(f.String(), "Canonical", "Canonical", 1) + ","
	}
	repeatedStringForDataStream += "}"
	s := strings.Join([]string{`&CanonicalStream{`,
		`DataStream:` + repeatedStringForDataStream + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCanonicalMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Canonical) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonicalMessage
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
			return fmt.Errorf("proto: Canonical: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Canonical: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonicalMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data1", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Data1 = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data2", wireType)
			}
			m.Data2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonicalMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Data2 |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCanonicalMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonicalMessage
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
func (m *CanonicalRepeatedColumn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonicalMessage
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
			return fmt.Errorf("proto: CanonicalRepeatedColumn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalRepeatedColumn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCanonicalMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TimeStamp = append(m.TimeStamp, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCanonicalMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCanonicalMessage
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCanonicalMessage
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.TimeStamp) == 0 {
					m.TimeStamp = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCanonicalMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TimeStamp = append(m.TimeStamp, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
		case 2:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Data1 = append(m.Data1, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCanonicalMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCanonicalMessage
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCanonicalMessage
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.Data1) == 0 {
					m.Data1 = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Data1 = append(m.Data1, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Data1", wireType)
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCanonicalMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Data2 = append(m.Data2, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCanonicalMessage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCanonicalMessage
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCanonicalMessage
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Data2) == 0 {
					m.Data2 = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCanonicalMessage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Data2 = append(m.Data2, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Data2", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCanonicalMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonicalMessage
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
func (m *CanonicalStream) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCanonicalMessage
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
			return fmt.Errorf("proto: CanonicalStream: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CanonicalStream: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataStream", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCanonicalMessage
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
				return ErrInvalidLengthCanonicalMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCanonicalMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataStream = append(m.DataStream, &Canonical{})
			if err := m.DataStream[len(m.DataStream)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCanonicalMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCanonicalMessage
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
func skipCanonicalMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCanonicalMessage
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
					return 0, ErrIntOverflowCanonicalMessage
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
					return 0, ErrIntOverflowCanonicalMessage
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
				return 0, ErrInvalidLengthCanonicalMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCanonicalMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCanonicalMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCanonicalMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCanonicalMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCanonicalMessage = fmt.Errorf("proto: unexpected end of group")
)
