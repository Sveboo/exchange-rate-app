package parser

import (
	"io"

	"github.com/sveboo/exchangerate/internal/models"
)

// Parser is used for parsing data and currency codes.
type Parser interface {
	// Parse parse data from an io.ReadCloser which contains an exchange rate
	// and return a map where currency is a key and its exchange rates is a value.
	Parse(data io.ReadCloser) (map[string]float32, error)
	// ParseCurrencyISO parse currency from io.ReadCloser and return a slice of
	// `models.Currency` objects and an error if any issue occurs during the retrieval process.
	ParseCurrencyISO(data io.ReadCloser) ([]models.Currency, error)
}
