package currency

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	expectALL = Currency{
		AlphabeticCode: "ALL",
		NumericCode:    "008",
		MinorUnits:     NewMinorUnits(2),
		Name:           "Lek",
		CountryNames:   []string{"ALBANIA"},
	}
	expectBHD = Currency{
		AlphabeticCode: "BHD",
		NumericCode:    "048",
		MinorUnits:     NewMinorUnits(3),
		Name:           "Bahraini Dinar",
		CountryNames:   []string{"BAHRAIN"},
	}
	expectCLF = Currency{
		AlphabeticCode: "CLF",
		NumericCode:    "990",
		MinorUnits:     NewMinorUnits(4),
		Name:           "Unidad de Fomento",
		CountryNames:   []string{"CHILE"},
	}
	expectVND = Currency{
		AlphabeticCode: "VND",
		NumericCode:    "704",
		MinorUnits:     NewMinorUnits(0),
		Name:           "Dong",
		CountryNames:   []string{"VIET NAM"},
	}
	expectXAU = Currency{
		AlphabeticCode: "XAU",
		NumericCode:    "959",
		Name:           "Gold",
		CountryNames:   []string{"ZZ08_Gold"},
	}
)

func equal(t *testing.T, exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("Not equal:\nexp: %v\ngot: %v", exp, got)
	}
}

func TestAlphabeticCode_Get(t *testing.T) {
	var tests = []struct {
		alphabeticCode alphabeticCode
		expectCcy      Currency
	}{
		{
			alphabeticCode: ALL,
			expectCcy:      expectALL,
		},
		{
			alphabeticCode: BHD,
			expectCcy:      expectBHD,
		},
		{
			alphabeticCode: CLF,
			expectCcy:      expectCLF,
		},
		{
			alphabeticCode: VND,
			expectCcy:      expectVND,
		},
		{
			alphabeticCode: XAU,
			expectCcy:      expectXAU,
		},
	}
	for _, tt := range tests {
		iso := tt.alphabeticCode.Get()
		equal(t, iso, tt.expectCcy)
	}
}

func TestAlphabeticCode_String(t *testing.T) {
	var tests = []struct {
		alphabeticCode alphabeticCode
		expectString   string
	}{
		{
			alphabeticCode: ALL,
			expectString:   "ALL",
		},
		{
			alphabeticCode: BHD,
			expectString:   "BHD",
		},
		{
			alphabeticCode: CLF,
			expectString:   "CLF",
		},
		{
			alphabeticCode: VND,
			expectString:   "VND",
		},
		{
			alphabeticCode: XAU,
			expectString:   "XAU",
		},
	}
	for _, tt := range tests {
		s := tt.alphabeticCode.String()
		equal(t, s, tt.expectString)
	}
}

func TestNewMinorUnits(t *testing.T) {
	var tests = []struct {
		minorUnits   MinorUnits
		expectString string
		expectValue  uint8
	}{
		{
			minorUnits:   NewMinorUnits(2),
			expectString: "2",
			expectValue:  2,
		},
		{
			minorUnits:   NewMinorUnits(3),
			expectString: "3",
			expectValue:  3,
		},
		{
			minorUnits:   NewMinorUnits(4),
			expectString: "4",
			expectValue:  4,
		},
		{
			minorUnits:   NewMinorUnits(0),
			expectString: "0",
			expectValue:  0,
		},
		{
			minorUnits:   MinorUnits{},
			expectString: NotApplicable,

			expectValue: 0,
		},
	}
	for _, tt := range tests {
		equal(t, tt.minorUnits.String(), tt.expectString)
		equal(t, tt.minorUnits.Value(), tt.expectValue)

	}
}

func TestMinorUnits_String(t *testing.T) {
	var tests = []struct {
		alphabeticCode alphabeticCode
		expectString   string
	}{
		{
			alphabeticCode: ALL,
			expectString:   "2",
		},
		{
			alphabeticCode: BHD,
			expectString:   "3",
		},
		{
			alphabeticCode: CLF,
			expectString:   "4",
		},
		{
			alphabeticCode: VND,
			expectString:   "0",
		},
		{
			alphabeticCode: XAU,
			expectString:   NotApplicable,
		},
	}
	for _, tt := range tests {
		s := tt.alphabeticCode.Get().MinorUnits.String()
		equal(t, s, tt.expectString)
	}
}

