package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func HandleGroup(c *gin.Context) {
	groups := group.g
	//data1 := make(map[string]interface{})

	for index, key := range groups {
		//data1[index] = key.m
		fmt.Println(index, " ----- ")
		fmt.Printf("%+v\n", key)

	}

	//c.JSON(200, gin.H{
	//	"status": "ok",
	//	"groups": data1,
	//})
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

	SendMsg(&Msg{
		Uid:      "NOTICE",
		ToUid:    uid,
		Gid:      gid,
		Msg:      msg,
		SendType: sendtype,
	}, nil)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
