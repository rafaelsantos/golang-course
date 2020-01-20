package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	url := fmt.Sprintf("root:%s@/golang_course", password)
	db, error := sql.Open("mysql", url)

	if error != nil {
		panic(error)
	}

	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO users(name) VALUES(?)")
	statement.Exec("Johnny")
	result, _ := statement.Exec("Mary")

	id, _ := result.LastInsertId()
	fmt.Println(id)

	lines, _ := result.RowsAffected()
	fmt.Println(lines)
}
