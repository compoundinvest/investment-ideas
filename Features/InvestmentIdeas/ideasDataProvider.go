package investmentideas

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/compoundinvest/invest-core/quote/moexapi"
	"github.com/compoundinvest/invest-core/quote/yahooapi"
)

func readIdeasFile() ([]byte, error) {
	ideasFileContent, err := os.Open("Features/InvestmentIdeas/investmentideas.json")
	if err != nil {
		fmt.Printf("Error while attempting to read the investment ideas list file: %v\n", err)
		return nil, fmt.Errorf("unable to open the investment ideas file")
	}

	byteValue, err := io.ReadAll(ideasFileContent)
	if err != nil {
		fmt.Println("Error while attempting to convert the ideas list file into a byte array: ", err)
		return nil, fmt.Errorf("unable to read the contents of the investment ideas file")
	}

	return byteValue, nil
}

func getLocalIdeasList() InvestmentIdeas {
	unparsedIdeasData, err := readIdeasFile()
	if err != nil {
		return InvestmentIdeas{"Error", []InvestmentIdea{}}
	}

	var ideas InvestmentIdeas
	err = json.Unmarshal(unparsedIdeasData, &ideas)
	if err != nil {
		fmt.Println("Unable to unmarshal the JSON containing the list of investment ideas: ", err)
	}

	return ideas
}

func fetchQuotesForIdeas(ideas *InvestmentIdeas) []moexapi.SimpleQuote {
	quotes := []moexapi.SimpleQuote{}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		yahooQuotes := yahooapi.FetchQuotes(ideas.TickersOfSpecificCurrency("USD"))
		quotes = append(quotes, yahooQuotes...)
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		moexQuotes := moexapi.FetchQuotes(ideas.TickersOfSpecificCurrency("RUB"))
		quotes = append(quotes, moexQuotes...)
	}()
	waitGroup.Wait()

	return quotes
}

func GetInvestmentIdeas() InvestmentIdeas {
	ideas := getLocalIdeasList()
	quotes := fetchQuotesForIdeas(&ideas)
	ideas.CalculateUpsides(quotes)

	return ideas
}
