// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/host_zone.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// next id: 13
type HostZone struct {
	ChainId               string       `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	ConnectionId          string       `protobuf:"bytes,2,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	TransferChannelId     string       `protobuf:"bytes,12,opt,name=transferChannelId,proto3" json:"transferChannelId,omitempty"`
	Validators            []*Validator `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators,omitempty"`
	BlacklistedValidators []*Validator `protobuf:"bytes,4,rep,name=blacklistedValidators,proto3" json:"blacklistedValidators,omitempty"`
	WithdrawalAccount     *ICAAccount  `protobuf:"bytes,5,opt,name=withdrawalAccount,proto3" json:"withdrawalAccount,omitempty"`
	FeeAccount            *ICAAccount  `protobuf:"bytes,6,opt,name=feeAccount,proto3" json:"feeAccount,omitempty"`
	DelegationAccount     *ICAAccount  `protobuf:"bytes,7,opt,name=delegationAccount,proto3" json:"delegationAccount,omitempty"`
	RedemptionAccount     *ICAAccount  `protobuf:"bytes,13,opt,name=redemptionAccount,proto3" json:"redemptionAccount,omitempty"`
	// ibc denom on stride
	IBCDenom string `protobuf:"bytes,8,opt,name=IBCDenom,proto3" json:"IBCDenom,omitempty"`
	// native denom on host zone
	HostDenom string `protobuf:"bytes,9,opt,name=HostDenom,proto3" json:"HostDenom,omitempty"`
	// TODO(TEST-68): Should we make this an array and store the last n redemption rates
	// then calculate a TWARR?
	LastRedemptionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=LastRedemptionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"LastRedemptionRate"`
	RedemptionRate     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=RedemptionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"RedemptionRate"`
	// stores how many days we should wait before issuing unbondings
	UnbondingFrequency uint64 `protobuf:"varint,14,opt,name=unbondingFrequency,proto3" json:"unbondingFrequency,omitempty"`
}

func (m *HostZone) Reset()         { *m = HostZone{} }
func (m *HostZone) String() string { return proto.CompactTextString(m) }
func (*HostZone) ProtoMessage()    {}
func (*HostZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1d300c62c2b2d54, []int{0}
}
func (m *HostZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostZone.Merge(m, src)
}
func (m *HostZone) XXX_Size() int {
	return m.Size()
}
func (m *HostZone) XXX_DiscardUnknown() {
	xxx_messageInfo_HostZone.DiscardUnknown(m)
}

var xxx_messageInfo_HostZone proto.InternalMessageInfo

func (m *HostZone) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *HostZone) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *HostZone) GetTransferChannelId() string {
	if m != nil {
		return m.TransferChannelId
	}
	return ""
}

func (m *HostZone) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *HostZone) GetBlacklistedValidators() []*Validator {
	if m != nil {
		return m.BlacklistedValidators
	}
	return nil
}

func (m *HostZone) GetWithdrawalAccount() *ICAAccount {
	if m != nil {
		return m.WithdrawalAccount
	}
	return nil
}

func (m *HostZone) GetFeeAccount() *ICAAccount {
	if m != nil {
		return m.FeeAccount
	}
	return nil
}

func (m *HostZone) GetDelegationAccount() *ICAAccount {
	if m != nil {
		return m.DelegationAccount
	}
	return nil
}

func (m *HostZone) GetRedemptionAccount() *ICAAccount {
	if m != nil {
		return m.RedemptionAccount
	}
	return nil
}

func (m *HostZone) GetIBCDenom() string {
	if m != nil {
		return m.IBCDenom
	}
	return ""
}

func (m *HostZone) GetHostDenom() string {
	if m != nil {
		return m.HostDenom
	}
	return ""
}

func (m *HostZone) GetUnbondingFrequency() uint64 {
	if m != nil {
		return m.UnbondingFrequency
	}
	return 0
}

func init() {
	proto.RegisterType((*HostZone)(nil), "Stridelabs.stride.stakeibc.HostZone")
}

func init() { proto.RegisterFile("stakeibc/host_zone.proto", fileDescriptor_a1d300c62c2b2d54) }

