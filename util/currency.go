package util

const (
	EUR = "EUR"
	USD = "USD"
	IDR = "IDR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case EUR, USD, IDR:
		return true
	}
	return false
}
