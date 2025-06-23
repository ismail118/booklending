package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DBSOURCE = "mysql://root:password@tcp(localhost:3306)/booklending?tls=false"
)

func main() {
	dbConn, err := sql.Open("mysql", DBSOURCE)
	if err != nil {
		log.Fatalf("Error cannot connect to database: %v", err)
	}
	defer dbConn.Close()

	fmt.Println("test connect db")
}
