package printNumbers

func convertToString(digits []byte, n int) string {
	index := n - 1
	for i := 0; i < n; i++ {
		if digits[i] != '0' {
			index = i
		}
	}
	return string(digits[index:])
}

// // 迭代方法
// func printNumbers(n int) []string {
// 	result := []string{}
// 	digits := make([]byte, n)
// 	for i := 0; i < n; i++ {
// 		digits[i] = '0'
// 	}
// 	var forward byte
// 	for !isHighest(digits, n) {
// 		for i := n - 1; i >= 0; i-- {
// 			if i == n-1 {
// 				digits[i] = digits[i] + 1
// 				forward = 0
// 			} else {
// 				digits[i] = digits[i] + forward
// 			}
// 			if digits[i] > '9' {
// 				digits[i] = '0'
// 				forward = 1
// 			}
// 		}
// 		result = append(result, convertToString(digits, n))
// 	}
// 	return result
// }

// func isHighest(digits []byte, n int) bool {
// 	for i := 0; i < n; i++ {
// 		if digits[i] != '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

// 递归方法
func printNumbers(n int) []string {
	result := []string{}
	if n < 0 {
		return result
	}
	digits := make([]byte, n)
	addNumRecursively(&result, &digits, n, 0)
	return result[1:]
}

func addNumRecursively(result *[]string, digits *[]byte, n int, index int) {
	if index == n {
		*result = append(*result, convertToString(*digits, n))
	} else {
		for i := 0; i < 10; i++ {
			(*digits)[index] = '0' + uint8(i)
			addNumRecursively(result, digits, n, index+1)
		}
	}
}
