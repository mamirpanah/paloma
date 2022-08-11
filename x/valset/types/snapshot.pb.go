// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: valset/snapshot.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorState int32

const (
	ValidatorState_NONE   ValidatorState = 0
	ValidatorState_ACTIVE ValidatorState = 1
	ValidatorState_JAILED ValidatorState = 2
)

var ValidatorState_name = map[int32]string{
	0: "NONE",
	1: "ACTIVE",
	2: "JAILED",
}

var ValidatorState_value = map[string]int32{
	"NONE":   0,
	"ACTIVE": 1,
	"JAILED": 2,
}

func (x ValidatorState) String() string {
	return proto.EnumName(ValidatorState_name, int32(x))
}

func (ValidatorState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0adf39ac9845c7db, []int{0}
}

type Validator struct {
	ShareCount         github_com_cosmos_cosmos_sdk_types.Int        `protobuf:"bytes,1,opt,name=shareCount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shareCount"`
	State              ValidatorState                                `protobuf:"varint,2,opt,name=state,proto3,enum=volumefi.paloma.valset.ValidatorState" json:"state,omitempty"`
	ExternalChainInfos []*ExternalChainInfo                          `protobuf:"bytes,3,rep,name=externalChainInfos,proto3" json:"externalChainInfos,omitempty"`
	Address            github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,4,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"address,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_0adf39ac9845c7db, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetState() ValidatorState {
	if m != nil {
		return m.State
	}
	return ValidatorState_NONE
}

func (m *Validator) GetExternalChainInfos() []*ExternalChainInfo {
	if m != nil {
		return m.ExternalChainInfos
	}
	return nil
}

func (m *Validator) GetAddress() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

