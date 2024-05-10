package repository

// CurrencyRepository defines methods for writing currency data and checking if a currency exists.
type CurrencyRepository interface {
	// WriteCurrency write a currency ISO code in the repository and returns an error if the operation fails.
	WriteCurrency(ISOName string) error
	// IsCurrencyExist reports whether currency with the given ISOName exists in the repository.
	IsCurrencyExist(ISOName string) bool
}
