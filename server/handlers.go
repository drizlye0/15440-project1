package main

import (
	"bufio"
	"bytes"
	"log"
	"net"
)

func handleMessage(conn *net.Conn) *RPCRequest {
	reader := bufio.NewReader(*conn)

	data, err := reader.ReadBytes(byte('\n'))
	if err != nil {
		log.Printf("Failed to read conn message %e", err)
		return nil
	}

	buf := bytes.NewBuffer(data)
	req := decodeRPCRequest(buf)

	if req == nil {
		return nil
	}

	return req
}

func handleRequest(req *RPCRequest) *RPCResponse {
	return nil
}

func sendResponse(conn *net.Conn, res *RPCResponse) {
	_, err := (*conn).Write(res.encode().Bytes())
	if err != nil {
		log.Printf("Failed to send response: %v", err)
		return
	}
}
