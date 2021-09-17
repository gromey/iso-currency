package currency

import (
	"fmt"
	"strings"
)

//go:generate go run ./cmd

// ISO represents one currency in ISO 4217 format
type ISO struct {
	// The alphabetic code is based on another ISO standard, ISO 3166, which lists the codes for country names.
	// The first two letters of the ISO 4217 three-letter code are the same as the code for the country name.
	// Where possible, the third letter corresponds to the first letter of the currency name.
	AlphabeticCode string
	// The three-digit numeric code is useful when currency codes need to be understood in countries
	// that do not use Latin scripts and for computerized systems.
	// Where possible, the three-digit numeric code is the same as the numeric country code.
	NumericCode string
	// The number of digits after the decimal separator.
	MinorUnits string
	// Currency name.
	Name string
	// Locations listed for this currency.
	CountryNames []string
}

// ByAlphabeticCode returns ISO currency by alphabetic code or 'unknown' for invalid code.
func ByAlphabeticCode(code string) (*ISO, error) {
	code = strings.ToUpper(code)
	iso, ok := currencyByAlphabeticCode[code]
	if ok {
		return &iso, nil
	}
	return nil, fmt.Errorf("%s: %s", ErrAlphaCode, code)
}

// ByNumericCode returns ISO currency by numeric code or 'unknown' for invalid code.
func ByNumericCode(code string) (*ISO, error) {
	alphabeticCode, ok := alphabeticCodeByNumericCode[code]
	if ok {
		return ByAlphabeticCode(alphabeticCode)
	}
	return nil, fmt.Errorf("%s: %s", ErrNumCode, code)
}
