package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8083")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send data over the connection
	username := "exampleUser"
	password := "examplePassword"
	email := "example@example.com"

	data := fmt.Sprintf("%s,%s,%s", username, password, email)
	conn.Write([]byte(data))

	fmt.Println("Data sent:", data)
}
