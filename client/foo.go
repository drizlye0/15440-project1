package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func foo() {
	addr := "localhost:8080"
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)

	if err != nil {
		log.Printf("Failed to connect to server: %e\n", err)
		return
	}

	defer conn.Close()

	conn.Write([]byte("foo\n"))

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	if err != nil {
		log.Printf("Failed to read server response: %e\n", err)
		return
	}

	response = strings.ReplaceAll(response, " ", "\n")

	fmt.Println(response)
}
