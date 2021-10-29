package reverseVowels

func reverseVowels(s string) string {
	charMap := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	sArray := []byte(s)
	for pos, neg := 0, len(s)-1; pos < neg; {
		for {
			if _, ok := charMap[sArray[pos]]; !ok && pos < len(s)-1 {
				pos++
			} else {
				break
			}
		}
		for {
			if _, ok := charMap[sArray[neg]]; !ok && neg > 0 {
				neg--
			} else {
				break
			}
		}
		if pos < neg {
			sArray[pos], sArray[neg] = sArray[neg], sArray[pos]
			pos++
			neg--
		} else {
			break
		}

	}
	return string(sArray)
}

// // 有更简洁的写法，调用strings.Contains
// func reverseVowels(s string) string {
// 	t := []byte(s)
// 	n := len(t)
// 	i, j := 0, n-1
// 	for i < j {
// 		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
// 			i++
// 		}
// 		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
// 			j--
// 		}
// 		if i < j {
// 			t[i], t[j] = t[j], t[i]
// 			i++
// 			j--
// 		}
// 	}
// 	return string(t)
// }
