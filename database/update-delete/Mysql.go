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

	insert, _ := db.Prepare("INSERT INTO users(name) VALUES(?)")
	insert.Exec("Paty")

	update, _ := db.Prepare("UPDATE users SET name = ? WHERE name = ?")
	update.Exec("Paty")

	delete, _ := db.Prepare("DELETE FROM users WHERE name = ?")
	delete.Exec("Paty")
}
