package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Collection struct {
	Name string `json:"name" bson:"Name"`
}

func main() {
	log.SetPrefix("character service: ")
	log.SetFlags(0)

	http.HandleFunc("POST /characters/create", createCharacterHandler)
	http.HandleFunc("POST /collections/create", createCollectionHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func createCharacterHandler(w http.ResponseWriter, r *http.Request) {

}

func createCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var collection Collection

	err := json.NewDecoder(r.Body).Decode(&collection)

	if err != nil {
		fmt.Println("Error reading body")
	}

	fmt.Println(collection.Name)
}
