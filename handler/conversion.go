package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Conversation(ctx *gin.Context){	
	// Get path params
	request := CreateConversionRequest{}
	amount := ctx.Param("amount")
	from := ctx.Param("from")
	to := ctx.Param("to")
	rate := ctx.Param("rate")
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		logger.Errf("erro recebido : %v" , err)
	}
	amountRate, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		logger.Errf("erro recebido : %v" , err)
	}
	 // convert valors in funciton Convert , he returns result and symbol
	result, symbol,err := Convert(amountFloat,from,to,amountRate)
	
	if result ==  0|| symbol  == "" {
		response := gin.H{
			"err" : "Algum operador esta errado",
		}
		ctx.JSON(400, response)
	}else {
		response := gin.H{
			"valorConvertido" : result,
			"simboloMoeda" : symbol,
		}
		
		ctx.JSON(200, response)
	}

	ctx.BindJSON(&request)
	logger.Infof("request reciviced : %+v " ,request)
}