package httpgin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sveboo/exchangerate/internal/app"
	"github.com/sveboo/exchangerate/internal/errs"
)

// getInfo takes an `App` interface and retrieves information in JSON format using GetInfo.
//
//	@Summary		Get info
//	@Description	Get info about exchenge rate app
//	@Produce		json
//	@Router			/ [get]
//
//	@Success		200	{object}	InfoResponse	"Information has been received successfully"
func getInfo(a app.App) gin.HandlerFunc {
	infoApp := a.GetInfo()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, createInfoResp(infoApp))
	}
}

// The function `getExchangeRate`retrieves currency exchange rates based on user input.
//
//	@Summary		Get exchange rate
//	@Description	Returns exchange rates based on user input
//	@Produce		json
//	@Param			currency	query	string	false	"ISO currency name"			example(USD)
//	@Param			date		query	string	false	"date in format YYYY-MM-DD"	example(2016-01-02)
//	@Router			/currency{currency}{date} [get]
//
//	@Success		200	{object}	Response	"Exchange rate has been received successfully"
//
//	@Failure		400	{object}	Response	"Invalid query parameters"
//	@Failure		500	{object}	Response	"Exchange rate receiving caused error"
func getExchangeRate(a app.App) gin.HandlerFunc {
	service := "currency"
	return func(c *gin.Context) {
		var (
			currency, date string
			ok             bool
		)

		if currency, ok = c.GetQuery("currency"); ok {
			if !a.IsCurrencyValid(currency) {
				c.JSON(http.StatusBadRequest, createErrorCurResp(service, errs.ErrInvalidCurrency))
				return
			}
		}

		inputLayout := "2006-01-02"
		correctLayout := "02/01/2006"
		if date, ok = c.GetQuery("date"); ok {
			time, err := time.Parse(inputLayout, date)
			if err != nil {
				c.JSON(http.StatusBadRequest, createErrorCurResp(service, errs.ErrInvalidDate))
				return
			}

			date = time.Format(correctLayout)
		}

		exchangeRate, err := a.GetCurrencyRate(date, currency)
		if err != nil {
			c.JSON(http.StatusInternalServerError, createErrorCurResp(service, err))
			return
		}

		c.JSON(http.StatusOK, createSuccessCurResp(service, exchangeRate))
	}
}
