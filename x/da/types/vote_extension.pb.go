// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/da/vote_extension.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type VoteExtension struct {
	Data   []*PublishedData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Shares []*DataShares    `protobuf:"bytes,2,rep,name=shares,proto3" json:"shares,omitempty"`
}

func (m *VoteExtension) Reset()         { *m = VoteExtension{} }
func (m *VoteExtension) String() string { return proto.CompactTextString(m) }
func (*VoteExtension) ProtoMessage()    {}
func (*VoteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_3368f8b6bbfc2668, []int{0}
}
func (m *VoteExtension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteExtension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteExtension.Merge(m, src)
}
func (m *VoteExtension) XXX_Size() int {
	return m.Size()
}
func (m *VoteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_VoteExtension proto.InternalMessageInfo

func (m *VoteExtension) GetData() []*PublishedData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *VoteExtension) GetShares() []*DataShares {
	if m != nil {
		return m.Shares
	}
	return nil
}

func init() {
	proto.RegisterType((*VoteExtension)(nil), "sunrise.da.VoteExtension")
}

func init() { proto.RegisterFile("sunrise/da/vote_extension.proto", fileDescriptor_3368f8b6bbfc2668) }

var fileDescriptor_3368f8b6bbfc2668 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x1b, 0x95, 0x1e, 0x22, 0x1e, 0x0c, 0xa2, 0x6d, 0x0f, 0xab, 0x78, 0x12, 0xc1, 0x5d,
	0xd4, 0x37, 0x90, 0x7a, 0x97, 0x0a, 0x1e, 0xbc, 0x84, 0x49, 0xb3, 0x24, 0x81, 0x26, 0x13, 0x76,
	0x36, 0xa5, 0x7d, 0x0b, 0x1f, 0xcb, 0x63, 0x8f, 0x1e, 0x25, 0x79, 0x11, 0xd9, 0x7f, 0x98, 0xcb,
	0xee, 0xce, 0x7e, 0xdf, 0x0c, 0xbf, 0x89, 0xaf, 0xa9, 0x6b, 0x54, 0x45, 0x52, 0xe4, 0x20, 0xb6,
	0xa8, 0x65, 0x2a, 0x77, 0x5a, 0x36, 0x54, 0x61, 0xc3, 0x5b, 0x85, 0x1a, 0x93, 0xd8, 0x0b, 0x3c,
	0x87, 0xc5, 0x39, 0xd4, 0x55, 0x83, 0xc2, 0x9e, 0x0e, 0x2f, 0xae, 0xd6, 0x48, 0x35, 0x92, 0xa8,
	0xa9, 0x10, 0xdb, 0x47, 0x73, 0x79, 0x30, 0x77, 0x20, 0xb5, 0x95, 0x70, 0x85, 0x47, 0x17, 0x05,
	0x16, 0xe8, 0xfe, 0xcd, 0x2b, 0x4c, 0x1a, 0x25, 0x69, 0x41, 0x41, 0x1d, 0xf4, 0x71, 0xc4, 0xb6,
	0xcb, 0x36, 0x15, 0x95, 0x32, 0x4f, 0x73, 0xd0, 0xe0, 0x84, 0xdb, 0x26, 0x3e, 0xfb, 0x40, 0x2d,
	0x5f, 0x43, 0xf2, 0xe4, 0x21, 0x3e, 0x31, 0x78, 0x16, 0xdd, 0x1c, 0xdf, 0x9d, 0x3e, 0xcd, 0xf9,
	0xff, 0x0a, 0xfc, 0x2d, 0x0c, 0x58, 0x82, 0x86, 0x95, 0xd5, 0x12, 0x1e, 0x4f, 0xa9, 0x04, 0x25,
	0x69, 0x76, 0x64, 0x1b, 0x2e, 0xc7, 0x0d, 0xc6, 0x7b, 0xb7, 0x74, 0xe5, 0xad, 0x97, 0xe5, 0x77,
	0xcf, 0xa2, 0x43, 0xcf, 0xa2, 0xdf, 0x9e, 0x45, 0x5f, 0x03, 0x9b, 0x1c, 0x06, 0x36, 0xf9, 0x19,
	0xd8, 0xe4, 0xf3, 0xbe, 0xa8, 0x74, 0xd9, 0x65, 0x7c, 0x8d, 0xb5, 0xf0, 0x33, 0x36, 0xb0, 0x97,
	0x2a, 0x14, 0x62, 0x67, 0x96, 0xd0, 0xfb, 0x56, 0x52, 0x36, 0xb5, 0xe1, 0x9f, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x15, 0x63, 0xe9, 0x82, 0x01, 0x00, 0x00,
}

func (m *VoteExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteExtension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteExtension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVoteExtension(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVoteExtension(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintVoteExtension(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoteExtension(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteExtension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovVoteExtension(uint64(l))
		}
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovVoteExtension(uint64(l))
		}
	}
	return n
}

func sovVoteExtension(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoteExtension(x uint64) (n int) {
	return sovVoteExtension(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteExtension
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
			return fmt.Errorf("proto: VoteExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteExtension
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
				return ErrInvalidLengthVoteExtension
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVoteExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &PublishedData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteExtension
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
				return ErrInvalidLengthVoteExtension
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVoteExtension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, &DataShares{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteExtension(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteExtension
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
func skipVoteExtension(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoteExtension
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
					return 0, ErrIntOverflowVoteExtension
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
					return 0, ErrIntOverflowVoteExtension
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
				return 0, ErrInvalidLengthVoteExtension
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoteExtension
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoteExtension
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoteExtension        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoteExtension          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoteExtension = fmt.Errorf("proto: unexpected end of group")
)
