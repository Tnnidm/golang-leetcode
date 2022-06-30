package simplifyPath

import (
	"strings"
)

func simplifyPath(path string) string {
	result := []string{}
	for _, section := range strings.Split(path, "/") {
		if section == ".." {
			resultLen := len(result)
			if resultLen > 0 {
				result = result[:resultLen-1]
			}
		} else if section != "." && section != "" {
			result = append(result, section)
		}
	}
	return "/" + strings.Join(result, "/")
}
