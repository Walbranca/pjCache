package src

import "github.com/xMota14/postgresAux"

var PostgresServer = postgresAux.Server{
	Host:     "postgres",
	Port:     5432,
	User:     "postgres",
	Password: "postgres",
	Database: "pjcache",
}
