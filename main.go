package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
	CreatedAt           time.Time  `gorm:"column:createdAt"`
	UpdatedAt           time.Time  `gorm:"column:updatedAt"`
	DeletedAt           time.Time  `gorm:"column:deletedAt"`
}

func main() {
	connect()
	log.SetPrefix("character service: ")
	log.SetFlags(0)

	//Character Handlers
	http.HandleFunc("POST /characters", createCharacterHandler)

	//Collection Handlers
	http.HandleFunc("GET /collections", getCollectionsHandler)
	http.HandleFunc("POST /collections", createCollectionHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
	getAllCharacters()
}

func connect() {
	connectionStr := os.Getenv("DATABASE_URL")
	postgres, err := gorm.Open(postgres.Open(connectionStr))
	if err != nil {
		log.Fatal("could not connect to dmkt database")
	}

	db = postgres
	fmt.Println("Connected to DMKT database!")
	db.AutoMigrate(&Collection{})
}

func createCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var character Character
	reqBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &character)
	if err != nil {
		fmt.Println("error reading body")
	}
}

func getCollectionsHandler(w http.ResponseWriter, r *http.Request) {
	getAllCharacters()
}

func createCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var collection Collection

	err := json.NewDecoder(r.Body).Decode(&collection)

	if err != nil {
		fmt.Println("Error reading body")
	}

	db.Create(&collection)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(collection)
}

func getAllCharacters() []Collection {
	var collections []Collection
	result := db.Find(&collections)
	if result.Error != nil {
		fmt.Println("could not obtain collections")
	}

	for row := range result.RowsAffected {
		fmt.Println(collections[row])
	}

	return collections
}
