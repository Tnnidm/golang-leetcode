package strStr

// // My method 1: brute force
// // 执行用时：212 ms, 在所有 Go 提交中击败了23.48%的用户
// // 内存消耗：2.6 MB, 在所有 Go 提交中击败了37.06%的用户
// func strStr(haystack string, needle string) int {
// 	haystackLen := len(haystack)
// 	needleLen := len(needle)
// 	var result int = -1
// 	var i int
// 	var j int
// 	if needleLen > haystackLen {
// 		result = -1
// 	} else if needleLen == 0 {
// 		result = 0
// 	} else {
// 		for i = 0; i <= haystackLen-needleLen; i++ {
// 			for j = 0; j < needleLen; j++ {
// 				if needle[j] != haystack[i+j] {
// 					break
// 				}
// 			}
// 			if j == needleLen {
// 				result = i
// 				break
// 			}
// 		}
// 	}
// 	return result
// }

// // Other's method: use package "strings"
// // 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// // 内存消耗：2.5 MB, 在所有 Go 提交中击败了73.36%的用户
// func strStr(haystack string, needle string) int {
// 	if needle == "" {
// 		return 0
// 	}
// 	return strings.Index(haystack, needle)
// }

// My method 2: KMP algorithm
// 执行用时：280 ms, 在所有 Go 提交中击败了17.83%的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了6.41%的用户
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen > haystackLen {
		return -1
	} else if needleLen == 0 {
		return 0
	} else {
		dismatch_table := make([]int, needleLen, needleLen)
		var index int
		var index2 int
		index2 = -1
		for index = 0; index < needleLen; index++ {
			for index2 > 0 {
				if needle[index-index2:index] == needle[0:index2] {
					break
				} else {
					index2 -= 1
				}
			}
			dismatch_table[index] = index2
			index2 += 1
		}
		for index = 0; index <= haystackLen-needleLen; {
			for index2 = 0; index2 < needleLen; index2++ {
				if needle[index2] != haystack[index+index2] {
					index = index + index2 - dismatch_table[index2]
					break
				}
			}
			if index2 == needleLen {
				return index
			}
		}
	}
	return -1
}
