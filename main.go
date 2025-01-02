package main

import (
	"fmt"

	"github.com/ivivanov18/go-currency-converter/ui"
)

func main() {
	ui.Execute()
	fmt.Printf("Base currency: %s\n", ui.BaseCurrencyInput)
	fmt.Printf("Target currency: %s\n", ui.TargetCurrencyInput)
	fmt.Printf("Amount: %s\n", ui.AmountInput)
}
