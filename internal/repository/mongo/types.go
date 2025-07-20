package mongo

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"otus/internal/model/catalog"
)

type Repository[T catalog.HasId] struct {
	сollection  *mongo.Collection
	СlientRedis *redis.Client
}
