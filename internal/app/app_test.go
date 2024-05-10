package app

import (
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sveboo/exchangerate/internal/errs"
	"github.com/sveboo/exchangerate/internal/models"
	mock_interfaces "github.com/sveboo/exchangerate/mocks"
)

func TestGetCurrencyRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockParser := mock_interfaces.NewMockParser(ctrl)
	mockCurrencyService := mock_interfaces.NewMockCurrencyServicer(ctrl)
	mockRepo := mock_interfaces.NewMockRepository(ctrl)
	mockInfo := models.InfoApp{}

	respBody := `
	<ValCurs>
		<Valute>
			<CharCode>USD</CharCode>
			<Value>75,5</Value>
		</Valute>
	</ValCurs>`
	mockData := io.NopCloser(strings.NewReader(respBody))
	date := "01/02/2016"
	someError := fmt.Errorf("some error")
	parsedRate := map[string]float32{
		"USD": 75.5,
		"EUR": 88.2,
	}

	currency := "USD"
	curencyRate := map[string]float32{
		"USD": 75.5,
	}

	mockCurrencyService.EXPECT().GetCurrencyInfo().Return(mockData, nil)
	mockParser.EXPECT().ParseCurrencyISO(gomock.Any()).Return([]models.Currency{}, nil)
	app, err := newExchangeRateApp(mockParser, mockCurrencyService, mockRepo, mockInfo)

	if err != nil {
		log.Println("Enable to create exchange rate app")
		return
	}
	tests := []struct {
		errService   error
		parsedRate   map[string]float32
		errParser    error
		currency     string
		exchangeRate map[string]float32
		errApp       error
	}{
		{
			errService:   nil,
			parsedRate:   nil,
			errParser:    someError,
			currency:     "",
			exchangeRate: nil,
			errApp:       someError,
		},
		{
			errService:   nil,
			parsedRate:   map[string]float32{},
			errParser:    nil,
			currency:     "",
			exchangeRate: nil,
			errApp:       errs.ErrRateNotFound,
		},
		{
			errService:   nil,
			parsedRate:   parsedRate,
			errParser:    nil,
			currency:     "",
			exchangeRate: parsedRate,
			errApp:       nil,
		},
		{
			errService:   nil,
			parsedRate:   parsedRate,
			errParser:    nil,
			currency:     currency,
			exchangeRate: curencyRate,
			errApp:       nil,
		},
	}

	for _, test := range tests {
		mockCurrencyService.EXPECT().GetExchangeRate(gomock.Any()).Return(mockData, test.errService)
		mockParser.EXPECT().Parse(gomock.Any()).Return(test.parsedRate, test.errParser)
		rate, err := app.GetCurrencyRate(date, test.currency)
		assert.Equal(t, test.exchangeRate, rate)
		assert.Equal(t, test.errApp, err)
	}
}

func TestIsCurrencyValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockParser := mock_interfaces.NewMockParser(ctrl)
	mockCurrencyService := mock_interfaces.NewMockCurrencyServicer(ctrl)
	mockRepo := mock_interfaces.NewMockRepository(ctrl)
	mockInfo := models.InfoApp{}

	mockData := io.NopCloser(strings.NewReader("respBody"))
	mockCurrencyService.EXPECT().GetCurrencyInfo().Return(mockData, nil)
	mockParser.EXPECT().ParseCurrencyISO(gomock.Any()).Return([]models.Currency{}, nil)
	app, err := newExchangeRateApp(mockParser, mockCurrencyService, mockRepo, mockInfo)

	if err != nil {
		log.Printf("Enable to create exchange rate app. Error:%s", err)
	}
	tests := []struct {
		currency string
		isExist  bool
	}{
		{
			currency: "USD",
			isExist:  true,
		},

		{
			currency: "UUU",
			isExist:  false,
		},
	}

	for _, test := range tests {
		mockRepo.EXPECT().IsCurrencyExist(gomock.Any()).Return(test.isExist)
		ok := app.IsCurrencyValid(test.currency)
		assert.Equal(t, test.isExist, ok)
	}
}

func TestGetInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInfo := models.InfoApp{
		Service: "service",
		Version: "0.1.1",
		Author:  "n.lastname",
	}

	app := ExchangeRateApp{
		info: mockInfo,
	}

	info := app.GetInfo()

	assert.Equal(t, mockInfo, info)
}
