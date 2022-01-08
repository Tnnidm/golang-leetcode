package areAlmostEqual

func areAlmostEqual(s1 string, s2 string) bool {
	difference := make([]int, 0, 2)
	for i := range s1 {
		if s1[i] != s2[i] {
			difference = append(difference, i)
		}
		if len(difference) > 2 {
			return false
		}
	}
	if len(difference) == 2 {
		return s1[difference[0]] == s2[difference[1]] && s1[difference[1]] == s2[difference[0]]
	}
	return len(difference) == 0
}
