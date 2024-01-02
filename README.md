# iso-currency

![https://img.shields.io/github/v/tag/gromey/iso-currency](https://img.shields.io/github/v/tag/gromey/iso-currency)
![https://img.shields.io/github/license/gromey/iso-currency](https://img.shields.io/github/license/gromey/iso-currency)

`iso-currency` library of currencies based on the [ISO 4217 standard](https://www.iso.org/iso-4217-currency-codes.html)
published: ***2024-01-01***

## Installation

`iso-currency` can be installed like any other Go library through `go get`:

```console
go get github.com/gromey/iso-currency
```

Or, if you are already using
[Go Modules](https://github.com/golang/go/wiki/Modules), you may specify a version number as well:

```console
go get github.com/gromey/iso-currency@latest
```

## Getting Started

```go
package main

import (
	"fmt"

	"github.com/gromey/iso-currency"
)

func main() {
	jpy := currency.JPY.Get()

	fmt.Printf("jpy: %+v\n", jpy)
	// jpy: {AlphabeticCode:JPY NumericCode:392 MinorUnits:0 Name:Yen CountryNames:[JAPAN]}
	fmt.Printf("jpy.MinorUnits: Value() = %v, String() = %v\n", jpy.MinorUnits.Value(), jpy.MinorUnits.String())
	// jpy.MinorUnits: Value() = 0; String() = 0

	ccy, err := currency.ByAlphabeticCode("XAU")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ccy: %+v\n", ccy)
	// ccy: {AlphabeticCode:XAU NumericCode:959 MinorUnits:N.A. Name:Gold CountryNames:[ZZ08_Gold]}
	fmt.Printf("ccy.MinorUnits: Value() = %v, String() = %v\n", ccy.MinorUnits.Value(), ccy.MinorUnits.String())
	// ccy.MinorUnits: Value() = 0; String() = N.A.

	ccy, err = currency.ByNumericCode("008")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ccy: %+v\n", ccy)
	// ccy: {AlphabeticCode:ALL NumericCode:008 MinorUnits:2 Name:Lek CountryNames:[ALBANIA]}
	fmt.Printf("ccy.MinorUnits: Value() = %v, String() = %v\n", ccy.MinorUnits.Value(), ccy.MinorUnits.String())
	// ccy.MinorUnits: Value() = 2; String() = 2
}
```

If you need, you can create your own currency like this:

```go
	ccy := currency.Currency{
		AlphabeticCode: "YAC",
		NumericCode:    "123",
		MinorUnits:     currency.NewMinorUnits(2),
		Name:           "YouCurrencyName",
		CountryNames:   []string{"YouCountryName"},
	}

	fmt.Printf("ccy: %+v\n", ccy)
	// ccy: {AlphabeticCode:YAC NumericCode:123 MinorUnits:2 Name:YouCurrencyName CountryNames:[YouCountryName]}

	// If you do not specify MinorUnits, then when you create the currency, a default value will be assigned,
	// which will give "N.A." when displayed, like "XAU".

	ccy = currency.Currency{
		AlphabeticCode: "YAC",
		NumericCode:    "123",
		Name:           "YouCurrencyName",
		CountryNames:   []string{"YouCountryName"},
	}

	fmt.Printf("ccy: %+v\n", ccy)
	// ccy: {AlphabeticCode:YAC NumericCode:123 MinorUnits:N.A. Name:YouCurrencyName CountryNames:[YouCountryName]}
```

### Internal

To update the library:

```console
go generate
```
