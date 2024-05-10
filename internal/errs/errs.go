// Package errs contains internal errors
package errs

import "fmt"

var (
	ErrRateNotFound    = fmt.Errorf("no rate for this date")
	ErrInvalidCurrency = fmt.Errorf("invalid currency")
	ErrInvalidDate     = fmt.Errorf("invalid date")
)
