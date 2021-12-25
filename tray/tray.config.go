package tray

import "App/tray/locales"

type Config struct {
	Locale      locales.Locale `comment:"Locale options: en, zh"`
	StartServer bool
	EnableSwag  bool
}

func DefaultConfig() *Config {
	return &Config{
		Locale:      locales.Zh,
		StartServer: true,
		EnableSwag:  true,
	}
}
