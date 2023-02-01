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

func newServer(addr string, port int) *Server {

	var server Server

	server.addr = addr
	server.port = port

	router := gin.Default()

	router.GET("/events", getAllEvents)
	router.GET("/events/:id", getEventByID)
	router.GET("/events/delete/:id", deleteEventByID)
	router.GET("/events/update/:id", updateEventByID)
	router.POST("/events/create", createEvent)

	server.router = router

	return &server
}

func (server *Server) start() {
	host := fmt.Sprintf("%s:%d", server.addr, server.port)
	server.router.Run(host)
}
