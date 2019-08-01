package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	Name string
	Age  int
	City string
	Dob  string
}

func insert(client *mongo.Client, collection *mongo.Collection) {
	person := createStud()

	result, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		fmt.Println("error in insertion", err)
	}
	fmt.Println("Insert successful with id:", result.InsertedID)
}

func createStud() interface{} {
	var name, dob, city string
	var age int

	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter age: ")
	fmt.Scan(&age)
	fmt.Print("Enter date of birth (dd-mm-yyyy format): ")
	fmt.Scan(&dob)
	fmt.Print("Enter your City: ")
	fmt.Scan(&city)

	person := user{Name: name, Age: age, Dob: dob, City: city}

	return person
}
