package httpgin

import (
	"github.com/gin-gonic/gin"
	"github.com/sveboo/exchangerate/internal/models"
)

type Response struct {
	Service string             `json:"service" example:"application"`
	Data    map[string]float32 `json:"data" example:"USD:33.4013"`
}

type InfoResponse struct {
	Version string `json:"version" example:"0.0.0"`
	Service string `json:"service" example:"application"`
	Author  string `json:"author" example:"n.lastname"`
}

func createInfoResp(info models.InfoApp) InfoResponse {
	return InfoResponse{
		Version: info.Version,
		Service: info.Service,
		Author:  info.Author,
	}
}

func createSuccessCurResp(service string, currencies map[string]float32) *gin.H {
	return &gin.H{
		"service": service,
		"data":    currencies,
	}
}

func createErrorCurResp(service string, err error) *gin.H {
	return &gin.H{
		"service": service,
		"data":    nil,
		"error":   err.Error(),
	}
}
