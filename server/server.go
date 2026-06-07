package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type TCPServer struct {
	port    int
	conns   []net.Conn
	maxConn int
}

func NewTCPServer(port, maxConn int) *TCPServer {
	conns := []net.Conn{}
	return &TCPServer{port, conns, maxConn}
}

func (srv *TCPServer) Listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", srv.port))
	if err != nil {
		log.Fatal("Error listening", err)
	}

	defer listener.Close()

	log.Printf("server listening on port %d", srv.port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error handling conn", err)
			continue
		}

		go srv.handleConn(conn)
	}
}

func (srv *TCPServer) handleConn(conn net.Conn) {
	defer conn.Close()
	defer func() {
		srv.conns = srv.conns[:len(srv.conns) - 1]
	}()

	if len(srv.conns) >= srv.maxConn {
		log.Println("Server has max number of connections")
		return
	}

	srv.conns = append(srv.conns, conn)

	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')
	message = strings.TrimSpace(message)
	log.Println(message)

	if err != nil {
		log.Printf("Failed to read conn message %e", err)
		return
	}

	var response = "\n"
	if message == "foo" {
		response = foo()
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Failed to send response: %e", err)
	}
}

func foo() string {
	return "open read read close \n"
}
