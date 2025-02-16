// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: authzed/api/v0/developer.proto

package v0

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

// DeveloperServiceClient is the client API for DeveloperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeveloperServiceClient interface {
	EditCheck(ctx context.Context, in *EditCheckRequest, opts ...grpc.CallOption) (*EditCheckResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	LookupShared(ctx context.Context, in *LookupShareRequest, opts ...grpc.CallOption) (*LookupShareResponse, error)
	UpgradeSchema(ctx context.Context, in *UpgradeSchemaRequest, opts ...grpc.CallOption) (*UpgradeSchemaResponse, error)
	FormatSchema(ctx context.Context, in *FormatSchemaRequest, opts ...grpc.CallOption) (*FormatSchemaResponse, error)
}

type developerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeveloperServiceClient(cc grpc.ClientConnInterface) DeveloperServiceClient {
	return &developerServiceClient{cc}
}

func (c *developerServiceClient) EditCheck(ctx context.Context, in *EditCheckRequest, opts ...grpc.CallOption) (*EditCheckResponse, error) {
	out := new(EditCheckResponse)
	err := c.cc.Invoke(ctx, "/authzed.api.v0.DeveloperService/EditCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/authzed.api.v0.DeveloperService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/authzed.api.v0.DeveloperService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) LookupShared(ctx context.Context, in *LookupShareRequest, opts ...grpc.CallOption) (*LookupShareResponse, error) {
	out := new(LookupShareResponse)
	err := c.cc.Invoke(ctx, "/authzed.api.v0.DeveloperService/LookupShared", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) UpgradeSchema(ctx context.Context, in *UpgradeSchemaRequest, opts ...grpc.CallOption) (*UpgradeSchemaResponse, error) {
	out := new(UpgradeSchemaResponse)
	err := c.cc.Invoke(ctx, "/authzed.api.v0.DeveloperService/UpgradeSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developerServiceClient) FormatSchema(ctx context.Context, in *FormatSchemaRequest, opts ...grpc.CallOption) (*FormatSchemaResponse, error) {
	out := new(FormatSchemaResponse)
	err := c.cc.Invoke(ctx, "/authzed.api.v0.DeveloperService/FormatSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeveloperServiceServer is the server API for DeveloperService service.
// All implementations must embed UnimplementedDeveloperServiceServer
// for forward compatibility
type DeveloperServiceServer interface {
	EditCheck(context.Context, *EditCheckRequest) (*EditCheckResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	Share(context.Context, *ShareRequest) (*ShareResponse, error)
	LookupShared(context.Context, *LookupShareRequest) (*LookupShareResponse, error)
	UpgradeSchema(context.Context, *UpgradeSchemaRequest) (*UpgradeSchemaResponse, error)
	FormatSchema(context.Context, *FormatSchemaRequest) (*FormatSchemaResponse, error)
	mustEmbedUnimplementedDeveloperServiceServer()
}

// UnimplementedDeveloperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeveloperServiceServer struct {
}

func (UnimplementedDeveloperServiceServer) EditCheck(context.Context, *EditCheckRequest) (*EditCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCheck not implemented")
}
func (UnimplementedDeveloperServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedDeveloperServiceServer) Share(context.Context, *ShareRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedDeveloperServiceServer) LookupShared(context.Context, *LookupShareRequest) (*LookupShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupShared not implemented")
}
func (UnimplementedDeveloperServiceServer) UpgradeSchema(context.Context, *UpgradeSchemaRequest) (*UpgradeSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeSchema not implemented")
}
func (UnimplementedDeveloperServiceServer) FormatSchema(context.Context, *FormatSchemaRequest) (*FormatSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormatSchema not implemented")
}
func (UnimplementedDeveloperServiceServer) mustEmbedUnimplementedDeveloperServiceServer() {}

// UnsafeDeveloperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeveloperServiceServer will
// result in compilation errors.
type UnsafeDeveloperServiceServer interface {
	mustEmbedUnimplementedDeveloperServiceServer()
}

func RegisterDeveloperServiceServer(s grpc.ServiceRegistrar, srv DeveloperServiceServer) {
	s.RegisterService(&DeveloperService_ServiceDesc, srv)
}

func _DeveloperService_EditCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).EditCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authzed.api.v0.DeveloperService/EditCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).EditCheck(ctx, req.(*EditCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authzed.api.v0.DeveloperService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authzed.api.v0.DeveloperService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).Share(ctx, req.(*ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_LookupShared_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).LookupShared(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authzed.api.v0.DeveloperService/LookupShared",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).LookupShared(ctx, req.(*LookupShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_UpgradeSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).UpgradeSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authzed.api.v0.DeveloperService/UpgradeSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).UpgradeSchema(ctx, req.(*UpgradeSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeveloperService_FormatSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormatSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeveloperServiceServer).FormatSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authzed.api.v0.DeveloperService/FormatSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeveloperServiceServer).FormatSchema(ctx, req.(*FormatSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeveloperService_ServiceDesc is the grpc.ServiceDesc for DeveloperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeveloperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authzed.api.v0.DeveloperService",
	HandlerType: (*DeveloperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditCheck",
			Handler:    _DeveloperService_EditCheck_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _DeveloperService_Validate_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _DeveloperService_Share_Handler,
		},
		{
			MethodName: "LookupShared",
			Handler:    _DeveloperService_LookupShared_Handler,
		},
		{
			MethodName: "UpgradeSchema",
			Handler:    _DeveloperService_UpgradeSchema_Handler,
		},
		{
			MethodName: "FormatSchema",
			Handler:    _DeveloperService_FormatSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authzed/api/v0/developer.proto",
}
