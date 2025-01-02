package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ivivanov18/go-currency-converter/api"
	"github.com/ivivanov18/go-currency-converter/ui"
)

func main() {
	ui.Execute()
	fmt.Printf("Base currency: %s\n", ui.BaseCurrencyInput)
	fmt.Printf("Target currency: %s\n", ui.TargetCurrencyInput)
	fmt.Printf("Amount: %s\n", ui.AmountInput)

	floatInput, err := strconv.ParseFloat(ui.AmountInput, 64)
	if err != nil {
		log.Fatalf("Failed to convert input to float: %v", err)
		return
	}
	converted := api.Convert(ui.BaseCurrencyInput, ui.TargetCurrencyInput, floatInput)
	fmt.Printf("Converted amount: %f\n", converted)
}
