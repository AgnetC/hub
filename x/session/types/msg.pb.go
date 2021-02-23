// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	types "github.com/sentinel-official/hub/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgUpsertRequest struct {
	From      string          `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Id        uint64          `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Address   string          `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Duration  time.Duration   `protobuf:"bytes,4,opt,name=duration,proto3,stdduration" json:"duration"`
	Bandwidth types.Bandwidth `protobuf:"bytes,5,opt,name=bandwidth,proto3" json:"bandwidth"`
}

func (m *MsgUpsertRequest) Reset()         { *m = MsgUpsertRequest{} }
func (m *MsgUpsertRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpsertRequest) ProtoMessage()    {}
func (*MsgUpsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfac64e08a3e319, []int{0}
}
func (m *MsgUpsertRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpsertRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpsertRequest.Merge(m, src)
}
func (m *MsgUpsertRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpsertRequest proto.InternalMessageInfo

type MsgUpsertResponse struct {
}

func (m *MsgUpsertResponse) Reset()         { *m = MsgUpsertResponse{} }
func (m *MsgUpsertResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpsertResponse) ProtoMessage()    {}
func (*MsgUpsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfac64e08a3e319, []int{1}
}
func (m *MsgUpsertResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpsertResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpsertResponse.Merge(m, src)
}
func (m *MsgUpsertResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpsertResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpsertRequest)(nil), "sentinel.session.v1.MsgUpsertRequest")
	proto.RegisterType((*MsgUpsertResponse)(nil), "sentinel.session.v1.MsgUpsertResponse")
}

func init() { proto.RegisterFile("sentinel/session/v1/msg.proto", fileDescriptor_1cfac64e08a3e319) }

var fileDescriptor_1cfac64e08a3e319 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x8f, 0xd4, 0x30,
	0x10, 0x8d, 0x97, 0x00, 0xb7, 0x06, 0x21, 0xf0, 0x51, 0x84, 0x15, 0x38, 0xcb, 0x49, 0xa0, 0x6d,
	0xb0, 0xc9, 0xd1, 0xd1, 0x80, 0x22, 0x2a, 0xa4, 0x6b, 0x82, 0x68, 0xae, 0xcb, 0x87, 0xe3, 0x58,
	0xda, 0xc4, 0x21, 0xe3, 0x04, 0xf8, 0x07, 0x94, 0x94, 0x57, 0xee, 0xcf, 0xb9, 0xf2, 0x4a, 0x2a,
	0x3e, 0x76, 0x1b, 0x7e, 0x06, 0x8a, 0xb3, 0x09, 0x08, 0x90, 0xae, 0x9b, 0xf1, 0x7b, 0xf3, 0x3c,
	0xef, 0xd9, 0xf8, 0x01, 0x88, 0xca, 0xa8, 0x4a, 0xac, 0x39, 0x08, 0x00, 0xa5, 0x2b, 0xde, 0x05,
	0xbc, 0x04, 0xc9, 0xea, 0x46, 0x1b, 0x4d, 0x0e, 0x47, 0x98, 0xed, 0x61, 0xd6, 0x05, 0x0b, 0x9a,
	0x6a, 0x28, 0x35, 0xf0, 0x24, 0x06, 0xc1, 0xbb, 0x20, 0x11, 0x26, 0x0e, 0x78, 0xaa, 0x55, 0x35,
	0x0c, 0x2d, 0xee, 0x4a, 0x2d, 0xb5, 0x2d, 0x79, 0x5f, 0xed, 0x4f, 0xa9, 0xd4, 0x5a, 0xae, 0x05,
	0xb7, 0x5d, 0xd2, 0xe6, 0x3c, 0x6b, 0x9b, 0xd8, 0xf4, 0x92, 0x03, 0xfe, 0x70, 0xda, 0xc4, 0x7c,
	0xac, 0x05, 0xf4, 0x7b, 0x24, 0x71, 0x95, 0xbd, 0x57, 0x99, 0x29, 0x46, 0x89, 0x7f, 0x29, 0x60,
	0x62, 0xd3, 0xc2, 0x80, 0x1f, 0xfd, 0x40, 0xf8, 0xf6, 0x09, 0xc8, 0xb7, 0x35, 0x88, 0xc6, 0x44,
	0xe2, 0x5d, 0x2b, 0xc0, 0x10, 0x82, 0xdd, 0xbc, 0xd1, 0xa5, 0x87, 0x96, 0x68, 0x35, 0x8f, 0x6c,
	0x4d, 0x6e, 0xe1, 0x99, 0xca, 0xbc, 0xd9, 0x12, 0xad, 0xdc, 0x68, 0xa6, 0x32, 0xe2, 0xe1, 0xeb,
	0x71, 0x96, 0x35, 0x02, 0xc0, 0xbb, 0x62, 0x69, 0x63, 0x4b, 0x5e, 0xe0, 0x83, 0x71, 0x4f, 0xcf,
	0x5d, 0xa2, 0xd5, 0x8d, 0xe3, 0x7b, 0x6c, 0x30, 0xc2, 0x46, 0x23, 0xec, 0xd5, 0x9e, 0x10, 0x1e,
	0x9c, 0x7f, 0xf5, 0x9d, 0xb3, 0x6f, 0x3e, 0x8a, 0xa6, 0x21, 0xf2, 0x12, 0xcf, 0x27, 0x1b, 0xde,
	0x55, 0xab, 0x70, 0x9f, 0x4d, 0xa9, 0x5a, 0x1f, 0xac, 0x0b, 0x58, 0x38, 0x72, 0x42, 0xb7, 0x17,
	0x89, 0x7e, 0x0f, 0x3d, 0xbf, 0xf9, 0x69, 0xe3, 0x3b, 0x67, 0x1b, 0x1f, 0xfd, 0xdc, 0xf8, 0xce,
	0xd1, 0x21, 0xbe, 0xf3, 0x87, 0x45, 0xa8, 0x75, 0x05, 0xe2, 0xb8, 0xc0, 0xf8, 0x04, 0xe4, 0x1b,
	0xd1, 0x74, 0x2a, 0x15, 0xe4, 0x14, 0xcf, 0x27, 0x0a, 0x79, 0xc4, 0xfe, 0xf3, 0x84, 0xec, 0xef,
	0x94, 0x16, 0x8f, 0x2f, 0xa3, 0x0d, 0x37, 0x85, 0xaf, 0xcf, 0xb7, 0x14, 0x5d, 0x6c, 0x29, 0xfa,
	0xbe, 0xa5, 0xe8, 0xf3, 0x8e, 0x3a, 0x17, 0x3b, 0xea, 0x7c, 0xd9, 0x51, 0xe7, 0xf4, 0xa9, 0x54,
	0xa6, 0x68, 0x13, 0x96, 0xea, 0x92, 0x8f, 0x5a, 0x4f, 0x74, 0x9e, 0xab, 0x54, 0xc5, 0x6b, 0x5e,
	0xb4, 0x09, 0xff, 0x30, 0xfd, 0x31, 0x6b, 0x3b, 0xb9, 0x66, 0x13, 0x7c, 0xf6, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x38, 0x35, 0x40, 0x6a, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgUpsert(ctx context.Context, in *MsgUpsertRequest, opts ...grpc.CallOption) (*MsgUpsertResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgUpsert(ctx context.Context, in *MsgUpsertRequest, opts ...grpc.CallOption) (*MsgUpsertResponse, error) {
	out := new(MsgUpsertResponse)
	err := c.cc.Invoke(ctx, "/sentinel.session.v1.MsgService/MsgUpsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgUpsert(context.Context, *MsgUpsertRequest) (*MsgUpsertResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgUpsert(ctx context.Context, req *MsgUpsertRequest) (*MsgUpsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUpsert not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgUpsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgUpsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.session.v1.MsgService/MsgUpsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgUpsert(ctx, req.(*MsgUpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.session.v1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgUpsert",
			Handler:    _MsgService_MsgUpsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/session/v1/msg.proto",
}

func (m *MsgUpsertRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpsertRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpsertRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintMsg(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpsertResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpsertResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpsertResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpsertRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovMsg(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovMsg(uint64(l))
	l = m.Bandwidth.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgUpsertResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpsertRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpsertRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpsertRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUpsertResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpsertResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpsertResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)