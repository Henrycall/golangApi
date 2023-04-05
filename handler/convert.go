package handler

import "fmt"

func Convert(amount float64, from string, to string, rate float64) (float64, string, error) {
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
		} else {
			return 0, "", fmt.Errorf("unsupported 'to' currency code: %s", to)
		}
	case "USD":
		if to == "BRL" {
			result = amount * rate
			symbol = "R$"
		} else {
			return 0, "", fmt.Errorf("unsupported 'to' currency code: %s", to)
		}
	case "EUR":
		if to == "BRL" {
			result = amount * (rate / 1.17)
			symbol = "R$"
		} else {
			return 0, "", fmt.Errorf("unsupported 'to' currency code: %s", to)
		}
	case "BTC":
		if to == "USD" {
			result = amount * rate
			symbol = "$"
		} else if to == "BRL" {
			result = amount * (rate * 280000)
			symbol = "R$"
		} else {
			return 0, "", fmt.Errorf("unsupported 'to' currency code: %s", to)
		}
	default:
		return 0, "", fmt.Errorf("unsupported 'from' currency code: %s", from)
	}

	return result, symbol, nil
}

