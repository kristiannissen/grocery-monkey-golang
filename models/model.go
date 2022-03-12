package models

import (
    "os"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

type (
  Model struct {}
)

var (
    DB *sql.DB
)

// Is loaded when any model is loaded
func init() {
    var err error
    DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }
}