var fileDescriptor_a1d300c62c2b2d54 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x63, 0x5a, 0x9a, 0x64, 0x53, 0x2a, 0x75, 0x05, 0x92, 0xb1, 0x90, 0x1b, 0x55, 0xa2,
	0xca, 0x81, 0xd8, 0x6a, 0xb9, 0x72, 0x69, 0x12, 0xaa, 0x46, 0xea, 0xc9, 0xa0, 0x1e, 0xca, 0x21,
	0x5a, 0xef, 0x4e, 0xed, 0x55, 0x9c, 0xdd, 0xe0, 0xdd, 0xb4, 0x94, 0xa7, 0xe0, 0x19, 0x78, 0x06,
	0x1e, 0xa2, 0xc7, 0x8a, 0x13, 0xe2, 0x50, 0xa1, 0xe4, 0x45, 0x90, 0x3f, 0x13, 0x9a, 0x80, 0x14,
	0x89, 0x53, 0x66, 0xe6, 0x3f, 0xf3, 0xdb, 0xc9, 0xce, 0xac, 0x91, 0xa9, 0x34, 0x19, 0x02, 0xf7,
	0xa9, 0x1b, 0x4a, 0xa5, 0x07, 0x9f, 0xa5, 0x00, 0x67, 0x1c, 0x4b, 0x2d, 0xb1, 0xf5, 0x4e, 0xc7,
	0x9c, 0x41, 0x44, 0x7c, 0xe5, 0xa8, 0xd4, 0x74, 0x8a, 0x5c, 0x6b, 0x5e, 0x75, 0x45, 0x22, 0xce,
	0x88, 0x96, 0x71, 0x56, 0x65, 0x59, 0xa5, 0xc2, 0x29, 0x19, 0x10, 0x4a, 0xe5, 0x44, 0xe8, 0x5c,
	0x7b, 0x1a, 0xc8, 0x40, 0xa6, 0xa6, 0x9b, 0x58, 0x79, 0xf4, 0x39, 0x95, 0x6a, 0x24, 0xd5, 0x20,
	0x13, 0x32, 0x27, 0x97, 0xec, 0xcc, 0x73, 0x7d, 0xa2, 0xc0, 0xbd, 0x3a, 0xf4, 0x41, 0x93, 0x43,
	0x97, 0x4a, 0x2e, 0x32, 0x7d, 0xff, 0x6b, 0x15, 0xd5, 0x4e, 0xa5, 0xd2, 0x17, 0x52, 0x00, 0x36,
	0x51, 0x95, 0x86, 0x84, 0x8b, 0x3e, 0x33, 0x8d, 0xa6, 0xd1, 0xaa, 0x7b, 0x85, 0x8b, 0xf7, 0xd1,
	0x36, 0x95, 0x42, 0x00, 0xd5, 0x5c, 0x26, 0xf2, 0xa3, 0x54, 0xfe, 0x23, 0x86, 0x5f, 0xa1, 0x5d,
	0x1d, 0x13, 0xa1, 0x2e, 0x21, 0xee, 0x86, 0x44, 0x08, 0x88, 0xfa, 0xcc, 0xdc, 0x4e, 0x13, 0x97,
	0x05, 0xfc, 0x16, 0xa1, 0xf2, 0x8f, 0x2b, 0x73, 0xa3, 0xb9, 0xd1, 0x6a, 0x1c, 0xbd, 0x74, 0xfe,
	0x7e, 0x61, 0xce, 0x79, 0x91, 0xed, 0x2d, 0x14, 0xe2, 0x0f, 0xe8, 0x99, 0x1f, 0x11, 0x3a, 0x8c,
	0xb8, 0xd2, 0xc0, 0xce, 0xe7, 0xc4, 0xcd, 0x75, 0x88, 0xab, 0x19, 0xf8, 0x3d, 0xda, 0xbd, 0xe6,
	0x3a, 0x64, 0x31, 0xb9, 0x26, 0xd1, 0x71, 0x36, 0x08, 0xf3, 0x71, 0xd3, 0x68, 0x35, 0x8e, 0x0e,
	0xfe, 0x05, 0xee, 0x77, 0x8f, 0xf3, 0x6c, 0x6f, 0x19, 0x80, 0x4f, 0x10, 0xba, 0x04, 0x28, 0x70,
	0x5b, 0x6b, 0xe1, 0x16, 0x2a, 0x93, 0xee, 0x18, 0x44, 0x10, 0x90, 0xe4, 0xfe, 0x0b, 0x5c, 0x75,
	0xbd, 0xee, 0x96, 0x00, 0x09, 0x35, 0x06, 0x06, 0xa3, 0xf1, 0x22, 0xf5, 0xc9, 0x7a, 0xd4, 0x25,
	0x00, 0xb6, 0x50, 0xad, 0xdf, 0xe9, 0xf6, 0x40, 0xc8, 0x91, 0x59, 0x4b, 0x57, 0xa2, 0xf4, 0xf1,
	0x0b, 0x54, 0x4f, 0x36, 0x30, 0x13, 0xeb, 0xa9, 0x38, 0x0f, 0xe0, 0x08, 0xe1, 0x33, 0xa2, 0xb4,
	0x57, 0x22, 0x3d, 0xa2, 0xc1, 0x44, 0x49, 0x5a, 0xe7, 0xcd, 0xed, 0xfd, 0x5e, 0xe5, 0xe7, 0xfd,
	0xde, 0x41, 0xc0, 0x75, 0x38, 0xf1, 0x1d, 0x2a, 0x47, 0xf9, 0xf6, 0xe7, 0x3f, 0x6d, 0xc5, 0x86,
	0xae, 0xbe, 0x19, 0x83, 0x72, 0x7a, 0x40, 0xbf, 0x7f, 0x6b, 0xa3, 0xfc, 0x71, 0xf4, 0x80, 0x7a,
	0x2b, 0xb8, 0x98, 0xa1, 0x9d, 0x07, 0x27, 0x35, 0xfe, 0xc3, 0x49, 0x0f, 0x98, 0xd8, 0x41, 0x78,
	0x22, 0x7c, 0x29, 0x18, 0x17, 0xc1, 0x49, 0x0c, 0x1f, 0x27, 0x20, 0xe8, 0x8d, 0xb9, 0xd3, 0x34,
	0x5a, 0x9b, 0xde, 0x0a, 0xa5, 0x73, 0x7a, 0x3b, 0xb5, 0x8d, 0xbb, 0xa9, 0x6d, 0xfc, 0x9a, 0xda,
	0xc6, 0x97, 0x99, 0x5d, 0xb9, 0x9b, 0xd9, 0x95, 0x1f, 0x33, 0xbb, 0x72, 0xe1, 0x2c, 0xf4, 0x93,
	0x0d, 0xa7, 0x7d, 0x46, 0x7c, 0xe5, 0x66, 0xd3, 0x71, 0x3f, 0xb9, 0xe5, 0xb7, 0x24, 0xed, 0xcd,
	0xdf, 0x4a, 0x5f, 0xfd, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0xab, 0x17, 0x61, 0xb4,
	0x04, 0x00, 0x00,
}

