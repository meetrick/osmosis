// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/superfluid.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type SuperfluidAssetType int32

const (
	// TODO: what superfluid asset types will be here?
	SuperfluidAssetTypeNative  SuperfluidAssetType = 0
	SuperfluidAssetTypeLPShare SuperfluidAssetType = 1
)

var SuperfluidAssetType_name = map[int32]string{
	0: "SuperfluidAssetTypeNative",
	1: "SuperfluidAssetTypeLPShare",
}

var SuperfluidAssetType_value = map[string]int32{
	"SuperfluidAssetTypeNative":  0,
	"SuperfluidAssetTypeLPShare": 1,
}

func (x SuperfluidAssetType) String() string {
	return proto.EnumName(SuperfluidAssetType_name, int32(x))
}

func (SuperfluidAssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{0}
}

// SuperfluidAsset stores the pair of superfluid asset type and denom pair
type SuperfluidAsset struct {
	AssetType SuperfluidAssetType `protobuf:"varint,1,opt,name=asset_type,json=assetType,proto3,enum=osmosis.superfluid.SuperfluidAssetType" json:"asset_type,omitempty"`
	Denom     string              `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *SuperfluidAsset) Reset()      { *m = SuperfluidAsset{} }
func (*SuperfluidAsset) ProtoMessage() {}
func (*SuperfluidAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{0}
}
func (m *SuperfluidAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidAsset.Merge(m, src)
}
func (m *SuperfluidAsset) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidAsset.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidAsset proto.InternalMessageInfo

// SuperfluidAssetInfo stores the information of superfluid asset - real time changes
type SuperfluidAssetInfo struct {
	SuperfluidAssetId          uint64                                 `protobuf:"varint,1,opt,name=superfluid_asset_id,json=superfluidAssetId,proto3" json:"superfluid_asset_id,omitempty"`
	TotalStakedAmount          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=total_staked_amount,json=totalStakedAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_staked_amount" yaml:"superfluid_staked_amount"`
	RiskAdjustedOsmoEquivalent github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=risk_adjusted_osmo_equivalent,json=riskAdjustedOsmoEquivalent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"risk_adjusted_osmo_equivalent" yaml:"risk_adjusted_osmo_equivalent"`
}

func (m *SuperfluidAssetInfo) Reset()         { *m = SuperfluidAssetInfo{} }
func (m *SuperfluidAssetInfo) String() string { return proto.CompactTextString(m) }
func (*SuperfluidAssetInfo) ProtoMessage()    {}
func (*SuperfluidAssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{1}
}
func (m *SuperfluidAssetInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidAssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidAssetInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidAssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidAssetInfo.Merge(m, src)
}
func (m *SuperfluidAssetInfo) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidAssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidAssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidAssetInfo proto.InternalMessageInfo

func (m *SuperfluidAssetInfo) GetSuperfluidAssetId() uint64 {
	if m != nil {
		return m.SuperfluidAssetId
	}
	return 0
}

func init() {
	proto.RegisterEnum("osmosis.superfluid.SuperfluidAssetType", SuperfluidAssetType_name, SuperfluidAssetType_value)
	proto.RegisterType((*SuperfluidAsset)(nil), "osmosis.superfluid.SuperfluidAsset")
	proto.RegisterType((*SuperfluidAssetInfo)(nil), "osmosis.superfluid.SuperfluidAssetInfo")
}

func init() {
	proto.RegisterFile("osmosis/superfluid/superfluid.proto", fileDescriptor_79d3c29d82dbb734)
}

var fileDescriptor_79d3c29d82dbb734 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0xb5, 0x0a, 0x1d, 0x44, 0xdb, 0x6c, 0x0f, 0x6b, 0xa0, 0x93, 0x12, 0x45, 0x8b,
	0xd0, 0x0c, 0xad, 0xb7, 0xde, 0xb6, 0xa0, 0x50, 0x28, 0xfe, 0xc8, 0x8a, 0x87, 0x5e, 0xc2, 0x64,
	0x67, 0x36, 0x1d, 0x37, 0xc9, 0xc4, 0xcc, 0xcc, 0xe2, 0x1e, 0x3c, 0x78, 0xeb, 0xd1, 0xa3, 0xde,
	0x16, 0xfc, 0x67, 0x7a, 0xec, 0x51, 0x3c, 0x2c, 0xb2, 0x8b, 0xe0, 0xd9, 0xbf, 0x40, 0x32, 0xc9,
	0xb6, 0x6b, 0xba, 0x08, 0x9e, 0xf2, 0xde, 0xfb, 0xce, 0x7c, 0xdf, 0xe7, 0x4d, 0x66, 0xe0, 0x7d,
	0x21, 0x53, 0x21, 0xb9, 0xc4, 0x52, 0xe7, 0xac, 0xe8, 0x27, 0x9a, 0xd3, 0x85, 0xd0, 0xcf, 0x0b,
	0xa1, 0x84, 0x6d, 0xd7, 0x8b, 0xfc, 0x2b, 0xc5, 0xd9, 0x8c, 0x45, 0x2c, 0x8c, 0x8c, 0xcb, 0xa8,
	0x5a, 0xe9, 0xa0, 0x58, 0x88, 0x38, 0x61, 0xd8, 0x64, 0x91, 0xee, 0x63, 0xaa, 0x0b, 0xa2, 0xb8,
	0xc8, 0x6a, 0xdd, 0x6d, 0xea, 0x8a, 0xa7, 0x4c, 0x2a, 0x92, 0xe6, 0x73, 0x83, 0x9e, 0xe9, 0x85,
	0x23, 0x22, 0x19, 0x1e, 0xee, 0x45, 0x4c, 0x91, 0x3d, 0xdc, 0x13, 0xbc, 0x36, 0xf0, 0x3e, 0xc0,
	0xbb, 0xdd, 0x4b, 0x88, 0x8e, 0x94, 0x4c, 0xd9, 0xcf, 0x20, 0x24, 0x65, 0x10, 0xaa, 0x51, 0xce,
	0xda, 0x60, 0x1b, 0xec, 0xdc, 0xd9, 0x7f, 0xe4, 0x5f, 0x47, 0xf6, 0x1b, 0x1b, 0x5f, 0x8f, 0x72,
	0x16, 0xac, 0x91, 0x79, 0x68, 0x6f, 0xc2, 0x9b, 0x94, 0x65, 0x22, 0x6d, 0xaf, 0x6c, 0x83, 0x9d,
	0xb5, 0xa0, 0x4a, 0x0e, 0x6e, 0x9f, 0x8d, 0x5d, 0xeb, 0xf3, 0xd8, 0xb5, 0x7e, 0x8d, 0x5d, 0xe0,
	0xfd, 0x5c, 0x81, 0xad, 0x86, 0xcd, 0x51, 0xd6, 0x17, 0xb6, 0x0f, 0x5b, 0x57, 0x8d, 0xc2, 0x0a,
	0x87, 0x53, 0x03, 0xb3, 0x1a, 0x6c, 0xc8, 0xc6, 0x0e, 0x6a, 0x7f, 0x04, 0xb0, 0xa5, 0x84, 0x22,
	0x49, 0x28, 0x15, 0x19, 0x30, 0x1a, 0x92, 0x54, 0xe8, 0x4c, 0x55, 0xad, 0x0f, 0x5f, 0x9d, 0x4f,
	0x5c, 0xeb, 0xfb, 0xc4, 0x7d, 0x18, 0x73, 0x75, 0xaa, 0x23, 0xbf, 0x27, 0x52, 0x5c, 0x9f, 0x4b,
	0xf5, 0xd9, 0x95, 0x74, 0x80, 0xcb, 0x71, 0xa5, 0x7f, 0x94, 0xa9, 0xdf, 0x13, 0xd7, 0x1d, 0x91,
	0x34, 0x39, 0xf0, 0x16, 0x18, 0xfe, 0xf2, 0xf5, 0x82, 0x0d, 0xd3, 0xad, 0x6b, 0x8a, 0x1d, 0x53,
	0xb3, 0xbf, 0x00, 0xb8, 0x55, 0x70, 0x39, 0x08, 0x09, 0x7d, 0xab, 0xa5, 0x62, 0x34, 0x2c, 0xbd,
	0x43, 0xf6, 0x4e, 0xf3, 0x21, 0x49, 0x58, 0xa6, 0xda, 0x37, 0x0c, 0xcd, 0x9b, 0xff, 0xa6, 0x79,
	0x50, 0xd1, 0xfc, 0xd3, 0xdc, 0x0b, 0x9c, 0x52, 0xef, 0xd4, 0xf2, 0x0b, 0x99, 0x8a, 0xa7, 0x97,
	0xe2, 0xe3, 0x93, 0x6b, 0xc7, 0x6c, 0x7e, 0xd1, 0x16, 0xbc, 0xb7, 0xa4, 0xfc, 0x9c, 0x28, 0x3e,
	0x64, 0xeb, 0x96, 0x8d, 0xa0, 0xb3, 0x44, 0x3e, 0x7e, 0xd9, 0x3d, 0x25, 0x05, 0x5b, 0x07, 0xce,
	0xea, 0xd9, 0x57, 0x64, 0x1d, 0x1e, 0x9f, 0x4f, 0x11, 0xb8, 0x98, 0x22, 0xf0, 0x63, 0x8a, 0xc0,
	0xa7, 0x19, 0xb2, 0x2e, 0x66, 0xc8, 0xfa, 0x36, 0x43, 0xd6, 0xc9, 0xfe, 0xc2, 0x84, 0xf5, 0xfd,
	0xd9, 0x4d, 0x48, 0x24, 0xe7, 0x09, 0x7e, 0xbf, 0xf8, 0x4c, 0xcc, 0xc4, 0xd1, 0x2d, 0x73, 0x2f,
	0x9f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x57, 0x05, 0xc3, 0x49, 0x03, 0x00, 0x00,
}

func (this *SuperfluidAsset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SuperfluidAsset)
	if !ok {
		that2, ok := that.(SuperfluidAsset)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AssetType != that1.AssetType {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	return true
}
func (m *SuperfluidAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.AssetType != 0 {
		i = encodeVarintSuperfluid(dAtA, i, uint64(m.AssetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SuperfluidAssetInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidAssetInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidAssetInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RiskAdjustedOsmoEquivalent.Size()
		i -= size
		if _, err := m.RiskAdjustedOsmoEquivalent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSuperfluid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TotalStakedAmount.Size()
		i -= size
		if _, err := m.TotalStakedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSuperfluid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SuperfluidAssetId != 0 {
		i = encodeVarintSuperfluid(dAtA, i, uint64(m.SuperfluidAssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSuperfluid(dAtA []byte, offset int, v uint64) int {
	offset -= sovSuperfluid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SuperfluidAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetType != 0 {
		n += 1 + sovSuperfluid(uint64(m.AssetType))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	return n
}

func (m *SuperfluidAssetInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SuperfluidAssetId != 0 {
		n += 1 + sovSuperfluid(uint64(m.SuperfluidAssetId))
	}
	l = m.TotalStakedAmount.Size()
	n += 1 + l + sovSuperfluid(uint64(l))
	l = m.RiskAdjustedOsmoEquivalent.Size()
	n += 1 + l + sovSuperfluid(uint64(l))
	return n
}

func sovSuperfluid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSuperfluid(x uint64) (n int) {
	return sovSuperfluid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SuperfluidAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetType", wireType)
			}
			m.AssetType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetType |= SuperfluidAssetType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func (m *SuperfluidAssetInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidAssetInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidAssetInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidAssetId", wireType)
			}
			m.SuperfluidAssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuperfluidAssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalStakedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RiskAdjustedOsmoEquivalent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RiskAdjustedOsmoEquivalent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func skipSuperfluid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
				return 0, ErrInvalidLengthSuperfluid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSuperfluid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSuperfluid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSuperfluid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSuperfluid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSuperfluid = fmt.Errorf("proto: unexpected end of group")
)
