package student

import "strings"

type Student struct {
	Email string
	Name  string
	Days  []*Day
}

type Day struct {
	Name     string
	Schedule map[string]bool
}

func newDay(name string, schedulesStr string) *Day {
	schedules := strings.Split(schedulesStr, ",")
	m := make(map[string]bool)

	for _, schedule := range schedules {
		scheduleSanit := strings.TrimSpace(schedule)
		if scheduleSanit != "" {
			m[scheduleSanit] = true
		}
	}

	day := Day{
		Name:     name,
		Schedule: m,
	}

	return &day
}

func newDays(dayNames []string, daySchedules []string) []*Day {
	var days []*Day

	for _, dayName := range dayNames {
		dayName = strings.TrimSpace(dayName)

		var index int

		switch dayName {
		case "Lunes":
			index = 0
		case "Martes":
			index = 1
		case "Miércoles":
			index = 2
		case "Jueves":
			index = 3
		case "Viernes":
			index = 4
		}

		schedule := daySchedules[index]
		if schedule != "" {
			day := newDay(dayName, schedule)
			days = append(days, day)
		}
	}

	return days
}

func NewStudent(email string, name string, dayNames []string, daySchedules []string) *Student {
	return &Student{
		Email: email,
		Name:  name,
		Days:  newDays(dayNames, daySchedules),
	}
}
