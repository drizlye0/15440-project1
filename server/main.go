package main

func main() {
	srv := NewTCPServer(8080, 5)
	srv.Listen()
}
