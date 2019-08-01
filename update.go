package main

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func update(client *mongo.Client, collection *mongo.Collection) {
	var objID string
	var person user

	fmt.Print("Enter the _id: ")
	fmt.Scan(&objID)
	id, _ := primitive.ObjectIDFromHex(objID)
	result := collection.FindOne(context.TODO(), bson.M{"_id": id})
	result.Decode(&person)

	person = updateWhat(person)
	collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, bson.M{"$set": person})
	fmt.Println(person, "person updated successfully")
}

func updateWhat(person user) user {
	var name, dob, city string
	var age int
	var entry string

	fmt.Println("\nWhat do you want to Update? \n  (Name / Age / Dob / City)")
	fmt.Print("Enter one of the above options: ")
	fmt.Scan(&entry)
	entry = strings.ToLower(entry)
	fmt.Println()

	switch entry {
	case "name":
		fmt.Print("Enter new name: ")
		fmt.Scan(&name)
		person.Name = name
	case "age":
		fmt.Print("Enter new age: ")
		fmt.Scan(&age)
		person.Age = age
	case "dob":
		fmt.Print("Enter new dob (in dd-mm-yyyy format): ")
		fmt.Scan(&dob)
		person.Dob = dob
	case "city":
		fmt.Print("Enter new home: ")
		fmt.Scan(&city)
		person.City = city
	}
	return person
}
