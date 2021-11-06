package convertToBase7

import "strconv"

// func convertToBase7(num int) string {
// 	var result string
// 	if num < 0 {
// 		result = "-"
// 		num = -num
// 	} else if num == 0 {
// 		return "0"
// 	}
// 	temp := []string{}
// 	for num != 0 {
// 		temp = append(temp, strconv.Itoa(num%7))
// 		num /= 7
// 	}
// 	for i := len(temp) - 1; i >= 0; i-- {
// 		result += temp[i]
// 	}
// 	return result
// }

// 这道题可以直接一行秒杀的
func convertToBase7(num int) string {
	res := strconv.FormatInt(int64(num), 7)
	return string(res)
}
