package main

import (
	"fmt"
	"log"
	"net"
)

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
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("Error when read request from client")
		}
		fmt.Println(string(buf[:n]))
		fmt.Fprintf(conn, "Hello client")
		conn.Close()
	}
}
