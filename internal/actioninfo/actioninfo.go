package actioninfo
import "fmt"
type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
  if err := dp.Parse(data); err != nil {
   continue
  }

  info, err := dp.ActionInfo()
  if err != nil {
   continue
  }

  fmt.Print(info) // важно: Print, не Println (у info уже есть \n в конце)
 }
}
