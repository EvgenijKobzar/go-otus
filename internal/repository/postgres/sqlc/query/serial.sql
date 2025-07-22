-- name: AddSerial :one
INSERT INTO movies_online.serial (
                                  SORT,
                                  ACTIVE,
                                  FILE_ID,
                                  TITLE,
                                  PRODUCTION_PERIOD,
                                  RATING,
                                  QUALITY,
                                  DURATION
)
VALUES
    (
     $1,
     $2,
     $3,
     $4,
     $5,
     $6,
     $7,
     $8
    )
RETURNING id;

-- name: UpdateSerial :exec
UPDATE movies_online.serial SET
                                SORT=$2,
                                ACTIVE=$3,
                                FILE_ID=$4,
                                TITLE=$5,
                                PRODUCTION_PERIOD=$6,
                                RATING=$7,
                                QUALITY=$8,
                                DURATION=$9
WHERE id = $1;

-- name: GetAllSerial :many
select * from movies_online.serial;

-- name: GetByIdSerial :one
select *
from movies_online.serial
where id = $1;

-- name: DeleteSerial :exec
delete from movies_online.serial
where id = $1;