package router
import ("github.com/gin-gonic/gin")

func Initialize(){
	// initialize router
	r := gin.Default()
	// initialize routes
	initializeRouters(r)

	// run server
	r.Run(":3000") 
}

