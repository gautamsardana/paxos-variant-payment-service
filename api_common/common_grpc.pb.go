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
	Paxos_ProcessTxnSet_FullMethodName    = "/common.Paxos/ProcessTxnSet"
	Paxos_EnqueueTxn_FullMethodName       = "/common.Paxos/EnqueueTxn"
	Paxos_Prepare_FullMethodName          = "/common.Paxos/Prepare"
	Paxos_Promise_FullMethodName          = "/common.Paxos/Promise"
	Paxos_Accept_FullMethodName           = "/common.Paxos/Accept"
	Paxos_Accepted_FullMethodName         = "/common.Paxos/Accepted"
	Paxos_Commit_FullMethodName           = "/common.Paxos/Commit"
	Paxos_Sync_FullMethodName             = "/common.Paxos/Sync"
	Paxos_IsAlive_FullMethodName          = "/common.Paxos/IsAlive"
	Paxos_PrintBalance_FullMethodName     = "/common.Paxos/PrintBalance"
	Paxos_GetServerBalance_FullMethodName = "/common.Paxos/GetServerBalance"
	Paxos_PrintLogs_FullMethodName        = "/common.Paxos/PrintLogs"
	Paxos_PrintDB_FullMethodName          = "/common.Paxos/PrintDB"
)

// PaxosClient is the client API for Paxos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaxosClient interface {
	ProcessTxnSet(ctx context.Context, in *TxnSet, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EnqueueTxn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Prepare(ctx context.Context, in *Prepare, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Promise(ctx context.Context, in *Promise, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Accept(ctx context.Context, in *Accept, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Accepted(ctx context.Context, in *Accepted, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Commit(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	IsAlive(ctx context.Context, in *IsAliveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PrintBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetServerBalance(ctx context.Context, in *GetServerBalanceRequest, opts ...grpc.CallOption) (*GetServerBalanceResponse, error)
	PrintLogs(ctx context.Context, in *PrintLogsRequest, opts ...grpc.CallOption) (*PrintLogsResponse, error)
	PrintDB(ctx context.Context, in *PrintDBRequest, opts ...grpc.CallOption) (*PrintDBResponse, error)
}

type paxosClient struct {
	cc grpc.ClientConnInterface
}

func NewPaxosClient(cc grpc.ClientConnInterface) PaxosClient {
	return &paxosClient{cc}
}

func (c *paxosClient) ProcessTxnSet(ctx context.Context, in *TxnSet, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_ProcessTxnSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) EnqueueTxn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_EnqueueTxn_FullMethodName, in, out, cOpts...)
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

func (c *paxosClient) Accept(ctx context.Context, in *Accept, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_Accept_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) Accepted(ctx context.Context, in *Accepted, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_Accepted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) Commit(ctx context.Context, in *Commit, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_Commit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_Sync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) IsAlive(ctx context.Context, in *IsAliveRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Paxos_IsAlive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) PrintBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, Paxos_PrintBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) GetServerBalance(ctx context.Context, in *GetServerBalanceRequest, opts ...grpc.CallOption) (*GetServerBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServerBalanceResponse)
	err := c.cc.Invoke(ctx, Paxos_GetServerBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) PrintLogs(ctx context.Context, in *PrintLogsRequest, opts ...grpc.CallOption) (*PrintLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrintLogsResponse)
	err := c.cc.Invoke(ctx, Paxos_PrintLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosClient) PrintDB(ctx context.Context, in *PrintDBRequest, opts ...grpc.CallOption) (*PrintDBResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrintDBResponse)
	err := c.cc.Invoke(ctx, Paxos_PrintDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaxosServer is the server API for Paxos service.
// All implementations must embed UnimplementedPaxosServer
// for forward compatibility.
type PaxosServer interface {
	ProcessTxnSet(context.Context, *TxnSet) (*emptypb.Empty, error)
	EnqueueTxn(context.Context, *TxnRequest) (*emptypb.Empty, error)
	Prepare(context.Context, *Prepare) (*emptypb.Empty, error)
	Promise(context.Context, *Promise) (*emptypb.Empty, error)
	Accept(context.Context, *Accept) (*emptypb.Empty, error)
	Accepted(context.Context, *Accepted) (*emptypb.Empty, error)
	Commit(context.Context, *Commit) (*emptypb.Empty, error)
	Sync(context.Context, *SyncRequest) (*emptypb.Empty, error)
	IsAlive(context.Context, *IsAliveRequest) (*emptypb.Empty, error)
	PrintBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetServerBalance(context.Context, *GetServerBalanceRequest) (*GetServerBalanceResponse, error)
	PrintLogs(context.Context, *PrintLogsRequest) (*PrintLogsResponse, error)
	PrintDB(context.Context, *PrintDBRequest) (*PrintDBResponse, error)
	mustEmbedUnimplementedPaxosServer()
}

// UnimplementedPaxosServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaxosServer struct{}

func (UnimplementedPaxosServer) ProcessTxnSet(context.Context, *TxnSet) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessTxnSet not implemented")
}
func (UnimplementedPaxosServer) EnqueueTxn(context.Context, *TxnRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueTxn not implemented")
}
func (UnimplementedPaxosServer) Prepare(context.Context, *Prepare) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedPaxosServer) Promise(context.Context, *Promise) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Promise not implemented")
}
func (UnimplementedPaxosServer) Accept(context.Context, *Accept) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (UnimplementedPaxosServer) Accepted(context.Context, *Accepted) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accepted not implemented")
}
func (UnimplementedPaxosServer) Commit(context.Context, *Commit) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedPaxosServer) Sync(context.Context, *SyncRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedPaxosServer) IsAlive(context.Context, *IsAliveRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAlive not implemented")
}
func (UnimplementedPaxosServer) PrintBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintBalance not implemented")
}
func (UnimplementedPaxosServer) GetServerBalance(context.Context, *GetServerBalanceRequest) (*GetServerBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerBalance not implemented")
}
func (UnimplementedPaxosServer) PrintLogs(context.Context, *PrintLogsRequest) (*PrintLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintLogs not implemented")
}
func (UnimplementedPaxosServer) PrintDB(context.Context, *PrintDBRequest) (*PrintDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintDB not implemented")
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

func _Paxos_ProcessTxnSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).ProcessTxnSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_ProcessTxnSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).ProcessTxnSet(ctx, req.(*TxnSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_EnqueueTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).EnqueueTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_EnqueueTxn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).EnqueueTxn(ctx, req.(*TxnRequest))
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

