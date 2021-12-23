package main

import (
	"App/logger"
	"App/models"
	"App/routers"
	"App/tray"
	"flag"
	"os"

	"gopkg.in/ini.v1"
)

var hasTray = flag.Int("t", 1, "set to enable system tray")

func main() {
	flag.Parse()

	config := GetConfig()
	loadIni(config)

	logger.Setup(config.Logger)

	models.SetConfig(config.Models)
	routers.SetConfig(config.Routers)
	tray.SetConfig(config.Tray)

	models.Run()
	if *hasTray == 1 {
		tray.Start()
	} else {
		routers.Start(true, true)
	}
}

func loadIni(config *Config) {
	_, err := os.Stat(config.IniPath)
	if err != nil {
		if os.IsNotExist(err) {
			iniFile := ini.Empty()
			err := ini.ReflectFrom(iniFile, config)
			if err != nil {
				panic("Failed to create ini: " + config.IniPath)
			}
			iniFile.SaveTo(config.IniPath)
			return
		}
		panic("Error when loading ini: " + err.Error())
	}

	iniFile, err := ini.Load(config.IniPath)
	if err != nil {
		panic("Failed to log ini: " + config.IniPath)
	}
	iniFile.MapTo(config)
}
