package lib

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

type RPCRequest struct {
	CallSignature string
	SysFd         uintptr
	FileName      string
	Data          []byte
}

func (r *RPCRequest) Encode() *bytes.Buffer {
	hexData := hex.EncodeToString(r.Data)
	format := fmt.Sprintf("%s %d %s %s\n", r.CallSignature, r.SysFd, r.FileName, hexData)
	fmt.Println(format)
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return &buf
}

func DecodeRPCRequest(b *bytes.Buffer) *RPCRequest {
	cleanedBytes := bytes.TrimSpace(b.Bytes())
	parts := bytes.Split(cleanedBytes, []byte(" "))

	if len(parts) < 4 {
		log.Println("Failed to parse request: invalid format or missing arguments")
		return nil
	}

	sysFd := ByteArrToPtr(parts[1])

	hexData := string(parts[3])
	data, err := hex.DecodeString(hexData)
	if err != nil {
		log.Printf("Failed to decode hex string data: %v\n", err)
		return nil
	}

	return &RPCRequest{
		CallSignature: string(parts[0]),
		SysFd:         sysFd,
		FileName:      string(parts[2]),
		Data:          data,
	}
}
