package main

import (
	"database/sql"
	"github.com/ZeinabAshjaei/go-master/api"
	db "github.com/ZeinabAshjaei/go-master/db/sqlc"
	"github.com/ZeinabAshjaei/go-master/utils"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Connection to the DB could not be established")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.SeverAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