func TestMinorUnits_Value(t *testing.T) {
	var tests = []struct {
		alphabeticCode alphabeticCode
		expectValue    uint8
	}{
		{
			alphabeticCode: ALL,
			expectValue:    2,
		},
		{
			alphabeticCode: BHD,
			expectValue:    3,
		},
		{
			alphabeticCode: CLF,
			expectValue:    4,
		},
		{
			alphabeticCode: VND,
			expectValue:    0,
		},
		{
			alphabeticCode: XAU,
			expectValue:    0,
		},
	}
	for _, tt := range tests {
		v := tt.alphabeticCode.Get().MinorUnits.Value()
		equal(t, v, tt.expectValue)
	}
}

func TestByAlphabeticCode(t *testing.T) {
	var tests = []struct {
		alphabeticCode string
		expectCcy      Currency
		err            error
	}{
		{
			alphabeticCode: "ALL",
			expectCcy:      expectALL,
		},
		{
			alphabeticCode: "bhd",
			expectCcy:      expectBHD,
		},
		{
			alphabeticCode: "Clf",
			expectCcy:      expectCLF,
		},
		{
			alphabeticCode: "vNd",
			expectCcy:      expectVND,
		},
		{
			alphabeticCode: "xaU",
			expectCcy:      expectXAU,
		},
		{
			alphabeticCode: "ER0",
			expectCcy:      Currency{},
			err:            fmt.Errorf("%s: %s", InvalidAlphabeticCode, "ER0"),
		},
		{
			alphabeticCode: "ERU",
			expectCcy:      Currency{},
			err:            fmt.Errorf("%s: %s", UnknownAlphabeticCode, "ERU"),
		},
	}
	for _, tt := range tests {
		iso, err := ByAlphabeticCode(tt.alphabeticCode)
		equal(t, tt.err, err)
		equal(t, iso, tt.expectCcy)
	}
}

func TestByNumericCode(t *testing.T) {
	var tests = []struct {
		numericCode string
		expectCcy   Currency
		err         error
	}{
		{
			numericCode: "008",
			expectCcy:   expectALL,
		},
		{
			numericCode: "048",
			expectCcy:   expectBHD,
		},
		{
			numericCode: "990",
			expectCcy:   expectCLF,
		},
		{
			numericCode: "704",
			expectCcy:   expectVND,
		},
		{
			numericCode: "959",
			expectCcy:   expectXAU,
		},
		{
			numericCode: "A51",
			expectCcy:   Currency{},
			err:         fmt.Errorf("%s: %s", InvalidNumericCode, "A51"),
		},
		{
			numericCode: "111",
			expectCcy:   Currency{},
			err:         fmt.Errorf("%s: %s", UnknownNumericCode, "111"),
		},
	}
	for _, tt := range tests {
		iso, err := ByNumericCode(tt.numericCode)
		equal(t, tt.err, err)
		equal(t, iso, tt.expectCcy)
	}
}

func TestOnlyLetters(t *testing.T) {
	var tests = []struct {
		alphabeticCode string
		expectBool     bool
	}{
		{
			alphabeticCode: "AAA",
			expectBool:     true,
		},
		{
			alphabeticCode: "AA1",
			expectBool:     false,
		},
		{
			alphabeticCode: "AA!",
			expectBool:     false,
		},
	}
	for _, tt := range tests {
		equal(t, onlyLetters(tt.alphabeticCode), tt.expectBool)
	}
}

func TestOnlyNumbers(t *testing.T) {
	var tests = []struct {
		numericCode string
		expectBool  bool
	}{
		{
			numericCode: "000",
			expectBool:  true,
		},
		{
			numericCode: "00A",
			expectBool:  false,
		},
		{
			numericCode: "00!",
			expectBool:  false,
		},
	}
	for _, tt := range tests {
		equal(t, onlyNumbers(tt.numericCode), tt.expectBool)
	}
}
