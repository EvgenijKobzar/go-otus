-- name: AddEpisode :one
INSERT INTO movies_online.episode (
                                   QUALITY,
                                   RATING,
                                   PRODUCTION_PERIOD,
                                   SERIAL_ID,
                                   SEASON_ID,
                                   SORT,
                                   ACTIVE,
                                   FILE_ID,
                                   TITLE,
                                   DURATION,
                                   DESCRIPTION
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
     $8,
     $9,
     $10,
     $11
    )
RETURNING id;

-- name: UpdateEpisode :exec
UPDATE movies_online.episode SET
                                 QUALITY = $2,
                                 RATING = $3,
                                 PRODUCTION_PERIOD = $4,
                                 SERIAL_ID = $5,
                                 SEASON_ID = $6,
                                 SORT = $7,
                                 ACTIVE = $8,
                                 FILE_ID = $9,
                                 TITLE = $10,
                                 DURATION = $11,
                                 DESCRIPTION = $12
WHERE id = $1;

-- name: GetAllEpisode :many
select * from movies_online.episode;

-- name: GetByIdEpisode :one
select *
from movies_online.episode
where id = $1;

-- name: DeleteEpisode :exec
delete from movies_online.episode
where id = $1;