// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: im/im.proto

package im

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IM_CreateChatLog_FullMethodName           = "/api.im.IM/CreateChatLog"
	IM_GetChatLog_FullMethodName              = "/api.im.IM/GetChatLog"
	IM_SetUpUserConversation_FullMethodName   = "/api.im.IM/SetUpUserConversation"
	IM_GetConversations_FullMethodName        = "/api.im.IM/GetConversations"
	IM_PutConversations_FullMethodName        = "/api.im.IM/PutConversations"
	IM_CreateGroupConversation_FullMethodName = "/api.im.IM/CreateGroupConversation"
)

// IMClient is the client API for IM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IMClient interface {
	// 创建会话记录
	CreateChatLog(ctx context.Context, in *ChatLog, opts ...grpc.CallOption) (*CreateChatLogResp, error)
	// 获取会话记录
	GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error)
	// 建立会话: 群聊, 私聊
	SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error)
	// 获取会话
	GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error)
	// 更新会话
	PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error)
	CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error)
}

type iMClient struct {
	cc grpc.ClientConnInterface
}

func NewIMClient(cc grpc.ClientConnInterface) IMClient {
	return &iMClient{cc}
}

func (c *iMClient) CreateChatLog(ctx context.Context, in *ChatLog, opts ...grpc.CallOption) (*CreateChatLogResp, error) {
	out := new(CreateChatLogResp)
	err := c.cc.Invoke(ctx, IM_CreateChatLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error) {
	out := new(GetChatLogResp)
	err := c.cc.Invoke(ctx, IM_GetChatLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error) {
	out := new(SetUpUserConversationResp)
	err := c.cc.Invoke(ctx, IM_SetUpUserConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error) {
	out := new(GetConversationsResp)
	err := c.cc.Invoke(ctx, IM_GetConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error) {
	out := new(PutConversationsResp)
	err := c.cc.Invoke(ctx, IM_PutConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMClient) CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error) {
	out := new(CreateGroupConversationResp)
	err := c.cc.Invoke(ctx, IM_CreateGroupConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IMServer is the server API for IM service.
// All implementations must embed UnimplementedIMServer
// for forward compatibility
type IMServer interface {
	// 创建会话记录
	CreateChatLog(context.Context, *ChatLog) (*CreateChatLogResp, error)
	// 获取会话记录
	GetChatLog(context.Context, *GetChatLogReq) (*GetChatLogResp, error)
	// 建立会话: 群聊, 私聊
	SetUpUserConversation(context.Context, *SetUpUserConversationReq) (*SetUpUserConversationResp, error)
	// 获取会话
	GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error)
	// 更新会话
	PutConversations(context.Context, *PutConversationsReq) (*PutConversationsResp, error)
	CreateGroupConversation(context.Context, *CreateGroupConversationReq) (*CreateGroupConversationResp, error)
	mustEmbedUnimplementedIMServer()
}

// UnimplementedIMServer must be embedded to have forward compatible implementations.
type UnimplementedIMServer struct {
}

func (UnimplementedIMServer) CreateChatLog(context.Context, *ChatLog) (*CreateChatLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatLog not implemented")
}
func (UnimplementedIMServer) GetChatLog(context.Context, *GetChatLogReq) (*GetChatLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatLog not implemented")
}
func (UnimplementedIMServer) SetUpUserConversation(context.Context, *SetUpUserConversationReq) (*SetUpUserConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUpUserConversation not implemented")
}
func (UnimplementedIMServer) GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversations not implemented")
}
func (UnimplementedIMServer) PutConversations(context.Context, *PutConversationsReq) (*PutConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutConversations not implemented")
}
func (UnimplementedIMServer) CreateGroupConversation(context.Context, *CreateGroupConversationReq) (*CreateGroupConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupConversation not implemented")
}
func (UnimplementedIMServer) mustEmbedUnimplementedIMServer() {}

// UnsafeIMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IMServer will
// result in compilation errors.
type UnsafeIMServer interface {
	mustEmbedUnimplementedIMServer()
}

func RegisterIMServer(s grpc.ServiceRegistrar, srv IMServer) {
	s.RegisterService(&IM_ServiceDesc, srv)
}

func _IM_CreateChatLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServer).CreateChatLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IM_CreateChatLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServer).CreateChatLog(ctx, req.(*ChatLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _IM_GetChatLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServer).GetChatLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IM_GetChatLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServer).GetChatLog(ctx, req.(*GetChatLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IM_SetUpUserConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUpUserConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServer).SetUpUserConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IM_SetUpUserConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServer).SetUpUserConversation(ctx, req.(*SetUpUserConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IM_GetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServer).GetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IM_GetConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServer).GetConversations(ctx, req.(*GetConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IM_PutConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServer).PutConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IM_PutConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServer).PutConversations(ctx, req.(*PutConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IM_CreateGroupConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServer).CreateGroupConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IM_CreateGroupConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServer).CreateGroupConversation(ctx, req.(*CreateGroupConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IM_ServiceDesc is the grpc.ServiceDesc for IM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.im.IM",
	HandlerType: (*IMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChatLog",
			Handler:    _IM_CreateChatLog_Handler,
		},
		{
			MethodName: "GetChatLog",
			Handler:    _IM_GetChatLog_Handler,
		},
		{
			MethodName: "SetUpUserConversation",
			Handler:    _IM_SetUpUserConversation_Handler,
		},
		{
			MethodName: "GetConversations",
			Handler:    _IM_GetConversations_Handler,
		},
		{
			MethodName: "PutConversations",
			Handler:    _IM_PutConversations_Handler,
		},
		{
			MethodName: "CreateGroupConversation",
			Handler:    _IM_CreateGroupConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im/im.proto",
}
