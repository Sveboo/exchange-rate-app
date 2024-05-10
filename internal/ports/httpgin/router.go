package httpgin

import (
	"github.com/gin-gonic/gin"
	"github.com/sveboo/exchangerate/internal/app"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// AppRouter sets up routes for getting information and exchange rates in application.
func AppRouter(r gin.IRouter, a app.App) {
	r.GET("", getInfo(a))
	r.GET("/currency", getExchangeRate(a))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
