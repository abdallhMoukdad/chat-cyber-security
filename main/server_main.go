package main

import (
	"awesomeProject1/enc"
	"awesomeProject1/ent"
	"awesomeProject1/ent/professor"
	"awesomeProject1/ent/student"
	"bufio"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/pem"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // add this
	"io"
	"log"
	"net"
	"os"
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
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			// break
		}

		data := strings.TrimSpace(string(buffer[:n]))
		switch data {
		case "signup":
			fmt.Println("signup")
			handleNotCompleteSignupPro(conn, client)

		case "login":
			fmt.Println("login")
			handleProfessorLogin(client, conn)

			//handleLogin(conn, client)
		case "exit":
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		default:
			fmt.Println("Unknown command:", data)
		}

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
func handleNotCompleteSignupPro(conn net.Conn, client *ent.Client) {
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
	username = strings.Replace(username, "\n", "", -1)

	password := splitValues[1]
	password = strings.Replace(password, "\n", "", -1)

	exsits, err := RegisterProfessor(client, username, password)
	if err != nil && exsits == false {
		log.Printf("failed to register user: %v", err)
		conn.Write([]byte("register failed. Invalid username or password.\n"))
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
		role, _ := bufio.NewReader(conn).ReadString('\n')
		// fmt.Println("hey ya you now we are working")
		result := strings.Replace(role, "\n", "", -1)
		role = result
		if role == "chat" {
			fmt.Println("chatttttttttttttttttttttttttttttttttt")
			handleProfessorChat(conn, username)

		} else if role == "Q3" {

			fmt.Println("QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ3")
			handleQuestion3(conn, client, username, password)
		} else if role == "Q4" {

			fmt.Println("QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ4")
			handleQuestion4(conn, client, username)
		} else if role == "Q5" {

			fmt.Println("QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ5")
			handleQuestion5(conn, client, username)
		}

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
func handleQuestion3(conn net.Conn, client *ent.Client, userName, password string) {
	serverPrivateKey, serverPublicKey, err := enc.GenerateKeyPair()
	//_, serverPublicKey, err := enc.GenerateKeyPair()

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
	//message, err := readFromConnection(conn)
	//session, err := bufio.NewReader(conn).ReadString('\n')
	//message, err := bufio.NewReader(conn).ReadString('\n')
	//fmt.Println(string(message))

	session, err := bufio.NewReader(conn).ReadBytes('\n')
	session = []byte(strings.TrimSpace(string(session)))

	uDec, err := base64.StdEncoding.DecodeString(string(session))
	data, _ := enc.Decrypt([]byte(uDec), serverPrivateKey)

	session = data
	fmt.Printf("Session Key: %x\n", session)

	//fmt.Println("this the session key", string(session))
	//
	_, err = StorePublicKey(client, userName, password, string(clientPublicKey))
	if err != nil {
		log.Println("something went wrong")

		//return
	}
	log.Println("the public key stored")

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
			data1, _ := enc.GetAESDecrypted(string(uDec), string(session))
			fmt.Println(string(data1))

			//data, _ := enc.Decrypt([]byte(uDec), serverPrivateKey)
			//fmt.Println(string(data))
			cipherText, err := enc.Encrypt([]byte("Done"), clientPubDecoded)
			s64 := base64.StdEncoding.EncodeToString(cipherText)
			conn.Write([]byte(s64))
			conn.Write([]byte("\n"))

			//fmt.Print(message)

			//fmt.Print(message)
		}
	}()

	// Read professor's input and send messages to the server
	for {
		fmt.Print("Enter your message: ")

		message := readInput()
		//conn.Write([]byte(message + "\n"))
		//cipherText, err := enc.Encrypt([]byte(message), clientPubDecoded)
		//if err != nil {
		//	fmt.Println("Error encrypt message:", err)
		//	return
		//}
		cipherText, err := enc.GetAESEncrypted(message, string(session))
		if err != nil {
			fmt.Println("Error encrypt message:", err)
			return
		}

		//fmt.Println("the cipher text:", cipherText)
		//fmt.Println("the cipher text:", string(cipherText))
		s64 := base64.StdEncoding.EncodeToString([]byte(cipherText))

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
func readInput() string {
	var input string
	reader := bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')

	//fmt.Scanln(&input)
	return input
}
func readFromConnection(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
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
		role, _ := bufio.NewReader(conn).ReadString('\n')
		// fmt.Println("hey ya you now we are working")
		result := strings.Replace(role, "\n", "", -1)
		role = result
		if role == "chat" {
			fmt.Println("chatttttttttttttttttttttttttttttttttt")
			handleChat(conn, username)

		} else if role == "Q3" {

			fmt.Println("QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ3")
			handleQuestion3(conn, client, username, password)
		} else if role == "Q4" {

			fmt.Println("QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ4")
			handleQuestion4(conn, client, username)
		} else if role == "Q5" {

			fmt.Println("QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ5")
			handleQuestion5(conn, client, username)
		}

		//handleChat(conn, username)
	}

	// Perform login authentication
	//if authenticate(username, password) {
	//	conn.Write([]byte("Login successful. Welcome, " + username + "!\n"))
	//} else {
	//	conn.Write([]byte("Login failed. Invalid username or password.\n"))
	//}
}
func receiveCertificate(conn net.Conn) ([]byte, error) {
	lengthBytes := make([]byte, 4)
	_, err := conn.Read(lengthBytes)
	if err != nil {
		return nil, err
	}

	certLength := int(binary.BigEndian.Uint32(lengthBytes))
	certBytes := make([]byte, certLength)
	_, err = conn.Read(certBytes)
	if err != nil {
		return nil, err
	}

	return certBytes, nil
}
func receiveCertificate1(conn net.Conn) ([]byte, error) {
	lengthBytes := make([]byte, 4)
	_, err := io.ReadFull(conn, lengthBytes)
	if err != nil {
		return nil, fmt.Errorf("error reading certificate length: %w", err)
	}

	certLength := int(binary.BigEndian.Uint32(lengthBytes))
	certBytes := make([]byte, certLength)

	n, err := io.ReadFull(conn, certBytes)
	if err != nil {
		return nil, fmt.Errorf("error reading certificate data: %w", err)
	}

	// Check if all expected bytes have been read
	if n != certLength {
		return nil, fmt.Errorf("expected %d bytes, but only read %d bytes", certLength, n)
	}

	return certBytes, nil
}

func handleQuestion5(conn net.Conn, client *ent.Client, username string) {
	// Receive client certificate
	clientCertificate, err := receiveCertificate(conn)
	if err != nil {
		fmt.Println("Error receiving client certificate:", err)
		return
	}
	clientCertFile, err := os.Create("client_certificate.pem")

	//clientCertFile, err := os.Create(username + "_client_certificate.pem")
	if err != nil {
		fmt.Println("Error creating client certificate file:", err)
		return
	}
	_, err = clientCertFile.Write(clientCertificate)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	//// Create a PEM block
	//pemBlock := &pem.Block{
	//	Type:  "CERTIFICATE",
	//	Bytes: clientCertificate,
	//}
	//
	//// Encode the PEM block to the file
	//err = pem.Encode(clientCertFile, pemBlock)
	//if err != nil {
	//	fmt.Println("Error encoding client certificate:", err)
	//	return
	//}
	stat, err := clientCertFile.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat.Size())
	fmt.Println(len(clientCertificate))
	clientCertFile.Close()
	fmt.Println("Client certificate saved to file.")
	//// Save client certificate to file (optional)
	//clientCertFile, err := os.Create(username + "_client_certificate.pem")
	//if err != nil {
	//	fmt.Println("Error creating client certificate file:", err)
	//	return
	//}
	//pem.Encode(clientCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: clientCertificate})
	//clientCertFile.Close()
	//fmt.Println("done")
	// Decode the PEM-encoded certificate
	block, _ := pem.Decode(clientCertificate)
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Println("Failed to decode PEM block containing certificate")
		return
	}

	// Parse the certificate
	receivedCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing received certificate:", err)
		return
	}

	// Load the root certificate (replace "rootCA.pem" with the path to your root CA certificate)
	rootCertPEM, err := os.ReadFile("ca_certificate.pem")
	if err != nil {
		fmt.Println("Error reading root certificate:", err)
		return
	}

	// Decode the PEM-encoded root certificate
	block, _ = pem.Decode(rootCertPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Println("Failed to decode PEM block containing root certificate")
		return
	}

	// Parse the root certificate
	rootCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing root certificate:", err)
		return
	}

	// Create a certificate pool containing the root certificate
	rootCertPool := x509.NewCertPool()
	rootCertPool.AddCert(rootCert)

	// Verify the received certificate against the root certificate
	opts := x509.VerifyOptions{
		Roots:         rootCertPool,
		Intermediates: x509.NewCertPool(),
	}

	if _, err := receivedCert.Verify(opts); err != nil {
		fmt.Println("Certificate verification failed:", err)
		return
	}

	fmt.Println("Certificate verification successful.")
}

