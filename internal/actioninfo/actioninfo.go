// Package actioninfo provides a common activity-processing interface.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser is implemented by both workouts and walks.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info parses the dataset and prints information for each valid activity.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print(info)
	}
}
