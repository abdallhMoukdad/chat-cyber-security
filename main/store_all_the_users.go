package main

import (
	"fmt"
	"net"
)

type User struct {
	Connection net.Conn
	Username   string
}

var connectedUsers []User

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	user := User{Connection: conn}
	connectedUsers = append(connectedUsers, user)

	fmt.Println("New user connected")

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		message := string(buffer[:n])

		// Handle the message received from the user

		// Example:
		// if message == "disconnect" {
		//     disconnectUser(user)
		//     break
		// }

		fmt.Printf("Received message from %s: %s\n", user.Username, message)
	}
}

func disconnectUser(user User) {
	for i, u := range connectedUsers {
		if u.Connection == user.Connection {
			connectedUsers = append(connectedUsers[:i], connectedUsers[i+1:]...)
			fmt.Printf("%s disconnected\n", u.Username)
			break
		}
	}
}
