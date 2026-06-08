package main

import (
	"fmt"
	"log"
	"net"
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
		srv.conns = srv.conns[:len(srv.conns)-1]
	}()

	if len(srv.conns) >= srv.maxConn {
		log.Println("Server has max number of connections")
		return
	}

	srv.conns = append(srv.conns, conn)

	var res *RPCResponse
	req := handleMessage(&conn)
	if req == nil {
		res = &RPCResponse{error: fmt.Errorf("Malformed message")}
		sendResponse(&conn, res)
		return
	}

	res = handleRequest(req)
	if res == nil {
		res = &RPCResponse{error: fmt.Errorf("Failed to handle request")}
		sendResponse(&conn, res)
	}

	sendResponse(&conn, res)
}

func foo() string {
	return "open read read close \n"
}
