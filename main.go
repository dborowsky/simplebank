package main

import (
	"database/sql"
	"log"

	"github.com/dborowsky/simplebank/util"

	_ "github.com/lib/pq"

	"github.com/dborowsky/simplebank/api"
	db "github.com/dborowsky/simplebank/db/sqlc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %s\n", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connection to DB", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %s\n", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
