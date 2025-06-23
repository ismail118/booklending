package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
	DeleteBook(ctx context.Context, id int64) error
	GetListBook(ctx context.Context, arg GetListBookParams) ([]Book, error)
}

func NewQuerier(db *sql.DB) Querier {
	return &Queries{db: db}
}

type Queries struct {
	db *sql.DB
}
