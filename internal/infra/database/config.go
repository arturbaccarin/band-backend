package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func OpenConnection(dbUser, dbPass, dbHost, dbName, webServerPort string) *sql.DB {
	dbAddress := dbHost + ":" + webServerPort

	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPass,
		Net:    "tcp",
		Addr:   dbAddress,
		DBName: dbName,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db
}