func (m *HostZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnbondingFrequency != 0 {
		i = encodeVarintHostZone(dAtA, i, uint64(m.UnbondingFrequency))
		i--
		dAtA[i] = 0x70
	}
	if m.RedemptionAccount != nil {
		{
			size, err := m.RedemptionAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if len(m.TransferChannelId) > 0 {
		i -= len(m.TransferChannelId)
		copy(dAtA[i:], m.TransferChannelId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.TransferChannelId)))
		i--
		dAtA[i] = 0x62
	}
	{
		size := m.RedemptionRate.Size()
		i -= size
		if _, err := m.RedemptionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintHostZone(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.LastRedemptionRate.Size()
		i -= size
		if _, err := m.LastRedemptionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintHostZone(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.HostDenom) > 0 {
		i -= len(m.HostDenom)
		copy(dAtA[i:], m.HostDenom)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.HostDenom)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.IBCDenom) > 0 {
		i -= len(m.IBCDenom)
		copy(dAtA[i:], m.IBCDenom)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.IBCDenom)))
		i--
		dAtA[i] = 0x42
	}
	if m.DelegationAccount != nil {
		{
			size, err := m.DelegationAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.FeeAccount != nil {
		{
			size, err := m.FeeAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.WithdrawalAccount != nil {
		{
			size, err := m.WithdrawalAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHostZone(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BlacklistedValidators) > 0 {
		for iNdEx := len(m.BlacklistedValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlacklistedValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHostZone(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHostZone(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintHostZone(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHostZone(dAtA []byte, offset int, v uint64) int {
	offset -= sovHostZone(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovHostZone(uint64(l))
		}
	}
	if len(m.BlacklistedValidators) > 0 {
		for _, e := range m.BlacklistedValidators {
			l = e.Size()
			n += 1 + l + sovHostZone(uint64(l))
		}
	}
	if m.WithdrawalAccount != nil {
		l = m.WithdrawalAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.FeeAccount != nil {
		l = m.FeeAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.DelegationAccount != nil {
		l = m.DelegationAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.IBCDenom)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = len(m.HostDenom)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	l = m.LastRedemptionRate.Size()
	n += 1 + l + sovHostZone(uint64(l))
	l = m.RedemptionRate.Size()
	n += 1 + l + sovHostZone(uint64(l))
	l = len(m.TransferChannelId)
	if l > 0 {
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.RedemptionAccount != nil {
		l = m.RedemptionAccount.Size()
		n += 1 + l + sovHostZone(uint64(l))
	}
	if m.UnbondingFrequency != 0 {
		n += 1 + sovHostZone(uint64(m.UnbondingFrequency))
	}
	return n
}

func sovHostZone(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHostZone(x uint64) (n int) {
	return sovHostZone(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostZone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHostZone
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
			return fmt.Errorf("proto: HostZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlacklistedValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlacklistedValidators = append(m.BlacklistedValidators, &Validator{})
			if err := m.BlacklistedValidators[len(m.BlacklistedValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WithdrawalAccount == nil {
				m.WithdrawalAccount = &ICAAccount{}
			}
			if err := m.WithdrawalAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeeAccount == nil {
				m.FeeAccount = &ICAAccount{}
			}
			if err := m.FeeAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DelegationAccount == nil {
				m.DelegationAccount = &ICAAccount{}
			}
			if err := m.DelegationAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IBCDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRedemptionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastRedemptionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RedemptionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
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
				return ErrInvalidLengthHostZone
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHostZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RedemptionAccount == nil {
				m.RedemptionAccount = &ICAAccount{}
			}
			if err := m.RedemptionAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingFrequency", wireType)
			}
			m.UnbondingFrequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostZone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingFrequency |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHostZone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHostZone
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
func skipHostZone(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHostZone
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
					return 0, ErrIntOverflowHostZone
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
					return 0, ErrIntOverflowHostZone
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
				return 0, ErrInvalidLengthHostZone
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHostZone
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHostZone
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHostZone        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHostZone          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHostZone = fmt.Errorf("proto: unexpected end of group")
)
