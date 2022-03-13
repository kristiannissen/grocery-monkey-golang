package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type (
	Model struct{}
)

var (
	db *sql.DB
)

// Is loaded when any model is loaded
func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Models %q", (err))
	}
}
