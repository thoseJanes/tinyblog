package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
)

type MySQLDSN struct{
	User string
	Password string
	Host string
	Database string
}

type MySQLOptions struct{
	MySQLDSN
	MaxIdleConnections int
	MaxOpenConnections int
	MaxConnectionLifeTime time.Duration
	LogLevel int
}

func (dsn *MySQLDSN) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		dsn.User,
		dsn.Password,
		dsn.Host,
		dsn.Database,
		true,
		"Local")
}

func NewMySQL(opt *MySQLOptions) (*gorm.DB, error) {
	if opt.LogLevel == 0 {
		opt.LogLevel = int(logger.Silent)
	}

	db, err := gorm.Open(mysql.Open(opt.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(opt.LogLevel)),
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetConnMaxLifetime(opt.MaxConnectionLifeTime)
	sqlDb.SetMaxIdleConns(opt.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(opt.MaxOpenConnections)

	return db, nil
}

