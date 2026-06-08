package main

import (
	"bytes"
	"fmt"
)

type RPCResponse struct {
	Value any
	Error error
}

func (r *RPCResponse) encode() *bytes.Buffer {
	format := fmt.Sprintf("%s %s", r.Value, r.Error)
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return &buf
}
