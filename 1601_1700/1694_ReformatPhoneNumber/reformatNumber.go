package reformatNumber

import "strings"

func reformatNumber(number string) string {
	remainNum := 0
	numberNum := len(number)
	for i := 0; i < numberNum; i++ {
		if number[i] >= '0' && number[i] <= '9' {
			remainNum++
		}
	}
	result := strings.Builder{}
	index, counter := 0, 0
	for remainNum > 4 {
		counter = 0
		for ; index < numberNum && counter < 3; index++ {
			if number[index] >= '0' && number[index] <= '9' {
				result.WriteByte(number[index])
				counter++
				remainNum--
			}
		}
		result.WriteByte('-')
	}
	if remainNum == 4 {
		counter = 0
		for ; index < numberNum; index++ {
			if number[index] >= '0' && number[index] <= '9' {
				result.WriteByte(number[index])
				counter++
				remainNum--
				if counter == 2 {
					result.WriteByte('-')
				}
			}
		}
	} else {
		for ; index < numberNum; index++ {
			if number[index] >= '0' && number[index] <= '9' {
				result.WriteByte(number[index])
				counter++
				remainNum--
			}
		}
	}
	return result.String()
}
