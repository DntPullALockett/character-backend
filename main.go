package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Collection struct {
	gorm.Model
	Name string `json:"name" bson:"Name"`
}

type Character struct {
	gorm.Model
	Name                string     `json:"name" bson:"Name"`
	CurrentlyObtainable bool       `json:"currentlyObtainable" bson:"CurrentlyObtainable"`
	Premium             bool       `json:"premium" bson:"Premium"`
	LimitedTime         bool       `json:"limitedTime" bson:"LimitedTime"`
	MaxLevel            int        `json:"maxLevel" bson:"MaxLevel"`
	Collection          Collection `json:"collection" bson:"Colleciton"`
}

func main() {
	connect()
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

func connect() {
	connectionStr := os.Getenv("DATABASE_URL")
	postgres, err := gorm.Open(postgres.Open(connectionStr))
	if err != nil {
		log.Fatal("could not connect to dmkt database")
	}

	db = postgres
}

func createCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var character Character
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &character)
	if err != nil {
		fmt.Println("error reading body")
	}

}

func createCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var collection Collection

	err := json.NewDecoder(r.Body).Decode(&collection)

	if err != nil {
		fmt.Println("Error reading body")
	}

	db.Create(collection)
}

func getAllCharacters() {

}
