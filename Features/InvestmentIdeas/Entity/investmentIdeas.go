package investmentideas

import (
	"fmt"
	"sort"
)

type InvestmentIdeas struct {
	Author string           `json:"author"`
	Ideas  []InvestmentIdea `json:"ideas"`
}

func (ideas InvestmentIdeas) TickersOfSpecificCurrency(currency string) []string {
	targetTickers := []string{}
	for _, idea := range ideas.Ideas {
		if idea.Currency == currency {
			targetTickers = append(targetTickers, idea.Ticker)
		}
	}
	return targetTickers
}

func (ideas *InvestmentIdeas) CalculateUpsides(quotes []SimpleQuote) {
	for i := 0; i < len(ideas.Ideas); i++ {
		quote, err := findQuoteFor(ideas.Ideas[i].Ticker, quotes)
		if err != nil {
			continue
		}
		ideas.Ideas[i].CalculateUpside(quote)
	}
	sort.Slice(ideas.Ideas, func(i, j int) bool {
		return ideas.Ideas[i].Upside > ideas.Ideas[j].Upside
	})
}

func findQuoteFor(ticker string, quotes []SimpleQuote) (SimpleQuote, error) {
	for _, quote := range quotes {
		if quote.Ticker() == ticker {
			return quote, nil
		}
	}
	return nil, fmt.Errorf("no quote was found for %v", ticker)
}
