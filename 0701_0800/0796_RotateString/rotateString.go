package rotateString

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	s = strings.Join([]string{s, s}, "")
	return strings.Contains(s, goal)
}
