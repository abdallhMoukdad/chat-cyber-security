package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
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

func (cr *ChatRoom) Broadcast(sender, message string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	for username, conn := range cr.clients {
		if username != sender {
			conn.Write([]byte(fmt.Sprintf("[%s]: %s\n", sender, message)))
		}
	}
}

func handleConnection(username string, conn net.Conn, chatRoom *ChatRoom) {
	defer func() {
		fmt.Printf("Client %s disconnected.\n", username)
		chatRoom.RemoveClient(username)
		conn.Close()
	}()

	fmt.Printf("Client %s connected.\n", username)
	chatRoom.AddClient(username, conn)

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		message = strings.TrimSpace(message)
		if message == "exit" {
			return
		}

		fmt.Printf("[%s]: %s\n", username, message)
		chatRoom.Broadcast(username, message)
	}
}

func main() {
	chatRoom := NewChatRoom()

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

		go func(conn net.Conn) {
			defer conn.Close()

			// Ask for username
			conn.Write([]byte("Enter your username: "))
			username, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading username:", err)
				return
			}
			username = strings.TrimSpace(username)

			handleConnection(username, conn, chatRoom)
		}(conn)
	}
}
