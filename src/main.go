package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(mysql:3306)/performance_schema")
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		os.Exit(2)
	}

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("db.Query error err: %v", err)
	}

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("rows.Err error err: %v", err)
	}
}
