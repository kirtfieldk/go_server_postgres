package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "orders"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	/* Actually connects to DB */
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	makeTables(db)
	// http.HandleFunc("/", lp_routes.MainPage)
	// http.HandleFunc("/search", lp_routes.SearchRoute)
	// http.HandleFunc("/order", lp_routes.AddOrder)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeTables(db *sql.DB) {
	stm := `CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);`
	_, err := db.Exec(stm)
	if err != nil {
		panic(err)
	}
}
