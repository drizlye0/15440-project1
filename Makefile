.PHONY: server client

clean:
	rm -rf bin

build:
	mkdir bin
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

test:
	go test -v -race -count=1 ./lib

# Target to run all Go files in the server directory
server:
	go run ./server

# Target to run all Go files in the client directory
client:
	go run ./client
