package currency

import (
	"fmt"
	"strconv"
	"strings"
)

//go:generate sh ./internal/scripts/generate.sh

const NotApplicable string = "N.A."

type numericCode struct {
	Value uint
}

func (n numericCode) String() string {
	s := strconv.Itoa(int(n.Value))
	if len(s) != 3 {
		s = strings.Repeat("0", 3-len(s)) + s
	}
	return s
}

type minorUnits struct {
	Value      uint
	applicable bool
}

func (m minorUnits) String() string {
	if !m.applicable {
		return NotApplicable
	}
	return strconv.Itoa(int(m.Value))
}

// iso represents one currency in ISO 4217 format
type iso struct {
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

// ByAlphabeticCode returns iso currency by alphabetic code or 'unknown' for invalid code.
func ByAlphabeticCode(code string) (*iso, error) {
	code = strings.ToUpper(code)
	iso, ok := currencyByAlphabeticCode[code]
	if ok {
		return &iso, nil
	}
	return nil, fmt.Errorf("%s: %s", ErrAlphaCode, code)
}

// ByNumericCode returns iso currency by numeric code or 'unknown' for invalid code.
func ByNumericCode(code uint) (*iso, error) {
	alphabeticCode, ok := alphabeticCodeByNumericCode[code]
	if ok {
		return ByAlphabeticCode(alphabeticCode)
	}
	return nil, fmt.Errorf("%s: %d", ErrNumCode, code)
}
