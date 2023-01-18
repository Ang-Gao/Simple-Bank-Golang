package util

const (
	USD = "USD"
	EUR = "EUR"
	CNY = "CNY"
	AED = "AED"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	/*
		The switch statement is used to perform different actions based on different conditions.
		Each case branch is unique, and it is tested one by one from top to bottom until it matches.
	*/
	switch currency {
	case USD, EUR, CNY, AED, CAD:
		return true
	default:
		return false
	}

}
