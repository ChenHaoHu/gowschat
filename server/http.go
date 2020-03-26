package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleHttp(c *gin.Context) {
	e := groupqueue.g

	for a, b := range e {
		log.Print(a)
		for _, d := range b.m {
			log.Println(d.Uid, d.Name, d.LoginTime)
		}

	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
