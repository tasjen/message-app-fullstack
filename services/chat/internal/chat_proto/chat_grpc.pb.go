// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: internal/chat_proto/chat.proto

package chat_proto

import (
	context "context"
	empty "google.golang.org/protobuf/types/known/emptypb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Chat_GetByUsername_FullMethodName = "/Chat/GetByUsername"
	Chat_GetUserById_FullMethodName   = "/Chat/GetUserById"
	Chat_GetContacts_FullMethodName   = "/Chat/GetContacts"
	Chat_GetMessages_FullMethodName   = "/Chat/GetMessages"
	Chat_CreateUser_FullMethodName    = "/Chat/CreateUser"
	Chat_CreateGroup_FullMethodName   = "/Chat/CreateGroup"
	Chat_AddFriend_FullMethodName     = "/Chat/AddFriend"
	Chat_AddMember_FullMethodName     = "/Chat/AddMember"
	Chat_SendMessage_FullMethodName   = "/Chat/SendMessage"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	GetByUsername(ctx context.Context, in *GetByUsernameReq, opts ...grpc.CallOption) (*GetByUsernameRes, error)
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdRes, error)
	GetContacts(ctx context.Context, in *GetContactsReq, opts ...grpc.CallOption) (*GetContactsRes, error)
	GetMessages(ctx context.Context, in *GetMessagesReq, opts ...grpc.CallOption) (*GetMessagesRes, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error)
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*empty.Empty, error)
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*empty.Empty, error)
	AddMember(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*empty.Empty, error)
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageRes, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) GetByUsername(ctx context.Context, in *GetByUsernameReq, opts ...grpc.CallOption) (*GetByUsernameRes, error) {
	out := new(GetByUsernameRes)
	err := c.cc.Invoke(ctx, Chat_GetByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdRes, error) {
	out := new(GetUserByIdRes)
	err := c.cc.Invoke(ctx, Chat_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetContacts(ctx context.Context, in *GetContactsReq, opts ...grpc.CallOption) (*GetContactsRes, error) {
	out := new(GetContactsRes)
	err := c.cc.Invoke(ctx, Chat_GetContacts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetMessages(ctx context.Context, in *GetMessagesReq, opts ...grpc.CallOption) (*GetMessagesRes, error) {
	out := new(GetMessagesRes)
	err := c.cc.Invoke(ctx, Chat_GetMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error) {
	out := new(CreateUserRes)
	err := c.cc.Invoke(ctx, Chat_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Chat_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Chat_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) AddMember(ctx context.Context, in *AddMemberReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Chat_AddMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageRes, error) {
	out := new(SendMessageRes)
	err := c.cc.Invoke(ctx, Chat_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	GetByUsername(context.Context, *GetByUsernameReq) (*GetByUsernameRes, error)
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdRes, error)
	GetContacts(context.Context, *GetContactsReq) (*GetContactsRes, error)
	GetMessages(context.Context, *GetMessagesReq) (*GetMessagesRes, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	CreateGroup(context.Context, *CreateGroupReq) (*empty.Empty, error)
	AddFriend(context.Context, *AddFriendReq) (*empty.Empty, error)
	AddMember(context.Context, *AddMemberReq) (*empty.Empty, error)
	SendMessage(context.Context, *SendMessageReq) (*SendMessageRes, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) GetByUsername(context.Context, *GetByUsernameReq) (*GetByUsernameRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (UnimplementedChatServer) GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedChatServer) GetContacts(context.Context, *GetContactsReq) (*GetContactsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedChatServer) GetMessages(context.Context, *GetMessagesReq) (*GetMessagesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedChatServer) CreateGroup(context.Context, *CreateGroupReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedChatServer) AddFriend(context.Context, *AddFriendReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedChatServer) AddMember(context.Context, *AddMemberReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (UnimplementedChatServer) SendMessage(context.Context, *SendMessageReq) (*SendMessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetByUsername(ctx, req.(*GetByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetContacts(ctx, req.(*GetContactsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetMessages(ctx, req.(*GetMessagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateGroup(ctx, req.(*CreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_AddMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).AddMember(ctx, req.(*AddMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByUsername",
			Handler:    _Chat_GetByUsername_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Chat_GetUserById_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _Chat_GetContacts_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Chat_GetMessages_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Chat_CreateUser_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _Chat_CreateGroup_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _Chat_AddFriend_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _Chat_AddMember_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Chat_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/chat_proto/chat.proto",
}
