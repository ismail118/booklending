package db

import "context"

const createUserQuery = `
INSERT INTO users ( name, email, hashed_password) 
VALUES (?, ?, ?)`

type CreateUserParams struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	res, err := q.db.ExecContext(ctx, createUserQuery, arg.Name, arg.Email, arg.HashedPassword)
	if err != nil {
		return 0, err
	}

	newId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return newId, nil
}

const getUserQuery = `
SELECT id, name, email, hashed_password FROM users
WHERE email = ?
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserQuery, email)
	var d User
	err := row.Scan(
		&d.ID,
		&d.Name,
		&d.Email,
		&d.HashedPassword,
	)

	return d, err
}
