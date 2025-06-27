package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

// error ref: https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
// used: db.IsSqlErr(err, db.ErrNumForeignKeyViolation)
var (
	ErrNumForeignKeyViolation uint16 = 1216
	ErrNumUniqueViolation     uint16 = 1062
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
	ReturnBook(ctx context.Context, arg ReturnBookParams) error
	GetListLendingRecordByBorrower(ctx context.Context, borrower int64) ([]LendingRecord, error)
	GetLendingRecord(ctx context.Context, id int64) (LendingRecord, error)
	UpdateBookQty(ctx context.Context, arg UpdateBookQtyParams) error
}

func NewQuerier(db *sql.DB) Querier {
	return &Queries{db: db}
}

type Queries struct {
	db *sql.DB
}

func IsSqlErr(err error, errNum uint16) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == errNum
	}

	return false
}
