package main

import (
	"awesomeProject1/enc"
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func handleMiniSignupPro(conn net.Conn) {
	// Read username from the user
	fmt.Print("Enter your username: ")
	username := readInput()
	//conn.Write([]byte(username + "\n"))
	//conn.Write([]byte("signup" + "\n"))

	fmt.Print("Enter your password: ")
	password := readInput()
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

func readFromConnection(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}
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
	conn.Write([]byte("1\n"))
	fmt.Println("1. signup")
	fmt.Println("2. login")

	fmt.Print("Choose an option: ")
	reader := bufio.NewReader(os.Stdin)

	option, _ := reader.ReadString('\n')

	switch option {
	case "1\n":
		conn.Write([]byte("signup" + "\n"))

		fmt.Println("You chose signup.")
		handleMiniSignupPro(conn)
		//handleUserLogin1()

		// Add your login logic here
	case "2\n":
		fmt.Println("You chose login.")
		//conn.Write([]byte("Q3" + "\n"))
		conn.Write([]byte("login\n"))

		logninProf(conn, err)
		//handleQuetion3U(conn)

	default:
		fmt.Println("Invalid option. Please choose a valid option.")
	}

}

//logninProf(conn, err)

func logninProf(conn net.Conn, err error) {
	// Read username from the professor
	fmt.Print("Enter your username (e.g., Professor): ")
	username := readInput()
	//conn.Write([]byte("login\n"))

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
		fmt.Println("1. chat")
		fmt.Println("2. Q3")
		fmt.Println("3. Q4")
		fmt.Println("4. Q5")

		fmt.Print("Choose an option: ")
		reader := bufio.NewReader(os.Stdin)

		option, _ := reader.ReadString('\n')

		switch option {
		case "1\n":
			conn.Write([]byte("chat" + "\n"))

			fmt.Println("You chose chat.")
			handleProChat(conn)
			//handleUserLogin1()

			// Add your login logic here
		case "2\n":
			fmt.Println("You chose Q3.")
			conn.Write([]byte("Q3" + "\n"))
			handleQuetion3P(conn)
			//handleQuetion3U(conn)
		case "3\n":
			fmt.Println("You chose Q4.")
			conn.Write([]byte("Q4" + "\n"))

			handleQuetion4P(conn)
		case "4\n":
			fmt.Println("You chose Q5.")
			conn.Write([]byte("Q5" + "\n"))

			handleQuetion5P(conn)

		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}

	}
}

func handleQuetion5P(conn net.Conn) {
	file, err := os.ReadFile("client_certificate.pem")
	if err != nil {
		return
	}
	err = sendCertificate(conn, file)
	if err != nil {
		fmt.Println("Error sending client certificate:", err)
		return
	}
	fmt.Printf("done")
	//pem.Decode(file)
}
func sendCertificate(conn net.Conn, certificate []byte) error {
	err := sendLength(conn, len(certificate))
	if err != nil {
		return err
	}

	_, err = conn.Write(certificate)
	return err
}

func sendLength(conn net.Conn, length int) error {
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))
	_, err := conn.Write(lengthBytes)
	return err
}

func handleQuetion4P(conn net.Conn) {
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

	if err != nil {
		fmt.Println("Error reading the server public key:", err)
		return
	}
	fmt.Print("start\n" + string(serverPublicKey) + "end\n")
	serverPubDecoded, err := enc.DecodePublicKey(serverPublicKey)
	fmt.Println("the professor public key", serverPubDecoded)
	fmt.Printf("Server Public Key: %+v\n", serverPubDecoded)
	session := enc.CreateSessionKey()

	cipherText, err := enc.Encrypt([]byte(session), serverPubDecoded)
	if err != nil {
		fmt.Println("Error encrypt the session key:", err)
		return
	}
	s64 := base64.StdEncoding.EncodeToString([]byte(cipherText))

	conn.Write([]byte(s64))
	conn.Write([]byte("\n"))

	// Generate RSA key pair for signing and verification
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA key pair:", err)
		return
	}

	// fmt.Println(&privKey.PublicKey)
	// fmt.Println(string(privKey.PublicKey))
	fmt.Printf("Server Public Key: %+v\n", privKey.PublicKey)

	// Send the public key to the client
	err = gob.NewEncoder(conn).Encode(&privKey.PublicKey)
	if err != nil {
		fmt.Println("Error sending public key:", err)
		return
	}
	// Read the PDF file to be sent
	fileContents, err := os.ReadFile("example.pdf")
	if err != nil {
		fmt.Println("Error reading PDF file:", err)
		return
	}

	// Sign the PDF file
	hashedPDF := sha256.Sum256(fileContents)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashedPDF[:])
	if err != nil {
		fmt.Println("Error signing PDF file:", err)
		return
	}
	fmt.Println("the sign len", len(signature))
	//
	// Send the signed PDF file to the server
	signedPDF := append(fileContents, signature...)
	signedPDF, err = enc.GetAESEncrypted1(signedPDF, string(session))

	//signedPDF, err = enc.Encrypt([]byte(signedPDF), serverPubDecoded)
	if err != nil {
		fmt.Println("Error encrypt the signedPDF:", err)
		return
	}

	err = gob.NewEncoder(conn).Encode(&signedPDF)

	// err = gob.NewEncoder(conn).Encode(fileContents)
	if err != nil {
		fmt.Println("Error sending signed PDF:", err)
		return
	}

	fmt.Println("File sent successfully.")
	message, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message)
}
func PrintAllPdfNamesInMyWorkingDir() {

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// List all files with the ".pdf" extension in the current directory
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pdf") {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}

}
func SendTheFile(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of the file with the ext")

	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)
	file, err := os.Open(option)
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		bytesRead, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading file:", err.Error())
			return
		}

		conn.Write(buffer[:bytesRead])
	}

	fmt.Println("File sent successfully!")
}

//func signMessage(message []byte, privateKey *rsa.PrivateKey) (string, error) {
//	hashed := sha256.Sum256(message)
//
//	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
//	if err != nil {
//		return "", fmt.Errorf("failed to sign message: %w", err)
//	}
//
//	return base64.StdEncoding.EncodeToString(signature), nil
//}

//	func verifySignature(message []byte, signature string, publicKey *rsa.PublicKey) (bool, error) {
//		hashed := sha256.Sum256(message)
//
//		signatureBytes, err := base64.StdEncoding.DecodeString(signature)
//		if err != nil {
//			return false, err
//		}
//		err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, &hashed, &signatureBytes)
//
//		if err != nil {
//			return false, err
//		}
//
//		return true, nil
//	}
func handleProChat(conn net.Conn) {
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
func readInput() string {
	var input string
	reader := bufio.NewReader(os.Stdin)

	input, _ = reader.ReadString('\n')

	//fmt.Scanln(&input)
	return input
}
func handleQuetion3P(conn net.Conn) {
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

		message := readInput()
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
