package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/saltchang/church-music-api/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database struct
type Database struct {
	Songs           *mongo.Collection // The collection of songs data in MongoDB
	SongsForTesting *mongo.Collection // The collection of songs data for testing in MongoDB
	Tokens          *mongo.Collection // The collection of tokens
}

// InitDB function
func (db *Database) InitDB() *Database {

	// MongoDB
	fmt.Println("Connected to MongoDB...")

	mongoURI := fmt.Sprintf("%s", env.ENV.MongoURI)

	// Make a context with timeout for 10s for create the client for MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Create the client at port:27017 (MongoDB default)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	// If it fails...
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	// Make a context with timeout for 10s for connect to MongoDB
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// Try to connect to MongoDB and catch the error if it fails
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	// If success
	fmt.Printf("Success!\n\n")

	// Get MongoDB collection "songs" from database "caten-worship" as a
	// *mongo.Collection type

	if env.ENV.AppEnv == "Development" {
		db.Songs = client.Database(env.ENV.SongsDBName).Collection(env.ENV.SongsCollectionNameForDev)
		fmt.Println("DB environment: SongsCollectionNameForDev")
	} else if env.ENV.AppEnv == "Production" {
		db.Songs = client.Database(env.ENV.SongsDBName).Collection(env.ENV.SongsCollectionName)
		fmt.Println("DB environment: SongsCollectionName")
	} else {
		fmt.Println("Error: AppEnv not definded.")
	}

	db.SongsForTesting = client.Database(env.ENV.SongsDBName).Collection(env.ENV.SongsCollectionNameForDev)
	db.Tokens = client.Database(env.ENV.SongsDBName).Collection(env.ENV.TokensCollectionName)

	cancel()

	return db
}

var (
	// DB var
	DB = new(Database)
)
