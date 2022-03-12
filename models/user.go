package models

import (
    "log"
    "database/sql"
)

type (
    User struct {
        NickName string
        Uuid string
    }
)

// The specific model should return something or an error
func (m *Model) GetUser(nickname string) *User {
    var err error

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
    defer DB.Close()

    row := DB.QueryRow("SELECT nickname, uuid FROM user WHERE nickname = $1", nickname)
    user := new(User)

    err = row.Scan(&user.NickName, &user.Uuid)
    if err == sql.ErrNoRows {
        log.Fatal(err)
        return user
    }
    return user
}
