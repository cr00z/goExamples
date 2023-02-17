package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

type AuthInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(jwtManager *JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager:      jwtManager,
		accessibleRoles: accessibleRoles,
	}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		log.Print("--> unary interceptor: ", info.FullMethod)
		return handler(ctx, req)
	}
}

func (i *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {

		err := i.authorize(ss.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		log.Print("--> stream interceptor: ")
		return handler(srv, ss)
	}
}

func (i *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, inMap := i.accessibleRoles[method]
	if !inMap {
		// full access
		return nil
	}

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	auth := meta["authorization"]
	if len(auth) == 0 {
		return status.Error(codes.Unauthenticated, "authorization token is empty")
	}

	accessToken := auth[0]
	log.Print(accessToken)
	claims, err := i.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if claims.Role == role {
			return nil
		}
	}

	return status.Error(codes.Unauthenticated, "no permissions to access this rpc")
}
