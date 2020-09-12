package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"demo-company/config"
)

var (
	DB *mongo.Database
	client *mongo.Client
)

// Connect ...
func Connect(database string) *mongo.Database{
	envVars :=config.GetEnv()

	// connect to database
	client,err :=mongo.NewClient(options.Client().ApplyURI(envVars.Database.URI))
	
	// err 
	if err != nil {
		log.Fatal("Cannot connect to database:",err)
	}

	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err !=nil {
		log.Fatal("Cannot connect ",err)	
	}
	
	DB = client.Database(database)
	fmt.Println("Database connected",database)
	return DB
}

