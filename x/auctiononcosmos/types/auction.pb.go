// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auctiononcosmos/auctiononcosmos/auction.proto

package types

import (
	fmt "fmt"
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

type Auction struct {
	Index           uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	AuctionOwner    string `protobuf:"bytes,2,opt,name=auctionOwner,proto3" json:"auctionOwner,omitempty"`
	ItemDescription string `protobuf:"bytes,3,opt,name=itemDescription,proto3" json:"itemDescription,omitempty"`
	InitialPrice    uint64 `protobuf:"varint,4,opt,name=initialPrice,proto3" json:"initialPrice,omitempty"`
	CreatedAt       uint64 `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	EndTime         uint64 `protobuf:"varint,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
	AuctionStatus   string `protobuf:"bytes,7,opt,name=auctionStatus,proto3" json:"auctionStatus,omitempty"`
}

func (m *Auction) Reset()         { *m = Auction{} }
func (m *Auction) String() string { return proto.CompactTextString(m) }
func (*Auction) ProtoMessage()    {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_38ceb729739022b8, []int{0}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

func (m *Auction) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Auction) GetAuctionOwner() string {
	if m != nil {
		return m.AuctionOwner
	}
	return ""
}

func (m *Auction) GetItemDescription() string {
	if m != nil {
		return m.ItemDescription
	}
	return ""
}

func (m *Auction) GetInitialPrice() uint64 {
	if m != nil {
		return m.InitialPrice
	}
	return 0
}

func (m *Auction) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Auction) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Auction) GetAuctionStatus() string {
	if m != nil {
		return m.AuctionStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*Auction)(nil), "auctiononcosmos.auctiononcosmos.Auction")
}

func init() {
	proto.RegisterFile("auctiononcosmos/auctiononcosmos/auction.proto", fileDescriptor_38ceb729739022b8)
}

var fileDescriptor_38ceb729739022b8 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x2c, 0x4d, 0x2e,
	0xc9, 0xcc, 0xcf, 0xcb, 0xcf, 0x4b, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0xc7, 0xc1, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x47, 0x93, 0xd6, 0x43, 0xe3, 0x2b, 0xbd, 0x67, 0xe4, 0x62,
	0x77, 0x84, 0x88, 0x09, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x04, 0x41, 0x38, 0x42, 0x4a, 0x5c, 0x3c, 0x50, 0x4d, 0xfe, 0xe5, 0x79, 0xa9, 0x45,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x28, 0x62, 0x42, 0x1a, 0x5c, 0xfc, 0x99, 0x25, 0xa9,
	0xb9, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x20, 0x71, 0x09, 0x66, 0xb0, 0x32, 0x74, 0x61,
	0x90, 0x69, 0x99, 0x79, 0x99, 0x25, 0x99, 0x89, 0x39, 0x01, 0x45, 0x99, 0xc9, 0xa9, 0x12, 0x2c,
	0x60, 0xab, 0x50, 0xc4, 0x84, 0x64, 0xb8, 0x38, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x53, 0x1c,
	0x4b, 0x24, 0x58, 0xc1, 0x0a, 0x10, 0x02, 0x42, 0x12, 0x5c, 0xec, 0xa9, 0x79, 0x29, 0x21, 0x99,
	0xb9, 0xa9, 0x12, 0x6c, 0x60, 0x39, 0x18, 0x57, 0x48, 0x85, 0x8b, 0x17, 0xea, 0xaa, 0xe0, 0x92,
	0xc4, 0x92, 0xd2, 0x62, 0x09, 0x76, 0xb0, 0x1b, 0x50, 0x05, 0x9d, 0x22, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x2e, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0xdf, 0x39, 0xb1, 0xa0, 0x24, 0xbf, 0x38, 0x23, 0x3b, 0xd1, 0xc7, 0x51, 0x1f, 0x1a,
	0x3e, 0xfe, 0x79, 0xce, 0x90, 0x20, 0xae, 0xc0, 0x08, 0xf4, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24,
	0x36, 0x70, 0x98, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x03, 0xde, 0xf7, 0xee, 0xa4, 0x01,
	0x00, 0x00,
}

func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuctionStatus) > 0 {
		i -= len(m.AuctionStatus)
		copy(dAtA[i:], m.AuctionStatus)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.AuctionStatus)))
		i--
		dAtA[i] = 0x3a
	}
	if m.EndTime != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x30
	}
	if m.CreatedAt != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x28
	}
	if m.InitialPrice != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.InitialPrice))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ItemDescription) > 0 {
		i -= len(m.ItemDescription)
		copy(dAtA[i:], m.ItemDescription)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.ItemDescription)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AuctionOwner) > 0 {
		i -= len(m.AuctionOwner)
		copy(dAtA[i:], m.AuctionOwner)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.AuctionOwner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovAuction(uint64(m.Index))
	}
	l = len(m.AuctionOwner)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = len(m.ItemDescription)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.InitialPrice != 0 {
		n += 1 + sovAuction(uint64(m.InitialPrice))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovAuction(uint64(m.CreatedAt))
	}
	if m.EndTime != 0 {
		n += 1 + sovAuction(uint64(m.EndTime))
	}
	l = len(m.AuctionStatus)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemDescription", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemDescription = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPrice", wireType)
			}
			m.InitialPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
