package endpoint

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) JoinRoom(c *gin.Context) {
	intId, err := e.GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	roomId, err := e.services.JoinRoom(intId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"room_id": roomId,
	})
}
func (e *Endpoint) LeaveRoom(c *gin.Context) {
	id, err := e.GetUserId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	roomId := c.Param("id")
	intId, err := strconv.Atoi(roomId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	currentRoomId, err := e.services.LeaveRoom(id, intId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"room_id": currentRoomId,
	})
}
func (e *Endpoint) LeaveMatchMaking(c *gin.Context) {
	id, err := e.GetUserId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	roomId := c.Param("id")
	intId, err := strconv.Atoi(roomId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = e.services.LeaveMatchMaking(id, intId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
