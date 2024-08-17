package gapi

import (
	"context"

	db "github.com/MRD1920/Gobank.git/db/sqlc"
	"github.com/MRD1920/Gobank.git/pb"
	"github.com/MRD1920/Gobank.git/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	HashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash the password: %v", err)
	}
	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: HashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exists: %v", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	response := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return response, nil
	// return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
