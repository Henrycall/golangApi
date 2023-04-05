package schemas

import (
	"gorm.io/gorm"
  )
  
  type Conversion struct {
	gorm.Model
	ID          int     
    From        string  
    To          string  
    Rate        float64 
    Converted   float64
  }

  type ConversionResponse struct {
	ID          int     `json:"id"`
    From        string  `json:"from"`
    To          string  `json:"to"`
    Rate        float64 `json:"rate"`
    Converted   float64 `json:"converted"`

  }

