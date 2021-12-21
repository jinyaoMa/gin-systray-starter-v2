package main

import (
	"App/routers"
	"App/tray"
	"log"
)

func main() {
	// routers.SetConfig(config.RoutersConfig)
	// routers.StartWithLoop(true)

	// tray.SetConfig(config.TrayConfig)
	tray.Run(&tray.Listener{
		OnReady: func() {
			log.Println("App starting!")
		},
		OnExit: func() {
			log.Println("App exiting!")
		},
		OnServer: func(run bool, withSwag bool) bool {
			if run {
				return routers.Start(withSwag)
			} else {
				routers.Stop()
			}
			return false
		},
	})
}
