package countBalls

func countBalls(lowLimit int, highLimit int) int {
	box := make([]int, 46)
	for i := lowLimit; i <= highLimit; i++ {
		box[boxIndex(i)]++
	}
	maxNum := 0
	for i := 0; i < 46; i++ {
		if box[i] > maxNum {
			maxNum = box[i]
		}
	}
	return maxNum
}

func boxIndex(ballIndex int) int {
	boxindex := 0
	for ballIndex != 0 {
		boxindex = boxindex + ballIndex%10
		ballIndex = ballIndex / 10
	}
	return boxindex
}
