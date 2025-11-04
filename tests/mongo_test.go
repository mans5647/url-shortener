package tests

// import (
// 	"context"
// 	"testing"
// 	"url-shortener/models/url"
// 	"url-shortener/service"
// 	"go.mongodb.org/mongo-driver/v2/bson"
// 	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
// )

// func TestConnection(t * testing.T) {

// 	ctx := context.Background()

// 	client, err := service.NewMongoConnection("localhost", 27017)
// 	if err != nil {
// 		t.Error()
// 	}

// 	defer func() {
//     if err := client.Disconnect(ctx); err != nil {
//         panic(err)
//     }
// 	}()

// 	if client.Ping(ctx, readpref.Primary()) != nil {
// 		t.Error()
// 	}

// }

// func TestInsertOne(t * testing.T) {
// 	ctx := context.Background()

// 	oldUrl := url.PlainUrl{
// 		Url: "https://www.google.com",
// 	}

// 	shortener := url.SequenceShortener{}

// 	newUrl, err := shortener.Shorten(ctx, url.DefaultCutSize, &oldUrl)

// 	if err != nil {
// 		t.Error()
// 	}

// 	client, err := service.NewMongoConnection("localhost",27017)

// 	if err != nil {
// 		t.Errorf("failed to connect to mongo!")
// 	}

// 	defer func() {
//     if err := client.Disconnect(ctx); err != nil {
//         panic(err)
//     }
// 	}()

// 	if client.Ping(ctx, readpref.Primary()) != nil {
// 		t.Error()
// 	}

// 	db := client.Database("url_shortener")

// 	coll := db.Collection("urls")


// 	_, err = coll.InsertOne(ctx, newUrl)

// 	if err != nil {
// 		t.Errorf("insert one error!")
// 	}

// 	_, err = coll.DeleteMany(ctx,bson.D{{"_id", newUrl.Id},})

// 	if err != nil {
// 		t.Error("error clearing entries")
// 	}

// }

// func TestInsertMany(t * testing.T) {

// 	ctx := context.Background()

// 	oldUrls := []url.PlainUrl{
// 		{Url: "https://www.google.com"},
// 		{Url: "https://yandex.ru"},
// 		{Url: "https://https://www.mongodb.com/docs/drivers/go/"},
// 	}

// 	shortener := url.SequenceShortener{}

// 	client, err := service.NewMongoConnection("localhost",27017)

// 	if err != nil {
// 		t.Errorf("failed to connect to mongo!")
// 	}

// 	defer func() {
//     if err := client.Disconnect(ctx); err != nil {
//         panic(err)
//     }
// 	}()

// 	if client.Ping(ctx, readpref.Primary()) != nil {
// 		t.Error()
// 	}

// 	db := client.Database("url_shortener")

// 	coll := db.Collection("urls")

// 	newUrls := []*url.ShortUrl{}
// 	for _, old := range oldUrls {
// 		newUrl, err := shortener.Shorten(ctx, url.DefaultCutSize, &old)

// 		if err != nil {
// 			t.Errorf("error in shortening")
// 		}

// 		newUrls = append(newUrls, newUrl)
// 	}

// 	_, err = coll.InsertMany(ctx, newUrls)

// 	if err != nil {
// 		t.Errorf("InsertMany() error")
// 	}

// 	_, err = coll.DeleteMany(ctx, bson.D{})

// 	if err != nil {
// 		t.Errorf("error clearing entries")
// 	}

// }