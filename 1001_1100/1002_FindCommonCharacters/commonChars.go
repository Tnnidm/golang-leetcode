package commonChars

import "math"

func commonChars(words []string) []string {
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range words {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	ans := []string{}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return ans
}
