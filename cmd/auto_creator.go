package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const (
	url = "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list_one.xml"

	mapsTemplate   = "template/maps.tmpl"
	readmeTemplate = "template/README.tmpl"

	mapsFileName   = "maps.go"
	readmeFileName = "README.md"
)

type currencyRow struct {
	AlphabeticCode string `xml:"Ccy"`
	NumericCode    uint   `xml:"CcyNbr"`
	MinorUnitsS    string `xml:"CcyMnrUnts"`
	MinorUnitsI    uint   `xml:"-"`
	Valid          bool   `xml:"-"`
	Name           string `xml:"CcyNm"`
	CountryName    string `xml:"CtryNm"`
}

type currencyTable struct {
	CurrencyRows []currencyRow `xml:"CcyNtry"`
}

type iso4217 struct {
	Published     string        `xml:"Pblshd,attr"`
	CurrencyTable currencyTable `xml:"CcyTbl"`
}

type numeric struct {
	NumericCode    uint
	AlphabeticCode string
}

type data struct {
	Alphabetic []currencyRow
	Numeric    []numeric
	Published  string
}

type isoResult struct {
	alphabeticCode string
	numericCode    uint
	minorUnits     uint
	valid          bool
	name           string
	countryNames   []string
}

func run() error {
	iso := new(iso4217)

	if err := iso.get(); err != nil {
		return err
	}

	return iso.makeFiles()
}

func (iso *iso4217) get() error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
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

		v, ok := mapISO[strings.TrimSpace(strings.TrimSpace(ccyRow.AlphabeticCode))]
		if ok {
			v.countryNames = append(v.countryNames, strings.TrimSpace(ccyRow.CountryName))
			mapISO[strings.TrimSpace(ccyRow.AlphabeticCode)] = v
			continue
		}

		isoResultRow := isoResult{
			alphabeticCode: strings.TrimSpace(ccyRow.AlphabeticCode),
			numericCode:    ccyRow.NumericCode,
			name:           strings.TrimSpace(ccyRow.Name),
			countryNames:   []string{strings.TrimSpace(ccyRow.CountryName)},
		}

		if strings.TrimSpace(ccyRow.MinorUnitsS) == "N.A." {
			isoResultRow.minorUnits = 0
			isoResultRow.valid = false
		} else {
			m, err := strconv.Atoi(ccyRow.MinorUnitsS)
			if err != nil {
				return err
			}
			isoResultRow.minorUnits = uint(m)
			isoResultRow.valid = true
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
			MinorUnitsI:    v.minorUnits,
			Valid:          v.valid,
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

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return temp.Execute(file, data)
}
