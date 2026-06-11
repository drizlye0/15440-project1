# RPC Transparent Remote File Operations

Project 1 of [15-440](https://www.andrew.cmu.edu/course/15-440/) course.

This project is a simple RPC System for remote file operations built in Go.

Currently only support this file operations:
- Open
- Close
- Read
- Write

Project Structure:
```bash
├── client                          # client side app
│   ├── go.mod
│   └── main.go                     # entry point
├── lib                             # layer for communication between server and client
│   ├── calls.go
│   ├── calls_test.go
│   ├── go.mod
│   ├── rpc_request.go              # marshall and unmarshall logic for RPCRequest
│   ├── rpc_response.go             # marshall and unmarshall logic for RPCResponse
│   └── utils.go
├── server
│   ├── calls.go                    # native system calls for file operations
│   ├── go.mod
│   ├── handlers.go                 
│   ├── main.go                     # entry point
│   └── server.go                   # tcp server
```
