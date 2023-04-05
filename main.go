package main

import (
	"fmt"

	"github.com/Henrycall/golangApi/config"
	"github.com/Henrycall/golangApi/router"
)

var (
	logger *config.Logger
)
	
func main()  {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		panic(err)
		logger.Errf("config initializaion error : %v" , err )
		fmt.Println(err)
		return
	}
	router.Initialize()
}
