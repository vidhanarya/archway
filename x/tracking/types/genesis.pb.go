// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/tracking/v1beta1/genesis.proto

package types

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

// GenesisState defines the initial state of the tracking module.
type GenesisState struct {
	// tx_info_last_id defines the last unique ID for a TxInfo objs.
	TxInfoLastId uint64 `protobuf:"varint,1,opt,name=tx_info_last_id,json=txInfoLastId,proto3" json:"tx_info_last_id,omitempty"`
	// tx_infos defines a list of all the tracked transactions.
	TxInfos []TxInfo `protobuf:"bytes,2,rep,name=tx_infos,json=txInfos,proto3" json:"tx_infos"`
	// contract_op_info_last_id defines the last unique ID for ContractOperationInfo objs.
	ContractOpInfoLastId uint64 `protobuf:"varint,3,opt,name=contract_op_info_last_id,json=contractOpInfoLastId,proto3" json:"contract_op_info_last_id,omitempty"`
	// contract_op_infos defines a list of all the tracked contract operations.
	ContractOpInfos []ContractOperationInfo `protobuf:"bytes,4,rep,name=contract_op_infos,json=contractOpInfos,proto3" json:"contract_op_infos"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f994a880eb2f25, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetTxInfoLastId() uint64 {
	if m != nil {
		return m.TxInfoLastId
	}
	return 0
}

func (m *GenesisState) GetTxInfos() []TxInfo {
	if m != nil {
		return m.TxInfos
	}
	return nil
}

func (m *GenesisState) GetContractOpInfoLastId() uint64 {
	if m != nil {
		return m.ContractOpInfoLastId
	}
	return 0
}

func (m *GenesisState) GetContractOpInfos() []ContractOperationInfo {
	if m != nil {
		return m.ContractOpInfos
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "archway.tracking.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("archway/tracking/v1beta1/genesis.proto", fileDescriptor_40f994a880eb2f25)
}

var fileDescriptor_40f994a880eb2f25 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xdb, 0x6d, 0xa8, 0xc4, 0xc1, 0xb0, 0xec, 0x50, 0x76, 0x88, 0x43, 0x50, 0x77, 0x31,
	0x61, 0x0e, 0xbc, 0x3b, 0x0f, 0x32, 0x50, 0x84, 0xe9, 0xc9, 0x4b, 0xc9, 0xba, 0xac, 0x0b, 0xd3,
	0xa4, 0x34, 0x4f, 0xd7, 0x7d, 0x0b, 0xc1, 0x2f, 0xb5, 0xe3, 0x8e, 0x9e, 0x44, 0xda, 0x2f, 0x22,
	0x4b, 0xdb, 0xb9, 0x09, 0xbd, 0x25, 0xbc, 0xdf, 0xfb, 0xff, 0xfe, 0x3c, 0x74, 0xc6, 0x22, 0x7f,
	0x3a, 0x67, 0x0b, 0x0a, 0x11, 0xf3, 0x67, 0x42, 0x06, 0xf4, 0xbd, 0x3b, 0xe2, 0xc0, 0xba, 0x34,
	0xe0, 0x92, 0x6b, 0xa1, 0x49, 0x18, 0x29, 0x50, 0x8e, 0x9b, 0x73, 0xa4, 0xe0, 0x48, 0xce, 0xb5,
	0x9a, 0x81, 0x0a, 0x94, 0x81, 0xe8, 0xfa, 0x95, 0xf1, 0xad, 0xf3, 0xd2, 0xdc, 0x4d, 0x80, 0x01,
	0x4f, 0x3e, 0x2b, 0xa8, 0x7e, 0x9b, 0xa9, 0x1e, 0x81, 0x01, 0x77, 0x4e, 0x51, 0x03, 0x62, 0x4f,
	0xc8, 0x89, 0xf2, 0x5e, 0x98, 0x06, 0x4f, 0x8c, 0x5d, 0xbb, 0x6d, 0x77, 0x6a, 0xc3, 0x3a, 0xc4,
	0x03, 0x39, 0x51, 0x77, 0x4c, 0xc3, 0x60, 0xec, 0x5c, 0xa3, 0x83, 0x1c, 0xd3, 0x6e, 0xa5, 0x5d,
	0xed, 0x1c, 0x5e, 0xb6, 0x49, 0x59, 0x47, 0xf2, 0x64, 0x36, 0xfb, 0xb5, 0xe5, 0xf7, 0xb1, 0x35,
	0xdc, 0xcf, 0x72, 0xb4, 0x73, 0x85, 0x5c, 0x5f, 0xc9, 0x35, 0x0c, 0x9e, 0x0a, 0x77, 0x95, 0x55,
	0xa3, 0x6c, 0x16, 0xf3, 0x87, 0x70, 0x4b, 0xcd, 0xd0, 0xd1, 0xff, 0x3d, 0xed, 0xd6, 0x4c, 0x07,
	0x5a, 0xde, 0xe1, 0x66, 0x13, 0xc5, 0x23, 0x06, 0x42, 0xc9, 0xad, 0x4a, 0x8d, 0x5d, 0x8f, 0xee,
	0xdf, 0x2f, 0x13, 0x6c, 0xaf, 0x12, 0x6c, 0xff, 0x24, 0xd8, 0xfe, 0x48, 0xb1, 0xb5, 0x4a, 0xb1,
	0xf5, 0x95, 0x62, 0xeb, 0xb9, 0x17, 0x08, 0x98, 0xbe, 0x8d, 0x88, 0xaf, 0x5e, 0x69, 0xee, 0xba,
	0x90, 0x1c, 0xe6, 0x2a, 0x9a, 0x15, 0x7f, 0x1a, 0xff, 0x5d, 0x1d, 0x16, 0x21, 0xd7, 0xa3, 0x3d,
	0x73, 0xeb, 0xde, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0xe2, 0x96, 0x47, 0xee, 0x01, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractOpInfos) > 0 {
		for iNdEx := len(m.ContractOpInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractOpInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ContractOpInfoLastId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ContractOpInfoLastId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TxInfos) > 0 {
		for iNdEx := len(m.TxInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.TxInfoLastId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TxInfoLastId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxInfoLastId != 0 {
		n += 1 + sovGenesis(uint64(m.TxInfoLastId))
	}
	if len(m.TxInfos) > 0 {
		for _, e := range m.TxInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ContractOpInfoLastId != 0 {
		n += 1 + sovGenesis(uint64(m.ContractOpInfoLastId))
	}
	if len(m.ContractOpInfos) > 0 {
		for _, e := range m.ContractOpInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxInfoLastId", wireType)
			}
			m.TxInfoLastId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxInfoLastId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxInfos = append(m.TxInfos, TxInfo{})
			if err := m.TxInfos[len(m.TxInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOpInfoLastId", wireType)
			}
			m.ContractOpInfoLastId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractOpInfoLastId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOpInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractOpInfos = append(m.ContractOpInfos, ContractOperationInfo{})
			if err := m.ContractOpInfos[len(m.ContractOpInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
