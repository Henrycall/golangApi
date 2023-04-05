package router

import (
	"github.com/Henrycall/golangApi/handler"
	"github.com/gin-gonic/gin"
)
func initializeRouters(router *gin.Engine){
	v1 := router.Group("/api/v1")
	
		v1.GET("/opening" , handler.ShowOpeningHandler)
		v1.POST("/opening" , handler.CreateOpeningHandler)
		v1.PUT("/opening" , handler.CreateOpeningHandler)
		v1.DELETE("/opening" , handler.DeletepeningHandler)
		v1.GET("/openings" , handler.ListOpeningHandler)
	
}