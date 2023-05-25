package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Error when create server")
	}

	// Print the message
	fmt.Println(string(buf[:n]))

	// Send a response to the client
	fmt.Fprintf(conn, "Hello, client!")

	// Close the connection
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error when create server")
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error when accept connect")
		}
		// create go routine to handle multi client request
		go handleConnection(conn)
	}
}
