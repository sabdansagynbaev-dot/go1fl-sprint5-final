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
	Weight float64
	Height float64

	Personal personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")

 if len(parts) != 2 && len(parts) != 4 {
  return errors.New("invalid data string format")
 }

 for _, p := range parts {
  if strings.TrimSpace(p) != p {
   return errors.New("invalid data string format")
  }
 }

 steps, err := strconv.Atoi(parts[0])
 if err != nil || steps <= 0 {
  return errors.New("invalid steps")
 }

 duration, err := time.ParseDuration(parts[1])
 if err != nil || duration <= 0 {
  return errors.New("invalid duration")
 }

 ds.Steps = steps
 ds.Duration = duration

 if len(parts) == 4 {
  weightInt, err := strconv.Atoi(parts[2])
  if err != nil || weightInt <= 0 {
   return errors.New("invalid weight")
  }

  heightInt, err := strconv.Atoi(parts[3])
  if err != nil || heightInt <= 0 {
   return errors.New("invalid height")
  }

  ds.Weight = float64(weightInt)
  ds.Height = float64(heightInt)
  ds.Personal.Weight = ds.Weight
  ds.Personal.Height = ds.Height
 }

 return nil
}
func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	 distance := spentenergy.Distance(ds.Steps, float64(ds.Personal.Height))

 calories, err := spentenergy.WalkingSpentCalories(
  ds.Steps,
  float64(ds.Personal.Weight),
  float64(ds.Personal.Height),
  ds.Duration,
 )
 if err != nil {
  return "", err
 }

 info := fmt.Sprintf(
  "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
  ds.Steps, distance, calories,
 )
 return info, nil
}

func (ds DaySteps) Print() {
 info, err := ds.ActionInfo()
 if err != nil {
  fmt.Println(err)
  return
 }
 fmt.Print(info)
}
