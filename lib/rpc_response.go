package lib

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
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

	sysFd := ByteArrToPtr(parts[0])
	fmt.Println(string(parts[2]))
	written, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		return &RPCResponse{Error: err}
	}

	var resErr error = nil
	if string(parts[3]) != "<nil>" {
		resErr = fmt.Errorf(string(parts[3]))
	}

	if len(parts[1]) <= 2 {
		return &RPCResponse{
			SysFd: sysFd,
			Read: []byte{},
			Written: int(written),
			Error: resErr,
		}
	}

	fmt.Printf("decode string: %s %d\n", string(parts[1]), len(parts[1]))
	read, err := hex.DecodeString(string(parts[1]))
	if err != nil {
		return &RPCResponse{Error: fmt.Errorf("Failed to decode read hex data")}
	}

	return &RPCResponse{
		SysFd:   sysFd,
		Read:    read,
		Written: int(written),
		Error:   resErr,
	}
}
