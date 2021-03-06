package protobuf

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Channel  string `protobuf:"bytes,1,opt,name=Channel" json:"Channel,omitempty"`
	Metadata string `protobuf:"bytes,2,opt,name=Metadata" json:"Metadata,omitempty"`
	Body     []byte `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Message) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SubscribeRequest struct {
	Channel string `protobuf:"bytes,1,opt,name=Channel" json:"Channel,omitempty"`
	Group   string `protobuf:"bytes,2,opt,name=Group" json:"Group,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SubscribeRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SubscribeRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type Request struct {
	ID           string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Channel      string `protobuf:"bytes,2,opt,name=Channel" json:"Channel,omitempty"`
	Metadata     string `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
	Body         []byte `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
	ReplyChannel string `protobuf:"bytes,5,opt,name=ReplyChannel" json:"ReplyChannel,omitempty"`
	Timeout      int32  `protobuf:"varint,6,opt,name=Timeout" json:"Timeout,omitempty"`
	CacheKey     string `protobuf:"bytes,7,opt,name=CacheKey" json:"CacheKey,omitempty"`
	CacheTTL     int32  `protobuf:"varint,8,opt,name=CacheTTL" json:"CacheTTL,omitempty"`
	Context      []byte `protobuf:"bytes,9,opt,name=Context,proto3" json:"Context,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Request) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Request) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Request) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetReplyChannel() string {
	if m != nil {
		return m.ReplyChannel
	}
	return ""
}

func (m *Request) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Request) GetCacheKey() string {
	if m != nil {
		return m.CacheKey
	}
	return ""
}

func (m *Request) GetCacheTTL() int32 {
	if m != nil {
		return m.CacheTTL
	}
	return 0
}

func (m *Request) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

type Response struct {
	RequestID    string `protobuf:"bytes,1,opt,name=RequestID" json:"RequestID,omitempty"`
	ReplyChannel string `protobuf:"bytes,2,opt,name=ReplyChannel" json:"ReplyChannel,omitempty"`
	Metadata     string `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
	Body         []byte `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
	CacheHit     bool   `protobuf:"varint,5,opt,name=CacheHit" json:"CacheHit,omitempty"`
	Context      []byte `protobuf:"bytes,6,opt,name=Context,proto3" json:"Context,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Response) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *Response) GetReplyChannel() string {
	if m != nil {
		return m.ReplyChannel
	}
	return ""
}

func (m *Response) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetCacheHit() bool {
	if m != nil {
		return m.CacheHit
	}
	return false
}

func (m *Response) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "warp.Empty")
	proto.RegisterType((*Message)(nil), "warp.Message")
	proto.RegisterType((*SubscribeRequest)(nil), "warp.SubscribeRequest")
	proto.RegisterType((*Request)(nil), "warp.Request")
	proto.RegisterType((*Response)(nil), "warp.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Warp service

type WarpClient interface {
	// SendMessage - publish single message
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	// message_timeout - Optional, set timeout in millisecnd for verification of delivery
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
	// SendMessageStream - publish constant stream of pub Message
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SendMessageStream(ctx context.Context, opts ...grpc.CallOption) (Warp_SendMessageStreamClient, error)
	// SubscribeToChannel - listening to pub messages
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SubscribeToChannel(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Warp_SubscribeToChannelClient, error)
	// SendRequest - sending request with timeout
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// SendResponse - sending single response in case a client cannot support bi-di streaming
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SendResponse(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Empty, error)
	// RequestResponseStream - bi-di streams of getting request / sending replies
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	// channel - the channel we subscribe for getting requests
	// group - the gropu we subscirbe for getting requests
	RequestResponseStream(ctx context.Context, opts ...grpc.CallOption) (Warp_RequestResponseStreamClient, error)
}

type warpClient struct {
	cc *grpc.ClientConn
}

func NewWarpClient(cc *grpc.ClientConn) WarpClient {
	return &warpClient{cc}
}

func (c *warpClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/warp.warp/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warpClient) SendMessageStream(ctx context.Context, opts ...grpc.CallOption) (Warp_SendMessageStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Warp_serviceDesc.Streams[0], c.cc, "/warp.warp/SendMessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &warpSendMessageStreamClient{stream}
	return x, nil
}

type Warp_SendMessageStreamClient interface {
	Send(*Message) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type warpSendMessageStreamClient struct {
	grpc.ClientStream
}

func (x *warpSendMessageStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *warpSendMessageStreamClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *warpClient) SubscribeToChannel(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Warp_SubscribeToChannelClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Warp_serviceDesc.Streams[1], c.cc, "/warp.warp/SubscribeToChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &warpSubscribeToChannelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Warp_SubscribeToChannelClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type warpSubscribeToChannelClient struct {
	grpc.ClientStream
}

func (x *warpSubscribeToChannelClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *warpClient) SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/warp.warp/SendRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warpClient) SendResponse(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/warp.warp/SendResponse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warpClient) RequestResponseStream(ctx context.Context, opts ...grpc.CallOption) (Warp_RequestResponseStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Warp_serviceDesc.Streams[2], c.cc, "/warp.warp/RequestResponseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &warpRequestResponseStreamClient{stream}
	return x, nil
}

type Warp_RequestResponseStreamClient interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ClientStream
}

