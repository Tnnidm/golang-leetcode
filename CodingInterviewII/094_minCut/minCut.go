package minCut

// // 非常省内存但是时间复杂度高
// // O(n^3)时间复杂度和O(1)空间复杂度
// func minCut(s string) int {
// 	sLen := len(s)
// 	endWithIndex := make([]int, sLen)
// 	for i := 0; i < sLen; i++ {
// 		if isPalindrome(&s, 0, i) {
// 			endWithIndex[i] = 0
// 		} else {
// 			endWithIndex[i] = i
// 		}
// 	}
// 	for i := 1; i < sLen; i++ {
// 		for j := i + 1; j < sLen; j++ {
// 			if isPalindrome(&s, i, j) {
// 				endWithIndex[j] = min(endWithIndex[i]+1, endWithIndex[j])
// 			}
// 		}
// 	}
// 	return endWithIndex[sLen-1]
// }

// func isPalindrome(s *string, start int, end int) bool {
// 	for start < end {
// 		if (*s)[start] != (*s)[end] {
// 			return false
// 		}
// 		start++
// 		end--
// 	}
// 	return true
// }

// func min(x int, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// 换一种空间换时间的方法，先把哪些是回文记下来
// O(n^2)时间复杂度和O(n^2)空间复杂度
func minCut(s string) int {
	sLen := len(s)
	palindromeTwoHead := make([][]int, sLen)
	// palindrome with odd num
	for center := 1; center < sLen-1; center++ {
		for radius := 1; center-radius >= 0 && center+radius < sLen; radius++ {
			if s[center-radius] == s[center+radius] {
				palindromeTwoHead[center-radius] = append(palindromeTwoHead[center-radius], center+radius)
			} else {
				break
			}
		}
	}
	// palindrome with even num
	for center := 0; center < sLen-1; center++ {
		for radius := 0; center-radius >= 0 && center+1+radius < sLen; radius++ {
			if s[center-radius] == s[center+radius+1] {
				palindromeTwoHead[center-radius] = append(palindromeTwoHead[center-radius], center+1+radius)
			} else {
				break
			}
		}
	}
	endWithIndex := make([]int, sLen)
	for i := 0; i < sLen; i++ {
		endWithIndex[i] = i
	}
	for _, index := range palindromeTwoHead[0] {
		endWithIndex[index] = 0
	}
	for i := 1; i < sLen; i++ {
		endWithIndex[i] = min(endWithIndex[i], endWithIndex[i-1]+1)
		for _, index := range palindromeTwoHead[i] {
			endWithIndex[index] = min(endWithIndex[index], endWithIndex[i-1]+1)
		}
	}
	return endWithIndex[sLen-1]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
