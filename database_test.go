package learn_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenCon(t *testing.T) {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/go_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
