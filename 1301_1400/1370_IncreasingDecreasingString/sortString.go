package sortString

import "strings"

func sortString(s string) string {
	bucket := make([]int, 26)
	for _, v := range s {
		bucket[v-'a']++
	}
	result := strings.Builder{}
	for result.Len() != len(s) {
		for i := 0; i < 26; i++ {
			if bucket[i] > 0 {
				result.WriteByte(uint8(i) + 'a')
				bucket[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if bucket[i] > 0 {
				result.WriteByte(uint8(i) + 'a')
				bucket[i]--
			}
		}
	}
	return result.String()
}
