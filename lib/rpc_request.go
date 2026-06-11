package lib

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

const (
	REQ_SIGNATURE int = iota
	REQ_SYSFD
	REQ_FILENAME
	REQ_DATA
)

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
	buf := bytes.Buffer{}
	buf.Write([]byte(format))
	return &buf
}

func DecodeRPCRequest(b *bytes.Buffer) *RPCRequest {
	cleanedBytes := bytes.TrimSpace(b.Bytes())
	parts := bytes.Split(cleanedBytes, []byte(" "))

	signature := string(parts[REQ_SIGNATURE])
	switch signature {
	case "open":
		return newOpenReadRequest(signature, parts[REQ_FILENAME])

	case "close":
		return newCloseRequest(signature, parts[REQ_SYSFD], parts[REQ_FILENAME])

	case "read":
		return newOpenReadRequest(signature, parts[REQ_FILENAME])

	case "write":
		return newWriteRequest(signature, parts[REQ_SYSFD], parts[REQ_FILENAME], parts[REQ_DATA])

	default:
		return nil
	}
}

func newOpenReadRequest(signature string, fileName []byte) *RPCRequest {
	return &RPCRequest{
		CallSignature: signature,
		FileName:      string(fileName),
	}
}

func newCloseRequest(signature string, fdBytes []byte, fileName []byte) *RPCRequest {
	sysFd, err := ByteArrToPtr(fdBytes)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &RPCRequest{
		CallSignature: signature,
		SysFd:         sysFd,
		FileName:      string(fileName),
	}
}

func newWriteRequest(signature string, fdBytes []byte, fileName []byte, data []byte) *RPCRequest {
	sysFd, err := ByteArrToPtr(fdBytes)
	if err != nil {
		log.Println(err)
		return nil
	}

	decodedData, err := hex.DecodeString(string(data))
	if err != nil {
		log.Println(err)
		return nil
	}

	return &RPCRequest{
		CallSignature: signature,
		SysFd:         sysFd,
		FileName:      string(fileName),
		Data:          decodedData,
	}
}
