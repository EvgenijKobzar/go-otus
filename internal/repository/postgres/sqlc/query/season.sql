-- name: AddSeason :one
INSERT INTO movies_online.season (
                                  SERIAL_ID,
                                  SORT,
                                  ACTIVE,
                                  TITLE
)
VALUES
    (
        $1,
        $2,
        $3,
        $4
    )
RETURNING id;

-- name: UpdateSeason :exec
UPDATE movies_online.season SET
                                 SERIAL_ID=$2,
                                 SORT=$3,
                                 ACTIVE=$4,
                                 TITLE=$5
WHERE id = $1;

-- name: GetAllSeason :many
select * from movies_online.season;

-- name: GetByIdSeason :one
select *
from movies_online.season
where id = $1;

-- name: DeleteSeason :exec
delete from movies_online.season
where id = $1;