package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (int64, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) error
	DeleteBook(ctx context.Context, id int64) error
	GetListBook(ctx context.Context, arg GetListBookParams) ([]Book, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetUser(ctx context.Context, email string) (User, error)
	CreateLendingRecords(ctx context.Context, arg CreateLendingRecordsParams) (int64, error)
	ReturnBook(ctx context.Context, arg returnBookParams) error
	GetListLendingRecordByBorrower(ctx context.Context, borrower int64) ([]LendingRecord, error)
	GetLendingRecord(ctx context.Context, id int64) (LendingRecord, error)
}

func NewQuerier(db *sql.DB) Querier {
	return &Queries{db: db}
}

type Queries struct {
	db *sql.DB
}
