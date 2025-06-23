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

const getListBookQuery = `
SELECT id, title, author, ISBN, quantity, category FROM books
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetListBookParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) GetListBook(ctx context.Context, arg GetListBookParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, getBookQuery, arg.Limit, arg.Offset)
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
