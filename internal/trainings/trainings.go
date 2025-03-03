package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Cornpop456/go-5-sprint-final/internal/personaldata"
	"github.com/Cornpop456/go-5-sprint-final/internal/spentenergy"
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
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return errors.New("wrong format")
	}

	steps, err := strconv.Atoi(parts[0])

	if err != nil {
		return err
	}

	t.Steps = steps

	training := parts[1]

	switch training {
	case "Бег", "Ходьба":
		t.TrainingType = training
	default:
		return errors.New("unknown training type")
	}

	duration, err := time.ParseDuration(parts[2])

	if err != nil {
		return err
	}

	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() string {
	distance := spentenergy.Distance(t.Steps)
	calories := 0.0

	switch t.TrainingType {
	case "Бег":
		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "неизвестный тип тренировки"
	}

	return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`,
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		spentenergy.MeanSpeed(t.Steps, t.Duration),
		calories,
	)
}
