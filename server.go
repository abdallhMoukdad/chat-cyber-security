package awesomeProject1

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	//start the tcp server
	listener, err := net.Listen("tcp", ":8080")
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
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			break
		}

		data := strings.TrimSpace(string(buffer[:n]))
		temp := data
		if strings.HasPrefix(data, "chat") {
			temp = "chat"
		}

		switch temp {
		case "login":
			go handleLogin(conn)
		case "chat":
			go handleChat(conn)
		case "exit":
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		default:
			fmt.Println("Unknown command:", data)
		}
	}
}

func handleChat(conn net.Conn /* username string*/) {
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
		fmt.Printf("[%s]: %s\n", username, message)
	}
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

func handleLogin(conn net.Conn) {
	fmt.Println("Client connected:", conn.RemoteAddr())

	// Read username from the client_user
	username, err := readLine(conn)
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	// Read password from the client_user
	password, err := readLine(conn)
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	// Perform login authentication
	if authenticate(username, password) {
		conn.Write([]byte("Login successful. Welcome, " + username + "!\n"))
	} else {
		conn.Write([]byte("Login failed. Invalid username or password.\n"))
	}
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