type warpRequestResponseStreamClient struct {
	grpc.ClientStream
}

func (x *warpRequestResponseStreamClient) Send(m *Response) error {
	return x.ClientStream.SendMsg(m)
}

func (x *warpRequestResponseStreamClient) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Warp service

type WarpServer interface {
	// SendMessage - publish single message
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	// message_timeout - Optional, set timeout in millisecnd for verification of delivery
	SendMessage(context.Context, *Message) (*Empty, error)
	// SendMessageStream - publish constant stream of pub Message
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SendMessageStream(Warp_SendMessageStreamServer) error
	// SubscribeToChannel - listening to pub messages
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SubscribeToChannel(*SubscribeRequest, Warp_SubscribeToChannelServer) error
	// SendRequest - sending request with timeout
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SendRequest(context.Context, *Request) (*Response, error)
	// SendResponse - sending single response in case a client cannot support bi-di streaming
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	SendResponse(context.Context, *Response) (*Empty, error)
	// RequestResponseStream - bi-di streams of getting request / sending replies
	// Metadata Paramters:
	// client_tag - a string that represent the client connection
	// channel - the channel we subscribe for getting requests
	// group - the gropu we subscirbe for getting requests
	RequestResponseStream(Warp_RequestResponseStreamServer) error
}

func RegisterWarpServer(s *grpc.Server, srv WarpServer) {
	s.RegisterService(&_Warp_serviceDesc, srv)
}

func _Warp_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarpServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warp.warp/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarpServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warp_SendMessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WarpServer).SendMessageStream(&warpSendMessageStreamServer{stream})
}

type Warp_SendMessageStreamServer interface {
	SendAndClose(*Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type warpSendMessageStreamServer struct {
	grpc.ServerStream
}

func (x *warpSendMessageStreamServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *warpSendMessageStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Warp_SubscribeToChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WarpServer).SubscribeToChannel(m, &warpSubscribeToChannelServer{stream})
}

type Warp_SubscribeToChannelServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type warpSubscribeToChannelServer struct {
	grpc.ServerStream
}

func (x *warpSubscribeToChannelServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Warp_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarpServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warp.warp/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarpServer).SendRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warp_SendResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarpServer).SendResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warp.warp/SendResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarpServer).SendResponse(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warp_RequestResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WarpServer).RequestResponseStream(&warpRequestResponseStreamServer{stream})
}

type Warp_RequestResponseStreamServer interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ServerStream
}

type warpRequestResponseStreamServer struct {
	grpc.ServerStream
}

