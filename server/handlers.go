package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"

	"github.com/drizlye0/15440-project1/lib"
)

func handleMessage(conn *net.Conn) *lib.RPCRequest {
	reader := bufio.NewReader(*conn)

	data, err := reader.ReadBytes(byte('\n'))
	if err != nil {
		log.Printf("Failed to read conn message %e", err)
		return nil
	}

	buf := bytes.NewBuffer(data)
	req := lib.DecodeRPCRequest(buf)

	if req == nil {
		return nil
	}

	return req
}

func handleRequest(req *lib.RPCRequest) *lib.RPCResponse {
	signature := req.CallSignature
	if signature == "" {
		return &lib.RPCResponse{Error: fmt.Errorf("Invalid call signature")}
	}

	var res *lib.RPCResponse
	switch signature {
	case "open":
		res = open(req.FileName)
	case "close":
		res = close(req.SysFd, req.FileName)
	case "read":
		res = read(req.FileName)
	case "write":
		res = write(req.SysFd, req.FileName, req.Data)
	default:
		res = &lib.RPCResponse{Error: fmt.Errorf("Invalid call signature")}
	}

	return res
}

func sendResponse(conn *net.Conn, res *lib.RPCResponse) {
	_, err := (*conn).Write(res.Encode().Bytes())
	if err != nil {
		log.Printf("Failed to send response: %v", err)
		return
	}
}
