package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8082")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server.")
	//conn.Write([]byte("0"))

	reader := bufio.NewReader(os.Stdin)
	//for {
	fmt.Println("Welcome to the Student Portal!")
	fmt.Println("1. Login")
	fmt.Println("2. Signup")
	fmt.Println("3. Chat")
	fmt.Print("Choose an option: ")

	option, _ := reader.ReadString('\n')

	switch option {
	case "1\n":
		conn.Write([]byte("0" + "\n"))

		fmt.Println("You chose to login.")
		handleUserLogin(conn)
		//handleUserLogin1()

		// Add your login logic here
	case "2\n":
		fmt.Println("You chose to signup.")
		conn.Write([]byte("0" + "\n"))

		handleUserSignup(conn)

		// Add your signup logic here
	case "3\n":
		fmt.Println("You chose to chat.")
		handleUserChat(conn)

		// Add your chat logic here
	default:
		fmt.Println("Invalid option. Please choose a valid option.")
	}
	//}

}

func handleUserChat(conn net.Conn) {

	fmt.Print("Enter  professor name: ")
	proname := readInput1()
	conn.Write([]byte(proname + "\n"))
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}
	fmt.Print(message)
	if message == "0\n" {
		log.Println("the professor not exsits yet or not contected")
		return
	}

	// Start a goroutine to read and display messages from the server
	go func() {
		for {
			//fmt.Println("the student goroutine started")
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			fmt.Print(message)

			//fmt.Print(message)
		}
	}()

	// Read user input and send messages to the server
	for {
		fmt.Print("Enter your message: ")

		message := readInput1()
		conn.Write([]byte(message + "\n"))
		//message, err := bufio.NewReader(conn).ReadString('\n')
		//if err != nil {
		//	fmt.Println("Error reading message:", err)
		//	return
		//}
		//fmt.Print(message)

		if message == "exit" {
			break
		}
	}

}

func handleUserSignup(conn net.Conn) {
	//conn.Write([]byte("1" + "\n"))

	// Read username from the user
	fmt.Print("Enter your username: ")
	username := readInput1()
	//conn.Write([]byte(username + "\n"))
	conn.Write([]byte("signup" + "\n"))

	fmt.Print("Enter your password: ")
	password := readInput1()
	//conn.Write([]byte(password + "\n"))
	fmt.Print("Enter your national number: ")
	nationalNumber := readInput1()
	//conn.Write([]byte(nationalNumber + "\n"))
	fmt.Print("Enter your home location: ")
	homeLocation := readInput1()
	//conn.Write([]byte(homeLocation + "\n"))
	fmt.Print("Enter your phone number: ")
	phoneNumber := readInput1()
	//conn.Write([]byte(phoneNumber+ "\n"))
	data := fmt.Sprintf("%s,%s,%s,%s,%s", username, password, phoneNumber, nationalNumber, homeLocation)
	conn.Write([]byte(data))

	fmt.Println("Data sent:", data)
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}
	fmt.Println(message)

}

func handleUserLogin(conn net.Conn) {
	//conn.Write([]byte("1" + "\n"))

	// Read username from the user
	fmt.Print("Enter your username: ")
	username := readInput1()
	//conn.Write([]byte(username + "\n"))
	conn.Write([]byte("login" + "\n"))

	fmt.Print("Enter your password: ")
	password := readInput1()
	//conn.Write([]byte(password + "\n"))
	data := fmt.Sprintf("%s,%s", username, password)
	conn.Write([]byte(data))
	fmt.Println("Data sent:", data)
	time.Sleep(2 * time.Second)

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}
	fmt.Println("this the message form the backend ", message)
	if message == "1\n" {
		handleUserChat(conn)

		//fmt.Println("hey ya you now we are working")
	}
}
func readInput1() string {
	var input string
	//fmt.Scanln(&input)
	reader := bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')

	return input
}
func handleUserLogin1() {
	//conn.Write([]byte("1" + "\n"))
	//conn.Write([]byte("login" + "\n"))

	// Read username from the user
	fmt.Print("Enter your username: ")
	username := readInput1()
	//conn.Write([]byte(username + "\n"))

	fmt.Print("Enter your password: ")
	password := readInput1()
	//conn.Write([]byte(password + "\n"))
	data := fmt.Sprintf("%s,%s", username, password)
	//conn.Write([]byte(data))
	//
	fmt.Println("Data sent:", data)

}
