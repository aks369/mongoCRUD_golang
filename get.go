package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func get(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var person user
	fmt.Print("Enter the _id: ")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)
	if person.Age != 0 {
		fmt.Print(person, "  : found successfully !!")
	} else {
		fmt.Print("Oops!! Search failed")
	}
}