func (x *warpRequestResponseStreamServer) Send(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

func (x *warpRequestResponseStreamServer) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Warp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "warp.warp",
	HandlerType: (*WarpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Warp_SendMessage_Handler,
		},
		{
			MethodName: "SendRequest",
			Handler:    _Warp_SendRequest_Handler,
		},
		{
			MethodName: "SendResponse",
			Handler:    _Warp_SendResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessageStream",
			Handler:       _Warp_SendMessageStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SubscribeToChannel",
			Handler:       _Warp_SubscribeToChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RequestResponseStream",
			Handler:       _Warp_RequestResponseStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "warp.proto",
}

func init() { proto.RegisterFile("warp.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0xba, 0x49, 0x9c, 0x4c, 0x43, 0x05, 0x0b, 0x54, 0xab, 0x88, 0x43, 0xb4, 0xa7, 0x48,
	0x08, 0xab, 0xa2, 0x57, 0x24, 0xa4, 0xb4, 0xa8, 0x54, 0x50, 0x0e, 0xb6, 0x5f, 0x60, 0x63, 0x8f,
	0xa8, 0xa5, 0xc6, 0x6b, 0xec, 0x35, 0xe0, 0x57, 0xe2, 0xc8, 0x43, 0xf0, 0x42, 0xbc, 0x00, 0xf2,
	0xfe, 0xb8, 0xb1, 0xa1, 0x91, 0xb8, 0xf9, 0xdb, 0x9d, 0xef, 0x9b, 0x6f, 0x66, 0x3f, 0x03, 0x7c,
	0x13, 0x65, 0x11, 0x14, 0xa5, 0x54, 0x92, 0x8e, 0xdb, 0x6f, 0xee, 0xc3, 0xe4, 0xdd, 0xae, 0x50,
	0x0d, 0x8f, 0xc0, 0xbf, 0xc1, 0xaa, 0x12, 0x9f, 0x91, 0x32, 0xf0, 0x2f, 0x6e, 0x45, 0x9e, 0xe3,
	0x1d, 0x23, 0x2b, 0xb2, 0x9e, 0x87, 0x0e, 0xd2, 0x25, 0xcc, 0x6e, 0x50, 0x89, 0x54, 0x28, 0xc1,
	0x3c, 0x7d, 0xd5, 0x61, 0x4a, 0x61, 0xbc, 0x91, 0x69, 0xc3, 0x8e, 0x56, 0x64, 0xbd, 0x08, 0xf5,
	0x37, 0xdf, 0xc0, 0xe3, 0xa8, 0xde, 0x56, 0x49, 0x99, 0x6d, 0x31, 0xc4, 0x2f, 0x35, 0x56, 0xea,
	0x80, 0xfa, 0x33, 0x98, 0x5c, 0x95, 0xb2, 0x2e, 0xac, 0xb4, 0x01, 0xfc, 0x37, 0x01, 0xdf, 0x71,
	0x4f, 0xc0, 0xbb, 0xbe, 0xb4, 0x34, 0xef, 0xfa, 0x72, 0x5f, 0xcb, 0x7b, 0xd8, 0xe9, 0xd1, 0x03,
	0x4e, 0xc7, 0xf7, 0x4e, 0x29, 0x87, 0x45, 0x88, 0xc5, 0x5d, 0xe3, 0xe4, 0x26, 0x9a, 0xd3, 0x3b,
	0x6b, 0xbb, 0xc5, 0xd9, 0x0e, 0x65, 0xad, 0xd8, 0x74, 0x45, 0xd6, 0x93, 0xd0, 0xc1, 0xb6, 0xdb,
	0x85, 0x48, 0x6e, 0xf1, 0x03, 0x36, 0xcc, 0x37, 0xdd, 0x1c, 0xee, 0xee, 0xe2, 0xf8, 0x23, 0x9b,
	0x69, 0x5a, 0x87, 0xb5, 0x7f, 0x99, 0x2b, 0xfc, 0xae, 0xd8, 0x5c, 0x9b, 0x71, 0x90, 0xff, 0x24,
	0x30, 0x0b, 0xb1, 0x2a, 0x64, 0x5e, 0x21, 0x7d, 0x01, 0x73, 0xbb, 0x81, 0x6e, 0xfa, 0xfb, 0x83,
	0xbf, 0xac, 0x7b, 0xff, 0xb0, 0xfe, 0xbf, 0xeb, 0x70, 0xa6, 0xdf, 0x67, 0x4a, 0xaf, 0x62, 0x16,
	0x76, 0x78, 0xdf, 0xf4, 0xb4, 0x67, 0xfa, 0xf5, 0x2f, 0x0f, 0x74, 0xaa, 0xe8, 0x4b, 0x38, 0x8e,
	0x30, 0x4f, 0x5d, 0xa0, 0x1e, 0x05, 0x3a, 0x77, 0x16, 0x2e, 0x8f, 0x0d, 0x34, 0xb9, 0x1b, 0xd1,
	0x73, 0x78, 0xb2, 0x57, 0x1c, 0xa9, 0x12, 0xc5, 0xee, 0x30, 0x65, 0x4d, 0xe8, 0x5b, 0xa0, 0x5d,
	0xb2, 0x62, 0xe9, 0xc6, 0x3c, 0x35, 0x65, 0xc3, 0xcc, 0x2d, 0xfb, 0x6a, 0x7c, 0x74, 0x46, 0x68,
	0x60, 0x2c, 0xba, 0x64, 0xd9, 0x0a, 0x47, 0x38, 0x71, 0xd0, 0xbc, 0x00, 0x1f, 0xd1, 0x57, 0xb0,
	0x30, 0xf5, 0xf6, 0x4d, 0x06, 0x15, 0xc3, 0xa1, 0xde, 0xc0, 0x73, 0xab, 0xe5, 0x2a, 0xec, 0x60,
	0x43, 0x5e, 0xbf, 0x71, 0x3b, 0xdb, 0x19, 0xd9, 0x9c, 0xfe, 0xf0, 0x9e, 0xc6, 0xa5, 0x48, 0x31,
	0x4f, 0x9a, 0xe0, 0x93, 0xf8, 0x9a, 0xc9, 0xe0, 0xaa, 0x2c, 0x92, 0xed, 0x54, 0xff, 0xba, 0xe7,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x6f, 0xf8, 0xd6, 0xc8, 0x03, 0x00, 0x00,
}
