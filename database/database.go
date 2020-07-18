package database

import "database/sql"

type DatabaseBuilder interface {
	// Main Exportable Builder Function
	BuildTables(db *sql.DB)
	buildUsers(db *sql.DB)
	buildItems(db *sql.DB)
	buildOrders(db *sql.DB)
}

func BuildTables(db *sql.DB) {
	buildUsers(db)
	buildItems(db)
	buildOrders(db)
}

// Build the users table
func buildUsers(db *sql.DB) {
	stm := `CREATE TABLE IF NOT EXISTS users (
		customer_id SERIAL PRIMARY KEY,
		age INT,
		first_name varchar(80),
		last_name varchar(80),
		email varchar(80) UNIQUE NOT NULL
		);`
	_, err := db.Exec(stm)
	if err != nil {
		panic(err)
	}
}

// Build The Items Table
func buildItems(db *sql.DB) {
	stm := `CREATE TABLE IF NOT EXISTS items (
		item_id SERIAL PRIMARY KEY,
		name varchar(80),
		price INT
		);`
	_, err := db.Exec(stm)
	if err != nil {
		panic(err)
	}
}

// Build The Orders Table
func buildOrders(db *sql.DB) {
	stm := `CREATE TABLE IF NOT EXISTS orders (
		order_id SERIAL PRIMARY KEY,
		item_id INT,
		order_group_id INT,
		customer_id varchar(80),
		CONSTRAINT fk_customer
		FORIGN KEY(customer_id)
		REFERENCES users(customer_id)
		);`
	_, err := db.Exec(stm)
	if err != nil {
		panic(err)
	}
}
