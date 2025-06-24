package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

const DBSOURCE = "root:password@tcp(localhost:3306)/booklending?tls=false&parseTime=true&loc=Local"

var testQuerier Querier

func TestMain(m *testing.M) {
	dbConn, err := sql.Open("mysql", DBSOURCE)
	if err != nil {
		log.Fatalf("Error cannot connect to database: %v", err)
	}

	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Error ping to database: %v", err)
	}

	testQuerier = NewQuerier(dbConn)

	os.Exit(m.Run())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const characters = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(characters)

	for i := 0; i < n; i++ {
		c := characters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}
