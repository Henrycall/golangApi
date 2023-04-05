package handler

type CreateConversionRequest struct {
	Amount      float64 `json:"amount"`
    From        string  `json:"from"`
    To          string  `json:"to"`
    Rate        float64 `json:"rate"`
    Converted   float64 `json:"converted"`
}