func handleQuestion4(conn net.Conn, client *ent.Client, username string) {

	serverPrivateKey, serverPublicKey, err := enc.GenerateKeyPair()
	//_, serverPublicKey, err := enc.GenerateKeyPair()

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

	session, err := bufio.NewReader(conn).ReadBytes('\n')
	session = []byte(strings.TrimSpace(string(session)))

	uDec, err := base64.StdEncoding.DecodeString(string(session))
	data, _ := enc.Decrypt([]byte(uDec), serverPrivateKey)

	session = data
	fmt.Printf("Session Key: %x\n", session)

	var PublicKey rsa.PublicKey
	err = gob.NewDecoder(conn).Decode(&PublicKey)
	if err != nil {
		fmt.Println("Error receiving PublicKey :", err)
		return
	}

	// fmt.Println(PublicKey)
	fmt.Printf("Server Public Key: %+v\n", PublicKey)
	var signedPDF []byte

	err = gob.NewDecoder(conn).Decode(&signedPDF)
	// err = gob.NewDecoder(conn).Decode(&signedPDF)
	if err != nil {
		fmt.Println("Error receiving signed PDF:", err)
		return
	}
	//signedPDF, _ = enc.Decrypt([]byte(signedPDF), serverPrivateKey)
	signedPDF, _ = enc.GetAESDecrypted1(signedPDF, string(session))

	sig := signedPDF[len(signedPDF)-256:]
	//hashedPDF := sha256.Sum256(fileContents)

	hash := sha256.Sum256(signedPDF[:len(signedPDF)-256])
	// Verify the signature
	//hashedPDF := sha256.Sum256(signedPDF)
	//err = rsa.VerifyPKCS1v15(&PublicKey, crypto.SHA256, hashedPDF[:], signedPDF[len(signedPDF)-256:])
	err = rsa.VerifyPKCS1v15(&PublicKey, crypto.SHA256, hash[:], sig)

	// fmt.Println("the sign len", (len(signedPDF) - 256))
	fmt.Println("the sign len", len(signedPDF[len(signedPDF)-256:]))
	// fmt.Printf("%v\n", signedPDF[len(signedPDF)-256:])
	// err = rsa.VerifyPKCS1v15(&PublicKey, crypto.SHA256, hashedPDF[:], signedPDF[:])
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return
	}

	// Save the verified PDF file
	fileName := "received.pdf"
	// err = os.WriteFile("received.pdf", signedPDF[:], 0644)
	err = os.WriteFile(fileName, signedPDF[:len(signedPDF)-256], 0644)
	if err != nil {
		fmt.Println("Error saving received PDF:", err)
		return
	}

	fmt.Println("File received and verified successfully.")
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set the log output to the file
	log.SetOutput(file)

	// Write logs to the file

	log.Println("This is a log message")
	log.Printf("This is a formatted log message: %s", "Hello, World!")
	ui := uuid.New().String()
	log.Printf("the user: %s did send the file : %s and the action with uuid: %s", username, fileName, ui)
	conn.Write([]byte(ui))
	conn.Write([]byte("\n"))

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
func StorePublicKey(client *ent.Client, userName, password, publicKey string) (bool, error) {
	//user, err := client.Student.Query().Where(student.Name(userName),
	//	student.Password(password)).Only(context.Background())
	user, err := client.Student.Query().Where(student.Name(userName),
		student.Password(password)).First(context.Background())

	if err != nil {
		log.Printf("failed creating user: %v", err)
		return false, err
	} else {
		user.Update().SetEnycrptionKey(publicKey).Save(context.Background())
		return true, nil
	}

	return false, nil
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

func RegisterProfessor(client *ent.Client, username, password string) (bool, error) {
	_, err := client.Professor.
		Create().SetName(username).
		SetPassword(password).
		Save(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return false, err
	}
	return true, nil
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
