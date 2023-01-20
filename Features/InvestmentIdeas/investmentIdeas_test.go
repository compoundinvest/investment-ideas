package investmentideas

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGetIdeasList(t *testing.T) {

	ideasFileContent, err := os.ReadFile("investmentideas.json")
	if err != nil {
		t.Errorf("Error while attempting to read the investment ideas list file: %v\n", err)
	}

	var ideas InvestmentIdeas
	err = json.Unmarshal(ideasFileContent, &ideas)
	if err != nil {
		t.Errorf("Unable to unmarshal the JSON containing the list of investment ideas: %v", err)
	}

	quotes := fetchQuotesForIdeas(&ideas)
	ideas.CalculateUpsides(quotes)

	for _, idea := range ideas.Ideas {
		fmt.Println("Ticker: ", idea.Ticker, ". Opened on: ", idea.OpeningDate)
	}

	for _, idea := range ideas.Ideas {
		if idea.Ticker == "" {
			t.Error("Found an idea with an empty ticker")
		}
		if idea.Currency == "" {
			t.Errorf("Found an idea with no currency: %v", idea.Ticker)
		}
		if idea.TargetPrice <= 0 {
			t.Errorf("Found an idea with an invalid target price: %v. Ticker: %v", idea.TargetPrice, idea.Ticker)
		}
		if idea.InvestmentThesis == "" {
			t.Errorf("Found an idea with missing thesis: %v", idea.Ticker)
		}
		if idea.CompanyName == "" {
			t.Errorf("Found an idea with missing company name:: %v", idea.Ticker)
		}

		//Validate the opening date by checking if it's earlier than 2 AD
		layout := "2006-01-02T15:04:05.000Z"
		ancientDateAsString := "0002-01-01T15:04:05.000Z"
		ancientDate, _ := time.Parse(layout, ancientDateAsString)
		if idea.OpeningDate.Before(ancientDate) {
			t.Errorf("Found an idea with invalid date: %v. Ticker: %v", idea.OpeningDate, idea.Ticker)
		}

		if idea.PriceOnOpening <= 0 {
			t.Errorf("Found an idea with an invalid price on opening: %v. Ticker: %v", idea.PriceOnOpening, idea.Ticker)
		}
	}
}
