package findAnagrams

func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return []int{}
	}
	var pMap [26]int
	for i := range p {
		pMap[p[i]-'a']++
	}
	var sMap [26]int
	for i := 0; i < pLen; i++ {
		sMap[s[i]-'a']++
	}
	result := []int{}
	if sMap == pMap {
		result = append(result, 0)
	}
	for i := 1; i <= sLen-pLen; i++ {
		sMap[s[i-1]-'a']--
		sMap[s[i+pLen-1]-'a']++
		if sMap == pMap {
			result = append(result, i)
		}
	}
	return result
}
