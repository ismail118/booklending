package db

import "database/sql"

type Querier interface{}

func NewQuerier(db *sql.DB) Querier {
	return &Queries{db: db}
}

type Queries struct {
	db *sql.DB
}
