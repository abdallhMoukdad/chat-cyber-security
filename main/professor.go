package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")
	conn.Write([]byte("1\n"))

	// Read username from the professor
	fmt.Print("Enter your username (e.g., Professor): ")
	username := readInput()

	//conn.Write([]byte(username + "\n"))

	fmt.Print("Enter your password: ")
	password := readInput()

	//conn.Write([]byte(password + "\n"))
	data := fmt.Sprintf("%s,%s", username, password)
	conn.Write([]byte(data))
	fmt.Println("Data sent:", data)
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}
	fmt.Println("this the message form the backend ", message)
	if message != "Login failed. Invalid username or password.\n" {
		// Start a goroutine to read and display messages from the server
		go func() {
			for {
				//fmt.Println("the pro goroutine started")
				message, err := bufio.NewReader(conn).ReadString('\n')
				//message, err := bufio.NewReader(conn).ReadString('.')

				if err != nil {
					fmt.Println("Error reading message:", err)
					return
				}
				fmt.Print(message)
			}
		}()

		// Read professor's input and send messages to the server
		for {
			fmt.Print("Enter your message: ")

			message := readInput()
			conn.Write([]byte(message + "\n"))
			//message, err := bufio.NewReader(conn).ReadString('\n')
			//if err != nil {
			//	fmt.Println("Error reading message:", err)
			//	return
			//}
			//fmt.Print("this form the backend", message)

			if message == "exit" {
				break
			}
		}
	}

}

func readInput() string {
	var input string
	reader := bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')

	//fmt.Scanln(&input)
	return input
}
