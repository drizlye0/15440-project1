package main

import (
	"bytes"
	"fmt"
	"log"
)

type RPCRequest struct {
	callSignature string
	fileName      string
}

func (r *RPCRequest) encode() bytes.Buffer {
	format := fmt.Sprintf("%s %s", r.callSignature, r.fileName)
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return buf
}

func decodeRPCRequest(b *bytes.Buffer) *RPCRequest {
	cleanedBytes := bytes.TrimSpace(b.Bytes())
	parts := bytes.Split(cleanedBytes, []byte(" "))

	if len(parts) < 2 {
		log.Println("Failed to parse request: invalid format or missing arguments")
		return nil
	}

	return &RPCRequest{
		callSignature: string(parts[0]),
		fileName:      string(parts[1]),
	}
}
