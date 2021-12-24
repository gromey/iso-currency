package currency

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	expectALL = &currency{
		AlphabeticCode: "ALL",
		NumericCode:    numericCode(8),
		MinorUnits:     minorUnits{value: 2, applicable: true},
		Name:           "Lek",
		CountryNames:   []string{"ALBANIA"},
	}
	expectBHD = &currency{
		AlphabeticCode: "BHD",
		NumericCode:    numericCode(48),
		MinorUnits:     minorUnits{value: 3, applicable: true},
		Name:           "Bahraini Dinar",
		CountryNames:   []string{"BAHRAIN"},
	}
	expectCLF = &currency{
		AlphabeticCode: "CLF",
		NumericCode:    numericCode(990),
		MinorUnits:     minorUnits{value: 4, applicable: true},
		Name:           "Unidad de Fomento",
		CountryNames:   []string{"CHILE"},
	}
	expectVND = &currency{
		AlphabeticCode: "VND",
		NumericCode:    numericCode(704),
		MinorUnits:     minorUnits{value: 0, applicable: true},
		Name:           "Dong",
		CountryNames:   []string{"VIET NAM"},
	}
	expectXAU = &currency{
		AlphabeticCode: "XAU",
		NumericCode:    numericCode(959),
		MinorUnits:     minorUnits{value: 0, applicable: false},
		Name:           "Gold",
		CountryNames:   []string{"ZZ08_Gold"},
	}
)

func equal(t *testing.T, exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("Not equal:\nexp: %v\ngot: %v", exp, got)
	}
}

func TestByAlphabeticCode(t *testing.T) {
	var tests = []struct {
		AlphabeticCode string
		expectISO      *currency
		err            error
	}{
		{
			AlphabeticCode: "ALL",
			expectISO:      expectALL,
		},
		{
			AlphabeticCode: "bhd",
			expectISO:      expectBHD,
		},
		{
			AlphabeticCode: "Clf",
			expectISO:      expectCLF,
		},
		{
			AlphabeticCode: "vNd",
			expectISO:      expectVND,
		},
		{
			AlphabeticCode: "xaU",
			expectISO:      expectXAU,
		},
		{
			AlphabeticCode: "ERU",
			expectISO:      nil,
			err:            fmt.Errorf("%s: %s", ErrAlphaCode, "ERU"),
		},
	}
	for _, tt := range tests {
		iso, err := ByAlphabeticCode(tt.AlphabeticCode)
		equal(t, tt.err, err)
		equal(t, iso, tt.expectISO)
	}
}

func TestByNumericCode(t *testing.T) {
	var tests = []struct {
		numericCode uint16
		expectISO   *currency
		err         error
	}{
		{
			numericCode: 8,
			expectISO:   expectALL,
		},
		{
			numericCode: 48,
			expectISO:   expectBHD,
		},
		{
			numericCode: 990,
			expectISO:   expectCLF,
		},
		{
			numericCode: 704,
			expectISO:   expectVND,
		},
		{
			numericCode: 959,
			expectISO:   expectXAU,
		},
		{
			numericCode: 111,
			expectISO:   nil,
			err:         fmt.Errorf("%s: %s", ErrNumCode, "111"),
		},
	}
	for _, tt := range tests {
		iso, err := ByNumericCode(tt.numericCode)
		equal(t, tt.err, err)
		equal(t, iso, tt.expectISO)
	}
}
