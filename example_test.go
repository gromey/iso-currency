package currency_test

import (
	"fmt"
	"github.com/gromey/iso-currency"
)

func ExampleByAlphabeticCode() {
	ccy, err := currency.ByAlphabeticCode("JPY")
	fmt.Println(ccy, err)
	// Output: &{JPY 392 0 Yen [JAPAN]} <nil>
}

func ExampleByNumericCode() {
	ccy, err := currency.ByNumericCode("392")
	fmt.Println(ccy, err)
	// Output: &{JPY 392 0 Yen [JAPAN]} <nil>
}
