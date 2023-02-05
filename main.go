package main

import "simpleGoService/server"

func main() {

	server := server.NewServer("", 8080)
	server.Start()
}
