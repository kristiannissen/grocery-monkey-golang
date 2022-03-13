package models

import (
    "log"
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

func GroceryListSetUp(m *Model) {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS grocerylist (grocerylist_id serial PRIMARY KEY, content jsonb, user_uid VARCHAR(255) NOT NULL)") 
    if err != nil {
        log.Fatalf("Create statement %q", err)
    }
}

func GroceryListTearDown(m *Model) {
    _, err := db.Exec("DROP TABLE IF EXISTS grocerylist")
    if err != nil {
        log.Fatalf("Drop statement %q", err)
    }
}

func GetGroceryList(m *Model) *GroceryList {
    groceryList := new(GroceryList)

   return groceryList
}

func CreateGroceryList(m *Model) *GroceryList {
    groceryList := new(GroceryList)

   return groceryList
}

func UpdateGroceryList(m *Model) *GroceryList {
    groceryList := new(GroceryList)

   return groceryList
}
