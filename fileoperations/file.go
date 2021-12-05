package fileoperations

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadKeys(db *map[string]string) {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	json.Unmarshal(data, &db)
}
