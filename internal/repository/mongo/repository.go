package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"otus/internal/model"
	"otus/internal/model/catalog"
)

func NewRepository[T catalog.HasId]() *Repository[T] {
	return &Repository[T]{
		items:      make(map[int]T),
		Collection: getDBCollection[T](),
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
	result, _ := r.Collection.InsertOne(ctx, entity)
	insertedID := result.InsertedID

	err := r.Collection.FindOne(
		context.TODO(),
		bson.M{"_id": insertedID},
	).Decode(&entity)
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

	err := r.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errors.New("entity not found")
		}
	}
	return err
}

func (r *Repository[T]) Delete(id int) error {
	ctx := context.TODO()
	filter := bson.M{"_id": id}
	var entity T
	err := r.Collection.FindOneAndDelete(ctx, filter).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errors.New("entity not found")
		}
	}

	return err
}

func (r *Repository[T]) GetAll() ([]T, error) {
	r.imx.RLock()
	defer r.imx.RUnlock()

	var items []T

	ctx := context.TODO()
	cursor, err := r.Collection.Find(
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

	err = r.Collection.FindOne(
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

	// Пинг сервера для проверки соединения
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err

	}
	fmt.Println("Подключено к MongoDB!")

	// Создание или переключение на базу данных
	return client.Database(os.Getenv("MONGO_DB_NAME")), nil
}

func getDBCollection[T catalog.HasId]() *mongo.Collection {
	var err error
	var db *mongo.Database
	var collectionName string

	if collectionName, err = resolveCollectionNameByEntityType[T](); err == nil {
		if db, err = dbConnect(); err == nil {
			return db.Collection(collectionName)
		}
	}
	log.Fatal(err)
	return nil
}
