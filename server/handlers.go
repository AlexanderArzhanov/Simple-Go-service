package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	internallogic "simpleGoService/InternalLogic"
	"simpleGoService/dataAccess"
)

func getDefault(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Default")
}

func getAllEvents(ctx *gin.Context) {
	handlerRequest := dataAccess.ServiceDataConn.ReadAll()
	ctx.IndentedJSON(http.StatusOK, &handlerRequest)
}

func getEventByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	if len(id) > 1 {
		id = id[1:]
	}

	handlerRequest := dataAccess.ServiceDataConn.GetEventByID(id)
	ctx.JSON(http.StatusOK, &handlerRequest)
}

func deleteEventByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	dataAccess.ServiceDataConn.DeleteEventByID(id)
}

func createEvent(ctx *gin.Context) {

	event := new(internallogic.Event)
	ctx.BindJSON(event)

	handlerRequest := dataAccess.ServiceDataConn.CreateEvent(event)

	ctx.JSON(http.StatusOK, &handlerRequest)
}

func updateEventByID(ctx *gin.Context) {

	event := new(internallogic.Event)
	ctx.BindJSON(event)

	handlerRequest := dataAccess.ServiceDataConn.UpdateEvent(event)

	ctx.JSON(http.StatusOK, &handlerRequest)
}
