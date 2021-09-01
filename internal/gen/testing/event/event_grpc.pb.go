// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package event

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

// TestAuthMethodServiceClient is the client API for TestAuthMethodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestAuthMethodServiceClient interface {
	// TestAuthenticate validates credentials provided and returns an Auth Token.
	// Deprecated: use AuthenticateLogin instead.
	TestAuthenticate(ctx context.Context, in *TestAuthenticateRequest, opts ...grpc.CallOption) (*TestAuthenticateResponse, error)
}

type testAuthMethodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestAuthMethodServiceClient(cc grpc.ClientConnInterface) TestAuthMethodServiceClient {
	return &testAuthMethodServiceClient{cc}
}

func (c *testAuthMethodServiceClient) TestAuthenticate(ctx context.Context, in *TestAuthenticateRequest, opts ...grpc.CallOption) (*TestAuthenticateResponse, error) {
	out := new(TestAuthenticateResponse)
	err := c.cc.Invoke(ctx, "/testing.event.v1.TestAuthMethodService/TestAuthenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestAuthMethodServiceServer is the server API for TestAuthMethodService service.
// All implementations must embed UnimplementedTestAuthMethodServiceServer
// for forward compatibility
type TestAuthMethodServiceServer interface {
	// TestAuthenticate validates credentials provided and returns an Auth Token.
	// Deprecated: use AuthenticateLogin instead.
	TestAuthenticate(context.Context, *TestAuthenticateRequest) (*TestAuthenticateResponse, error)
	mustEmbedUnimplementedTestAuthMethodServiceServer()
}

// UnimplementedTestAuthMethodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestAuthMethodServiceServer struct {
}

func (UnimplementedTestAuthMethodServiceServer) TestAuthenticate(context.Context, *TestAuthenticateRequest) (*TestAuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAuthenticate not implemented")
}
func (UnimplementedTestAuthMethodServiceServer) mustEmbedUnimplementedTestAuthMethodServiceServer() {}

// UnsafeTestAuthMethodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestAuthMethodServiceServer will
// result in compilation errors.
type UnsafeTestAuthMethodServiceServer interface {
	mustEmbedUnimplementedTestAuthMethodServiceServer()
}

func RegisterTestAuthMethodServiceServer(s grpc.ServiceRegistrar, srv TestAuthMethodServiceServer) {
	s.RegisterService(&TestAuthMethodService_ServiceDesc, srv)
}

func _TestAuthMethodService_TestAuthenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestAuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestAuthMethodServiceServer).TestAuthenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testing.event.v1.TestAuthMethodService/TestAuthenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestAuthMethodServiceServer).TestAuthenticate(ctx, req.(*TestAuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestAuthMethodService_ServiceDesc is the grpc.ServiceDesc for TestAuthMethodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestAuthMethodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testing.event.v1.TestAuthMethodService",
	HandlerType: (*TestAuthMethodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestAuthenticate",
			Handler:    _TestAuthMethodService_TestAuthenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testing/event/v1/event.proto",
}
