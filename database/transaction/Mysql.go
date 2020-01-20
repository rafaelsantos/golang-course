package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	url := fmt.Sprintf("root:%s@/golang_course", password)
	db, error := sql.Open("mysql", url)

	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	transaction, _ := db.Begin()
	statement, _ := transaction.Prepare("INSERT INTO users(id, name) VALUES(?,?)")
	statement.Exec(234, "Paul")
	statement.Exec(245, "Marissa")

	_, error = statement.Exec(1, "July")

	if error != nil {
		transaction.Rollback()
		log.Fatal(error)
	}

	transaction.Commit()
}
