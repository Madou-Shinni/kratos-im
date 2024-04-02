// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.0
// source: gateway/gateway.proto

package gateway

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationGatewayGroupPutin = "/api.gateway.Gateway/GroupPutin"

type GatewayHTTPServer interface {
	GroupPutin(context.Context, *GroupPutinReq) (*GroupPutinResp, error)
}

func RegisterGatewayHTTPServer(s *http.Server, srv GatewayHTTPServer) {
	r := s.Route("/")
	r.PUT("/group/putin", _Gateway_GroupPutin0_HTTP_Handler(srv))
}

func _Gateway_GroupPutin0_HTTP_Handler(srv GatewayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupPutinReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGatewayGroupPutin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupPutin(ctx, req.(*GroupPutinReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GroupPutinResp)
		return ctx.Result(200, reply)
	}
}

type GatewayHTTPClient interface {
	GroupPutin(ctx context.Context, req *GroupPutinReq, opts ...http.CallOption) (rsp *GroupPutinResp, err error)
}

type GatewayHTTPClientImpl struct {
	cc *http.Client
}

func NewGatewayHTTPClient(client *http.Client) GatewayHTTPClient {
	return &GatewayHTTPClientImpl{client}
}

func (c *GatewayHTTPClientImpl) GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...http.CallOption) (*GroupPutinResp, error) {
	var out GroupPutinResp
	pattern := "/group/putin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGatewayGroupPutin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
