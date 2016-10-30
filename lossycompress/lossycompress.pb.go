// Code generated by protoc-gen-go.
// source: lossycompress.proto
// DO NOT EDIT!

/*
Package lossycompress is a generated protocol buffer package.

It is generated from these files:
	lossycompress.proto

It has these top-level messages:
	Request
	Result
*/
package lossycompress

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

type Request struct {
	Data    string `protobuf:"bytes,1,opt,name=Data,json=data" json:"Data,omitempty"`
	Quality int32  `protobuf:"varint,2,opt,name=Quality,json=quality" json:"Quality,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	Link string `protobuf:"bytes,1,opt,name=Link,json=link" json:"Link,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Result)(nil), "Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Img service

type ImgClient interface {
	Compress(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type imgClient struct {
	cc *grpc.ClientConn
}

func NewImgClient(cc *grpc.ClientConn) ImgClient {
	return &imgClient{cc}
}

func (c *imgClient) Compress(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/Img/Compress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Img service

type ImgServer interface {
	Compress(context.Context, *Request) (*Result, error)
}

func RegisterImgServer(s *grpc.Server, srv ImgServer) {
	s.RegisterService(&_Img_serviceDesc, srv)
}

func _Img_Compress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImgServer).Compress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Img/Compress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImgServer).Compress(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Img_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Img",
	HandlerType: (*ImgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compress",
			Handler:    _Img_Compress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("lossycompress.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xc9, 0x2f, 0x2e,
	0xae, 0x4c, 0xce, 0xcf, 0x2d, 0x28, 0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57,
	0xf2, 0xe6, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x71, 0x49,
	0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x14, 0x92,
	0xe0, 0x62, 0x0f, 0x2c, 0x4d, 0xcc, 0xc9, 0x2c, 0xa9, 0x94, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d,
	0x62, 0x2f, 0x84, 0x70, 0x41, 0xaa, 0xfd, 0x12, 0x73, 0x53, 0x25, 0x98, 0x21, 0xaa, 0xf3, 0x12,
	0x73, 0x53, 0x95, 0x0c, 0xb8, 0xd8, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0xc0, 0x66, 0xf9, 0x64, 0xe6,
	0x65, 0xc3, 0xcc, 0xca, 0xc9, 0xcc, 0xcb, 0x86, 0xeb, 0x60, 0x42, 0xe8, 0x30, 0x52, 0xe3, 0x62,
	0xf6, 0xcc, 0x4d, 0x17, 0x92, 0xe7, 0xe2, 0x70, 0x86, 0xba, 0x4b, 0x88, 0x43, 0x0f, 0xea, 0x20,
	0x29, 0x76, 0x3d, 0x88, 0x69, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xd7, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x2e, 0x53, 0x35, 0xc4, 0x00, 0x00, 0x00,
}