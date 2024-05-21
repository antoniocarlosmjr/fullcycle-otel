package config

import (
	"os"
)

type Conf struct {
	UrlZipKin string `mapstructure:"URL_ZIPKIN"`
}

func LoadConfig() *Conf {
	return &Conf{
		UrlZipKin: os.Getenv("URL_ZIPKIN"),
	}
}
