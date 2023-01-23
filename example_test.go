package currency_test

import (
	"fmt"

	"github.com/gromey/iso-currency"
)

func ExampleByAlphabeticCode() {
	ccy, err := currency.ByAlphabeticCode("JPY")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", ccy)
	// Output: {AlphabeticCode:JPY NumericCode:392 MinorUnits:0 Name:Yen CountryNames:[JAPAN]}
}

func ExampleByNumericCode() {
	ccy, err := currency.ByNumericCode("392")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", ccy)
	// Output: {AlphabeticCode:JPY NumericCode:392 MinorUnits:0 Name:Yen CountryNames:[JAPAN]}
}
