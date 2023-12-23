package main

import (
	"awesomeProject1/enc"
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
)

func readFromConnection(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}
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
		serverPrivateKey, serverPublicKey, err := enc.GenerateKeyPair()
		if err != nil {
			fmt.Println("Error generating key pair for server:", err)
			return
		}
		serverPubPEM, err := enc.EncodePublicKey(serverPublicKey)
		if err != nil {
			fmt.Println("Error encoding server public key:", err)
			return
		}
		//clientPublicKey, err := bufio.NewReader(conn).ReadBytes('\n')
		clientPublicKey, err := readFromConnection(conn)
		fmt.Print("start\n" + string(clientPublicKey) + "end\n")

		clientPubDecoded, err := enc.DecodePublicKey(clientPublicKey)
		conn.Write([]byte(string(serverPubPEM) + "\n"))
		fmt.Println("the student public key", clientPubDecoded)
		fmt.Printf("Server Public Key: %+v\n", clientPubDecoded)
		// Start a goroutine to read and display messages from the server
		i := 0
		message, err := readFromConnection(conn)

		fmt.Println(string(message))

		go func() {
			for {
				//fmt.Println("the pro goroutine started")
				//message, err := bufio.NewReader(conn).ReadString('\n')
				message, err := bufio.NewReader(conn).ReadString('\n')
				//fmt.Println("the message is", message)

				//fmt.Println(message)

				//message, err := bufio.NewReader(conn).ReadString('.')
				//buffer := make([]byte, 2048)
				//n, err := conn.Read(buffer)
				//message := string(buffer[:n])
				//uDec, _ := base64.URLEncoding.DecodeString(message)
				message = strings.TrimSpace(message)

				uDec, err := base64.StdEncoding.DecodeString(message)
				//fmt.Println("uDec is:", string(uDec))

				if err != nil {
					fmt.Println("Error decoding base64:", err)
					return
				}

				//fmt.Println(string(uDec))

				if err != nil {
					fmt.Println("Error reading message:", err)
					return
				}
				i++
				//fmt.Println("i am here for the time number", i)

				data, _ := enc.Decrypt([]byte(uDec), serverPrivateKey)
				fmt.Println(string(data))
				//fmt.Print(message)

				//fmt.Print(message)
			}
		}()

		// Read professor's input and send messages to the server
		for {
			fmt.Print("Enter your message: ")

			message := readInput()
			//conn.Write([]byte(message + "\n"))
			cipherText, err := enc.Encrypt([]byte(message), clientPubDecoded)
			if err != nil {
				fmt.Println("Error encrypt message:", err)
				return
			}
			//fmt.Println("the cipher text:", cipherText)
			//fmt.Println("the cipher text:", string(cipherText))
			s64 := base64.StdEncoding.EncodeToString(cipherText)

			//fmt.Println("the cipher text:", s64)
			//conn.Write([]byte(string(cipherText) + "\n"))
			conn.Write([]byte(s64))
			//conn.Write([]byte(string('\n')))
			conn.Write([]byte("\n"))

			//conn.Write(cipherText + byte('\n'))
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
