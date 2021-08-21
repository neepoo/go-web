package main

import (
	"database/sql"
	"github.com/neepoo/go-web/util"
	"log"

	_ "github.com/lib/pq"
	"github.com/neepoo/go-web/api"

	db "github.com/neepoo/go-web/db/sqlc"
)

 func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("load config error", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
