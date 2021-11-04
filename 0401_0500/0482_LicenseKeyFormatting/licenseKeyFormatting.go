package licenseKeyFormatting

import "strings"

func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(s)
	result := []byte{}
	sLen := len(s)
	counter := 0
	for i := sLen - 1; i >= 0; i-- {
		if s[i] != '-' {
			result = append(result, s[i])
			counter++
			if counter == k && i != 0 {
				result = append(result, '-')
				counter = 0
			}
		}
	}
	resultLen := len(result)
	if resultLen == 0 {
		return ""
	} else {
		if result[resultLen-1] == '-' {
			result = result[:resultLen-1]
			resultLen -= 1
		}
		for i := 0; i < resultLen/2; i++ {
			result[i], result[resultLen-1-i] = result[resultLen-1-i], result[i]
		}
		return string(result)
	}
}
