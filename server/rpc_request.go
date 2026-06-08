package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

type RPCRequest struct {
	callSignature string
	sysFd         uintptr
	fileName      string
	data          []byte
}

func (r *RPCRequest) encode() bytes.Buffer {
	hexData := hex.EncodeToString(r.data)
	format := fmt.Sprintf("%s %d %s %s", r.callSignature, r.sysFd, r.fileName, hexData)
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return buf
}

func decodeRPCRequest(b *bytes.Buffer) *RPCRequest {
	cleanedBytes := bytes.TrimSpace(b.Bytes())
	parts := bytes.Split(cleanedBytes, []byte(" "))

	if len(parts) < 4 {
		log.Println("Failed to parse request: invalid format or missing arguments")
		return nil
	}

	ptrArr := parts[1]
	for i := 0; i < 8-len(ptrArr); i++ {
		ptrArr = append([]byte{0x0, 0x0}, ptrArr...)
	}
	ptr := binary.BigEndian.Uint64(ptrArr)

	hexStr := string(parts[3])
	decodedData, err := hex.DecodeString(hexStr)
	if err != nil {
		log.Printf("Failed to decode hex string data: %v\n", err)
		return nil
	}

	return &RPCRequest{
		callSignature: string(parts[0]),
		sysFd:         uintptr(ptr),
		fileName:      string(parts[2]),
		data:          decodedData,
	}
}
