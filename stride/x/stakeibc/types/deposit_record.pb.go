// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/deposit_record.proto

package types

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

type DepositRecord struct {
	Id       uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount   int32     `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom    string    `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	HostZone *HostZone `protobuf:"bytes,4,opt,name=hostZone,proto3" json:"hostZone,omitempty"`
	Sender   string    `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
	Purpose  int32     `protobuf:"varint,6,opt,name=purpose,proto3" json:"purpose,omitempty"`
}

func (m *DepositRecord) Reset()         { *m = DepositRecord{} }
func (m *DepositRecord) String() string { return proto.CompactTextString(m) }
func (*DepositRecord) ProtoMessage()    {}
func (*DepositRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc20567faf76ca44, []int{0}
}
func (m *DepositRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositRecord.Merge(m, src)
}
func (m *DepositRecord) XXX_Size() int {
	return m.Size()
}
func (m *DepositRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DepositRecord proto.InternalMessageInfo

func (m *DepositRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DepositRecord) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DepositRecord) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DepositRecord) GetHostZone() *HostZone {
	if m != nil {
		return m.HostZone
	}
	return nil
}

func (m *DepositRecord) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *DepositRecord) GetPurpose() int32 {
	if m != nil {
		return m.Purpose
	}
	return 0
}

func init() {
	proto.RegisterType((*DepositRecord)(nil), "Stridelabs.stride.stakeibc.DepositRecord")
}

func init() { proto.RegisterFile("stakeibc/deposit_record.proto", fileDescriptor_cc20567faf76ca44) }

var fileDescriptor_cc20567faf76ca44 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4b, 0xc4, 0x30,
	0x18, 0xc6, 0x9b, 0x7a, 0xad, 0x1a, 0xd1, 0x21, 0x88, 0x84, 0x03, 0x43, 0x11, 0x87, 0x2e, 0xa6,
	0xa0, 0x5f, 0x40, 0xc4, 0xe1, 0xe6, 0xba, 0xdd, 0x72, 0xb4, 0xcd, 0x8b, 0x17, 0xb4, 0x7d, 0x4b,
	0x92, 0x82, 0xfa, 0x29, 0xfc, 0x4a, 0x6e, 0x8e, 0x37, 0x3a, 0x4a, 0xfb, 0x45, 0xa4, 0x7f, 0xae,
	0x9b, 0x5b, 0x1e, 0xc8, 0xef, 0x79, 0x7f, 0x3c, 0xf4, 0xd2, 0xba, 0xec, 0x05, 0x74, 0x5e, 0x24,
	0x0a, 0x6a, 0xb4, 0xda, 0x6d, 0x0c, 0x14, 0x68, 0x94, 0xac, 0x0d, 0x3a, 0x64, 0xcb, 0x27, 0x67,
	0xb4, 0x82, 0xd7, 0x2c, 0xb7, 0xd2, 0x0e, 0x4f, 0xb9, 0x07, 0x96, 0x7c, 0x46, 0xb7, 0x68, 0xdd,
	0xe6, 0x03, 0x2b, 0x18, 0xa9, 0xab, 0x2f, 0x42, 0x4f, 0x1f, 0xc7, 0xba, 0x74, 0x68, 0x63, 0x67,
	0xd4, 0xd7, 0x8a, 0x93, 0x88, 0xc4, 0x8b, 0xd4, 0xd7, 0x8a, 0x5d, 0xd0, 0x30, 0x2b, 0xb1, 0xa9,
	0x1c, 0xf7, 0x23, 0x12, 0x07, 0xe9, 0x94, 0xd8, 0x39, 0x0d, 0x14, 0x54, 0x58, 0xf2, 0x83, 0x88,
	0xc4, 0xc7, 0xe9, 0x18, 0xd8, 0x3d, 0x3d, 0xea, 0x4f, 0xac, 0xb1, 0x02, 0xbe, 0x88, 0x48, 0x7c,
	0x72, 0x7b, 0x2d, 0xff, 0x17, 0x93, 0xab, 0xe9, 0x6f, 0x3a, 0x53, 0xfd, 0x3d, 0x0b, 0x95, 0x02,
	0xc3, 0x83, 0xa1, 0x78, 0x4a, 0x8c, 0xd3, 0xc3, 0xba, 0x31, 0x35, 0x5a, 0xe0, 0xe1, 0x20, 0xb2,
	0x8f, 0x0f, 0xab, 0xef, 0x56, 0x90, 0x5d, 0x2b, 0xc8, 0x6f, 0x2b, 0xc8, 0x67, 0x27, 0xbc, 0x5d,
	0x27, 0xbc, 0x9f, 0x4e, 0x78, 0x6b, 0xf9, 0xac, 0xdd, 0xb6, 0xc9, 0x65, 0x81, 0x65, 0x32, 0x5a,
	0xdc, 0xf4, 0x1a, 0xc9, 0xa8, 0x91, 0xbc, 0x25, 0xf3, 0x2e, 0xee, 0xbd, 0x06, 0x9b, 0x87, 0xc3,
	0x28, 0x77, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0xb5, 0x36, 0xcd, 0x6b, 0x01, 0x00, 0x00,
}

func (m *DepositRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Purpose != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.Purpose))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintDepositRecord(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x2a
	}
	if m.HostZone != nil {
		{
			size, err := m.HostZone.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDepositRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDepositRecord(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDepositRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovDepositRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovDepositRecord(uint64(m.Id))
	}
	if m.Amount != 0 {
		n += 1 + sovDepositRecord(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDepositRecord(uint64(l))
	}
	if m.HostZone != nil {
		l = m.HostZone.Size()
		n += 1 + l + sovDepositRecord(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovDepositRecord(uint64(l))
	}
	if m.Purpose != 0 {
		n += 1 + sovDepositRecord(uint64(m.Purpose))
	}
	return n
}

func sovDepositRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDepositRecord(x uint64) (n int) {
	return sovDepositRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDepositRecord
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
			return fmt.Errorf("proto: DepositRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HostZone == nil {
				m.HostZone = &HostZone{}
			}
			if err := m.HostZone.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Purpose", wireType)
			}
			m.Purpose = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Purpose |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDepositRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDepositRecord
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
func skipDepositRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
				return 0, ErrInvalidLengthDepositRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDepositRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDepositRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDepositRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDepositRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDepositRecord = fmt.Errorf("proto: unexpected end of group")
)
