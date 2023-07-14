package config

import "time"

const (
	DefaultLimit              = 5
	DefaultOffset             = 0
	DurationOfTheUpdatePeriod = time.Minute
	DriverName                = "postgres"
	Host                      = "localhost"
	Port                      = 5432
	Dbname                    = "db"
	SslMode                   = "disable"
	HttpAddress               = "localhost:8080"
)
