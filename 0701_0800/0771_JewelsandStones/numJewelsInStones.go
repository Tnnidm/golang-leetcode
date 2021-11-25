package numJewelsInStones

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	counter := 0
	for _, v := range stones {
		if strings.ContainsRune(jewels, v) {
			counter++
		}
	}
	return counter
}
