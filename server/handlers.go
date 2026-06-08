package main

import (
	"bufio"
	"bytes"
	"fmt"
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
	signature := req.callSignature
	if signature == "" {
		return &RPCResponse{error: fmt.Errorf("Invalid call signature")}
	}

	var res *RPCResponse
	switch signature {
	case "open":
		res = open(req.fileName)
	case "close":
		res = close(req.sysFd, req.fileName)
	case "read":
		res = read(req.fileName)
	case "write":
		res = write(req.sysFd, req.fileName, req.data)
	default:
		res = &RPCResponse{error: fmt.Errorf("Invalid call signature")}
	}

	return res
}

func sendResponse(conn *net.Conn, res *RPCResponse) {
	_, err := (*conn).Write(res.encode().Bytes())
	if err != nil {
		log.Printf("Failed to send response: %v", err)
		return
	}
}
