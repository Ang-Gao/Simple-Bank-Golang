package main

import (
	"database/sql"
	"log"
	api "myproject/api"
	db "myproject/db/sqlc"

	_ "github.com/lib/pq" //for driver detected as postgres
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgresql://root:abc123@localhost:5432/simple_bank?sslmode=disable"
	serverAddr = "localhost:8080"
)

func main() {
	//setup connection
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
