// Package app provides functionality for exchange rate app
package app

import "github.com/sveboo/exchangerate/internal/models"

// App provides currency rate information and validation functionality.
type App interface {
	//	GetCurrencyRate retrieve currency exchange rates for a specific date and currency and
	// 	an error if any issue occurs during the retrieval process.
	// If the date is not specified rate returns for the current day
	// 	The returned map contain currency as a key and its exchange rates as a value.
	GetCurrencyRate(date, currency string) (map[string]float32, error)
	//	IsCurrencyValid reports whether currency is valid or not.
	IsCurrencyValid(currency string) bool
	// GetInfo retrieve information about the application.
	GetInfo() models.InfoApp
}
