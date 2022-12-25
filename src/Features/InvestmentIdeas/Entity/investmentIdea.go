package investmentideas

import (
	quote "compound/Core/Quotes/Common"
	"fmt"
	"time"
)

type SimpleQuote = quote.SimpleQuote

type InvestmentIdea struct {
	Ticker           string `json:"ticker"`
	companyName      string
	Currency         string `json:"currency"`
	priceOnOpening   float64
	TargetPrice      float64 `json:"targetPrice"`
	Upside           float64 `json:"upside"`
	openingDate      time.Time
	investmentThesis string
}

func (idea *InvestmentIdea) CalculateUpside(currentQuote SimpleQuote) {
	if currentQuote.Quote() <= 0 {
		fmt.Printf("invalid quote for %v: %v", currentQuote.Ticker(), currentQuote.Quote())
	}

	upside := (idea.TargetPrice - currentQuote.Quote()) / currentQuote.Quote()
	upsideAsPercentage := upside * 100

	idea.Upside = upsideAsPercentage
}
