// utils project utils.go
package utils

import (
	"log"
	"strconv"
)

func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetInteger(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		result = 0
		LogErr(err)
	}
	return result
}
