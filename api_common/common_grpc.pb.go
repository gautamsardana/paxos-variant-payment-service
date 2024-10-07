// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: common.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Paxos_ProcessTxn_FullMethodName = "/common.Paxos/ProcessTxn"
	Paxos_Prepare_FullMethodName    = "/common.Paxos/Prepare"
	Paxos_Promise_FullMethodName    = "/common.Paxos/Promise"
)

// PaxosClient is the client API for Paxos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaxosClient interface {
	ProcessTxn(ctx context.Context, in *ProcessTxnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Prepare(ctx context.Context, in *Prepare, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Promise(ctx context.Context, in *Promise, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type paxosClient struct {
	cc grpc.ClientConnInterface
}

func NewPaxosClient(cc grpc.ClientConnInterface) PaxosClient {
	return &paxosClient{cc}
}

func (c *paxosClient) ProcessTxn(ctx context.Context, in *ProcessTxnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_ProcessTxn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) Prepare(ctx context.Context, in *Prepare, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_Prepare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) Promise(ctx context.Context, in *Promise, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_Promise_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaxosServer is the server API for Paxos service.
// All implementations must embed UnimplementedPaxosServer
// for forward compatibility.
type PaxosServer interface {
	ProcessTxn(context.Context, *ProcessTxnRequest) (*emptypb.Empty, error)
	Prepare(context.Context, *Prepare) (*emptypb.Empty, error)
	Promise(context.Context, *Promise) (*emptypb.Empty, error)
	mustEmbedUnimplementedPaxosServer()
}

// UnimplementedPaxosServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaxosServer struct{}

func (UnimplementedPaxosServer) ProcessTxn(context.Context, *ProcessTxnRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessTxn not implemented")
}
func (UnimplementedPaxosServer) Prepare(context.Context, *Prepare) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedPaxosServer) Promise(context.Context, *Promise) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Promise not implemented")
}
func (UnimplementedPaxosServer) mustEmbedUnimplementedPaxosServer() {}
func (UnimplementedPaxosServer) testEmbeddedByValue()               {}

// UnsafePaxosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaxosServer will
// result in compilation errors.
type UnsafePaxosServer interface {
	mustEmbedUnimplementedPaxosServer()
}

func RegisterPaxosServer(s grpc.ServiceRegistrar, srv PaxosServer) {
	// If the following call pancis, it indicates UnimplementedPaxosServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Paxos_ServiceDesc, srv)
}

func _Paxos_ProcessTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessTxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).ProcessTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_ProcessTxn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).ProcessTxn(ctx, req.(*ProcessTxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prepare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_Prepare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).Prepare(ctx, req.(*Prepare))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_Promise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Promise)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).Promise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_Promise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).Promise(ctx, req.(*Promise))
	}
	return interceptor(ctx, in, info, handler)
}

// Paxos_ServiceDesc is the grpc.ServiceDesc for Paxos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Paxos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.Paxos",
	HandlerType: (*PaxosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessTxn",
			Handler:    _Paxos_ProcessTxn_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Paxos_Prepare_Handler,
		},
		{
			MethodName: "Promise",
			Handler:    _Paxos_Promise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}