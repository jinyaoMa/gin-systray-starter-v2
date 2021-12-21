package tray

import (
	"App/tray/locales"
)

type Config struct {
	Locale        *locales.Dictionary
	StartServer   bool
	EnableSwagger bool
}

func DefaultConfig() *Config {
	return &Config{
		Locale:        locales.En,
		StartServer:   true,
		EnableSwagger: true,
	}
}
