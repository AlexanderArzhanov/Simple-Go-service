package main

import (
	"simpleGoService/dataAccess"
	"simpleGoService/server"
)

func main() {

	dataAccess.ServiceDataConn = &dataAccess.PostgresConnection{}

	dataAccess.NewPostgresConnection("admin", "password", "localhost", 5432, "pgdb")

	server := server.NewServer("", 8080)
	server.Start()

	server.Stop()
}
