// utils.go
package structs

import (
	"game/utils"
	"regexp"
)

func isStringMatches(line, regex string) []string {
	r, err := regexp.Compile(regex)
	utils.LogErr(err)
	result := r.FindStringSubmatch(line)
	return result
}
