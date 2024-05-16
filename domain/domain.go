package domain

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Post struct {
	ID            string    `bson:"_id"`
	Title         string    `bson:"Title"`
	Content       string    `bson:"Content"`
	Description   string    `bson:"Description"`
	Created       time.Time `bson:"Created"`
	Tags          []string  `bson:"Tags"`
	ImageLocation string    `bson:"ImageLocation"`
	Author        string    `bson:"Author"`
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

func (mongo *MongoDbService) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mongo.client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

func (mongo *MongoDbService) GetPosts() ([]Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.client.Database("blog").Collection("posts")

	opts := options.Find().SetSort(map[string]int{"created": -1})

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

func (mongo *MongoDbService) GetPostById(id string) (Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongo.client.Database("blog").Collection("posts")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid ID format: ", id)
		return Post{}, err
	}
	filter := bson.M{"_id": objectID}

	var post = Post{}

	err = collection.FindOne(ctx, filter).Decode(&post)

	if err != nil {
		fmt.Println("Could not decode post for id: ", id)
		return Post{}, err
	}

	return post, nil
}
