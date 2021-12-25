package main

import (
	"App/logger"
	"App/models"
	"App/routers"
	"App/tray"
	"os"
	"path/filepath"
)

type Config struct {
	IniPath string `ini:"-"`
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
		IniPath: filepath.Join(filepath.Dir(path), "App.ini"),
		Logger:  logger.DefaultConfig(),
		Tray:    tray.DefaultConfig(),
		Routers: routers.DefaultConfig(),
		Models:  models.DefaultSqliteConfig(),
	}
}
