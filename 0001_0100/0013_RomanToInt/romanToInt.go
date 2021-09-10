package romanToInt

// my method 逆向读string, 不够简洁
// 8ms 3MB
func romanToInt(s string) int {
	var result int
	num := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	strlen := len(s)
	flag := true
	for i := strlen - 1; i >= 1; i-- {
		if num[s[i-1]] < num[s[i]] {
			result += num[s[i]] - num[s[i-1]]
			if i-1 == 0 {
				flag = false
				break
			}
			i = i - 1
		} else {
			result += num[s[i]]
		}
	}
	if flag {
		result += num[s[0]]
	}
	return result
}

// // others' method 正向读string，比较简洁
// // 8ms 3MB
// func romanToInt(s string) int {
// 	var result int
// 	num := map[byte]int{
// 		'I': 1,
// 		'V': 5,
// 		'X': 10,
// 		'L': 50,
// 		'C': 100,
// 		'D': 500,
// 		'M': 1000,
// 	}
// 	strlen := len(s)
// 	s = s + "I"
// 	for i := 0; i < strlen; i++ {
// 		if num[s[i]] < num[s[i+1]] {
// 			result -= num[s[i]]
// 		} else {
// 			result += num[s[i]]
// 		}
// 	}
// 	return result
// }
