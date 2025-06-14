package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        DROP TABLE IF EXISTS users;
        DROP TABLE IF EXISTS travels;
        DROP TABLE IF EXISTS savedTravels;
        CREATE TABLE IF NOT EXISTS users (
            email TEXT PRIMARY KEY,
            password TEXT NOT NULL
        );
        CREATE TABLE IF NOT EXISTS travels (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            cities TEXT,
            totalCost INTEGER,
            duration INTEGER,
            departure TEXT,
            arrival TEXT
        );
        CREATE TABLE IF NOT EXISTS savedTravels (
            tid INTEGER,
            uid TEXT
        );
        CREATE TABLE IF NOT EXISTS routes (
            fromCity TEXT,
            toCity TEXT,
            departure INTEGER,
            arrival INTEGER,
            duration INTEGER,
            price INTEGER
        );
    `)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
