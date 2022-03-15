package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type (
	User struct {
		NickName string `json:"nickname"`
		Uuid     string `json:"uuid"`
	}
)

// GetUser finds a user based on nickname
// and either returns a User or an error ir nothing was found
func (m *Model) GetUser(nickname string) (*User, error) {
	row := db.QueryRow(
		"SELECT nickname, uuid FROM users WHERE nickname = $1 LIMIT 1", nickname)

	user := new(User)

	err := row.Scan(&user.NickName, &user.Uuid)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("User not found"))
	}

	return user, nil
}

func (m *Model) CreateUser(nickname string) (*User, error) {
	user := m.NewUser()
	user.NickName = nickname

	_, err := db.Exec(
		"INSERT INTO users (nickname, uuid) VALUES ($1, $2)", user.NickName, user.Uuid)
	if err != nil {
		log.Fatalf("Insert statement error %q", err)
		return nil, err
	}

	return user, nil
}

func (m *Model) NewUser() *User {
	user := new(User)
	user.Uuid = uuid.New().String()

	return user
}
