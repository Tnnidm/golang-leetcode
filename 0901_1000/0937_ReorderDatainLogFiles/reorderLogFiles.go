package reorderLogFiles

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		split1 := strings.SplitN(logs[i], " ", 2)
		split2 := strings.SplitN(logs[j], " ", 2)
		isDigit1 := unicode.IsDigit(rune(split1[1][0]))
		isDigit2 := unicode.IsDigit(rune(split2[1][0]))
		if !isDigit1 && !isDigit2 {
			return split1[1] < split2[1] || (split1[1] == split2[1] && split1[0] < split2[0])
		}
		return !isDigit1 && isDigit2
	})
	return logs
}
