package minNumberOfHours

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	// N := len(energy)
	needEnergy, remainExperience := 1, initialExperience
	result := 0
	for i := range energy {
		needEnergy += energy[i]
		if remainExperience <= experience[i] {
			diff := experience[i] + 1 - remainExperience
			remainExperience += diff
			result += diff
		}
		remainExperience += experience[i]
	}
	return max(needEnergy-initialEnergy, 0) + result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
