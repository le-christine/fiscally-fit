package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


type TaxBrackets struct {
	TaxBrackets []struct {
		Min  float64     `json:"min"`
		Max  *float64     `json:"max,omitempty"`
		Rate float64 `json:"rate"`
	} `json:"tax_brackets"`
}

var AllTaxBrackets map[int]TaxBrackets

func init() {
	AllTaxBrackets = make(map[int]TaxBrackets)

	fmt.Println("Application initializing... querying tax brackets data")

	// Years supported by Mock API
	years := []int{2019, 2020, 2021, 2022}

	for _, year := range years {
		url := fmt.Sprintf("http://localhost:5000/tax-calculator/tax-year/%d", year)

		//try 3 times in case error occurs
		for i := 0; i < 3; i++ {

			resp, err := http.Get(url)

			if err != nil {
				fmt.Println("There was an error fetching from", url, "err: ", err);
			}

			if resp.StatusCode == 200 {

				var taxBracket TaxBrackets
				err = json.NewDecoder(resp.Body).Decode(&taxBracket)
				if err != nil {
					fmt.Println("Error while reading response body", err)
				} else {
					AllTaxBrackets[year] = taxBracket
					fmt.Println("Successfully added tax bracket for year: ", year)
					break // successful response, exit loop
				}
			} else {
				fmt.Println("There was an error fetching from", url, "resp code: ", resp.StatusCode)
			}
			fmt.Println("Retrying ", url)
			time.Sleep(2 * time.Second) // wait before retrying
		}
	}
}
