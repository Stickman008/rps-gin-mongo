package database

import (
	"context"
	"log"
	"rock-paper-scissor/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient interface {
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database() *mongo.Database
}

type Config struct {
	Host string
	Db   string
}

type mongoClient struct {
	client *mongo.Client
	config Config
}

func Init() {
	config := Config{Host: config.DB.Host, Db: config.DB.Name}
	client := NewMongoClient(config)
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Mongodb connected!! 🎉")
}

// newMongoClient initial mongo connection
func NewMongoClient(c Config) MongoClient {
	clientOptions := options.Client().ApplyURI(c.Host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// return client.Database(c.Db)
	return &mongoClient{
		client: client,
		config: c,
	}
}

func (c *mongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.client.Ping(ctx, nil)
}

func (c *mongoClient) Database() *mongo.Database {
	return c.client.Database(c.config.Db)
}
