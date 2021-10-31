package toHex

import (
	"strings"
)

// func toHex(num int) string {
// 	if num < 0 {
// 		num += 0x100000000
// 	} else if num == 0 {
// 		return "0"
// 	}
// 	result := []byte{}
// 	hexMap := map[int]byte{
// 		0:  '0',
// 		1:  '1',
// 		2:  '2',
// 		3:  '3',
// 		4:  '4',
// 		5:  '5',
// 		6:  '6',
// 		7:  '7',
// 		8:  '8',
// 		9:  '9',
// 		10: 'a',
// 		11: 'b',
// 		12: 'c',
// 		13: 'd',
// 		14: 'e',
// 		15: 'f',
// 	}
// 	for num != 0 {
// 		result = append([]byte{hexMap[num%16]}, result...)
// 		num /= 16
// 	}
// 	return string(result)
// }

// 使用位运算+strings.Builder{}运算可以消耗更少的内存
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	result := &strings.Builder{}
	for i := 7; i >= 0; i-- {

		val := num >> (4 * i) & 0xf
		if val > 0 || result.Len() > 0 {
			var digit byte
			if val < 10 {
				digit = '0' + byte(val)
			} else {
				digit = 'a' + byte(val-10)
			}
			result.WriteByte(digit)
		}
	}
	// return result.String()
	return "0"
}
