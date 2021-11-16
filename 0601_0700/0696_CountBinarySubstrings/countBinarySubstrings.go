package countBinarySubstrings

func countBinarySubstrings(s string) int {
	counter := []int{}
	ptr, sLen := 0, len(s)
	for ptr < sLen {
		c := s[ptr]
		count := 0
		for ptr < sLen && s[ptr] == c {
			ptr++
			count++
		}
		counter = append(counter, count)
	}
	result := 0
	for i := 1; i < len(counter); i++ {
		result += min(counter[i-1], counter[i])
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// // 更省内存的做法
// func countBinarySubstrings(s string) int {
// 	last, curr, res := 0, 1, 0
// 	for i := 1; i < len(s); i++ {
// 		if s[i] == s[i-1] {
// 			curr++
// 		} else {
// 			last, curr = curr, 1
// 		}
// 		if last >= curr {
// 			res++
// 		}
// 	}
// 	return res
// }
