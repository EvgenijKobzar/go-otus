package gorm

import (
	"gorm.io/gorm"
	"otus/internal/model/catalog"
)

type Repository[T catalog.HasId] struct {
	db *gorm.DB
}
