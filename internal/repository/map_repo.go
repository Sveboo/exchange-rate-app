// Package repository provides type to store data
package repository

// MapRepo contains a map with string keys and empty struct values.
type MapRepo struct {
	repo map[string]struct{}
}

// The NewMapRepo function creates a new MapRepo instance with an empty map repository.
func NewMapRepo() MapRepo {
	r := make(map[string]struct{})
	return MapRepo{
		repo: r,
	}
}

// WriteCurrency add a new entry to the map repository with the key `ISOName` and an empty struct value.
func (m MapRepo) WriteCurrency(ISOName string) error {
	m.repo[ISOName] = struct{}{}
	return nil
}

// IsCurrencyExist reports whether currency with the given ISOName exists in the map repository.
func (m MapRepo) IsCurrencyExist(ISOName string) bool {
	_, ok := m.repo[ISOName]
	return ok
}
