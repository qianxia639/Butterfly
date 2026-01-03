package db

import (
	"Butterfly/db/models"
	"context"
)

const createUser = `
INSERT INTO users (
	username, nickname, password, email, gender
) VALUES (
	$1, $2, $3, $4, $5
)
RETURNING id, username, nickname, password, email, gender, brithday, avatar_url, signature, password_changed_at, created_at, updated_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   int8   `json:"gender"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (models.User, error) {
	row := q.db.QueryRowxContext(ctx, createUser,
		arg.Username,
		arg.Nickname,
		arg.Password,
		arg.Email,
		arg.Gender,
	)
	var i models.User
	err := row.StructScan(&i)
	return i, err
}

const existsUsername = `
SELECT EXISTS (
	SELECT 1 FROM users 
	WHERE username = $1
)
`

func (q *Queries) ExistsUsername(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsUsername, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsEmail = `
SELECT EXISTS (
	SELECT 1 FROM users
	WHERE email = $1
)
`

func (q *Queries) ExistsEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUser = `
SELECT id, username, nickname, password, email, gender, brithday, avatar_url, signature, password_changed_at, created_at, updated_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (models.User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Nickname,
		&i.Password,
		&i.Email,
		&i.Gender,
		&i.Brithday,
		&i.AvatarUrl,
		&i.Signature,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
