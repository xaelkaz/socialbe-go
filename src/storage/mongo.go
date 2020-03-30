package storage

import (
	"context"
	"social-golang/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Database *mongo.Database
var Comment *mongo.Collection
var Faculty *mongo.Collection
var FavoriteSubject *mongo.Collection
var Feedback *mongo.Collection
var Major *mongo.Collection
var Post *mongo.Collection
var Request *mongo.Collection
var Subject *mongo.Collection
var User *mongo.Collection

func Init() {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(config.ServerConfig.Storage.Uri)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		log.Println("Connect to MongoDB fail")
		return
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	Database = client.Database(config.ServerConfig.Storage.Name)
	initCollection()
}

func initCollection() {
	Comment = Database.Collection("Comment")
	Faculty = Database.Collection("Faculty")
	FavoriteSubject = Database.Collection("FavoriteSubject")

	Feedback = Database.Collection("Feedback")
	Major = Database.Collection("Major")
	Post = Database.Collection("Post")
	Request = Database.Collection("Request")
	Subject = Database.Collection("Subject")
	User = Database.Collection("User")
}
