package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	r := gin.Default()
	r.Use(Cors())
	r.Static("/wstest", "./dist/vuews/dist")
	r.Any("/ws/:token", HandleWS)
	r.POST("/chat/notice", HandleNotice)
	r.GET("/chat/group", HandleGroup)
	r.Run(":" + strconv.Itoa(port))
}
