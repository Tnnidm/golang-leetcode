package daysBetweenDates

import (
	"strconv"
	"strings"
)

func daysBetweenDates(date1 string, date2 string) int {
	if date1 == date2 {
		return 0
	}
	date1Split := strings.Split(date1, "-")
	date2Split := strings.Split(date2, "-")
	date1Year, _ := strconv.Atoi(date1Split[0])
	date1Month, _ := strconv.Atoi(date1Split[1])
	date1Day, _ := strconv.Atoi(date1Split[2])
	date2Year, _ := strconv.Atoi(date2Split[0])
	date2Month, _ := strconv.Atoi(date2Split[1])
	date2Day, _ := strconv.Atoi(date2Split[2])
	days1 := daysTo1971(date1Year, date1Month, date1Day)
	days2 := daysTo1971(date2Year, date2Month, date2Day)
	if days1 > days2 {
		return days1 - days2
	} else {
		return days2 - days1
	}
}

func daysTo1971(year, month, day int) int {
	days := 0
	for i := 1971; i < year; i++ {
		if (i%4 == 0 && i%100 != 0) || i%400 == 0 {
			days = days + 366
		} else {
			days = days + 365
		}
	}
	for i := 1; i < month; i++ {
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			days = days + 31
		case 4, 6, 9, 11:
			days = days + 30
		case 2:
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				days = days + 29
			} else {
				days = days + 28
			}
		}
	}
	days += day
	return days
}
