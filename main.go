package main

import (
	"App/models"
	"App/routers"
	"App/tray"
	"flag"
)

var hasTray = flag.Int("t", 1, "set to enable system tray")

func main() {
	flag.Parse()
	// config := GetConfig()

	// models.SetConfig(config.Models)
	models.Run()

	if *hasTray == 1 {
		// tray.SetConfig(config.Tray)
		tray.Start()
	} else {
		routers.Start(true, true)
	}
}
