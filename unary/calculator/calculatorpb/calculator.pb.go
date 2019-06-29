// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unary/calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Calculator struct {
	Number1              int32    `protobuf:"varint,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              int32    `protobuf:"varint,2,opt,name=number2,proto3" json:"number2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculator) Reset()         { *m = Calculator{} }
func (m *Calculator) String() string { return proto.CompactTextString(m) }
func (*Calculator) ProtoMessage()    {}
func (*Calculator) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd71c1ca20eeece, []int{0}
}

func (m *Calculator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculator.Unmarshal(m, b)
}
func (m *Calculator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculator.Marshal(b, m, deterministic)
}
func (m *Calculator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculator.Merge(m, src)
}
func (m *Calculator) XXX_Size() int {
	return xxx_messageInfo_Calculator.Size(m)
}
func (m *Calculator) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculator.DiscardUnknown(m)
}

var xxx_messageInfo_Calculator proto.InternalMessageInfo

func (m *Calculator) GetNumber1() int32 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *Calculator) GetNumber2() int32 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type CalculatorRequest struct {
	Values               *Calculator `protobuf:"bytes,1,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd71c1ca20eeece, []int{1}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetValues() *Calculator {
	if m != nil {
		return m.Values
	}
	return nil
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd71c1ca20eeece, []int{2}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculator)(nil), "calculator.Calculator")
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
}

func init() {
	proto.RegisterFile("unary/calculator/calculatorpb/calculator.proto", fileDescriptor_2bd71c1ca20eeece)
}

var fileDescriptor_2bd71c1ca20eeece = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0xcd, 0x4b, 0x2c,
	0xaa, 0xd4, 0x4f, 0x4e, 0xcc, 0x49, 0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0x42, 0x62, 0x16, 0x24,
	0x21, 0x71, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x10, 0x22, 0x4a, 0x0e, 0x5c, 0x5c,
	0xce, 0x70, 0x9e, 0x90, 0x04, 0x17, 0x7b, 0x5e, 0x69, 0x6e, 0x52, 0x6a, 0x91, 0xa1, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x6b, 0x10, 0x8c, 0x8b, 0x90, 0x31, 0x92, 0x60, 0x42, 0x96, 0x31, 0x52, 0x72,
	0xe6, 0x12, 0x44, 0x98, 0x10, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc7, 0xc5, 0x56,
	0x96, 0x98, 0x53, 0x9a, 0x5a, 0x0c, 0x36, 0x87, 0xdb, 0x48, 0x4c, 0x0f, 0xc9, 0x15, 0x48, 0xca,
	0xa1, 0xaa, 0x94, 0x74, 0xb8, 0x84, 0x90, 0x0d, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12,
	0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x81, 0xba, 0x06, 0xca, 0x33, 0x4a, 0x44, 0xb6,
	0x32, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x87, 0x8b, 0x13, 0x26, 0x98, 0x2a, 0x24,
	0x8b, 0xc3, 0x3e, 0x88, 0xf3, 0xa4, 0xe4, 0x70, 0x49, 0x43, 0x2c, 0x56, 0x62, 0x70, 0xe2, 0x8b,
	0xe2, 0x41, 0x0e, 0xc4, 0x24, 0x36, 0x70, 0xd0, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x18,
	0xd1, 0xf6, 0x44, 0x6c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the calculator_client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Calculate(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Calculate(ctx context.Context, req *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculate(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculatorService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unary/calculator/calculatorpb/calculator.proto",
}