package server

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	addr string
	port int

	router *gin.Engine
}

func NewServer(addr string, port int) *Server {

	var server Server

	server.addr = addr
	server.port = port

	router := gin.Default()

	router.GET("/", getDefault)
	router.GET("/events", getAllEvents)
	router.GET("/events/get/:id", getEventByID)
	router.DELETE("/events/delete/:id", deleteEventByID)
	router.POST("/events/update", updateEventByID)
	router.POST("/events/create", createEvent)

	server.router = router

	return &server
}

func (server *Server) Start() {
	host := fmt.Sprintf("%s:%d", server.addr, server.port)
	server.router.Run(host)
}

func (server *Server) Stop() {
	fmt.Println("Server stopped")
}
