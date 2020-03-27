package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGroup(c *gin.Context) {
	groups := groupqueue.g
	data1 := make(map[string]interface{})

	for index, key := range groups {
		data1[index] = key.m
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"groups": data1,
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
