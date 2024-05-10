// Package currencyservice implements utilities currency servicer for retrieving currency
// information and exchange rates from an external service (CBR - Central Bank
// of Russia).
package currencyservice

import (
	"io"
	"net/http"
	"net/url"

	"github.com/corpix/uarand"
)

// CBRService represents a service with URLs for currency and exchange rates, along with an HTTP client.
type CBRService struct {
	currencierURL    string // URL for accessing currency-related information
	exchangeRatesURL string // URL endpoint for fetching exchange rates from CBR
	client           Client // struct to make HTTP requests to external service CBR
}

// The NewCBRService function creates a new CBRService instance with specified URLs and an HTTP client
// with a timeout of 10 seconds.
func NewCBRService(currencierURL, exchangeRatesURL string, client Client) CBRService {
	return CBRService{
		currencierURL:    currencierURL,
		exchangeRatesURL: exchangeRatesURL,
		client:           client,
	}
}

// GetCurrencyInfo makes an HTTP GET request to currencierURL to retrieve currency-related information.
func (s CBRService) GetCurrencyInfo() (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", s.currencierURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", uarand.GetRandom())

	resp, err := s.client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// GetExchangeRate make an HTTP GET request to exchangeRatesURL with a specific date to retrieve exchange rate information.
func (s CBRService) GetExchangeRate(date string) (io.ReadCloser, error) {
	u, err := url.Parse(s.exchangeRatesURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("date_req", date)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", uarand.GetRandom())
	resp, err := s.client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
