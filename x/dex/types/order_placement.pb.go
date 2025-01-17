// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/order_placement.proto

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

type OrderPlacement struct {
	Long       bool   `protobuf:"varint,1,opt,name=long,proto3" json:"long,omitempty"`
	Price      uint64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Quantity   uint64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PriceDenom string `protobuf:"bytes,4,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	AssetDenom string `protobuf:"bytes,5,opt,name=assetDenom,proto3" json:"assetDenom,omitempty"`
	Open       bool   `protobuf:"varint,6,opt,name=open,proto3" json:"open,omitempty"`
	Limit      bool   `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Leverage   string `protobuf:"bytes,8,opt,name=leverage,proto3" json:"leverage,omitempty"`
}

func (m *OrderPlacement) Reset()         { *m = OrderPlacement{} }
func (m *OrderPlacement) String() string { return proto.CompactTextString(m) }
func (*OrderPlacement) ProtoMessage()    {}
func (*OrderPlacement) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2471ad330c401e, []int{0}
}
func (m *OrderPlacement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderPlacement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderPlacement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderPlacement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPlacement.Merge(m, src)
}
func (m *OrderPlacement) XXX_Size() int {
	return m.Size()
}
func (m *OrderPlacement) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPlacement.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPlacement proto.InternalMessageInfo

func (m *OrderPlacement) GetLong() bool {
	if m != nil {
		return m.Long
	}
	return false
}

func (m *OrderPlacement) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderPlacement) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderPlacement) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *OrderPlacement) GetAssetDenom() string {
	if m != nil {
		return m.AssetDenom
	}
	return ""
}

func (m *OrderPlacement) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *OrderPlacement) GetLimit() bool {
	if m != nil {
		return m.Limit
	}
	return false
}

func (m *OrderPlacement) GetLeverage() string {
	if m != nil {
		return m.Leverage
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderPlacement)(nil), "seiprotocol.seichain.dex.OrderPlacement")
}

func init() { proto.RegisterFile("dex/order_placement.proto", fileDescriptor_ca2471ad330c401e) }

var fileDescriptor_ca2471ad330c401e = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x63, 0x48, 0x4b, 0xf0, 0xc0, 0x60, 0x31, 0x18, 0x06, 0x2b, 0x62, 0xca, 0xd2, 0x64,
	0xe0, 0x0d, 0x10, 0x12, 0x23, 0x28, 0x23, 0x0b, 0x4a, 0x93, 0x5f, 0xa9, 0x25, 0xc7, 0x36, 0xb6,
	0x8b, 0xd2, 0xb7, 0xe0, 0xb1, 0x18, 0x3b, 0x32, 0x30, 0xa0, 0xe4, 0x45, 0x90, 0x6d, 0xb5, 0x74,
	0xbb, 0xbb, 0xcf, 0xbf, 0x75, 0x3a, 0x7c, 0xd3, 0xc1, 0x58, 0x29, 0xd3, 0x81, 0x79, 0xd3, 0xa2,
	0x69, 0x61, 0x00, 0xe9, 0x4a, 0x6d, 0x94, 0x53, 0x84, 0x5a, 0xe0, 0x41, 0xb5, 0x4a, 0x94, 0x16,
	0x78, 0xbb, 0x69, 0xb8, 0x2c, 0x3b, 0x18, 0xef, 0x7e, 0x10, 0xbe, 0x7a, 0xf6, 0x37, 0x2f, 0x87,
	0x13, 0x42, 0x70, 0x2a, 0x94, 0xec, 0x29, 0xca, 0x51, 0x91, 0xd5, 0x41, 0x93, 0x6b, 0xbc, 0xd0,
	0x86, 0xb7, 0x40, 0xcf, 0x72, 0x54, 0xa4, 0x75, 0x34, 0xe4, 0x16, 0x67, 0xef, 0xdb, 0x46, 0x3a,
	0xee, 0x76, 0xf4, 0x3c, 0x80, 0xa3, 0x27, 0x0c, 0xe3, 0xf0, 0xe8, 0x11, 0xa4, 0x1a, 0x68, 0x9a,
	0xa3, 0xe2, 0xb2, 0x3e, 0x49, 0x3c, 0x6f, 0xac, 0x05, 0x17, 0xf9, 0x22, 0xf2, 0xff, 0xc4, 0xb7,
	0x50, 0x1a, 0x24, 0x5d, 0xc6, 0x16, 0x5e, 0xfb, 0x16, 0x82, 0x0f, 0xdc, 0xd1, 0x8b, 0x10, 0x46,
	0xe3, 0x5b, 0x08, 0xf8, 0x00, 0xd3, 0xf4, 0x40, 0xb3, 0xf0, 0xcf, 0xd1, 0x3f, 0x3c, 0x7d, 0x4d,
	0x0c, 0xed, 0x27, 0x86, 0x7e, 0x27, 0x86, 0x3e, 0x67, 0x96, 0xec, 0x67, 0x96, 0x7c, 0xcf, 0x2c,
	0x79, 0x5d, 0xf5, 0xdc, 0x6d, 0xb6, 0xeb, 0xb2, 0x55, 0x43, 0x65, 0x81, 0xaf, 0x0e, 0xf3, 0x04,
	0x13, 0xf6, 0xa9, 0xc6, 0xca, 0x2f, 0xea, 0x76, 0x1a, 0xec, 0x7a, 0x19, 0xf8, 0xfd, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x70, 0x07, 0x95, 0x65, 0x01, 0x00, 0x00,
}

