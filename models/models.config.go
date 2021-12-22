package models

import (
	"os"
	"path/filepath"
)

const (
	DriverSqlite   Driver = "sqlite"
	DriverMysql    Driver = "mysql"
	DriverPostgres Driver = "postgres"
)

type Driver string

type Config struct {
	Driver   Driver
	Host     string
	Port     uint16
	User     string
	Password string
	Database string
	Tail     string
}

func DefaultSqliteConfig() *Config {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return &Config{
		Driver:   DriverSqlite,
		Host:     "",
		Port:     0,
		User:     "admin",
		Password: "admin",
		Database: filepath.Join(filepath.Dir(path), "database.db"),
		Tail:     "_auth_crypt=sha1",
	}
}
