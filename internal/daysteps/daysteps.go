package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Cornpop456/go-5-sprint-final/internal/personaldata"
	"github.com/Cornpop456/go-5-sprint-final/internal/spentenergy"
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
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return errors.New("wrong format")
	}

	steps, err := strconv.Atoi(parts[0])

	if err != nil {
		return err
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])

	if err != nil {
		return err
	}

	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() string {
	distInKilometers := (float64(ds.Steps) * StepLength) / 1000
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	return fmt.Sprintf(`Количество шагов: %d.
Дистанция: %.2f км.
Вы сожгли %.2f ккал.`,
		ds.Steps,
		distInKilometers,
		calories)
}
