package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Post struct {
	ID      string    `bson:"_id"`
	Author  string    `bson:"author"`
	Content string    `bson:"content"`
	Created time.Time `bson:"created"`
}

type PostsService interface {
	GetLatestsPosts() ([]Post, error)
}

type MongoDbService struct {
	client *mongo.Client
}

func NewMongoDbService(connectionString string) (*MongoDbService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return &MongoDbService{client: client}, nil
}

func (mongo *MongoDbService) GetLatestsPosts() ([]Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.client.Database("blog").Collection("posts")

	opts := options.Find().SetSort(map[string]int{"created": -1}).SetLimit(10)

	cursor, err := collection.Find(ctx, bson.D{}, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []Post
	for cursor.Next(ctx) {
		var post Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
