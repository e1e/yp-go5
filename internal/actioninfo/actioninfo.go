package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		err := dp.Parse(datastring)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		output, err := dp.ActionInfo()
		if err != nil {
			continue
		}
		fmt.Println(output)
	}
}
