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
	Models  *log.Logger
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
		Tray = newLog(os.Stdout, "TRAY: ")
		Routers = newLog(os.Stdout, "ROUTERS: ")
		Models = newLog(os.Stdout, "MODELS: ")

	} else {
		logTray := getLogFile(config.LogTrayPath)
		logRouters := getLogFile(config.LogRoutersPath)
		logModels := getLogFile(config.LogModelsPath)

		Tray = newLog(io.MultiWriter(logTray, os.Stdout), "TRAY: ")
		Routers = newLog(io.MultiWriter(logRouters, os.Stdout), "ROUTERS: ")
		Models = newLog(io.MultiWriter(logModels, os.Stdout), "MODELS: ")
	}
}

func newLog(out io.Writer, prefix string) *log.Logger {
	return log.New(out, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}

func getLogFile(path string) (file *os.File) {
	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log: " + path)
	}
	return
}
