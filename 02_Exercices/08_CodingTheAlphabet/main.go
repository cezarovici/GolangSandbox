package main

import (
	"math"
)

// If a = 1, b = 2, c = 3, ..., z = 26.
// Given a string find all possible alphabet codes that the string can generate.
// Provide the count and the strings.

// Input:
// example: "1123"

// Output:

// "aabc" - a = 1, a = 1, b = 2, c = 3
// "kbc"  - k = 11, b = 2, c = 3
// "alc"  - a = 1, l = 12, c = 3
// "aaw"  - a = 1, a = 1, w = 23
// "kw"   - k = 11, w = 23

func fromAsciiToIndex(character rune) int {
	return int(character) - 96
}

func convertStringToIndex(input string) int {
	var res float64
	var base float64
	var power int

	base = 10

	for i := len(input) - 1; i >= 0; i-- {
		power = i
		value := fromAsciiToIndex(rune(input[len(input)-i-1]))
		if value > 9 {
			res *= 10
		}
		res = res + math.Pow(base, float64(power))*float64(value)
	}

	return int(res)
}

func main() {
}
