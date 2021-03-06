// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: money.proto

package msg

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

/// 用户货币信息
type Money struct {
	/// 类型, 1=金币,2=钻石,3=积分,
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	/// 数量
	Num int64 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_2917341c8b38eaf8, []int{0}
}
func (m *Money) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Money.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(m, src)
}
func (m *Money) XXX_Size() int {
	return m.Size()
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Money) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

/// 更新货币
//@msg
type S2C_UpdateMoney struct {
	/// 发生变更的货币列表
	Monies []*Money `protobuf:"bytes,1,rep,name=Monies,proto3" json:"Monies,omitempty"`
}

func (m *S2C_UpdateMoney) Reset()         { *m = S2C_UpdateMoney{} }
func (m *S2C_UpdateMoney) String() string { return proto.CompactTextString(m) }
func (*S2C_UpdateMoney) ProtoMessage()    {}
func (*S2C_UpdateMoney) Descriptor() ([]byte, []int) {
	return fileDescriptor_2917341c8b38eaf8, []int{1}
}
func (m *S2C_UpdateMoney) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C_UpdateMoney) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C_UpdateMoney.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S2C_UpdateMoney) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_UpdateMoney.Merge(m, src)
}
func (m *S2C_UpdateMoney) XXX_Size() int {
	return m.Size()
}
func (m *S2C_UpdateMoney) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_UpdateMoney.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_UpdateMoney proto.InternalMessageInfo

func (m *S2C_UpdateMoney) GetMonies() []*Money {
	if m != nil {
		return m.Monies
	}
	return nil
}

func init() {
	proto.RegisterType((*Money)(nil), "msg.Money")
	proto.RegisterType((*S2C_UpdateMoney)(nil), "msg.S2C_UpdateMoney")
}

func init() { proto.RegisterFile("money.proto", fileDescriptor_2917341c8b38eaf8) }

var fileDescriptor_2917341c8b38eaf8 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0xcf, 0x4b,
	0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xd2, 0xe5, 0x62,
	0xf5, 0x05, 0x89, 0x09, 0x09, 0x71, 0xb1, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x06, 0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0x7e, 0xa5, 0xb9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0xcc, 0x41, 0x20, 0xa6, 0x92, 0x29, 0x17, 0x7f, 0xb0, 0x91, 0x73, 0x7c, 0x68, 0x41, 0x4a, 0x62,
	0x49, 0x2a, 0x44, 0xa3, 0x12, 0x17, 0x9b, 0x6f, 0x7e, 0x5e, 0x66, 0x6a, 0xb1, 0x04, 0xa3, 0x02,
	0xb3, 0x06, 0xb7, 0x11, 0x97, 0x5e, 0x6e, 0x71, 0xba, 0x1e, 0x58, 0x2e, 0x08, 0x2a, 0xe3, 0x24,
	0x71, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0xb7, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x4a, 0x76, 0xfe, 0x9a, 0x00, 0x00, 0x00,
}

func (m *Money) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Money) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Money) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Num != 0 {
		i = encodeVarintMoney(dAtA, i, uint64(m.Num))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintMoney(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *S2C_UpdateMoney) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_UpdateMoney) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *S2C_UpdateMoney) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Monies) > 0 {
		for iNdEx := len(m.Monies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Monies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMoney(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMoney(dAtA []byte, offset int, v uint64) int {
	offset -= sovMoney(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Money) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMoney(uint64(m.Type))
	}
	if m.Num != 0 {
		n += 1 + sovMoney(uint64(m.Num))
	}
	return n
}

func (m *S2C_UpdateMoney) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Monies) > 0 {
		for _, e := range m.Monies {
			l = e.Size()
			n += 1 + l + sovMoney(uint64(l))
		}
	}
	return n
}

func sovMoney(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMoney(x uint64) (n int) {
	return sovMoney(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Money) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMoney
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
			return fmt.Errorf("proto: Money: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Money: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMoney
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMoney
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMoney(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMoney
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMoney
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
func (m *S2C_UpdateMoney) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMoney
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
			return fmt.Errorf("proto: S2C_UpdateMoney: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_UpdateMoney: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Monies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMoney
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
				return ErrInvalidLengthMoney
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMoney
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Monies = append(m.Monies, &Money{})
			if err := m.Monies[len(m.Monies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMoney(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMoney
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMoney
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
func skipMoney(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMoney
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
					return 0, ErrIntOverflowMoney
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
					return 0, ErrIntOverflowMoney
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
				return 0, ErrInvalidLengthMoney
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMoney
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMoney
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMoney        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMoney          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMoney = fmt.Errorf("proto: unexpected end of group")
)
