package mongorm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PrinceNarteh/go-events-api/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DBConfigs *dbConfigs

type dbConfigs struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvConfigs.MongoURI))
	if err != nil {
		log.Fatal("Error connecting database")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully!")

	DBConfigs = &dbConfigs{
		Client:   client,
		Database: client.Database(configs.EnvConfigs.DBName),
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return DBConfigs.Database.Collection(collectionName)
}
