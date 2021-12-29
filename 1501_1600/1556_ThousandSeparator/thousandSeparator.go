package thousandSeparator

import (
	"strconv"
	"strings"
)

func thousandSeparator(n int) string {
	nString := strconv.Itoa(n)
	count := len(nString) % 3
	if count == 0 {
		count = 3
	}
	result := strings.Builder{}
	for i := 0; i < len(nString); i++ {
		result.WriteByte(nString[i])
		count--
		if count == 0 && i != len(nString)-1 {
			result.WriteByte('.')
			count = 3
		}
	}
	return result.String()
}
