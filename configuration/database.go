package configuration

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:hbx@tcp(172.17.0.1:3306)/go-htmx")

	if err != nil {
		panic(err.Error())
	}

	return db
}
