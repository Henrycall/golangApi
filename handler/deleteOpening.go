package handler

import (
	"github.com/gin-gonic/gin"
)


func DeletepeningHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message" : "DELETE JOB",
	})
}