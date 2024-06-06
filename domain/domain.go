package domain

import (
	"context"
	"embed"
	"io/fs"
	"time"
"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
)

//go:embed creds.yaml
var configFile embed.FS

type Config struct {
	MongoDb struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Url      string `yaml:"url"`
	} `yaml:"mongodb"`
}

func ReadConfig() Config {
	var config Config
	yamlFile, err := fs.ReadFile(configFile, "creds.yaml")

	if err != nil {
		log.Println("Error reading cr.yaml", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Println("Could not unmarshal cr.yaml", err)
	}

	return config

}

type Post struct {
	ID            primitive.ObjectID    `bson:"_id, omitempty"`
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

func NewMongoDbService() (*MongoDbService, error) {
	var config = ReadConfig();

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mongoConnection = "mongodb+srv://"+config.MongoDb.Username+":"+config.MongoDb.Password+"@"+config.MongoDb.Url


	log.Println("MongoConnection: ", mongoConnection)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnection))

	if err != nil {
		log.Printf("Could not connect to mongo: ", err)
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("Could not ping  mongo: ", err)
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

func(mongo *MongoDbService) InsertPost(post Post) (interface{}, error){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongo.client.Database("blog").Collection("posts")
	result, error:=	 collection.InsertOne(ctx,  post)
	return result.InsertedID, error
	
}

func (mongo *MongoDbService) GetPostsByTag(tag string) ([]Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := mongo.client.Database("blog").Collection("posts")

	filter := bson.M{"Tags": bson.M{"$in": []string{tag}}}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println("Could not find a collection for a tag", tag)
		return nil, err
	}

	defer cursor.Close(ctx)

	var posts []Post
	for cursor.Next(ctx) {

		var post Post

		err = cursor.Decode(&post)
		if err != nil {
			log.Println("Could not decode a tag result", tag)
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
		log.Println("Invalid ID format: ", id)
		return Post{}, err
	}
	filter := bson.M{"_id": objectID}

	var post = Post{}

	err = collection.FindOne(ctx, filter).Decode(&post)

	if err != nil {
		log.Println("Could not decode post for id: ", id)
		return Post{}, err
	}

	return post, nil
}
