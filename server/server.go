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

	// websocket echo
	r.Any("/ws/:token", HandleWS)

	// http echo
	r.GET("/http", HandleHttp)

	r.Run(":" + strconv.Itoa(port))
}
