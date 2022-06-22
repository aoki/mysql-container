package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(0.0.0.0:3306)/admin")
	if err != nil {
		os.Exit(1)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		os.Exit(2)
	}
}
