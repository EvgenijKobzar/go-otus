// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: account.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addAccount = `-- name: AddAccount :one
INSERT INTO movies_online.account (
    firstName,
    lastName,
    login,
    password
)
VALUES
    (
     $1,
     $2,
     $3,
     $4
    )
RETURNING id
`

type AddAccountParams struct {
	Firstname sql.NullString
	Lastname  sql.NullString
	Login     sql.NullString
	Password  sql.NullString
}

func (q *Queries) AddAccount(ctx context.Context, arg AddAccountParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, addAccount,
		arg.Firstname,
		arg.Lastname,
		arg.Login,
		arg.Password,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteAccount = `-- name: DeleteAccount :exec
delete from movies_online.account
where id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAllAccount = `-- name: GetAllAccount :many
select id, firstname, lastname, login, password, created_at, updated_at from movies_online.account
`

func (q *Queries) GetAllAccount(ctx context.Context) ([]MoviesOnlineAccount, error) {
	rows, err := q.db.QueryContext(ctx, getAllAccount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MoviesOnlineAccount
	for rows.Next() {
		var i MoviesOnlineAccount
		if err := rows.Scan(
			&i.ID,
			&i.Firstname,
			&i.Lastname,
			&i.Login,
			&i.Password,
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

const getByIdAccount = `-- name: GetByIdAccount :one
select id, firstname, lastname, login, password, created_at, updated_at
from movies_online.account
where id = $1
`

func (q *Queries) GetByIdAccount(ctx context.Context, id int32) (MoviesOnlineAccount, error) {
	row := q.db.QueryRowContext(ctx, getByIdAccount, id)
	var i MoviesOnlineAccount
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Login,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE movies_online.account SET
                                 firstName=$2,
                                 lastName=$3,
                                 login=$4,
                                 password=$5
WHERE id = $1
`

type UpdateAccountParams struct {
	ID        int32
	Firstname sql.NullString
	Lastname  sql.NullString
	Login     sql.NullString
	Password  sql.NullString
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.db.ExecContext(ctx, updateAccount,
		arg.ID,
		arg.Firstname,
		arg.Lastname,
		arg.Login,
		arg.Password,
	)
	return err
}
