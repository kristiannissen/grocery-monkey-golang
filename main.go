package main

import (
  "fmt"
  "github.com/kristiannissen/grocery-monkey-golang/models"
    "os"
)

func main() {
  m := models.Model{}
	fmt.Println(m.GetUser("Kitty") + os.Getenv("DATABASE_URL"))
}
