package slowestKey

func slowestKey(releaseTimes []int, keysPressed string) byte {
	key := keysPressed[0]
	pressTime := releaseTimes[0]
	for i := 1; i < len(keysPressed); i++ {
		if releaseTimes[i]-releaseTimes[i-1] > pressTime {
			pressTime = releaseTimes[i] - releaseTimes[i-1]
			key = keysPressed[i]
		} else if releaseTimes[i]-releaseTimes[i-1] == pressTime {
			if keysPressed[i] > key {
				key = keysPressed[i]
			}
		}
	}
	return key
}
