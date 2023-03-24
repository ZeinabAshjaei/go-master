package utils

const (
	EUR = "EUR"
	CAD = "CAD"
	USD = "USD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
