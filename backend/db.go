package main

import (
  "database/sql"
  "log"
)

func initDB() *sql.DB {
  db, err := sql.Open("sqlite3", "./app.db")
  if err != nil {
    log.Fatal(err)
  }
  return db
}
