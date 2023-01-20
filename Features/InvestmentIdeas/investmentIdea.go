package investmentideas

import (
	quote "compound/Core/Quotes/Common"
	"fmt"
	"math"
	"time"
)

type SimpleQuote = quote.SimpleQuote
type Time = time.Time

type InvestmentIdea struct {
	Ticker           string  `json:"ticker"`
	CompanyName      string  `json:"companyName,omitempty"`
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

	idea.CalculateReturn()
}

func (idea *InvestmentIdea) CalculateUpside(currentQuote SimpleQuote) {
	if currentQuote.Quote() <= 0 {
		fmt.Printf("invalid quote for %v: %v", currentQuote.Ticker(), currentQuote.Quote())
	}

	upside := (idea.TargetPrice - currentQuote.Quote()) / currentQuote.Quote()
	upsideAsPercentage := upside * 100

	idea.Upside = upsideAsPercentage
}

func (idea *InvestmentIdea) CalculateReturn() (currentReturn float64, annualizedReturn float64, err error) {
	if idea.TargetPrice <= 0 || idea.PriceOnOpening <= 0 || idea.CurrentQuote <= 0 {
		return 0, 0, fmt.Errorf("unable to calculate the idea's return due to invalid data: target price (%v); price on opening (%v), current quote (%v)", idea.TargetPrice, idea.PriceOnOpening, idea.CurrentQuote)
	}

	currentReturn = (idea.CurrentQuote/idea.PriceOnOpening - 1) * 100                   //as percentage
	positionLifetime := float64(time.Since(idea.OpeningDate)) / float64(time.Hour) / 24 //in days
	annualizedReturn = calcAnnualizedReturn(currentReturn, positionLifetime)            //as percentage

	return currentReturn, annualizedReturn, nil

}

func calcAnnualizedReturn(totalReturn float64, daysHeld float64) float64 {
	return math.Pow((1+totalReturn/100), 365/daysHeld) - 1
}
