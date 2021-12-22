package logger

import (
	"os"
	"path/filepath"
)

type Config struct {
	IsDev          bool
	LogTrayPath    string
	LogRoutersPath string
}

func DefaultConfig() *Config {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return &Config{
		IsDev:          true,
		LogTrayPath:    filepath.Join(filepath.Dir(path), "log.tray.txt"),
		LogRoutersPath: filepath.Join(filepath.Dir(path), "log.routers.txt"),
	}
}
