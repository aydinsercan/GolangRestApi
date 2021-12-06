package store

import (
	"encoding/json"
	fileoperations "go-api/fileoperations"
	"log"
	"os"
	"strconv"
	"time"
)

type Store struct {
	Db map[string]string
}

func New() Store {
	var db = map[string]string{}
	//fmt.Printf("%T\n", db)
	var store = Store{Db: db}

	fileoperations.LoadKeys(&db)

	defer func() {
		go Timer(store)
	}()

	return store
}

// Get Item
func (k Store) Get(key string) string {
	return k.Db[key]
}

// Add Item
func (k Store) Add(value string) string {
	count := len(k.Db)
	count += 1

	newId := strconv.Itoa(count)

	k.Db[newId] = value
	return newId
}

// Flush Items
func (k Store) Flush() {
	for k2 := range k.Db {
		delete(k.Db, k2)
	}
}

//Save result to the file at a specified interval (every 5 seconds)
func Timer(store Store) {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		log.Println("")

		dbJson, _ := json.Marshal(store.Db)
		os.WriteFile("data.json", dbJson, os.ModeAppend)
	}
}