func (m *OrderPlacement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderPlacement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderPlacement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leverage) > 0 {
		i -= len(m.Leverage)
		copy(dAtA[i:], m.Leverage)
		i = encodeVarintOrderPlacement(dAtA, i, uint64(len(m.Leverage)))
		i--
		dAtA[i] = 0x42
	}
	if m.Limit {
		i--
		if m.Limit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Open {
		i--
		if m.Open {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.AssetDenom) > 0 {
		i -= len(m.AssetDenom)
		copy(dAtA[i:], m.AssetDenom)
		i = encodeVarintOrderPlacement(dAtA, i, uint64(len(m.AssetDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintOrderPlacement(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x22
	}
	if m.Quantity != 0 {
		i = encodeVarintOrderPlacement(dAtA, i, uint64(m.Quantity))
		i--
		dAtA[i] = 0x18
	}
	if m.Price != 0 {
		i = encodeVarintOrderPlacement(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x10
	}
	if m.Long {
		i--
		if m.Long {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderPlacement(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderPlacement(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderPlacement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Long {
		n += 2
	}
	if m.Price != 0 {
		n += 1 + sovOrderPlacement(uint64(m.Price))
	}
	if m.Quantity != 0 {
		n += 1 + sovOrderPlacement(uint64(m.Quantity))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovOrderPlacement(uint64(l))
	}
	l = len(m.AssetDenom)
	if l > 0 {
		n += 1 + l + sovOrderPlacement(uint64(l))
	}
	if m.Open {
		n += 2
	}
	if m.Limit {
		n += 2
	}
	l = len(m.Leverage)
	if l > 0 {
		n += 1 + l + sovOrderPlacement(uint64(l))
	}
	return n
}

func sovOrderPlacement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderPlacement(x uint64) (n int) {
	return sovOrderPlacement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderPlacement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderPlacement
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
			return fmt.Errorf("proto: OrderPlacement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderPlacement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Long", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Long = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			m.Quantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
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
				return ErrInvalidLengthOrderPlacement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderPlacement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
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
				return ErrInvalidLengthOrderPlacement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderPlacement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Open", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Open = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Limit = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderPlacement
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
				return ErrInvalidLengthOrderPlacement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderPlacement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leverage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrderPlacement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderPlacement
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
func skipOrderPlacement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderPlacement
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
					return 0, ErrIntOverflowOrderPlacement
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
					return 0, ErrIntOverflowOrderPlacement
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
				return 0, ErrInvalidLengthOrderPlacement
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderPlacement
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderPlacement
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderPlacement        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderPlacement          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderPlacement = fmt.Errorf("proto: unexpected end of group")
)
