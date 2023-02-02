package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 1)
}

func getEventByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 2)
}

func deleteEventByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 3)
}

func updateEventByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 4)
}

func createEvent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 5)
}
