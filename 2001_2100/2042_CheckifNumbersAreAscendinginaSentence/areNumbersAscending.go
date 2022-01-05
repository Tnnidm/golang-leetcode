package areNumbersAscending

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	tokens := strings.Fields(s)
	lastNum := -1
	for _, token := range tokens {
		num, error := strconv.Atoi(token)
		if error == nil {
			if num > lastNum {
				lastNum = num
			} else {
				return false
			}
		}
	}
	return true
}
