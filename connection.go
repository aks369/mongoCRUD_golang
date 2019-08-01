package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connection() (client *mongo.Client) {
	conn := options.Client().ApplyURI("mongodb+srv://<username>:<password>@<clustername>.mongodb.net/test")
	//Replace '<username>', '<password>', and '<clustername>'

	client, err := mongo.Connect(context.TODO(), conn)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB !!!")

	return
}
