package client_user

import (
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Println("Welcome to the University Chat App!")
		fmt.Println("1. Login")
		fmt.Println("2. Chat with Professor")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			login()
		case 2:
			chatWithProfessor()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func login() {
	fmt.Println("Logging in...")
}

func chatWithProfessor() {
	fmt.Println("Chatting with professor...")
}
