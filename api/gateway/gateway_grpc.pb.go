// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: gateway/gateway.proto

package gateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	user "kratos-im/api/user"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Gateway_GroupCreate_FullMethodName           = "/api.gateway.Gateway/GroupCreate"
	Gateway_GroupPutin_FullMethodName            = "/api.gateway.Gateway/GroupPutin"
	Gateway_GroupPutinList_FullMethodName        = "/api.gateway.Gateway/GroupPutinList"
	Gateway_GroupPutInHandle_FullMethodName      = "/api.gateway.Gateway/GroupPutInHandle"
	Gateway_GroupList_FullMethodName             = "/api.gateway.Gateway/GroupList"
	Gateway_GroupUsers_FullMethodName            = "/api.gateway.Gateway/GroupUsers"
	Gateway_FriendPutIn_FullMethodName           = "/api.gateway.Gateway/FriendPutIn"
	Gateway_FriendPutInHandle_FullMethodName     = "/api.gateway.Gateway/FriendPutInHandle"
	Gateway_FriendPutInList_FullMethodName       = "/api.gateway.Gateway/FriendPutInList"
	Gateway_FriendList_FullMethodName            = "/api.gateway.Gateway/FriendList"
	Gateway_FriendsOnline_FullMethodName         = "/api.gateway.Gateway/FriendsOnline"
	Gateway_GroupMembersOnline_FullMethodName    = "/api.gateway.Gateway/GroupMembersOnline"
	Gateway_SetUpUserConversation_FullMethodName = "/api.gateway.Gateway/SetUpUserConversation"
	Gateway_GetConversations_FullMethodName      = "/api.gateway.Gateway/GetConversations"
	Gateway_PutConversations_FullMethodName      = "/api.gateway.Gateway/PutConversations"
	Gateway_GetChatLog_FullMethodName            = "/api.gateway.Gateway/GetChatLog"
	Gateway_GetReadChatRecords_FullMethodName    = "/api.gateway.Gateway/GetReadChatRecords"
	Gateway_UserLogin_FullMethodName             = "/api.gateway.Gateway/UserLogin"
	Gateway_UserSignUp_FullMethodName            = "/api.gateway.Gateway/UserSignUp"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	// 创建群
	GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error)
	// 入群申请
	GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error)
	// 入群申请列表
	GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error)
	// 入群申请处理
	GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error)
	// 群列表
	GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error)
	// 群成员列表
	GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error)
	// 好友申请
	FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error)
	// 好友申请处理
	FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error)
	// 好友申请列表
	FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error)
	// 好友列表
	FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error)
	// 在线好友情况
	FriendsOnline(ctx context.Context, in *FriendsOnlineReq, opts ...grpc.CallOption) (*FriendsOnlineResp, error)
	// 在线群成员情况
	GroupMembersOnline(ctx context.Context, in *GroupMembersOnlineReq, opts ...grpc.CallOption) (*GroupMembersOnlineResp, error)
	// 建立会话
	SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error)
	// 获取会话列表
	GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error)
	// 更新会话
	PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error)
	// 获取聊天记录
	GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error)
	// 获取消息已读记录
	GetReadChatRecords(ctx context.Context, in *GetReadChatRecordsReq, opts ...grpc.CallOption) (*GetReadChatRecordsResp, error)
	// 用户登录
	UserLogin(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*UserLoginResp, error)
	// 用户注册
	UserSignUp(ctx context.Context, in *user.Account, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error) {
	out := new(GroupCreateResp)
	err := c.cc.Invoke(ctx, Gateway_GroupCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error) {
	out := new(GroupPutinResp)
	err := c.cc.Invoke(ctx, Gateway_GroupPutin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error) {
	out := new(GroupPutinListResp)
	err := c.cc.Invoke(ctx, Gateway_GroupPutinList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error) {
	out := new(GroupPutInHandleResp)
	err := c.cc.Invoke(ctx, Gateway_GroupPutInHandle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error) {
	out := new(GroupListResp)
	err := c.cc.Invoke(ctx, Gateway_GroupList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error) {
	out := new(GroupUsersResp)
	err := c.cc.Invoke(ctx, Gateway_GroupUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error) {
	out := new(FriendPutInResp)
	err := c.cc.Invoke(ctx, Gateway_FriendPutIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error) {
	out := new(FriendPutInHandleResp)
	err := c.cc.Invoke(ctx, Gateway_FriendPutInHandle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error) {
	out := new(FriendPutInListResp)
	err := c.cc.Invoke(ctx, Gateway_FriendPutInList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error) {
	out := new(FriendListResp)
	err := c.cc.Invoke(ctx, Gateway_FriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) FriendsOnline(ctx context.Context, in *FriendsOnlineReq, opts ...grpc.CallOption) (*FriendsOnlineResp, error) {
	out := new(FriendsOnlineResp)
	err := c.cc.Invoke(ctx, Gateway_FriendsOnline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GroupMembersOnline(ctx context.Context, in *GroupMembersOnlineReq, opts ...grpc.CallOption) (*GroupMembersOnlineResp, error) {
	out := new(GroupMembersOnlineResp)
	err := c.cc.Invoke(ctx, Gateway_GroupMembersOnline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error) {
	out := new(SetUpUserConversationResp)
	err := c.cc.Invoke(ctx, Gateway_SetUpUserConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error) {
	out := new(GetConversationsResp)
	err := c.cc.Invoke(ctx, Gateway_GetConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error) {
	out := new(PutConversationsResp)
	err := c.cc.Invoke(ctx, Gateway_PutConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error) {
	out := new(GetChatLogResp)
	err := c.cc.Invoke(ctx, Gateway_GetChatLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetReadChatRecords(ctx context.Context, in *GetReadChatRecordsReq, opts ...grpc.CallOption) (*GetReadChatRecordsResp, error) {
	out := new(GetReadChatRecordsResp)
	err := c.cc.Invoke(ctx, Gateway_GetReadChatRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UserLogin(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*UserLoginResp, error) {
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, Gateway_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UserSignUp(ctx context.Context, in *user.Account, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Gateway_UserSignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	// 创建群
	GroupCreate(context.Context, *GroupCreateReq) (*GroupCreateResp, error)
	// 入群申请
	GroupPutin(context.Context, *GroupPutinReq) (*GroupPutinResp, error)
	// 入群申请列表
	GroupPutinList(context.Context, *GroupPutinListReq) (*GroupPutinListResp, error)
	// 入群申请处理
	GroupPutInHandle(context.Context, *GroupPutInHandleReq) (*GroupPutInHandleResp, error)
	// 群列表
	GroupList(context.Context, *GroupListReq) (*GroupListResp, error)
	// 群成员列表
	GroupUsers(context.Context, *GroupUsersReq) (*GroupUsersResp, error)
	// 好友申请
	FriendPutIn(context.Context, *FriendPutInReq) (*FriendPutInResp, error)
	// 好友申请处理
	FriendPutInHandle(context.Context, *FriendPutInHandleReq) (*FriendPutInHandleResp, error)
	// 好友申请列表
	FriendPutInList(context.Context, *FriendPutInListReq) (*FriendPutInListResp, error)
	// 好友列表
	FriendList(context.Context, *FriendListReq) (*FriendListResp, error)
	// 在线好友情况
	FriendsOnline(context.Context, *FriendsOnlineReq) (*FriendsOnlineResp, error)
	// 在线群成员情况
	GroupMembersOnline(context.Context, *GroupMembersOnlineReq) (*GroupMembersOnlineResp, error)
	// 建立会话
	SetUpUserConversation(context.Context, *SetUpUserConversationReq) (*SetUpUserConversationResp, error)
	// 获取会话列表
	GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error)
	// 更新会话
	PutConversations(context.Context, *PutConversationsReq) (*PutConversationsResp, error)
	// 获取聊天记录
	GetChatLog(context.Context, *GetChatLogReq) (*GetChatLogResp, error)
	// 获取消息已读记录
	GetReadChatRecords(context.Context, *GetReadChatRecordsReq) (*GetReadChatRecordsResp, error)
	// 用户登录
	UserLogin(context.Context, *user.LoginRequest) (*UserLoginResp, error)
	// 用户注册
	UserSignUp(context.Context, *user.Account) (*emptypb.Empty, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GroupCreate(context.Context, *GroupCreateReq) (*GroupCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupCreate not implemented")
}
func (UnimplementedGatewayServer) GroupPutin(context.Context, *GroupPutinReq) (*GroupPutinResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPutin not implemented")
}
func (UnimplementedGatewayServer) GroupPutinList(context.Context, *GroupPutinListReq) (*GroupPutinListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPutinList not implemented")
}
func (UnimplementedGatewayServer) GroupPutInHandle(context.Context, *GroupPutInHandleReq) (*GroupPutInHandleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPutInHandle not implemented")
}
func (UnimplementedGatewayServer) GroupList(context.Context, *GroupListReq) (*GroupListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupList not implemented")
}
func (UnimplementedGatewayServer) GroupUsers(context.Context, *GroupUsersReq) (*GroupUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupUsers not implemented")
}
func (UnimplementedGatewayServer) FriendPutIn(context.Context, *FriendPutInReq) (*FriendPutInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendPutIn not implemented")
}
func (UnimplementedGatewayServer) FriendPutInHandle(context.Context, *FriendPutInHandleReq) (*FriendPutInHandleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendPutInHandle not implemented")
}
func (UnimplementedGatewayServer) FriendPutInList(context.Context, *FriendPutInListReq) (*FriendPutInListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendPutInList not implemented")
}
func (UnimplementedGatewayServer) FriendList(context.Context, *FriendListReq) (*FriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendList not implemented")
}
func (UnimplementedGatewayServer) FriendsOnline(context.Context, *FriendsOnlineReq) (*FriendsOnlineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendsOnline not implemented")
}
func (UnimplementedGatewayServer) GroupMembersOnline(context.Context, *GroupMembersOnlineReq) (*GroupMembersOnlineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupMembersOnline not implemented")
}
func (UnimplementedGatewayServer) SetUpUserConversation(context.Context, *SetUpUserConversationReq) (*SetUpUserConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUpUserConversation not implemented")
}
func (UnimplementedGatewayServer) GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversations not implemented")
}
func (UnimplementedGatewayServer) PutConversations(context.Context, *PutConversationsReq) (*PutConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutConversations not implemented")
}
func (UnimplementedGatewayServer) GetChatLog(context.Context, *GetChatLogReq) (*GetChatLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatLog not implemented")
}
func (UnimplementedGatewayServer) GetReadChatRecords(context.Context, *GetReadChatRecordsReq) (*GetReadChatRecordsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadChatRecords not implemented")
}
func (UnimplementedGatewayServer) UserLogin(context.Context, *user.LoginRequest) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedGatewayServer) UserSignUp(context.Context, *user.Account) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUp not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_GroupCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GroupCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GroupCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GroupCreate(ctx, req.(*GroupCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GroupPutin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPutinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GroupPutin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GroupPutin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GroupPutin(ctx, req.(*GroupPutinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GroupPutinList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPutinListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GroupPutinList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GroupPutinList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GroupPutinList(ctx, req.(*GroupPutinListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GroupPutInHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPutInHandleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GroupPutInHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GroupPutInHandle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GroupPutInHandle(ctx, req.(*GroupPutInHandleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GroupList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GroupList(ctx, req.(*GroupListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GroupUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GroupUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GroupUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GroupUsers(ctx, req.(*GroupUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_FriendPutIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendPutInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).FriendPutIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_FriendPutIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).FriendPutIn(ctx, req.(*FriendPutInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_FriendPutInHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendPutInHandleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).FriendPutInHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_FriendPutInHandle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).FriendPutInHandle(ctx, req.(*FriendPutInHandleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_FriendPutInList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendPutInListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).FriendPutInList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_FriendPutInList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).FriendPutInList(ctx, req.(*FriendPutInListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_FriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).FriendList(ctx, req.(*FriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_FriendsOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendsOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).FriendsOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_FriendsOnline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).FriendsOnline(ctx, req.(*FriendsOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GroupMembersOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupMembersOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GroupMembersOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GroupMembersOnline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GroupMembersOnline(ctx, req.(*GroupMembersOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_SetUpUserConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUpUserConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).SetUpUserConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_SetUpUserConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).SetUpUserConversation(ctx, req.(*SetUpUserConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetConversations(ctx, req.(*GetConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_PutConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).PutConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_PutConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).PutConversations(ctx, req.(*PutConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetChatLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetChatLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetChatLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetChatLog(ctx, req.(*GetChatLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetReadChatRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadChatRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetReadChatRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetReadChatRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetReadChatRecords(ctx, req.(*GetReadChatRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UserLogin(ctx, req.(*user.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UserSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UserSignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UserSignUp(ctx, req.(*user.Account))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.gateway.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GroupCreate",
			Handler:    _Gateway_GroupCreate_Handler,
		},
		{
			MethodName: "GroupPutin",
			Handler:    _Gateway_GroupPutin_Handler,
		},
		{
			MethodName: "GroupPutinList",
			Handler:    _Gateway_GroupPutinList_Handler,
		},
		{
			MethodName: "GroupPutInHandle",
			Handler:    _Gateway_GroupPutInHandle_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Gateway_GroupList_Handler,
		},
		{
			MethodName: "GroupUsers",
			Handler:    _Gateway_GroupUsers_Handler,
		},
		{
			MethodName: "FriendPutIn",
			Handler:    _Gateway_FriendPutIn_Handler,
		},
		{
			MethodName: "FriendPutInHandle",
			Handler:    _Gateway_FriendPutInHandle_Handler,
		},
		{
			MethodName: "FriendPutInList",
			Handler:    _Gateway_FriendPutInList_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _Gateway_FriendList_Handler,
		},
		{
			MethodName: "FriendsOnline",
			Handler:    _Gateway_FriendsOnline_Handler,
		},
		{
			MethodName: "GroupMembersOnline",
			Handler:    _Gateway_GroupMembersOnline_Handler,
		},
		{
			MethodName: "SetUpUserConversation",
			Handler:    _Gateway_SetUpUserConversation_Handler,
		},
		{
			MethodName: "GetConversations",
			Handler:    _Gateway_GetConversations_Handler,
		},
		{
			MethodName: "PutConversations",
			Handler:    _Gateway_PutConversations_Handler,
		},
		{
			MethodName: "GetChatLog",
			Handler:    _Gateway_GetChatLog_Handler,
		},
		{
			MethodName: "GetReadChatRecords",
			Handler:    _Gateway_GetReadChatRecords_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _Gateway_UserLogin_Handler,
		},
		{
			MethodName: "UserSignUp",
			Handler:    _Gateway_UserSignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/gateway.proto",
}
