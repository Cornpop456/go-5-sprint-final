package actioninfo

import "fmt"

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() string
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		err := dp.Parse(datastring)

		if err != nil {
			continue
		}

		fmt.Println(dp.ActionInfo())
	}
}
