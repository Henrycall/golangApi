package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Conversation(ctx *gin.Context){	
	// Cria a conversão

	request := CreateConversionRequest{}
	amount := ctx.Param("amount")
	from := ctx.Param("from")
	to := ctx.Param("to")
	rate := ctx.Param("rate")
	amountFloat, err := strconv.ParseFloat(amount, 64)
	amountRate, err := strconv.ParseFloat(rate, 64)
	if err != nil {}
	result, symbol := Convert(amountFloat,from,to,amountRate)
	print(result,symbol)


	ctx.BindJSON(&request)
	logger.Infof("request reciviced : %+v " ,request)
}