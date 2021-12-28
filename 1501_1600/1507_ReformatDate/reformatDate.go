package reformatDate

import "strings"

func reformatDate(date string) string {
	dateList := strings.Fields(date)
	result := strings.Builder{}
	result.WriteString(dateList[2])
	result.WriteByte('-')
	switch dateList[1] {
	case "Jan":
		result.WriteString("01")
	case "Feb":
		result.WriteString("02")
	case "Mar":
		result.WriteString("03")
	case "Apr":
		result.WriteString("04")
	case "May":
		result.WriteString("05")
	case "Jun":
		result.WriteString("06")
	case "Jul":
		result.WriteString("07")
	case "Aug":
		result.WriteString("08")
	case "Sep":
		result.WriteString("09")
	case "Oct":
		result.WriteString("10")
	case "Nov":
		result.WriteString("11")
	case "Dec":
		result.WriteString("12")
	}
	result.WriteByte('-')
	if len(dateList[0]) == 3 {
		result.WriteByte('0')
		result.WriteByte(dateList[0][0])
	} else {
		result.WriteString(dateList[0][0:2])
	}
	return result.String()
}
