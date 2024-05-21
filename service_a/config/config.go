package config

import (
	"os"
)

type Conf struct {
	UrlServiceB string `mapstructure:"URL_SERVICE_B"`
	UrlZipKin   string `mapstructure:"URL_ZIPKIN"`
}

func LoadConfig() *Conf {
	return &Conf{
		UrlServiceB: os.Getenv("URL_SERVICE_B"),
		UrlZipKin:   os.Getenv("URL_ZIPKIN"),
	}
}
