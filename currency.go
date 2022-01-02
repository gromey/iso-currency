package currency

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:generate sh ./internal/scripts/generate.sh

// Currency in ISO 4217 format.
type Currency struct {
	// The alphabetic code is based on another ISO standard, ISO 3166, which lists the codes for country names.
	// The first two letters of the ISO 4217 three-letter code are the same as the code for the country name.
	// Where possible, the third letter corresponds to the first letter of the currency name.
	AlphabeticCode string
	// The three-digit numeric code is useful when currency codes need to be understood in countries
	// that do not use Latin scripts and for computerized systems.
	// Where possible, the three-digit numeric code is the same as the numeric country code.
	NumericCode string
	// The number of digits after the decimal separator.
	MinorUnits MinorUnits
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

// Get returns Currency for the alphabetic currency code.
func (a alphabeticCode) Get() *Currency {
	ccy := currencyByAlphabeticCode[a]
	return &ccy
}

// MinorUnits represents the number of digits after the decimal separator.
// Use the NewMinorUnits func to create a MinorUnits.
type MinorUnits struct {
	value      uint8
	applicable bool
}

// NewMinorUnits returns a new minor units of the currency.
func NewMinorUnits(value uint8) MinorUnits {
	return MinorUnits{
		value:      value,
		applicable: true,
	}
}

// String returns the minor units of the currency in a string.
func (m MinorUnits) String() string {
	if !m.applicable {
		return NotApplicable
	}
	return strconv.Itoa(int(m.value))
}

// Value returns value of the minor units of the currency.
func (m MinorUnits) Value() uint8 {
	return m.value
}

// ByAlphabeticCode returns Currency by alphabetic code
// or error: 'unknown' for invalid code.
func ByAlphabeticCode(code string) (*Currency, error) {
	code = strings.ToUpper(code)
	if len(code) != 3 || !onlyLetters(code) {
		return nil, fmt.Errorf("%s: %s", InvalidAlphabeticCode, code)
	}
	ccy, ok := currencyByAlphabeticCode[alphabeticCode(code)]
	if ok {
		return &ccy, nil
	}
	return nil, fmt.Errorf("%s: %s", UnknownAlphabeticCode, code)
}

// ByNumericCode returns Currency by numeric code
// or error: 'unknown' for invalid code.
func ByNumericCode(code string) (*Currency, error) {
	if len(code) != 3 || !onlyNumbers(code) {
		return nil, fmt.Errorf("%s: %s", InvalidNumericCode, code)
	}
	alphaCode, ok := alphabeticCodeByNumericCode[code]
	if ok {
		return alphaCode.Get(), nil
	}
	return nil, fmt.Errorf("%s: %s", UnknownNumericCode, code)
}

func onlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func onlyNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}
