package client_user

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type 'login' or 'chat' to start.")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		// Send the command to the server
		conn.Write([]byte(command))
		if command != "chat" {
			// Start chat
			go receiveMessages(conn)
			sendMessages(conn)
		} else if command == "login" {
			handleLogin(conn)
		} else if command == "exit" {
			// Handle Ctrl+C to gracefully exit
			signalCh := make(chan os.Signal, 1)
			signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				<-signalCh
				fmt.Println("Closing connection...")
				conn.Close()
				os.Exit(0)
			}()
			break
		} else {
			fmt.Println("Unknown command")
		}
	}
}
func readInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func sendMessages(conn net.Conn) {
	for {
		fmt.Print("Enter message (type 'exit' to quit): ")
		message := readInput()

		// Send the message to the server
		conn.Write([]byte(message))

		// Check for exit command
		if strings.ToLower(message) == "exit" {
			break
		}
	}
}

func receiveMessages(conn net.Conn) {
	for {
		// Read and display messages from the server
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		message := strings.TrimSpace(string(buffer[:n]))
		fmt.Println("Server:", message)
	}
}

// -----------------------------------------
func handleLogin(conn net.Conn) {
	// Read username and password from the user
	fmt.Print("Enter username: ")
	username := readInput()

	fmt.Print("Enter password: ")
	password := readInput()

	// Send username and password to the server
	conn.Write([]byte(username + "\n"))
	conn.Write([]byte(password + "\n"))

	// Read and display server response
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Print(string(response[:n]))
}
