// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: llm.proto

package llm

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LLMService_GetEvaluation_FullMethodName = "/llm.LLMService/GetEvaluation"
	LLMService_GetArea_FullMethodName       = "/llm.LLMService/GetArea"
	LLMService_GetDomain_FullMethodName     = "/llm.LLMService/GetDomain"
)

// LLMServiceClient is the client API for LLMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义服务
type LLMServiceClient interface {
	GetEvaluation(ctx context.Context, in *GetEvaluationRequest, opts ...grpc.CallOption) (*GetEvaluationResponse, error)
	GetArea(ctx context.Context, in *GetAreaRequest, opts ...grpc.CallOption) (*GetAreaResponse, error)
	GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*GetDomainResponse, error)
}

type lLMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLLMServiceClient(cc grpc.ClientConnInterface) LLMServiceClient {
	return &lLMServiceClient{cc}
}

func (c *lLMServiceClient) GetEvaluation(ctx context.Context, in *GetEvaluationRequest, opts ...grpc.CallOption) (*GetEvaluationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEvaluationResponse)
	err := c.cc.Invoke(ctx, LLMService_GetEvaluation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lLMServiceClient) GetArea(ctx context.Context, in *GetAreaRequest, opts ...grpc.CallOption) (*GetAreaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAreaResponse)
	err := c.cc.Invoke(ctx, LLMService_GetArea_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lLMServiceClient) GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*GetDomainResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDomainResponse)
	err := c.cc.Invoke(ctx, LLMService_GetDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LLMServiceServer is the server API for LLMService service.
// All implementations must embed UnimplementedLLMServiceServer
// for forward compatibility.
//
// 定义服务
type LLMServiceServer interface {
	GetEvaluation(context.Context, *GetEvaluationRequest) (*GetEvaluationResponse, error)
	GetArea(context.Context, *GetAreaRequest) (*GetAreaResponse, error)
	GetDomain(context.Context, *GetDomainRequest) (*GetDomainResponse, error)
	mustEmbedUnimplementedLLMServiceServer()
}

// UnimplementedLLMServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLLMServiceServer struct{}

func (UnimplementedLLMServiceServer) GetEvaluation(context.Context, *GetEvaluationRequest) (*GetEvaluationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvaluation not implemented")
}
func (UnimplementedLLMServiceServer) GetArea(context.Context, *GetAreaRequest) (*GetAreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArea not implemented")
}
func (UnimplementedLLMServiceServer) GetDomain(context.Context, *GetDomainRequest) (*GetDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomain not implemented")
}
func (UnimplementedLLMServiceServer) mustEmbedUnimplementedLLMServiceServer() {}
func (UnimplementedLLMServiceServer) testEmbeddedByValue()                    {}

// UnsafeLLMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LLMServiceServer will
// result in compilation errors.
type UnsafeLLMServiceServer interface {
	mustEmbedUnimplementedLLMServiceServer()
}

func RegisterLLMServiceServer(s grpc.ServiceRegistrar, srv LLMServiceServer) {
	// If the following call pancis, it indicates UnimplementedLLMServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LLMService_ServiceDesc, srv)
}

func _LLMService_GetEvaluation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEvaluationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).GetEvaluation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LLMService_GetEvaluation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).GetEvaluation(ctx, req.(*GetEvaluationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LLMService_GetArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).GetArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LLMService_GetArea_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).GetArea(ctx, req.(*GetAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LLMService_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LLMService_GetDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).GetDomain(ctx, req.(*GetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LLMService_ServiceDesc is the grpc.ServiceDesc for LLMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LLMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "llm.LLMService",
	HandlerType: (*LLMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvaluation",
			Handler:    _LLMService_GetEvaluation_Handler,
		},
		{
			MethodName: "GetArea",
			Handler:    _LLMService_GetArea_Handler,
		},
		{
			MethodName: "GetDomain",
			Handler:    _LLMService_GetDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "llm.proto",
}
