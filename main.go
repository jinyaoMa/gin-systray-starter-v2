package main

import (
	"App/routers"
	"App/tray"
	"flag"
)

var hasTray = flag.Int("t", 1, "set to enable system tray")

func main() {
	flag.Parse()

	if *hasTray == 1 {
		// config := GetConfig()

		// tray.SetConfig(config.Tray)
		tray.Start()
	} else {
		routers.Start(true, true)
	}
}
