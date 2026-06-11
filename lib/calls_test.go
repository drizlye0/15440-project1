package lib

import (
	"strings"
	"testing"
)

var file_name string = "test.txt"
var file_content string = "This is a test file"

var write_value string = "(written)"
var content_written string = file_content + write_value

func openTestFile(t *testing.T) uintptr {
	t.Helper()

	sysFd, err := Open(file_name)
	if err != nil {
		t.Errorf("Open() error: %v", err)
	}

	if sysFd == 0 {
		t.Error("invalid file descriptor")
	}

	return sysFd
}

func TestOpen(t *testing.T) {
	_ = openTestFile(t)
}

func TestClose(t *testing.T) {
	sysFd := openTestFile(t)

	err := Close(sysFd, file_name)
	if err != nil {
		t.Errorf("Close() error: %v", err)
	}
}

func TestRead(t *testing.T) {
	data, err := Read(file_name)
	if err != nil {
		t.Errorf("Read() error: %v", err)
	}

	dataStr := strings.TrimSpace(string(data))
	if dataStr != file_content {
		t.Errorf("Invalid file content")
	}
}

func TestWrite(t *testing.T) {
	lines, err := Write(file_name, []byte(write_value))
	if err != nil {
		t.Errorf("Write() error: %v", err)
	}

	if lines < 0 {
		t.Error("Error: Invalid number of lines written")
	}
}
