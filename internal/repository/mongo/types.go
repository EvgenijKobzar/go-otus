package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"otus/internal/model/catalog"
	"sync"
)

type Repository[T catalog.HasId] struct {
	imx        sync.RWMutex
	items      map[int]T
	nextId     int
	Collection *mongo.Collection
}
