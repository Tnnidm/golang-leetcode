package distributeCandies

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	for i := 0; candies > 0; i++ {
		ans[i%num_people] += min(i+1, candies)
		candies -= i + 1
	}
	return ans
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
