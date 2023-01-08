package moexquote

import (
	quote "compound/Core/Quotes/Common"
	utils "compound/Core/Utilities"
	"encoding/json"
	"fmt"
	"net/http"
)

type SimpleQuote = quote.SimpleQuote

func FetchQuotes(tickers []string) []SimpleQuote {
	simpleQuotes := []quote.SimpleQuote{}

	quotes := fetchQuotes(tickers)
	simpleQuotes = append(simpleQuotes, quote.ConvertToSimpleQuote(quotes)...)

	return simpleQuotes
}

func fetchQuotes(tickers []string) []MoexQuote {
	quotesURL := "https://iss.moex.com/iss/engines/stock/markets/shares/boards/tqbr/securities.json?iss.meta=on&iss.only=marketdata&marketdata.columns=SECID,LAST,ISSUECAPITALIZATION"

	response, err := http.Get(quotesURL)
	if err != nil {
		fmt.Println("Unable to fetch quotes from ", quotesURL)
		return nil
	}
	defer response.Body.Close()

	var quotesDTO MoexQuotesDTO
	json.NewDecoder(response.Body).Decode(&quotesDTO)

	quotes := []MoexQuote{}
	for _, quoteDTO := range quotesDTO.MarketData.Data {
		var ticker string
		var quote float64
		var marketCap float64

		//The first element in the array is the quote's ticker (string)
		switch quoteDTO[0].(type) {
		case string:
			ticker = quoteDTO[0].(string)
			if !utils.Contains(tickers, ticker) {
				continue
			}

		default:
			fmt.Printf("MoexQuoteService: Unable to cast %v as a ticker (string)", quoteDTO[0])
			continue
		}

		//The second element in the array is the quote itself (number)
		switch quoteDTO[1].(type) {
		case float64, float32, int64, int32, int:
			quote = quoteDTO[1].(float64)
		default:
			continue
		}
		//The third element in the array is the stock's market cap (number)
		switch quoteDTO[2].(type) {
		case float64, float32, int64, int32, int:
			marketCap = quoteDTO[2].(float64)
		default:
		}

		quotes = append(quotes, MoexQuote{ticker, quote, marketCap})
	}

	return quotes
}
