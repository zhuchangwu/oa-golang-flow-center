// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flowCenter.proto

package flow_center

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

// 定义消息，消息为service中发送或者是接受的内容
type Flow struct {
	FlowType             string            `protobuf:"bytes,1,opt,name=flowType,proto3" json:"flowType,omitempty"`
	RoleMap              map[string]string `protobuf:"bytes,2,rep,name=roleMap,proto3" json:"roleMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Applicate            string            `protobuf:"bytes,3,opt,name=applicate,proto3" json:"applicate,omitempty"`
	DepartmentId         int32             `protobuf:"varint,4,opt,name=departmentId,proto3" json:"departmentId,omitempty"`
	RecordId             int32             `protobuf:"varint,5,opt,name=recordId,proto3" json:"recordId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Flow) Reset()         { *m = Flow{} }
func (m *Flow) String() string { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()    {}
func (*Flow) Descriptor() ([]byte, []int) {
	return fileDescriptor_84ee69aac8680289, []int{0}
}

func (m *Flow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flow.Unmarshal(m, b)
}
func (m *Flow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flow.Marshal(b, m, deterministic)
}
func (m *Flow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flow.Merge(m, src)
}
func (m *Flow) XXX_Size() int {
	return xxx_messageInfo_Flow.Size(m)
}
func (m *Flow) XXX_DiscardUnknown() {
	xxx_messageInfo_Flow.DiscardUnknown(m)
}

var xxx_messageInfo_Flow proto.InternalMessageInfo

func (m *Flow) GetFlowType() string {
	if m != nil {
		return m.FlowType
	}
	return ""
}

func (m *Flow) GetRoleMap() map[string]string {
	if m != nil {
		return m.RoleMap
	}
	return nil
}

func (m *Flow) GetApplicate() string {
	if m != nil {
		return m.Applicate
	}
	return ""
}

func (m *Flow) GetDepartmentId() int32 {
	if m != nil {
		return m.DepartmentId
	}
	return 0
}

func (m *Flow) GetRecordId() int32 {
	if m != nil {
		return m.RecordId
	}
	return 0
}

type Response struct {
	FlowId               int32    `protobuf:"varint,1,opt,name=flowId,proto3" json:"flowId,omitempty"`
	ResponseMsg          string   `protobuf:"bytes,2,opt,name=responseMsg,proto3" json:"responseMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_84ee69aac8680289, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetFlowId() int32 {
	if m != nil {
		return m.FlowId
	}
	return 0
}

func (m *Response) GetResponseMsg() string {
	if m != nil {
		return m.ResponseMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Flow)(nil), "proto.Flow")
	proto.RegisterMapType((map[string]string)(nil), "proto.Flow.RoleMapEntry")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("flowCenter.proto", fileDescriptor_84ee69aac8680289) }

var fileDescriptor_84ee69aac8680289 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0xef, 0x14, 0xca, 0x85, 0x53, 0x72, 0x2f, 0x99, 0x18, 0xd3, 0x10, 0x13, 0x9b, 0xae,
	0xba, 0x50, 0x24, 0xb8, 0x31, 0x2c, 0x5b, 0x30, 0x61, 0xc1, 0xa6, 0xb8, 0x72, 0x37, 0xb6, 0x47,
	0xd3, 0x58, 0x3a, 0xe3, 0x74, 0x80, 0xf4, 0x21, 0x7c, 0x57, 0x1f, 0xc1, 0xcc, 0x50, 0xa0, 0xae,
	0x64, 0xd5, 0x7e, 0xdf, 0x7c, 0xf3, 0x3b, 0x7f, 0x32, 0x30, 0x78, 0xcd, 0xf9, 0x2e, 0xc2, 0x42,
	0xa1, 0x1c, 0x09, 0xc9, 0x15, 0xa7, 0xb6, 0xf9, 0xf8, 0x5f, 0x04, 0xda, 0x8f, 0x39, 0xdf, 0xd1,
	0x21, 0x74, 0x75, 0xe6, 0xa9, 0x12, 0xe8, 0x12, 0x8f, 0x04, 0xbd, 0xf8, 0xa8, 0xe9, 0x04, 0xfe,
	0x4a, 0x9e, 0xe3, 0x92, 0x09, 0xd7, 0xf2, 0x5a, 0x81, 0x33, 0x71, 0xf7, 0x90, 0x91, 0xbe, 0x39,
	0x8a, 0xf7, 0x47, 0xf3, 0x42, 0xc9, 0x2a, 0x3e, 0x04, 0xe9, 0x15, 0xf4, 0x98, 0x10, 0x79, 0x96,
	0x30, 0x85, 0x6e, 0xcb, 0x00, 0x4f, 0x06, 0xf5, 0xa1, 0x9f, 0xa2, 0x60, 0x52, 0xad, 0xb1, 0x50,
	0x8b, 0xd4, 0x6d, 0x7b, 0x24, 0xb0, 0xe3, 0x1f, 0x9e, 0xee, 0x48, 0x62, 0xc2, 0x65, 0xba, 0x48,
	0x5d, 0xdb, 0x9c, 0x1f, 0xf5, 0x70, 0x0a, 0xfd, 0x66, 0x59, 0x3a, 0x80, 0xd6, 0x3b, 0x56, 0x75,
	0xe3, 0xfa, 0x97, 0x5e, 0x80, 0xbd, 0x65, 0xf9, 0x06, 0x5d, 0xcb, 0x78, 0x7b, 0x31, 0xb5, 0x1e,
	0x88, 0x3f, 0x83, 0x6e, 0x8c, 0xa5, 0xe0, 0x45, 0x89, 0xf4, 0x12, 0x3a, 0x7a, 0xca, 0x45, 0x6a,
	0xae, 0xda, 0x71, 0xad, 0xa8, 0x07, 0x8e, 0xac, 0x33, 0xcb, 0xf2, 0xad, 0x66, 0x34, 0xad, 0xc9,
	0xa7, 0x05, 0x8e, 0x1e, 0x7f, 0x85, 0x72, 0x9b, 0x25, 0x48, 0x6f, 0x00, 0x22, 0x89, 0x4c, 0xa1,
	0xd9, 0xa6, 0xd3, 0x58, 0xd0, 0xf0, 0x7f, 0x2d, 0x0e, 0x55, 0xfd, 0x3f, 0x74, 0x06, 0xde, 0x29,
	0x7d, 0xf0, 0xc3, 0x4a, 0xc3, 0x50, 0xae, 0x94, 0x44, 0xb6, 0xfe, 0x8d, 0x31, 0x26, 0x34, 0x82,
	0xeb, 0x26, 0xe5, 0x63, 0x83, 0xa5, 0x0a, 0xab, 0x28, 0xcf, 0xb0, 0x50, 0xe7, 0x41, 0x02, 0x42,
	0xe7, 0x4d, 0x48, 0x58, 0x85, 0x59, 0x9a, 0x49, 0x4c, 0x54, 0xc6, 0x0b, 0x96, 0x9f, 0x0b, 0x19,
	0x93, 0x70, 0xf0, 0xfc, 0x8f, 0xb3, 0x3b, 0xbd, 0xbe, 0xdb, 0xc4, 0xbc, 0xb3, 0x97, 0x8e, 0xc9,
	0xdd, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xad, 0x30, 0xfa, 0x40, 0x7c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlowServiceClient is the client API for FlowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlowServiceClient interface {
	CreateFlow(ctx context.Context, in *Flow, opts ...grpc.CallOption) (*Response, error)
	// 服务端返回stream
	CreateFlowResponseByServerStream(ctx context.Context, in *Flow, opts ...grpc.CallOption) (FlowService_CreateFlowResponseByServerStreamClient, error)
	// 客户端往服务端发送stream
	CreateFlowRequestByClientStream(ctx context.Context, opts ...grpc.CallOption) (FlowService_CreateFlowRequestByClientStreamClient, error)
	// 客户端，服务端双向流通信
	CreateFlowByBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (FlowService_CreateFlowByBidirectionalStreamClient, error)
}

type flowServiceClient struct {
	cc *grpc.ClientConn
}

func NewFlowServiceClient(cc *grpc.ClientConn) FlowServiceClient {
	return &flowServiceClient{cc}
}

func (c *flowServiceClient) CreateFlow(ctx context.Context, in *Flow, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.FlowService/CreateFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) CreateFlowResponseByServerStream(ctx context.Context, in *Flow, opts ...grpc.CallOption) (FlowService_CreateFlowResponseByServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FlowService_serviceDesc.Streams[0], "/proto.FlowService/CreateFlowResponseByServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceCreateFlowResponseByServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowService_CreateFlowResponseByServerStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type flowServiceCreateFlowResponseByServerStreamClient struct {
	grpc.ClientStream
}

func (x *flowServiceCreateFlowResponseByServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowServiceClient) CreateFlowRequestByClientStream(ctx context.Context, opts ...grpc.CallOption) (FlowService_CreateFlowRequestByClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FlowService_serviceDesc.Streams[1], "/proto.FlowService/CreateFlowRequestByClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceCreateFlowRequestByClientStreamClient{stream}
	return x, nil
}

type FlowService_CreateFlowRequestByClientStreamClient interface {
	Send(*Flow) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type flowServiceCreateFlowRequestByClientStreamClient struct {
	grpc.ClientStream
}

func (x *flowServiceCreateFlowRequestByClientStreamClient) Send(m *Flow) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceCreateFlowRequestByClientStreamClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowServiceClient) CreateFlowByBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (FlowService_CreateFlowByBidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FlowService_serviceDesc.Streams[2], "/proto.FlowService/CreateFlowByBidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceCreateFlowByBidirectionalStreamClient{stream}
	return x, nil
}

type FlowService_CreateFlowByBidirectionalStreamClient interface {
	Send(*Flow) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type flowServiceCreateFlowByBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *flowServiceCreateFlowByBidirectionalStreamClient) Send(m *Flow) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceCreateFlowByBidirectionalStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlowServiceServer is the server API for FlowService service.
type FlowServiceServer interface {
	CreateFlow(context.Context, *Flow) (*Response, error)
	// 服务端返回stream
	CreateFlowResponseByServerStream(*Flow, FlowService_CreateFlowResponseByServerStreamServer) error
	// 客户端往服务端发送stream
	CreateFlowRequestByClientStream(FlowService_CreateFlowRequestByClientStreamServer) error
	// 客户端，服务端双向流通信
	CreateFlowByBidirectionalStream(FlowService_CreateFlowByBidirectionalStreamServer) error
}

func RegisterFlowServiceServer(s *grpc.Server, srv FlowServiceServer) {
	s.RegisterService(&_FlowService_serviceDesc, srv)
}

func _FlowService_CreateFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Flow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).CreateFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FlowService/CreateFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).CreateFlow(ctx, req.(*Flow))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_CreateFlowResponseByServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Flow)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowServiceServer).CreateFlowResponseByServerStream(m, &flowServiceCreateFlowResponseByServerStreamServer{stream})
}

type FlowService_CreateFlowResponseByServerStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type flowServiceCreateFlowResponseByServerStreamServer struct {
	grpc.ServerStream
}

func (x *flowServiceCreateFlowResponseByServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowService_CreateFlowRequestByClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).CreateFlowRequestByClientStream(&flowServiceCreateFlowRequestByClientStreamServer{stream})
}

type FlowService_CreateFlowRequestByClientStreamServer interface {
	SendAndClose(*Response) error
	Recv() (*Flow, error)
	grpc.ServerStream
}

type flowServiceCreateFlowRequestByClientStreamServer struct {
	grpc.ServerStream
}

func (x *flowServiceCreateFlowRequestByClientStreamServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceCreateFlowRequestByClientStreamServer) Recv() (*Flow, error) {
	m := new(Flow)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowService_CreateFlowByBidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).CreateFlowByBidirectionalStream(&flowServiceCreateFlowByBidirectionalStreamServer{stream})
}

type FlowService_CreateFlowByBidirectionalStreamServer interface {
	Send(*Response) error
	Recv() (*Flow, error)
	grpc.ServerStream
}

type flowServiceCreateFlowByBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *flowServiceCreateFlowByBidirectionalStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceCreateFlowByBidirectionalStreamServer) Recv() (*Flow, error) {
	m := new(Flow)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FlowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FlowService",
	HandlerType: (*FlowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFlow",
			Handler:    _FlowService_CreateFlow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateFlowResponseByServerStream",
			Handler:       _FlowService_CreateFlowResponseByServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateFlowRequestByClientStream",
			Handler:       _FlowService_CreateFlowRequestByClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateFlowByBidirectionalStream",
			Handler:       _FlowService_CreateFlowByBidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "flowCenter.proto",
}
