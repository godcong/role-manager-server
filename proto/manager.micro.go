// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: manager.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	manager.proto

It has these top-level messages:
	ManagerNodeRequest
	ManagerCensorRequest
	ManagerReply
	ManagerReplyDetail
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ManagerService service

type ManagerService interface {
	NodeBack(ctx context.Context, in *ManagerNodeRequest, opts ...client.CallOption) (*ManagerReply, error)
	CensorBack(ctx context.Context, in *ManagerCensorRequest, opts ...client.CallOption) (*ManagerReply, error)
}

type managerService struct {
	c    client.Client
	name string
}

func NewManagerService(name string, c client.Client) ManagerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &managerService{
		c:    c,
		name: name,
	}
}

func (c *managerService) NodeBack(ctx context.Context, in *ManagerNodeRequest, opts ...client.CallOption) (*ManagerReply, error) {
	req := c.c.NewRequest(c.name, "ManagerService.NodeBack", in)
	out := new(ManagerReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerService) CensorBack(ctx context.Context, in *ManagerCensorRequest, opts ...client.CallOption) (*ManagerReply, error) {
	req := c.c.NewRequest(c.name, "ManagerService.CensorBack", in)
	out := new(ManagerReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ManagerService service

type ManagerServiceHandler interface {
	NodeBack(context.Context, *ManagerNodeRequest, *ManagerReply) error
	CensorBack(context.Context, *ManagerCensorRequest, *ManagerReply) error
}

func RegisterManagerServiceHandler(s server.Server, hdlr ManagerServiceHandler, opts ...server.HandlerOption) error {
	type managerService interface {
		NodeBack(ctx context.Context, in *ManagerNodeRequest, out *ManagerReply) error
		CensorBack(ctx context.Context, in *ManagerCensorRequest, out *ManagerReply) error
	}
	type ManagerService struct {
		managerService
	}
	h := &managerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ManagerService{h}, opts...))
}

type managerServiceHandler struct {
	ManagerServiceHandler
}

func (h *managerServiceHandler) NodeBack(ctx context.Context, in *ManagerNodeRequest, out *ManagerReply) error {
	return h.ManagerServiceHandler.NodeBack(ctx, in, out)
}

func (h *managerServiceHandler) CensorBack(ctx context.Context, in *ManagerCensorRequest, out *ManagerReply) error {
	return h.ManagerServiceHandler.CensorBack(ctx, in, out)
}