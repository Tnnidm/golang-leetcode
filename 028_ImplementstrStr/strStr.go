package strStr

// My method 1: brute force
// 执行用时：212 ms, 在所有 Go 提交中击败了23.48%的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了37.06%的用户
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	var result int = -1
	var i int
	var j int
	if needleLen > haystackLen {
		result = -1
	} else if needleLen == 0 {
		result = 0
	} else {
		for i = 0; i <= haystackLen-needleLen; i++ {
			for j = 0; j < needleLen; j++ {
				if needle[j] != haystack[i+j] {
					break
				}
			}
			if j == needleLen {
				result = i
				break
			}
		}
	}
	return result
}

// // Other's method: use package "strings"
// // 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// // 内存消耗：2.5 MB, 在所有 Go 提交中击败了73.36%的用户
// func strStr(haystack string, needle string) int {
// 	if needle == "" {
// 		return 0
// 	}
// 	return strings.Index(haystack, needle)
// }

// // My method 2: KMP algorithm
// func strStr(haystack string, needle string) int {

// }
