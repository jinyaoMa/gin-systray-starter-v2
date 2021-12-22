package logger

import (
	"io"
	"log"
	"os"
)

var config *Config

var (
	Tray    *log.Logger
	Routers *log.Logger
)

func init() {
	Setup(DefaultConfig())
}

func Setup(conf *Config) {
	config = conf
	handleConfig()
}

func handleConfig() {
	if config.IsDev {
		Tray = log.New(os.Stdout,
			"TRAY: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Routers = log.New(os.Stdout,
			"ROUTERS: ",
			log.Ldate|log.Ltime|log.Lshortfile)

	} else {
		logTray, err := os.OpenFile(config.LogTrayPath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Failed to open error log (tray) file")
		}

		logRouters, err := os.OpenFile(config.LogRoutersPath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Failed to open error log (routers) file")
		}

		Tray = log.New(io.MultiWriter(logTray, os.Stdout),
			"TRAY: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Routers = log.New(io.MultiWriter(logRouters, os.Stdout),
			"ROUTERS: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}
}
