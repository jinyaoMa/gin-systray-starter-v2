package models

import (
	"App/logger"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	config *Config
	db     *gorm.DB
)

func init() {
	config = DefaultSqliteConfig()
}

func Run() {
	switch config.Driver {
	case DriverSqlite:
		initSqlite()
	case DriverMysql:
		initMysql()
	case DriverPostgres:
		initPostgres()
	}

	db.AutoMigrate(&Record{})
}

func SetConfig(conf *Config) {
	config = conf
}

func initSqlite() {
	var err error
	db, err = gorm.Open(sqlite.Open(config.Database), &gorm.Config{})
	if err != nil {
		logger.Models.Fatalf("database (sqlite) connect error %v\n", err)
	}
	if db.Error != nil {
		logger.Models.Fatalf("database (sqlite) error %v\n", db.Error)
	}
}

func initMysql() {
	var err error
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Tail)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Models.Fatalf("database (mysql) connect error %v\n", err)
	}
	if db.Error != nil {
		logger.Models.Fatalf("database (mysql) error %v\n", db.Error)
	}
}

func initPostgres() {
	var err error
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s",
		config.Host,
		config.User,
		config.Password,
		config.Database,
		config.Port,
		config.Tail)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Models.Fatalf("database (postgres) connect error %v\n", err)
	}
	if db.Error != nil {
		logger.Models.Fatalf("database (postgres) error %v\n", db.Error)
	}
}
