package models

import (
    "log"
    "github.com/google/uuid"
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
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS grocerylist (grocerylist_id serial PRIMARY KEY, content jsonb, user_uid VARCHAR(255) NOT NULL)") 
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

func (m *Model) CreateGroceryList(g *GroceryList) *GroceryList {
    groceryList := g

   return groceryList
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
