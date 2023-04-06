package schemas

import (
	"gorm.io/gorm"
  )
  type Conversion struct {
    gorm.Model
    Amount       float64
    FromCurrency string
    ToCurrency   string
    Rate         float64
    Result       float64
    Symbol       string
  }

