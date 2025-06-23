package db

import (
	"context"
)

const createBookQuery = `
INSERT INTO books ( title, author, ISBN, quantity, category ) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING id, title, author, ISBN, quantity, category`

type CreateBookParams struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Quantity int64  `json:"quantity"`
	Category string `json:"category"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBookQuery, arg.Title, arg.Author, arg.ISBN, arg.Quantity, arg.Category)
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

const getBookQuery = `
SELECT id, title, author, ISBN, quantity, category FROM books
WHERE id = $1
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
SET Title=$1, Author=$2, ISBN=$3, Quantity=$4, category=$5
WHERE id = $6
RETURNING id, title, author, ISBN, quantity, category
`

type UpdateBookParams struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Quantity int64  `json:"quantity"`
	Category string `json:"category"`
	ID       int64  `json:"id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBookQuery, arg.Title, arg.Author, arg.ISBN, arg.Quantity, arg.Category, arg.ID)
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

const deleteBookQuery = `
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBookQuery, id)
	return err
}
