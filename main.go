package main

import (
	"simpleGoService/dataAccess"
	"simpleGoService/server"
)

func main() {

	connData := make(map[string]string)

	connData["user"] = "admin"
	connData["password"] = "password"
	connData["addr"] = "localhost"
	connData["port"] = "5432"
	connData["dbname"] = "pgdb"

	dataAccess.ServiceDataConn = dataAccess.NewConnection()
	dataAccess.ServiceDataConn.Connect(connData)

	server := server.NewServer("", 8080)
	server.Start()

	server.Stop()
}
