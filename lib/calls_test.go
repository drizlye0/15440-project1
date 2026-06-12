package lib

import (
	"strings"
	"sync"
	"testing"
)

var testNumGoroutines = 5

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

func closeTestFile(t *testing.T, sysFd uintptr) {
	t.Helper()

	err := Close(sysFd, file_name)
	if err != nil {
		t.Errorf("Close() error: %v", err)
	}
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

func TestConcurrentOpen(t *testing.T) {
	var wg sync.WaitGroup
	fds := make([]uintptr, testNumGoroutines)
	errs := make([]error, testNumGoroutines)

	for i := 0; i < testNumGoroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			sysFd, err := Open(file_name)
			fds[idx] = sysFd
			errs[idx] = err
		}(i)
	}
	wg.Wait()

	for i := 0; i < testNumGoroutines; i++ {
		if errs[i] != nil {
			t.Errorf("goroutine %d: Open() error: %v", i, errs[i])
		}
		if fds[i] == 0 {
			t.Errorf("goroutine %d: invalid file descriptor", i)
		}
	}

	seen := make(map[uintptr]bool)
	for i, fd := range fds {
		if seen[fd] && errs[i] == nil {
			t.Errorf("goroutine %d: duplicate file descriptor %d", i, fd)
		}
		seen[fd] = true
	}

	for _, fd := range fds {
		if fd != 0 {
			closeTestFile(t, fd)
		}
	}
}

func TestConcurrentOpenOverload(t *testing.T) {
	var wg sync.WaitGroup
	numGoroutines := 20
	fds := make([]uintptr, numGoroutines)
	errs := make([]error, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			sysFd, err := Open(file_name)
			fds[idx] = sysFd
			errs[idx] = err
		}(i)
	}
	wg.Wait()

	successCount := 0
	failCount := 0
	for i := 0; i < numGoroutines; i++ {
		if errs[i] != nil || fds[i] == 0 {
			failCount++
		} else {
			successCount++
		}
	}

	if successCount == 0 {
		t.Error("expected at least one successful Open")
	}
	if failCount == 0 {
		t.Error("expected at least one failed Open (server maxConn=5 limit)")
	}

	for _, fd := range fds {
		if fd != 0 {
			closeTestFile(t, fd)
		}
	}
}

func TestConcurrentOpenClose(t *testing.T) {
	var wg sync.WaitGroup

	for round := 0; round < 3; round++ {
		for i := 0; i < testNumGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sysFd, err := Open(file_name)
				if err != nil {
					t.Errorf("Open() error: %v", err)
				}
				if sysFd == 0 {
					t.Error("invalid file descriptor")
				}
				if sysFd != 0 {
					closeTestFile(t, sysFd)
				}
			}()
		}
		wg.Wait()
	}
}
