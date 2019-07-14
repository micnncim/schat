// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chat.proto

package chatpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SendMessageRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{0}
}

func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SendMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendMessageResponse struct {
	Message              string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SendMessageResponse) Reset()         { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()    {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{1}
}

func (m *SendMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponse.Unmarshal(m, b)
}
func (m *SendMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponse.Marshal(b, m, deterministic)
}
func (m *SendMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponse.Merge(m, src)
}
func (m *SendMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponse.Size(m)
}
func (m *SendMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResponse proto.InternalMessageInfo

func (m *SendMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendMessageResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*SendMessageRequest)(nil), "chat.SendMessageRequest")
	proto.RegisterType((*SendMessageResponse)(nil), "chat.SendMessageResponse")
}

func init() { proto.RegisterFile("proto/chat.proto", fileDescriptor_ed7e7dde45555b7d) }

var fileDescriptor_ed7e7dde45555b7d = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0x48, 0x2c, 0xd1, 0x03, 0x33, 0x85, 0x58, 0x40, 0x6c, 0x29, 0xf9, 0xf4,
	0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0xb0, 0x58, 0x52, 0x69, 0x9a, 0x7e, 0x49, 0x66, 0x6e, 0x6a,
	0x71, 0x49, 0x62, 0x6e, 0x01, 0x44, 0x99, 0x92, 0x17, 0x97, 0x50, 0x70, 0x6a, 0x5e, 0x8a, 0x6f,
	0x6a, 0x71, 0x71, 0x62, 0x7a, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x14, 0x17,
	0x47, 0x69, 0x71, 0x6a, 0x51, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x9c, 0x2f, 0x24, 0xc1, 0xc5, 0x9e, 0x0b, 0x51, 0x2d, 0xc1, 0x04, 0x96, 0x82, 0x71, 0x95, 0x32,
	0xb9, 0x84, 0x51, 0xcc, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x46, 0xd1, 0xc0, 0x88, 0xa2, 0x41, 0xc8,
	0x82, 0x8b, 0x13, 0xee, 0x1e, 0xb0, 0x61, 0xdc, 0x46, 0x52, 0x7a, 0x10, 0x17, 0xeb, 0xc1, 0x5c,
	0xac, 0x17, 0x02, 0x53, 0x11, 0x84, 0x50, 0x6c, 0x14, 0xc0, 0xc5, 0xe2, 0x9c, 0x91, 0x58, 0x22,
	0xe4, 0xc1, 0xc5, 0x8d, 0x64, 0xa5, 0x90, 0x84, 0x1e, 0x38, 0x04, 0x30, 0x7d, 0x24, 0x25, 0x89,
	0x45, 0x06, 0xe2, 0x3e, 0x25, 0x06, 0x0d, 0x46, 0x03, 0x46, 0x27, 0x8e, 0x28, 0x36, 0x90, 0x8a,
	0x82, 0xa4, 0x24, 0x36, 0xb0, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x93, 0xec,
	0xf8, 0x54, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/chat.Chat/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatSendMessageClient{stream}
	return x, nil
}

type Chat_SendMessageClient interface {
	Send(*SendMessageRequest) error
	Recv() (*SendMessageResponse, error)
	grpc.ClientStream
}

type chatSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatSendMessageClient) Send(m *SendMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatSendMessageClient) Recv() (*SendMessageResponse, error) {
	m := new(SendMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	SendMessage(Chat_SendMessageServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) SendMessage(srv Chat_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).SendMessage(&chatSendMessageServer{stream})
}

type Chat_SendMessageServer interface {
	Send(*SendMessageResponse) error
	Recv() (*SendMessageRequest, error)
	grpc.ServerStream
}

type chatSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatSendMessageServer) Send(m *SendMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatSendMessageServer) Recv() (*SendMessageRequest, error) {
	m := new(SendMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _Chat_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chat.proto",
}