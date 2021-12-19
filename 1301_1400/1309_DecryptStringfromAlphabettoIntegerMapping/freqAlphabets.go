package freqAlphabets

import (
	"strconv"
	"strings"
)

func freqAlphabets(s string) string {
	result := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			decode, _ := strconv.Atoi(s[i-2 : i])
			result = append(result, 'a'+uint8(decode-1))
			i = i - 2
		} else {
			decode, _ := strconv.Atoi(s[i : i+1])
			result = append(result, 'a'+uint8(decode-1))
		}
	}
	ans := strings.Builder{}
	for i := len(result) - 1; i >= 0; i-- {
		ans.WriteByte(result[i])
	}
	return ans.String()
}
