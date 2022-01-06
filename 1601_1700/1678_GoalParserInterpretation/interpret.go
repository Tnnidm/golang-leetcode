package interpret

import "strings"

func interpret(command string) string {
	i := 0
	result := strings.Builder{}
	for i < len(command) {
		if command[i:i+1] == "G" {
			result.WriteString("G")
			i = i + 1
		} else if command[i:i+2] == "()" {
			result.WriteString("o")
			i = i + 2
		} else {
			result.WriteString("al")
			i = i + 4
		}
	}
	return result.String()
}
