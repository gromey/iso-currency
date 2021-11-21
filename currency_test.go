package currency

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	expectALL = &iso{
		AlphabeticCode: "ALL",
		NumericCode:    numericCode{Value: 8},
		MinorUnits:     minorUnits{Value: 2, applicable: true},
		Name:           "Lek",
		CountryNames:   []string{"ALBANIA"},
	}
	expectBHD = &iso{
		AlphabeticCode: "BHD",
		NumericCode:    numericCode{Value: 48},
		MinorUnits:     minorUnits{Value: 3, applicable: true},
		Name:           "Bahraini Dinar",
		CountryNames:   []string{"BAHRAIN"},
	}
	expectCLF = &iso{
		AlphabeticCode: "CLF",
		NumericCode:    numericCode{Value: 990},
		MinorUnits:     minorUnits{Value: 4, applicable: true},
		Name:           "Unidad de Fomento",
		CountryNames:   []string{"CHILE"},
	}
	expectVND = &iso{
		AlphabeticCode: "VND",
		NumericCode:    numericCode{Value: 704},
		MinorUnits:     minorUnits{Value: 0, applicable: true},
		Name:           "Dong",
		CountryNames:   []string{"VIET NAM"},
	}
	expectXAU = &iso{
		AlphabeticCode: "XAU",
		NumericCode:    numericCode{Value: 959},
		MinorUnits:     minorUnits{Value: 0, applicable: false},
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
		expectISO      *iso
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
		numericCode uint
		expectISO   *iso
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
