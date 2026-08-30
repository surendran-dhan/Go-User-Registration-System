package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	connStr := "host=localhost port=5432 user=postgres password=YOUR_PASSWORD dbname=login_db sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Database is not connected:", err)
		return
	}

	fmt.Println("Database connected successfully!")

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/register", registerHandler)

	fmt.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	_, err := db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		username,
		password,
	)

	if err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User registered successfully!")
}
