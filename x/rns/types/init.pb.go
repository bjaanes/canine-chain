// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rns/init.proto

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

type Init struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Complete bool   `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
}

func (m *Init) Reset()         { *m = Init{} }
func (m *Init) String() string { return proto.CompactTextString(m) }
func (*Init) ProtoMessage()    {}
func (*Init) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f4e519bccb6bfd, []int{0}
}
func (m *Init) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Init) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Init.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Init) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Init.Merge(m, src)
}
func (m *Init) XXX_Size() int {
	return m.Size()
}
func (m *Init) XXX_DiscardUnknown() {
	xxx_messageInfo_Init.DiscardUnknown(m)
}

var xxx_messageInfo_Init proto.InternalMessageInfo

func (m *Init) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Init) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func init() {
	proto.RegisterType((*Init)(nil), "jackaldao.canine.rns.Init")
}

func init() { proto.RegisterFile("rns/init.proto", fileDescriptor_40f4e519bccb6bfd) }

var fileDescriptor_40f4e519bccb6bfd = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xca, 0x2b, 0xd6,
	0xcf, 0xcc, 0xcb, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc9, 0x4a, 0x4c, 0xce,
	0x4e, 0xcc, 0x49, 0x49, 0xcc, 0xd7, 0x4b, 0x4e, 0xcc, 0xcb, 0xcc, 0x4b, 0xd5, 0x2b, 0xca, 0x2b,
	0x56, 0xb2, 0xe1, 0x62, 0xf1, 0xcc, 0xcb, 0x2c, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xa4, 0xb8, 0x38,
	0x92, 0xf3, 0x73, 0x0b, 0x72, 0x52, 0x4b, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x82, 0xe0,
	0x7c, 0x27, 0xa7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x48, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x58, 0xac, 0x9b, 0x92, 0x98, 0xaf,
	0x0f, 0xb1, 0x59, 0xbf, 0x42, 0x1f, 0xe4, 0xb8, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0,
	0xf3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xdb, 0x39, 0xdc, 0xb0, 0x00, 0x00, 0x00,
}

func (m *Init) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Init) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Init) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Complete {
		i--
		if m.Complete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintInit(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInit(dAtA []byte, offset int, v uint64) int {
	offset -= sovInit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Init) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovInit(uint64(l))
	}
	if m.Complete {
		n += 2
	}
	return n
}

func sovInit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInit(x uint64) (n int) {
	return sovInit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Init) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInit
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
			return fmt.Errorf("proto: Init: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Init: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInit
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
				return ErrInvalidLengthInit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Complete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInit
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
			m.Complete = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipInit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInit
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
func skipInit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInit
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
					return 0, ErrIntOverflowInit
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
					return 0, ErrIntOverflowInit
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
				return 0, ErrInvalidLengthInit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInit = fmt.Errorf("proto: unexpected end of group")
)
