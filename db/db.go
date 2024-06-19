package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)

var gormDB *gorm.DB

func DB() *gorm.DB {
	return gormDB
}

func InitDB(dsn string) {
	InitDBWithLogLevel(dsn, Info)
}

func InitDBWithLogLevel(dsn string, logLevel LogLevel) {
	log.Info().Msg("InitDB:数据库连接中...")
	level := func(logLevel LogLevel) logger.LogLevel {
		switch logLevel {
		case Silent:
			return logger.Silent
		case Error:
			return logger.Error
		case Warn:
			return logger.Warn
		case Info:
			return logger.Info
		}
		return logger.Silent
	}(logLevel)
	var err error
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("InitDB:数据库连接失败")
	} else {
		log.Info().Msg("InitDB:数据库连接成功")
	}
}
