// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/fee/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type EventFeeBurn struct {
	Fees github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=fees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fees"`
}

func (m *EventFeeBurn) Reset()         { *m = EventFeeBurn{} }
func (m *EventFeeBurn) String() string { return proto.CompactTextString(m) }
func (*EventFeeBurn) ProtoMessage()    {}
func (*EventFeeBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa464a73c5ec94d, []int{0}
}
func (m *EventFeeBurn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFeeBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFeeBurn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFeeBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFeeBurn.Merge(m, src)
}
func (m *EventFeeBurn) XXX_Size() int {
	return m.Size()
}
func (m *EventFeeBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFeeBurn.DiscardUnknown(m)
}

var xxx_messageInfo_EventFeeBurn proto.InternalMessageInfo

func (m *EventFeeBurn) GetFees() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fees
	}
	return nil
}

func init() {
	proto.RegisterType((*EventFeeBurn)(nil), "sunrise.fee.EventFeeBurn")
}

func init() { proto.RegisterFile("sunrise/fee/events.proto", fileDescriptor_9fa464a73c5ec94d) }

var fileDescriptor_9fa464a73c5ec94d = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2e, 0xcd, 0x2b,
	0xca, 0x2c, 0x4e, 0xd5, 0x4f, 0x4b, 0x4d, 0xd5, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xca, 0xe8, 0xa5, 0xa5, 0xa6, 0x4a, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0xc5, 0xf5, 0x41, 0x2c, 0x88, 0x12, 0x29, 0xc9, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc,
	0xe2, 0x78, 0x88, 0x04, 0x84, 0x03, 0x95, 0x92, 0x83, 0xf0, 0xf4, 0x93, 0x12, 0x8b, 0x53, 0xf5,
	0xcb, 0x0c, 0x93, 0x52, 0x4b, 0x12, 0x0d, 0xf5, 0x93, 0xf3, 0x33, 0xf3, 0x20, 0xf2, 0x4a, 0xf9,
	0x5c, 0x3c, 0xae, 0x20, 0xdb, 0xdc, 0x52, 0x53, 0x9d, 0x4a, 0x8b, 0xf2, 0x84, 0xe2, 0xb9, 0x58,
	0xd2, 0x52, 0x53, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x24, 0xf5, 0xa0, 0x86, 0x81,
	0xb4, 0xeb, 0x41, 0xb5, 0xeb, 0x39, 0xe7, 0x67, 0xe6, 0x39, 0x19, 0x9c, 0xb8, 0x27, 0xcf, 0xb0,
	0xea, 0xbe, 0xbc, 0x46, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0xd4, 0x66,
	0x28, 0xa5, 0x5b, 0x9c, 0x92, 0xad, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0x0c, 0xd6, 0x50, 0x1c, 0x04,
	0x36, 0xd8, 0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0x91,
	0x4c, 0x82, 0xfa, 0x39, 0x27, 0xb1, 0x32, 0xb5, 0x08, 0xc6, 0xd1, 0xaf, 0x00, 0x07, 0x0e, 0xd8,
	0xc8, 0x24, 0x36, 0xb0, 0xf3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x93, 0x00, 0x14,
	0x38, 0x01, 0x00, 0x00,
}

func (m *EventFeeBurn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFeeBurn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFeeBurn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventFeeBurn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventFeeBurn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventFeeBurn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFeeBurn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, types.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
