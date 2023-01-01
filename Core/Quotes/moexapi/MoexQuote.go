package moexquote

type MoexQuotesDTO struct {
	MarketData MarketData `json:"marketdata"`
}

type MarketData struct {
	Data [][]interface{} `json:"data"`
}

type MoexQuote struct {
	stockTicker  string
	CurrentQuote float64
	MarketCap    float64
}

// MARK: Implementation of the SimpleQuote interface
func (quote MoexQuote) Quote() float64 {
	return quote.CurrentQuote
}

func (quote MoexQuote) Ticker() string {
	return quote.stockTicker
}
