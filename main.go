package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/MRD1920/Gobank.git/api"
	db "github.com/MRD1920/Gobank.git/db/sqlc"
	"github.com/MRD1920/Gobank.git/gapi"
	"github.com/MRD1920/Gobank.git/pb"
	"github.com/MRD1920/Gobank.git/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Open a database connection
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// Create a new store
	store := db.NewStore(conn)

	//run the Gin server
	// runGinServer(config, store)
	runGrpcServer(config, store)

}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create a grpc server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGoBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	log.Printf("start gRPC server at %s", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
