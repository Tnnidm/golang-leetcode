package intToRoman

import (
	"strings"
)

func intToRoman(num int) string {
	transformMap := map[int]byte{
		1000: 'M',
		500:  'D',
		100:  'C',
		50:   'L',
		10:   'X',
		5:    'V',
		1:    'I',
	}
	result := strings.Builder{}
	for base := 1000; base != 0; base /= 10 {
		baseNum := num / base
		if baseNum <= 3 {
			for i := 0; i < baseNum; i++ {
				result.WriteByte(transformMap[base])
			}
		} else if baseNum == 4 {
			result.WriteByte(transformMap[base])
			result.WriteByte(transformMap[5*base])
		} else if baseNum >= 5 && baseNum < 9 {
			result.WriteByte(transformMap[5*base])
			for i := 5; i < baseNum; i++ {
				result.WriteByte(transformMap[base])
			}
		} else {
			result.WriteByte(transformMap[base])
			result.WriteByte(transformMap[10*base])
		}
		num = num % base
	}
	return result.String()
}
