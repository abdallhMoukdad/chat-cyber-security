package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleClient11(conn net.Conn) {
	defer conn.Close()

	numberChanged := make(chan string)
	//locationChanged := make(chan string)
	//done := make(chan struct{}) // Create a channel to signal when the goroutines should stop
	//var w sync.WaitGroup
	//var m sync.Mutex

	var number = "nil"
	//var homeLoc = "nil"

	fmt.Println("number is ", number)
	fmt.Print("Server: ")
	conn.Write([]byte("enter your phone number\n"))

	reader := bufio.NewReader(conn)

	number, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("number is ", number)

	//go func() {
	//	for {
	//		select {
	//		case newValue := <-numberChanged:
	//			w.Add(1)
	//
	//			fmt.Printf("Value changed to %s\n", newValue)
	//			m.Lock()
	//			conn.Write([]byte("enter your home location\n"))
	//			var err1 error
	//			homeLoc, err1 = reader.ReadString('\n')
	//			m.Unlock()
	//			w.Done()
	//			if err1 != nil {
	//				fmt.Println("Error reading:", err)
	//				close(done) // Signal that the goroutine should stop
	//				return
	//			}
	//			//locationChanged <- homeLoc
	//			locationChanged <- "hiiiiiii"
	//
	//			fmt.Print("Client home location : " + homeLoc)
	//			//close(done) // Signal that the goroutine should stop
	//
	//		case <-done: // Check if the done channel has been closed
	//		case <-numberChanged:
	//			return
	//		}
	//	}
	//}()

	numberChanged <- number
	fmt.Println(number)

	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Print("Client phone number : " + number)
	//go func() {
	//	for {
	//		select {
	//		case newValue := <-locationChanged:
	//			w.Add(1)
	//			fmt.Printf("Value changed to %s\n", newValue)
	//			m.Lock()
	//			conn.Write([]byte("enter your nationalNumber\n"))
	//			nationalNumber, err := reader.ReadString('\n')
	//			m.Unlock()
	//			w.Done()
	//			if err != nil {
	//				fmt.Println("Error reading:", err)
	//				close(done) // Signal that the goroutine should stop
	//				return
	//			}
	//			fmt.Print("Client nationalNumber : " + nationalNumber)
	//		case <-done: // Check if the done channel has been closed
	//		case <-locationChanged:
	//			return
	//		}
	//	}
	//}()

}
func handleClient(conn net.Conn) {
	defer conn.Close()

	//reader := bufio.NewReader(conn)
	//for {
	//message, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("Error reading:", err)
	//	break
	//}
	//fmt.Print("Client: " + message)
	numberChanged := make(chan string, 10)
	locationChanged := make(chan string, 10)
	var number = "nil"
	var homeLoc = "nil"

	fmt.Println("number is ", number)
	fmt.Print("Server: ")
	//response, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//conn.Write([]byte(response))
	conn.Write([]byte("enter your phone number" + "\n"))
	reader := bufio.NewReader(conn)
	number, err := reader.ReadString('\n')
	numberChanged <- number
	fmt.Println(number)
	fmt.Println(homeLoc)

	go func() {
		for {
			fmt.Println("start the number go")
			newValue := <-numberChanged
			fmt.Printf("Value changed to %s\n", newValue)
			conn.Write([]byte("enter your home location" + "\n"))
			homeLoc, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading:", err)
				//break
			}
			//locationChanged <- homeLoc
			locationChanged <- "hiiiiii"

			fmt.Print("Client home location : " + homeLoc)

			fmt.Println("the end the number go")
		}

	}()
	go func() {
		for {
			fmt.Println("start of lock")
			newValue := <-locationChanged
			fmt.Printf("Value changed to %s\n", newValue)
			conn.Write([]byte("enter your nationalNumber" + "\n"))
			nationalNumber, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading:", err)
				//break
			}
			fmt.Print("Client nationalNumber : " + nationalNumber)

			fmt.Println("the end of lock")
		}
	}()

	if err != nil {
		fmt.Println("Error reading:", err)
		//break
	}
	fmt.Print("Client phone number : " + number)
	//conn.Write([]byte("enter your home location" + "\n"))
	//homeLoc, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("Error reading:", err)
	//	//break
	//}
	//fmt.Print("Client home location : " + homeLoc)
	//conn.Write([]byte("enter your nationalNumber" + "\n"))
	//nationalNumber, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("Error reading:", err)
	//	//break
	//}
	//fmt.Print("Client nationalNumber : " + nationalNumber)

	//}
}
func handleClient10(conn net.Conn) {
	defer conn.Close()

	numberChanged := make(chan string)
	locationChanged := make(chan string)

	go func() {
		for {
			conn.Write([]byte("enter your phone number\n"))
			reader := bufio.NewReader(conn)
			number, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading:", err)
				break
			}
			numberChanged <- number
		}
	}()

	go func() {
		for {
			newValue := <-numberChanged
			fmt.Printf("Value changed to %s\n", newValue)
			conn.Write([]byte("enter your home location\n"))
			reader := bufio.NewReader(conn)
			homeLoc, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading:", err)
				break
			}
			fmt.Print("Client home location : " + homeLoc)
		}
	}()

	go func() {
		for {
			newValue := <-locationChanged
			fmt.Printf("Value changed to %s\n", newValue)
			conn.Write([]byte("enter your nationalNumber\n"))
			reader := bufio.NewReader(conn)
			nationalNumber, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Error reading:", err)
				break
			}
			fmt.Print("Client nationalNumber : " + nationalNumber)
		}
	}()
}
func handleClient9(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	conn.Write([]byte("enter your phone number\n"))
	number, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Print("Client phone number : " + number)

	conn.Write([]byte("enter your home location\n"))
	homeLoc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Print("Client home location : " + homeLoc)

	conn.Write([]byte("enter your nationalNumber\n"))
	nationalNumber, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Print("Client nationalNumber : " + nationalNumber)
}

//	func main() {
//		listener, err := net.Listen("tcp", ":8082")
//		if err != nil {
//			fmt.Println("Error listening:", err)
//			return
//		}
//		defer listener.Close()
//
//		fmt.Println("Server started. Listening on :8080")
//
//		for {
//			conn, err := listener.Accept()
//			if err != nil {
//				fmt.Println("Error accepting:", err)
//				break
//			}
//
//			go handleClient11(conn)
//		}
//	}
func main() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			break
		}
		go handleClient11(conn)
	}
}
