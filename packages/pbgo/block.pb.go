// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

package pbgo

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

//BlockHeader is a structure of the block's header
type BlockHeader struct {
	BlockId       int64  `protobuf:"varint,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp     int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EcosystemId   int64  `protobuf:"varint,3,opt,name=ecosystem_id,json=ecosystemId,proto3" json:"ecosystem_id,omitempty"`
	KeyId         int64  `protobuf:"varint,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	NodePosition  int64  `protobuf:"varint,5,opt,name=node_position,json=nodePosition,proto3" json:"node_position,omitempty"`
	Sign          []byte `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	BlockHash     []byte `protobuf:"bytes,7,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	RollbacksHash []byte `protobuf:"bytes,8,opt,name=rollbacks_hash,json=rollbacksHash,proto3" json:"rollbacks_hash,omitempty"`
	Version       int32  `protobuf:"varint,9,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetBlockId() int64 {
	if m != nil {
		return m.BlockId
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetEcosystemId() int64 {
	if m != nil {
		return m.EcosystemId
	}
	return 0
}

func (m *BlockHeader) GetKeyId() int64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *BlockHeader) GetNodePosition() int64 {
	if m != nil {
		return m.NodePosition
	}
	return 0
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *BlockHeader) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *BlockHeader) GetRollbacksHash() []byte {
	if m != nil {
		return m.RollbacksHash
	}
	return nil
}

func (m *BlockHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "pbgo.BlockHeader")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0xeb, 0xfe, 0xf7, 0xda, 0x32, 0x58, 0x42, 0x0a, 0x12, 0x44, 0x05, 0x84, 0x94, 0xa5,
	0xcd, 0xc0, 0x13, 0x90, 0xa9, 0xd9, 0x50, 0x27, 0xc4, 0x52, 0x39, 0xb1, 0x95, 0x58, 0x69, 0x7a,
	0x51, 0x6c, 0x10, 0x7d, 0x0b, 0x56, 0xde, 0x88, 0xb1, 0x23, 0x23, 0x6a, 0x5e, 0x04, 0xf9, 0x0a,
	0x65, 0xf3, 0x7d, 0xdf, 0x67, 0xd9, 0x3a, 0x18, 0x27, 0x1b, 0x4c, 0x8b, 0x45, 0x55, 0xa3, 0x45,
	0xde, 0xad, 0x92, 0x0c, 0x6f, 0x3e, 0xda, 0x30, 0x8e, 0x1c, 0x5d, 0x2a, 0x21, 0x55, 0xcd, 0x2f,
	0x60, 0x48, 0xd1, 0x5a, 0x4b, 0x8f, 0xcd, 0x58, 0xd0, 0x59, 0x0d, 0x68, 0x8e, 0x25, 0xbf, 0x84,
	0x91, 0xd5, 0xa5, 0x32, 0x56, 0x94, 0x95, 0xd7, 0x26, 0xf7, 0x0f, 0xf8, 0x35, 0x4c, 0x54, 0x8a,
	0x66, 0x67, 0xac, 0x2a, 0xdd, 0xe5, 0x0e, 0x05, 0xe3, 0x13, 0x8b, 0x25, 0x3f, 0x87, 0x7e, 0xa1,
	0x76, 0x4e, 0x76, 0x49, 0xf6, 0x0a, 0xb5, 0x8b, 0x25, 0xbf, 0x85, 0xe9, 0x16, 0xa5, 0x5a, 0x57,
	0x68, 0xb4, 0xd5, 0xb8, 0xf5, 0x7a, 0x64, 0x27, 0x0e, 0x3e, 0xfe, 0x32, 0xce, 0xa1, 0x6b, 0x74,
	0xb6, 0xf5, 0xfa, 0x33, 0x16, 0x4c, 0x56, 0x74, 0xe6, 0x57, 0x00, 0xc7, 0xbf, 0xe6, 0xc2, 0xe4,
	0xde, 0x80, 0xcc, 0x88, 0xc8, 0x52, 0x98, 0x9c, 0xdf, 0xc1, 0x59, 0x8d, 0x9b, 0x4d, 0x22, 0xd2,
	0xc2, 0x1c, 0x93, 0x21, 0x25, 0xd3, 0x13, 0xa5, 0xcc, 0x83, 0xc1, 0xab, 0xaa, 0x8d, 0x7b, 0x78,
	0x34, 0x63, 0x41, 0x6f, 0xf5, 0x37, 0x46, 0xd1, 0xe7, 0xc1, 0x67, 0xfb, 0x83, 0xcf, 0xbe, 0x0f,
	0x3e, 0x7b, 0x6f, 0xfc, 0xd6, 0xbe, 0xf1, 0x5b, 0x5f, 0x8d, 0xdf, 0x7a, 0x0e, 0x32, 0x6d, 0xf3,
	0x97, 0x64, 0x91, 0x62, 0x19, 0xc6, 0xd1, 0xc3, 0xd3, 0x5c, 0x63, 0x98, 0xe1, 0x5c, 0x27, 0xe2,
	0x2d, 0xac, 0x44, 0x5a, 0x88, 0x4c, 0x99, 0xd0, 0xed, 0x37, 0xe9, 0xd3, 0xb2, 0xef, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0xe2, 0x98, 0x91, 0x7b, 0x01, 0x00, 0x00,
}

func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x48
	}
	if len(m.RollbacksHash) > 0 {
		i -= len(m.RollbacksHash)
		copy(dAtA[i:], m.RollbacksHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.RollbacksHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0x32
	}
	if m.NodePosition != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.NodePosition))
		i--
		dAtA[i] = 0x28
	}
	if m.KeyId != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x20
	}
	if m.EcosystemId != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.EcosystemId))
		i--
		dAtA[i] = 0x18
	}
	if m.Timestamp != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockId != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.BlockId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockId != 0 {
		n += 1 + sovBlock(uint64(m.BlockId))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	if m.EcosystemId != 0 {
		n += 1 + sovBlock(uint64(m.EcosystemId))
	}
	if m.KeyId != 0 {
		n += 1 + sovBlock(uint64(m.KeyId))
	}
	if m.NodePosition != 0 {
		n += 1 + sovBlock(uint64(m.NodePosition))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.RollbacksHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovBlock(uint64(m.Version))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			m.BlockId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EcosystemId", wireType)
			}
			m.EcosystemId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EcosystemId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodePosition", wireType)
			}
			m.NodePosition = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodePosition |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollbacksHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollbacksHash = append(m.RollbacksHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RollbacksHash == nil {
				m.RollbacksHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)