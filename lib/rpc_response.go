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
	hexRead := "-"
	if len(r.Read) > 0 {
		hexRead = hex.EncodeToString(r.Read)
	}

	errStr := "-"
	if r.Error != nil {
		errStr = hex.EncodeToString([]byte(r.Error.Error()))
	}

	format := fmt.Sprintf("%d %s %d %s\n", r.SysFd, hexRead, r.Written, errStr)
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

	if string(parts[RES_READ]) != "-" {
		read, err := hex.DecodeString(string(parts[RES_READ]))
		if err == nil {
			res.Read = read
		}
	}

	written, err := strconv.ParseInt(string(parts[RES_WRITTEN]), 10, 0)
	if err == nil {
		res.Written = int(written)
	}

	if string(parts[RES_ERROR]) != "-" {
		errBytes, err := hex.DecodeString(string(parts[RES_ERROR]))
		if err == nil {
			res.Error = fmt.Errorf("%s", string(errBytes))
		}
	}

	return res
}
