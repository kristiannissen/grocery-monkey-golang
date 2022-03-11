package main

import (
  "fmt"
  "github.com/kristiannissen/grocery-monkey-golang/models"
)

func main() {
  m := models.Model{}
	fmt.Println(m.GetUser("Kitty"))
}
