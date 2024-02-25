package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetPrefix("character service: ")
	log.SetFlags(0)

	http.HandleFunc("POST /characters/create", createCharacterHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func createCharacterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("character created")
}
