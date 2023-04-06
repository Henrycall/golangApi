package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Henrycall/golangApi/schemas"
	"strconv"
)

func Conversation(ctx *gin.Context){	
	// Get path params
	// request := CreateConversionRequest{}
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
	if err != nil {
		ctx.JSON(400, err.Error())
	}else {
		response := gin.H{
			"valorConvertido" : result,
			"simboloMoeda" : symbol,
		}
		conversion := &schemas.Conversion{Amount: amountFloat, FromCurrency: from, ToCurrency: to, Rate: amountRate, Result: result, Symbol: symbol}

		if err := db.Create(conversion).Error; err != nil {
			logger.Errf("error create opening :  %v",  err.Error() )
		}
		ctx.JSON(200, response)
		
	}

	//ctx.BindJSON(&request)
	 //logger.Infof("request reciviced : %+v " ,request)
}