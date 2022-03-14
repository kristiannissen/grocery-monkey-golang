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
		Uuid        string    `json:"uuid"`
		Subscribers []string  `json:"subscribers"`
		Groceries   []Grocery `json:"groceries"`
		UserUuid    string    `json:"useruuid"`
	}
)

func (m *Model) GroceryListSetUp() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS grocerylist (grocerylist_id serial PRIMARY KEY, groceries jsonb NOT NULL, user_uuid VARCHAR(255) NOT NULL, created_at TIMESTAMP, uuid VARCHAR(255))")
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

func (m *Model) GetGroceryList(user_uuid string) (*GroceryList, error) {
	g := new(GroceryList)
	var groceries string

	row := db.QueryRow("SELECT groceries FROM grocerylist WHERE user_uuid = $1", user_uuid)
	err := row.Scan(&groceries)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(groceries), &g)

	return g, nil
}

func (m *Model) CreateGroceryList(g *GroceryList) (*GroceryList, error) {
	// Encode struct to string
	str, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	// Store grocerylist
	_, err = db.Exec("INSERT INTO grocerylist (groceries, user_uuid) VALUES ($1, $2)", str, g.UserUuid)
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
	groceryList.Groceries = []Grocery{}

	return groceryList
}
