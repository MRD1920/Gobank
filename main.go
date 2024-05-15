package main

import (
	"database/sql"
	"log"

	"github.com/MRD1920/Gobank.git/api"
	db "github.com/MRD1920/Gobank.git/db/sqlc"
	"github.com/MRD1920/Gobank.git/util"
	_ "github.com/lib/pq"
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
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
