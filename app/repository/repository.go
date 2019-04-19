package repository

import (
	"context"
	"fmt"
	"spamtrawler/app/models"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var MongoDB *mongo.Database

func MongoGetDbConnection(configuration models.Configuration) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	MongoDB = db.Database("local")

	fmt.Println(configuration)
	fmt.Println("Connected to MongoDB!")

	return MongoDB
}
