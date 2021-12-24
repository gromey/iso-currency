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

	fmt.Println(ccy)
	// Output: &{JPY 392 0 Yen [JAPAN]}
}

func ExampleByNumericCode() {
	ccy, err := currency.ByNumericCode(392)
	if err != nil {
		panic(err)
	}

	fmt.Println(ccy)
	// Output: &{JPY 392 0 Yen [JAPAN]}
}
