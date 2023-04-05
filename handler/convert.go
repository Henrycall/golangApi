package handler

func Convert(amount float64, from string, to string, rate float64) (float64, string) {
	var result float64
	var symbol string
	
	switch from {
	case "BRL":
		if to == "USD" {
			result = amount / rate
			symbol = "$"
		} else if to == "EUR" {
			result = amount / (rate / 1.17)
			symbol = "€"
		}
	case "USD":
		if to == "BRL" {
			result = amount * rate
			symbol = "R$"
		}
	case "EUR":
		if to == "BRL" {
			result = amount * (rate / 1.17)
			symbol = "R$"
		}
	case "BTC":
		if to == "USD" {
			result = amount * rate
			symbol = "$"
		} else if to == "BRL" {
			result = amount * (rate * 280000)
			symbol = "R$"
		}
	}

	return result, symbol
}
