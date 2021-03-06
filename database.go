package learn_golang_database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/go_db?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
