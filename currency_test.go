package currency_test

import (
	"fmt"
	"github.com/gromey/iso-currency"
	"reflect"
	"testing"
)

var (
	expectEUR = &currency.ISO{
		AlphabeticCode: "EUR",
		NumericCode:    "978",
		MinorUnits:     "2",
		Name:           "Euro",
		CountryNames:   []string{"ÅLAND ISLANDS", "ANDORRA", "AUSTRIA", "BELGIUM", "CYPRUS", "ESTONIA", "EUROPEAN UNION", "FINLAND", "FRANCE", "FRENCH GUIANA", "FRENCH SOUTHERN TERRITORIES (THE)", "GERMANY", "GREECE", "GUADELOUPE", "HOLY SEE (THE)", "IRELAND", "ITALY", "LATVIA", "LITHUANIA", "LUXEMBOURG", "MALTA", "MARTINIQUE", "MAYOTTE", "MONACO", "MONTENEGRO", "NETHERLANDS (THE)", "PORTUGAL", "RÉUNION", "SAINT BARTHÉLEMY", "SAINT MARTIN (FRENCH PART)", "SAINT PIERRE AND MIQUELON", "SAN MARINO", "SLOVAKIA", "SLOVENIA", "SPAIN"},
	}
	expectBOV = &currency.ISO{
		AlphabeticCode: "BOV",
		NumericCode:    "984",
		MinorUnits:     "2",
		Name:           "Mvdol",
		CountryNames:   []string{"BOLIVIA (PLURINATIONAL STATE OF)"},
	}
	expectUSD = &currency.ISO{
		AlphabeticCode: "USD",
		NumericCode:    "840",
		MinorUnits:     "2",
		Name:           "US Dollar",
		CountryNames:   []string{"AMERICAN SAMOA", "BONAIRE, SINT EUSTATIUS AND SABA", "BRITISH INDIAN OCEAN TERRITORY (THE)", "ECUADOR", "EL SALVADOR", "GUAM", "HAITI", "MARSHALL ISLANDS (THE)", "MICRONESIA (FEDERATED STATES OF)", "NORTHERN MARIANA ISLANDS (THE)", "PALAU", "PANAMA", "PUERTO RICO", "TIMOR-LESTE", "TURKS AND CAICOS ISLANDS (THE)", "UNITED STATES MINOR OUTLYING ISLANDS (THE)", "UNITED STATES OF AMERICA (THE)", "VIRGIN ISLANDS (BRITISH)", "VIRGIN ISLANDS (U.S.)"},
	}
)

func equal(t *testing.T, exp, got interface{}) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("Not equal:\nexp: %v\ngot: %v", exp, got)
	}
}

func TestByAlphabeticCode(t *testing.T) {
	var tests = []struct {
		alphabeticCode string
		expectISO      *currency.ISO
		err            error
	}{
		{
			alphabeticCode: "EUR",
			expectISO:      expectEUR,
			err:            nil,
		},
		{
			alphabeticCode: "bov",
			expectISO:      expectBOV,
			err:            nil,
		},
		{
			alphabeticCode: "usD",
			expectISO:      expectUSD,
			err:            nil,
		},
		{
			alphabeticCode: "ERU",
			expectISO:      nil,
			err:            fmt.Errorf("%s: %s", currency.ErrAlphaCode, "ERU"),
		},
	}
	for _, tt := range tests {
		iso, err := currency.ByAlphabeticCode(tt.alphabeticCode)
		equal(t, tt.err, err)
		equal(t, iso, tt.expectISO)
	}
}

func TestByNumericCode(t *testing.T) {
	var tests = []struct {
		numericCode string
		expectISO   *currency.ISO
		err         error
	}{
		{
			numericCode: "978",
			expectISO:   expectEUR,
			err:         nil,
		},
		{
			numericCode: "984",
			expectISO:   expectBOV,
			err:         nil,
		},
		{
			numericCode: "840",
			expectISO:   expectUSD,
			err:         nil,
		},
		{
			numericCode: "111",
			expectISO:   nil,
			err:         fmt.Errorf("%s: %s", currency.ErrNumCode, "111"),
		},
	}
	for _, tt := range tests {
		iso, err := currency.ByNumericCode(tt.numericCode)
		equal(t, tt.err, err)
		equal(t, iso, tt.expectISO)
	}
}
