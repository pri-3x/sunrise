// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquidityincentive/events.proto

package types

import (
	fmt "fmt"
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

type EventSetVote struct {
	Address string       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Weights []PoolWeight `protobuf:"bytes,2,rep,name=weights,proto3" json:"weights"`
}

func (m *EventSetVote) Reset()         { *m = EventSetVote{} }
func (m *EventSetVote) String() string { return proto.CompactTextString(m) }
func (*EventSetVote) ProtoMessage()    {}
func (*EventSetVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7ff3c117acf07e, []int{0}
}
func (m *EventSetVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetVote.Merge(m, src)
}
func (m *EventSetVote) XXX_Size() int {
	return m.Size()
}
func (m *EventSetVote) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetVote.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetVote proto.InternalMessageInfo

func (m *EventSetVote) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventSetVote) GetWeights() []PoolWeight {
	if m != nil {
		return m.Weights
	}
	return nil
}

func init() {
	proto.RegisterType((*EventSetVote)(nil), "sunrise.liquidityincentive.EventSetVote")
}

func init() {
	proto.RegisterFile("sunrise/liquidityincentive/events.proto", fileDescriptor_6c7ff3c117acf07e)
}

var fileDescriptor_6c7ff3c117acf07e = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x2e, 0xcd, 0x2b,
	0xca, 0x2c, 0x4e, 0xd5, 0xcf, 0xc9, 0x2c, 0x2c, 0xcd, 0x4c, 0xc9, 0x2c, 0xa9, 0xcc, 0xcc, 0x4b,
	0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x82, 0x2a, 0xd4, 0xc3, 0x54, 0x28, 0x25, 0x92, 0x9e, 0x9f, 0x9e,
	0x0f, 0x56, 0xa6, 0x0f, 0x62, 0x41, 0x74, 0x48, 0xa9, 0xe1, 0x31, 0x3a, 0x3d, 0xb1, 0x34, 0x3d,
	0x15, 0xa2, 0x4e, 0xa9, 0x80, 0x8b, 0xc7, 0x15, 0x64, 0x53, 0x70, 0x6a, 0x49, 0x58, 0x7e, 0x49,
	0xaa, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x8c, 0x2b, 0xe4, 0xc6, 0xc5, 0x5e, 0x9e, 0x9a, 0x99, 0x9e, 0x51, 0x52, 0x2c,
	0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa6, 0x87, 0xdb, 0x55, 0x7a, 0x01, 0xf9, 0xf9, 0x39,
	0xe1, 0x60, 0xe5, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0xc1, 0x34, 0x3b, 0x85, 0x9c, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c,
	0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x55, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd4, 0xe8, 0x9c, 0xc4, 0xca, 0xd4, 0x22, 0x18, 0x47, 0xbf, 0x02,
	0x9b, 0x6f, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xde, 0x31, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0x63, 0xc8, 0x14, 0x53, 0x01, 0x00, 0x00,
}

func (m *EventSetVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Weights) > 0 {
		for iNdEx := len(m.Weights) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Weights[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *EventSetVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Weights) > 0 {
		for _, e := range m.Weights {
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
func (m *EventSetVote) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventSetVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weights", wireType)
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
			m.Weights = append(m.Weights, PoolWeight{})
			if err := m.Weights[len(m.Weights)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
