// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package core_grpc is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	RequestPing
	RequestBroadcastTx
	ResponsePing
	ResponseBroadcastTx
*/
package core_grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/bdc/abci/types"

import (
	"context"

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

type RequestPing struct {
}

func (m *RequestPing) Reset()                    { *m = RequestPing{} }
func (m *RequestPing) String() string            { return proto.CompactTextString(m) }
func (*RequestPing) ProtoMessage()               {}
func (*RequestPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RequestBroadcastTx struct {
	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *RequestBroadcastTx) Reset()                    { *m = RequestBroadcastTx{} }
func (m *RequestBroadcastTx) String() string            { return proto.CompactTextString(m) }
func (*RequestBroadcastTx) ProtoMessage()               {}
func (*RequestBroadcastTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RequestBroadcastTx) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type ResponsePing struct {
}

func (m *ResponsePing) Reset()                    { *m = ResponsePing{} }
func (m *ResponsePing) String() string            { return proto.CompactTextString(m) }
func (*ResponsePing) ProtoMessage()               {}
func (*ResponsePing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ResponseBroadcastTx struct {
	CheckTx   *types.ResponseCheckTx   `protobuf:"bytes,1,opt,name=check_tx,json=checkTx" json:"check_tx,omitempty"`
	DeliverTx *types.ResponseDeliverTx `protobuf:"bytes,2,opt,name=deliver_tx,json=deliverTx" json:"deliver_tx,omitempty"`
}

func (m *ResponseBroadcastTx) Reset()                    { *m = ResponseBroadcastTx{} }
func (m *ResponseBroadcastTx) String() string            { return proto.CompactTextString(m) }
func (*ResponseBroadcastTx) ProtoMessage()               {}
func (*ResponseBroadcastTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResponseBroadcastTx) GetCheckTx() *types.ResponseCheckTx {
	if m != nil {
		return m.CheckTx
	}
	return nil
}

func (m *ResponseBroadcastTx) GetDeliverTx() *types.ResponseDeliverTx {
	if m != nil {
		return m.DeliverTx
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestPing)(nil), "core_grpc.RequestPing")
	proto.RegisterType((*RequestBroadcastTx)(nil), "core_grpc.RequestBroadcastTx")
	proto.RegisterType((*ResponsePing)(nil), "core_grpc.ResponsePing")
	proto.RegisterType((*ResponseBroadcastTx)(nil), "core_grpc.ResponseBroadcastTx")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BroadcastAPI service

type BroadcastAPIClient interface {
	Ping(ctx context.Context, in *RequestPing, opts ...grpc.CallOption) (*ResponsePing, error)
	BroadcastTx(ctx context.Context, in *RequestBroadcastTx, opts ...grpc.CallOption) (*ResponseBroadcastTx, error)
}

type broadcastAPIClient struct {
	cc *grpc.ClientConn
}

func NewBroadcastAPIClient(cc *grpc.ClientConn) BroadcastAPIClient {
	return &broadcastAPIClient{cc}
}

func (c *broadcastAPIClient) Ping(ctx context.Context, in *RequestPing, opts ...grpc.CallOption) (*ResponsePing, error) {
	out := new(ResponsePing)
	err := grpc.Invoke(ctx, "/core_grpc.BroadcastAPI/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastAPIClient) BroadcastTx(ctx context.Context, in *RequestBroadcastTx, opts ...grpc.CallOption) (*ResponseBroadcastTx, error) {
	out := new(ResponseBroadcastTx)
	err := grpc.Invoke(ctx, "/core_grpc.BroadcastAPI/BroadcastTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BroadcastAPI service

type BroadcastAPIServer interface {
	Ping(context.Context, *RequestPing) (*ResponsePing, error)
	BroadcastTx(context.Context, *RequestBroadcastTx) (*ResponseBroadcastTx, error)
}

func RegisterBroadcastAPIServer(s *grpc.Server, srv BroadcastAPIServer) {
	s.RegisterService(&_BroadcastAPI_serviceDesc, srv)
}

func _BroadcastAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core_grpc.BroadcastAPI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastAPIServer).Ping(ctx, req.(*RequestPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadcastAPI_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBroadcastTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastAPIServer).BroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core_grpc.BroadcastAPI/BroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastAPIServer).BroadcastTx(ctx, req.(*RequestBroadcastTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _BroadcastAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core_grpc.BroadcastAPI",
	HandlerType: (*BroadcastAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BroadcastAPI_Ping_Handler,
		},
		{
			MethodName: "BroadcastTx",
			Handler:    _BroadcastAPI_BroadcastTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xce, 0x2f, 0x4a, 0x8d, 0x4f, 0x2f,
	0x2a, 0x48, 0x96, 0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f,
	0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0x4c, 0x4a, 0xce, 0xd4, 0x07,
	0x6b, 0xd1, 0x47, 0xd2, 0xa8, 0xc4, 0xcb, 0xc5, 0x1d, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x12,
	0x90, 0x99, 0x97, 0xae, 0xa4, 0xc2, 0x25, 0x04, 0xe5, 0x3a, 0x15, 0xe5, 0x27, 0xa6, 0x24, 0x27,
	0x16, 0x97, 0x84, 0x54, 0x08, 0xf1, 0x71, 0x31, 0x95, 0x54, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x04, 0x31, 0x95, 0x54, 0x28, 0xf1, 0x71, 0xf1, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x82, 0x75, 0x35, 0x32, 0x72, 0x09, 0xc3, 0x04, 0x90, 0xf5, 0x19, 0x72, 0x71, 0x24, 0x67, 0xa4,
	0x26, 0x67, 0xc7, 0x43, 0x75, 0x73, 0x1b, 0x89, 0xe9, 0x41, 0x2c, 0x87, 0xa9, 0x76, 0x06, 0x49,
	0x87, 0x54, 0x04, 0xb1, 0x27, 0x43, 0x18, 0x42, 0xe6, 0x5c, 0x5c, 0x29, 0xa9, 0x39, 0x99, 0x65,
	0xa9, 0x45, 0x20, 0x4d, 0x4c, 0x60, 0x4d, 0x12, 0x68, 0x9a, 0x5c, 0x20, 0x0a, 0x42, 0x2a, 0x82,
	0x38, 0x53, 0x60, 0x4c, 0xa3, 0xa9, 0x8c, 0x5c, 0x3c, 0x70, 0xbb, 0x1d, 0x03, 0x3c, 0x85, 0xcc,
	0xb9, 0x58, 0x40, 0x8e, 0x13, 0x12, 0xd3, 0x83, 0x87, 0x8d, 0x1e, 0x92, 0x57, 0xa5, 0xc4, 0x51,
	0xc4, 0x11, 0xbe, 0x11, 0xf2, 0xe1, 0xe2, 0x46, 0xf6, 0x84, 0x2c, 0xa6, 0x7e, 0x24, 0x69, 0x29,
	0x39, 0x2c, 0xc6, 0x20, 0xc9, 0x27, 0xb1, 0x81, 0xc3, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x92, 0x29, 0xd9, 0x42, 0xaf, 0x01, 0x00, 0x00,
}
