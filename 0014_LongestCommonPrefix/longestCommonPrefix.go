package longestCommonPrefix

// My method
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00 % 的用户
// 内存消耗：2.4 MB, 在所有 Go 提交中击败了 10.46 % 的用户
func longestCommonPrefix(strs []string) string {
	var common_prefix string
	str_num := len(strs)
	if str_num == 0 {
		common_prefix = ""
	} else if str_num == 1 {
		common_prefix = strs[0]
	} else {
		for i := range strs {
			strs[i] = strs[i] + "X"
		}
		var prefix byte
		var flag bool
		for string_index := 0; ; string_index++ {
			flag = true
			prefix = strs[0][string_index]
			for i := 0; i < str_num; i++ {
				if strs[i][string_index] != prefix || strs[i][string_index] == 'X' {
					flag = false
					break
				}
			}
			if flag {
				common_prefix = common_prefix + string(prefix)
			} else {
				break
			}
		}
	}
	return common_prefix
}

// // 如何优化内存呢？
// // 执行用时：0 ms, 在所有 Go 提交中击败了 100.00 % 的用户
// // 内存消耗：2.3 MB, 在所有 Go 提交中击败了 58.66 % 的用户
// func longestCommonPrefix(strs []string) string {
// 	// find smallest
// 	min := ""
// 	limit := 200
// 	for _, str := range strs {
// 		if len(str) < limit {
// 			min = str
// 			limit = len(str)
// 		}
// 	}

// 	// compare
// 	equalMinLen := 200
// 	for _, str := range strs {
// 		equalLen := 0
// 		for i := range min {
// 			if str[i] != min[i] {
// 				break
// 			} else {
// 				equalLen = i + 1
// 			}
// 		}
// 		if equalLen < equalMinLen {
// 			equalMinLen = equalLen
// 		}
// 	}

// 	return min[:equalMinLen]
// }
