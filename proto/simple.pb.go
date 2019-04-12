// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/simple.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SimpleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4285f8f75e6ba5cf, []int{0}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

func (m *SimpleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SimpleResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4285f8f75e6ba5cf, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "proto.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "proto.SimpleResponse")
}

func init() { proto.RegisterFile("proto/simple.proto", fileDescriptor_4285f8f75e6ba5cf) }

var fileDescriptor_4285f8f75e6ba5cf = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x92,
	0x32, 0x17, 0x6f, 0x30, 0x58, 0x38, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b,
	0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x52, 0xe1,
	0xe2, 0x83, 0x29, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xc5, 0xa6, 0xca, 0xc8, 0x8b, 0x8b, 0x1b,
	0xa2, 0x2a, 0x00, 0x6c, 0x81, 0x35, 0x17, 0x97, 0x7b, 0x6a, 0x89, 0x6f, 0x6a, 0x71, 0x71, 0x62,
	0x7a, 0xaa, 0x90, 0x08, 0xc4, 0x5a, 0x3d, 0x14, 0xcb, 0xa4, 0x44, 0xd1, 0x44, 0x21, 0xa6, 0x2b,
	0x31, 0x24, 0xb1, 0x81, 0xc5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0x32, 0xea, 0x59,
	0xba, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimpleProtoClient is the client API for SimpleProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleProtoClient interface {
	GetMessage(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type simpleProtoClient struct {
	cc *grpc.ClientConn
}

func NewSimpleProtoClient(cc *grpc.ClientConn) SimpleProtoClient {
	return &simpleProtoClient{cc}
}

func (c *simpleProtoClient) GetMessage(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/proto.SimpleProto/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleProtoServer is the server API for SimpleProto service.
type SimpleProtoServer interface {
	GetMessage(context.Context, *SimpleRequest) (*SimpleResponse, error)
}

func RegisterSimpleProtoServer(s *grpc.Server, srv SimpleProtoServer) {
	s.RegisterService(&_SimpleProto_serviceDesc, srv)
}

func _SimpleProto_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleProtoServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SimpleProto/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleProtoServer).GetMessage(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimpleProto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SimpleProto",
	HandlerType: (*SimpleProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _SimpleProto_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/simple.proto",
}
