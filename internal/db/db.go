package db

import (
	"cache-ahead-task/internal/config"
	. "cache-ahead-task/internal/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func Connect() *sqlx.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, os.Getenv("USER"), os.Getenv("PASSWORD"), config.Dbname, config.SslMode)
	db, err := sqlx.Connect(config.DriverName, dataSourceName)
	if err != nil {
		Logger.Fatal(err.Error())
		return db
	}
	Logger.Info(fmt.Sprintf("Сreated a connection to %s", config.DriverName))
	return db
}
