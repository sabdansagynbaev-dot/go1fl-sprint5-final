package daysteps
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)
type DaySteps struct {
	// TODO: добавить поля
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
    if len(parts) != 2 {
        return errors.New("invalid data string format")
    }

    steps, err := strconv.Atoi(parts[0])
if err != nil {
 return err
}
if steps <= 0 {
 return errors.New("invalid steps")
}

duration, err := time.ParseDuration(parts[1])
if err != nil {
 return err
}
if duration <= 0 {
 return errors.New("invalid duration")
}

ds.Steps = steps
ds.Duration = duration
return nil
}
func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

    calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
    if err != nil {
        return "", err
    }

    info := fmt.Sprintf(
        "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
        ds.Steps, distance, calories,
    )

    return info, nil
}
