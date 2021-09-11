package lengthOfLastWord

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了70.08%的用户
func lengthOfLastWord(s string) int {
	sLen := len(s)
	counter := 0
	for i := sLen - 1; i >= 0; i-- {
		if s[i] != ' ' {
			counter += 1
		} else if counter != 0 {
			break
		}
	}
	return counter
}
