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
		fmt.Println("Error connecting:", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	readerFromTheServer := bufio.NewReader(conn)

	for {
		//fmt.Print("Enter text: ")
		message, err := readerFromTheServer.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			//break
		}
		fmt.Print("server : " + message + "\n")

		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n")

		//message1, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Print("Server: " + message1)
	}
}
