package schemas

import (
	"gorm.io/gorm"
  )
  
  type Opening struct {
	gorm.Model
	Role  string
	Comapany string
	Location string
	Remite bool
	Link string
	Salary int64
  }

