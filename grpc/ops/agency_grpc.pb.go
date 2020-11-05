// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ops

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AgencyClient is the client API for Agency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgencyClient interface {
	PSMHook(ctx context.Context, in *DataHook, opts ...grpc.CallOption) (Agency_PSMHookClient, error)
}

type agencyClient struct {
	cc grpc.ClientConnInterface
}

func NewAgencyClient(cc grpc.ClientConnInterface) AgencyClient {
	return &agencyClient{cc}
}

func (c *agencyClient) PSMHook(ctx context.Context, in *DataHook, opts ...grpc.CallOption) (Agency_PSMHookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Agency_serviceDesc.Streams[0], "/ops.Agency/PSMHook", opts...)
	if err != nil {
		return nil, err
	}
	x := &agencyPSMHookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Agency_PSMHookClient interface {
	Recv() (*AgencyStatus, error)
	grpc.ClientStream
}

type agencyPSMHookClient struct {
	grpc.ClientStream
}

func (x *agencyPSMHookClient) Recv() (*AgencyStatus, error) {
	m := new(AgencyStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgencyServer is the server API for Agency service.
// All implementations must embed UnimplementedAgencyServer
// for forward compatibility
type AgencyServer interface {
	PSMHook(*DataHook, Agency_PSMHookServer) error
	mustEmbedUnimplementedAgencyServer()
}

// UnimplementedAgencyServer must be embedded to have forward compatible implementations.
type UnimplementedAgencyServer struct {
}

func (UnimplementedAgencyServer) PSMHook(*DataHook, Agency_PSMHookServer) error {
	return status.Errorf(codes.Unimplemented, "method PSMHook not implemented")
}
func (UnimplementedAgencyServer) mustEmbedUnimplementedAgencyServer() {}

// UnsafeAgencyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgencyServer will
// result in compilation errors.
type UnsafeAgencyServer interface {
	mustEmbedUnimplementedAgencyServer()
}

func RegisterAgencyServer(s *grpc.Server, srv AgencyServer) {
	s.RegisterService(&_Agency_serviceDesc, srv)
}

func _Agency_PSMHook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataHook)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgencyServer).PSMHook(m, &agencyPSMHookServer{stream})
}

type Agency_PSMHookServer interface {
	Send(*AgencyStatus) error
	grpc.ServerStream
}

type agencyPSMHookServer struct {
	grpc.ServerStream
}

func (x *agencyPSMHookServer) Send(m *AgencyStatus) error {
	return x.ServerStream.SendMsg(m)
}

var _Agency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ops.Agency",
	HandlerType: (*AgencyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PSMHook",
			Handler:       _Agency_PSMHook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "agency.proto",
}

// DevOpsClient is the client API for DevOps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevOpsClient interface {
	Enter(ctx context.Context, in *Cmd, opts ...grpc.CallOption) (*CmdReturn, error)
}

type devOpsClient struct {
	cc grpc.ClientConnInterface
}

func NewDevOpsClient(cc grpc.ClientConnInterface) DevOpsClient {
	return &devOpsClient{cc}
}

func (c *devOpsClient) Enter(ctx context.Context, in *Cmd, opts ...grpc.CallOption) (*CmdReturn, error) {
	out := new(CmdReturn)
	err := c.cc.Invoke(ctx, "/ops.DevOps/Enter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevOpsServer is the server API for DevOps service.
// All implementations must embed UnimplementedDevOpsServer
// for forward compatibility
type DevOpsServer interface {
	Enter(context.Context, *Cmd) (*CmdReturn, error)
	mustEmbedUnimplementedDevOpsServer()
}

// UnimplementedDevOpsServer must be embedded to have forward compatible implementations.
type UnimplementedDevOpsServer struct {
}

func (UnimplementedDevOpsServer) Enter(context.Context, *Cmd) (*CmdReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enter not implemented")
}
func (UnimplementedDevOpsServer) mustEmbedUnimplementedDevOpsServer() {}

// UnsafeDevOpsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevOpsServer will
// result in compilation errors.
type UnsafeDevOpsServer interface {
	mustEmbedUnimplementedDevOpsServer()
}

func RegisterDevOpsServer(s *grpc.Server, srv DevOpsServer) {
	s.RegisterService(&_DevOps_serviceDesc, srv)
}

func _DevOps_Enter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevOpsServer).Enter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ops.DevOps/Enter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevOpsServer).Enter(ctx, req.(*Cmd))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevOps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ops.DevOps",
	HandlerType: (*DevOpsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enter",
			Handler:    _DevOps_Enter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agency.proto",
}
