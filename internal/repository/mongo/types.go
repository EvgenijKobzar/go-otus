package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"otus/internal/model/catalog"
)

type Repository[T catalog.HasId] struct {
	сollection *mongo.Collection
}
