package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pcbook/pb"
)

type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unknown user")
	}

	if !user.IsPasswordCorrect(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "unknown user")
	}

	token, err := s.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "jwt generate error")
	}

	res := &pb.LoginResponse{
		AccessToken: token,
	}
	return res, nil
}
