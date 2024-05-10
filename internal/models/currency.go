// Package models contains data structures related to the application's domain model.
package models

// The Currency type in Go represents a currency with an ISO character code.
type Currency struct {
	ISOName string `xml:"ISO_Char_Code"`
}
