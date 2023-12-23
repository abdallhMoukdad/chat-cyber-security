package main

import (
	"awesomeProject1/enc"
	"awesomeProject1/ent"
	"awesomeProject1/ent/professor"
	"awesomeProject1/ent/student"
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

func main() {
	// start the tcp server
	client, err := ent.Open(dialect.Postgres, "user=postgres password=postgres dbname=ISS sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println("Server started. Waiting for connections...")

	listener, err := net.Listen("tcp", ":8082")
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

		go handleConnection(conn, client)
	}
}

func handleConnection(conn net.Conn, client *ent.Client) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	role, err := bufio.NewReader(conn).ReadString('\n')
	// fmt.Println("hey ya you now we are working")
	result := strings.Replace(role, "\n", "", -1)
	role = result
	if err != nil {
		fmt.Println("Error reading role:", err)
		return
	}
	if role == "0" {
		// for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			// break
		}

		data := strings.TrimSpace(string(buffer[:n]))
		temp := data
		if strings.HasPrefix(data, "chat") {
			temp = "chat"
		}

		switch temp {
		case "signup":
			fmt.Println("signup")
			handleSignup(conn, client)

		case "login":
			fmt.Println("login")
			handleLogin(conn, client)
		case "chat":
			fmt.Println("chat")
			// handleChat(conn)
		case "NotCompleteSignup":
			fmt.Println("NotCompleteSignup")
			handleNotCompleteSignup(conn, client)
			//handleSignup(conn, client)

		case "exit":
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		default:
			fmt.Println("Unknown command:", data)
		}
		//}
	} else {
		handleProfessorLogin(client, conn)
	}
}

func checkIfComplete(client *ent.Client, username, password string) (bool, error) {
	exsits, err := client.Student.Query().Where(student.Name(username),
		student.Password(password),
		student.NationalNumberContains("unknown")).Exist(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return exsits, err
	}
	return exsits, nil

}
func contuineSignUP(conn net.Conn, client *ent.Client, username, password string) {
	fmt.Println("Client connected:", conn.RemoteAddr())
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := string(buffer[:n])
	fmt.Println(" encrypted Data received:", data)
	decryptedData, _ := enc.GetAESDecrypted(data, "my32digitkey12345678901234567890")
	fmt.Println("  Data received:", string(decryptedData))

	// Split the values using the comma as the delimiter
	splitValues := strings.Split(string(decryptedData), ",")
	naionalNumber := splitValues[0]
	phoneNumber := splitValues[1]
	homeLocation := splitValues[2]
	exsits, err := ContuineRegisterationStudent(client, naionalNumber, homeLocation, phoneNumber, username, password)
	if err != nil && exsits == false {
		log.Printf("failed to login user: %v", err)
		conn.Write([]byte("Login failed. Invalid username or password.\n"))
		return
	} else {
		conn.Write([]byte("register successful., " + username + "!\n"))
	}

}
func handleNotCompleteSignup(conn net.Conn, client *ent.Client) {
	fmt.Println("Client connected:", conn.RemoteAddr())
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := string(buffer[:n])
	fmt.Println("Data received:", data)
	// Split the values using the comma as the delimiter
	splitValues := strings.Split(data, ",")
	username := splitValues[0]
	password := splitValues[1]
	exsits, err := RegisterNotCompleteStudent(client, username, password)
	if err != nil && exsits == false {
		log.Printf("failed to register user: %v", err)
		conn.Write([]byte("Login failed. Invalid username or password.\n"))
		return
	} else {
		conn.Write([]byte("register successful., " + username + "!\n"))
	}

}

func handleProfessorLogin(client *ent.Client, conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := string(buffer[:n])
	fmt.Println("Data received:", data)
	// Split the values using the comma as the delimiter
	splitValues := strings.Split(data, ",")

	// Assign each value to a new variable
	username := splitValues[0]
	password := splitValues[1]
	fmt.Println(username)
	fmt.Println(password)
	result := strings.Replace(username, "\n", "", -1)
	username = result
	result = strings.Replace(password, "\n", "", -1)
	password = result

	//exampleEmail := splitValues[2]
	// Read username from the client_user
	//username, err := readLine(conn)
	//if err != nil {
	//	fmt.Println("Error reading username:", err)
	//	return
	//}
	//
	//// Read password from the client_user
	//password, err := readLine(conn)
	//if err != nil {
	//	fmt.Println("Error reading password:", err)
	//	return
	//}
	exsits, err := LoginProfessor(client, username, password)
	if err != nil || exsits == false {
		log.Printf("failed to login user: %v", err)
		conn.Write([]byte("Login failed. Invalid username or password.\n"))
		return
	} else {
		user := User{Connection: conn, Username: username + "\n"}
		connectedUsers = append(connectedUsers, user)
		for _, user := range connectedUsers {
			fmt.Println("printing all the users:", user.Username)
		}

		conn.Write([]byte("Login successful. Welcome, " + username + "!\n"))
		// handleChat(conn, username)
		handleProfessorChat(conn, username)
	}
}

