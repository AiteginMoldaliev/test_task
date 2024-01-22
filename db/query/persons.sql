-- name: GetPerson :one
SELECT  *
FROM persons
WHERE id = $1
LIMIT 1;    

-- name: CreatePerson :one
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
RETURNING *;

-- name: UpdatePerson :one
UPDATE persons
    set firstname = $2,
    surname = $3, 
    patronymic = $4,
    gender = $5,
    age = $6,
    nationality = $7
WHERE id = $1
RETURNING *;

-- name: DeletePerson :exec
DELETE from persons
WHERE id = $1;

-- name: GetPersonsByFilter :many
SELECT id, firstname, surname, patronymic, gender, age, nationality, created_at
FROM persons
WHERE
	gender = $3
	AND (age <= $4 OR age >= $5)
	AND nationality = $6
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetPersonsList :many
SELECT id, firstname, surname, patronymic, gender, age, nationality, created_at
FROM persons
ORDER BY id
LIMIT $1
OFFSET $2;