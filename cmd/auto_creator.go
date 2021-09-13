package main

import (
	"encoding/xml"
	"fmt"
	currency "github.com/gromey/iso-currency"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
)

const (
	url = "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/amendments/lists/list_one.xml"

	mapsTemplate   = "template/maps.tmpl"
	readmeTemplate = "template/README.tmpl"

	mapsFileName   = "maps.go"
	readmeFileName = "README.md"
)

type currencyRow struct {
	AlphabeticCode string `xml:"Ccy"`
	NumericCode    string `xml:"CcyNbr"`
	MinorUnits     string `xml:"CcyMnrUnts"`
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
	NumericCode    string
	AlphabeticCode string
}

type data struct {
	Alphabetic []currencyRow
	Numeric    []numeric
	Published  string
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
	mapISO := make(map[string]currency.ISO, len(iso.CurrencyTable.CurrencyRows))

	for _, ccyRow := range iso.CurrencyTable.CurrencyRows {
		if strings.TrimSpace(ccyRow.AlphabeticCode) == "" {
			continue
		}

		v, ok := mapISO[strings.TrimSpace(strings.TrimSpace(ccyRow.AlphabeticCode))]
		if ok {
			v.CountryNames = append(v.CountryNames, strings.TrimSpace(ccyRow.CountryName))
			mapISO[strings.TrimSpace(ccyRow.AlphabeticCode)] = v
			continue
		}

		mapISO[ccyRow.AlphabeticCode] = currency.ISO{
			AlphabeticCode: strings.TrimSpace(ccyRow.AlphabeticCode),
			NumericCode:    strings.TrimSpace(ccyRow.NumericCode),
			MinorUnits:     strings.TrimSpace(ccyRow.MinorUnits),
			Name:           strings.TrimSpace(ccyRow.Name),
			CountryNames:   []string{strings.TrimSpace(ccyRow.CountryName)},
		}
	}

	alphabeticSlice := make([]currencyRow, len(mapISO))
	numericSlice := make([]numeric, len(mapISO))

	i := 0
	for _, v := range mapISO {
		alphabeticSlice[i] = currencyRow{
			AlphabeticCode: v.AlphabeticCode,
			NumericCode:    v.NumericCode,
			MinorUnits:     v.MinorUnits,
			Name:           v.Name,
			CountryName:    fmt.Sprintf("%#v", v.CountryNames),
		}
		numericSlice[i] = numeric{
			AlphabeticCode: v.AlphabeticCode,
			NumericCode:    v.NumericCode,
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