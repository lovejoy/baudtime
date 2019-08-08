// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StatusCode int32

const (
	StatusCode_Succeed StatusCode = 0
	StatusCode_Failed  StatusCode = 1
)

var StatusCode_name = map[int32]string{
	0: "Succeed",
	1: "Failed",
}
var StatusCode_value = map[string]int32{
	"Succeed": 0,
	"Failed":  1,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pb_684b3f305cfcd3a3, []int{0}
}

type Label struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_684b3f305cfcd3a3, []int{0}
}
func (m *Label) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Label.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(dst, src)
}
func (m *Label) XXX_Size() int {
	return m.Size()
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Point struct {
	T int64   `protobuf:"zigzag64,1,opt,name=T,proto3" json:"T,omitempty"`
	V float64 `protobuf:"fixed64,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_684b3f305cfcd3a3, []int{1}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Point.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return m.Size()
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetT() int64 {
	if m != nil {
		return m.T
	}
	return 0
}

func (m *Point) GetV() float64 {
	if m != nil {
		return m.V
	}
	return 0
}

type Series struct {
	Labels []Label `protobuf:"bytes,1,rep,name=labels" json:"labels"`
	Points []Point `protobuf:"bytes,2,rep,name=points" json:"points"`
}

func (m *Series) Reset()         { *m = Series{} }
func (m *Series) String() string { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()    {}
func (*Series) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_684b3f305cfcd3a3, []int{2}
}
func (m *Series) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Series) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Series.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Series) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Series.Merge(dst, src)
}
func (m *Series) XXX_Size() int {
	return m.Size()
}
func (m *Series) XXX_DiscardUnknown() {
	xxx_messageInfo_Series.DiscardUnknown(m)
}

var xxx_messageInfo_Series proto.InternalMessageInfo

func (m *Series) GetLabels() []Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Series) GetPoints() []Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type LabelValuesResponse struct {
	Values   []string   `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	Status   StatusCode `protobuf:"varint,2,opt,name=status,proto3,enum=pb.StatusCode" json:"status,omitempty"`
	ErrorMsg string     `protobuf:"bytes,3,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (m *LabelValuesResponse) Reset()         { *m = LabelValuesResponse{} }
func (m *LabelValuesResponse) String() string { return proto.CompactTextString(m) }
func (*LabelValuesResponse) ProtoMessage()    {}
func (*LabelValuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_684b3f305cfcd3a3, []int{3}
}
func (m *LabelValuesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelValuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelValuesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LabelValuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelValuesResponse.Merge(dst, src)
}
func (m *LabelValuesResponse) XXX_Size() int {
	return m.Size()
}
func (m *LabelValuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelValuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LabelValuesResponse proto.InternalMessageInfo

func (m *LabelValuesResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *LabelValuesResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_Succeed
}

func (m *LabelValuesResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

type GeneralResponse struct {
	Status  StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=pb.StatusCode" json:"status,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_684b3f305cfcd3a3, []int{4}
}
func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(dst, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return m.Size()
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

func (m *GeneralResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_Succeed
}

func (m *GeneralResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Label)(nil), "pb.Label")
	proto.RegisterType((*Point)(nil), "pb.Point")
	proto.RegisterType((*Series)(nil), "pb.Series")
	proto.RegisterType((*LabelValuesResponse)(nil), "pb.LabelValuesResponse")
	proto.RegisterType((*GeneralResponse)(nil), "pb.GeneralResponse")
	proto.RegisterEnum("pb.StatusCode", StatusCode_name, StatusCode_value)
}
func (m *Label) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Label) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *Point) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Point) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.T != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPb(dAtA, i, uint64((uint64(m.T)<<1)^uint64((m.T>>63))))
	}
	if m.V != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.V))))
		i += 8
	}
	return i, nil
}

func (m *Series) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Series) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, msg := range m.Labels {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Points) > 0 {
		for _, msg := range m.Points {
			dAtA[i] = 0x12
			i++
			i = encodeVarintPb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *LabelValuesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelValuesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPb(dAtA, i, uint64(m.Status))
	}
	if len(m.ErrorMsg) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.ErrorMsg)))
		i += copy(dAtA[i:], m.ErrorMsg)
	}
	return i, nil
}

func (m *GeneralResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GeneralResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPb(dAtA, i, uint64(m.Status))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintPb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Label) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	return n
}

func (m *Point) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.T != 0 {
		n += 1 + sozPb(uint64(m.T))
	}
	if m.V != 0 {
		n += 9
	}
	return n
}

func (m *Series) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	if len(m.Points) > 0 {
		for _, e := range m.Points {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	return n
}

func (m *LabelValuesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovPb(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovPb(uint64(m.Status))
	}
	l = len(m.ErrorMsg)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	return n
}

func (m *GeneralResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovPb(uint64(m.Status))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	return n
}

func sovPb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPb(x uint64) (n int) {
	return sovPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Label) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Label: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Label: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *Point) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Point: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Point: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.T = int64(v)
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field V", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.V = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *Series) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Series: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Series: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, Label{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Points = append(m.Points, Point{})
			if err := m.Points[len(m.Points)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *LabelValuesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelValuesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelValuesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (StatusCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func (m *GeneralResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GeneralResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GeneralResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (StatusCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
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
func skipPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPb
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
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPb(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb.proto", fileDescriptor_pb_684b3f305cfcd3a3) }

var fileDescriptor_pb_684b3f305cfcd3a3 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x3b, 0x05, 0x0a, 0x5c, 0xbe, 0xf0, 0x91, 0xd1, 0x98, 0x86, 0x98, 0x4a, 0x6a, 0x54,
	0x62, 0x62, 0x89, 0xf8, 0x06, 0x98, 0xe8, 0x46, 0x13, 0xd3, 0x12, 0x16, 0xec, 0x3a, 0x70, 0xad,
	0x4d, 0x4a, 0xa7, 0x76, 0x5a, 0x9f, 0xc3, 0xc7, 0x62, 0xc9, 0xd2, 0x95, 0x31, 0xf0, 0x22, 0xa6,
	0xb7, 0xfc, 0xd9, 0xb9, 0x9b, 0xdf, 0x9c, 0x73, 0xcf, 0xb9, 0xed, 0x40, 0x23, 0x11, 0x4e, 0x92,
	0xca, 0x4c, 0x72, 0x3d, 0x11, 0xdd, 0x9b, 0x20, 0xcc, 0xde, 0x72, 0xe1, 0xcc, 0xe4, 0x62, 0x10,
	0xc8, 0x40, 0x0e, 0x48, 0x12, 0xf9, 0x2b, 0x11, 0x01, 0x9d, 0xca, 0x11, 0xfb, 0x16, 0x6a, 0x4f,
	0xbe, 0xc0, 0x88, 0x73, 0xa8, 0xc6, 0xfe, 0x02, 0x4d, 0xd6, 0x63, 0xfd, 0xa6, 0x4b, 0x67, 0x7e,
	0x0c, 0xb5, 0x0f, 0x3f, 0xca, 0xd1, 0xd4, 0xe9, 0xb2, 0x04, 0xfb, 0x1c, 0x6a, 0x2f, 0x32, 0x8c,
	0x33, 0xfe, 0x0f, 0xd8, 0x98, 0xfc, 0xdc, 0x65, 0xe3, 0x82, 0x26, 0x64, 0x64, 0x2e, 0x9b, 0xd8,
	0x53, 0x30, 0x3c, 0x4c, 0x43, 0x54, 0xfc, 0x0a, 0x8c, 0xa8, 0x68, 0x50, 0x26, 0xeb, 0x55, 0xfa,
	0xad, 0x61, 0xd3, 0x49, 0x84, 0x43, 0x9d, 0xa3, 0xea, 0xf2, 0xfb, 0x4c, 0x73, 0xb7, 0x72, 0x61,
	0x4c, 0x8a, 0x5c, 0x65, 0xea, 0x07, 0x23, 0x35, 0xed, 0x8c, 0xa5, 0x6c, 0xbf, 0xc3, 0x11, 0xcd,
	0x4f, 0x8a, 0x75, 0x94, 0x8b, 0x2a, 0x91, 0xb1, 0x42, 0x7e, 0x02, 0x06, 0x2d, 0x58, 0x16, 0x35,
	0xdd, 0x2d, 0xf1, 0x4b, 0x30, 0x54, 0xe6, 0x67, 0xb9, 0xa2, 0xed, 0xda, 0xc3, 0x76, 0x91, 0xeb,
	0xd1, 0xcd, 0xbd, 0x9c, 0xa3, 0xbb, 0x55, 0x79, 0x17, 0x1a, 0x98, 0xa6, 0x32, 0x7d, 0x56, 0x81,
	0x59, 0xa1, 0x0f, 0xde, 0xb3, 0xed, 0xc1, 0xff, 0x47, 0x8c, 0x31, 0xf5, 0xa3, 0x7d, 0xdd, 0x21,
	0x96, 0xfd, 0x19, 0x6b, 0x42, 0x7d, 0x81, 0x4a, 0xf9, 0xc1, 0xee, 0x37, 0xee, 0xf0, 0xfa, 0x02,
	0xe0, 0xe0, 0xe7, 0x2d, 0xa8, 0x7b, 0xf9, 0x6c, 0x86, 0x38, 0xef, 0x68, 0x1c, 0xc0, 0x78, 0xf0,
	0xc3, 0x08, 0xe7, 0x1d, 0x36, 0x3a, 0x5d, 0xae, 0x2d, 0xb6, 0x5a, 0x5b, 0xec, 0x67, 0x6d, 0xb1,
	0xcf, 0x8d, 0xa5, 0xad, 0x36, 0x96, 0xf6, 0xb5, 0xb1, 0xb4, 0xa9, 0x9e, 0x08, 0x61, 0xd0, 0x3b,
	0xde, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x23, 0xc7, 0xd8, 0xa9, 0x06, 0x02, 0x00, 0x00,
}
