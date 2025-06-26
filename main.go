package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ismail118/booklending/api"
	db "github.com/ismail118/booklending/db/sql"
	"github.com/ismail118/booklending/token"
	"github.com/ismail118/booklending/util"
	"github.com/spf13/viper"
	"log"
)

func main() {
	config, err := readConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := sql.Open("mysql", config.DbSource)
	if err != nil {
		log.Fatalf("Error cannot connect to database: %v", err)
	}

	defer dbConn.Close()

	querier := db.NewQuerier(dbConn)

	paseto, err := token.NewPaseto(config.SecretKey)
	if err != nil {
		log.Fatalf("Error cannot make paseto %v", err)
	}

	server := api.NewServer(querier, paseto)
	err = server.Start(config.AddrServer)
	if err != nil {
		log.Fatalf("Error cannot start server: %v", err)
	}

	log.Println("Starting server on", config.AddrServer)
}

func readConfig(path string) (util.Config, error) {
	var config util.Config

	viper.SetConfigName("APP")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("cannot read file config: %v", err)
	}
	viper.AutomaticEnv()
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("cannot unmarshal config: %v", err)
	}

	return config, nil
}
