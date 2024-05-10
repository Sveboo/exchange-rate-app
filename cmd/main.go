package main

import (
	"log"

	_ "github.com/sveboo/exchangerate/docs"

	"github.com/sveboo/exchangerate/internal/app"
	"github.com/sveboo/exchangerate/internal/config"
	"github.com/sveboo/exchangerate/internal/ports/httpgin"
)

//	@title			Exchange rate application documentation
//	@version		1.0.0
//	@description	A collection of endpoints available to retrieve the exchange rate for a specific date and currency.

// @contact.name	Maintainer
// @contact.url		https://github.com/Sveboo/exchange-rate-app
// @contact.email	svebo3348@gmail.com
// @host			localhost:8000
// @accept			json
// @produce		json
// @schemes		http
func main() {
	// get app configuration
	confPath := "config.yml"
	conf, err := config.GetConfig(confPath)
	if err != nil {
		log.Printf("Enable to get configuration. Error: %s", err)
		log.Printf("Enable to get configuration. Error: %s", err)
		return
	}
	// create app info using configuration
	a, err := app.Run(conf)
	if err != nil {
		log.Printf("Enable to create exchange rate app. Error: %s", err)
		log.Printf("Enable to create exchange rate app. Error: %s", err)
	}

	err = httpgin.Run(a, conf.Port)
	if err != nil {
		log.Printf("Run server error %s", err)
		log.Printf("Run server error %s", err)
		return
	}
}
