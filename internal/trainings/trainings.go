package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"yp-go5/internal/personaldata"
	"yp-go5/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	sData := strings.Split(datastring, ",")
	if len(sData) != 3 {
		return errors.New("invalid input data")
	}

	steps, err := strconv.Atoi(sData[0])
	if err != nil {
		return
	}
	t.Steps = steps

	if sData[1] != "Бег" && sData[1] != "Ходьба" {
		return errors.New("invalid input data")
	}
	t.TrainingType = sData[1]

	duration, err := time.ParseDuration(sData[2])
	if err != nil {
		return
	}
	t.Duration = duration

	return
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)

	if t.Duration <= 0 {
		return "", errors.New("duration < 0")
	}

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var spentCalories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories)

	return output, nil
}