type ValidatorExternalAccounts struct {
	Address           github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"address,omitempty"`
	ExternalChainInfo []*ExternalChainInfo                          `protobuf:"bytes,2,rep,name=ExternalChainInfo,proto3" json:"ExternalChainInfo,omitempty"`
}

func (m *ValidatorExternalAccounts) Reset()         { *m = ValidatorExternalAccounts{} }
func (m *ValidatorExternalAccounts) String() string { return proto.CompactTextString(m) }
func (*ValidatorExternalAccounts) ProtoMessage()    {}
func (*ValidatorExternalAccounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_0adf39ac9845c7db, []int{1}
}
func (m *ValidatorExternalAccounts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorExternalAccounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorExternalAccounts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorExternalAccounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorExternalAccounts.Merge(m, src)
}
func (m *ValidatorExternalAccounts) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorExternalAccounts) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorExternalAccounts.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorExternalAccounts proto.InternalMessageInfo

func (m *ValidatorExternalAccounts) GetAddress() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ValidatorExternalAccounts) GetExternalChainInfo() []*ExternalChainInfo {
	if m != nil {
		return m.ExternalChainInfo
	}
	return nil
}

type Snapshot struct {
	Id          uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Height      int64                                  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Validators  []Validator                            `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators"`
	TotalShares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=totalShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalShares"`
	CreatedAt   time.Time                              `protobuf:"bytes,5,opt,name=createdAt,proto3,stdtime" json:"createdAt"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_0adf39ac9845c7db, []int{2}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Snapshot) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Snapshot) GetValidators() []Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *Snapshot) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

type ExternalChainInfo struct {
	ChainType        string `protobuf:"bytes,1,opt,name=chainType,proto3" json:"chainType,omitempty"`
	ChainReferenceID string `protobuf:"bytes,2,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	Address          string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Pubkey           []byte `protobuf:"bytes,4,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Balance          string `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *ExternalChainInfo) Reset()         { *m = ExternalChainInfo{} }
func (m *ExternalChainInfo) String() string { return proto.CompactTextString(m) }
func (*ExternalChainInfo) ProtoMessage()    {}
func (*ExternalChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0adf39ac9845c7db, []int{3}
}
func (m *ExternalChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExternalChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExternalChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExternalChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalChainInfo.Merge(m, src)
}
func (m *ExternalChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ExternalChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalChainInfo proto.InternalMessageInfo

func (m *ExternalChainInfo) GetChainType() string {
	if m != nil {
		return m.ChainType
	}
	return ""
}

func (m *ExternalChainInfo) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *ExternalChainInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ExternalChainInfo) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *ExternalChainInfo) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func init() {
	proto.RegisterEnum("volumefi.paloma.valset.ValidatorState", ValidatorState_name, ValidatorState_value)
	proto.RegisterType((*Validator)(nil), "volumefi.paloma.valset.Validator")
	proto.RegisterType((*ValidatorExternalAccounts)(nil), "volumefi.paloma.valset.ValidatorExternalAccounts")
	proto.RegisterType((*Snapshot)(nil), "volumefi.paloma.valset.Snapshot")
	proto.RegisterType((*ExternalChainInfo)(nil), "volumefi.paloma.valset.ExternalChainInfo")
}

func init() { proto.RegisterFile("valset/snapshot.proto", fileDescriptor_0adf39ac9845c7db) }

var fileDescriptor_0adf39ac9845c7db = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x3a, 0x69, 0xa8, 0xb7, 0xa8, 0x0a, 0x2b, 0xa8, 0xdc, 0x0a, 0x39, 0x21, 0x87, 0x2a,
	0x54, 0xd4, 0x56, 0xc3, 0x95, 0x4b, 0xd2, 0x06, 0x14, 0x40, 0x05, 0x6d, 0xab, 0x22, 0xb8, 0xa0,
	0xb5, 0xbd, 0xb1, 0xad, 0xda, 0x5e, 0xcb, 0xbb, 0x89, 0x9a, 0xbf, 0x28, 0x9f, 0xc1, 0x67, 0xf4,
	0xd6, 0x63, 0x8f, 0x88, 0x43, 0x41, 0xed, 0x5f, 0x70, 0x42, 0x5e, 0xdb, 0x49, 0xda, 0x04, 0x41,
	0xc5, 0x29, 0x3b, 0x9e, 0x79, 0x6f, 0x66, 0xde, 0xcc, 0x04, 0x3e, 0x1a, 0x91, 0x80, 0x53, 0x61,
	0xf2, 0x88, 0xc4, 0xdc, 0x63, 0xc2, 0x88, 0x13, 0x26, 0x18, 0x5a, 0x1b, 0xb1, 0x60, 0x18, 0xd2,
	0x81, 0x6f, 0xc4, 0x24, 0x60, 0x21, 0x31, 0xb2, 0xb0, 0x8d, 0x87, 0x2e, 0x73, 0x99, 0x0c, 0x31,
	0xd3, 0x57, 0x16, 0xbd, 0xb1, 0x6e, 0x33, 0x1e, 0x32, 0xfe, 0x39, 0x73, 0x64, 0x46, 0xee, 0xaa,
	0xbb, 0x8c, 0xb9, 0x01, 0x35, 0xa5, 0x65, 0x0d, 0x07, 0xa6, 0xf0, 0x43, 0xca, 0x05, 0x09, 0xe3,
	0x02, 0x7b, 0x3b, 0x80, 0x44, 0xe3, 0xdc, 0xa5, 0x67, 0x4c, 0xa6, 0x45, 0x38, 0x35, 0x47, 0x3b,
	0x16, 0x15, 0x64, 0xc7, 0xb4, 0x99, 0x1f, 0x65, 0xfe, 0xe6, 0x99, 0x02, 0xd5, 0x23, 0x12, 0xf8,
	0x0e, 0x11, 0x2c, 0x41, 0xfb, 0x10, 0x72, 0x8f, 0x24, 0x74, 0x97, 0x0d, 0x23, 0xa1, 0x81, 0x06,
	0x68, 0xdd, 0xef, 0x1a, 0xe7, 0x97, 0xf5, 0xd2, 0xf7, 0xcb, 0xfa, 0xa6, 0xeb, 0x0b, 0x6f, 0x68,
	0x19, 0x36, 0x0b, 0xf3, 0xf2, 0xf2, 0x9f, 0x6d, 0xee, 0x1c, 0x9b, 0x62, 0x1c, 0x53, 0x6e, 0xf4,
	0x23, 0x81, 0x67, 0x18, 0xd0, 0x0b, 0xb8, 0xc4, 0x05, 0x11, 0x54, 0x53, 0x1a, 0xa0, 0xb5, 0xda,
	0xde, 0x34, 0x16, 0x4b, 0x62, 0x4c, 0x2a, 0x38, 0x48, 0xa3, 0x71, 0x06, 0x42, 0x1f, 0x21, 0xa2,
	0x27, 0x82, 0x26, 0x11, 0x09, 0x76, 0x3d, 0xe2, 0x47, 0xfd, 0x68, 0xc0, 0xb8, 0x56, 0x6e, 0x94,
	0x5b, 0x2b, 0xed, 0xa7, 0x7f, 0xa2, 0xea, 0xdd, 0x46, 0xe0, 0x05, 0x24, 0xe8, 0x0d, 0xbc, 0x47,
	0x1c, 0x27, 0xa1, 0x9c, 0x6b, 0x15, 0xd9, 0xe5, 0xce, 0xaf, 0xcb, 0xfa, 0xf6, 0x3f, 0x74, 0x78,
	0x44, 0x82, 0x4e, 0x06, 0xc4, 0x05, 0x43, 0xf3, 0x0c, 0xc0, 0xf5, 0x49, 0x07, 0x45, 0xfe, 0x8e,
	0x6d, 0xa7, 0x12, 0xdc, 0x48, 0x05, 0xfe, 0x37, 0x15, 0xfa, 0x00, 0x1f, 0xcc, 0x35, 0xa8, 0x29,
	0x77, 0x55, 0x64, 0x9e, 0xa3, 0xf9, 0x45, 0x81, 0xcb, 0x07, 0xf9, 0xfe, 0xa2, 0x55, 0xa8, 0xf8,
	0x8e, 0xac, 0xb6, 0x82, 0x15, 0xdf, 0x41, 0x6b, 0xb0, 0xea, 0x51, 0xdf, 0xf5, 0x84, 0x9c, 0x63,
	0x19, 0xe7, 0x16, 0x7a, 0x05, 0xe1, 0xa8, 0xe8, 0xbb, 0x18, 0xcc, 0x93, 0xbf, 0xce, 0xb8, 0x5b,
	0x49, 0x37, 0x0a, 0xcf, 0x40, 0xd1, 0x7b, 0xb8, 0x22, 0x98, 0x20, 0xc1, 0x41, 0xba, 0x3a, 0xc5,
	0x48, 0xee, 0xba, 0x78, 0xb3, 0x14, 0xa8, 0x0b, 0x55, 0x3b, 0xa1, 0x44, 0x50, 0xa7, 0x23, 0xb4,
	0xa5, 0x06, 0x68, 0xad, 0xb4, 0x37, 0x8c, 0xec, 0x4c, 0x8c, 0xe2, 0x4c, 0x8c, 0xc3, 0xe2, 0x8e,
	0xba, 0xcb, 0x69, 0xae, 0xd3, 0x1f, 0x75, 0x80, 0xa7, 0xb0, 0xe6, 0x57, 0xb0, 0x40, 0x6d, 0xf4,
	0x18, 0xaa, 0x76, 0x6a, 0x1c, 0x8e, 0x63, 0x2a, 0x35, 0x52, 0xf1, 0xf4, 0x03, 0xda, 0x82, 0x35,
	0x69, 0x60, 0x3a, 0xa0, 0x09, 0x8d, 0x6c, 0xda, 0xdf, 0x93, 0xa2, 0xa9, 0x78, 0xee, 0x3b, 0xd2,
	0xa6, 0x9b, 0x51, 0x96, 0x21, 0x93, 0x31, 0xaf, 0xc1, 0x6a, 0x3c, 0xb4, 0x8e, 0xe9, 0x38, 0x93,
	0x02, 0xe7, 0x56, 0x8a, 0xb0, 0x48, 0x40, 0x22, 0x9b, 0xca, 0x9e, 0x54, 0x5c, 0x98, 0x5b, 0x6d,
	0xb8, 0x7a, 0xf3, 0x88, 0xd0, 0x32, 0xac, 0xec, 0xbf, 0xdb, 0xef, 0xd5, 0x4a, 0x08, 0xc2, 0x6a,
	0x67, 0xf7, 0xb0, 0x7f, 0xd4, 0xab, 0x81, 0xf4, 0xfd, 0xba, 0xd3, 0x7f, 0xdb, 0xdb, 0xab, 0x29,
	0xdd, 0x97, 0xe7, 0x57, 0x3a, 0xb8, 0xb8, 0xd2, 0xc1, 0xcf, 0x2b, 0x1d, 0x9c, 0x5e, 0xeb, 0xa5,
	0x8b, 0x6b, 0xbd, 0xf4, 0xed, 0x5a, 0x2f, 0x7d, 0x7a, 0x36, 0x23, 0x79, 0x36, 0x45, 0x59, 0x7c,
	0xfe, 0x36, 0x4f, 0xcc, 0xfc, 0x1f, 0x4f, 0x8a, 0x6f, 0x55, 0xa5, 0xa0, 0xcf, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0x80, 0x6d, 0xe6, 0x08, 0x05, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ExternalChainInfos) > 0 {
		for iNdEx := len(m.ExternalChainInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExternalChainInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSnapshot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.State != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.ShareCount.Size()
		i -= size
		if _, err := m.ShareCount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSnapshot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ValidatorExternalAccounts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorExternalAccounts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorExternalAccounts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExternalChainInfo) > 0 {
		for iNdEx := len(m.ExternalChainInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExternalChainInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSnapshot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSnapshot(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	{
		size := m.TotalShares.Size()
		i -= size
		if _, err := m.TotalShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSnapshot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSnapshot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExternalChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExternalChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainType) > 0 {
		i -= len(m.ChainType)
		copy(dAtA[i:], m.ChainType)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.ChainType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ShareCount.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	if m.State != 0 {
		n += 1 + sovSnapshot(uint64(m.State))
	}
	if len(m.ExternalChainInfos) > 0 {
		for _, e := range m.ExternalChainInfos {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func (m *ValidatorExternalAccounts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	if len(m.ExternalChainInfo) > 0 {
		for _, e := range m.ExternalChainInfo {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSnapshot(uint64(m.Id))
	}
	if m.Height != 0 {
		n += 1 + sovSnapshot(uint64(m.Height))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	l = m.TotalShares.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovSnapshot(uint64(l))
	return n
}

func (m *ExternalChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainType)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ShareCount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= ValidatorState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalChainInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalChainInfos = append(m.ExternalChainInfos, &ExternalChainInfo{})
			if err := m.ExternalChainInfos[len(m.ExternalChainInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *ValidatorExternalAccounts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: ValidatorExternalAccounts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorExternalAccounts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalChainInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalChainInfo = append(m.ExternalChainInfo, &ExternalChainInfo{})
			if err := m.ExternalChainInfo[len(m.ExternalChainInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *ExternalChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: ExternalChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSnapshot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSnapshot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSnapshot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSnapshot = fmt.Errorf("proto: unexpected end of group")
)
