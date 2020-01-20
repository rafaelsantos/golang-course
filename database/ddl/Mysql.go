package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, error := db.Exec(sql)

	if error != nil {
		panic(error)
	}

	return result
}

func main() {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	url := fmt.Sprintf("root:%s@/", password)
	db, error := sql.Open("mysql", url)

	if error != nil {
		panic(error)
	}

	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS golang_course")
	exec(db, "USE golang_course")
	exec(db, "DROP TABLE IF EXISTS users")
	exec(db, `CREATE TABLE users(
		id INTEGER AUTO_INCREMENT,
		name VARCHAR(80),
		PRIMARY KEY(id))`,
	)
}
