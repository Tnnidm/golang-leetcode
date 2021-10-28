package firstBadVersion

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	if n == 1 {
		return n
	}
	begin := 1
	end := n
	mid := 0
	for {
		mid = (begin + end) / 2
		if mid == begin {
			if isBadVersion(mid) {
				return mid
			} else {
				return end
			}
		} else {
			if isBadVersion(mid) {
				end = mid
			} else {
				begin = mid
			}
		}

	}
}
