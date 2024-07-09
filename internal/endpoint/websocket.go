package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	clients map[int]*websocket.Conn
}

type WsMessage struct {
	Text   string `json:"text"`
	UserId int    `json:"user_id"`
}

var rooms = make(map[int]*Room)

func (e *Endpoint) WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	userId, err := strconv.Atoi(c.Param("user"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	roomId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if _, ok := rooms[roomId]; !ok {
		rooms[roomId] = &Room{
			clients: map[int]*websocket.Conn{
				userId: conn,
			},
		}
	} else {
		rooms[roomId].clients[userId] = conn
	}

	conn.SetCloseHandler(func(code int, text string) error {
		delete(rooms[roomId].clients, userId)
		return nil
	})
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(rooms[roomId].clients, userId)
			break
		}
		data, err := json.Marshal(WsMessage{string(msg), userId})
		if err != nil {
			delete(rooms[roomId].clients, userId)
			break
		}
		for _, client := range rooms[roomId].clients {
			//if id != userId {
			err = client.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				delete(rooms[roomId].clients, userId)
				break
			}
			//}

		}
	}

}
