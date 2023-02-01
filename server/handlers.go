package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 1)
}

func getEventByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 1)
}

func deleteEventByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 1)
}

func updateEventByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 1)
}

func createEvent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 1)
}
