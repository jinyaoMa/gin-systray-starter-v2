package main

import (
	"App/logger"
	"App/routers"
	"App/tray"
	"App/tray/locales"
	"os"
	"path/filepath"
)

type Config struct {
	Logger  *logger.Config
	Tray    *tray.Config
	Routers *routers.Config
}

func GetConfig() *Config {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return &Config{
		Logger: &logger.Config{
			IsDev:          false,
			LogTrayPath:    filepath.Join(filepath.Dir(path), "log.tray.txt"),
			LogRoutersPath: filepath.Join(filepath.Dir(path), "log.routers.txt"),
		},
		Tray: &tray.Config{
			Locale:      locales.Zh,
			StartServer: true,
			EnableSwag:  true,
		},
		Routers: &routers.Config{
			IsDev:   false,
			Port:    18080,
			PortTls: 18443,
			CertDir: filepath.Dir(path),
		},
	}
}
