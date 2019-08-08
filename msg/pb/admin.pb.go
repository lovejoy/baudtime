// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: admin.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type AdminCmdRequest struct {
	// Types that are valid to be assigned to Command:
	//	*AdminCmdRequest_Info
	//	*AdminCmdRequest_JoinCluster
	Command isAdminCmdRequest_Command `protobuf_oneof:"command"`
}

func (m *AdminCmdRequest) Reset()         { *m = AdminCmdRequest{} }
func (m *AdminCmdRequest) String() string { return proto.CompactTextString(m) }
func (*AdminCmdRequest) ProtoMessage()    {}
func (*AdminCmdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b9c1af4f7059a473, []int{0}
}
func (m *AdminCmdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdminCmdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdminCmdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AdminCmdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminCmdRequest.Merge(dst, src)
}
func (m *AdminCmdRequest) XXX_Size() int {
	return m.Size()
}
func (m *AdminCmdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminCmdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdminCmdRequest proto.InternalMessageInfo

type isAdminCmdRequest_Command interface {
	isAdminCmdRequest_Command()
	MarshalTo([]byte) (int, error)
	Size() int
}

type AdminCmdRequest_Info struct {
	Info *Info `protobuf:"bytes,1,opt,name=info,oneof"`
}
type AdminCmdRequest_JoinCluster struct {
	JoinCluster *JoinCluster `protobuf:"bytes,2,opt,name=joinCluster,oneof"`
}

func (*AdminCmdRequest_Info) isAdminCmdRequest_Command()        {}
func (*AdminCmdRequest_JoinCluster) isAdminCmdRequest_Command() {}

func (m *AdminCmdRequest) GetCommand() isAdminCmdRequest_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *AdminCmdRequest) GetInfo() *Info {
	if x, ok := m.GetCommand().(*AdminCmdRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *AdminCmdRequest) GetJoinCluster() *JoinCluster {
	if x, ok := m.GetCommand().(*AdminCmdRequest_JoinCluster); ok {
		return x.JoinCluster
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdminCmdRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdminCmdRequest_OneofMarshaler, _AdminCmdRequest_OneofUnmarshaler, _AdminCmdRequest_OneofSizer, []interface{}{
		(*AdminCmdRequest_Info)(nil),
		(*AdminCmdRequest_JoinCluster)(nil),
	}
}

func _AdminCmdRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdminCmdRequest)
	// command
	switch x := m.Command.(type) {
	case *AdminCmdRequest_Info:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Info); err != nil {
			return err
		}
	case *AdminCmdRequest_JoinCluster:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JoinCluster); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AdminCmdRequest.Command has unexpected type %T", x)
	}
	return nil
}

func _AdminCmdRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdminCmdRequest)
	switch tag {
	case 1: // command.info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Info)
		err := b.DecodeMessage(msg)
		m.Command = &AdminCmdRequest_Info{msg}
		return true, err
	case 2: // command.joinCluster
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JoinCluster)
		err := b.DecodeMessage(msg)
		m.Command = &AdminCmdRequest_JoinCluster{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AdminCmdRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdminCmdRequest)
	// command
	switch x := m.Command.(type) {
	case *AdminCmdRequest_Info:
		s := proto.Size(x.Info)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdminCmdRequest_JoinCluster:
		s := proto.Size(x.JoinCluster)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Info struct {
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b9c1af4f7059a473, []int{1}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Info.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return m.Size()
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

type JoinCluster struct {
}

func (m *JoinCluster) Reset()         { *m = JoinCluster{} }
func (m *JoinCluster) String() string { return proto.CompactTextString(m) }
func (*JoinCluster) ProtoMessage()    {}
func (*JoinCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_b9c1af4f7059a473, []int{2}
}
func (m *JoinCluster) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinCluster.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *JoinCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinCluster.Merge(dst, src)
}
func (m *JoinCluster) XXX_Size() int {
	return m.Size()
}
func (m *JoinCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinCluster.DiscardUnknown(m)
}

var xxx_messageInfo_JoinCluster proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdminCmdRequest)(nil), "pb.AdminCmdRequest")
	proto.RegisterType((*Info)(nil), "pb.Info")
	proto.RegisterType((*JoinCluster)(nil), "pb.JoinCluster")
}
func (m *AdminCmdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdminCmdRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Command != nil {
		nn1, err := m.Command.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *AdminCmdRequest_Info) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Info != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.Info.Size()))
		n2, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *AdminCmdRequest_JoinCluster) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.JoinCluster != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAdmin(dAtA, i, uint64(m.JoinCluster.Size()))
		n3, err := m.JoinCluster.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *JoinCluster) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinCluster) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AdminCmdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Command != nil {
		n += m.Command.Size()
	}
	return n
}

func (m *AdminCmdRequest_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}
func (m *AdminCmdRequest_JoinCluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.JoinCluster != nil {
		l = m.JoinCluster.Size()
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}
func (m *Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *JoinCluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovAdmin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AdminCmdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: AdminCmdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdminCmdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Info{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Command = &AdminCmdRequest_Info{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoinCluster", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &JoinCluster{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Command = &AdminCmdRequest_JoinCluster{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
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
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
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
func (m *JoinCluster) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: JoinCluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinCluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
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
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
				return 0, ErrInvalidLengthAdmin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAdmin
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
				next, err := skipAdmin(dAtA[start:])
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
	ErrInvalidLengthAdmin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("admin.proto", fileDescriptor_admin_b9c1af4f7059a473) }

var fileDescriptor_admin_b9c1af4f7059a473 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0xc9, 0xcd,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x2a, 0xe4, 0xe2, 0x77,
	0x04, 0x09, 0x39, 0xe7, 0xa6, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x71, 0xb1,
	0x64, 0xe6, 0xa5, 0xe5, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71, 0xe8, 0x15, 0x24, 0xe9,
	0x79, 0xe6, 0xa5, 0xe5, 0x7b, 0x30, 0x04, 0x81, 0xc5, 0x85, 0x8c, 0xb9, 0xb8, 0xb3, 0xf2, 0x33,
	0xf3, 0x9c, 0x73, 0x4a, 0x8b, 0x4b, 0x52, 0x8b, 0x24, 0x98, 0xc0, 0xca, 0xf8, 0x41, 0xca, 0xbc,
	0x10, 0xc2, 0x1e, 0x0c, 0x41, 0xc8, 0xaa, 0x9c, 0x38, 0xb9, 0xd8, 0x93, 0xf3, 0x73, 0x73, 0x13,
	0xf3, 0x52, 0x94, 0xd8, 0xb8, 0x58, 0x40, 0xe6, 0x29, 0xf1, 0x72, 0x71, 0x23, 0x69, 0x70, 0x92,
	0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xa6, 0x82, 0xa4, 0x24, 0x36,
	0xb0, 0x93, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0x22, 0xb8, 0x8f, 0xc1, 0x00, 0x00,
	0x00,
}
