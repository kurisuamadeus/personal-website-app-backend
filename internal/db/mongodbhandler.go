package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const username string = "Admin"
const password string = "zWJ9NnC8iWFL8PqR"

var DB *mongo.Client

func MongoDB() {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.h5l4h.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", username, password))
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		fmt.Println(err)
	}
	// test := client.Database("db_personalwebsiteapp").Collection("page_data").FindOne(context.TODO(), bson.D{{"lang", "en"}})
	// value, err := test.Raw()
	// fmt.Println(value)
	DB = client

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}
