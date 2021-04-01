// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package connector

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	common "github.com/telepresenceio/telepresence/rpc/v2/common"
	manager "github.com/telepresenceio/telepresence/rpc/v2/manager"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ConnectorClient is the client API for Connector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorClient interface {
	// Returns version information from the Connector
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.VersionInfo, error)
	// Connects to the cluster and connects the laptop's network (via
	// the daemon process) to the cluster's network.  A result code of
	// UNSPECIFIED indicates that the connection was successfully
	// initiated; if already connected, then either ALREADY_CONNECTED or
	// MUST_RESTART is returned, based on whether the current connection
	// is in agreement with the ConnectionRequest.
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectInfo, error)
	// Status is much like Connect, except that it doesn't actually do
	// anything.  It's a dry-run.
	Status(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectInfo, error)
	// Adds an intercept to a workload.  Requires having already called
	// Connect.
	CreateIntercept(ctx context.Context, in *CreateInterceptRequest, opts ...grpc.CallOption) (*InterceptResult, error)
	// Deactivates and removes an existent workload intercept.
	// Requires having already called Connect.
	RemoveIntercept(ctx context.Context, in *manager.RemoveInterceptRequest2, opts ...grpc.CallOption) (*InterceptResult, error)
	// Uninstalls traffic-agents and traffic-manager from the cluster.
	// Requires having already called Connect.
	Uninstall(ctx context.Context, in *UninstallRequest, opts ...grpc.CallOption) (*UninstallResult, error)
	// Returns a list of workloads and their current intercept status.
	// Requires having already called Connect.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*WorkloadInfoSnapshot, error)
	// Quits (terminates) the connector process.
	Quit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type connectorClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorClient(cc grpc.ClientConnInterface) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.VersionInfo, error) {
	out := new(common.VersionInfo)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectInfo, error) {
	out := new(ConnectInfo)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Status(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectInfo, error) {
	out := new(ConnectInfo)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) CreateIntercept(ctx context.Context, in *CreateInterceptRequest, opts ...grpc.CallOption) (*InterceptResult, error) {
	out := new(InterceptResult)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/CreateIntercept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) RemoveIntercept(ctx context.Context, in *manager.RemoveInterceptRequest2, opts ...grpc.CallOption) (*InterceptResult, error) {
	out := new(InterceptResult)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/RemoveIntercept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Uninstall(ctx context.Context, in *UninstallRequest, opts ...grpc.CallOption) (*UninstallResult, error) {
	out := new(UninstallResult)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/Uninstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*WorkloadInfoSnapshot, error) {
	out := new(WorkloadInfoSnapshot)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Quit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.connector.Connector/Quit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorServer is the server API for Connector service.
// All implementations must embed UnimplementedConnectorServer
// for forward compatibility
type ConnectorServer interface {
	// Returns version information from the Connector
	Version(context.Context, *empty.Empty) (*common.VersionInfo, error)
	// Connects to the cluster and connects the laptop's network (via
	// the daemon process) to the cluster's network.  A result code of
	// UNSPECIFIED indicates that the connection was successfully
	// initiated; if already connected, then either ALREADY_CONNECTED or
	// MUST_RESTART is returned, based on whether the current connection
	// is in agreement with the ConnectionRequest.
	Connect(context.Context, *ConnectRequest) (*ConnectInfo, error)
	// Status is much like Connect, except that it doesn't actually do
	// anything.  It's a dry-run.
	Status(context.Context, *ConnectRequest) (*ConnectInfo, error)
	// Adds an intercept to a workload.  Requires having already called
	// Connect.
	CreateIntercept(context.Context, *CreateInterceptRequest) (*InterceptResult, error)
	// Deactivates and removes an existent workload intercept.
	// Requires having already called Connect.
	RemoveIntercept(context.Context, *manager.RemoveInterceptRequest2) (*InterceptResult, error)
	// Uninstalls traffic-agents and traffic-manager from the cluster.
	// Requires having already called Connect.
	Uninstall(context.Context, *UninstallRequest) (*UninstallResult, error)
	// Returns a list of workloads and their current intercept status.
	// Requires having already called Connect.
	List(context.Context, *ListRequest) (*WorkloadInfoSnapshot, error)
	// Quits (terminates) the connector process.
	Quit(context.Context, *empty.Empty) (*empty.Empty, error)
	mustEmbedUnimplementedConnectorServer()
}

// UnimplementedConnectorServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorServer struct {
}

func (UnimplementedConnectorServer) Version(context.Context, *empty.Empty) (*common.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedConnectorServer) Connect(context.Context, *ConnectRequest) (*ConnectInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedConnectorServer) Status(context.Context, *ConnectRequest) (*ConnectInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedConnectorServer) CreateIntercept(context.Context, *CreateInterceptRequest) (*InterceptResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIntercept not implemented")
}
func (UnimplementedConnectorServer) RemoveIntercept(context.Context, *manager.RemoveInterceptRequest2) (*InterceptResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveIntercept not implemented")
}
func (UnimplementedConnectorServer) Uninstall(context.Context, *UninstallRequest) (*UninstallResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uninstall not implemented")
}
func (UnimplementedConnectorServer) List(context.Context, *ListRequest) (*WorkloadInfoSnapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedConnectorServer) Quit(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Quit not implemented")
}
func (UnimplementedConnectorServer) mustEmbedUnimplementedConnectorServer() {}

// UnsafeConnectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorServer will
// result in compilation errors.
type UnsafeConnectorServer interface {
	mustEmbedUnimplementedConnectorServer()
}

func RegisterConnectorServer(s grpc.ServiceRegistrar, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Status(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_CreateIntercept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInterceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).CreateIntercept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/CreateIntercept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).CreateIntercept(ctx, req.(*CreateInterceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_RemoveIntercept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(manager.RemoveInterceptRequest2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).RemoveIntercept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/RemoveIntercept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).RemoveIntercept(ctx, req.(*manager.RemoveInterceptRequest2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Uninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Uninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/Uninstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Uninstall(ctx, req.(*UninstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Quit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Quit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.connector.Connector/Quit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Quit(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telepresence.connector.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Connector_Version_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Connector_Connect_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Connector_Status_Handler,
		},
		{
			MethodName: "CreateIntercept",
			Handler:    _Connector_CreateIntercept_Handler,
		},
		{
			MethodName: "RemoveIntercept",
			Handler:    _Connector_RemoveIntercept_Handler,
		},
		{
			MethodName: "Uninstall",
			Handler:    _Connector_Uninstall_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Connector_List_Handler,
		},
		{
			MethodName: "Quit",
			Handler:    _Connector_Quit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/connector/connector.proto",
}
