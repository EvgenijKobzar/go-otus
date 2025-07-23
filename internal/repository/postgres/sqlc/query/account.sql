-- name: AddAccount :one
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
RETURNING id;

-- name: UpdateAccount :exec
UPDATE movies_online.account SET
                                 firstName=$2,
                                 lastName=$3,
                                 login=$4,
                                 password=$5
WHERE id = $1;

-- name: GetAllAccount :many
select * from movies_online.account;

-- name: GetByIdAccount :one
select *
from movies_online.account
where id = $1;

-- name: DeleteAccount :exec
delete from movies_online.account
where id = $1;