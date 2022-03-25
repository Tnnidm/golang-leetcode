package letterCombinations

func letterCombinations(digits string) []string {
	letterLen := len(digits)
	if letterLen == 0 {
		return []string{}
	}
	letters := []byte(digits)
	ansLen := 1
	for i := 0; i < letterLen; i++ {
		if letters[i] == '7' || letters[i] == '9' {
			ansLen *= 4
		} else {
			ansLen *= 3
		}
	}
	ans := make([]string, ansLen)
	ansIndex := 0
	letterMap := map[byte]([]byte){
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	dfs(&letters, &letterMap, 0, letterLen, &ans, &ansIndex)
	return ans
}

func dfs(letters *[]byte, letterMap *map[byte]([]byte), index int, length int, ans *[]string, ansIndex *int) {
	if index == length {
		(*ans)[*ansIndex] = string(*letters)
		*ansIndex++
		return
	}
	temp := (*letters)[index]
	for _, char := range (*letterMap)[(*letters)[index]] {
		(*letters)[index] = char
		dfs(letters, letterMap, index+1, length, ans, ansIndex)
		(*letters)[index] = temp
	}
}
