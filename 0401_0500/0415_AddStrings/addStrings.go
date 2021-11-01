package addStrings

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	result := ""
	sum := 0
	add := 0
	digit1, digit2 := 0, 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		digit1, digit2 = 0, 0
		if i >= 0 {
			digit1 = int(num1[i] - '0')
		}
		if j >= 0 {
			digit2 = int(num2[j] - '0')
		}
		sum = digit1 + digit2 + add
		result = strconv.Itoa(sum%10) + result
		add = sum / 10
	}
	return result
}
