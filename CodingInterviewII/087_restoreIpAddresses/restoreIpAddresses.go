package restoreIpAddresses

import "strings"

func restoreIpAddresses(s string) []string {
	sLen := len(s)
	if sLen < 4 || sLen > 12 {
		return []string{}
	}
	result := []string{}
	combination := []string{}
	var dfs func(index int, seg int)
	dfs = func(index int, seg int) {
		if index == sLen || seg == 4 {
			if index == sLen && seg == 4 {
				result = append(result, strings.Join(combination, "."))
			}
			return
		}
		if s[index] == '0' {
			combination = append(combination, "0")
			dfs(index+1, seg+1)
			combination = combination[:seg]
		} else {
			for i, number := index, 0; i < sLen; i++ {
				number = 10*number + int(s[i]-'0')
				if number > 255 {
					break
				}
				combination = append(combination, s[index:i+1])
				dfs(i+1, seg+1)
				combination = combination[:seg]
			}
		}
	}
	dfs(0, 0)
	return result
}
