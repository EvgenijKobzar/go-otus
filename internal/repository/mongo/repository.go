package mongo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"otus/internal/model"
	"otus/internal/model/catalog"
	"strconv"
	"time"
)

const AddAction = "add"
const UpdateAction = "update"
const DeleteAction = "delete"

func NewRepository[T catalog.HasId]() *Repository[T] {
	сollection, err := getDBCollection[T]()
	if err != nil {
		log.Fatal(err)
	}

	client, err := getClientRedis()
	if err != nil {
		log.Fatal(err)
	}

	return &Repository[T]{
		сollection:  сollection,
		СlientRedis: client,
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
	ctx := context.TODO()
	entity.SetId(r.Count() + 1)
	result, _ := r.сollection.InsertOne(ctx, entity)
	insertedID := result.InsertedID

	err := r.сollection.FindOne(
		context.TODO(),
		bson.M{"_id": insertedID},
	).Decode(&entity)

	if err == nil {
		err = r.addLogMessage(entity, AddAction)
	}

	return err
}

func (r *Repository[T]) update(entity T) error {
	ctx := context.TODO()
	filter := bson.M{"_id": entity.GetId()}
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After). // Возвращаем документ ПОСЛЕ обновления
		SetUpsert(false)                  // Не создавать новый, если не найден

	update := bson.M{
		"$set": entity,
	}

	err := r.сollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errors.New("entity not found")
		}
	}

	if err == nil {
		err = r.addLogMessage(entity, UpdateAction)
	}

	return err
}

func (r *Repository[T]) Delete(id int) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var entity T
	err := r.сollection.FindOneAndDelete(ctx, filter).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errors.New("entity not found")
		}
	}

	if err == nil {
		err = r.addLogMessage(entity, DeleteAction)
	}

	return err
}

func (r *Repository[T]) GetAll() ([]T, error) {
	var items []T

	ctx := context.TODO()
	cursor, err := r.сollection.Find(
		ctx,
		bson.M{})
	defer cursor.Close(ctx)

	if err == nil {
		err = cursor.All(ctx, &items)
	}

	return items, err
}

func (r *Repository[T]) GetById(id int) (T, error) {
	var entity T
	var err error

	err = r.сollection.FindOne(
		context.TODO(),
		bson.M{"_id": id},
	).Decode(&entity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errors.New("entity not found")
		}
	}

	return entity, err
}

func (r *Repository[T]) Count() int {
	items, _ := r.GetAll()
	return len(items)
}

func resolveCollectionNameByEntityType[T catalog.HasId]() (string, error) {
	var entity T

	var name string
	switch any(entity).(type) {
	case *catalog.Serial:
		name = "serial"
	case *catalog.Season:
		name = "season"
	case *catalog.Episode:

		name = "episode"
	case *model.Account:

		name = "account"
	default:
		return "", errors.New("invalid entity type")
	}
	return name, nil
}

func dbConnect() (*mongo.Database, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(os.Getenv("MONGO_DB_NAME")), nil
}

func getDBCollection[T catalog.HasId]() (*mongo.Collection, error) {
	var err error
	var db *mongo.Database
	var collectionName string

	if collectionName, err = resolveCollectionNameByEntityType[T](); err == nil {
		if db, err = dbConnect(); err == nil {
			return db.Collection(collectionName), nil
		}
	}

	return nil, err
}

func getClientRedis() (*redis.Client, error) {
	ctx := context.Background()

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB_NAME"))
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func makeKey[T catalog.HasId](entity catalog.HasId, action string) (string, error) {
	now := time.Now().String()
	buffer := bytes.Buffer{}
	if name, err := resolveCollectionNameByEntityType[T](); err == nil {
		var id int
		if entity.GetId() == 0 {
			id = 0
		} else {
			id = entity.GetId()
		}

		buffer.WriteString(action)
		buffer.WriteString(":")
		buffer.WriteString(name)
		buffer.WriteString(":")
		buffer.WriteString(strconv.Itoa(id))
		buffer.WriteString(":")
		buffer.WriteString(now)

		return buffer.String(), nil

	} else {
		return "", err
	}
}

func (r *Repository[T]) addLogMessage(entity catalog.HasId, action string) error {
	ctx := context.Background()

	var err error
	var key string
	key, err = makeKey[T](entity, action)
	if err == nil {

		jsonData, err := json.Marshal(entity)
		if err == nil {
			err = r.СlientRedis.Set(ctx, key, jsonData, time.Second*10).Err()
		}
	}
	return err
}
