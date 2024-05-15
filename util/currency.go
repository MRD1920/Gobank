package util

//consants for all supported currencies
const (
	USD = "USD"
	INT = "INR"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INT, CAD, EUR:
		return true
	}
	return false
}
