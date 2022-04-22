package countAndSay

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	result := "1"
	for ; n > 1; n-- {
		result = countAndSaycore(result)
	}
	return result
}

func countAndSaycore(input string) string {
	count := 0
	prev := input[0]
	result := strings.Builder{}
	for i := 0; i < len(input); i++ {
		if input[i] != prev {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(prev)
			count = 1
			prev = input[i]
		} else {
			count++
		}
	}
	result.WriteString(strconv.Itoa(count))
	result.WriteByte(prev)
	return result.String()
}
