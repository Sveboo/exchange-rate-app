package httpgin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sveboo/exchangerate/internal/models"
	mock_interfaces "github.com/sveboo/exchangerate/mocks"
)

func TestGetInfoHandler(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockApp := mock_interfaces.NewMockApp(ctr)

	Info := models.InfoApp{
		Service: "service",
		Version: "0.1.1",
		Author:  "n.lastname",
	}

	expectedInfo :=
		`{"version":"0.1.1","service":"service","author":"n.lastname"}`
	mockApp.EXPECT().GetInfo().Return(Info)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler := getInfo(mockApp)
	handler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedInfo, w.Body.String())
}

func TestGetExchangeRateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := mock_interfaces.NewMockApp(ctrl)

	ExchangeRate := map[string]float32{"USD": 33.4013}
	mockApp.EXPECT().IsCurrencyValid(gomock.Any()).Return(true).AnyTimes()
	mockApp.EXPECT().GetCurrencyRate(gomock.Any(), gomock.Any()).Return(ExchangeRate, nil).AnyTimes()

	expectedJSON := `{"service":"currency","data":{"USD":33.4013}}`
	tests := []struct {
		name           string
		queryParams    map[string]string
		expectedStatus int
	}{
		{
			name:           "No query parameters",
			queryParams:    map[string]string{},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Valid date format",
			queryParams:    map[string]string{"date": "2024-03-30"},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Valid currency",
			queryParams:    map[string]string{"currency": "XYZ"},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest(http.MethodGet, "/currency", nil)
			q := req.URL.Query()
			for k, v := range tc.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()

			c.Request = req
			getExchangeRate(mockApp)(c)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, expectedJSON, w.Body.String())
		})
	}
}
