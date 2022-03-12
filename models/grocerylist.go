package models

import (

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
