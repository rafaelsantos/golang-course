package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//User : user struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//UserHandler : analyses the request and delegate to properly function
func UserHandler(writer http.ResponseWriter, request *http.Request) {
	//Remove URL path and /users/ from string
	sid := strings.TrimPrefix(request.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case request.Method == "GET" && id > 0:
		userByID(writer, request, id)
	case request.Method == "GET":
		allUsers(writer, request)
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "Sorry... Not Found!")
	}
}

func userByID(writer http.ResponseWriter, request *http.Request, id int) {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	url := fmt.Sprintf("root:%s@/golang_course", password)
	db, error := sql.Open("mysql", url)

	if error != nil {
		panic(error)
	}

	defer db.Close()

	var user User
	db.QueryRow("SELECT id, name FROM users WHERE id= ?", id).Scan(&user.ID, &user.Name)

	json, _ := json.Marshal(user)

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(writer, string(json))
}

func allUsers(writer http.ResponseWriter, request *http.Request) {
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	url := fmt.Sprintf("root:%s@/golang_course", password)
	db, error := sql.Open("mysql", url)

	if error != nil {
		panic(error)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT id, name FROM users")
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)

		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(writer, string(json))
}
