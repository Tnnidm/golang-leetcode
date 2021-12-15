package dayOfTheWeek

func dayOfTheWeek(day int, month int, year int) string {
	res := []string{"Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
	days := -1
	for i := 1971; i < year; i++ {
		if isLeapYear(i) {
			days += 366
		} else {
			days += 365
		}
	}
	for i := 1; i < month; i++ {
		switch i {
		case 2:
			if isLeapYear(year) {
				days += 29
			} else {
				days += 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			days += 31
		default:
			days += 30
		}
	}
	days += day
	return res[days%7]
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%4 == 0 && year%100 != 0 {
		return true
	}
	return false
}
