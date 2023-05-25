package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	if err != nil {
		log.Fatal("Error when connect to server")
	}
	fmt.Fprint(conn, "Hello server")

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Error when read response from server")
	}
	fmt.Println(string(buf[:n]))

}
