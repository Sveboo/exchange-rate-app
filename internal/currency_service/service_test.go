package currencyservice

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_interfaces "github.com/sveboo/exchangerate/mocks"
)

func TestGetCurrencyInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_interfaces.NewMockClient(ctrl)
	service := NewCBRService("", "", client)
	respBody := "mock response body"

	tests := []struct {
		response *http.Response
		err      error
	}{
		{
			response: &http.Response{
				Body:       io.NopCloser(bytes.NewBufferString(respBody)),
				StatusCode: http.StatusOK,
			},
			err: nil,
		},
		{
			response: &http.Response{
				Body:       nil,
				StatusCode: http.StatusBadRequest,
			},
			err: fmt.Errorf("some error"),
		},
	}

	for _, test := range tests {
		client.EXPECT().Do(gomock.Any()).Return(test.response, test.err)
		body, err := service.GetCurrencyInfo()
		if err == nil {
			defer body.Close()

			info, _ := io.ReadAll(body)
			assert.Equal(t, respBody, string(info))
		}

		assert.Equal(t, test.err, err)

	}
}

func TestGetExchangeRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_interfaces.NewMockClient(ctrl)
	service := NewCBRService("", "", client)
	respBody := "mock response body"
	date := "02/01/2006"

	tests := []struct {
		response *http.Response
		err      error
	}{
		{
			response: &http.Response{
				Body:       io.NopCloser(bytes.NewBufferString(respBody)),
				StatusCode: http.StatusOK,
			},
			err: nil,
		},
		{
			response: &http.Response{
				Body:       nil,
				StatusCode: http.StatusBadRequest,
			},
			err: fmt.Errorf("some error"),
		},
	}

	for _, test := range tests {
		client.EXPECT().Do(gomock.Any()).Return(test.response, test.err)
		body, err := service.GetExchangeRate(date)
		if err == nil {
			defer body.Close()

			info, _ := io.ReadAll(body)
			assert.Equal(t, respBody, string(info))
		}

		assert.Equal(t, test.err, err)

	}
}
