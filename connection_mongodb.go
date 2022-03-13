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

	clientOptions.SetAuth(options.Credential{Username: Godotenv("logger_user"), Password: Godotenv("logger_password")})
	clientOptions.ApplyURI(Godotenv("logger_url"))
	client, err := mongo.Connect(ctx, &clientOptions)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return client.Database(Godotenv("logger_database"))
}
