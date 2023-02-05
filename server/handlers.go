package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

    dataAccess "simpleGoService/dataAccess"
)

func getDefault(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Default")
}

func getAllEvents(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, dataAccess.ReadAll)
}

func getEventByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("det by id %s", id))
}

func deleteEventByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("delete by id %s", id))
}

func updateEventByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("update by id %s", id))
}

func createEvent(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "create")
}
