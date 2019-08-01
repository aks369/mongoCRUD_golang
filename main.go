package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	client := connection()
	collection := client.Database("<database>").Collection("<collection-name>")
	//Replace '<database-name>', '<collection-name>'

	var choice string
	fmt.Println("\nWhat operation do you want to perform? \n  (insert / delete / update / get / exit)")
	fmt.Print("Enter one of the above options: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid format;", err)
		os.Exit(0)
	}

	choice = strings.ToLower(choice)

	switch choice {
	case "insert":
		insert(client, collection)

	case "delete":
		delete(client, collection)

	case "update":
		update(client, collection)

	case "get":
		get(client, collection)

	case "exit":
		os.Exit(0)

	default:
		fmt.Println("invalid input\n use appropriate values")
	}
}
