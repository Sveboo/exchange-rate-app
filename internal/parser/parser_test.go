package parser

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sveboo/exchangerate/internal/models"
)

func TestParseCurrencyISO(t *testing.T) {

	parser := XMLParser{}

	tests := []struct {
		xmlData      string
		outputStruct []models.Currency
	}{
		{
			xmlData: `
			<Valuta>
				<Item>
					<ISO_Char_Code>USD</ISO_Char_Code>
					<Value>75.5</Value>
				</Item>
				<Item>
					<ISO_Char_Code>EUR</ISO_Char_Code>
					<Value>88.2</Value>
				</Item>
			</Valuta>`,
			outputStruct: []models.Currency{
				{ISOName: "USD"},
				{ISOName: "EUR"},
			},
		},
		{
			xmlData:      `data: black`,
			outputStruct: nil,
		},
	}

	for _, test := range tests {
		mockData := io.NopCloser(strings.NewReader(test.xmlData))
		currencies, _ := parser.ParseCurrencyISO(mockData)
		assert.Equal(t, test.outputStruct, currencies)
	}

}

func TestParse(t *testing.T) {
	parser := XMLParser{}

	tests := []struct {
		xmlData      string
		outputStruct map[string]float32
	}{
		{
			xmlData: `
			<ValCurs>
				<Valute>
					<CharCode>USD</CharCode>
					<Value>75,5</Value>
				</Valute>
				<Valute>
					<CharCode>EUR</CharCode>
					<Value>88,2</Value>
				</Valute>
			</ValCurs>`,
			outputStruct: map[string]float32{
				"USD": 75.5,
				"EUR": 88.2,
			},
		},
		{
			xmlData:      `data: black`,
			outputStruct: nil,
		},
		{
			xmlData: `
			<Valuta>
				<Item>
					<ISO_Char_Code>USD</ISO_Char_Code>
					<Value>75,,</Value>
				</Item>
				<Item>
					<ISO_Char_Code>EUR</ISO_Char_Code>
					<Value>88,,</Value>
				</Item>
			</Valuta>`,
			outputStruct: nil,
		},
	}

	for _, test := range tests {
		mockData := io.NopCloser(strings.NewReader(test.xmlData))
		rates, _ := parser.Parse(mockData)
		assert.Equal(t, test.outputStruct, rates)
	}
}
