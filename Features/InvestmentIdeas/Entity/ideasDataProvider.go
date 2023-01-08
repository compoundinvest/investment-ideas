package investmentideas

import (
	moexapi "compound/Core/Quotes/moexapi"
	yahooapi "compound/Core/Quotes/yahooapi"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

func getLocalIdeasList() InvestmentIdeas {
	ideasFileContent, err := os.Open("Features/InvestmentIdeas/Entity/investmentideas.json")
	if err != nil {
		fmt.Printf("Error while attempting to read the investment ideas list file: %v", err)
		return InvestmentIdeas{"Error", []InvestmentIdea{}}
	}

	byteValue, err := io.ReadAll(ideasFileContent)
	if err != nil {
		fmt.Println("Error while attempting to convert the ideas list file into a byte array: ", err)
		return InvestmentIdeas{"Error", []InvestmentIdea{}}
	}

	var ideas InvestmentIdeas
	err = json.Unmarshal(byteValue, &ideas)
	if err != nil {
		fmt.Println("Unable to unmarshal the JSON containing the list of investment ideas: ", err)
	}

	return ideas
}

func GetInvestmentIdeas() InvestmentIdeas {
	ideas := getLocalIdeasList()
	waitGroup := &sync.WaitGroup{}

	quotes := []moexapi.SimpleQuote{}

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

	ideas.CalculateUpsides(quotes)
	return ideas
}
