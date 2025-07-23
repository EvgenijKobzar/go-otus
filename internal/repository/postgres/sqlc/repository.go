package sqlc

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"os"
	"otus/internal/model"
	"otus/internal/model/catalog"
	"otus/pkg/lib/mapstructure"
)

const AccountType = 1
const SerialType = 2
const SeasonType = 3
const EpisodeType = 4

func NewRepository[T catalog.HasId]() *Repository[T] {
	db, err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}

	repo := New(db)

	return &Repository[T]{
		repo: repo,
	}
}

func (r *Repository[T]) Save(entity T) error {
	var err error

	if entity.GetId() == 0 {
		err = r.add(entity)
	} else {
		err = r.update(entity)
	}

	return err
}

func (r *Repository[T]) add(entity T) error {
	var err error
	var typeId int

	m, _ := mapstructure.StructToMap(entity)
	if typeId, err = resolveTypeIdByEntityType[T](); err == nil {
		switch typeId {
		case 2:
			id, err := r.repo.AddSerial(context.Background(), AddSerialParams{
				Sort: sql.NullInt32{
					Int32: int32(m["sort"].(float64)),
					Valid: true,
				},
				Title: sql.NullString{
					String: m["title"].(string),
					Valid:  true,
				},
				ProductionPeriod: sql.NullString{
					String: m["production_period"].(string),
					Valid:  true,
				},
				Rating: sql.NullFloat64{
					Float64: m["rating"].(float64),
					Valid:   true,
				},
				Quality: sql.NullString{
					String: m["quality"].(string),
					Valid:  true,
				},
				Duration: sql.NullFloat64{
					Float64: m["duration"].(float64),
					Valid:   true,
				},
			})

			if err == nil {
				item, _ := r.GetById(int(id))
				m, _ := mapstructure.StructToMap(item)
				mapstructure.MapToStruct(m, &entity)
			}
		case 1, 3, 4:
			// TODO
		}
	}

	return err
}

func (r *Repository[T]) update(entity T) error {
	var err error
	var typeId int
	var id int

	id = entity.GetId()

	_, err = r.GetById(id)
	if err == nil {
		if typeId, err = resolveTypeIdByEntityType[T](); err == nil {
			switch typeId {
			case 2:
				m, _ := mapstructure.StructToMap(entity)

				err = r.repo.UpdateSerial(context.Background(), UpdateSerialParams{
					ID: int32(entity.GetId()),
					Sort: sql.NullInt32{
						Int32: int32(m["sort"].(float64)),
						Valid: true,
					},
					Title: sql.NullString{
						String: m["title"].(string),
						Valid:  true,
					},
					ProductionPeriod: sql.NullString{
						String: m["production_period"].(string),
						Valid:  true,
					},
					Rating: sql.NullFloat64{
						Float64: m["rating"].(float64),
						Valid:   true,
					},
					Quality: sql.NullString{
						String: m["quality"].(string),
						Valid:  true,
					},
					Duration: sql.NullFloat64{
						Float64: m["duration"].(float64),
						Valid:   true,
					},
				})
			case 1, 3, 4:
				// TODO
			}
		}
	}

	return err
}

func (r *Repository[T]) Delete(id int) error {
	var err error
	var typeId int

	_, err = r.GetById(id)
	if err == nil {
		if typeId, err = resolveTypeIdByEntityType[T](); err == nil {
			switch typeId {
			case 2:
				err = r.repo.DeleteSerial(context.Background(), int32(id))
			case 1, 3, 4:
				// TODO
			}
		}
	}

	return err
}

func (r *Repository[T]) GetAll() ([]T, error) {
	var err error
	var items []T

	var rows []MoviesOnlineSerial
	rows, err = r.repo.GetAllSerial(context.Background())
	if err == nil {
		for _, item := range rows {
			entity := new(T)

			m := structToMap(item)
			mapstructure.MapToStruct(m, &entity)

			items = append(items, *entity)
		}
	}

	return items, err
}

func (r *Repository[T]) GetById(id int) (T, error) {
	var entity T
	var err error

	var i MoviesOnlineSerial
	i, err = r.repo.GetByIdSerial(context.Background(), int32(id))
	if err == nil {
		result := structToMap(i)
		mapstructure.MapToStruct(result, &entity)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("entity not found")
		}
	}

	return entity, err
}

func (r *Repository[T]) Count() int {
	items, _ := r.GetAll()
	return len(items)
}

func structToMap(i MoviesOnlineSerial) map[string]interface{} {
	var result = make(map[string]interface{})

	result["id"] = int(i.ID)

	if i.Sort.Valid {
		result["sort"] = int(i.Sort.Int32)
	}
	if i.FileID.Valid {
		result["fileId"] = int(i.FileID.Int32)
	}
	if i.Title.Valid {
		result["title"] = i.Title.String
	}
	if i.ProductionPeriod.Valid {
		result["productionPeriod"] = i.ProductionPeriod.String
	}
	if i.Rating.Valid {
		result["rating"] = i.Rating.Float64
	}
	if i.Quality.Valid {
		result["quality"] = i.Quality.String
	}
	if i.Duration.Valid {
		result["duration"] = i.Duration.Float64
	}
	if i.Description.Valid {
		result["description"] = i.Description.String
	}
	return result
}

func resolveTypeIdByEntityType[T catalog.HasId]() (int, error) {
	var entity T

	var typeId int
	switch any(entity).(type) {
	case *catalog.Serial:
		typeId = SerialType
	case *catalog.Season:
		typeId = SeasonType
	case *catalog.Episode:

		typeId = EpisodeType
	case *model.Account:

		typeId = AccountType
	default:
		return 0, errors.New("invalid entity type")
	}
	return typeId, nil
}

func dbConnect() (*sql.DB, error) {
	if db, err := sql.Open("postgres", os.Getenv("POSTGRES_DB_URL")); err == nil {
		return db, nil
	} else {
		return nil, err
	}
}
