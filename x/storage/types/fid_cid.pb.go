// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/fid_cid.proto

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

type FidCid struct {
	Fid  string `protobuf:"bytes,1,opt,name=fid,proto3" json:"fid,omitempty"`
	Cids string `protobuf:"bytes,2,opt,name=cids,proto3" json:"cids,omitempty"`
}

func (m *FidCid) Reset()         { *m = FidCid{} }
func (m *FidCid) String() string { return proto.CompactTextString(m) }
func (*FidCid) ProtoMessage()    {}
func (*FidCid) Descriptor() ([]byte, []int) {
	return fileDescriptor_a811c4f39ed19f5a, []int{0}
}
func (m *FidCid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FidCid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FidCid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FidCid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FidCid.Merge(m, src)
}
func (m *FidCid) XXX_Size() int {
	return m.Size()
}
func (m *FidCid) XXX_DiscardUnknown() {
	xxx_messageInfo_FidCid.DiscardUnknown(m)
}

var xxx_messageInfo_FidCid proto.InternalMessageInfo

func (m *FidCid) GetFid() string {
	if m != nil {
		return m.Fid
	}
	return ""
}

func (m *FidCid) GetCids() string {
	if m != nil {
		return m.Cids
	}
	return ""
}

func init() {
	proto.RegisterType((*FidCid)(nil), "jackaldao.canine.storage.FidCid")
}

func init() { proto.RegisterFile("storage/fid_cid.proto", fileDescriptor_a811c4f39ed19f5a) }

var fileDescriptor_a811c4f39ed19f5a = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x4f, 0xcb, 0x4c, 0x89, 0x4f, 0xce, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0xc8, 0x4a, 0x4c, 0xce, 0x4e, 0xcc, 0x49, 0x49, 0xcc, 0xd7, 0x4b, 0x4e, 0xcc,
	0xcb, 0xcc, 0x4b, 0xd5, 0x83, 0xaa, 0x53, 0xd2, 0xe3, 0x62, 0x73, 0xcb, 0x4c, 0x71, 0xce, 0x4c,
	0x11, 0x12, 0xe0, 0x62, 0x4e, 0xcb, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31,
	0x85, 0x84, 0xb8, 0x58, 0x92, 0x33, 0x53, 0x8a, 0x25, 0x98, 0xc0, 0x42, 0x60, 0xb6, 0x93, 0xdb,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0xac, 0xd3, 0x4d, 0x49, 0xcc, 0xd7, 0x87, 0xd8, 0xa7,
	0x5f, 0xa1, 0x0f, 0x73, 0x59, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x61, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xb2, 0xc7, 0x02, 0xb1, 0x00, 0x00, 0x00,
}

func (m *FidCid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FidCid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FidCid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cids) > 0 {
		i -= len(m.Cids)
		copy(dAtA[i:], m.Cids)
		i = encodeVarintFidCid(dAtA, i, uint64(len(m.Cids)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Fid) > 0 {
		i -= len(m.Fid)
		copy(dAtA[i:], m.Fid)
		i = encodeVarintFidCid(dAtA, i, uint64(len(m.Fid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFidCid(dAtA []byte, offset int, v uint64) int {
	offset -= sovFidCid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FidCid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fid)
	if l > 0 {
		n += 1 + l + sovFidCid(uint64(l))
	}
	l = len(m.Cids)
	if l > 0 {
		n += 1 + l + sovFidCid(uint64(l))
	}
	return n
}

func sovFidCid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFidCid(x uint64) (n int) {
	return sovFidCid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FidCid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFidCid
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
			return fmt.Errorf("proto: FidCid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FidCid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFidCid
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
				return ErrInvalidLengthFidCid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFidCid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cids", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFidCid
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
				return ErrInvalidLengthFidCid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFidCid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cids = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFidCid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFidCid
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
func skipFidCid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFidCid
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
					return 0, ErrIntOverflowFidCid
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
					return 0, ErrIntOverflowFidCid
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
				return 0, ErrInvalidLengthFidCid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFidCid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFidCid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFidCid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFidCid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFidCid = fmt.Errorf("proto: unexpected end of group")
)
