package bestSchedule

type BestSchedule struct {
	Days map[string]*Day
}

type Day struct {
	Schedule map[string][]string
}

func NewDay() *Day {
	m := make(map[string][]string)

	return &Day {
		Schedule: m,
	}
}

func NewBestSchedule() *BestSchedule {
	m := make(map[string]*Day)
	return &BestSchedule{
		Days: m,
	}
}
