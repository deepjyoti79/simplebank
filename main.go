package main

import (
	"database/sql"
	"log"

	"github.com/deepjyoti79/simplebank/api"
	db "github.com/deepjyoti79/simplebank/db/sqlc"
	"github.com/deepjyoti79/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("error loading env file")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdress)
}
