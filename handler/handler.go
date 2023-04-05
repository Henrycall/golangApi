package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message" : "POST JOB",
	})
}

func ShowOpeningHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message" : "GET JOB",
	})
}

func DeletepeningHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message" : "DELETE JOB",
	})
}

func UpdateOpeningHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message" : "PUT JOB",
	})
}

func ListOpeningHandler(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"message" : "GET JOB",
	})
}