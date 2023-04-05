package config

import (
	"os"
	"github.com/Henrycall/golangApi/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB , error){
	logger := GetLogger("sqlite")
	dbpath := "./db/main.db"
	// check if database file exists
	_,err := os.Stat(dbpath)
	if os.IsNotExist(err){
		logger.Info("database file notfoud, creating ...")
		// crete database file and directory
		err = os.MkdirAll("./db" , os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbpath)
		if err != nil {
			return nil,err
		}
		file.Close()
	}

	// create database and connect
	db , err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		logger.Errf("sqlite opning error:  %v" , err)
		return nil,err
	}
	err = db.AutoMigrate(schemas.Conversion{})
	if err != nil {
		logger.Errf("sqlite automigration error : %v" , err)
	}
	return db,nil
}