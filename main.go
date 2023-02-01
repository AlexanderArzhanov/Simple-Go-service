package main

import "simpleGoService/server"

func main() {
	server := server.NewServer("localhost", 3030)
	server.Start()
}
