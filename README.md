# iso-currency

![https://img.shields.io/github/v/tag/gromey/iso-currency](https://img.shields.io/github/v/tag/gromey/iso-currency)
![https://img.shields.io/github/license/gromey/iso-currency](https://img.shields.io/github/license/gromey/iso-currency)

`iso-currency` library of currencies based on the [ISO 4217 standard](https://www.iso.org/iso-4217-currency-codes.html)
published: ***October 1, 2021***

## Installation

`iso-currency` can be installed like any other Go library through `go get`:

```console
$ go get github.com/gromey/iso-currency
```

Or, if you are already using
[Go Modules](https://github.com/golang/go/wiki/Modules), you may specify a version number as well:

```console
$ go get github.com/gromey/iso-currency@latest
```

## Getting Started

```go
package main

import (
	"fmt"
	"github.com/gromey/iso-currency"
)

func main() {
	ccy, err := currency.ByAlphabeticCode("All")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ccy: %v;\nccy.NumericCode.Value: %d;\nccy.NumericCode.String(): %s\n",
		ccy, ccy.NumericCode.Value, ccy.NumericCode.String())
	// ccy: &{ALL 008 2 Lek [ALBANIA]};
	// ccy.NumericCode.Value: 8;
	// ccy.NumericCode.String(): 008

	ccy, err = currency.ByNumericCode(959)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ccy: %v;\nccy.MinorUnits.Value: %d;\nccy.MinorUnits.String(): %s\n",
		ccy, ccy.MinorUnits.Value, ccy.MinorUnits.String())
	// ccy: &{XAU 959 N.A. Gold [ZZ08_Gold]};
	// ccy.MinorUnits.Value: 0;
	// ccy.MinorUnits.String(): N.A.
}
```

## Updating

If the information on currencies has been updated, you can update the library files yourself by executing:

```console
$ go generate
```