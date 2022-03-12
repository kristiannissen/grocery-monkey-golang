package models

import (
    "log"
    "os"
    // "database/sql"
    // _ "github.com/lib/pq"
)

type (
  Model struct {}
)

var (
    DB string
)

// Is loaded when any model is loaded
func init() {
    log.Println("Hello from models.init()")
    DB = os.Getenv("DATABASE_URL") 
}
