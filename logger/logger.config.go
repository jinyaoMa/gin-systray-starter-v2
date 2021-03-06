package logger

import (
	"os"
	"path/filepath"
)

type Config struct {
	IsDev          bool
	LogTrayPath    string
	LogRoutersPath string
	LogModelsPath  string
}

func DefaultConfig() *Config {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return &Config{
		IsDev:          true,
		LogTrayPath:    filepath.Join(filepath.Dir(path), "tray.log"),
		LogRoutersPath: filepath.Join(filepath.Dir(path), "routers.log"),
		LogModelsPath:  filepath.Join(filepath.Dir(path), "models.log"),
	}
}
