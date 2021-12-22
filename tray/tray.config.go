package tray

import "App/tray/locales"

type Config struct {
	Locale      locales.Locale
	StartServer bool
	EnableSwag  bool
}

func DefaultConfig() *Config {
	return &Config{
		Locale:      locales.En,
		StartServer: false,
		EnableSwag:  false,
	}
}
