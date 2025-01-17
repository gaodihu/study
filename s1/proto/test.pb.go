// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package github_com_gaodihu_study_s1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type TestRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_4569d1102b9b6e3a, []int{0}
}
func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (dst *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(dst, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_4569d1102b9b6e3a, []int{1}
}
func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (dst *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(dst, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "github.com.gaodihu.study.s1.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "github.com.gaodihu.study.s1.TestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TestService service

type TestServiceClient interface {
	Hello(ctx context.Context, in *TestRequest, opts ...client.CallOption) (*TestResponse, error)
}

type testServiceClient struct {
	c           client.Client
	serviceName string
}

func NewTestServiceClient(serviceName string, c client.Client) TestServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "github.com.gaodihu.study.s1"
	}
	return &testServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *testServiceClient) Hello(ctx context.Context, in *TestRequest, opts ...client.CallOption) (*TestResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TestService.Hello", in)
	out := new(TestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceHandler interface {
	Hello(context.Context, *TestRequest, *TestResponse) error
}

func RegisterTestServiceHandler(s server.Server, hdlr TestServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&TestService{hdlr}, opts...))
}

type TestService struct {
	TestServiceHandler
}

func (h *TestService) Hello(ctx context.Context, in *TestRequest, out *TestResponse) error {
	return h.TestServiceHandler.Hello(ctx, in, out)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_4569d1102b9b6e3a) }

var fileDescriptor_test_4569d1102b9b6e3a = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4b, 0x4f, 0xcc, 0x4f, 0xc9, 0xcc, 0x28, 0xd5, 0x2b, 0x2e, 0x29, 0x4d, 0xa9,
	0xd4, 0x2b, 0x36, 0x54, 0x52, 0xe4, 0xe2, 0x0e, 0x49, 0x2d, 0x2e, 0x09, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x14, 0xb8, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85,
	0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0xa1, 0x4a, 0x40, 0x4c, 0xa3, 0x5c, 0x88, 0x21, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x71, 0x5c, 0xac, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x42, 0x1a,
	0x7a, 0x78, 0xac, 0xd6, 0x43, 0xb2, 0x57, 0x4a, 0x93, 0x08, 0x95, 0x10, 0xeb, 0x95, 0x18, 0x92,
	0xd8, 0xc0, 0xfe, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x56, 0x3a, 0x9c, 0xe5, 0x00,
	0x00, 0x00,
}
