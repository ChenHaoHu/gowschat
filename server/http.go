package server

import (
	"io"

	"github.com/gin-gonic/gin"
)

func HandleHttp(c *gin.Context) {
	io.Copy(c.Writer, c.Request.Body)
}
