package gapi

import (
	"fmt"

	db "github.com/MRD1920/Gobank.git/db/sqlc"
	"github.com/MRD1920/Gobank.git/pb"
	"github.com/MRD1920/Gobank.git/token"
	"github.com/MRD1920/Gobank.git/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedGoBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
