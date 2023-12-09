package main

import (
	"log"
	"sort"
	"time"
)

// Trainning plan
type Plan struct {
	Distance      float32
	EstimatedTime int // In minutes
	Active        bool
	RaceDay       time.Time
	Microcycles   []Microcycle
}
type MesocycleType int

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	Undefined MesocycleType = iota
	// Summer
	// Autumn
	// Winter
	// Spring
)

type Microcycle struct {
	MesocycleType    MesocycleType
	MicrocycleNumber int
	StartDate        time.Time
	EndDate          time.Time
}

func (mt MesocycleType) String() string {
	return mt.String()
	// switch mt {
	// case Summer:
	// 	return "ñlk"
	// default:
	// 	return ""
	// }
}

// Calculating a list of microcycles
func CalculatingMicrocycles(plan *Plan) error {
	err := CalculatingMicrocyclesDates(plan, Microcycle{Undefined, 0, time.Now(), plan.RaceDay})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CalculatingMicrocyclesDates(plan *Plan, lastMicro Microcycle) error {
	today := time.Now()
	var targetDay time.Time = lastMicro.EndDate

	// Hay que buscar el lunes si todavía podemos programar microciclos
	if targetDay.After(today) {
		mondayDay := targetDay
		for mondayDay.Weekday() != time.Monday {
			mondayDay = mondayDay.AddDate(0, 0, -1)
		}
		if mondayDay.Before(today) {
			mondayDay = today
		}
		//microcycles := make([]Microcycle, 6)
		thisMicrocycle := Microcycle{Undefined, lastMicro.MicrocycleNumber, mondayDay, targetDay}
		plan.Microcycles = append(plan.Microcycles, thisMicrocycle)

		return CalculatingMicrocyclesDates(plan, Microcycle{Undefined, lastMicro.MicrocycleNumber + 1, time.Now(), mondayDay.AddDate(0, 0, -1)})
	}

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.

	//Ordenar por fechas
	sort.Slice(plan.Microcycles, func(i, j int) bool {
		return plan.Microcycles[i].StartDate.Before(plan.Microcycles[j].StartDate)
	})

	//Modificar el número de microciclo
	mcCount := len(plan.Microcycles)
	for i, v := range plan.Microcycles {
		plan.Microcycles[i].MicrocycleNumber = mcCount - v.MicrocycleNumber - 1
	}

	return nil
}
