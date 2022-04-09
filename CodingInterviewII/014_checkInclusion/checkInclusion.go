package checkInclusion

func checkInclusion(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s2Len < s1Len {
		return false
	}
	stat := make([]int, 26)
	for i := 0; i < s1Len; i++ {
		stat[s1[i]-'a']++
		stat[s2[i]-'a']--
	}
	if isAllZero(stat) {
		return true
	}
	for i := s1Len; i < s2Len; i++ {
		stat[s2[i]-'a']--
		stat[s2[i-s1Len]-'a']++
		if isAllZero(stat) {
			return true
		}
	}
	return false
}

func isAllZero(arr []int) bool {
	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}
