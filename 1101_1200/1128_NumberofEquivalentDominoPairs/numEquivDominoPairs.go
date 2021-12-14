package numEquivDominoPairs

func numEquivDominoPairs(dominoes [][]int) (ans int) {
	cnt := [90]int{}
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1] - 10
		ans += cnt[v]
		cnt[v]++
	}
	return
}
