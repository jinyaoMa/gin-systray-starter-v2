package main

import (
	"App/routers"
	"App/tray"
)

type Config struct {
	RoutersConfig routers.Config
	TrayConfig    tray.Config
}

var config = &Config{
	RoutersConfig: routers.DefaultConfig(),
	TrayConfig:    tray.DefaultConfig(),
}
