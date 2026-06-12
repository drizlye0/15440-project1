package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/drizlye0/15440-project1/lib"
)

type TCPServer struct {
	port    int
	conns   []net.Conn
	maxConn int
	mu      sync.Mutex
}

func NewTCPServer(port, maxConn int) *TCPServer {
	return &TCPServer{
		port:    port,
		conns:   []net.Conn{},
		maxConn: maxConn,
	}
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

	srv.mu.Lock()
	if len(srv.conns) >= srv.maxConn {
		srv.mu.Unlock()
		log.Println("Server has max number of connections")
		return
	}
	srv.conns = append(srv.conns, conn)
	srv.mu.Unlock()

	defer func() {
		srv.mu.Lock()
		srv.conns = srv.conns[:len(srv.conns)-1]
		srv.mu.Unlock()
	}()

	var res *lib.RPCResponse
	req := handleMessage(&conn)
	if req == nil {
		res = &lib.RPCResponse{Error: fmt.Errorf("Malformed message")}
		sendResponse(&conn, res)
		return
	}

	res = handleRequest(req)
	if res == nil {
		res = &lib.RPCResponse{Error: fmt.Errorf("Failed to handle request")}
		sendResponse(&conn, res)
	}

	sendResponse(&conn, res)
}
