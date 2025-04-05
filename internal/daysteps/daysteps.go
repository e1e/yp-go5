package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"yp-go5/internal/personaldata"
	"yp-go5/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	sData := strings.Split(datastring, ",")
	if len(sData) != 2 {
		return errors.New("invalid input data")
	}

	steps, err := strconv.Atoi(sData[0])
	if err != nil {
		return
	}
	if steps <= 0 {
		return fmt.Errorf("count steps %d < 0", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(sData[1])
	if err != nil {
		return
	}
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("duration < 0")
	}

	distance := spentenergy.Distance(ds.Steps)

	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила: %.2f км.
Вы сожгли: %.2f.`, ds.Steps, distance, spentCalories)

	return output, nil
}
