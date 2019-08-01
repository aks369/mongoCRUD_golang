package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func delete(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var person user
	fmt.Print("Enter the _id: ")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)
	fmt.Println(person, "  : deleted successfully !!")
}
