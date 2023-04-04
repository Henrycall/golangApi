package router

import (
	"github.com/gin-gonic/gin"
)
func initializeRouters(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening" , func(ctx *gin.Context){
			ctx.JSON(200,gin.H{
				"message" : "GET JOB",
			})

		})
		v1.POST("/opening" , func(ctx *gin.Context){
			ctx.JSON(200,gin.H{
				"message" : "POST JOB",
			})

		})
		v1.PUT("/opening" , func(ctx *gin.Context){
			ctx.JSON(200,gin.H{
				"message" : "PUT JOB",
			})

		})
		v1.DELETE("/opening" , func(ctx *gin.Context){
			ctx.JSON(200,gin.H{
				"message" : "DELETE JOB",
			})

		})
		v1.GET("/openings" , func(ctx *gin.Context){
			ctx.JSON(200,gin.H{
				"message" : "GET JOB",
			})

		})

	}
	
}