package handler

import (
	"github.com/gin-gonic/gin"
)

func Conversation(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message" : "POST JOB",
	})
}