package main

import (
	"App/logger"
	"App/models"
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
	Models  *models.Config
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
			LogModelsPath:  filepath.Join(filepath.Dir(path), "log.models.txt"),
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
		Models: &models.Config{
			Driver:   models.DriverSqlite,
			Host:     "",
			Port:     0,
			User:     "admin",
			Password: "admin",
			Database: filepath.Join(filepath.Dir(path), "database.db"),
			Tail:     "_auth_crypt=sha1",
		},
	}
}
