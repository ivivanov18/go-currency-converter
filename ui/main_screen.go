package ui

import (
	"log"

	"github.com/charmbracelet/huh"
)

var (
	BaseCurrencyInput   string
	TargetCurrencyInput string
	AmountInput         string
)

func Execute() {
	mainUI := createMainUI()

	err := mainUI.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func createMainUI() huh.Form {
	mainUI := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose the currency you want to convert from").
				Options(
					huh.NewOption("USD", "USD"),
				).
				Value(&BaseCurrencyInput),
			huh.NewSelect[string]().
				Title("Choose the currency you want to convert to").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("JPY", "JPY"),
					huh.NewOption("BGN", "BGN"),
				).
				Value(&TargetCurrencyInput),
			huh.NewInput().
				Title("Enter the amount you want to convert").
				Value(&AmountInput),
		),
	)

	return *mainUI
}
