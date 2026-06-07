package main

import (
	"os"
)

func main() {
	signature := os.Args[1:]
	if signature[0] == "foo" {
		foo()
	}
}
