package findContentChildren

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] > s[j] {
			j++
		} else {
			i++
			j++
			count++
		}
	}
	return count
}