func _Paxos_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Accept)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_Accept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).Accept(ctx, req.(*Accept))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_Accepted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Accepted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).Accepted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_Accepted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).Accepted(ctx, req.(*Accepted))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_Commit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).Commit(ctx, req.(*Commit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_IsAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).IsAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_IsAlive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).IsAlive(ctx, req.(*IsAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_PrintBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).PrintBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_PrintBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).PrintBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_GetServerBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).GetServerBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_GetServerBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).GetServerBalance(ctx, req.(*GetServerBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_PrintLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).PrintLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_PrintLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).PrintLogs(ctx, req.(*PrintLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paxos_PrintDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServer).PrintDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Paxos_PrintDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServer).PrintDB(ctx, req.(*PrintDBRequest))
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
			MethodName: "ProcessTxnSet",
			Handler:    _Paxos_ProcessTxnSet_Handler,
		},
		{
			MethodName: "EnqueueTxn",
			Handler:    _Paxos_EnqueueTxn_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Paxos_Prepare_Handler,
		},
		{
			MethodName: "Promise",
			Handler:    _Paxos_Promise_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _Paxos_Accept_Handler,
		},
		{
			MethodName: "Accepted",
			Handler:    _Paxos_Accepted_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Paxos_Commit_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Paxos_Sync_Handler,
		},
		{
			MethodName: "IsAlive",
			Handler:    _Paxos_IsAlive_Handler,
		},
		{
			MethodName: "PrintBalance",
			Handler:    _Paxos_PrintBalance_Handler,
		},
		{
			MethodName: "GetServerBalance",
			Handler:    _Paxos_GetServerBalance_Handler,
		},
		{
			MethodName: "PrintLogs",
			Handler:    _Paxos_PrintLogs_Handler,
		},
		{
			MethodName: "PrintDB",
			Handler:    _Paxos_PrintDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}
