package main

import (
	"bytes"
	"fmt"
)

type RPCResponse struct {
	sysFd   uintptr
	read    []byte
	written int
	error   error
}

func (r *RPCResponse) encode() *bytes.Buffer {
	format := fmt.Sprintf("%d %v %d %v", r.sysFd, r.read, r.written, r.error)
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return &buf
}
