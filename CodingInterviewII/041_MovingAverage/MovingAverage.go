package movingAverage

type MovingAverage struct {
	volume []int
	count  int
	sum    int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{make([]int, size), 0, 0}
}

func (ma *MovingAverage) Next(val int) float64 {
	if ma.count < len(ma.volume) {
		ma.volume[ma.count] = val
		ma.sum += val
		ma.count++
	} else {
		ma.sum = ma.sum - ma.volume[0] + val
		ma.volume = ma.volume[1:]
		ma.volume = append(ma.volume, val)
	}
	return float64(ma.sum) / float64(ma.count)
}
