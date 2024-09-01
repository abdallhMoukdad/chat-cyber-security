package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Perform TLS handshake
	config := &tls.Config{
		// Configure TLS options here (e.g., certificates, cipher suites)
		// ...
	}
	tlsConn := tls.Server(conn, config)
	if err := tlsConn.Handshake(); err != nil {
		fmt.Println("TLS handshake error:", err)
		return
	}

	// Read data from client
	buf := make([]byte, 1024)
	n, err := tlsConn.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Read error:", err)
		return
	}

	// Process received data
	fmt.Println("Received data:", string(buf[:n]))

	// Send response back to client
	response := []byte("Hello from server!")
	_, err = tlsConn.Write(response)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
}
func main() {
	//go func() {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
	//}()

	connect()
}
func connect() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}

	// Perform TLS handshake
	config := &tls.Config{
		InsecureSkipVerify: true,
		// Configure other TLS options here (e.g., certificates, cipher suites)
		// ...
	}
	tlsConn := tls.Client(conn, config)

	if err = tlsConn.Handshake(); err != nil {
		fmt.Println("TLS handshake error:", err)
		return
	}

	// Send data to server
	message := []byte("Hello from client!")
	_, err = tlsConn.Write(message)

	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	// Read response from server
	buf := make([]byte, 1024)
	n, _ := tlsConn.Read(buf)

	fmt.Println("Received response:", string(buf[:n]))
}
