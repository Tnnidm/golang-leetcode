package recentCounter

type RecentCounter struct {
	recorder []int
}

func Constructor() RecentCounter {
	rc := RecentCounter{}
	return rc
}

func (this *RecentCounter) Ping(t int) int {
	this.recorder = append(this.recorder, t)
	for i := 0; i < len(this.recorder); i++ {
		if this.recorder[i] >= t-3000 {
			this.recorder = this.recorder[i:]
			break
		}
	}
	return len(this.recorder)
}
