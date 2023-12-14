package main

import (
	"fmt"
)

func main() {
	var value int
	//var value1 int
	valueChanged := make(chan int)
	valueChanged1 := make(chan int)
	// Goroutine to listen for the change in value
	go func() {
		for {
			newValue := <-valueChanged
			fmt.Printf("Value changed to %d\n", newValue)
			valueChanged1 <- 64

		}
	}()
	go func() {
		for {
			newValue := <-valueChanged1
			fmt.Printf("Value1 changed to %d\n", newValue)
		}
	}()

	// Simulate a delay until the event happens
	//time.Sleep(2 * time.Second)

	// Change the value and notify the listener
	value = 10
	valueChanged <- value

	//time.Sleep(2 * time.Second) // Wait for the listener to print the message

	fmt.Println("Done", value)
}

//func handleConnection(conn net.Conn) {
//	// Wait for a signal to send data
//	signal := make(chan bool)
//	go func() {
//		// Simulate waiting for some event or condition
//		fmt.Println("Waiting for signal to send data...")
//		<-signal
//		fmt.Println("Sending data...")
//
//		// Send the data
//		conn.Write([]byte("Hello, client!"))
//
//		conn.Close()
//	}()
//
//	// Do other processing or handle other requests while waiting
//}
//
//func main() {
//	listener, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		fmt.Println("Error listening:", err.Error())
//		return
//	}
//
//	defer listener.Close()
//
//	fmt.Println("Server started. Listening on :8080")
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Println("Error accepting: ", err.Error())
//			return
//		}
//
//		go handleConnection(conn)
//	}
//}
