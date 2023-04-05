package handler

import (
	"gorm.io/gorm"
	"github.com/Henrycall/golangApi/config"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler(){
	logger = config.GetLogger("Handler")
	db = config.GetSQLite()
}
	
