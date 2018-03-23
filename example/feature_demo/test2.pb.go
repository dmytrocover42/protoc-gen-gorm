// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/feature_demo/test2.proto

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IntPoint struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	X  int32  `protobuf:"varint,2,opt,name=x" json:"x,omitempty"`
	Y  int32  `protobuf:"varint,3,opt,name=y" json:"y,omitempty"`
}

func (m *IntPoint) Reset()                    { *m = IntPoint{} }
func (m *IntPoint) String() string            { return proto.CompactTextString(m) }
func (*IntPoint) ProtoMessage()               {}
func (*IntPoint) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *IntPoint) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IntPoint) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *IntPoint) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*IntPoint)(nil), "example.int_point")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PointService service

type PointServiceClient interface {
	CreateIntPoint(ctx context.Context, in *IntPoint, opts ...grpc.CallOption) (*IntPoint, error)
	ReadIntPoint(ctx context.Context, in *IntPoint, opts ...grpc.CallOption) (*IntPoint, error)
	UpdateIntPoint(ctx context.Context, in *IntPoint, opts ...grpc.CallOption) (*IntPoint, error)
}

type pointServiceClient struct {
	cc *grpc.ClientConn
}

func NewPointServiceClient(cc *grpc.ClientConn) PointServiceClient {
	return &pointServiceClient{cc}
}

func (c *pointServiceClient) CreateIntPoint(ctx context.Context, in *IntPoint, opts ...grpc.CallOption) (*IntPoint, error) {
	out := new(IntPoint)
	err := grpc.Invoke(ctx, "/example.PointService/CreateIntPoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) ReadIntPoint(ctx context.Context, in *IntPoint, opts ...grpc.CallOption) (*IntPoint, error) {
	out := new(IntPoint)
	err := grpc.Invoke(ctx, "/example.PointService/ReadIntPoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) UpdateIntPoint(ctx context.Context, in *IntPoint, opts ...grpc.CallOption) (*IntPoint, error) {
	out := new(IntPoint)
	err := grpc.Invoke(ctx, "/example.PointService/UpdateIntPoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PointService service

type PointServiceServer interface {
	CreateIntPoint(context.Context, *IntPoint) (*IntPoint, error)
	ReadIntPoint(context.Context, *IntPoint) (*IntPoint, error)
	UpdateIntPoint(context.Context, *IntPoint) (*IntPoint, error)
}

func RegisterPointServiceServer(s *grpc.Server, srv PointServiceServer) {
	s.RegisterService(&_PointService_serviceDesc, srv)
}

func _PointService_CreateIntPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntPoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).CreateIntPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PointService/CreateIntPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).CreateIntPoint(ctx, req.(*IntPoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_ReadIntPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntPoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).ReadIntPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PointService/ReadIntPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).ReadIntPoint(ctx, req.(*IntPoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_UpdateIntPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntPoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).UpdateIntPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PointService/UpdateIntPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).UpdateIntPoint(ctx, req.(*IntPoint))
	}
	return interceptor(ctx, in, info, handler)
}

var _PointService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.PointService",
	HandlerType: (*PointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIntPoint",
			Handler:    _PointService_CreateIntPoint_Handler,
		},
		{
			MethodName: "ReadIntPoint",
			Handler:    _PointService_ReadIntPoint_Handler,
		},
		{
			MethodName: "UpdateIntPoint",
			Handler:    _PointService_UpdateIntPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/feature_demo/test2.proto",
}

func init() { proto.RegisterFile("example/feature_demo/test2.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x3f, 0x4b, 0xc4, 0x40,
	0x14, 0xc4, 0xdd, 0x88, 0xa7, 0x2e, 0x31, 0xc5, 0x56, 0x31, 0x8d, 0xe1, 0xaa, 0x6b, 0x2e, 0x0b,
	0x67, 0x23, 0xb1, 0x10, 0x54, 0x04, 0x3b, 0x89, 0xd8, 0xd8, 0x1c, 0x49, 0xf6, 0x5d, 0x5c, 0xb8,
	0xec, 0x5b, 0xf6, 0xde, 0x49, 0xee, 0xa3, 0x69, 0xe5, 0x47, 0x93, 0xfc, 0xc1, 0xea, 0x84, 0x23,
	0xe5, 0xcc, 0xfe, 0x66, 0xe7, 0x31, 0x3c, 0x86, 0x26, 0xaf, 0xed, 0x1a, 0xe4, 0x0a, 0x72, 0xda,
	0x3a, 0x58, 0x2a, 0xa8, 0x51, 0x12, 0x6c, 0x68, 0x91, 0x58, 0x87, 0x84, 0xe2, 0x74, 0x20, 0xa2,
	0xb4, 0xd2, 0xf4, 0xb1, 0x2d, 0x92, 0x12, 0x6b, 0xa9, 0xcd, 0x0a, 0x8b, 0x35, 0x36, 0x68, 0xc1,
	0xc8, 0x8e, 0x2b, 0xe7, 0x15, 0x98, 0x79, 0x85, 0xae, 0x96, 0x68, 0x49, 0xa3, 0xd9, 0xc8, 0x56,
	0xf4, 0x9f, 0x44, 0x57, 0xff, 0xd6, 0xf4, 0xc0, 0xf4, 0x8e, 0x9f, 0x6b, 0x43, 0x4b, 0x8b, 0xda,
	0x90, 0x08, 0xb8, 0xa7, 0x55, 0xc8, 0x62, 0x36, 0xbb, 0xc8, 0x3c, 0xad, 0x84, 0xcf, 0x59, 0x13,
	0x7a, 0x31, 0x9b, 0x9d, 0x64, 0xac, 0x69, 0xd5, 0x2e, 0x3c, 0xee, 0xd5, 0x2e, 0x9d, 0x7c, 0x7f,
	0x5d, 0x7a, 0x67, 0x6c, 0xf1, 0xc3, 0xb8, 0xff, 0xd2, 0xa6, 0x5f, 0xc1, 0x7d, 0xea, 0x12, 0x44,
	0xca, 0x83, 0x07, 0x07, 0x39, 0xc1, 0xb3, 0xa1, 0xee, 0x41, 0x88, 0x64, 0xb8, 0x22, 0xf9, 0xab,
	0x8a, 0xf6, 0x78, 0xd3, 0x23, 0x71, 0xc3, 0xfd, 0x0c, 0x72, 0x35, 0x22, 0x99, 0xf2, 0xe0, 0xcd,
	0xaa, 0x51, 0xad, 0xf7, 0x4f, 0xef, 0x8f, 0x87, 0x4e, 0xbc, 0x6f, 0xce, 0xdb, 0xc1, 0x2c, 0x26,
	0x1d, 0x7d, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x30, 0x2e, 0x1a, 0xa0, 0xdc, 0x01, 0x00, 0x00,
}
