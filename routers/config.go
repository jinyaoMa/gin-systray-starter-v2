package routers

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	HttpPort     string
	HttpsPort    string
	CertDirCache string
}

func DefaultConfig() *Config {
	certDirCache, err := os.Executable()
	if err != nil {
		log.Panic(err)
	}
	return &Config{
		HttpPort:     ":8080",
		HttpsPort:    ":8443",
		CertDirCache: filepath.Dir(certDirCache),
	}
}
