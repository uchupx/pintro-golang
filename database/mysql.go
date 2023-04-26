package database

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	HostName string
	Username string
	Password string
	Database string
}

func NewConnection(config Config) (*sqlx.DB, error) {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = config.HostName
	mysqlConfig.User = config.Username
	mysqlConfig.Passwd = config.Password
	mysqlConfig.DBName = config.Database

	db, err := sqlx.Open("mysql", mysqlConfig.FormatDSN())

	if err != nil {
		return nil, err
	}

	db = db.Unsafe()
	db.SetConnMaxLifetime(3 * time.Minute)

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
