package main

import (
	quote "compound/Core/Quotes/Common"
	moexapi "compound/Core/Quotes/moexapi"
	yahooapi "compound/Core/Quotes/yahooapi"
	investmentideas "compound/Features/InvestmentIdeas/Entity"
	"encoding/json"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func main() {
	//No need to perform any operations in main() because Yandex Cloud Function will use YandexCFHandler()
}

// Entry point for the Yandex Cloud Function
func YandexCFHandler() ([]byte, error) {
	ideas := investmentideas.GetLocalIdeasList()
	quotes := []moexapi.SimpleQuote{}

	yahooQuotes := yahooapi.FetchQuotes(ideas.TickersOfSpecificCurrency("USD"))
	quotes = append(quotes, quote.ConvertToSimpleQuote(yahooQuotes)...)

	moexQuotes := moexapi.FetchQuotes(ideas.TickersOfSpecificCurrency("RUB"))
	quotes = append(quotes, quote.ConvertToSimpleQuote(moexQuotes)...)

	ideas.CalculateUpsides(quotes)

	response := Response{
		200,
		ideas,
	}
	responseJSON, err := json.Marshal(response)

	return responseJSON, err
}
