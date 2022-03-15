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

// TODO: move to model.go
func (m *Model) UserSetUp() {
	_, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS users (users_id serial PRIMARY KEY, nickname varchar(255), uuid varchar(255), created_at TIMESTAMP)")
	if err != nil {
		log.Fatalf("Create statement %q", err)
	}
}

// TODO: move to model.go
func (m *Model) UserTearDown() {
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		log.Fatalf("Drop statement %q", err)
	}
}

func (m *Model) CleanUserTable() {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatalf("Delete statement %q", err)
	}
}
