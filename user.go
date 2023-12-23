package main

import (
	"encoding/json"
	"log"
	"os"
)

// User data
type User struct {
	Id   int
	Name string
}

func (user *User) LoadData() {
	ok := CheckFileExists(c_userFilePath)
	if !ok {
		f, err := os.Create(c_userFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	} else {
		data, err := os.ReadFile(c_userFilePath)
		if err != nil {
			log.Fatal(err)
		} else {
			json.Unmarshal(data, &user)
		}
	}
}

func (user *User) ToJson() ([]byte, error) {
	return json.Marshal(user)
}

func (user *User) FromJson(data []byte) error {
	json.Unmarshal(data, &user)
	return nil
}
