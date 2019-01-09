// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: consignment.proto

/*
Package cz_go_microservices_consignment is a generated protocol buffer package.

It is generated from these files:
	consignment.proto

It has these top-level messages:
	Consignment
	Container
	Response
	EmptyRequest
*/
package cz_go_microservices_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ConsignmentService service

type ConsignmentService interface {
	Create(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*Response, error)
}

type consignmentService struct {
	c    client.Client
	name string
}

func NewConsignmentService(name string, c client.Client) ConsignmentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "cz.go.microservices.consignment"
	}
	return &consignmentService{
		c:    c,
		name: name,
	}
}

func (c *consignmentService) Create(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConsignmentService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consignmentService) GetAll(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ConsignmentService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsignmentService service

type ConsignmentServiceHandler interface {
	Create(context.Context, *Consignment, *Response) error
	GetAll(context.Context, *EmptyRequest, *Response) error
}

func RegisterConsignmentServiceHandler(s server.Server, hdlr ConsignmentServiceHandler, opts ...server.HandlerOption) error {
	type consignmentService interface {
		Create(ctx context.Context, in *Consignment, out *Response) error
		GetAll(ctx context.Context, in *EmptyRequest, out *Response) error
	}
	type ConsignmentService struct {
		consignmentService
	}
	h := &consignmentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConsignmentService{h}, opts...))
}

type consignmentServiceHandler struct {
	ConsignmentServiceHandler
}

func (h *consignmentServiceHandler) Create(ctx context.Context, in *Consignment, out *Response) error {
	return h.ConsignmentServiceHandler.Create(ctx, in, out)
}

func (h *consignmentServiceHandler) GetAll(ctx context.Context, in *EmptyRequest, out *Response) error {
	return h.ConsignmentServiceHandler.GetAll(ctx, in, out)
}