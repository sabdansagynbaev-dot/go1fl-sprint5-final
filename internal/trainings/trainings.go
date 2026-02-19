package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data string format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("invalid steps")
	}

	trainingType := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("invalid duration")
	}

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, float64(t.Height))
	meanSpeed := spentenergy.MeanSpeed(t.Steps, float64(t.Height), t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	durationHours := t.Duration.Hours()

	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationHours, distance, meanSpeed, calories,
	)

	return info, nil
}
