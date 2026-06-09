package lib

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
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

	if len(parts) < 4 {
		return nil
	}

	sysFd := ByteArrToPtr(parts[0])

	hexRead := string(parts[1])
	read, err := hex.DecodeString(hexRead)
	if err != nil {
		return nil
	}

	written := binary.BigEndian.Uint32(parts[2])
	err = fmt.Errorf(string(parts[3]))

	return &RPCResponse{
		SysFd:   sysFd,
		Read:    read,
		Written: int(written),
		Error:   err,
	}
}
