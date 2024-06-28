package endpoint

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	signingKey  = "c38jxmk"
	userContext = "user_id"
)

func (e *Endpoint) Middleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid header")
		return
	}
	arrOfHeader := strings.Split(authHeader, " ")
	if len(arrOfHeader) != 2 || arrOfHeader[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid headerr")
		return
	}
	token := arrOfHeader[1]
	correctToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid headerrr")
		return
	}
	if claims, ok := correctToken.Claims.(jwt.MapClaims); ok && correctToken.Valid {
		c.Set(userContext, claims["user_id"])
	}
}

func (e *Endpoint) GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userContext)
	if !ok {
		return 0, errors.New("user id is not found")
	}
	fmt.Println(id, reflect.TypeOf(id))
	floatID, ok := id.(float64)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}
	intID := int(floatID)
	return intID, nil
}
