package models

import (
	_ "database/sql/driver"
	"encoding/json"
	_ "errors"
	"github.com/google/uuid"
	"log"
)

type (
	Grocery struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}

	GroceryList struct {
		Uuid        string    `json:"uuid" query:"uuid"`
		Subscribers []string  `json:"subscribers"`
		Groceries   []Grocery `json:"groceries"`
		UserUuid    string    `json:"useruuid"`
	}
)

func (m *Model) GetGroceryList(uuid string) (*GroceryList, error) {
	g := new(GroceryList)
	var groceries string
	// TODO: should be finding based on subscribers
	row := db.QueryRow("SELECT groceries FROM grocerylist WHERE uuid = $1", uuid)
	err := row.Scan(&groceries)
	if err != nil {
		log.Printf("Select error %s", err)
		return nil, err
	}
	json.Unmarshal([]byte(groceries), &g)

	return g, nil
}

func (m *Model) CreateGroceryList(g *GroceryList) (*GroceryList, error) {
	// Encode struct to string
	str, err := json.Marshal(g)
	if err != nil {
		log.Printf("Marshal error %s", err)
		return nil, err
	}
	// Store grocerylist
	_, err = db.Exec(
		"INSERT INTO grocerylist (groceries, user_uuid, uuid) VALUES ($1, $2, $3)", str, g.UserUuid, g.Uuid)
	if err != nil {
		log.Printf("Insert error %s", err)
		return nil, err
	}

	return m.GetGroceryList(g.Uuid)
}

func (m *Model) UpdateGroceryList(g *GroceryList) (*GroceryList, error) {
	// Encode the struct to a string
	str, err := json.Marshal(g)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// log.Printf("str %s", str)
	// Update the grocerylist
	// TODO: Should use subscriber instead
	_, err = db.Exec("UPDATE grocerylist SET groceries = $1 WHERE uuid = $2", str, g.Uuid)
	if err != nil {
		log.Printf("Update error %s", err)
		return nil, err
	}

	return m.GetGroceryList(g.Uuid)
}

func (m *Model) NewGroceryList() *GroceryList {
	groceryList := new(GroceryList)
	groceryList.Uuid = uuid.New().String()
	groceryList.Groceries = []Grocery{}
	groceryList.Subscribers = []string{}

	return groceryList
}
