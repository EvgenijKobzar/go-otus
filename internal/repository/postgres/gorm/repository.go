package gorm

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"otus/internal/model/catalog"
)

func NewRepository[T catalog.HasId]() *Repository[T] {
	db, err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	return &Repository[T]{
		db: db,
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
	result := r.db.Create(&entity)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) update(entity T) error {
	var err error

	result := r.db.Save(&entity)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("entity not found")
		}
	}
	return err
}

func (r *Repository[T]) Delete(id int) error {
	var err error
	var entity T

	result := r.db.Delete(&entity, id)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("entity not found")
		}
	}
	return err
}

func (r *Repository[T]) GetAll() ([]T, error) {
	var items []T
	var err error

	result := r.db.Find(&items)
	if result.Error != nil {
		err = result.Error
	}
	return items, err
}

func (r *Repository[T]) GetById(id int) (T, error) {
	var entity T
	var err error

	err = r.db.First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("entity not found")
	}
	return entity, err
}

func (r *Repository[T]) Count() int {
	items, _ := r.GetAll()
	return len(items)
}

func dbConnect() (*gorm.DB, error) {

	if db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DB_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "movies_online.",
		},
	}); err == nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(100)

		return db, nil
	} else {
		return nil, err
	}
}
