package models

import (
    "log"
    "database/sql"
    "github.com/google/uuid"
)

type (
    User struct {
        NickName string
        Uuid string
    }
)

func (m *Model) SetUp() {
    _, err := DB.Exec("CREATE TABLE IF NOT EXISTS users (users_id serial PRIMARY KEY, nickname varchar(255), uuid varchar(255), created_at TIMESTAMP)")
    if err != nil {
        log.Fatalf("Create statement %q", err)
    }
}

func (m *Model) TearDown() {
    _, err := DB.Exec("DROP TABLE IF EXISTS users")
    if err != nil {
        log.Fatalf("Drop statement %q", err)
    }
}

func (m *Model) GetUser(nickname string) *User {
    var err error

    if err = DB.Ping(); err != nil {
        log.Fatalf("User %q", err)
    }

    row := DB.QueryRow("SELECT nickname, uuid FROM users WHERE nickname = $1 LIMIT 1", nickname)

    user := new(User)

    err = row.Scan(&user.NickName, &user.Uuid)
    if err == sql.ErrNoRows {
        log.Fatalf("Select statement error %q", err)
        return user
    }
    return user
}

func (m *Model) CreateUser(nickname string) *User {
    user := new(User)
    user.NickName = nickname
    user.Uuid = uuid.New().String()

    res, err := DB.Exec("INSERT INTO users (nickname, uuid) VALUES ($1, $2)", user.NickName, user.Uuid)
    if err != nil {
        log.Fatalf("Insert statement error %q", err)
        return nil
    }

    rows, err := res.RowsAffected()
    if err != nil {
        log.Fatal(rows)
        return nil
    }

    return user
}
