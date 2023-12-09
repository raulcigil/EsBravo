package main

import (
	"encoding/json"
	"log"
	"os"
)

type DataBase struct {
	User User
	Plan Plan
}

func (database *DataBase) LoadData() {
	ok := CheckFileExists(c_DataBaseFilePath)
	if !ok {
		f, err := os.Create(c_DataBaseFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	} else {
		data, err := os.ReadFile(c_DataBaseFilePath)
		if err != nil {
			log.Fatal(err)
		} else {
			json.Unmarshal(data, &database)
		}
	}
}

func (database *DataBase) ToJson() ([]byte, error) {
	return json.MarshalIndent(database, "", "\t")
}
