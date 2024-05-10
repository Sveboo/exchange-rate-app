package app

import (
	"net/http"
	"time"

	"github.com/sveboo/exchangerate/internal/config"
	currencyservice "github.com/sveboo/exchangerate/internal/currency_service"
	"github.com/sveboo/exchangerate/internal/errs"
	"github.com/sveboo/exchangerate/internal/models"
	"github.com/sveboo/exchangerate/internal/parser"
	"github.com/sveboo/exchangerate/internal/repository"
)

// The ExchangeRateApp struct represents an application for exchanging currency rates.
type ExchangeRateApp struct {
	parser          parser.Parser                    // parser defines methods for parsing data
	currencyService currencyservice.CurrencyServicer // currencyService defines methods for interacting with currency-related functionality
	repo            repository.CurrencyRepository    // repo interact with a data repository
	info            models.InfoApp                   // information about app
}

// fillRepo writes currency ISO names to a repository and returns an error if encountered.
func fillRepo(repo repository.CurrencyRepository, info []models.Currency) error {
	for _, currency := range info {
		if err := repo.WriteCurrency(currency.ISOName); err != nil {
			return err
		}
	}

	return nil
}

// The newExchangeRateApp function initializes an ExchangeRateApp instance.
func newExchangeRateApp(parser parser.Parser, currencyService currencyservice.CurrencyServicer, repo repository.CurrencyRepository, info models.InfoApp) (ExchangeRateApp, error) {
	currencyInfo, err := currencyService.GetCurrencyInfo()
	if err != nil {
		return ExchangeRateApp{}, err
	}

	parsedInfo, err := parser.ParseCurrencyISO(currencyInfo)
	if err != nil {
		return ExchangeRateApp{}, err
	}

	err = fillRepo(repo, parsedInfo)
	if err != nil {
		return ExchangeRateApp{}, err
	}

	a := ExchangeRateApp{
		parser:          parser,
		currencyService: currencyService,
		repo:            repo,
		info:            info,
	}

	return a, nil
}

// GetCurrencyRate retrieve the exchange rate for a specific date and currency.
func (a ExchangeRateApp) GetCurrencyRate(date, currency string) (map[string]float32, error) {
	exchangeRate, err := a.currencyService.GetExchangeRate(date)
	if err != nil {
		return nil, err
	}

	parsedRate, err := a.parser.Parse(exchangeRate)
	if err != nil {
		return nil, err
	}

	if len(parsedRate) == 0 {
		return nil, errs.ErrRateNotFound
	}
	if currency != "" {
		currencyRate := make(map[string]float32, 1)
		currencyRate[currency] = parsedRate[currency]
		return currencyRate, nil
	}
	return parsedRate, nil
}

// IsCurrencyValid reports whether currency is valid or not. It calls the `IsCurrencyExist` method of a.repo
// and returns a boolean value indicating whether the currency exists in the repository or not.
func (a ExchangeRateApp) IsCurrencyValid(currency string) bool {
	return a.repo.IsCurrencyExist(currency)
}

// GetInfo returns the `info` which contains information about the application.
func (a ExchangeRateApp) GetInfo() models.InfoApp {
	return a.info
}

func makeInfoApp(conf config.Config) models.InfoApp {
	return models.InfoApp{
		Version: conf.Version,
		Service: conf.Service,
		Author:  conf.Author,
	}
}

func Run(conf config.Config) (ExchangeRateApp, error) {
	infoApp := makeInfoApp(conf)
	client := currencyservice.NewHTTPClient(http.Client{Timeout: 10 * time.Second})

	n := currencyservice.NewCBRService(conf.CurrenciesListURL, conf.CurrenciesRatesURL, client)
	p := parser.XMLParser{}
	rep := repository.NewMapRepo()
	a, err := newExchangeRateApp(p, n, rep, infoApp)
	if err != nil {

		return ExchangeRateApp{}, err
	}

	return a, nil
}
