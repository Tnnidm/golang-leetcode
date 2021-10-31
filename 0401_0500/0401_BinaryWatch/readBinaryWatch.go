package readBinaryWatch

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	result := []string{}
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return result
}
