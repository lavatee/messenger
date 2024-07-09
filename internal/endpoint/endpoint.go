package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavatee/messenger/internal/service"
	"github.com/sirupsen/logrus"
)

type Endpoint struct {
	services *service.Service
}

func NewEndpoint(services *service.Service) *Endpoint {
	return &Endpoint{
		services: services,
	}
}

func (e *Endpoint) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		ctx.Writer.Header().Set("Acces-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		ctx.Writer.Header().Set("Access-COntrol-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
	})
	auth := router.Group("/auth")
	{
		auth.POST("/signup", e.SignUp)
		auth.POST("/signin", e.SignIn)
		auth.POST("/refresh", e.Refresh)
	}
	api := router.Group("/api", e.Middleware)
	{
		api.PUT("/users", e.PutUser)
		api.GET("/users/:id", e.GetUser)
		api.POST("/chats/:id", e.PostChat)
		api.GET("/chats", e.GetChats)
		api.DELETE("/chats/:id", e.DeleteChat)
		api.POST("/messages", e.PostMessage)
		api.GET("/messages/:seconduserid", e.GetMessages)
		api.DELETE("/messages/:id", e.DeleteMessage)
		api.POST("/rooms/join", e.JoinRoom)
		api.POST("/rooms/leave/:id", e.LeaveRoom)
		api.POST("/rooms/leaveall/:id", e.LeaveMatchMaking)
	}
	ws := router.Group("/ws")
	{
		ws.GET("/rooms/:id/:user", e.WebSocketHandler)
		ws.GET("/chats/:id/:user", e.WebSocketHandler)
	}
	return router
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}
