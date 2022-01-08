package countMatches

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var k int
	switch ruleKey {
	case "type":
		k = 0
	case "color":
		k = 1
	case "name":
		k = 2
	}
	var ans int
	for _, item := range items {
		if item[k] == ruleValue {
			ans++
		}
	}
	return ans
}
