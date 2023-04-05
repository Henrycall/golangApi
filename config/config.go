package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db * gorm.DB
	logger *Logger
)
func Init() error{
	var err error 
	// initialize SQLITE
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initialize sqllite : %v" , err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
func GetLogger(p string) * Logger {
	logger := Newlogger(p)
	return logger
}