package goutils

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectionMongoDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var max uint64
	max = 10
	clientOptions := options.ClientOptions{MaxPoolSize: &max}

	user := Godotenv("mongo_user")
	if user == "" {
		user = Godotenv("logger_user")
	}
	password := Godotenv("mongo_password")
	if password == "" {
		password = Godotenv("logger_password")
	}
	url := Godotenv("mongo_url")
	if url == "" {
		url = Godotenv("logger_url")
	}
	database := Godotenv("mongo_database")
	if database == "" {
		database = Godotenv("logger_database")
	}

	Godotenv("logger_password")
	Godotenv("logger_url")
	Godotenv("logger_database")

	clientOptions.SetAuth(options.Credential{Username: user, Password: password})
	clientOptions.ApplyURI(url)
	client, err := mongo.Connect(ctx, &clientOptions)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return client.Database(database)
}
