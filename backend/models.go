package main

import "database/sql"

/* type Date struct {
	day, month, year string
} */

type Config struct {
	db *sql.DB
}

type User struct {
	email, password string
}

type Route struct {
	departure, arrival, cost int
}
