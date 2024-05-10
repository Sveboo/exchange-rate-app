// Package config cotains configuration settings for the application.
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	CurrenciesListURL  string `yaml:"currenciesListURL"`
	CurrenciesRatesURL string `yaml:"currenciesRatesURL"`
	Port               string `yaml:"port"`
	Service            string `yaml:"service"`
	Author             string `yaml:"author"`
	Version            string `yaml:"version"`
}

func parseConfig(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return &Config{}, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	var conf Config
	err = decoder.Decode(&conf)
	if err != nil {
		return &Config{}, err
	}

	return &conf, err
}

func GetConfig(fileName string) (Config, error) {
	cnf, err := parseConfig(fileName)
	if err != nil {
		return Config{}, err
	}
	version := os.Getenv("VERSION")
	if version != "" {

		cnf.Version = version
	}

	//get port from environment
	port := os.Getenv("PORT")
	if port != "" {
		cnf.Port = port
	}
	cnf.Port = ":" + cnf.Port
	return *cnf, nil
}
