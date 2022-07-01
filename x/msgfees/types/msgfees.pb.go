// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/msgfees/v1/msgfees.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the set of params for the msgfees module.
type Params struct {
	// constant used to calculate fees when gas fees shares denom with msg fee
	FloorGasPrice types.Coin `protobuf:"bytes,2,opt,name=floor_gas_price,json=floorGasPrice,proto3" json:"floor_gas_price" yaml:"floor_gas_price"`
	// total nhash per usd mil for converting usd to nhash
	NhashPerUsdMil uint64 `protobuf:"varint,3,opt,name=nhash_per_usd_mil,json=nhashPerUsdMil,proto3" json:"nhash_per_usd_mil,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6265859d114362, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetFloorGasPrice() types.Coin {
	if m != nil {
		return m.FloorGasPrice
	}
	return types.Coin{}
}

func (m *Params) GetNhashPerUsdMil() uint64 {
	if m != nil {
		return m.NhashPerUsdMil
	}
	return 0
}

// MsgFee is the core of what gets stored on the blockchain
// it consists of two parts
// 1. the msg type url, i.e. /cosmos.bank.v1beta1.MsgSend
// 2. minimum additional fees(can be of any denom)
type MsgFee struct {
	MsgTypeUrl string `protobuf:"bytes,1,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	// additional_fee can pay in any Coin( basically a Denom and Amount, Amount can be zero)
	AdditionalFee types.Coin `protobuf:"bytes,2,opt,name=additional_fee,json=additionalFee,proto3" json:"additional_fee" yaml:"additional_fee"`
}

func (m *MsgFee) Reset()         { *m = MsgFee{} }
func (m *MsgFee) String() string { return proto.CompactTextString(m) }
func (*MsgFee) ProtoMessage()    {}
func (*MsgFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6265859d114362, []int{1}
}
func (m *MsgFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFee.Merge(m, src)
}
func (m *MsgFee) XXX_Size() int {
	return m.Size()
}
func (m *MsgFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFee proto.InternalMessageInfo

func (m *MsgFee) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *MsgFee) GetAdditionalFee() types.Coin {
	if m != nil {
		return m.AdditionalFee
	}
	return types.Coin{}
}

// EventMsgFee final event property for msg fee on type
type EventMsgFee struct {
	MsgType   string `protobuf:"bytes,1,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Count     string `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	Total     string `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	Recipient string `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
}

func (m *EventMsgFee) Reset()         { *m = EventMsgFee{} }
func (m *EventMsgFee) String() string { return proto.CompactTextString(m) }
func (*EventMsgFee) ProtoMessage()    {}
func (*EventMsgFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6265859d114362, []int{2}
}
func (m *EventMsgFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMsgFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMsgFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMsgFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMsgFee.Merge(m, src)
}
func (m *EventMsgFee) XXX_Size() int {
	return m.Size()
}
func (m *EventMsgFee) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMsgFee.DiscardUnknown(m)
}

var xxx_messageInfo_EventMsgFee proto.InternalMessageInfo

func (m *EventMsgFee) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *EventMsgFee) GetCount() string {
	if m != nil {
		return m.Count
	}
	return ""
}

func (m *EventMsgFee) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func (m *EventMsgFee) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

// EventMsgFees event emitted with summary of msg fees
type EventMsgFees struct {
	MsgFees []EventMsgFee `protobuf:"bytes,1,rep,name=msg_fees,json=msgFees,proto3" json:"msg_fees"`
}

func (m *EventMsgFees) Reset()         { *m = EventMsgFees{} }
func (m *EventMsgFees) String() string { return proto.CompactTextString(m) }
func (*EventMsgFees) ProtoMessage()    {}
func (*EventMsgFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c6265859d114362, []int{3}
}
func (m *EventMsgFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMsgFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMsgFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMsgFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMsgFees.Merge(m, src)
}
func (m *EventMsgFees) XXX_Size() int {
	return m.Size()
}
func (m *EventMsgFees) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMsgFees.DiscardUnknown(m)
}

var xxx_messageInfo_EventMsgFees proto.InternalMessageInfo

func (m *EventMsgFees) GetMsgFees() []EventMsgFee {
	if m != nil {
		return m.MsgFees
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "provenance.msgfees.v1.Params")
	proto.RegisterType((*MsgFee)(nil), "provenance.msgfees.v1.MsgFee")
	proto.RegisterType((*EventMsgFee)(nil), "provenance.msgfees.v1.EventMsgFee")
	proto.RegisterType((*EventMsgFees)(nil), "provenance.msgfees.v1.EventMsgFees")
}

func init() {
	proto.RegisterFile("provenance/msgfees/v1/msgfees.proto", fileDescriptor_0c6265859d114362)
}

