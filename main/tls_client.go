package main

import (
	"crypto/tls"
	"fmt"
	"net"
)

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
