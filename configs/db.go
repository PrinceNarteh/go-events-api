package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PrinceNarteh/go-events-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoURI = os.Getenv("MONGO_URI")

type DB struct {
	client   *mongo.Client
	database *mongo.Database
}

func InitDB() *DB {
	utils.LoadEnv()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Error connecting database")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully!")
	return &DB{
		client:   client,
		database: client.Database(os.Getenv("DB_NAME")),
	}
}

func (db *DB) GetCollection(collectionName string) *mongo.Collection {
	return db.database.Collection(collectionName)
}
