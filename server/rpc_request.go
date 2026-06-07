package main

import (
	"bytes"
	"fmt"
)

type RPCRequest struct {
	id int
	callSignature string
	paramaters int
}

func (m *RPCRequest) encode() bytes.Buffer {
	format := fmt.Sprintf("%d %s %d", m.id, m.callSignature, m.paramaters)
	b := bytes.Buffer{}
	b.Write([]byte(format))
	return b
}

func decodeRPCRequest(b bytes.Buffer) *RPCRequest {
	parts := bytes.Split(b.Bytes(), []byte(" "))
	for _, part := range parts {
		fmt.Printf("%s\n", part)
	}

	return nil
}
