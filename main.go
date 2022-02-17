package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Kitty!")
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
