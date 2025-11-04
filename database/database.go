package database

import (
	"context"
	"errors"
	"url-shortener/models/url"
	"url-shortener/service"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage interface {
	Store(ctx context.Context, value *url.ShortUrl) error
	FindByCode(ctx context.Context, code string) (*url.ShortUrl, error)
	GenerateId(ctx context.Context) int
	Exists(ctx context.Context, code string) bool
}

type MemoryStorage struct {
	data map[string]*url.ShortUrl
}

// Exists implements Storage.
func (memoryStorage *MemoryStorage) Exists(ctx context.Context, code string) bool {
	panic("unimplemented")
}

// GenerateId implements Storage.
func (memoryStorage *MemoryStorage) GenerateId(ctx context.Context) int {
	return 1
}

type MongoStorage struct {
	collection *mongo.Collection
}

func NewMemoryStorage() Storage {
	return &MemoryStorage{
		data: make(map[string]*url.ShortUrl),
	}
}

func NewMongoStorage(dbName string, collectionName string, host string, port int) (*MongoStorage, error) {
	client, err := service.NewMongoConnection(host, port)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoStorage{
		collection: collection,
	}, nil
}

func (mongoStorage *MongoStorage) Store(ctx context.Context, value *url.ShortUrl) error {

	_, err := mongoStorage.collection.InsertOne(ctx, value)
	return err
}

func (mongoStorage *MongoStorage) FindByCode(ctx context.Context, code string) (*url.ShortUrl, error) {

	findRes := mongoStorage.collection.FindOne(ctx, bson.D{
		{"shortcode", code},
	})
	
	url := &url.ShortUrl{}
	err := findRes.Decode(url)

	if err != nil {
		return nil, err
	}

	return url, nil
}

func (mongoStorage *MongoStorage) Exists(ctx context.Context, code string) bool {

	filter := bson.M{"shortcode": code}

	cursor, err := mongoStorage.collection.Find(ctx, filter)

	if err != nil {
		panic("error requesting resource")
	}

	return cursor.TryNext(ctx)
}

func (mongoStorage *MongoStorage) GenerateId(ctx context.Context) int {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"uniq_number", -1}}).SetLimit(1)
	cursor, err := mongoStorage.collection.Find(ctx, filter, opts)

	if err != nil {
		return -1
	}

	if !cursor.TryNext(ctx) {
		return 1
	}

	var docs []url.ShortUrl
	cursor.All(ctx, &docs)
	cursor.Close(ctx)

	return docs[0].Id + 1
}

func (memoryStorage *MemoryStorage) Store(ctx context.Context, value *url.ShortUrl) error {
	memoryStorage.data[value.ShortCode] = value
	return nil
}

func (memoryStorage *MemoryStorage) FindByCode(ctx context.Context, code string) (*url.ShortUrl, error) {

	val, found := memoryStorage.data[code]

	if !found {
		return nil, errors.New("not found")
	}

	return val, nil
}
