package largestPalindromic

import (
	"strconv"
	"strings"
)

func largestPalindromic(num string) string {
	count := make([]int, 10)
	for i := range num {
		count[num[i]-'0']++
	}
	result := strings.Builder{}
	for i := 9; i >= 1; i-- {
		for ; count[i] >= 2; count[i] -= 2 {
			result.WriteByte('0' + uint8(i))
		}
	}
	if result.Len() > 0 {
		for ; count[0] >= 2; count[0] -= 2 {
			result.WriteByte('0')
		}
	}
	half := result.String()
	reverse := []byte(half)
	halfLen := len(reverse)
	for i := 0; i < halfLen>>1; i++ {
		reverse[i], reverse[halfLen-1-i] = reverse[halfLen-1-i], reverse[i]
	}

	for i := 9; i >= 0; i-- {
		if count[i] > 0 {
			return half + strconv.Itoa(i) + string(reverse)
		}
	}
	return half + string(reverse)
}
