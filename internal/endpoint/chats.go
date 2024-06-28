package endpoint

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e *Endpoint) GetChats(c *gin.Context) {
	id, err := e.GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userChats, err := e.services.GetUserChats(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"chats": userChats,
	})
}
func (e *Endpoint) PostChat(c *gin.Context) {
	id, err := e.GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	secondUserId := c.Param("id")
	intSecondUserId, err := strconv.Atoi(secondUserId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	chatId, err := e.services.CreateChat(id, intSecondUserId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"chat_id": chatId,
	})
}
func (e *Endpoint) DeleteChat(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = e.services.DeleteChat(intId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
