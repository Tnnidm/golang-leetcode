package recentCounter

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	target := t - 3000
	this.queue = append(this.queue, t)

	left, right := 0, len(this.queue)-1
	for left < right {
		mid := (left + right) / 2
		if this.queue[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	this.queue = this.queue[left:]
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
