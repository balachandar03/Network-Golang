package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CurrencyResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	// API endpoint for currency conversion
	url := "https://api.exchangerate-api.com/v4/latest/USD" // USD is the base currency, you can change it to your preferred base currency

	// Make an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse the JSON response
	var currencyResp CurrencyResponse
	err = json.Unmarshal(body, &currencyResp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the conversion rates
	fmt.Println("Conversion Rates:")
	for currency, rate := range currencyResp.Rates {
		fmt.Printf("%s: %f\n", currency, rate)
	}
}

