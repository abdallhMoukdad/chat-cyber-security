package main

import (
	"awesomeProject1/enc"
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	//arguments := os.Args
	//if len(arguments) == 1 {
	//	fmt.Println("Please provide host:port.")
	//	return
	//}
	//
	//CONNECT := arguments[1]
	//conn, err := net.Dial("tcp", CONNECT)
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
	fmt.Println("2. Full Signup")
	fmt.Println("3. Chat")
	fmt.Println("4. Signup")

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
	case "4\n":
		fmt.Println("You chose to  mini signup.")
		conn.Write([]byte("0" + "\n"))

		handleMiniSignup(conn)
		//handleNotCompleteSignup(conn)

		// Add your chat logic here
	default:
		fmt.Println("Invalid option. Please choose a valid option.")
	}
	//}

}

func handleMiniSignup(conn net.Conn) {
	// Read username from the user
	fmt.Print("Enter your username: ")
	username := readInput1()
	//conn.Write([]byte(username + "\n"))
	conn.Write([]byte("NotCompleteSignup" + "\n"))

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

}

func handleNotCompleteSignup(conn net.Conn) {
	fmt.Print("Enter your national_number: ")
	national_number := readInput1()
	//conn.Write([]byte(username + "\n"))
	//conn.Write([]byte("NotCompleteSignup" + "\n"))

	fmt.Print("Enter your phone_number: ")
	phone_number := readInput1()
	fmt.Print("Enter your home_location: ")
	home_location := readInput1()

	data := fmt.Sprintf("%s,%s,%s", national_number, phone_number, home_location)
	encryptedData, err := enc.GetAESEncrypted(data, "my32digitkey12345678901234567890")
	if err != nil {
		fmt.Println("Error during encryption", err)
	}

	conn.Write([]byte(encryptedData))

	//conn.Write([]byte(data))

	fmt.Println("Data sent:", data)
	fmt.Println("  encrypted Data sent:", encryptedData)

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}
	fmt.Println(message)

}
func readFromConnection(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}
func parse(pubKeyPEM []byte) (*rsa.PublicKey, error) {
	// Decode the public key from PEM format
	block, _ := pem.Decode(pubKeyPEM)
	if block == nil {
		fmt.Println("Error decoding PEM block containing public key")
		return nil, errors.New("Error decoding PEM block containing public key")
	}

	// Parse the public key
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing public key:", err)
		return nil, errors.New("Error parsing public key")
	}

	fmt.Println("Received Public Key:", publicKey)
	return publicKey, nil
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
	clientPrivateKey, clientPublicKey, err := enc.GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating key pair for client:", err)
		return
	}
	clientPubPEM, err := enc.EncodePublicKey(clientPublicKey)
	if err != nil {
		fmt.Println("Error encoding client public key:", err)
		return
	}
	conn.Write([]byte(string(clientPubPEM) + "\n"))
	//serverPublicKey, err := bufio.NewReader(conn).ReadBytes('\n')
	//serverPublicKey, err, _ := bufio.NewReader(conn).ReadLine()
	message, err = bufio.NewReader(conn).ReadString('\n')
	fmt.Println("the message is ", message)
	serverPublicKey, err := readFromConnection(conn)
	//serverPublicKey, err = bufio.NewReader(conn).ReadString("-----END PUBLIC KEY-----")
	//serverPublicKey, err = bufio.NewReader(conn).ReadBytes("-----END PUBLIC KEY-----")

	if err != nil {
		fmt.Println("Error reading the server public key:", err)
		return
	}
	fmt.Print("start\n" + string(serverPublicKey) + "end\n")
	serverPubDecoded, err := enc.DecodePublicKey(serverPublicKey)
	fmt.Println("the professor public key", serverPubDecoded)
	fmt.Printf("Server Public Key: %+v\n", serverPubDecoded)
	//serverPubDecoded, err := parse(serverPublicKey)
	// Start a goroutine to read and display messages from the server
	//go func() {
	//	for {
	//		//fmt.Println("the student goroutine started")
	//
	//		message, err := bufio.NewReader(conn).ReadString('\n')
	//		//fmt.Println(message)
	//
	//		//uDec, _ := base64.URLEncoding.DecodeString(message)
	//		//fmt.Println(string(uDec))
	//		message = strings.TrimSpace(message)
	//
	//		uDec, err := base64.StdEncoding.DecodeString(message)
	//		if err != nil {
	//			fmt.Println("Error decoding base64:", err)
	//			return
	//		}
	//		//fmt.Println(uDec)
	//		//message=uDec
	//		//buffer := make([]byte, 2048)
	//		//n, err := conn.Read(buffer)
	//		//message := string(buffer[:n])
	//
	//		//if err != nil {
	//		//	fmt.Println("Error reading message:", err)
	//		//	return
	//		//}
	//		data, _ := enc.Decrypt([]byte(uDec), clientPrivateKey)
	//		fmt.Println(string(data))
	//
	//		//fmt.Print(message)
	//
	//		//fmt.Print(message)
	//	}
	//}()
	i := 0
	go func() {
		for {
			//fmt.Println("the pro goroutine started")
			message, err := bufio.NewReader(conn).ReadString('\n')
			//fmt.Println(message)

			message = strings.TrimSpace(message)
			//fmt.Println("the message is", message)

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
			data, _ := enc.Decrypt([]byte(uDec), clientPrivateKey)
			fmt.Println(string(data))

		}
	}()
	// Read user input and send messages to the server
	for {
		fmt.Print("Enter your message: ")

		message := readInput1()
		cipherText, err := enc.Encrypt([]byte(message), serverPubDecoded)
		if err != nil {
			fmt.Println("Error encrypt message:", err)
			return
		}
		s64 := base64.StdEncoding.EncodeToString(cipherText)

		//fmt.Println("the cipher text:", s64)
		//conn.Write([]byte(string(cipherText) + "\n"))
		conn.Write([]byte(s64))
		conn.Write([]byte("\n"))

		//conn.Write([]byte(string(cipherText) + "\n"))

		//conn.Write([]byte(message + "\n"))
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
func handleQuetion3U(conn net.Conn) {
	//clientPrivateKey, clientPublicKey, err := enc.GenerateKeyPair()
	_, clientPublicKey, err := enc.GenerateKeyPair()

	if err != nil {
		fmt.Println("Error generating key pair for client:", err)
		return
	}
	clientPubPEM, err := enc.EncodePublicKey(clientPublicKey)
	if err != nil {
		fmt.Println("Error encoding client public key:", err)
		return
	}
	conn.Write([]byte(string(clientPubPEM) + "\n"))
	//serverPublicKey, err := bufio.NewReader(conn).ReadBytes('\n')
	//serverPublicKey, err, _ := bufio.NewReader(conn).ReadLine()
	//message, err := bufio.NewReader(conn).ReadString('\n')
	//fmt.Println("the message is ", message)
	serverPublicKey, err := readFromConnection(conn)
	//serverPublicKey, err = bufio.NewReader(conn).ReadString("-----END PUBLIC KEY-----")
	//serverPublicKey, err = bufio.NewReader(conn).ReadBytes("-----END PUBLIC KEY-----")

	if err != nil {
		fmt.Println("Error reading the server public key:", err)
		return
	}
	fmt.Print("start\n" + string(serverPublicKey) + "end\n")
	serverPubDecoded, err := enc.DecodePublicKey(serverPublicKey)
	fmt.Println("the professor public key", serverPubDecoded)
	fmt.Printf("Server Public Key: %+v\n", serverPubDecoded)
	i := 0
	session := enc.CreateSessionKey()
	//conn.Write([]byte("\n"))

	//conn.Write([]byte("first"))
	//conn.Write([]byte("\n"))
	//conn.Write([]byte(session))
	//conn.Write([]byte("\n"))
	cipherText, err := enc.Encrypt([]byte(session), serverPubDecoded)
	if err != nil {
		fmt.Println("Error encrypt the session key:", err)
		return
	}
	s64 := base64.StdEncoding.EncodeToString([]byte(cipherText))

	conn.Write([]byte(s64))
	conn.Write([]byte("\n"))
	go func() {
		for {
			//fmt.Println("the pro goroutine started")
			message, err := bufio.NewReader(conn).ReadString('\n')
			//fmt.Println(message)

			message = strings.TrimSpace(message)
			//fmt.Println("the message is", message)

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
			//data, _ := enc.Decrypt([]byte(uDec), clientPrivateKey)
			//fmt.Println(string(data))
			data1, _ := enc.GetAESDecrypted(string(uDec), string(session))
			fmt.Println(string(data1))

		}
	}()
	// Read user input and send messages to the server
	for {
		fmt.Print("Enter your message: ")

		message := readInput1()
		//cipherText, err := enc.Encrypt([]byte(message), serverPubDecoded)
		//if err != nil {
		//	fmt.Println("Error encrypt message:", err)
		//	return
		//}
		cipherText, err := enc.GetAESEncrypted(message, string(session))
		if err != nil {
			fmt.Println("Error encrypt message:", err)
			return
		}

		s64 := base64.StdEncoding.EncodeToString([]byte(cipherText))

		//fmt.Println("the cipher text:", s64)
		//conn.Write([]byte(string(cipherText) + "\n"))
		conn.Write([]byte(s64))
		conn.Write([]byte("\n"))

		//conn.Write([]byte(string(cipherText) + "\n"))

		//conn.Write([]byte(message + "\n"))
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

		//handleUserChat(conn)
		fmt.Println("1. chat")
		fmt.Println("2. Q3")
		fmt.Print("Choose an option: ")
		reader := bufio.NewReader(os.Stdin)

		option, _ := reader.ReadString('\n')

		switch option {
		case "1\n":
			conn.Write([]byte("chat" + "\n"))

			fmt.Println("You chose chat.")
			handleUserChat(conn)
			//handleUserLogin1()

			// Add your login logic here
		case "2\n":
			fmt.Println("You chose Q3.")
			conn.Write([]byte("Q3" + "\n"))

			handleQuetion3U(conn)

		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}

		//fmt.Println("hey ya you now we are working")
	} else if message == "2\n" {
		handleNotCompleteSignup(conn)
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