func LoginProfessor(client *ent.Client, username string, password string) (bool, error) {
	exsits, err := client.Professor.Query().Where(professor.Name(username), professor.Password(password)).Exist(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return exsits, err
	}
	return exsits, nil
}

func queryAllProfessor() {
}

func handleSignup(conn net.Conn, client *ent.Client) {
	fmt.Println("Client connected:", conn.RemoteAddr())
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := string(buffer[:n])
	fmt.Println("Data received:", data)
	// Split the values using the comma as the delimiter
	splitValues := strings.Split(data, ",")
	username := splitValues[0]
	password := splitValues[1]
	phoneNumber := splitValues[2]
	naionalNumber := splitValues[3]
	homeLocation := splitValues[4]
	exsits, err := RegisterStudent(client, username, password, naionalNumber, homeLocation, phoneNumber)
	if err != nil && exsits == false {
		log.Printf("failed to login user: %v", err)
		conn.Write([]byte("Login failed. Invalid username or password.\n"))
		return
	} else {
		conn.Write([]byte("register successful., " + username + "!\n"))
	}
}

func handleChat(conn net.Conn, username string) {
	// fmt.Println("enter the professor name you want to talk to :")
	// conn.Write([]byte("enter the professor name you want to talk to :\n"))

	professorName, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err)
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

	professorName = strings.TrimSpace(professorName)
	// fmt.Printf("Client %s connected.\n", username)
	// Check if there's an existing room with one user
	professorName = professorName + "\n"
	professorObject := searchUser(professorName)
	if professorObject != nil {
		fmt.Println("the pro found yahhhhhhhhhhhh")
		conn.Write([]byte("the professor  exsits  ...\n"))

		//for _, room := range rooms {
		//
		//	if len(room.GetClients()) == 1 {
		//		existingRoom = room
		//		break
		//	}
		//}

		//if existingRoom == nil {
		//	// Create a new room for this user
		//	existingRoom = NewChatRoom()
		//	rooms = append(rooms, existingRoom)
		//}

		// Add the user to the existing room
		existingRoom.AddClient(professorObject.Username, professorObject.Connection)
		conn.Write([]byte("you can chat now ...\n"))

	} else {
		fmt.Println("the professor not exsits yet or not contected ")
		conn.Write([]byte("0\n"))

		conn.Write([]byte("the professor not exsits yet or not contected...\n"))

	}

	// Notify the user that they are waiting for another participant
	conn.Write([]byte("Waiting for another participant...\n"))

	// Check if there is another user in the same chat room
	for len(existingRoom.GetClients()) < 2 {
		time.Sleep(1 * time.Second)
	}
	names := existingRoom.GetClientsNames()
	fmt.Println(names)

	reader := bufio.NewReader(conn)
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	//result := strings.Replace(username, "\n", "", -1)
	//username = result

	existingRoom.BroadcastKeys(username, string(buffer[:n]))

	for {
		//n, err := conn.Read(buffer)

		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		//message := string(buffer[:n])
		//message = strings.TrimSpace(message)
		if message == "exit" {
			return
		}
		result := strings.Replace(username, "\n", "", -1)
		username = result

		fmt.Printf("[%s]: %s\n", username, message)
		existingRoom.BroadcastKeys(username, message)
	}

	// conn.Write([]byte(professorName))
	for {

		// Read message from the client_user
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		message := strings.TrimSpace(string(buffer[:n]))

		// Handle special commands (e.g., exit)
		if strings.ToLower(message) == "exit" {
			break
		}

		// Process and broadcast the message
		// fmt.Printf("[%s]: %s\n", username, message)
	}
	for _, user := range connectedUsers {
		fmt.Println(user.Username)
	}
}

