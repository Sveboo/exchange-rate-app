package currencyservice

import "io"

// The CurrencyServicer defines methods for retrieving currency information and exchange rates.

type CurrencyServicer interface {
	// GetCurrencyInfo retrieve io.ReadCloser that can be used to read the currency information
	// and an error if any issue occurs during the retrieval process.
	GetCurrencyInfo() (io.ReadCloser, error)
	// GetExchangeRate retrieves io.ReadCloser that can be used to read the exchange rate for a specific date
	// and an error if any issue occurs during the retrieval process.
	// If the date is not specified rate returns for the current day
	GetExchangeRate(date string) (io.ReadCloser, error)
}
