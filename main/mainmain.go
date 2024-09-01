package main

import (
	"awesomeProject1/ent"
	"awesomeProject1/ent/professor"
	"awesomeProject1/ent/student"
	_ "awesomeProject1/ent/student"
	"bufio"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq" // add this
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type User struct {
	Connection net.Conn
	Username   string
}

var connectedUsers []User

func searchUser(username string) *User {
	for _, user := range connectedUsers {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

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
		//go handleRoom(newRoom)
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

func handleConnection(conn net.Conn, client *ent.Client) {
	defer conn.Close()
	fmt.Println("A new user connected.")
	role, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	if role == "0" {

	} else {
		username, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading username:", err)
			return
		}
		username = strings.TrimSpace(username)
		user := User{Connection: conn, Username: username}
		connectedUsers = append(connectedUsers, user)
		password, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading password:", err)
			return
		}
		password = strings.TrimSpace(password)
		if exsit, _ := QueryUser(context.Background(), client, username); exsit {
			//if hisInfoCompleted() {
			//
			//} else {
			//	conn.Write([]byte("enter your phone number" + "\n"))
			//
			//}

		} else {
			username, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading username:", err)
				return
			}
			username = strings.TrimSpace(username)
			user := User{Connection: conn, Username: username}
			connectedUsers = append(connectedUsers, user)
			password, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading password:", err)
				return
			}
			password = strings.TrimSpace(password)

		}
	}
	// Ask for username
	//conn.Write([]byte("Enter your username: "))
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
	username = strings.TrimSpace(username)
	user := User{Connection: conn, Username: username}
	connectedUsers = append(connectedUsers, user)

	fmt.Printf("Client %s connected.\n", username)
	//conn.Write([]byte("Enter the professor name: "))
	proName, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
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

	proName = strings.TrimSpace(proName)
	//fmt.Printf("Client %s connected.\n", username)
	// Check if there's an existing room with one user
	professorObject := searchUser(proName)
	if professorObject != nil {
		fmt.Println("the pro found yahhhhhhhhhhhh")
		conn.Write([]byte("the professor  exsits  ...\n"))

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
		existingRoom.AddClient(professorObject.Username, professorObject.Connection)
		conn.Write([]byte("you can chat now ...\n"))

	} else {
		fmt.Println("the professor not exsits yet or not contected ")
		conn.Write([]byte("the professor not exsits yet or not contected...\n"))

	}

	// Notify the user that they are waiting for another participant
	conn.Write([]byte("Waiting for another participant...\n"))

	// Check if there is another user in the same chat room
	for len(existingRoom.GetClients()) < 2 {
		time.Sleep(1 * time.Second)
	}
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
		existingRoom.Broadcast(username, message)
	}

	// Start handling the chat room
	//go handleRoom(existingRoom, conn)
}

func main() {
	client, err := ent.Open(dialect.Postgres, "user=postgres password=postgres dbname=ISS sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	listener, err := net.Listen("tcp", ":8082")
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

		go handleConnection(conn, client)
	}

}
func QueryUser(ctx context.Context, client *ent.Client, userName string) (bool, error) {
	u, err := client.Student.
		Query().
		Where(student.Name("abd")).Exist(ctx)
	// `Only` fails if no user found,
	// or more than 1 user returned.
	//Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
func RegisterStudent(client *ent.Client, username, password, email, nationalNumber, home_loc, phoneNumber string) error {

	_, err := client.Student.
		Create().SetName(username).
		SetPassword(password).SetNationalNumber(nationalNumber).
		SetHomeLocation(home_loc).SetPhoneNumber(phoneNumber).
		Save(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return err
	}
	return nil
}
func RegisterProfessor(client *ent.Client, username, password, email, nationalNumber, home_loc, phoneNumber string) error {

	_, err := client.Professor.
		Create().SetName(username).
		SetPassword(password).
		Save(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return err
	}
	return nil
}
func QueryProfessor(ctx context.Context, client *ent.Client) (bool, error) {
	u, err := client.Professor.
		Query().
		Where(professor.Name("abd")).Exist(ctx)
	// `Only` fails if no user found,
	// or more than 1 user returned.
	//Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
