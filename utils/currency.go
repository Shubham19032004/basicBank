package utils



const (
	INR = "INR"
)

func IsSupportedCurrency(currency string) bool{
	switch currency{
	case INR:
		return true
	}
	return false
}