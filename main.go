package main

import (
	"simpleGoService/dataAccess"
	"simpleGoService/server"
)

func main() {

	dataAccess.NewPostgresConnection()

	server := server.NewServer("", 8080)
	server.Start()

	server.Stop()

}
