package main

import (
	"fmt"
	"log"

	"github.com/drizlye0/15440-project1/lib"
)

func main() {
	sysFd, err := lib.Open("hello.txt")
	if err != nil {
		log.Panic("Open request failed")
	}

	fmt.Println("Remote File Descritptor: %d", sysFd)

	err = lib.Close(sysFd, "hello.txt")
	if err != nil {
		log.Panic("Failed to close remote file")
	}

	fmt.Println("Closed file")
}
