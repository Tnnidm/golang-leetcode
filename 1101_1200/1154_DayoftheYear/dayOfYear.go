package dayOfYear

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	ymd := strings.Split(date, "-")
	year, _ := strconv.Atoi(ymd[0])
	month, _ := strconv.Atoi(ymd[1])
	day, _ := strconv.Atoi(ymd[2])

	result := 0
	for i := 1; i <= month-1; i++ {
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			result = result + 31
		case 4, 6, 9, 11:
			result = result + 30
		case 2:
			if year%4 == 0 {
				result = result + 29
			} else {
				result = result + 28
			}
		}
	}
	result = result + day
	return result
}
