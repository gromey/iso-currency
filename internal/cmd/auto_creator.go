package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/gromey/iso-currency"
)

const (
	rawURL = "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml"

	constTemplate  = "internal/template/constants.tmpl"
	mapsTemplate   = "internal/template/maps.tmpl"
	readmeTemplate = "internal/template/README.tmpl"

	constFileName  = "constants.go"
	mapsFileName   = "maps.go"
	readmeFileName = "README.md"
)

type minorUnits struct {
	Value      uint
	Applicable bool
}

type currencyRow struct {
	AlphabeticCode string     `xml:"Ccy"`
	NumericCode    string     `xml:"CcyNbr"`
	MinorUnits     minorUnits `xml:"CcyMnrUnts"`
	Name           string     `xml:"CcyNm"`
	CountryName    string     `xml:"CtryNm"`
}

type currencyTable struct {
	CurrencyRows []currencyRow `xml:"CcyNtry"`
}

type iso4217 struct {
	Published     string        `xml:"Pblshd,attr"`
	CurrencyTable currencyTable `xml:"CcyTbl"`
}

type numeric struct {
	NumericCode    string
	AlphabeticCode string
}

type data struct {
	Alphabetic []currencyRow
	Numeric    []numeric
	Published  string
}

type isoResult struct {
	alphabeticCode string
	numericCode    string
	minorUnits     minorUnits
	name           string
	countryNames   []string
}

func (m *minorUnits) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	if strings.TrimSpace(v) != currency.NotApplicable {
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		m.Value = uint(n)
		m.Applicable = true
	} else {
		m.Value = 0
		m.Applicable = false
	}
	return nil
}

func run() error {
	iso := new(iso4217)

	if err := iso.get(); err != nil {
		return err
	}

	return iso.makeFiles()
}

func (iso *iso4217) get() error {
	response, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	defer func() { _ = response.Body.Close() }()

	var body []byte
	if body, err = io.ReadAll(response.Body); err != nil {
		return err
	}

	return xml.Unmarshal(body, iso)
}

func (iso *iso4217) makeFiles() error {
	mapISO := make(map[string]isoResult, len(iso.CurrencyTable.CurrencyRows))

	for _, ccyRow := range iso.CurrencyTable.CurrencyRows {
		if strings.TrimSpace(ccyRow.AlphabeticCode) == "" {
			continue
		}

		if v, ok := mapISO[strings.TrimSpace(ccyRow.AlphabeticCode)]; ok {
			v.countryNames = append(v.countryNames, strings.TrimSpace(ccyRow.CountryName))
			mapISO[strings.TrimSpace(ccyRow.AlphabeticCode)] = v
			continue
		}

		isoResultRow := isoResult{
			alphabeticCode: strings.TrimSpace(ccyRow.AlphabeticCode),
			numericCode:    ccyRow.NumericCode,
			minorUnits:     ccyRow.MinorUnits,
			name:           strings.TrimSpace(ccyRow.Name),
			countryNames:   []string{strings.TrimSpace(ccyRow.CountryName)},
		}

		mapISO[ccyRow.AlphabeticCode] = isoResultRow
	}

	alphabeticSlice := make([]currencyRow, len(mapISO))
	numericSlice := make([]numeric, len(mapISO))

	i := 0
	for _, v := range mapISO {
		alphabeticSlice[i] = currencyRow{
			AlphabeticCode: v.alphabeticCode,
			NumericCode:    v.numericCode,
			MinorUnits:     v.minorUnits,
			Name:           v.name,
			CountryName:    fmt.Sprintf("%#v", v.countryNames),
		}
		numericSlice[i] = numeric{
			AlphabeticCode: v.alphabeticCode,
			NumericCode:    v.numericCode,
		}
		i++
	}

	sort.Slice(alphabeticSlice, func(i, j int) bool {
		return alphabeticSlice[i].AlphabeticCode < alphabeticSlice[j].AlphabeticCode
	})

	sort.Slice(numericSlice, func(i, j int) bool {
		return numericSlice[i].NumericCode < numericSlice[j].NumericCode
	})

	result := data{
		Alphabetic: alphabeticSlice,
		Numeric:    numericSlice,
		Published:  iso.Published,
	}

	if err := createFileByTemplate(constTemplate, constFileName, result); err != nil {
		return err
	}

	if err := createFileByTemplate(mapsTemplate, mapsFileName, result); err != nil {
		return err
	}

	return createFileByTemplate(readmeTemplate, readmeFileName, result)
}

func createFileByTemplate(tempPath, filename string, data interface{}) error {
	temp, err := template.ParseFiles(tempPath)
	if err != nil {
		return err
	}

	var file *os.File
	if file, err = os.Create(filename); err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	return temp.Execute(file, data)
}
