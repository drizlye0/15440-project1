package main

import (
	"fmt"

	"github.com/drizlye0/15440-project1/lib"
)

func main() {
	fileName := "hello.txt"
	sysFd, err := lib.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(sysFd)
	// req := &lib.RPCRequest{
	// 	CallSignature: "open",
	// 	FileName: "hello.txt",
	// }

	// buf := req.Encode()
	// originalReq := lib.DecodeRPCRequest(buf)
	// fmt.Println(originalReq)
}
