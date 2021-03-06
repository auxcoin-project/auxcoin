// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auxcoin.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	auxcoin.proto

It has these top-level messages:
	StatusRequest
	StatusResponse
	AddBlockRequest
	AddBlockResponse
*/
package pb

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

type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusResponse struct {
	Head   string `protobuf:"bytes,1,opt,name=head" json:"head,omitempty"`
	Bits   uint32 `protobuf:"varint,2,opt,name=bits" json:"bits,omitempty"`
	Reward uint32 `protobuf:"varint,3,opt,name=reward" json:"reward,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusResponse) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

func (m *StatusResponse) GetBits() uint32 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *StatusResponse) GetReward() uint32 {
	if m != nil {
		return m.Reward
	}
	return 0
}

type AddBlockRequest struct {
	Block string `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *AddBlockRequest) Reset()                    { *m = AddBlockRequest{} }
func (m *AddBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*AddBlockRequest) ProtoMessage()               {}
func (*AddBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddBlockRequest) GetBlock() string {
	if m != nil {
		return m.Block
	}
	return ""
}

type AddBlockResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *AddBlockResponse) Reset()                    { *m = AddBlockResponse{} }
func (m *AddBlockResponse) String() string            { return proto.CompactTextString(m) }
func (*AddBlockResponse) ProtoMessage()               {}
func (*AddBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddBlockResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "pb.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "pb.StatusResponse")
	proto.RegisterType((*AddBlockRequest)(nil), "pb.AddBlockRequest")
	proto.RegisterType((*AddBlockResponse)(nil), "pb.AddBlockResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auxcoin service

type AuxcoinClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error)
}

type auxcoinClient struct {
	cc *grpc.ClientConn
}

func NewAuxcoinClient(cc *grpc.ClientConn) AuxcoinClient {
	return &auxcoinClient{cc}
}

func (c *auxcoinClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/pb.Auxcoin/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auxcoinClient) AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error) {
	out := new(AddBlockResponse)
	err := grpc.Invoke(ctx, "/pb.Auxcoin/AddBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auxcoin service

type AuxcoinServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error)
}

func RegisterAuxcoinServer(s *grpc.Server, srv AuxcoinServer) {
	s.RegisterService(&_Auxcoin_serviceDesc, srv)
}

func _Auxcoin_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuxcoinServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auxcoin/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuxcoinServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auxcoin_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuxcoinServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auxcoin/AddBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuxcoinServer).AddBlock(ctx, req.(*AddBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auxcoin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Auxcoin",
	HandlerType: (*AuxcoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Auxcoin_Status_Handler,
		},
		{
			MethodName: "AddBlock",
			Handler:    _Auxcoin_AddBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auxcoin.proto",
}

func init() { proto.RegisterFile("auxcoin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcb, 0x4e, 0x86, 0x30,
	0x10, 0x85, 0xe5, 0x57, 0x50, 0x27, 0x41, 0x74, 0x24, 0x86, 0xb0, 0x22, 0xdd, 0xc8, 0x8a, 0x44,
	0x5d, 0xb8, 0xc6, 0x27, 0x30, 0xf8, 0x04, 0x2d, 0x6d, 0x22, 0xd1, 0xd0, 0xda, 0x4b, 0xf4, 0xf1,
	0x0d, 0xbd, 0x44, 0xf9, 0x77, 0x73, 0xbe, 0x9c, 0xce, 0x99, 0x53, 0x28, 0xa9, 0xfb, 0x99, 0xe5,
	0xb2, 0x0e, 0x4a, 0x4b, 0x2b, 0xf1, 0xa0, 0x18, 0xa9, 0xa0, 0x7c, 0xb3, 0xd4, 0x3a, 0x33, 0x89,
	0x2f, 0x27, 0x8c, 0x25, 0xaf, 0x70, 0x95, 0x80, 0x51, 0x72, 0x35, 0x02, 0x11, 0xce, 0xde, 0x05,
	0xe5, 0x4d, 0xd6, 0x65, 0xfd, 0xe5, 0xe4, 0xe7, 0x8d, 0xb1, 0xc5, 0x9a, 0xe6, 0xd0, 0x65, 0x7d,
	0x39, 0xf9, 0x19, 0xef, 0xa0, 0xd0, 0xe2, 0x9b, 0x6a, 0xde, 0x9c, 0x7a, 0x1a, 0x15, 0xb9, 0x87,
	0x6a, 0xe4, 0xfc, 0xe5, 0x53, 0xce, 0x1f, 0x31, 0x04, 0x6b, 0xc8, 0xd9, 0xa6, 0xe3, 0xce, 0x20,
	0x48, 0x0f, 0xd7, 0x7f, 0xc6, 0x18, 0x5e, 0x43, 0x2e, 0xb4, 0x96, 0x3a, 0x39, 0xbd, 0x78, 0x74,
	0x70, 0x3e, 0x86, 0x2a, 0xf8, 0x00, 0x45, 0xb8, 0x17, 0x6f, 0x06, 0xc5, 0x86, 0x5d, 0x99, 0x16,
	0xff, 0xa3, 0xb0, 0x91, 0x9c, 0xe0, 0x33, 0x5c, 0xa4, 0x1c, 0xbc, 0xdd, 0x1c, 0x47, 0xe7, 0xb5,
	0xf5, 0x1e, 0xa6, 0x87, 0xac, 0xf0, 0xff, 0xf6, 0xf4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x86, 0xa3,
	0xa3, 0x9f, 0x48, 0x01, 0x00, 0x00,
}
