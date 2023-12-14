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
func (cr *ChatRoom) GetClientsNames() []string {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	var namesClients []string
	for names, _ := range cr.clients {
		namesClients = append(namesClients, names)
	}
	return namesClients
}

var rooms []*ChatRoom

func createNewRoom() {
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
}
func (cr *ChatRoom) Broadcast(sender, message string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	//var names = cr.GetClientsNames()

	for username, conn := range cr.clients {
		if username != sender {
			conn.Write([]byte(fmt.Sprintf("[%s]: %s\n", sender, message)))
		}
	}
}

func handleRoom(chatRoom *ChatRoom, conn net.Conn) {
	clients := chatRoom.GetClients()
	clientsNames := chatRoom.GetClientsNames()
	// Notify both clients that the chat room is active
	// Start relaying messages between the two clients
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			break
		}

		message = strings.TrimSpace(message)
		chatRoom.Broadcast(clientsNames[0], message)

		if message == "exit" {
			break
		}

		//Relay the message to the other client
		otherClient := clients[1]
		otherClient.Write([]byte(fmt.Sprintf("[%s]: %s\n", "User", message)))
	}
}

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
	fmt.Printf("Client %s connected.\n", username)

	// Check if there's an existing room with one user
	var existingRoom *ChatRoom
	for _, room := range rooms {
		if len(room.GetClients()) == 1 {
			existingRoom = room
			break
		}
	}

	if existingRoom == nil {
		// Create a new room for this user
		existingRoom = NewChatRoom()
		rooms = append(rooms, existingRoom)
	}

	// Add the user to the existing room
	existingRoom.AddClient(username, conn)

	// Notify the user that they are waiting for another participant
	conn.Write([]byte("Waiting for another participant...\n"))

	// Check if there is another user in the same chat room
	for len(existingRoom.GetClients()) < 2 {
		time.Sleep(1 * time.Second)
	}

	// Start handling the chat room
	go handleRoom(existingRoom, conn)
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	// Start creating new rooms in the background
	go createNewRoom()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
