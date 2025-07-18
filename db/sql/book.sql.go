package db

import (
	"context"
	"database/sql"
)

const createBookQuery = `
INSERT INTO books ( title, author, ISBN, quantity, category ) 
VALUES (?, ?, ?, ?, ?)`

type CreateBookParams struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Quantity int64  `json:"quantity"`
	Category string `json:"category"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (int64, error) {
	res, err := q.db.ExecContext(ctx, createBookQuery, arg.Title, arg.Author, arg.ISBN, arg.Quantity, arg.Category)
	if err != nil {
		return 0, err
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return newId, nil
}

const getBookQuery = `
SELECT id, title, author, ISBN, quantity, category FROM books
WHERE id = ?
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBookQuery, id)
	var d Book
	err := row.Scan(
		&d.ID,
		&d.Title,
		&d.Author,
		&d.ISBN,
		&d.Quantity,
		&d.Category,
	)

	return d, err
}

const updateBookQuery = `
UPDATE books
SET title=?, author=?, ISBN=?, quantity=?, category=?
WHERE id = ?
`

type UpdateBookParams struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Quantity int64  `json:"quantity"`
	Category string `json:"category"`
	ID       int64  `json:"id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	res, err := q.db.ExecContext(ctx, updateBookQuery, arg.Title, arg.Author, arg.ISBN, arg.Quantity, arg.Category, arg.ID)
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

const deleteBookQuery = `
DELETE FROM books
WHERE id = ?
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBookQuery, id)
	return err
}

const getListBookQuery = `
SELECT id, title, author, ISBN, quantity, category FROM books
ORDER BY id
LIMIT ? 
OFFSET ?
`

type GetListBookParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) GetListBook(ctx context.Context, arg GetListBookParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getListBookQuery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Book{}

	for rows.Next() {
		var d Book
		err := rows.Scan(
			&d.ID,
			&d.Title,
			&d.Author,
			&d.ISBN,
			&d.Quantity,
			&d.Category,
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

const updateBookQtyQuery = `
UPDATE books
SET quantity= quantity + ?
WHERE id = ?
`

type UpdateBookQtyParams struct {
	Qty int64 `json:"qty"`
	ID  int64 `json:"id"`
}

func (q *Queries) UpdateBookQty(ctx context.Context, arg UpdateBookQtyParams) error {
	res, err := q.db.ExecContext(ctx, updateBookQtyQuery, arg.Qty, arg.ID)
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
