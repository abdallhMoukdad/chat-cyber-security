package main

import (
	"awesomeProject1/ent"
	"context"
	//"entgo.io/ent"
	"entgo.io/ent/dialect"
	_ "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq" // add this
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func RegisterUser(client *ent.Client, username, password, email string) error {

	_, err := client.Student.
		Create().SetName(username).SetEmail(email).
		Save(context.Background())
	if err != nil {
		log.Printf("failed creating user: %v", err)
		return err
	}
	return nil
}
func main() {
	client, err := ent.Open(dialect.Postgres, "user=postgres password=postgres dbname=ISS sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	err = RegisterUser(client, "john_doe", "password123", "john@example.com")
	if err != nil {
		log.Printf("failed to register user: %v", err)
		return
	}

	log.Println("user registered successfully")

	////start the tcp server
	//listener, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	fmt.Println("Error starting server:", err)
	//	return
	//}
	//defer listener.Close()
	//
	//fmt.Println("Server started. Waiting for connections...")
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		fmt.Println("Error accepting connection:", err)
	//		continue
	//	}
	//
	//	go handleConnection(conn)
	//}
}
