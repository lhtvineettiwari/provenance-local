// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/msgfees/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgAssessCustomMsgFeeRequest defines an sdk.Msg type
type MsgAssessCustomMsgFeeRequest struct {
	Name      string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount    types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount" yaml:"amount"`
	Recipient string     `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// empty, whole amount goes to fee module
	From string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
}

func (m *MsgAssessCustomMsgFeeRequest) Reset()      { *m = MsgAssessCustomMsgFeeRequest{} }
func (*MsgAssessCustomMsgFeeRequest) ProtoMessage() {}
func (*MsgAssessCustomMsgFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6bb65eaf858b5f, []int{0}
}
func (m *MsgAssessCustomMsgFeeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAssessCustomMsgFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAssessCustomMsgFeeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAssessCustomMsgFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAssessCustomMsgFeeRequest.Merge(m, src)
}
func (m *MsgAssessCustomMsgFeeRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAssessCustomMsgFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAssessCustomMsgFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAssessCustomMsgFeeRequest proto.InternalMessageInfo

// MsgAssessCustomMsgFeeResponse defines the Msg/AssessCustomMsgFeee response type.
type MsgAssessCustomMsgFeeResponse struct {
}

func (m *MsgAssessCustomMsgFeeResponse) Reset()         { *m = MsgAssessCustomMsgFeeResponse{} }
func (m *MsgAssessCustomMsgFeeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAssessCustomMsgFeeResponse) ProtoMessage()    {}
func (*MsgAssessCustomMsgFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6bb65eaf858b5f, []int{1}
}
func (m *MsgAssessCustomMsgFeeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAssessCustomMsgFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAssessCustomMsgFeeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAssessCustomMsgFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAssessCustomMsgFeeResponse.Merge(m, src)
}
func (m *MsgAssessCustomMsgFeeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAssessCustomMsgFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAssessCustomMsgFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAssessCustomMsgFeeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAssessCustomMsgFeeRequest)(nil), "provenance.msgfees.v1.MsgAssessCustomMsgFeeRequest")
	proto.RegisterType((*MsgAssessCustomMsgFeeResponse)(nil), "provenance.msgfees.v1.MsgAssessCustomMsgFeeResponse")
}

func init() { proto.RegisterFile("provenance/msgfees/v1/tx.proto", fileDescriptor_4c6bb65eaf858b5f) }

var fileDescriptor_4c6bb65eaf858b5f = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x8f, 0xd3, 0x30,
	0x14, 0xc6, 0x63, 0xee, 0x74, 0xd2, 0x19, 0x58, 0x2c, 0x4e, 0x0a, 0xd5, 0xe1, 0x9c, 0x32, 0xdd,
	0x82, 0xad, 0x5c, 0x99, 0xba, 0xd1, 0x4a, 0x88, 0xa5, 0x12, 0xea, 0xc8, 0x96, 0x44, 0xaf, 0xc1,
	0x12, 0xf6, 0x0b, 0x79, 0x4e, 0xd4, 0x4e, 0x6c, 0x88, 0x91, 0x91, 0xb1, 0x7f, 0x08, 0x7f, 0x40,
	0xc7, 0x8e, 0x4c, 0x15, 0x6a, 0x17, 0x66, 0xfe, 0x02, 0x94, 0xa4, 0xa2, 0x48, 0x14, 0x24, 0xb6,
	0x2f, 0xf9, 0xbe, 0xcf, 0xbf, 0xa7, 0x67, 0x73, 0x59, 0x56, 0xd8, 0x80, 0x4b, 0x5d, 0x0e, 0xda,
	0x52, 0x31, 0x07, 0x20, 0xdd, 0x24, 0xda, 0x2f, 0x54, 0x59, 0xa1, 0x47, 0x71, 0x75, 0xf4, 0xd5,
	0xc1, 0x57, 0x4d, 0x32, 0x78, 0x54, 0x60, 0x81, 0x5d, 0x42, 0xb7, 0xaa, 0x0f, 0x0f, 0x64, 0x8e,
	0x64, 0x91, 0x74, 0x96, 0x12, 0xe8, 0x26, 0xc9, 0xc0, 0xa7, 0x89, 0xce, 0xd1, 0xb8, 0xde, 0x8f,
	0xbf, 0x30, 0x7e, 0x3d, 0xa5, 0xe2, 0x39, 0x11, 0x10, 0x4d, 0x6a, 0xf2, 0x68, 0xa7, 0x54, 0xbc,
	0x00, 0x98, 0xc1, 0xbb, 0x1a, 0xc8, 0x0b, 0xc1, 0xcf, 0x5d, 0x6a, 0x21, 0x64, 0x37, 0xec, 0xf6,
	0x72, 0xd6, 0x69, 0xf1, 0x92, 0x5f, 0xa4, 0x16, 0x6b, 0xe7, 0xc3, 0x7b, 0x37, 0xec, 0xf6, 0xfe,
	0xdd, 0x63, 0xd5, 0x53, 0x54, 0x4b, 0x51, 0x07, 0x8a, 0x9a, 0xa0, 0x71, 0xe3, 0xab, 0xf5, 0x36,
	0x0a, 0x7e, 0x6c, 0xa3, 0x87, 0xcb, 0xd4, 0xbe, 0x1d, 0xc5, 0x7d, 0x2d, 0x9e, 0x1d, 0xfa, 0xe2,
	0x9a, 0x5f, 0x56, 0x90, 0x9b, 0xd2, 0x80, 0xf3, 0xe1, 0x59, 0x87, 0x38, 0xfe, 0x68, 0xd9, 0xf3,
	0x0a, 0x6d, 0x78, 0xde, 0xb3, 0x5b, 0x3d, 0x7a, 0xf0, 0x71, 0x15, 0x05, 0x9f, 0x57, 0x51, 0xf0,
	0x7d, 0x15, 0x05, 0x71, 0xc4, 0x9f, 0xfc, 0x65, 0x7a, 0x2a, 0xd1, 0x11, 0xdc, 0x7d, 0x60, 0xfc,
	0x6c, 0x4a, 0x85, 0x78, 0xcf, 0xc5, 0x9f, 0x29, 0x31, 0x54, 0x27, 0x77, 0xa9, 0xfe, 0xb5, 0x91,
	0xc1, 0xb3, 0xff, 0x2b, 0xf5, 0x83, 0x8c, 0xcd, 0x7a, 0x27, 0xd9, 0x66, 0x27, 0xd9, 0xb7, 0x9d,
	0x64, 0x9f, 0xf6, 0x32, 0xd8, 0xec, 0x65, 0xf0, 0x75, 0x2f, 0x03, 0x1e, 0x1a, 0x3c, 0x7d, 0xe2,
	0x2b, 0xf6, 0x7a, 0x58, 0x18, 0xff, 0xa6, 0xce, 0x54, 0x8e, 0x56, 0x1f, 0x33, 0x4f, 0x0d, 0xfe,
	0xf6, 0xa5, 0x17, 0xbf, 0x9e, 0x89, 0x5f, 0x96, 0x40, 0xd9, 0x45, 0x77, 0xb5, 0xc3, 0x9f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x8c, 0x08, 0x4f, 0x89, 0x49, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// AssessCustomMsgFee endpoint executes the additional fee charges.
	// This will only emit the event and not persist it to the keeper.  Fees are handled with the custom msg fee handlers
	// Use Case: smart contracts will be able to charge additional fees and direct partial funds to specified recipient
	// for executing contracts
	AssessCustomMsgFee(ctx context.Context, in *MsgAssessCustomMsgFeeRequest, opts ...grpc.CallOption) (*MsgAssessCustomMsgFeeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AssessCustomMsgFee(ctx context.Context, in *MsgAssessCustomMsgFeeRequest, opts ...grpc.CallOption) (*MsgAssessCustomMsgFeeResponse, error) {
	out := new(MsgAssessCustomMsgFeeResponse)
	err := c.cc.Invoke(ctx, "/provenance.msgfees.v1.Msg/AssessCustomMsgFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AssessCustomMsgFee endpoint executes the additional fee charges.
	// This will only emit the event and not persist it to the keeper.  Fees are handled with the custom msg fee handlers
	// Use Case: smart contracts will be able to charge additional fees and direct partial funds to specified recipient
	// for executing contracts
	AssessCustomMsgFee(context.Context, *MsgAssessCustomMsgFeeRequest) (*MsgAssessCustomMsgFeeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AssessCustomMsgFee(ctx context.Context, req *MsgAssessCustomMsgFeeRequest) (*MsgAssessCustomMsgFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssessCustomMsgFee not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AssessCustomMsgFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAssessCustomMsgFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AssessCustomMsgFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.msgfees.v1.Msg/AssessCustomMsgFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AssessCustomMsgFee(ctx, req.(*MsgAssessCustomMsgFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.msgfees.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssessCustomMsgFee",
			Handler:    _Msg_AssessCustomMsgFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provenance/msgfees/v1/tx.proto",
}

func (m *MsgAssessCustomMsgFeeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAssessCustomMsgFeeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAssessCustomMsgFeeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAssessCustomMsgFeeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAssessCustomMsgFeeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAssessCustomMsgFeeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAssessCustomMsgFeeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAssessCustomMsgFeeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAssessCustomMsgFeeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAssessCustomMsgFeeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAssessCustomMsgFeeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAssessCustomMsgFeeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAssessCustomMsgFeeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAssessCustomMsgFeeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