func handleProfessorChat(conn net.Conn, username string) {
	//var existingRoom *ChatRoom
	//for _, room := range rooms {
	//
	//	if len(room.GetClients()) == 1 {
	//		existingRoom = room
	//		break
	//	}
	//}
	//
	//if existingRoom == nil {
	//	// Create a new room for this user
	//	existingRoom = NewChatRoom()
	//	rooms = append(rooms, existingRoom)
	//}

	// Add the user to the existing room
	//existingRoom.AddClient(username, conn)
	//
	//professorObject := searchUser(professorName)
	//if professorObject != nil {
	//	fmt.Println("the pro found yahhhhhhhhhhhh")
	//	conn.Write([]byte("the professor  exsits  ...\n"))
	//
	//	//for _, room := range rooms {
	//	//
	//	//	if len(room.GetClients()) == 1 {
	//	//		existingRoom = room
	//	//		break
	//	//	}
	//	//}
	//
	//	//if existingRoom == nil {
	//	//	// Create a new room for this user
	//	//	existingRoom = NewChatRoom()
	//	//	rooms = append(rooms, existingRoom)
	//	//}
	//
	//	// Add the user to the existing room
	//	existingRoom.AddClient(professorObject.Username, professorObject.Connection)
	//	conn.Write([]byte("you can chat now ...\n"))
	//
	//} else {
	//	fmt.Println("the professor not exsits yet or not contected ")
	//	conn.Write([]byte("0\n"))
	//
	//	conn.Write([]byte("the professor not exsits yet or not contected...\n"))
	//
	//}
	//
	//// Notify the user that they are waiting for another participant
	//conn.Write([]byte("Waiting for another participant...\n"))

	// Check if there is another user in the same chat room
	//for len(existingRoom.GetClients()) < 2 {
	//	time.Sleep(1 * time.Second)
	//}
	time.Sleep(7 * time.Second)
	existingRoom := findTheRoom(username)
	if existingRoom == nil {
		fmt.Println("the pro is not in any room")
		return
	}

	reader := bufio.NewReader(conn)
	buffer := make([]byte, 2048)
	n, _ := conn.Read(buffer)
	fmt.Printf("[%s]: %s\n", username, n)

	existingRoom.BroadcastKeys(username, string(buffer[:n]))

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		//message = strings.TrimSpace(message)
		if message == "exit" {
			return
		}

		fmt.Printf("[%s]: %s\n", username, message)
		existingRoom.BroadcastKeys(username, message)
	}

	//
	// reader := bufio.NewReader(conn)
	// for {
	// 	message, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		return
	// 	}
	//
	// 	message = strings.TrimSpace(message)
	// 	if message == "exit" {
	// 		return
	// 	}
	//
	// 	// fmt.Printf("[%s]: %s\n", username, message)
	// 	// existingRoom.Broadcast(username, message)
	// }
	//
	// conn.Write([]byte(professorName))
	for {

		// Read message from the client_user
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		message := strings.TrimSpace(string(buffer[:n]))

		// Handle special commands (e.g., exit)
		if strings.ToLower(message) == "exit" {
			break
		}

		// Process and broadcast the message
		// fmt.Printf("[%s]: %s\n", username, message)
	}
	for _, user := range connectedUsers {
		fmt.Println(user.Username)
	}
}

func findTheRoom(username string) *ChatRoom {
	for _, room := range rooms {
		names := room.GetClientsNames()
		for _, name := range names {
			if name == username+"\n" {
				return room
			}
		}

	}
	return nil
}

// -------------------------------------------------------
var professors = []string{"Professor A", "Professor B", "Professor C"}

func sendProfessorList(conn net.Conn) {
	conn.Write([]byte("Available Professors:\n"))
	for i, professor := range professors {
		conn.Write([]byte(fmt.Sprintf("%d. %s\n", i+1, professor)))
	}
	conn.Write([]byte("Choose a professor by typing 'chat <professor-number>'.\n"))
}

func handleChat1(conn net.Conn, command string) {
	parts := strings.Fields(command)
	if len(parts) != 2 {
		conn.Write([]byte("Invalid chat command format.\n"))
		return
	}

	professorNumber := parts[1]
	professorIndex := professorNumberToInt(professorNumber)
	if professorIndex == -1 {
		conn.Write([]byte("Invalid professor number.\n"))
		return
	}

	chosenProfessor := professors[professorIndex]
	conn.Write([]byte(fmt.Sprintf("You've chosen to chat with %s.\n", chosenProfessor)))

	// Implement chat logic with the chosen professor here...
}

func professorNumberToInt(professorNumber string) int {
	index := -1
	if n, err := fmt.Sscanf(professorNumber, "%d", &index); err != nil || n != 1 || index < 1 || index > len(professors) {
		return -1
	}
	return index - 1
}

