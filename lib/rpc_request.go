package lib

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

var CALL_SIGNATURE int = 0
var SYS_FD int = 1
var FILE_NAME int = 2
var DATA int = 3

type RPCRequest struct {
	CallSignature string
	SysFd         uintptr
	FileName      string
	Data          []byte
}

func (r *RPCRequest) Encode() *bytes.Buffer {
	hexData := "-"
	if len(r.Data) > 0 {
		hexData = hex.EncodeToString(r.Data)
	}

	format := fmt.Sprintf("%s %d %s %s\n", r.CallSignature, r.SysFd, r.FileName, hexData)
	fmt.Println(format)
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return &buf
}

func DecodeRPCRequest(b *bytes.Buffer) *RPCRequest {
	cleanedBytes := bytes.TrimSpace(b.Bytes())
	parts := bytes.Split(cleanedBytes, []byte(" "))

	signature := string(parts[CALL_SIGNATURE])
	fileName := string(parts[FILE_NAME])
	sysFd := ByteArrToPtr(parts[SYS_FD])
	hexData := string(parts[DATA])

	if hexData == "-" {
		return &RPCRequest{
			CallSignature: signature,
			FileName:      fileName,
			SysFd:         sysFd,
			Data:          []byte{},
		}
	}

	data, err := hex.DecodeString(hexData)
	if err != nil {
		return nil
	}

	return &RPCRequest{
		CallSignature: signature,
		FileName:      fileName,
		SysFd:         sysFd,
		Data:          data,
	}

	// if signature == "open" {
	// 	return &RPCRequest{
	// 		CallSignature: signature,
	// 		FileName:      fileName,
	// 	}
	// }

	// if signature == "close" {
	// 	sysFd := ByteArrToPtr(parts[SYS_FD])
	// 	return &RPCRequest{
	// 		CallSignature: signature,
	// 		SysFd:         sysFd,
	// 		FileName:      fileName,
	// 	}
	// }

	// if signature == "write" {
	// 	hexData := string(parts[DATA])
	// 	data, err := hex.DecodeString(hexData)
	// 	if err != nil {
	// 		return nil
	// 	}

	// 	sysFd := ByteArrToPtr(parts[SYS_FD])

	// 	return &RPCRequest{
	// 		CallSignature: signature,
	// 		SysFd:         sysFd,
	// 		FileName:      fileName,
	// 		Data:          data,
	// 	}
	// }

	// if signature == "read" {
	// 	return &RPCRequest{
	// 		CallSignature: signature,
	// 		FileName:      fileName,
	// 	}
	// }

	// return nil

	// sysFd := ByteArrToPtr(parts[1])

	// hexData := string(parts[3])
	// data, err := hex.DecodeString(hexData)
	// if err != nil {
	// 	log.Printf("Failed to decode hex string data: %v\n", err)
	// 	return nil
	// }

	// return &RPCRequest{
	// 	CallSignature: string(parts[0]),
	// 	SysFd:         sysFd,
	// 	FileName:      string(parts[2]),
	// 	Data:          data,
	// }
}
