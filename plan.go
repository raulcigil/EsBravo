package main

import (
	"log"
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
type TrainningPeriod struct {
	PeriodType       TrainningPeriodType
	DurationMin      int
	DurationMax      int
	Duration         int
	TrainningPurpose string
	TrainningDetails string
}

type TrainningPeriodType int

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	Undefined TrainningPeriodType = iota
	Preparation
	EarlyBase
	LateBase
	Build
	Peak
	Race
	Transition
	RestRecovery
)

type Microcycle struct {
	MesocycleType    TrainningPeriodType
	MicrocycleNumber int
	StartDate        time.Time
	EndDate          time.Time
}

func (mt TrainningPeriod) String() string {
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
	err = CalculatingMicrocyclesTypes(plan)
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

	// //Ordenar por fechas
	// sort.Slice(plan.Microcycles, func(i, j int) bool {
	// 	return plan.Microcycles[i].StartDate.Before(plan.Microcycles[j].StartDate)
	// })

	// //Modificar el número de microciclo
	// mcCount := len(plan.Microcycles)
	// for i, v := range plan.Microcycles {
	// 	plan.Microcycles[i].MicrocycleNumber = mcCount - v.MicrocycleNumber - 1
	// }

	return nil
}

func CalculatingMicrocyclesTypes(plan *Plan) error {
	var trainning []TrainningPeriod
	trainning, err := GetNewAnnualTrainningPeriod()
	if err != nil {
		return err
	}
	var currentPeriod int = len(trainning) - 1
	var currentMicrocycle int = 0
	var currentDuration int = 0

	//Recorremos los periodos en orden inverso desde la carrera hasta hoy
	for currentPeriod < len(trainning) {
		//En cada periodo recorremos las semanas del periodo
		for currentDuration < trainning[currentPeriod].Duration {
			//Comprobación de que no quedan semanas
			if currentMicrocycle < len(plan.Microcycles) {
				//Asigamos el tipo de periodo al microcyclo y pasamos a la siguiente semana
				plan.Microcycles[currentMicrocycle].MesocycleType = trainning[currentPeriod].PeriodType
				currentMicrocycle++
				currentDuration++

				if (trainning[currentPeriod].PeriodType == Peak) ||
					(trainning[currentPeriod].PeriodType == Build) ||
					(trainning[currentPeriod].PeriodType == LateBase) {
					//Cada dos semanas metemos una de recuperación
					if (currentDuration+1)%3 == 0 {
						plan.Microcycles[currentMicrocycle].MesocycleType = RestRecovery
						currentMicrocycle++
						//La contamos como que es del mismo periodo (2Build + RyR + 2Build +RyR)=6Semanas Min
						currentDuration++
					}
				}
			} else {
				break
			}
		}
		//Comprobación de que no quedan semanas
		if currentMicrocycle < len(plan.Microcycles) {
			//Si es pico de rendimiento, entonces una semana de RyR y la ultima semana no es de RyR
			if trainning[currentPeriod].PeriodType == Peak && plan.Microcycles[currentMicrocycle-1].MesocycleType != RestRecovery {
				plan.Microcycles[currentMicrocycle].MesocycleType = RestRecovery
				currentMicrocycle++
			}
			if trainning[currentPeriod].PeriodType == Build && plan.Microcycles[currentMicrocycle-1].MesocycleType != RestRecovery {
				plan.Microcycles[currentMicrocycle].MesocycleType = RestRecovery
				currentMicrocycle++
			}
			if trainning[currentPeriod].PeriodType == LateBase && plan.Microcycles[currentMicrocycle-1].MesocycleType != RestRecovery {
				plan.Microcycles[currentMicrocycle].MesocycleType = RestRecovery
				currentMicrocycle++
			}
		} else {
			break
		}
		//Pasamos de periodo
		currentPeriod--
		//Ponemos la cantidad de semanas a cero
		currentDuration = 0
	}

	return nil
}

func GetNewAnnualTrainningPeriod() ([]TrainningPeriod, error) {
	var trainning [6]TrainningPeriod
	trainning[0] = TrainningPeriod{Preparation, 1, 6, 1, "General Fitness", "Gradually reestablish a structured training routine."}
	trainning[1] = TrainningPeriod{EarlyBase, 3, 6, 3, "General Fitness", "Maximize strength and begin developing aerobic fitness."}
	trainning[2] = TrainningPeriod{LateBase, 3, 6, 3, "General Fitness", "Increase high-intensity training and maintain strength."}
	trainning[3] = TrainningPeriod{Build, 6, 9, 6, "Specific Fitness", "Train with racelike workouts and maintain strength."}
	trainning[4] = TrainningPeriod{Peak, 1, 2, 1, "Specific Fitness", "Taper and simulate small portions of the goal event and maintain strength."}
	trainning[5] = TrainningPeriod{Race, 1, 1, 1, "Specific Fitness", "Race week. Rest before the event while maintaining race fitness."}
	//Periodo de transición (después de carrera)
	//trainning[6] = TrainningPeriod{Transition, 1, 1, 1, "Specific Fitness", "Recover mentally and physically from the stress of serious training."}
	return trainning[:], nil
}

func searchPeriod() {
	// idx := slices.IndexFunc(myconfig, func(c Config) bool { return c.Key == "key1" })
	// return idx
}
