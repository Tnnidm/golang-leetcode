package calPoints

import "strconv"

func calPoints(ops []string) int {
	score := make([]int, len(ops))
	scoreLen := 0
	for _, v := range ops {
		switch v {
		case "+":
			score[scoreLen] = score[scoreLen-1] + score[scoreLen-2]
			scoreLen++
		case "D":
			score[scoreLen] = 2 * score[scoreLen-1]
			scoreLen++
		case "C":
			scoreLen--
		default:
			value, _ := strconv.Atoi(v)
			score[scoreLen] = value
			scoreLen++
		}
	}
	sum := 0
	for i := 0; i < scoreLen; i++ {
		sum += score[i]
	}
	return sum
}
