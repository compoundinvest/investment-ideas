package main

import (
	investmentideas "compound/Features/InvestmentIdeas/Entity"
	"encoding/json"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func main() {
	//No need to perform any operations in main() because Yandex Cloud Function will use YandexCFHandler()
	YandexCFHandler()
}

// Entry point for the Yandex Cloud Function
func YandexCFHandler() ([]byte, error) {
	ideas := investmentideas.GetInvestmentIdeas()

	response := Response{
		200,
		ideas,
	}
	responseJSON, err := json.Marshal(response)

	return responseJSON, err
}
