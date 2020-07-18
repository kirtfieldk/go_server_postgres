package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/keithkfield/pg_api/database"
	"github.com/keithkfield/pg_api/lp_routes"
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
	database.BuildTables(db)
	http.HandleFunc("/", lp_routes.MainPage)
	http.HandleFunc("/search", lp_routes.SearchRoute)
	http.HandleFunc("/order", lp_routes.AddOrder)
	http.HandleFunc("/api/v1/users", lp_routes.UsersRoute)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
