package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Path al fichero de datos de usuario
const c_userFilePath string = "data-user.json"

// Path al fichero de datos de usuario
const c_DataBaseFilePath string = "database.json"

// Datos del usuario
var DBData DataBase

// Aqui empieza todo
func main() {
	ClearConsole()
	LoadData()
	CheckUserData()
	CheckPlans()
	CalculatingMicrocycles(&DBData.Plan)
	SaveData()
}

func LoadData() {
	DBData.LoadData()
	//DEBUG, limpiar la lista
	DBData.Plan.Microcycles = make([]Microcycle, 0, 10)
}

func SaveData() {
	f, err := os.OpenFile(c_DataBaseFilePath, os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	} else {
		data, _ := DBData.ToJson()
		f.Write(data)
	}
}

func CheckUserData() {
	DBData.User.Name = "Raúl"
	return
	if IsEmpty(DBData.User.Name) {
		var name string
		fmt.Println(("¿What's your name?"))
		fmt.Scanln(&name)
		DBData.User.Name = name

	} else {
		fmt.Printf(" Hola %s!\n", DBData.User.Name)
	}
}
func CheckPlans() {
	if DBData.Plan.Distance == 0 {
		raceday, _ := time.Parse(time.DateOnly, "2024-03-24")
		DBData.Plan = Plan{26, 150, true, raceday, []Microcycle{}}
	}
}
