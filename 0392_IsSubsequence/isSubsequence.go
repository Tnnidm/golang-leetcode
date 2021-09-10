package isSubsequence

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了70.49%的用户
func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen == 0 {
		return true
	} else {
		tIndex := 0
		for i := 0; i < sLen; i++ {
			for ; tIndex < tLen; tIndex++ {
				if s[i] == t[tIndex] {
					if i == sLen-1 {
						return true
					} else {
						tIndex += 1
						break
					}

				}
			}
		}
		return false
	}
}
