package quote

type SimpleQuote interface {
	Quote() float64
	Ticker() string
}

// Converts slices of structs implementing SimpleQuote to a slice of SimpleQuote objects
func ConvertToSimpleQuote[T SimpleQuote](quotes []T) []SimpleQuote {
	quotesAsInterface := []SimpleQuote{}
	for _, quote := range quotes {
		quotesAsInterface = append(quotesAsInterface, quote)
	}
	return quotesAsInterface
}