var fileDescriptor_0c6265859d114362 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xd1, 0x10, 0x9a, 0x4b, 0x5b, 0x84, 0xd5, 0x22, 0xb7, 0x02, 0xc7, 0x32, 0x4b, 0x18,
	0x38, 0x2b, 0xed, 0xd6, 0x31, 0x15, 0x65, 0xaa, 0x14, 0x19, 0xba, 0xb0, 0x9c, 0x2e, 0xce, 0x8b,
	0x73, 0x92, 0x7d, 0x67, 0xdd, 0x5d, 0x2c, 0xf2, 0x37, 0xb0, 0x30, 0x32, 0xe6, 0xcf, 0xe9, 0xd8,
	0x91, 0xa9, 0x42, 0xc9, 0xc2, 0xcc, 0x5f, 0x80, 0xfc, 0x23, 0x38, 0xa0, 0x0c, 0xdd, 0xee, 0xbd,
	0xef, 0x7b, 0xdf, 0xfb, 0x3e, 0xfb, 0xe1, 0x37, 0x99, 0x92, 0x39, 0x08, 0x26, 0x22, 0x08, 0x52,
	0x1d, 0x4f, 0x01, 0x74, 0x90, 0x0f, 0x36, 0x4f, 0x92, 0x29, 0x69, 0xa4, 0x7d, 0xd2, 0x90, 0xc8,
	0x06, 0xc9, 0x07, 0x67, 0xc7, 0xb1, 0x8c, 0x65, 0xc9, 0x08, 0x8a, 0x57, 0x45, 0x3e, 0x73, 0x23,
	0xa9, 0x53, 0xa9, 0x83, 0x31, 0xd3, 0x10, 0xe4, 0x83, 0x31, 0x18, 0x36, 0x08, 0x22, 0xc9, 0x45,
	0x85, 0xfb, 0x4b, 0x84, 0xdb, 0x23, 0xa6, 0x58, 0xaa, 0x6d, 0x86, 0x9f, 0x4f, 0x13, 0x29, 0x15,
	0x8d, 0x99, 0xa6, 0x99, 0xe2, 0x11, 0x38, 0x4f, 0x3c, 0xd4, 0xef, 0x9e, 0x9f, 0x92, 0x4a, 0x84,
	0x14, 0x22, 0xa4, 0x16, 0x21, 0x57, 0x92, 0x8b, 0xa1, 0x7b, 0xf7, 0xd0, 0xb3, 0x7e, 0x3f, 0xf4,
	0x5e, 0x2e, 0x58, 0x9a, 0x5c, 0xfa, 0xff, 0xcd, 0xfb, 0xe1, 0x61, 0xd9, 0xf9, 0xc0, 0xf4, 0xa8,
	0xa8, 0xed, 0xb7, 0xf8, 0x85, 0x98, 0x31, 0x3d, 0xa3, 0x19, 0x28, 0x3a, 0xd7, 0x13, 0x9a, 0xf2,
	0xc4, 0xd9, 0xf3, 0x50, 0xbf, 0x15, 0x1e, 0x95, 0xc0, 0x08, 0xd4, 0xad, 0x9e, 0xdc, 0xf0, 0xe4,
	0x72, 0xff, 0xfb, 0xb2, 0x67, 0xfd, 0x5a, 0xf6, 0x2c, 0xff, 0x2b, 0xc2, 0xed, 0x1b, 0x1d, 0x5f,
	0x03, 0xd8, 0x1e, 0x3e, 0x48, 0x75, 0x4c, 0xcd, 0x22, 0x03, 0x3a, 0x57, 0x89, 0x83, 0x3c, 0xd4,
	0xef, 0x84, 0x38, 0xd5, 0xf1, 0xa7, 0x45, 0x06, 0xb7, 0x2a, 0xb1, 0x29, 0x3e, 0x62, 0x93, 0x09,
	0x37, 0x5c, 0x0a, 0x96, 0xd0, 0x29, 0x3c, 0x22, 0xc3, 0xeb, 0x3a, 0xc3, 0x49, 0x95, 0xe1, 0xdf,
	0x71, 0x3f, 0x3c, 0x6c, 0x1a, 0xd7, 0x00, 0xbe, 0xc2, 0xdd, 0xf7, 0x39, 0x08, 0x53, 0x3b, 0x3a,
	0xc5, 0xfb, 0x1b, 0x47, 0xb5, 0x9b, 0x67, 0xb5, 0x1b, 0xfb, 0x18, 0x3f, 0x8d, 0xe4, 0x5c, 0x98,
	0xd2, 0x41, 0x27, 0xac, 0x8a, 0xa2, 0x6b, 0xa4, 0x61, 0x55, 0xec, 0x4e, 0x58, 0x15, 0xf6, 0x2b,
	0xdc, 0x51, 0x10, 0xf1, 0x8c, 0x83, 0x30, 0x4e, 0xab, 0x44, 0x9a, 0x86, 0xff, 0x11, 0x1f, 0x6c,
	0xed, 0xd4, 0xf6, 0x55, 0xb5, 0xb4, 0xf8, 0xf3, 0x0e, 0xf2, 0xf6, 0xfa, 0xdd, 0x73, 0x9f, 0xec,
	0x3c, 0x0a, 0xb2, 0x35, 0x36, 0x6c, 0x15, 0x39, 0x4b, 0x7b, 0x85, 0xc8, 0x90, 0xdf, 0xad, 0x5c,
	0x74, 0xbf, 0x72, 0xd1, 0xcf, 0x95, 0x8b, 0xbe, 0xad, 0x5d, 0xeb, 0x7e, 0xed, 0x5a, 0x3f, 0xd6,
	0xae, 0x85, 0x1d, 0x2e, 0x77, 0xcb, 0x8d, 0xd0, 0xe7, 0x8b, 0x98, 0x9b, 0xd9, 0x7c, 0x4c, 0x22,
	0x99, 0x06, 0x0d, 0xe7, 0x1d, 0x97, 0x5b, 0x55, 0xf0, 0xe5, 0xef, 0xf1, 0x16, 0xdf, 0x45, 0x8f,
	0xdb, 0xe5, 0xad, 0x5d, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x25, 0x33, 0x27, 0xdf, 0x02,
	0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NhashPerUsdMil != 0 {
		i = encodeVarintMsgfees(dAtA, i, uint64(m.NhashPerUsdMil))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.FloorGasPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgfees(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *MsgFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AdditionalFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgfees(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintMsgfees(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventMsgFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMsgFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMsgFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintMsgfees(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Total) > 0 {
		i -= len(m.Total)
		copy(dAtA[i:], m.Total)
		i = encodeVarintMsgfees(dAtA, i, uint64(len(m.Total)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Count) > 0 {
		i -= len(m.Count)
		copy(dAtA[i:], m.Count)
		i = encodeVarintMsgfees(dAtA, i, uint64(len(m.Count)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MsgType) > 0 {
		i -= len(m.MsgType)
		copy(dAtA[i:], m.MsgType)
		i = encodeVarintMsgfees(dAtA, i, uint64(len(m.MsgType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventMsgFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMsgFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMsgFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgFees) > 0 {
		for iNdEx := len(m.MsgFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MsgFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgfees(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgfees(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgfees(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.FloorGasPrice.Size()
	n += 1 + l + sovMsgfees(uint64(l))
	if m.NhashPerUsdMil != 0 {
		n += 1 + sovMsgfees(uint64(m.NhashPerUsdMil))
	}
	return n
}

func (m *MsgFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovMsgfees(uint64(l))
	}
	l = m.AdditionalFee.Size()
	n += 1 + l + sovMsgfees(uint64(l))
	return n
}

func (m *EventMsgFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgType)
	if l > 0 {
		n += 1 + l + sovMsgfees(uint64(l))
	}
	l = len(m.Count)
	if l > 0 {
		n += 1 + l + sovMsgfees(uint64(l))
	}
	l = len(m.Total)
	if l > 0 {
		n += 1 + l + sovMsgfees(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovMsgfees(uint64(l))
	}
	return n
}

func (m *EventMsgFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MsgFees) > 0 {
		for _, e := range m.MsgFees {
			l = e.Size()
			n += 1 + l + sovMsgfees(uint64(l))
		}
	}
	return n
}

func sovMsgfees(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgfees(x uint64) (n int) {
	return sovMsgfees(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgfees
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloorGasPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FloorGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NhashPerUsdMil", wireType)
			}
			m.NhashPerUsdMil = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NhashPerUsdMil |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgfees(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgfees
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
func (m *MsgFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgfees
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
			return fmt.Errorf("proto: MsgFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgfees(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgfees
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
func (m *EventMsgFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgfees
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
			return fmt.Errorf("proto: EventMsgFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMsgFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Count = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Total = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgfees(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgfees
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
func (m *EventMsgFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgfees
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
			return fmt.Errorf("proto: EventMsgFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMsgFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgfees
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
				return ErrInvalidLengthMsgfees
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgfees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgFees = append(m.MsgFees, EventMsgFee{})
			if err := m.MsgFees[len(m.MsgFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgfees(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgfees
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
func skipMsgfees(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgfees
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
					return 0, ErrIntOverflowMsgfees
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
					return 0, ErrIntOverflowMsgfees
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
				return 0, ErrInvalidLengthMsgfees
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgfees
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgfees
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgfees        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgfees          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgfees = fmt.Errorf("proto: unexpected end of group")
)
