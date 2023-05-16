package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TaxBracket struct {
	Min  int     `json:"min"`
	Max  int     `json:"max"`
	Rate float64 `json:"rate"`
}

type TaxBracketsByYear struct {
	TaxBrackets []TaxBracket `json:"tax_brackets"`
}

var AllTaxBrackets []TaxBracketsByYear

func init() {
	fmt.Println("Application initializing... querying tax brackets data")

	// Years supported by Mock API
	years := []int{2019, 2020, 2021, 2022}

	for _, year := range years {
		url := fmt.Sprintf("http://localhost:5000/tax-calculator/tax-year/%d", year)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("There was an error fetching from", url, "err: ", err);
		}

		fmt.Println(resp.Body)
		decoder := json.NewDecoder(resp.Body)

		var t TaxBracketsByYear
		err = decoder.Decode(&t)
		if err != nil {
			fmt.Println("Error while reading response body", err)
		}
		fmt.Println(t.TaxBrackets)

		// try 3 times in case error occurs
		//for i := 0; i < 3; i++ { // try 3 times
		//	body, err := http.Get(url)
		//	if err == nil {
		//		defer body.Body.Close()
		//		var taxBrackets []TaxBracket
		//		err = json.NewDecoder(body.Body).Decode(&taxBrackets)
		//		if err == nil {
		//			AllTaxBrackets = append(AllTaxBrackets, TaxBracketsByYear{Year: year, TaxBrackets: taxBrackets})
		//			fmt.Println("Successfully added tax bracket for year: ", year)
		//			break // successful response, exit loop
		//		}
		//	}
		//	time.Sleep(2 * time.Second) // wait before retrying
		//}
	}
}
