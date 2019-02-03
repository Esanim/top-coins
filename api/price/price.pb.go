// Code generated by protoc-gen-go. DO NOT EDIT.
// source: price.proto

package pb_price

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PriceRequest struct {
	Limit                uint64   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceRequest) Reset()         { *m = PriceRequest{} }
func (m *PriceRequest) String() string { return proto.CompactTextString(m) }
func (*PriceRequest) ProtoMessage()    {}
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_price_62f082455f936a9f, []int{0}
}
func (m *PriceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceRequest.Unmarshal(m, b)
}
func (m *PriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceRequest.Marshal(b, m, deterministic)
}
func (dst *PriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceRequest.Merge(dst, src)
}
func (m *PriceRequest) XXX_Size() int {
	return xxx_messageInfo_PriceRequest.Size(m)
}
func (m *PriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PriceRequest proto.InternalMessageInfo

func (m *PriceRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Price struct {
	Id                   float64  `protobuf:"fixed64,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	PriceUSD             float64  `protobuf:"fixed64,4,opt,name=PriceUSD,proto3" json:"PriceUSD,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_price_62f082455f936a9f, []int{1}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (dst *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(dst, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetId() float64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Price) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Price) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Price) GetPriceUSD() float64 {
	if m != nil {
		return m.PriceUSD
	}
	return 0
}

type PriceResponse struct {
	Items                []*Price `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceResponse) Reset()         { *m = PriceResponse{} }
func (m *PriceResponse) String() string { return proto.CompactTextString(m) }
func (*PriceResponse) ProtoMessage()    {}
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_price_62f082455f936a9f, []int{2}
}
func (m *PriceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceResponse.Unmarshal(m, b)
}
func (m *PriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceResponse.Marshal(b, m, deterministic)
}
func (dst *PriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceResponse.Merge(dst, src)
}
func (m *PriceResponse) XXX_Size() int {
	return xxx_messageInfo_PriceResponse.Size(m)
}
func (m *PriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PriceResponse proto.InternalMessageInfo

func (m *PriceResponse) GetItems() []*Price {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*PriceRequest)(nil), "pb_price.PriceRequest")
	proto.RegisterType((*Price)(nil), "pb_price.Price")
	proto.RegisterType((*PriceResponse)(nil), "pb_price.PriceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PriceServiceClient interface {
	GetPrices(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error)
}

type priceServiceClient struct {
	cc *grpc.ClientConn
}

func NewPriceServiceClient(cc *grpc.ClientConn) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) GetPrices(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*PriceResponse, error) {
	out := new(PriceResponse)
	err := c.cc.Invoke(ctx, "/pb_price.PriceService/GetPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceServiceServer is the server API for PriceService service.
type PriceServiceServer interface {
	GetPrices(context.Context, *PriceRequest) (*PriceResponse, error)
}

func RegisterPriceServiceServer(s *grpc.Server, srv PriceServiceServer) {
	s.RegisterService(&_PriceService_serviceDesc, srv)
}

func _PriceService_GetPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_price.PriceService/GetPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetPrices(ctx, req.(*PriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PriceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_price.PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrices",
			Handler:    _PriceService_GetPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "price.proto",
}

func init() { proto.RegisterFile("price.proto", fileDescriptor_price_62f082455f936a9f) }

var fileDescriptor_price_62f082455f936a9f = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x4c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x48, 0x8a, 0x07, 0xf3, 0x95, 0x54,
	0xb8, 0x78, 0x02, 0x40, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6,
	0x9c, 0xcc, 0xdc, 0xcc, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x29, 0x9e,
	0x8b, 0x15, 0xac, 0x4a, 0x88, 0x8f, 0x8b, 0xc9, 0x33, 0x05, 0x2c, 0xc7, 0x18, 0xc4, 0xe4, 0x99,
	0x22, 0x24, 0xc4, 0xc5, 0xe2, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0x66, 0x0b, 0x89, 0x71, 0xb1, 0x05, 0x57, 0xe6, 0x26, 0xe5, 0xe7, 0x48, 0x30, 0x83, 0x45, 0xa1,
	0x3c, 0x21, 0x29, 0x2e, 0x0e, 0xb0, 0x21, 0xa1, 0xc1, 0x2e, 0x12, 0x2c, 0x60, 0x13, 0xe0, 0x7c,
	0x25, 0x33, 0x2e, 0x5e, 0xa8, 0x33, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x54, 0xb9, 0x58,
	0x33, 0x4b, 0x52, 0x73, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xf8, 0xf5, 0x60, 0x2e,
	0xd6, 0x83, 0xa8, 0x83, 0xc8, 0x1a, 0xf9, 0x41, 0x9d, 0x1f, 0x9c, 0x5a, 0x54, 0x06, 0x72, 0x9f,
	0x1d, 0x17, 0xa7, 0x7b, 0x6a, 0x09, 0x58, 0xa8, 0x58, 0x48, 0x0c, 0x5d, 0x13, 0xc4, 0x8f, 0x52,
	0xe2, 0x18, 0xe2, 0x10, 0x4b, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xe1, 0x63, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0x4a, 0xb2, 0x5a, 0x2e, 0x01, 0x00, 0x00,
}
