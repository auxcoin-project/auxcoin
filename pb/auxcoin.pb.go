// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auxcoin.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	auxcoin.proto

It has these top-level messages:
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auxcoin service

type AuxcoinClient interface {
}

type auxcoinClient struct {
	cc *grpc.ClientConn
}

func NewAuxcoinClient(cc *grpc.ClientConn) AuxcoinClient {
	return &auxcoinClient{cc}
}

// Server API for Auxcoin service

type AuxcoinServer interface {
}

func RegisterAuxcoinServer(s *grpc.Server, srv AuxcoinServer) {
	s.RegisterService(&_Auxcoin_serviceDesc, srv)
}

var _Auxcoin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Auxcoin",
	HandlerType: (*AuxcoinServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "auxcoin.proto",
}

func init() { proto.RegisterFile("auxcoin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 55 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xad, 0x48,
	0xce, 0xcf, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x32, 0xe2, 0xe4,
	0x62, 0x77, 0x84, 0x08, 0x26, 0xb1, 0x81, 0x45, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb,
	0x05, 0x6d, 0x0d, 0x26, 0x00, 0x00, 0x00,
}