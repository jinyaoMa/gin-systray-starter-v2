package tray

import (
	"App/tray/locales"
)

type Config struct {
	Locale          *locales.Dictionary
	EnableSwagger2  bool
	AutoStartServer bool
}

func DefaultConfig() Config {
	return Config{
		Locale:          locales.En(),
		EnableSwagger2:  true,
		AutoStartServer: true,
	}
}
