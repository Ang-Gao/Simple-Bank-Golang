package main

import (
	"database/sql"
	"log"
	api "myproject/api"
	db "myproject/db/sqlc"
	util "myproject/util"

	_ "github.com/lib/pq" //for driver detected as postgres
)

func main() {
	//load configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("config error")
	}

	//setup connection
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
