package models

import (
    "log"
    "github.com/google/uuid"
    "encoding/json"
)

type (
    Grocery struct {
        Name string
        Id string
    }

    GroceryList struct {
        Uuid string
        Subscribers []string
        Groceries []Grocery
        UserUuid string
    }
)

func (m *Model) GroceryListSetUp() {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS grocerylist (grocerylist_id serial PRIMARY KEY, content jsonb, user_uuid VARCHAR(255) NOT NULL)") 
    if err != nil {
        log.Fatalf("Create statement %q", err)
    }
}

func (m *Model) GroceryListTearDown() {
    _, err := db.Exec("DROP TABLE IF EXISTS grocerylist")
    if err != nil {
        log.Fatalf("Drop statement %q", err)
    }
}

func (m *Model) GetGroceryList() *GroceryList {
    groceryList := new(GroceryList)

   return groceryList
}

func (m *Model) CreateGroceryList(g *GroceryList) (*GroceryList, error) {
    // Encode struct to string
    str, err := json.Marshal(g)
    if err != nil {
        return nil, err
    }
    // Store grocerylist
    _, err = db.Exec("INSERT INTO grocerylist (content, user_uuid) VALUES ($1, $2)", str, g.UserUuid)
    if err != nil {
        return nil, err
    }

   return g, nil
}

func (m *Model) UpdateGroceryList() *GroceryList {
    groceryList := new(GroceryList)

   return groceryList
}

func (m *Model) NewGroceryList() *GroceryList {
    groceryList := new(GroceryList)
    groceryList.Uuid = uuid.New().String()

    return groceryList
}
