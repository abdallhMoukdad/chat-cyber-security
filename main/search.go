package main

import "fmt"

func main() {
	// Create a map list of professors and their connections
	professors := map[string]string{
		"John Smith":      "Computer Science",
		"Jane Doe":        "Mathematics",
		"Michael Johnson": "Physics",
	}

	// Search for a specific professor name
	professorName := "Jane Doedddd"
	connection, exists := professors[professorName]
	fmt.Println(connection)
	//fmt.Println(exists)
	// If the professor exists, return the name and connection
	if exists {
		fmt.Printf("Professor: %s\nConnection: %s\n", professorName, connection)
	} else {
		fmt.Println("Professor not found")
	}
}
