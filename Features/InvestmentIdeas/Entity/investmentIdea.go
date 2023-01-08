package investmentideas

import (
	quote "compound/Core/Quotes/Common"
	"fmt"
	"time"
)

type SimpleQuote = quote.SimpleQuote
type Time = time.Time

type InvestmentIdea struct {
	Ticker           string  `json:"ticker"`
	CompanyName      string  `json:"companyName"`
	Currency         string  `json:"currency"`
	TargetPrice      float64 `json:"targetPrice"`
	Upside           float64 `json:"upside"`
	PriceOnOpening   float64 `json:"priceOnOpening,omitempty"`
	CurrentQuote     float64 `json:"currentQuote,omitempty"`
	OpeningDate      Time    `json:"openingDate,omitempty"`
	InvestmentThesis string  `json:"thesis,omitempty"`
}

func (idea *InvestmentIdea) SetQuote(currentQuote SimpleQuote) {
	idea.CurrentQuote = currentQuote.Quote()
}

func (idea *InvestmentIdea) CalculateUpside(currentQuote SimpleQuote) {
	if currentQuote.Quote() <= 0 {
		fmt.Printf("invalid quote for %v: %v", currentQuote.Ticker(), currentQuote.Quote())
	}

	upside := (idea.TargetPrice - currentQuote.Quote()) / currentQuote.Quote()
	upsideAsPercentage := upside * 100

	idea.Upside = upsideAsPercentage
}
