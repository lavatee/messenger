package endpoint

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lavatee/messenger"
)

func (e *Endpoint) GetMessages(c *gin.Context) {
	secondUserId := c.Param("seconduserid")
	intId, err := strconv.Atoi(secondUserId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	firstUserId, err := e.GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	messages, err := e.services.GetChatMessages(firstUserId, intId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"messages": messages,
	})
}
func (e *Endpoint) PostMessage(c *gin.Context) {
	var message messenger.Message
	if err := c.BindJSON(&message); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := e.services.CreateMessage(message)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"chat_id": id,
	})
}
func (e *Endpoint) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = e.services.DeleteMessage(intId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
