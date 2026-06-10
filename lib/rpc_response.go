package lib

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
)

const (
	RES_SYSFD int = iota
	RES_READ
	RES_WRITTEN
	RES_ERROR
)

type RPCResponse struct {
	SysFd   uintptr
	Read    []byte
	Written int
	Error   error
}

func (r *RPCResponse) Encode() *bytes.Buffer {
	format := fmt.Sprintf("%d %v %d %v\n", r.SysFd, r.Read, r.Written, r.Error)
	fmt.Println(format)
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return &buf
}

func DecodeRPCResponse(b *bytes.Buffer) *RPCResponse {
	cleanedBytes := bytes.TrimSpace(b.Bytes())
	parts := bytes.Split(cleanedBytes, []byte(" "))

	// default response
	res := &RPCResponse{
		SysFd:   0,
		Read:    []byte{},
		Written: 0,
		Error:   nil,
	}

	sysFd, err := ByteArrToPtr(parts[RES_SYSFD])
	if err == nil {
		res.SysFd = sysFd
	}

	read, err := hex.DecodeString(string(parts[RES_READ]))
	if err == nil {
		res.Read = read
	}

	written, err := strconv.ParseInt(string(parts[RES_WRITTEN]), 10, 0)
	if err == nil {
		res.Written = int(written)
	}

	if string(parts[RES_ERROR]) != "<nil>" {
		res.Error = fmt.Errorf("%v", string(parts[RES_ERROR]))
	}

	return res
}
