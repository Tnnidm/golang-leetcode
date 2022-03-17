package lengthOfLongestSubstring

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// // 第一种方法是动态规划加哈希表
// // 时间复杂度是O(n)，空间复杂度是O(1)
// func lengthOfLongestSubstring(s string) int {
// 	sLen := len(s)
// 	if sLen <= 1 {
// 		return sLen
// 	}
// 	length := 0
// 	maxlength := 0
// 	lastIndex := make(map[byte]int, min(sLen, 128))
// 	for i := 0; i < sLen; i++ {
// 		if _, ok := lastIndex[s[i]]; ok && i-lastIndex[s[i]] <= length {
// 			length = i - lastIndex[s[i]]
// 		} else {
// 			length++
// 		}
// 		lastIndex[s[i]] = i
// 		if length > maxlength {
// 			maxlength = length
// 		}
// 	}
// 	return maxlength
// }

// 第二种方法是动态规划加数组
// 时间复杂度是O(n)，空间复杂度是O(1)
// 但是性能好多了
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return sLen
	}
	lastIndex := make([]int, 128)
	left := 0
	maxlength := 0
	for i := 0; i < len(s); i++ {
		left = max(lastIndex[s[i]], left)
		maxlength = max(maxlength, i-left+1)
		lastIndex[s[i]] = i + 1
	}
	return maxlength
}
