package average

func average(salary []int) float64 {
	sum, min, max := 0, salary[0], salary[0]
	workerNum := len(salary)
	for i := 0; i < workerNum; i++ {
		sum = sum + salary[i]
		if salary[i] < min {
			min = salary[i]
		}
		if salary[i] > max {
			max = salary[i]
		}
	}
	return float64(sum-min-max) / float64(workerNum-2)
}
