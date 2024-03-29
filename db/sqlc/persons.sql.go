// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: persons.sql

package db

import (
	"context"
	"database/sql"
)

const createPerson = `-- name: CreatePerson :one
INSERT INTO persons (
    firstname,
    surname,
    patronymic,
    gender,
    age,
    nationality 
)
VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, firstname, surname, patronymic, gender, age, nationality, created_at
`

type CreatePersonParams struct {
	Firstname   string         `json:"firstname"`
	Surname     string         `json:"surname"`
	Patronymic  sql.NullString `json:"patronymic"`
	Gender      string         `json:"gender"`
	Age         int64          `json:"age"`
	Nationality string         `json:"nationality"`
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, createPerson,
		arg.Firstname,
		arg.Surname,
		arg.Patronymic,
		arg.Gender,
		arg.Age,
		arg.Nationality,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Surname,
		&i.Patronymic,
		&i.Gender,
		&i.Age,
		&i.Nationality,
		&i.CreatedAt,
	)
	return i, err
}

const deletePerson = `-- name: DeletePerson :exec
DELETE from persons
WHERE id = $1
`

func (q *Queries) DeletePerson(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePerson, id)
	return err
}

const getPerson = `-- name: GetPerson :one
SELECT  id, firstname, surname, patronymic, gender, age, nationality, created_at
FROM persons
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetPerson(ctx context.Context, id int64) (Person, error) {
	row := q.db.QueryRowContext(ctx, getPerson, id)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Surname,
		&i.Patronymic,
		&i.Gender,
		&i.Age,
		&i.Nationality,
		&i.CreatedAt,
	)
	return i, err
}

const getPersonsByFilter = `-- name: GetPersonsByFilter :many
SELECT id, firstname, surname, patronymic, gender, age, nationality, created_at
FROM persons
WHERE
	gender = $3
	AND (age <= $4 OR age >= $5)
	AND nationality = $6
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetPersonsByFilterParams struct {
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
	Gender      string `json:"gender"`
	Age         int64  `json:"age"`
	Age_2       int64  `json:"age_2"`
	Nationality string `json:"nationality"`
}

func (q *Queries) GetPersonsByFilter(ctx context.Context, arg GetPersonsByFilterParams) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, getPersonsByFilter,
		arg.Limit,
		arg.Offset,
		arg.Gender,
		arg.Age,
		arg.Age_2,
		arg.Nationality,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.Firstname,
			&i.Surname,
			&i.Patronymic,
			&i.Gender,
			&i.Age,
			&i.Nationality,
			&i.CreatedAt,
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

const getPersonsList = `-- name: GetPersonsList :many
SELECT id, firstname, surname, patronymic, gender, age, nationality, created_at
FROM persons
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetPersonsListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPersonsList(ctx context.Context, arg GetPersonsListParams) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, getPersonsList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.Firstname,
			&i.Surname,
			&i.Patronymic,
			&i.Gender,
			&i.Age,
			&i.Nationality,
			&i.CreatedAt,
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

const updatePerson = `-- name: UpdatePerson :one
UPDATE persons
    set firstname = $2,
    surname = $3, 
    patronymic = $4,
    gender = $5,
    age = $6,
    nationality = $7
WHERE id = $1
RETURNING id, firstname, surname, patronymic, gender, age, nationality, created_at
`

type UpdatePersonParams struct {
	ID          int64          `json:"id"`
	Firstname   string         `json:"firstname"`
	Surname     string         `json:"surname"`
	Patronymic  sql.NullString `json:"patronymic"`
	Gender      string         `json:"gender"`
	Age         int64          `json:"age"`
	Nationality string         `json:"nationality"`
}

func (q *Queries) UpdatePerson(ctx context.Context, arg UpdatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, updatePerson,
		arg.ID,
		arg.Firstname,
		arg.Surname,
		arg.Patronymic,
		arg.Gender,
		arg.Age,
		arg.Nationality,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Surname,
		&i.Patronymic,
		&i.Gender,
		&i.Age,
		&i.Nationality,
		&i.CreatedAt,
	)
	return i, err
}
