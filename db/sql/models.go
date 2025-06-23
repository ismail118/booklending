package db

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type Book struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"ISBN"`
	Quantity int64  `json:"quantity"`
	Category string `json:"category"`
}

type LendingRecord struct {
	ID         int64        `json:"id"`
	Book       int64        `json:"book"`
	Borrower   int64        `json:"borrower"`
	IsReturn   bool         `json:"is_return"`
	BorrowDate time.Time    `json:"borrow_date"`
	ReturnDate sql.NullTime `json:"return_date"`
}
