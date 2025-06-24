package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ismail118/booklending/api"
	db "github.com/ismail118/booklending/db/sql"
	"log"
)

const (
	DBSOURCE      = "root:password@tcp(localhost:3306)/booklending?tls=false&parseTime=true&loc=Local"
	SERVERADDRESS = "0.0.0.0:8080"
)

func main() {
	dbConn, err := sql.Open("mysql", DBSOURCE)
	if err != nil {
		log.Fatalf("Error cannot connect to database: %v", err)
	}

	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Error ping to database: %v", err)
	}

	querier := db.NewQuerier(dbConn)

	server := api.NewServer(querier)
	err = server.Start(SERVERADDRESS)
	if err != nil {
		log.Fatalf("Error cannot start server: %v", err)
	}

	log.Println("Starting server on", SERVERADDRESS)
}