// -------------------------------------------------------
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func handleLogin(conn net.Conn, client *ent.Client) {
	fmt.Println("Client connected:", conn.RemoteAddr())
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := string(buffer[:n])
	fmt.Println("Data received:", data)
	// Split the values using the comma as the delimiter
	splitValues := strings.Split(data, ",")

	// Assign each value to a new variable
	username := splitValues[0]
	password := splitValues[1]
	//exampleEmail := splitValues[2]
	// Read username from the client_user
	//username, err := readLine(conn)
	//if err != nil {
	//	fmt.Println("Error reading username:", err)
	//	return
	//}
	//
	//// Read password from the client_user
	//password, err := readLine(conn)
	//if err != nil {
	//	fmt.Println("Error reading password:", err)
	//	return
	//}
	unComplete, err := checkIfComplete(client, username, password)

	exsits, err := LoginStudent(client, username, password)
	if err != nil || exsits == false {

		log.Printf("failed to login user: %v", err)
		conn.Write([]byte("Login failed. Invalid username or password.\n"))
		return
	} else if err != nil || unComplete == true {
		log.Printf("contuine user registreation %v", username)
		conn.Write([]byte("2\n"))

		conn.Write([]byte("contuine user registreation" + username + "!\n"))

		contuineSignUP(conn, client, username, password)
	} else {
		user := User{Connection: conn, Username: username}
		connectedUsers = append(connectedUsers, user)
		for _, user := range connectedUsers {
			fmt.Println("printing all the users:", user.Username)
		}

		log.Printf("login user done sucfully: %v", user.Username)
		conn.Write([]byte("1\n"))

		conn.Write([]byte("Login successful. Welcome, " + username + "!\n"))

		time.Sleep(1 * time.Second)
		// conn.Write([]byte("Login successful. Welcome, " + username + "!\n"))
		handleChat(conn, username)
	}

	// Perform login authentication
	//if authenticate(username, password) {
	//	conn.Write([]byte("Login successful. Welcome, " + username + "!\n"))
	//} else {
	//	conn.Write([]byte("Login failed. Invalid username or password.\n"))
	//}
}

func readLine(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func authenticate(username, password string) bool {
	// Check if the provided username and password match
	storedPassword, exists := users[username]
	return exists && storedPassword == password
}

func QueryUser(ctx context.Context, client *ent.Client, userName string) (bool, error) {
	u, err := client.Student.
		Query().
		Where(student.Name("abd")).Exist(ctx)
	// `Only` fails if no user found,
	// or more than 1 user returned.
	// Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func RegisterStudent(client *ent.Client, username, password, nationalNumber, home_loc, phoneNumber string) (bool, error) {
	_, err := client.Student.
		Create().SetName(username).
		SetPassword(password).SetNationalNumber(nationalNumber).
		SetHomeLocation(home_loc).SetPhoneNumber(phoneNumber).
		Save(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return false, err
	}
	return true, nil
}
func ContuineRegisterationStudent(client *ent.Client, nationalNumber, home_loc, phoneNumber, userName, password string) (bool, error) {
	//_, err := client.Student.
	//	Create().SetNationalNumber(nationalNumber).
	//	SetHomeLocation(home_loc).SetPhoneNumber(phoneNumber).
	//	Save(context.Background())
	//client.Student.Query().Select().
	// Assume you have a "User" model defined in your schema
	//user, err := client.Student.Query().Where(student.NameContains("dd")).Only(ctx)
	//if err != nil {
	//	// handle error
	//}
	user, err := client.Student.Query().Where(student.Name(userName),
		student.Password(password)).Only(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return false, err
	} else {
		user.Update().SetNationalNumber(nationalNumber).
			SetHomeLocation(home_loc).SetPhoneNumber(phoneNumber).Save(context.Background())

		return true, nil
	}
	//return exsits, nil

	// Update the user's fields based on certain conditions
}

func RegisterNotCompleteStudent(client *ent.Client, username, password string) (bool, error) {
	_, err := client.Student.
		Create().SetName(username).
		SetPassword(password).
		Save(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return false, err
	}
	return true, nil
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
	// Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func LoginStudent(client *ent.Client, username, password string) (bool, error) {
	//_, err := client.Student.
	//	Create().SetName(username).
	//	SetPassword(password).
	//	Save(context.Background())
	exsits, err := client.Student.Query().Where(student.Name(username), student.Password(password)).Exist(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return exsits, err
	}
	return exsits, nil
}

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
	for names := range cr.clients {
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
		// go handleRoom(newRoom)
	}
}

func (cr *ChatRoom) Broadcast(sender, message string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	// var names = cr.GetClientsNames()

	for username, conn := range cr.clients {
		if username != sender {
			conn.Write([]byte(fmt.Sprintf("[%s]: %s\n", sender, message)))
			//conn.Write([]byte(message))
		}
	}
}
func (cr *ChatRoom) BroadcastKeys(sender, message string) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	// var names = cr.GetClientsNames()

	for username, conn := range cr.clients {
		if username != sender {
			log.Printf("[%s]: %s", sender, username)
			conn.Write([]byte(message))
			//conn.Write([]byte(message))
		}
	}
}
