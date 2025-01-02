package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var APP_ID = ""

var BASE_URL = "https://openexchangerates.org/api/latest.json/"

func Convert(baseCurrency string, targetCurrency string, amount float64) float64 {
	url := BASE_URL + "?app_id=" + APP_ID
	fmt.Printf("Making request to: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	fmt.Printf("Response: %v\n", result)
	rates, ok := result["rates"].(map[string]interface{})
	if !ok {
		log.Fatalf("Unexpected response format")
	}

	targetRate, ok := rates[targetCurrency].(float64)
	if !ok {
		log.Fatalf("Error: target currency code %s not found", targetCurrency)
	}

	return targetRate * amount
}
