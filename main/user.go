package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")
	conn.Write([]byte("0"))

	// Read username from the user
	//fmt.Print("Enter your username: ")
	//username := readInput1()
	//conn.Write([]byte(username + "\n"))
	//
	//fmt.Print("Enter your password: ")
	//password := readInput1()
	//conn.Write([]byte(password + "\n"))

	//fmt.Print("Enter your pro name: ")
	//proname := readInput1()
	//conn.Write([]byte(proname + "\n"))

	// Start a goroutine to read and display messages from the server
	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			fmt.Print(message)
		}
	}()

	// Read user input and send messages to the server
	for {
		fmt.Print("Enter your message: ")

		message := readInput1()
		conn.Write([]byte(message + "\n"))

		if message == "exit" {
			break
		}
	}
}

func readInput1() string {
	var input string
	//fmt.Scanln(&input)
	reader := bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')

	return input
}
