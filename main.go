package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ismail118/booklending/api"
	db "github.com/ismail118/booklending/db/sql"
	"github.com/ismail118/booklending/token"
	"log"
)

const (
	DBSOURCE      = "root:password@tcp(localhost:3306)/booklending?tls=false&parseTime=true&loc=Local"
	SERVERADDRESS = "0.0.0.0:8080"
	TokenKey      = "abcdefghijklmnovqrstuvwxyz123456"
)

func main() {
	dbConn, err := sql.Open("mysql", DBSOURCE)
	if err != nil {
		log.Fatalf("Error cannot connect to database: %v", err)
	}

	defer dbConn.Close()

	querier := db.NewQuerier(dbConn)

	paseto, err := token.NewPaseto(TokenKey)
	if err != nil {
		log.Fatalf("Error cannot make paseto %v", err)
	}

	server := api.NewServer(querier, paseto)
	err = server.Start(SERVERADDRESS)
	if err != nil {
		log.Fatalf("Error cannot start server: %v", err)
	}

	log.Println("Starting server on", SERVERADDRESS)
}
