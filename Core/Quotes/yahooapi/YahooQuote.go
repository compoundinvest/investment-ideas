package yahooapi

type YahooQuotesDTO struct {
	QuoteResponse QuoteResponseDTO `json:"quoteResponse"`
}

type QuoteResponseDTO struct {
	Result []YahooQuote `json:"result"`
}

type YahooQuote struct {
	Symbol             string  `json:"symbol"`
	DisplayName        string  `json:"displayName"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
	PriceToBook        float64 `json:"priceToBook"`
	ForwardPE          float64 `json:"forwardPE"`
	MarketCap          float64 `json:"marketCap"`
	TrailingPE         float64 `json:"trailingPE"`
}

func (quote YahooQuote) Ticker() string {
	return quote.Symbol
}

func (quote YahooQuote) Quote() float64 {
	return quote.RegularMarketPrice
}
