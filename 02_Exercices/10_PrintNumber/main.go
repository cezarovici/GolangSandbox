package main

import "fmt"

func reverse(number int) int {
	var res int

	for number != 0 {
		res = res*10 + number%10

		number = number / 10
	}

	return res
}

func numberOfDigits(number int) int {
	var count int

	for number != 0 {
		count++

		number = number / 10
	}

	return count
}

func getRank(index int) string {
	ranks := []string{" ", " thousands ", " milions ", " bilions "}

	return ranks[index]
}

// 123 456 789
// 981 654 321

func printNumberAsString(number int) {
	number = reverse(number)

	for number != 0 {
		lenght := numberOfDigits(number)

		fmt.Print(reverse(number%1000), getRank(lenght/4))

		number = number / 1000
	}
}

func main() {
	printNumberAsString(123456789)
}
