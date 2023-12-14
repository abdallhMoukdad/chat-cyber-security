package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type ChatRoom struct {
	clients map[string]net.Conn
	mu      sync.Mutex
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		clients: make(map[string]net.Conn),
	}
}

func (cr *ChatRoom) AddClient(username string, conn net.Conn) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	cr.clients[username] = conn
}

func (cr *ChatRoom) RemoveClient(username string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	delete(cr.clients, username)
}

func (cr *ChatRoom) GetClients() []net.Conn {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	var clients []net.Conn
	for _, conn := range cr.clients {
		clients = append(clients, conn)
	}
	return clients
}

func (cr *ChatRoom) CreateNewRoom() {
	// Create a new chat room
	go func() {
		newRoom := NewChatRoom()
		rooms = append(rooms, newRoom)

		for {
			// Wait for two users to join the new room
			for len(newRoom.GetClients()) < 2 {
				time.Sleep(1 * time.Second)
			}

			// Start handling the new chat room
			go handleRoom(newRoom)
		}
	}()
}

func handleRoom(chatRoom *ChatRoom) {
	clients := chatRoom.GetClients()

	// Notify both clients that the chat room is active
	for _, conn := range clients {
		conn.Write([]byte("Chat room is now active. Start chatting!\n"))
	}

	// Start relaying messages between the two clients
	for {
		message, err := bufio.NewReader(clients[0]).ReadString('\n')
		if err != nil {
			break
		}

		message = strings.TrimSpace(message)
		if message == "exit" {
			break
		}

		// Relay the message to the other client
		otherClient := clients[1]
		otherClient.Write([]byte(fmt.Sprintf("[%s]: %s\n", "User", message)))
	}
}

var rooms []*ChatRoom

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("A new user connected.")

	// Ask for username
	conn.Write([]byte("Enter your username: "))
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
	username = strings.TrimSpace(username)

	// Create a new chat room for the user
	currentRoom := NewChatRoom()
	rooms = append(rooms, currentRoom)

	// Add the user to the current chat room
	currentRoom.AddClient(username, conn)

	// Notify the user that they are waiting for another participant
	conn.Write([]byte("Waiting for another participant...\n"))

	// Check if there is another user in the same chat room
	for len(currentRoom.GetClients()) < 2 {
		time.Sleep(1 * time.Second)
	}

	// Start handling the chat room
	go handleRoom(currentRoom)
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
