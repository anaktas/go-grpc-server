// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RecipesClient is the client API for Recipes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecipesClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetUserRecipes(ctx context.Context, in *GetUserRecipesRequest, opts ...grpc.CallOption) (*GetUserRecipesResponse, error)
}

type recipesClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipesClient(cc grpc.ClientConnInterface) RecipesClient {
	return &recipesClient{cc}
}

func (c *recipesClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/recipes.Recipes/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/recipes.Recipes/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesClient) GetUserRecipes(ctx context.Context, in *GetUserRecipesRequest, opts ...grpc.CallOption) (*GetUserRecipesResponse, error) {
	out := new(GetUserRecipesResponse)
	err := c.cc.Invoke(ctx, "/recipes.Recipes/GetUserRecipes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipesServer is the server API for Recipes service.
// All implementations must embed UnimplementedRecipesServer
// for forward compatibility
type RecipesServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetUserRecipes(context.Context, *GetUserRecipesRequest) (*GetUserRecipesResponse, error)
	mustEmbedUnimplementedRecipesServer()
}

// UnimplementedRecipesServer must be embedded to have forward compatible implementations.
type UnimplementedRecipesServer struct {
}

func (UnimplementedRecipesServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRecipesServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRecipesServer) GetUserRecipes(context.Context, *GetUserRecipesRequest) (*GetUserRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRecipes not implemented")
}
func (UnimplementedRecipesServer) mustEmbedUnimplementedRecipesServer() {}

// UnsafeRecipesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecipesServer will
// result in compilation errors.
type UnsafeRecipesServer interface {
	mustEmbedUnimplementedRecipesServer()
}

func RegisterRecipesServer(s grpc.ServiceRegistrar, srv RecipesServer) {
	s.RegisterService(&_Recipes_serviceDesc, srv)
}

func _Recipes_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipes.Recipes/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recipes_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipes.Recipes/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recipes_GetUserRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRecipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).GetUserRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipes.Recipes/GetUserRecipes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).GetUserRecipes(ctx, req.(*GetUserRecipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recipes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recipes.Recipes",
	HandlerType: (*RecipesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Recipes_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Recipes_Login_Handler,
		},
		{
			MethodName: "GetUserRecipes",
			Handler:    _Recipes_GetUserRecipes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/recipes.proto",
}
