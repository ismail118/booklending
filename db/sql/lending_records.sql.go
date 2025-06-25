package db

import (
	"context"
	"database/sql"
	"time"
)

const createLendingRecordsQuery = `
INSERT INTO lending_records ( book, borrower) 
VALUES (?, ?)`

type CreateLendingRecordsParams struct {
	Book     int64 `json:"book"`
	Borrower int64 `json:"borrower"`
}

func (q *Queries) CreateLendingRecords(ctx context.Context, arg CreateLendingRecordsParams) (int64, error) {
	res, err := q.db.ExecContext(ctx, createLendingRecordsQuery, arg.Book, arg.Borrower)
	if err != nil {
		return 0, err
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return newId, nil
}

const returnBookQuery = `
UPDATE lending_records
SET return_date= ?
WHERE id = ? AND return_date = NULL
`

type ReturnBookParams struct {
	ID         int64     `json:"id"`
	ReturnDate time.Time `json:"return_date"`
}

func (q *Queries) ReturnBook(ctx context.Context, arg ReturnBookParams) error {
	res, err := q.db.ExecContext(ctx, returnBookQuery, arg.ReturnDate, arg.ID)
	if err != nil {
		return err
	}

	rowsEff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsEff < 1 {
		return sql.ErrNoRows
	}

	return nil
}

const getListLendingRecordByBorrowerQuery = `
SELECT id, book, borrower, borrow_date, return_date FROM lending_records
WHERE borrower = ? AND return_date IS NULL
`

func (q *Queries) GetListLendingRecordByBorrower(ctx context.Context, borrower int64) ([]LendingRecord, error) {
	rows, err := q.db.QueryContext(ctx, getListLendingRecordByBorrowerQuery, borrower)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []LendingRecord{}

	for rows.Next() {
		var d LendingRecord
		err := rows.Scan(
			&d.ID,
			&d.Book,
			&d.Borrower,
			&d.BorrowDate,
			&d.ReturnDate,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, d)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return items, nil
}

const getLendingRecordQuery = `
SELECT id, book, borrower, borrow_date, return_date FROM lending_records
WHERE id = ?
`

func (q *Queries) GetLendingRecord(ctx context.Context, id int64) (LendingRecord, error) {
	row := q.db.QueryRowContext(ctx, getLendingRecordQuery, id)
	var d LendingRecord
	err := row.Scan(
		&d.ID,
		&d.Book,
		&d.Borrower,
		&d.BorrowDate,
		&d.ReturnDate,
	)

	return d, err
}
