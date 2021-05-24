package main

import (
	"analisis/internal/student"
	"analisis/internal/bestSchedule"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFile, err := os.Open("respuestas.csv")

	if err != nil {
		log.Panic(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		log.Panic(err)
	}

	var students []*student.Student

	for _, line := range csvLines {
		email := line[1]
		name := line[2]

		dayNames := strings.Split(line[5], ",")
		schedules := []string{line[6], line[7], line[8], line[9], line[10]}

		s := student.NewStudent(email, name, dayNames, schedules)

		students = append(students, s)
	}

	// Crear bestSchedule list
	bs := bestSchedule.NewBestSchedule()

	// Iteramos los students
	for _, student := range students {
		// x cada student obtenemos los days
		for _, day := range student.Days {
			
			dayName := (*day).Name
			if (*bs).Days[dayName] == nil {
				(*bs).Days[dayName] = bestSchedule.NewDay()
			}
			
			for key, _ := range day.Schedule {
				// Agrego el horario de c/dia/estudiante a la lista de horarios generales.
				sch := (*((*bs).Days[dayName])).Schedule[key]
				sch = append(sch, student.Email)
				(*((*bs).Days[dayName])).Schedule[key] = sch
			}
		}
	}

	// TODO: ANALIZAR BS => Hay que buscar los horarios con mayor número de coincidencias.

	fmt.Println(*bs)
}
