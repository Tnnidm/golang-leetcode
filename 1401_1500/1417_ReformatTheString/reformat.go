package reformat

func reformat(s string) string {
	// 计算两种类型各自总数  数字、字母
	var total0, total1 int
	var counter [36]int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' { // 26-25 数字
			counter[s[i]-'0'+26]++
			total0++
		} else {
			counter[s[i]-'a']++
			total1++
		}
	}
	diff := total0 - total1
	if diff > 1 || diff < -1 {
		return ""
	}
	bs := make([]uint8, len(s))
	var k int
	// 如果数字个数多，则第一个应该是数字，否则是字母
	if diff >= 0 { // 数字第一个
		k = 0
		for i := 26; i < 36; i++ { // 偶数位存放数字
			for counter[i] > 0 {
				bs[k] = uint8(i) - 26 + '0'
				k = k + 2
				counter[i]--
			}
		}
		k = 1
		for i := 0; i < 26; i++ { // 奇数位存放字母
			for counter[i] > 0 {
				bs[k] = uint8(i) + 'a'
				k = k + 2
				counter[i]--
			}
		}
	} else { // 字母第一个
		k = 0
		for i := 0; i < 26; i++ { // 奇数位存放字母
			for counter[i] > 0 {
				bs[k] = uint8(i) + 'a'
				k = k + 2
				counter[i]--
			}
		}
		k = 1
		for i := 26; i < 36; i++ { // 偶数位存放数字
			for counter[i] > 0 {
				bs[k] = uint8(i-26) + '0'
				k = k + 2
				counter[i]--
			}
		}
	}
	return string(bs)
}
