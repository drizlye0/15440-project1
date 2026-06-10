.PHONY: server client

# Target to run all Go files in the server directory
server:
	go run ./server

# Target to run all Go files in the client directory
client:
	go run ./client
