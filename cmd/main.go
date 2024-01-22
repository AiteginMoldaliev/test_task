package main

import (
	"database/sql"
	"log"

	"github.com/AiteginMoldaliev/test-task/api"
	db "github.com/AiteginMoldaliev/test-task/db/sqlc"
	"github.com/AiteginMoldaliev/test-task/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(config.Dbdriver, config.Dbsource)
	if err != nil {
		panic(err)
	}
	store := db.NewStore(conn)
	
	runGinServer(config, store)
}

func runGinServer(config util.Config, store *db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		panic(err)
	}

	log.Printf("HTTP server started on: %v", config.GinServerAddress)
	if err = server.Start(config.GinServerAddress); err != nil {
		panic(err)
	}
}
