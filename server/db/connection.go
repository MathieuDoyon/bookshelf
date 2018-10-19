package db

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spf13/viper"
)

var Client *mongo.Client
var Database *mongo.Database

// Configure create mongo database connection
func Configure() {
	var dbHost string = viper.GetString("db_host")
	client, err := mongo.Connect(context.Background(), dbHost, nil)
	// require.NoError(t, err)
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	db := client.Database("bookshelf")

	log.Println("DB initialized")
	Client = client
	Database = db
}
