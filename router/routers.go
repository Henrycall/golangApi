package router

import (
	"github.com/Henrycall/golangApi/handler"
	"github.com/gin-gonic/gin"
)
func initializeRouters(router *gin.Engine){
	handler.InitializeHandler()
	v1 := router.Group("/exchange")
	
		v1.GET("/:amount/:from/:to/:rate" , handler.Conversation)
	
}