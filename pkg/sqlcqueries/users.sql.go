// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package sqlcqueries

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, full_name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateUserParams struct {
	ID        uuid.UUID `json:"id"`
	FullName  *string   `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

type DeleteUserParams struct {
	ID uuid.UUID `json:"id"`
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) error {
	_, err := q.db.Exec(ctx, deleteUser, arg.ID)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT u.id AS user_id, u.email, u.password, u.full_name, u.created_at, u.updated_at, avatar_img, r.name AS role_name
FROM users u
LEFT JOIN user_roles ur ON ur.user_id = u.id
LEFT JOIN roles r ON r.id = ur.role_id
WHERE u.email = $1
`

type GetUserByEmailParams struct {
	Email string `json:"email"`
}

type GetUserByEmailRow struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  *string   `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AvatarImg *string   `json:"avatar_img"`
	RoleName  *string   `json:"role_name"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, arg GetUserByEmailParams) (*GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, arg.Email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AvatarImg,
		&i.RoleName,
	)
	return &i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, full_name, created_at, updated_at, avatar_img
FROM users
WHERE id = $1
`

type GetUserByIDParams struct {
	ID uuid.UUID `json:"id"`
}

type GetUserByIDRow struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FullName  *string   `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AvatarImg *string   `json:"avatar_img"`
}

func (q *Queries) GetUserByID(ctx context.Context, arg GetUserByIDParams) (*GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, arg.ID)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FullName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AvatarImg,
	)
	return &i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, email, full_name, created_at, updated_at, avatar_img
FROM users
`

type GetUsersRow struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FullName  *string   `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AvatarImg *string   `json:"avatar_img"`
}

func (q *Queries) GetUsers(ctx context.Context) ([]*GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FullName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AvatarImg,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET email = $1, password = $2, full_name = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4
`

type UpdateUserParams struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	FullName *string   `json:"full_name"`
	ID       uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Email,
		arg.Password,
		arg.FullName,
		arg.ID,
	)
	return err
}

const updateUserAvatar = `-- name: UpdateUserAvatar :exec
UPDATE users
SET avatar_img = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
`

type UpdateUserAvatarParams struct {
	AvatarImg *string   `json:"avatar_img"`
	ID        uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUserAvatar(ctx context.Context, arg UpdateUserAvatarParams) error {
	_, err := q.db.Exec(ctx, updateUserAvatar, arg.AvatarImg, arg.ID)
	return err
}
