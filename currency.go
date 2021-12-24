package currency

import (
	"fmt"
	"strconv"
	"strings"
)

//go:generate sh ./internal/scripts/generate.sh

// currency in ISO 4217 format.
type currency struct {
	// The alphabetic code is based on another ISO standard, ISO 3166, which lists the codes for country names.
	// The first two letters of the ISO 4217 three-letter code are the same as the code for the country name.
	// Where possible, the third letter corresponds to the first letter of the currency name.
	AlphabeticCode string
	// The three-digit numeric code is useful when currency codes need to be understood in countries
	// that do not use Latin scripts and for computerized systems.
	// Where possible, the three-digit numeric code is the same as the numeric country code.
	NumericCode numericCode
	// The number of digits after the decimal separator.
	MinorUnits minorUnits
	// Currency name.
	Name string
	// Locations listed for this currency.
	CountryNames []string
}

type alphabeticCode string

// String returns the alphabetic currency code in a string.
func (a alphabeticCode) String() string {
	return string(a)
}

// Get returns iso currency for the alphabetic currency code.
func (a alphabeticCode) Get() *currency {
	ccy := currencyByAlphabeticCode[a]
	return &ccy
}

type numericCode uint16

// String returns the numeric currency code in a string.
func (n numericCode) String() string {
	return fmt.Sprintf("%03d", n)
}

// Value returns value of the numeric currency code.
func (n numericCode) Value() uint16 {
	return uint16(n)
}

type minorUnits struct {
	value      uint8
	applicable bool
}

// String returns the minor units of the currency in a string.
func (m minorUnits) String() string {
	if !m.applicable {
		return NotApplicable
	}
	return strconv.Itoa(int(m.value))
}

// Value returns value of the minor units of the currency.
func (m minorUnits) Value() uint8 {
	return m.value
}

// ByAlphabeticCode returns iso currency by alphabetic code
// or error: 'unknown' for invalid code.
func ByAlphabeticCode(code string) (*currency, error) {
	ccy, ok := currencyByAlphabeticCode[alphabeticCode(strings.ToUpper(code))]
	if ok {
		return &ccy, nil
	}
	return nil, fmt.Errorf("%s: %s", ErrAlphaCode, code)
}

// ByNumericCode returns iso currency by numeric code
// or error: 'unknown' for invalid code.
func ByNumericCode(code uint16) (*currency, error) {
	alphaCode, ok := alphabeticCodeByNumericCode[code]
	if ok {
		return alphaCode.Get(), nil
	}
	return nil, fmt.Errorf("%s: %d", ErrNumCode, code)
}
