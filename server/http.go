package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGroup(c *gin.Context) {
	groups := groupqueue.g
	data := make(map[string]interface{})

	for index, key := range groups {
		data[index] = key.m
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": data,
	})
}

func HandleNotice(c *gin.Context) {

	sendtype, _ := strconv.Atoi(c.PostForm("sendtype"))
	gid := c.PostForm("gid")
	msg := c.PostForm("msg")
	uid := c.PostForm("uid")

	if sendtype < N2G {
		c.JSON(200, gin.H{
			"status": "sendtype can not smaller by " + string(N2G),
		})
		return
	}

	AddMsg(&Msg{
		Uid:      "NOTICE",
		ToUid:    uid,
		Gid:      gid,
		Msg:      msg,
		SendType: sendtype,
	})

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
