package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//User : user struct
type User struct {
	id   int
	name string
}

func main() {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	url := fmt.Sprintf("root:%s@/golang_course", password)
	db, error := sql.Open("mysql", url)

	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT * FROM users WHERE id > ?", 1)
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.id, &user.name)

		log.Println(user)
	}
}
