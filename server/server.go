package server

import (
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

func StartServer(port int) {
	r := gin.Default()

	r.Any("/ws/:token", HandleWS)

	r.POST("/chat/notice", HandleNotice)
	r.GET("/chat/group", HandleGroup)

	r.Run(":" + strconv.Itoa(port))
}
