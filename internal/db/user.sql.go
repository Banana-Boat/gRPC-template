// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
  username, password, gender, age
) VALUES (
  ?, ?, ?, ?
)
`

type CreateUserParams struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Gender   UsersGender `json:"gender"`
	Age      int32       `json:"age"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.Gender,
		arg.Age,
	)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, password, gender, age, created_at, updated_at FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Gender,
		&i.Age,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, gender, age, created_at, updated_at FROM users
WHERE username = ? LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Gender,
		&i.Age,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const isExistUser = `-- name: IsExistUser :one
SELECT EXISTS(
  SELECT 1 FROM users
  WHERE username = ? LIMIT 1
)
`

func (q *Queries) IsExistUser(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistUser, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, password, gender, age, created_at, updated_at FROM users
ORDER BY id
LIMIT ? OFFSET ?
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Gender,
			&i.Age,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET password = ?, gender = ?, age = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Password string      `json:"password"`
	Gender   UsersGender `json:"gender"`
	Age      int32       `json:"age"`
	ID       int32       `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Password,
		arg.Gender,
		arg.Age,
		arg.ID,
	)
	return err
}