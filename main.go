package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo"
)

// Path al fichero de datos de usuario
const c_userFilePath string = "data-user.json"

// Path al fichero de datos de usuario
const c_DataBaseFilePath string = "database.json"

// Datos del usuario
var DBData DataBase

// Aqui empieza todo
func main() {
	e := echo.New()
	RegisterRoutes(e)
	RegisterTemplates(e)
	e.Logger.Fatal(e.Start(":8080"))
	return
	//return c.Redirect(http.StatusMovedPermanently, "<URL>")
	LoadData()
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

// Comprobar que se ha configurado la cuenta de usuario.
func CheckUserData() bool {
	if DBData.User.Id == 0 {
		return true
	}
	if IsEmpty(DBData.User.Name) {
		return true
	}
	return false
}

func CheckPlans() {
	if DBData.Plan.Distance == 0 {
		raceday, _ := time.Parse(time.DateOnly, "2024-03-24")
		DBData.Plan = Plan{1, "Granadella", 26, 150, true, raceday, []Microcycle{}}
	}
}
