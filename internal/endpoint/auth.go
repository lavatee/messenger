package endpoint

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SignUpInput struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RefreshInput struct {
	RefreshToken string `json:"refresh"`
}

func (e *Endpoint) SignUp(c *gin.Context) {
	var input SignUpInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := e.services.SignUp(input.Username, input.Name, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}
func (e *Endpoint) SignIn(c *gin.Context) {
	var input SignInInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, name, access, refresh, err := e.services.SignIn(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"AccessToken":  access,
		"RefreshToken": refresh,
		"id":           id,
		"name":         name,
	})
}
func (e *Endpoint) GetUser(c *gin.Context) {
	userID, ok := c.Params.Get("id")
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "invalid request params")
		return
	}
	intID, err := strconv.Atoi(userID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	username, name, err := e.services.GetUserById(intID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"username": username,
		"name":     name,
	})
}
func (e *Endpoint) Refresh(c *gin.Context) {
	var input RefreshInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	access, refresh, err := e.services.Auth.Refresh(input.RefreshToken)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"AccessToken":  access,
		"RefreshToken": refresh,
	})
}

type PutInput struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func (e *Endpoint) PutUser(c *gin.Context) {
	var input PutInput
	err := c.BindJSON(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := e.GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = e.services.PutUser(input.Username, input.Name, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
