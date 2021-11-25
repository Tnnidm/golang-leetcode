package main

import "strings"

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	charCounter := make(map[byte]int)
	for i := range licensePlate {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			charCounter[licensePlate[i]-'a']++
		}
	}
	// fmt.Println(licensePlate, charCounter)
	result := "1234512345123451"
	for _, word := range words {
		if len(word) >= len(result) {
			continue
		}
		completingWordFlag := true
		charCounterWord := make(map[byte]int)
		for i := range word {
			charCounterWord[word[i]-'a']++
		}
		// fmt.Println(word, charCounterWord)
		for k, v := range charCounter {
			if charCounterWord[k]-v < 0 {
				completingWordFlag = false
				break
			}
		}
		if completingWordFlag {
			if len(word) < len(result) {
				result = word
			}
		}
	}
	return result
}
