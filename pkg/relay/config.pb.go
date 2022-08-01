// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package relay

import (
	fmt "fmt"
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

type ChainConfig struct {
	ChainId    string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	EthChainId int64  `protobuf:"varint,2,opt,name=eth_chain_id,json=ethChainId,proto3" json:"eth_chain_id,omitempty"`
	RpcAddr    string `protobuf:"bytes,3,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	// use for relayer
	HdwMnemonic       string `protobuf:"bytes,4,opt,name=hdw_mnemonic,json=hdwMnemonic,proto3" json:"hdw_mnemonic,omitempty"`
	HdwPath           string `protobuf:"bytes,5,opt,name=hdw_path,json=hdwPath,proto3" json:"hdw_path,omitempty"`
	IbcHostAddress    string `protobuf:"bytes,6,opt,name=ibc_host_address,json=ibcHostAddress,proto3" json:"ibc_host_address,omitempty"`
	IbcHandlerAddress string `protobuf:"bytes,7,opt,name=ibc_handler_address,json=ibcHandlerAddress,proto3" json:"ibc_handler_address,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainConfig)(nil), "relayer.chains.ethereum.config.ChainConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x06, 0xe0, 0x98, 0x42, 0x5b, 0xdc, 0x0a, 0x41, 0x60, 0x48, 0x19, 0xac, 0xc2, 0xd4, 0xa5,
	0xc9, 0x00, 0x17, 0x28, 0x5d, 0x60, 0x40, 0x42, 0x1d, 0x59, 0x22, 0xc7, 0xcf, 0xc4, 0x96, 0x92,
	0x38, 0xb2, 0x1d, 0x55, 0xbd, 0x05, 0xc7, 0xea, 0xd8, 0x91, 0x11, 0xda, 0x73, 0x20, 0xa1, 0xd8,
	0xa1, 0x9b, 0xfd, 0xde, 0xf7, 0x5b, 0xf2, 0x8f, 0xc7, 0x4c, 0x55, 0x1f, 0x32, 0x8f, 0x6b, 0xad,
	0xac, 0x0a, 0x89, 0xe6, 0x05, 0xdd, 0x70, 0x1d, 0x33, 0x41, 0x65, 0x65, 0x62, 0x6e, 0x05, 0xd7,
	0xbc, 0x29, 0x63, 0xaf, 0x6e, 0x6f, 0x72, 0x95, 0x2b, 0x47, 0x93, 0xf6, 0xe4, 0x53, 0xf7, 0xbf,
	0x08, 0x8f, 0x96, 0x6d, 0x60, 0xe9, 0x54, 0x38, 0xc1, 0x43, 0x97, 0x4f, 0x25, 0x44, 0x68, 0x8a,
	0x66, 0xe7, 0xab, 0x81, 0xbb, 0xbf, 0x40, 0x38, 0xc5, 0x63, 0x6e, 0x45, 0x7a, 0x5c, 0x9f, 0x4c,
	0xd1, 0xac, 0xb7, 0xc2, 0xdc, 0x8a, 0x65, 0x27, 0x26, 0x78, 0xa8, 0x6b, 0x96, 0x52, 0x00, 0x1d,
	0xf5, 0x7c, 0x58, 0xd7, 0x6c, 0x01, 0xa0, 0xc3, 0x3b, 0x3c, 0x16, 0xb0, 0x4e, 0xcb, 0x8a, 0x97,
	0xaa, 0x92, 0x2c, 0x3a, 0x75, 0xeb, 0x91, 0x80, 0xf5, 0x6b, 0x37, 0x6a, 0xd3, 0x2d, 0xa9, 0xa9,
	0x15, 0xd1, 0x99, 0x4f, 0x0b, 0x58, 0xbf, 0x51, 0x2b, 0xc2, 0x19, 0xbe, 0x94, 0x19, 0x4b, 0x85,
	0x32, 0xd6, 0xbd, 0xce, 0x8d, 0x89, 0xfa, 0x8e, 0x5c, 0xc8, 0x8c, 0x3d, 0x2b, 0x63, 0x17, 0x7e,
	0x1a, 0xc6, 0xf8, 0xda, 0x49, 0x5a, 0x41, 0xc1, 0xf5, 0x11, 0x0f, 0x1c, 0xbe, 0x6a, 0xb1, 0xdf,
	0x74, 0xfe, 0x69, 0xb5, 0xfd, 0x21, 0xc1, 0x76, 0x4f, 0xd0, 0x6e, 0x4f, 0xd0, 0xf7, 0x9e, 0xa0,
	0xcf, 0x03, 0x09, 0x76, 0x07, 0x12, 0x7c, 0x1d, 0x48, 0xf0, 0xfe, 0x98, 0x4b, 0x2b, 0x9a, 0x2c,
	0x66, 0xaa, 0x4c, 0x80, 0x5a, 0xea, 0xfe, 0x5e, 0xd0, 0x2c, 0xd9, 0x34, 0x72, 0xde, 0xf5, 0x3d,
	0x2f, 0x15, 0x34, 0x05, 0x37, 0xc9, 0x7f, 0xe1, 0x59, 0xdf, 0x55, 0xfb, 0xf0, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x9d, 0x84, 0x52, 0x2e, 0xa0, 0x01, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcHandlerAddress) > 0 {
		i -= len(m.IbcHandlerAddress)
		copy(dAtA[i:], m.IbcHandlerAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcHandlerAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.IbcHostAddress) > 0 {
		i -= len(m.IbcHostAddress)
		copy(dAtA[i:], m.IbcHostAddress)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcHostAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.HdwPath) > 0 {
		i -= len(m.HdwPath)
		copy(dAtA[i:], m.HdwPath)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.HdwPath)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.HdwMnemonic) > 0 {
		i -= len(m.HdwMnemonic)
		copy(dAtA[i:], m.HdwMnemonic)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.HdwMnemonic)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RpcAddr) > 0 {
		i -= len(m.RpcAddr)
		copy(dAtA[i:], m.RpcAddr)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RpcAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.EthChainId != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.EthChainId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.EthChainId != 0 {
		n += 1 + sovConfig(uint64(m.EthChainId))
	}
	l = len(m.RpcAddr)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.HdwMnemonic)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.HdwPath)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcHostAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.IbcHandlerAddress)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthChainId", wireType)
			}
			m.EthChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RpcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HdwMnemonic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HdwMnemonic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HdwPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HdwPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcHostAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcHostAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcHandlerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcHandlerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
