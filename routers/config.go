package routers

import (
	"os"
	"path/filepath"
)

type Config struct {
	IsDev   bool
	Port    uint16
	PortTls uint16
	CertDir string
}

func DefaultConfig() *Config {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return &Config{
		IsDev:   true,
		Port:    8080,
		PortTls: 8443,
		CertDir: filepath.Dir(path),
	}
}
