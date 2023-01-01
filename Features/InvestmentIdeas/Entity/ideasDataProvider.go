package investmentideas

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetLocalIdeasList() InvestmentIdeas {
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
	json.Unmarshal(byteValue, &ideas)

	return ideas
}
