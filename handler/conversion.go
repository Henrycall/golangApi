package handler

import (
	"github.com/gin-gonic/gin"
)

func Conversation(ctx *gin.Context){
	request := CreateConversionRequest{}
	ctx.BindJSON(&request)
	logger.Infof("request reciviced : %+v " ,request)
}