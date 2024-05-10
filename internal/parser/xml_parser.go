// Package parser contains functionality for parsing XML files.
package parser

import (
	"encoding/xml"
	"io"
	"strconv"
	"strings"

	"github.com/sveboo/exchangerate/internal/models"
	"golang.org/x/net/html/charset"
)

// Valute represents a collection of currency items.
type Valute struct {
	XMLName xml.Name          `xml:"Valuta"`
	Items   []models.Currency `xml:"Item"`
}

// ValCurs represents a collection of `tempCurrency` values with XML tags for serialization.
type ValCurs struct {
	XMLName xml.Name       `xml:"ValCurs"`
	Val     []tempCurrency `xml:"Valute"`
}

// tempCurrency represents a temporary currency with ISO code and value fields for XML encoding.
type tempCurrency struct {
	ISOName string `xml:"CharCode"`
	Value   string `xml:"Value"`
}

// XMLParser provides functionality for parsing xml-files
type XMLParser struct {
}

// ParseCurrencyISO parse currency ISO from an XML file.
func (p XMLParser) ParseCurrencyISO(data io.ReadCloser) ([]models.Currency, error) {
	defer data.Close()
	var v Valute
	d := xml.NewDecoder(data)
	d.CharsetReader = charset.NewReaderLabel
	err := d.Decode(&v)
	return v.Items, err
}

// Parse parse currency exchange rates  from an XML file.
func (p XMLParser) Parse(data io.ReadCloser) (map[string]float32, error) {
	defer data.Close()
	var vc ValCurs
	d := xml.NewDecoder(data)
	d.CharsetReader = charset.NewReaderLabel
	err := d.Decode(&vc)

	if err != nil {
		return nil, err
	}

	rates := make(map[string]float32, len(vc.Val))
	for _, val := range vc.Val {
		rate, err := strconv.ParseFloat(strings.Replace(val.Value, ",", ".", -1), 32)
		if err != nil {
			return nil, err
		}
		rates[val.ISOName] = float32(rate)
	}

	return rates, nil
}
