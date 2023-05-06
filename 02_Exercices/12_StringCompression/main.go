package main

import (
	"fmt"
)

func stringCompression(content string) string {
	var res string
	var lastLetter rune
	var lastFreq int

	for i, letter := range content {
		if i == 0 {
			lastFreq = 0
			lastLetter = letter

			continue
		}

		if letter != lastLetter {
			res = res + string(lastLetter) + fmt.Sprint(i-lastFreq)

			lastLetter = letter
			lastFreq = i
		}

		if i == len(content)-1 {
			res = res + string(letter) + fmt.Sprint(i-lastFreq+1)
		}
	}

	if len(res) > len(content) {
		return content
	}

	return res
}

func main() {
	fmt.Print(stringCompression("aabcccccaaa"))
}
