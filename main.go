package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Collection struct {
	Name string `json:"name" bson:"Name"`
}

type Character struct {
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
	postgres, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal("could not connect to postgres")
	}

	defer postgres.Close()
	err = postgres.Ping()
	if err != nil {
		panic(err)
	}

	db = postgres
	fmt.Println("Connection Successful!")
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

	var collectionId int
	rowErr := db.QueryRow(`INSERT INTO collections(name) VALUES(collection.Name) RETURNING id`).Scan(&collectionId)
	if rowErr != nil {
		fmt.Println("error inserting collection")
	}
	fmt.Println(collectionId)
}
