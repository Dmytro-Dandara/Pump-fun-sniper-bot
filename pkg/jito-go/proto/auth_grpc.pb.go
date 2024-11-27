// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: auth.proto

package proto

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

const (
	AuthService_GenerateAuthChallenge_FullMethodName = "/auth.AuthService/GenerateAuthChallenge"
	AuthService_GenerateAuthTokens_FullMethodName    = "/auth.AuthService/GenerateAuthTokens"
	AuthService_RefreshAccessToken_FullMethodName    = "/auth.AuthService/RefreshAccessToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// / Returns a challenge, client is expected to sign this challenge with an appropriate keypair in order to obtain access tokens.
	GenerateAuthChallenge(ctx context.Context, in *GenerateAuthChallengeRequest, opts ...grpc.CallOption) (*GenerateAuthChallengeResponse, error)
	// / Provides the client with the initial pair of auth tokens for API access.
	GenerateAuthTokens(ctx context.Context, in *GenerateAuthTokensRequest, opts ...grpc.CallOption) (*GenerateAuthTokensResponse, error)
	// / Call this method with a non-expired refresh token to obtain a new access token.
	RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GenerateAuthChallenge(ctx context.Context, in *GenerateAuthChallengeRequest, opts ...grpc.CallOption) (*GenerateAuthChallengeResponse, error) {
	out := new(GenerateAuthChallengeResponse)
	err := c.cc.Invoke(ctx, AuthService_GenerateAuthChallenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateAuthTokens(ctx context.Context, in *GenerateAuthTokensRequest, opts ...grpc.CallOption) (*GenerateAuthTokensResponse, error) {
	out := new(GenerateAuthTokensResponse)
	err := c.cc.Invoke(ctx, AuthService_GenerateAuthTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error) {
	out := new(RefreshAccessTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_RefreshAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// / Returns a challenge, client is expected to sign this challenge with an appropriate keypair in order to obtain access tokens.
	GenerateAuthChallenge(context.Context, *GenerateAuthChallengeRequest) (*GenerateAuthChallengeResponse, error)
	// / Provides the client with the initial pair of auth tokens for API access.
	GenerateAuthTokens(context.Context, *GenerateAuthTokensRequest) (*GenerateAuthTokensResponse, error)
	// / Call this method with a non-expired refresh token to obtain a new access token.
	RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) GenerateAuthChallenge(context.Context, *GenerateAuthChallengeRequest) (*GenerateAuthChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAuthChallenge not implemented")
}
func (UnimplementedAuthServiceServer) GenerateAuthTokens(context.Context, *GenerateAuthTokensRequest) (*GenerateAuthTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAuthTokens not implemented")
}
func (UnimplementedAuthServiceServer) RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_GenerateAuthChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAuthChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateAuthChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GenerateAuthChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateAuthChallenge(ctx, req.(*GenerateAuthChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateAuthTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAuthTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateAuthTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GenerateAuthTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateAuthTokens(ctx, req.(*GenerateAuthTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RefreshAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshAccessToken(ctx, req.(*RefreshAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAuthChallenge",
			Handler:    _AuthService_GenerateAuthChallenge_Handler,
		},
		{
			MethodName: "GenerateAuthTokens",
			Handler:    _AuthService_GenerateAuthTokens_Handler,
		},
		{
			MethodName: "RefreshAccessToken",
			Handler:    _AuthService_RefreshAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
