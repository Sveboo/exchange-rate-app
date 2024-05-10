package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteCurrency(t *testing.T) {
	m := NewMapRepo()
	ISOname := "USD"
	err := m.WriteCurrency(ISOname)
	assert.Equal(t, err, nil)

}

func TestIsCurrencyExist(t *testing.T) {
	m := NewMapRepo()
	ISOname := "USD"
	m.WriteCurrency(ISOname)
	tests := []struct {
		ISOname string
		ok      bool
	}{{
		ISOname: ISOname,
		ok:      true,
	},
		{
			ISOname: "not_existing",
			ok:      false,
		},
	}

	for _, test := range tests {
		ok := m.IsCurrencyExist(test.ISOname)
		assert.Equal(t, test.ok, ok)
	}
}
